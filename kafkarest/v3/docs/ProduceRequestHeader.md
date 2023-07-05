# ProduceRequestHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProduceRequestHeader

`func NewProduceRequestHeader(name string, ) *ProduceRequestHeader`

NewProduceRequestHeader instantiates a new ProduceRequestHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProduceRequestHeaderWithDefaults

`func NewProduceRequestHeaderWithDefaults() *ProduceRequestHeader`

NewProduceRequestHeaderWithDefaults instantiates a new ProduceRequestHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProduceRequestHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProduceRequestHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProduceRequestHeader) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ProduceRequestHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProduceRequestHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProduceRequestHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProduceRequestHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProduceRequestHeader) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProduceRequestHeader) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


