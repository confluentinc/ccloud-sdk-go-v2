# DecimalType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Precision** | **int32** | The precision of the decimal type (i.e. the number of digits in the number). | 
**Scale** | **int32** | The scale of the decimal type (i.e. the number of digits to the right of the decimal point). | 

## Methods

### NewDecimalType

`func NewDecimalType(nullable bool, type_ string, precision int32, scale int32, ) *DecimalType`

NewDecimalType instantiates a new DecimalType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecimalTypeWithDefaults

`func NewDecimalTypeWithDefaults() *DecimalType`

NewDecimalTypeWithDefaults instantiates a new DecimalType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *DecimalType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *DecimalType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *DecimalType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *DecimalType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecimalType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecimalType) SetType(v string)`

SetType sets Type field to given value.


### GetPrecision

`func (o *DecimalType) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *DecimalType) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *DecimalType) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.


### GetScale

`func (o *DecimalType) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *DecimalType) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *DecimalType) SetScale(v int32)`

SetScale sets Scale field to given value.



### AsDataType

`func (s *DecimalType) AsDataType() DataType`

Convenience method to wrap this instance of DecimalType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


