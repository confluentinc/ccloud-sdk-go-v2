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

// LoggingV1Source One source's capabilities: the fields it can be searched on, and the fields it returns. Field names match `SourceMetadata` in `cc-logging-service`.
type LoggingV1Source struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind string `json:"kind,omitempty"`
	// The catalog key that identifies this source. Pass this as `{name}` to `GET /logs/v1/sources/{name}`. Defaults to `resource_type` for a catalog entry with no separately configured source.
	Name string `json:"name,omitempty"`
	// Human-readable name for this source. Defaults to `name` when not separately configured.
	DisplayName string `json:"display_name,omitempty"`
	// The trailing CRN element this source's resources use. Matches a row of the Resource types table above.
	ResourceType string `json:"resource_type,omitempty"`
	// Prefix of this resource type's identifiers, when one is configured. Omitted when not configured.
	IdPrefix *string `json:"id_prefix,omitempty"`
	// The CRN element names a `crn` for this resource type must match, in order. See the Resource types table above.
	CrnElements []string `json:"crn_elements,omitempty"`
	// Fields `search` accepts for this resource type, other than the reserved `search_text` key (see `full_text_search_supported`).
	SearchFields []LoggingV1SearchFieldDescriptor `json:"search_fields,omitempty"`
	// Whether the reserved `search_text` key is accepted for this resource type.
	FullTextSearchSupported bool `json:"full_text_search_supported,omitempty"`
	// Fields `search_text` matches against for this resource type. Present only when `full_text_search_supported` is true.
	FullTextSearchFields *[]string `json:"full_text_search_fields,omitempty"`
	// Fields a log document for this resource type carries in `POST /logs/v1/search`'s response, with rendering metadata for each. See the Resource types table above.
	ReturnedFields []LoggingV1ReturnedFieldDescriptor `json:"returned_fields,omitempty"`
	// Lookback window in days for this resource type. An `end_time` older than this returns an empty 200 from `POST /logs/v1/search` without querying OpenSearch; a `start_time` older than this is silently clamped.
	MaxLookbackDays int32 `json:"max_lookback_days,omitempty"`
}

// NewLoggingV1Source instantiates a new LoggingV1Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingV1Source(apiVersion string, kind string, name string, displayName string, resourceType string, crnElements []string, searchFields []LoggingV1SearchFieldDescriptor, fullTextSearchSupported bool, returnedFields []LoggingV1ReturnedFieldDescriptor, maxLookbackDays int32) *LoggingV1Source {
	this := LoggingV1Source{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Name = name
	this.DisplayName = displayName
	this.ResourceType = resourceType
	this.CrnElements = crnElements
	this.SearchFields = searchFields
	this.FullTextSearchSupported = fullTextSearchSupported
	this.ReturnedFields = returnedFields
	this.MaxLookbackDays = maxLookbackDays
	return &this
}

// NewLoggingV1SourceWithDefaults instantiates a new LoggingV1Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingV1SourceWithDefaults() *LoggingV1Source {
	this := LoggingV1Source{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *LoggingV1Source) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *LoggingV1Source) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *LoggingV1Source) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *LoggingV1Source) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *LoggingV1Source) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoggingV1Source) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *LoggingV1Source) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LoggingV1Source) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetResourceType returns the ResourceType field value
func (o *LoggingV1Source) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *LoggingV1Source) SetResourceType(v string) {
	o.ResourceType = v
}

// GetIdPrefix returns the IdPrefix field value if set, zero value otherwise.
func (o *LoggingV1Source) GetIdPrefix() string {
	if o == nil || o.IdPrefix == nil {
		var ret string
		return ret
	}
	return *o.IdPrefix
}

// GetIdPrefixOk returns a tuple with the IdPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetIdPrefixOk() (*string, bool) {
	if o == nil || o.IdPrefix == nil {
		return nil, false
	}
	return o.IdPrefix, true
}

// HasIdPrefix returns a boolean if a field has been set.
func (o *LoggingV1Source) HasIdPrefix() bool {
	if o != nil && o.IdPrefix != nil {
		return true
	}

	return false
}

// SetIdPrefix gets a reference to the given string and assigns it to the IdPrefix field.
func (o *LoggingV1Source) SetIdPrefix(v string) {
	o.IdPrefix = &v
}

// GetCrnElements returns the CrnElements field value
func (o *LoggingV1Source) GetCrnElements() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CrnElements
}

// GetCrnElementsOk returns a tuple with the CrnElements field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetCrnElementsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrnElements, true
}

// SetCrnElements sets field value
func (o *LoggingV1Source) SetCrnElements(v []string) {
	o.CrnElements = v
}

// GetSearchFields returns the SearchFields field value
func (o *LoggingV1Source) GetSearchFields() []LoggingV1SearchFieldDescriptor {
	if o == nil {
		var ret []LoggingV1SearchFieldDescriptor
		return ret
	}

	return o.SearchFields
}

// GetSearchFieldsOk returns a tuple with the SearchFields field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetSearchFieldsOk() (*[]LoggingV1SearchFieldDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchFields, true
}

// SetSearchFields sets field value
func (o *LoggingV1Source) SetSearchFields(v []LoggingV1SearchFieldDescriptor) {
	o.SearchFields = v
}

// GetFullTextSearchSupported returns the FullTextSearchSupported field value
func (o *LoggingV1Source) GetFullTextSearchSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FullTextSearchSupported
}

// GetFullTextSearchSupportedOk returns a tuple with the FullTextSearchSupported field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetFullTextSearchSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullTextSearchSupported, true
}

// SetFullTextSearchSupported sets field value
func (o *LoggingV1Source) SetFullTextSearchSupported(v bool) {
	o.FullTextSearchSupported = v
}

// GetFullTextSearchFields returns the FullTextSearchFields field value if set, zero value otherwise.
func (o *LoggingV1Source) GetFullTextSearchFields() []string {
	if o == nil || o.FullTextSearchFields == nil {
		var ret []string
		return ret
	}
	return *o.FullTextSearchFields
}

// GetFullTextSearchFieldsOk returns a tuple with the FullTextSearchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetFullTextSearchFieldsOk() (*[]string, bool) {
	if o == nil || o.FullTextSearchFields == nil {
		return nil, false
	}
	return o.FullTextSearchFields, true
}

// HasFullTextSearchFields returns a boolean if a field has been set.
func (o *LoggingV1Source) HasFullTextSearchFields() bool {
	if o != nil && o.FullTextSearchFields != nil {
		return true
	}

	return false
}

// SetFullTextSearchFields gets a reference to the given []string and assigns it to the FullTextSearchFields field.
func (o *LoggingV1Source) SetFullTextSearchFields(v []string) {
	o.FullTextSearchFields = &v
}

// GetReturnedFields returns the ReturnedFields field value
func (o *LoggingV1Source) GetReturnedFields() []LoggingV1ReturnedFieldDescriptor {
	if o == nil {
		var ret []LoggingV1ReturnedFieldDescriptor
		return ret
	}

	return o.ReturnedFields
}

// GetReturnedFieldsOk returns a tuple with the ReturnedFields field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetReturnedFieldsOk() (*[]LoggingV1ReturnedFieldDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnedFields, true
}

// SetReturnedFields sets field value
func (o *LoggingV1Source) SetReturnedFields(v []LoggingV1ReturnedFieldDescriptor) {
	o.ReturnedFields = v
}

// GetMaxLookbackDays returns the MaxLookbackDays field value
func (o *LoggingV1Source) GetMaxLookbackDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxLookbackDays
}

// GetMaxLookbackDaysOk returns a tuple with the MaxLookbackDays field value
// and a boolean to check if the value has been set.
func (o *LoggingV1Source) GetMaxLookbackDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLookbackDays, true
}

// SetMaxLookbackDays sets field value
func (o *LoggingV1Source) SetMaxLookbackDays(v int32) {
	o.MaxLookbackDays = v
}

// Redact resets all sensitive fields to their zero value.
func (o *LoggingV1Source) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Name)
	o.recurseRedact(&o.DisplayName)
	o.recurseRedact(&o.ResourceType)
	o.recurseRedact(o.IdPrefix)
	o.recurseRedact(&o.CrnElements)
	o.recurseRedact(&o.SearchFields)
	o.recurseRedact(&o.FullTextSearchSupported)
	o.recurseRedact(o.FullTextSearchFields)
	o.recurseRedact(&o.ReturnedFields)
	o.recurseRedact(&o.MaxLookbackDays)
}

func (o *LoggingV1Source) recurseRedact(v interface{}) {
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

func (o LoggingV1Source) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o LoggingV1Source) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.IdPrefix != nil {
		toSerialize["id_prefix"] = o.IdPrefix
	}
	if true {
		toSerialize["crn_elements"] = o.CrnElements
	}
	if true {
		toSerialize["search_fields"] = o.SearchFields
	}
	if true {
		toSerialize["full_text_search_supported"] = o.FullTextSearchSupported
	}
	if o.FullTextSearchFields != nil {
		toSerialize["full_text_search_fields"] = o.FullTextSearchFields
	}
	if true {
		toSerialize["returned_fields"] = o.ReturnedFields
	}
	if true {
		toSerialize["max_lookback_days"] = o.MaxLookbackDays
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableLoggingV1Source struct {
	value *LoggingV1Source
	isSet bool
}

func (v NullableLoggingV1Source) Get() *LoggingV1Source {
	return v.value
}

func (v *NullableLoggingV1Source) Set(val *LoggingV1Source) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingV1Source) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingV1Source) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingV1Source(val *LoggingV1Source) *NullableLoggingV1Source {
	return &NullableLoggingV1Source{value: val, isSet: true}
}

func (v NullableLoggingV1Source) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableLoggingV1Source) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
