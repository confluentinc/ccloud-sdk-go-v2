// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Confluent Cloud Logs API

# Introduction  The Logs API is served by `cc-logging-service` and returns customer-facing operational log documents for one Confluent Cloud resource, filtered by time window and by the search fields declared for that resource type. Its consumers today are the Confluent Cloud Console log viewer and the Confluent CLI (`pkg/ccloudv2/logging.go`). This spec documents the API exactly as the service serves it on `master` as of 2026-10-15; it is the source for the generated Go client.  # Authentication  Requests carry a Confluent Cloud data-plane bearer token obtained from a user session (`Authorization: Bearer <token>`). Cloud API keys are not accepted. The organization in the request's `crn` must match the organization of the token's principal, otherwise the request is rejected with `422`.  Access is gated per organization by the `logging.external.api.enable` LaunchDarkly flag. An organization without the flag receives `403` with the detail `Feature is not enabled for this organization`.  # Resource types  The service resolves the trailing element of `crn` against its resource catalog, which is deployment configuration (`resourceCatalog` in `.halyard/cc-logging-service.yaml`). The catalog decides which OpenSearch indices are queried, which `search` keys are accepted, which fields a document carries in the response, and which RBAC operation is checked. As of 2026-10-15 the production catalog serves these resource types:  | CRN shape (trailing element) | Accepted `search` keys | Allowed values | Returned fields | RBAC operation | |------------------------------|------------------------|----------------|-----------------|----------------| | `organization/environment/cloud-cluster/connector` | `search_text`, `level`, `task_id`, `id` | `level`: `INFO`, `WARN`, `ERROR`, `FATAL` | `timestamp`, `level`, `task_id`, `message`, `exception.stacktrace`, `id`, `name` | `ReadLogs` | | `organization/environment/flink-region/statement` | `search_text`, `level`, `source`, `id` | `level`: `INFO`, `WARN`, `ERROR` | `id`, `timestamp`, `level`, `message`, `crn`, `data`, `exception`, `function_name`, `function_crn`, `source` | `Describe` | | `organization/environment/flink-region/cloud-cluster/materialized-table` | `search_text`, `level`, `source`, `id` | `level`: `INFO`, `WARN`, `ERROR`; `source`: `MaterializedTable`, `Function` | `id`, `timestamp`, `level`, `message`, `crn`, `data`, `source`, `materialized_table_name`, `materialized_table_version` | `Describe` | | `organization/environment/cloud-cluster/kafka` | `search_text`, `event_type`, `level`, `id` | `level`: `INFO`, `WARN`, `ERROR`; `event_type`: `SERVICE_MAINTENANCE`, `HEALTH_ALERT`, `CONFLUENT_STATUS_UPDATED`, `SERVICE_ISSUE_DETECTED` | `id`, `timestamp`, `level`, `message`, `crn`, `source`, `event_type`, `event_id`, `status`, `cloud_cluster_id` | `ReadLogs` |  A `crn` that cannot be parsed, or that ends in one of these resource types but does not have every element of its shape, is rejected with `422 Invalid CRN`, and an element that does not match its position in the shape with `422 Invalid CRN element {element}`; see Choosing a log source for the order of the checks. For connectors, the service first resolves the connector's logical cluster ID through the metadata service; a connector it cannot resolve is reported as `403`, never `404`, so that the response does not reveal whether the resource exists.  The lookback for every resource type above is 3 days. An `end_time` older than that returns an empty `200` without querying OpenSearch; a `start_time` older than that is silently clamped.  Audit logs (`audit_logs/v1`) are served by the same host under a separate path and are not part of this spec.  # Discovery  `GET /logs/v1/sources` and `GET /logs/v1/sources/{name}` let a client learn which log source types this deployment's catalog serves, and what each one supports (crn shape, accepted `search` keys and their allowed values, returned fields, and lookback).  These endpoints make no OpenSearch calls themselves and are not subject to the per-resource rate limit described below. They're instead subject to the per-IP limit plus their own (user, organization) limit.  # Choosing a log source  `source` is optional. The service infers the log source from the `crn`'s trailing element, so a client only sets it to pick one of several log sources for the same resource type.  The service checks the `crn` and `source` in this order and returns the first failure as a `422`:  1. The `crn` cannot be parsed: `Invalid CRN`. 2. There is no `source`, and the `crn`'s trailing element is a resource type with more than one    log source: `source is required for resource type {type}; one of [{sources}]`. 3. The log source cannot be resolved. With no `source` and a trailing element that is not a    catalog resource type: `Invalid resource type {type}`. With a `source` that does not exist:    `Unknown source {name}`. With a `source` that does not serve the `crn`'s trailing resource    type: `source {name} does not support resource type {type}`. 4. The `crn` does not have every element of the resolved resource type's shape: `Invalid CRN`.    An element that does not match its position: `Invalid CRN element {element}`.  Every query targets one resource, so the `crn` must be complete. A `crn` truncated to an organization or an environment ends in `organization` or `environment`, which is not a resource type, so it fails at step 3: `Invalid resource type environment` without a `source`, and `source cluster-health does not support resource type environment` with one. It never reaches step 4. A `crn` that ends in a real resource type but is missing elements, for example one that stops after `connector=my-connector` with no environment before it, fails at step 4 with `Invalid CRN`.  # Pagination  `POST /logs/v1/search` paginates with an opaque `page_token`. When a page is full, the response carries the URL of the next page in `metadata.next`; `metadata` is absent when the page was not full or was empty. The default page size is 100 documents; the maximum is 1000.  # Rate limiting  The rate-limit sidecar enforces, per (resource CRN, user, organization, client type), 1 request per second for UI clients (User-Agent containing `mozilla`) and 5 requests per minute for every other client, plus 5 requests per second per client IP before authentication. The per-resource limit applies only to `POST /logs/v1/search`, which is the only endpoint with a resource CRN to key it on; the per-IP limit applies to every endpoint in this spec, including the sources endpoints. The sources endpoints additionally enforce their own limit of 60 requests per minute per (user, organization) pair, independent of the per-IP check — discovery has no resource to key a per-resource limit on, and no UI-vs-other-client distinction the way search does. Every response that passed the per-IP limiter carries `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset` headers, including non-429 responses; a rate-limited request receives `429`. `Retry-After` is not sent. These numbers come from `ratelimit-config.yaml` and `discovery-ratelimit-config.yaml` in the service's Halyard spec as of 2026-09-15 and are deployment configuration, not part of the API contract.

API version: 0.1.0
Contact: logging-platform@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
	ContextHttpSignatureAuth = contextKey("httpsignature")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: ServerConfigurations{
			{
				URL:         "https://api.logging.confluent.cloud",
				Description: "Confluent Cloud production",
			},
		},
		OperationServers: map[string]ServerConfigurations{},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("Index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("The variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
