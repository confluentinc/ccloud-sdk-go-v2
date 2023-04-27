# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchParameters** | Pointer to [**SearchParams**](SearchParams.md) |  | [optional] 
**Types** | Pointer to **[]string** | The types | [optional] 
**Entities** | Pointer to [**[]EntityHeader**](EntityHeader.md) | The entities | [optional] 
**ReferredEntities** | Pointer to [**map[string]EntityHeader**](EntityHeader.md) | The referred entities | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult() *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchParameters

`func (o *SearchResult) GetSearchParameters() SearchParams`

GetSearchParameters returns the SearchParameters field if non-nil, zero value otherwise.

### GetSearchParametersOk

`func (o *SearchResult) GetSearchParametersOk() (*SearchParams, bool)`

GetSearchParametersOk returns a tuple with the SearchParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParameters

`func (o *SearchResult) SetSearchParameters(v SearchParams)`

SetSearchParameters sets SearchParameters field to given value.

### HasSearchParameters

`func (o *SearchResult) HasSearchParameters() bool`

HasSearchParameters returns a boolean if a field has been set.

### GetTypes

`func (o *SearchResult) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchResult) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchResult) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchResult) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetEntities

`func (o *SearchResult) GetEntities() []EntityHeader`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SearchResult) GetEntitiesOk() (*[]EntityHeader, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SearchResult) SetEntities(v []EntityHeader)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SearchResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetReferredEntities

`func (o *SearchResult) GetReferredEntities() map[string]EntityHeader`

GetReferredEntities returns the ReferredEntities field if non-nil, zero value otherwise.

### GetReferredEntitiesOk

`func (o *SearchResult) GetReferredEntitiesOk() (*map[string]EntityHeader, bool)`

GetReferredEntitiesOk returns a tuple with the ReferredEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredEntities

`func (o *SearchResult) SetReferredEntities(v map[string]EntityHeader)`

SetReferredEntities sets ReferredEntities field to given value.

### HasReferredEntities

`func (o *SearchResult) HasReferredEntities() bool`

HasReferredEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


