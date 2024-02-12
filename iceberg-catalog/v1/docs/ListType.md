# ListType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ElementId** | **int32** |  | 
**Element** | [**Type**](Type.md) |  | 
**ElementRequired** | **bool** |  | 

## Methods

### NewListType

`func NewListType(type_ string, elementId int32, element Type, elementRequired bool, ) *ListType`

NewListType instantiates a new ListType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTypeWithDefaults

`func NewListTypeWithDefaults() *ListType`

NewListTypeWithDefaults instantiates a new ListType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListType) SetType(v string)`

SetType sets Type field to given value.


### GetElementId

`func (o *ListType) GetElementId() int32`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ListType) GetElementIdOk() (*int32, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ListType) SetElementId(v int32)`

SetElementId sets ElementId field to given value.


### GetElement

`func (o *ListType) GetElement() Type`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ListType) GetElementOk() (*Type, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ListType) SetElement(v Type)`

SetElement sets Element field to given value.


### GetElementRequired

`func (o *ListType) GetElementRequired() bool`

GetElementRequired returns the ElementRequired field if non-nil, zero value otherwise.

### GetElementRequiredOk

`func (o *ListType) GetElementRequiredOk() (*bool, bool)`

GetElementRequiredOk returns a tuple with the ElementRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementRequired

`func (o *ListType) SetElementRequired(v bool)`

SetElementRequired sets ElementRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


