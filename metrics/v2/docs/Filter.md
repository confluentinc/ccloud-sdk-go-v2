# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Field** | Pointer to **string** | The field to filter on; see [here](#section/Object-Model/Labels) on using labels as filter fields.  | [optional] 
**Value** | [**FieldFilterValue**](FieldFilterValue.md) |  | 
**Filters** | [**[]Filter**](Filter.md) |  | 
**Filter** | [**Filter**](Filter.md) |  | 

## Methods

### NewFilter

`func NewFilter(op string, value FieldFilterValue, filters []Filter, filter Filter, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *Filter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Filter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Filter) SetOp(v string)`

SetOp sets Op field to given value.


### GetField

`func (o *Filter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Filter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Filter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Filter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *Filter) GetValue() FieldFilterValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Filter) GetValueOk() (*FieldFilterValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Filter) SetValue(v FieldFilterValue)`

SetValue sets Value field to given value.


### GetFilters

`func (o *Filter) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Filter) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Filter) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.


### GetFilter

`func (o *Filter) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Filter) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Filter) SetFilter(v Filter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


