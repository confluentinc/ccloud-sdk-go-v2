# UnaryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Filter** | [**Filter**](Filter.md) |  | 

## Methods

### NewUnaryFilter

`func NewUnaryFilter(op string, filter Filter, ) *UnaryFilter`

NewUnaryFilter instantiates a new UnaryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnaryFilterWithDefaults

`func NewUnaryFilterWithDefaults() *UnaryFilter`

NewUnaryFilterWithDefaults instantiates a new UnaryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *UnaryFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UnaryFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UnaryFilter) SetOp(v string)`

SetOp sets Op field to given value.


### GetFilter

`func (o *UnaryFilter) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UnaryFilter) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UnaryFilter) SetFilter(v Filter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


