# SortOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** |  | [readonly] 
**Fields** | [**[]SortField**](SortField.md) |  | 

## Methods

### NewSortOrder

`func NewSortOrder(orderId int32, fields []SortField, ) *SortOrder`

NewSortOrder instantiates a new SortOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortOrderWithDefaults

`func NewSortOrderWithDefaults() *SortOrder`

NewSortOrderWithDefaults instantiates a new SortOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *SortOrder) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SortOrder) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SortOrder) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetFields

`func (o *SortOrder) GetFields() []SortField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SortOrder) GetFieldsOk() (*[]SortField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SortOrder) SetFields(v []SortField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


