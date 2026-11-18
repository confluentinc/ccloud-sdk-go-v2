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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type LogsLoggingV1Api interface {

	/*
			DescribeLoggingV1Source Describe one log source

			Describes one source by its catalog key (the `name` field of an entry in
		`GET /logs/v1/sources`), with the same fields as that entry.


			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param name The source's catalog key, for example `connector` or `statement`.
			 @return ApiDescribeLoggingV1SourceRequest
	*/
	DescribeLoggingV1Source(ctx _context.Context, name string) ApiDescribeLoggingV1SourceRequest

	// DescribeLoggingV1SourceExecute executes the request
	//  @return LoggingV1Source
	DescribeLoggingV1SourceExecute(r ApiDescribeLoggingV1SourceRequest) (LoggingV1Source, *_nethttp.Response, error)

	/*
			ListLoggingV1Sources List available log sources

			Lists every log source `POST /logs/v1/search` can query, with the search fields and
		returned fields each one supports. See the Discovery section above.


			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiListLoggingV1SourcesRequest
	*/
	ListLoggingV1Sources(ctx _context.Context) ApiListLoggingV1SourcesRequest

	// ListLoggingV1SourcesExecute executes the request
	//  @return LoggingV1SourceList
	ListLoggingV1SourcesExecute(r ApiListLoggingV1SourcesRequest) (LoggingV1SourceList, *_nethttp.Response, error)

	/*
			SearchLoggingV1Logs Search logs

			Searches log documents for the resource identified by `crn`.

		`search` filters on the keys the resource type accepts (see the Resource types section
		above). An unknown key, or a value outside the key's allowed values, returns `422`. The
		reserved `search_text` key runs a phrase or wildcard match over the resource type's text
		fields (`message`, and `exception.stacktrace` for connectors).

		`start_time` and `end_time` define an inclusive UTC window. An `end_time` older than the
		3-day lookback returns an empty `200`; a `start_time` older than the lookback is silently
		clamped to it. Results page via `page_token`, see the Pagination section above. Requests are
		subject to the rate limits described in the Rate limiting section above.


			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiSearchLoggingV1LogsRequest
	*/
	SearchLoggingV1Logs(ctx _context.Context) ApiSearchLoggingV1LogsRequest

	// SearchLoggingV1LogsExecute executes the request
	//  @return LoggingV1SearchResponse
	SearchLoggingV1LogsExecute(r ApiSearchLoggingV1LogsRequest) (LoggingV1SearchResponse, *_nethttp.Response, error)
}

// LogsLoggingV1ApiService LogsLoggingV1Api service
type LogsLoggingV1ApiService service

type ApiDescribeLoggingV1SourceRequest struct {
	ctx        _context.Context
	ApiService LogsLoggingV1Api
	name       string
}

func (r ApiDescribeLoggingV1SourceRequest) Execute() (LoggingV1Source, *_nethttp.Response, error) {
	return r.ApiService.DescribeLoggingV1SourceExecute(r)
}

/*
DescribeLoggingV1Source Describe one log source

Describes one source by its catalog key (the `name` field of an entry in
`GET /logs/v1/sources`), with the same fields as that entry.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The source's catalog key, for example `connector` or `statement`.
	@return ApiDescribeLoggingV1SourceRequest
*/
func (a *LogsLoggingV1ApiService) DescribeLoggingV1Source(ctx _context.Context, name string) ApiDescribeLoggingV1SourceRequest {
	return ApiDescribeLoggingV1SourceRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return LoggingV1Source
func (a *LogsLoggingV1ApiService) DescribeLoggingV1SourceExecute(r ApiDescribeLoggingV1SourceRequest) (LoggingV1Source, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LoggingV1Source
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsLoggingV1ApiService.DescribeLoggingV1Source")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs/v1/sources/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListLoggingV1SourcesRequest struct {
	ctx        _context.Context
	ApiService LogsLoggingV1Api
}

func (r ApiListLoggingV1SourcesRequest) Execute() (LoggingV1SourceList, *_nethttp.Response, error) {
	return r.ApiService.ListLoggingV1SourcesExecute(r)
}

/*
ListLoggingV1Sources List available log sources

Lists every log source `POST /logs/v1/search` can query, with the search fields and
returned fields each one supports. See the Discovery section above.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListLoggingV1SourcesRequest
*/
func (a *LogsLoggingV1ApiService) ListLoggingV1Sources(ctx _context.Context) ApiListLoggingV1SourcesRequest {
	return ApiListLoggingV1SourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LoggingV1SourceList
func (a *LogsLoggingV1ApiService) ListLoggingV1SourcesExecute(r ApiListLoggingV1SourcesRequest) (LoggingV1SourceList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LoggingV1SourceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsLoggingV1ApiService.ListLoggingV1Sources")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs/v1/sources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchLoggingV1LogsRequest struct {
	ctx                    _context.Context
	ApiService             LogsLoggingV1Api
	loggingV1SearchRequest *LoggingV1SearchRequest
	pageSize               *int32
	pageToken              *string
}

func (r ApiSearchLoggingV1LogsRequest) LoggingV1SearchRequest(loggingV1SearchRequest LoggingV1SearchRequest) ApiSearchLoggingV1LogsRequest {
	r.loggingV1SearchRequest = &loggingV1SearchRequest
	return r
}

// Number of documents per page. Values outside 1..1000 are rejected with 422.
func (r ApiSearchLoggingV1LogsRequest) PageSize(pageSize int32) ApiSearchLoggingV1LogsRequest {
	r.pageSize = &pageSize
	return r
}

// Opaque continuation token taken from &#x60;metadata.next&#x60; of the previous response. The page size encoded in the token takes precedence over &#x60;page_size&#x60;. A token that cannot be decoded is ignored and the first page is returned.
func (r ApiSearchLoggingV1LogsRequest) PageToken(pageToken string) ApiSearchLoggingV1LogsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiSearchLoggingV1LogsRequest) Execute() (LoggingV1SearchResponse, *_nethttp.Response, error) {
	return r.ApiService.SearchLoggingV1LogsExecute(r)
}

/*
SearchLoggingV1Logs Search logs

Searches log documents for the resource identified by `crn`.

`search` filters on the keys the resource type accepts (see the Resource types section
above). An unknown key, or a value outside the key's allowed values, returns `422`. The
reserved `search_text` key runs a phrase or wildcard match over the resource type's text
fields (`message`, and `exception.stacktrace` for connectors).

`start_time` and `end_time` define an inclusive UTC window. An `end_time` older than the
3-day lookback returns an empty `200`; a `start_time` older than the lookback is silently
clamped to it. Results page via `page_token`, see the Pagination section above. Requests are
subject to the rate limits described in the Rate limiting section above.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchLoggingV1LogsRequest
*/
func (a *LogsLoggingV1ApiService) SearchLoggingV1Logs(ctx _context.Context) ApiSearchLoggingV1LogsRequest {
	return ApiSearchLoggingV1LogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LoggingV1SearchResponse
func (a *LogsLoggingV1ApiService) SearchLoggingV1LogsExecute(r ApiSearchLoggingV1LogsRequest) (LoggingV1SearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LoggingV1SearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsLoggingV1ApiService.SearchLoggingV1Logs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs/v1/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.loggingV1SearchRequest == nil {
		return localVarReturnValue, nil, reportError("loggingV1SearchRequest is required and must be specified")
	}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.loggingV1SearchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v LoggingV1ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
