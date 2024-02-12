# StructType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Fields** | [**[]StructField**](StructField.md) |  | 

## Methods

### NewStructType

`func NewStructType(type_ string, fields []StructField, ) *StructType`

NewStructType instantiates a new StructType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructTypeWithDefaults

`func NewStructTypeWithDefaults() *StructType`

NewStructTypeWithDefaults instantiates a new StructType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StructType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StructType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StructType) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *StructType) GetFields() []StructField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *StructType) GetFieldsOk() (*[]StructField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *StructType) SetFields(v []StructField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


