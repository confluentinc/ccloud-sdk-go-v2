# CompoundFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Filters** | [**[]Filter**](Filter.md) |  | 

## Methods

### NewCompoundFilter

`func NewCompoundFilter(op string, filters []Filter, ) *CompoundFilter`

NewCompoundFilter instantiates a new CompoundFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompoundFilterWithDefaults

`func NewCompoundFilterWithDefaults() *CompoundFilter`

NewCompoundFilterWithDefaults instantiates a new CompoundFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *CompoundFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CompoundFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CompoundFilter) SetOp(v string)`

SetOp sets Op field to given value.


### GetFilters

`func (o *CompoundFilter) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CompoundFilter) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CompoundFilter) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


