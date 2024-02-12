# Type

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Fields** | [**[]StructField**](StructField.md) |  | 
**ElementId** | **int32** |  | 
**Element** | [**Type**](Type.md) |  | 
**ElementRequired** | **bool** |  | 
**KeyId** | **int32** |  | 
**Key** | [**Type**](Type.md) |  | 
**ValueId** | **int32** |  | 
**Value** | [**Type**](Type.md) |  | 
**ValueRequired** | **bool** |  | 

## Methods

### NewType

`func NewType(type_ string, fields []StructField, elementId int32, element Type, elementRequired bool, keyId int32, key Type, valueId int32, value Type, valueRequired bool, ) *Type`

NewType instantiates a new Type object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeWithDefaults

`func NewTypeWithDefaults() *Type`

NewTypeWithDefaults instantiates a new Type object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Type) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Type) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Type) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *Type) GetFields() []StructField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Type) GetFieldsOk() (*[]StructField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Type) SetFields(v []StructField)`

SetFields sets Fields field to given value.


### GetElementId

`func (o *Type) GetElementId() int32`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *Type) GetElementIdOk() (*int32, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *Type) SetElementId(v int32)`

SetElementId sets ElementId field to given value.


### GetElement

`func (o *Type) GetElement() Type`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *Type) GetElementOk() (*Type, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *Type) SetElement(v Type)`

SetElement sets Element field to given value.


### GetElementRequired

`func (o *Type) GetElementRequired() bool`

GetElementRequired returns the ElementRequired field if non-nil, zero value otherwise.

### GetElementRequiredOk

`func (o *Type) GetElementRequiredOk() (*bool, bool)`

GetElementRequiredOk returns a tuple with the ElementRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementRequired

`func (o *Type) SetElementRequired(v bool)`

SetElementRequired sets ElementRequired field to given value.


### GetKeyId

`func (o *Type) GetKeyId() int32`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *Type) GetKeyIdOk() (*int32, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *Type) SetKeyId(v int32)`

SetKeyId sets KeyId field to given value.


### GetKey

`func (o *Type) GetKey() Type`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Type) GetKeyOk() (*Type, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Type) SetKey(v Type)`

SetKey sets Key field to given value.


### GetValueId

`func (o *Type) GetValueId() int32`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *Type) GetValueIdOk() (*int32, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *Type) SetValueId(v int32)`

SetValueId sets ValueId field to given value.


### GetValue

`func (o *Type) GetValue() Type`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Type) GetValueOk() (*Type, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Type) SetValue(v Type)`

SetValue sets Value field to given value.


### GetValueRequired

`func (o *Type) GetValueRequired() bool`

GetValueRequired returns the ValueRequired field if non-nil, zero value otherwise.

### GetValueRequiredOk

`func (o *Type) GetValueRequiredOk() (*bool, bool)`

GetValueRequiredOk returns a tuple with the ValueRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRequired

`func (o *Type) SetValueRequired(v bool)`

SetValueRequired sets ValueRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


