# TimestampWithLocalTimeZoneType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Precision** | Pointer to **int32** | The precision of the decimal type (i.e. the number of digits in the number). | [optional] 

## Methods

### NewTimestampWithLocalTimeZoneType

`func NewTimestampWithLocalTimeZoneType(nullable bool, type_ string, ) *TimestampWithLocalTimeZoneType`

NewTimestampWithLocalTimeZoneType instantiates a new TimestampWithLocalTimeZoneType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampWithLocalTimeZoneTypeWithDefaults

`func NewTimestampWithLocalTimeZoneTypeWithDefaults() *TimestampWithLocalTimeZoneType`

NewTimestampWithLocalTimeZoneTypeWithDefaults instantiates a new TimestampWithLocalTimeZoneType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *TimestampWithLocalTimeZoneType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *TimestampWithLocalTimeZoneType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *TimestampWithLocalTimeZoneType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *TimestampWithLocalTimeZoneType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimestampWithLocalTimeZoneType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimestampWithLocalTimeZoneType) SetType(v string)`

SetType sets Type field to given value.


### GetPrecision

`func (o *TimestampWithLocalTimeZoneType) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *TimestampWithLocalTimeZoneType) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *TimestampWithLocalTimeZoneType) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *TimestampWithLocalTimeZoneType) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.


### AsDataType

`func (s *TimestampWithLocalTimeZoneType) AsDataType() DataType`

Convenience method to wrap this instance of TimestampWithLocalTimeZoneType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


