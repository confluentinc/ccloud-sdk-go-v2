# LoggingV1Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Name** | **string** | The catalog key that identifies this source. Pass this as &#x60;{name}&#x60; to &#x60;GET /logs/v1/sources/{name}&#x60;. Defaults to &#x60;resource_type&#x60; for a catalog entry with no separately configured source. | 
**DisplayName** | **string** | Human-readable name for this source. Defaults to &#x60;name&#x60; when not separately configured. | 
**ResourceType** | **string** | The trailing CRN element this source&#39;s resources use. Matches a row of the Resource types table above. | 
**IdPrefix** | Pointer to **string** | Prefix of this resource type&#39;s identifiers, when one is configured. Omitted when not configured. | [optional] 
**CrnElements** | **[]string** | The CRN element names a &#x60;crn&#x60; for this resource type must match, in order. See the Resource types table above. | 
**SearchFields** | [**[]LoggingV1SearchFieldDescriptor**](LoggingV1SearchFieldDescriptor.md) | Fields &#x60;search&#x60; accepts for this resource type, other than the reserved &#x60;search_text&#x60; key (see &#x60;full_text_search_supported&#x60;). | 
**FullTextSearchSupported** | **bool** | Whether the reserved &#x60;search_text&#x60; key is accepted for this resource type. | 
**FullTextSearchFields** | Pointer to **[]string** | Fields &#x60;search_text&#x60; matches against for this resource type. Present only when &#x60;full_text_search_supported&#x60; is true. | [optional] 
**ReturnedFields** | [**[]LoggingV1ReturnedFieldDescriptor**](LoggingV1ReturnedFieldDescriptor.md) | Fields a log document for this resource type carries in &#x60;POST /logs/v1/search&#x60;&#39;s response, with rendering metadata for each. See the Resource types table above. | 
**MaxLookbackDays** | **int32** | Lookback window in days for this resource type. An &#x60;end_time&#x60; older than this returns an empty 200 from &#x60;POST /logs/v1/search&#x60; without querying OpenSearch; a &#x60;start_time&#x60; older than this is silently clamped. | 

## Methods

### NewLoggingV1Source

`func NewLoggingV1Source(apiVersion string, kind string, name string, displayName string, resourceType string, crnElements []string, searchFields []LoggingV1SearchFieldDescriptor, fullTextSearchSupported bool, returnedFields []LoggingV1ReturnedFieldDescriptor, maxLookbackDays int32, ) *LoggingV1Source`

NewLoggingV1Source instantiates a new LoggingV1Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1SourceWithDefaults

`func NewLoggingV1SourceWithDefaults() *LoggingV1Source`

NewLoggingV1SourceWithDefaults instantiates a new LoggingV1Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *LoggingV1Source) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LoggingV1Source) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LoggingV1Source) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *LoggingV1Source) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LoggingV1Source) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LoggingV1Source) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *LoggingV1Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingV1Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingV1Source) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *LoggingV1Source) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LoggingV1Source) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LoggingV1Source) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetResourceType

`func (o *LoggingV1Source) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *LoggingV1Source) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *LoggingV1Source) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetIdPrefix

`func (o *LoggingV1Source) GetIdPrefix() string`

GetIdPrefix returns the IdPrefix field if non-nil, zero value otherwise.

### GetIdPrefixOk

`func (o *LoggingV1Source) GetIdPrefixOk() (*string, bool)`

GetIdPrefixOk returns a tuple with the IdPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPrefix

`func (o *LoggingV1Source) SetIdPrefix(v string)`

SetIdPrefix sets IdPrefix field to given value.

### HasIdPrefix

`func (o *LoggingV1Source) HasIdPrefix() bool`

HasIdPrefix returns a boolean if a field has been set.

### GetCrnElements

`func (o *LoggingV1Source) GetCrnElements() []string`

GetCrnElements returns the CrnElements field if non-nil, zero value otherwise.

### GetCrnElementsOk

`func (o *LoggingV1Source) GetCrnElementsOk() (*[]string, bool)`

GetCrnElementsOk returns a tuple with the CrnElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrnElements

`func (o *LoggingV1Source) SetCrnElements(v []string)`

SetCrnElements sets CrnElements field to given value.


### GetSearchFields

`func (o *LoggingV1Source) GetSearchFields() []LoggingV1SearchFieldDescriptor`

GetSearchFields returns the SearchFields field if non-nil, zero value otherwise.

### GetSearchFieldsOk

`func (o *LoggingV1Source) GetSearchFieldsOk() (*[]LoggingV1SearchFieldDescriptor, bool)`

GetSearchFieldsOk returns a tuple with the SearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFields

`func (o *LoggingV1Source) SetSearchFields(v []LoggingV1SearchFieldDescriptor)`

SetSearchFields sets SearchFields field to given value.


### GetFullTextSearchSupported

`func (o *LoggingV1Source) GetFullTextSearchSupported() bool`

GetFullTextSearchSupported returns the FullTextSearchSupported field if non-nil, zero value otherwise.

### GetFullTextSearchSupportedOk

`func (o *LoggingV1Source) GetFullTextSearchSupportedOk() (*bool, bool)`

GetFullTextSearchSupportedOk returns a tuple with the FullTextSearchSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextSearchSupported

`func (o *LoggingV1Source) SetFullTextSearchSupported(v bool)`

SetFullTextSearchSupported sets FullTextSearchSupported field to given value.


### GetFullTextSearchFields

`func (o *LoggingV1Source) GetFullTextSearchFields() []string`

GetFullTextSearchFields returns the FullTextSearchFields field if non-nil, zero value otherwise.

### GetFullTextSearchFieldsOk

`func (o *LoggingV1Source) GetFullTextSearchFieldsOk() (*[]string, bool)`

GetFullTextSearchFieldsOk returns a tuple with the FullTextSearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextSearchFields

`func (o *LoggingV1Source) SetFullTextSearchFields(v []string)`

SetFullTextSearchFields sets FullTextSearchFields field to given value.

### HasFullTextSearchFields

`func (o *LoggingV1Source) HasFullTextSearchFields() bool`

HasFullTextSearchFields returns a boolean if a field has been set.

### GetReturnedFields

`func (o *LoggingV1Source) GetReturnedFields() []LoggingV1ReturnedFieldDescriptor`

GetReturnedFields returns the ReturnedFields field if non-nil, zero value otherwise.

### GetReturnedFieldsOk

`func (o *LoggingV1Source) GetReturnedFieldsOk() (*[]LoggingV1ReturnedFieldDescriptor, bool)`

GetReturnedFieldsOk returns a tuple with the ReturnedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedFields

`func (o *LoggingV1Source) SetReturnedFields(v []LoggingV1ReturnedFieldDescriptor)`

SetReturnedFields sets ReturnedFields field to given value.


### GetMaxLookbackDays

`func (o *LoggingV1Source) GetMaxLookbackDays() int32`

GetMaxLookbackDays returns the MaxLookbackDays field if non-nil, zero value otherwise.

### GetMaxLookbackDaysOk

`func (o *LoggingV1Source) GetMaxLookbackDaysOk() (*int32, bool)`

GetMaxLookbackDaysOk returns a tuple with the MaxLookbackDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLookbackDays

`func (o *LoggingV1Source) SetMaxLookbackDays(v int32)`

SetMaxLookbackDays sets MaxLookbackDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


