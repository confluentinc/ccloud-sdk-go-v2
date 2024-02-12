# SortField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **int32** |  | 
**Transform** | **string** |  | 
**Direction** | [**SortDirection**](SortDirection.md) |  | 
**NullOrder** | [**NullOrder**](NullOrder.md) |  | 

## Methods

### NewSortField

`func NewSortField(sourceId int32, transform string, direction SortDirection, nullOrder NullOrder, ) *SortField`

NewSortField instantiates a new SortField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortFieldWithDefaults

`func NewSortFieldWithDefaults() *SortField`

NewSortFieldWithDefaults instantiates a new SortField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SortField) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SortField) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SortField) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.


### GetTransform

`func (o *SortField) GetTransform() string`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *SortField) GetTransformOk() (*string, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *SortField) SetTransform(v string)`

SetTransform sets Transform field to given value.


### GetDirection

`func (o *SortField) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SortField) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SortField) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.


### GetNullOrder

`func (o *SortField) GetNullOrder() NullOrder`

GetNullOrder returns the NullOrder field if non-nil, zero value otherwise.

### GetNullOrderOk

`func (o *SortField) GetNullOrderOk() (*NullOrder, bool)`

GetNullOrderOk returns a tuple with the NullOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullOrder

`func (o *SortField) SetNullOrder(v NullOrder)`

SetNullOrder sets NullOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


