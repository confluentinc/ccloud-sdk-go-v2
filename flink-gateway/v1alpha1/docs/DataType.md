# DataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data type of the column. | 
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Length** | Pointer to **int32** | The length of the data type. | [optional] 
**Precision** | Pointer to **int32** | The precision of the data type. | [optional] 
**Scale** | Pointer to **int32** | The scale of the data type. | [optional] 
**KeyType** | Pointer to [**DataType**](DataType.md) | The type of the key in the data type (if applicable). | [optional] 
**ValueType** | Pointer to [**DataType**](DataType.md) | The type of the value in the data type (if applicable). | [optional] 

## Methods

### NewDataType

`func NewDataType(type_ string, nullable bool, ) *DataType`

NewDataType instantiates a new DataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeWithDefaults

`func NewDataTypeWithDefaults() *DataType`

NewDataTypeWithDefaults instantiates a new DataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataType) SetType(v string)`

SetType sets Type field to given value.


### GetNullable

`func (o *DataType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *DataType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *DataType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetLength

`func (o *DataType) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *DataType) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *DataType) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *DataType) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPrecision

`func (o *DataType) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *DataType) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *DataType) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *DataType) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetScale

`func (o *DataType) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *DataType) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *DataType) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *DataType) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetKeyType

`func (o *DataType) GetKeyType() DataType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *DataType) GetKeyTypeOk() (*DataType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *DataType) SetKeyType(v DataType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *DataType) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetValueType

`func (o *DataType) GetValueType() DataType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *DataType) GetValueTypeOk() (*DataType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *DataType) SetValueType(v DataType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *DataType) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


