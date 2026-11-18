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

// LoggingV1AggregationResponseTimestampHistogramBuckets struct for LoggingV1AggregationResponseTimestampHistogramBuckets
type LoggingV1AggregationResponseTimestampHistogramBuckets struct {
	// Bucket start, RFC 3339 UTC.
	Key *string `json:"key,omitempty"`
	// Number of matching documents in this bucket.
	Count           *int64                                                         `json:"count,omitempty"`
	AggregatedTerms *LoggingV1AggregationResponseTimestampHistogramAggregatedTerms `json:"aggregated_terms,omitempty"`
}

// NewLoggingV1AggregationResponseTimestampHistogramBuckets instantiates a new LoggingV1AggregationResponseTimestampHistogramBuckets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingV1AggregationResponseTimestampHistogramBuckets() *LoggingV1AggregationResponseTimestampHistogramBuckets {
	this := LoggingV1AggregationResponseTimestampHistogramBuckets{}
	return &this
}

// NewLoggingV1AggregationResponseTimestampHistogramBucketsWithDefaults instantiates a new LoggingV1AggregationResponseTimestampHistogramBuckets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingV1AggregationResponseTimestampHistogramBucketsWithDefaults() *LoggingV1AggregationResponseTimestampHistogramBuckets {
	this := LoggingV1AggregationResponseTimestampHistogramBuckets{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetKey(v string) {
	o.Key = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetCount(v int64) {
	o.Count = &v
}

// GetAggregatedTerms returns the AggregatedTerms field value if set, zero value otherwise.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetAggregatedTerms() LoggingV1AggregationResponseTimestampHistogramAggregatedTerms {
	if o == nil || o.AggregatedTerms == nil {
		var ret LoggingV1AggregationResponseTimestampHistogramAggregatedTerms
		return ret
	}
	return *o.AggregatedTerms
}

// GetAggregatedTermsOk returns a tuple with the AggregatedTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) GetAggregatedTermsOk() (*LoggingV1AggregationResponseTimestampHistogramAggregatedTerms, bool) {
	if o == nil || o.AggregatedTerms == nil {
		return nil, false
	}
	return o.AggregatedTerms, true
}

// HasAggregatedTerms returns a boolean if a field has been set.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) HasAggregatedTerms() bool {
	if o != nil && o.AggregatedTerms != nil {
		return true
	}

	return false
}

// SetAggregatedTerms gets a reference to the given LoggingV1AggregationResponseTimestampHistogramAggregatedTerms and assigns it to the AggregatedTerms field.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) SetAggregatedTerms(v LoggingV1AggregationResponseTimestampHistogramAggregatedTerms) {
	o.AggregatedTerms = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) Redact() {
	o.recurseRedact(o.Key)
	o.recurseRedact(o.Count)
	o.recurseRedact(o.AggregatedTerms)
}

func (o *LoggingV1AggregationResponseTimestampHistogramBuckets) recurseRedact(v interface{}) {
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

func (o LoggingV1AggregationResponseTimestampHistogramBuckets) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o LoggingV1AggregationResponseTimestampHistogramBuckets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.AggregatedTerms != nil {
		toSerialize["aggregated_terms"] = o.AggregatedTerms
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableLoggingV1AggregationResponseTimestampHistogramBuckets struct {
	value *LoggingV1AggregationResponseTimestampHistogramBuckets
	isSet bool
}

func (v NullableLoggingV1AggregationResponseTimestampHistogramBuckets) Get() *LoggingV1AggregationResponseTimestampHistogramBuckets {
	return v.value
}

func (v *NullableLoggingV1AggregationResponseTimestampHistogramBuckets) Set(val *LoggingV1AggregationResponseTimestampHistogramBuckets) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingV1AggregationResponseTimestampHistogramBuckets) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingV1AggregationResponseTimestampHistogramBuckets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingV1AggregationResponseTimestampHistogramBuckets(val *LoggingV1AggregationResponseTimestampHistogramBuckets) *NullableLoggingV1AggregationResponseTimestampHistogramBuckets {
	return &NullableLoggingV1AggregationResponseTimestampHistogramBuckets{value: val, isSet: true}
}

func (v NullableLoggingV1AggregationResponseTimestampHistogramBuckets) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableLoggingV1AggregationResponseTimestampHistogramBuckets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
