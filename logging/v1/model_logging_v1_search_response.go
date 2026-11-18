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
	"encoding/json"
)

import (
	"reflect"
)

// LoggingV1SearchResponse One page of matching log documents, together with the next-page link and any requested aggregation.
type LoggingV1SearchResponse struct {
	// Always `v1` on this endpoint.
	ApiVersion string `json:"api_version,omitempty"`
	// Always `Logs` on this endpoint.
	Kind string `json:"kind,omitempty"`
	// Log documents, each limited to the returned fields listed for its resource type in the Resource types section above. The shape differs per resource type.
	Data         []map[string]interface{}         `json:"data,omitempty"`
	Metadata     *LoggingV1SearchResponseMetadata `json:"metadata,omitempty"`
	Aggregations *LoggingV1AggregationResponse    `json:"aggregations,omitempty"`
}

// NewLoggingV1SearchResponse instantiates a new LoggingV1SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingV1SearchResponse(apiVersion string, kind string, data []map[string]interface{}) *LoggingV1SearchResponse {
	this := LoggingV1SearchResponse{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Data = data
	return &this
}

// NewLoggingV1SearchResponseWithDefaults instantiates a new LoggingV1SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingV1SearchResponseWithDefaults() *LoggingV1SearchResponse {
	this := LoggingV1SearchResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *LoggingV1SearchResponse) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *LoggingV1SearchResponse) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *LoggingV1SearchResponse) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *LoggingV1SearchResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *LoggingV1SearchResponse) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *LoggingV1SearchResponse) SetKind(v string) {
	o.Kind = v
}

// GetData returns the Data field value
func (o *LoggingV1SearchResponse) GetData() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LoggingV1SearchResponse) GetDataOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LoggingV1SearchResponse) SetData(v []map[string]interface{}) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LoggingV1SearchResponse) GetMetadata() LoggingV1SearchResponseMetadata {
	if o == nil || o.Metadata == nil {
		var ret LoggingV1SearchResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1SearchResponse) GetMetadataOk() (*LoggingV1SearchResponseMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LoggingV1SearchResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LoggingV1SearchResponseMetadata and assigns it to the Metadata field.
func (o *LoggingV1SearchResponse) SetMetadata(v LoggingV1SearchResponseMetadata) {
	o.Metadata = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *LoggingV1SearchResponse) GetAggregations() LoggingV1AggregationResponse {
	if o == nil || o.Aggregations == nil {
		var ret LoggingV1AggregationResponse
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1SearchResponse) GetAggregationsOk() (*LoggingV1AggregationResponse, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *LoggingV1SearchResponse) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given LoggingV1AggregationResponse and assigns it to the Aggregations field.
func (o *LoggingV1SearchResponse) SetAggregations(v LoggingV1AggregationResponse) {
	o.Aggregations = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *LoggingV1SearchResponse) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Data)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Aggregations)
}

func (o *LoggingV1SearchResponse) recurseRedact(v interface{}) {
	type redactor interface {
		Redact()
	}
	if r, ok := v.(redactor); ok {
		r.Redact()
	} else {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < val.Len(); i++ {
				// support data types declared without pointers
				o.recurseRedact(val.Index(i).Interface())
				// ... and data types that were declared without but need pointers (for Redact)
				if val.Index(i).CanAddr() {
					o.recurseRedact(val.Index(i).Addr().Interface())
				}
			}
		}
	}
}

func (o LoggingV1SearchResponse) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o LoggingV1SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableLoggingV1SearchResponse struct {
	value *LoggingV1SearchResponse
	isSet bool
}

func (v NullableLoggingV1SearchResponse) Get() *LoggingV1SearchResponse {
	return v.value
}

func (v *NullableLoggingV1SearchResponse) Set(val *LoggingV1SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingV1SearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingV1SearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingV1SearchResponse(val *LoggingV1SearchResponse) *NullableLoggingV1SearchResponse {
	return &NullableLoggingV1SearchResponse{value: val, isSet: true}
}

func (v NullableLoggingV1SearchResponse) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableLoggingV1SearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
