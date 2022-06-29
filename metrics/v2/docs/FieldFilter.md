# FieldFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The comparison operator for the filter. Note that labels are compared _lexicographically_.  The &#x60;GT&#x60; or &#x60;GTE&#x60; operators can be used to page through grouped result sets that exceed the query limit.  | 
**Field** | Pointer to **string** | The field to filter on; see [here](#section/Object-Model/Labels) on using labels as filter fields.  | [optional] 
**Value** | [**FieldFilterValue**](FieldFilterValue.md) |  | 

## Methods

### NewFieldFilter

`func NewFieldFilter(op string, value FieldFilterValue, ) *FieldFilter`

NewFieldFilter instantiates a new FieldFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldFilterWithDefaults

`func NewFieldFilterWithDefaults() *FieldFilter`

NewFieldFilterWithDefaults instantiates a new FieldFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *FieldFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *FieldFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *FieldFilter) SetOp(v string)`

SetOp sets Op field to given value.


### GetField

`func (o *FieldFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldFilter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FieldFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *FieldFilter) GetValue() FieldFilterValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldFilter) GetValueOk() (*FieldFilterValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldFilter) SetValue(v FieldFilterValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


