# SearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeDeleted** | Pointer to **bool** | Whether to include deleted | [optional] 
**Limit** | Pointer to **int32** | The limit | [optional] 
**Offset** | Pointer to **int32** | The offset | [optional] 

## Methods

### NewSearchParams

`func NewSearchParams() *SearchParams`

NewSearchParams instantiates a new SearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchParamsWithDefaults

`func NewSearchParamsWithDefaults() *SearchParams`

NewSearchParamsWithDefaults instantiates a new SearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeDeleted

`func (o *SearchParams) GetIncludeDeleted() bool`

GetIncludeDeleted returns the IncludeDeleted field if non-nil, zero value otherwise.

### GetIncludeDeletedOk

`func (o *SearchParams) GetIncludeDeletedOk() (*bool, bool)`

GetIncludeDeletedOk returns a tuple with the IncludeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeleted

`func (o *SearchParams) SetIncludeDeleted(v bool)`

SetIncludeDeleted sets IncludeDeleted field to given value.

### HasIncludeDeleted

`func (o *SearchParams) HasIncludeDeleted() bool`

HasIncludeDeleted returns a boolean if a field has been set.

### GetLimit

`func (o *SearchParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


