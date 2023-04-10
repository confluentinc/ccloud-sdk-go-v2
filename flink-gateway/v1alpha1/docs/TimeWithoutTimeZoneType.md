# TimeWithoutTimeZoneType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Scale** | Pointer to **int32** | The scale of the time type. | [optional] 
**Precision** | Pointer to **int32** | The precision of the time type. | [optional] 

## Methods

### NewTimeWithoutTimeZoneType

`func NewTimeWithoutTimeZoneType(nullable bool, type_ string, ) *TimeWithoutTimeZoneType`

NewTimeWithoutTimeZoneType instantiates a new TimeWithoutTimeZoneType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWithoutTimeZoneTypeWithDefaults

`func NewTimeWithoutTimeZoneTypeWithDefaults() *TimeWithoutTimeZoneType`

NewTimeWithoutTimeZoneTypeWithDefaults instantiates a new TimeWithoutTimeZoneType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *TimeWithoutTimeZoneType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *TimeWithoutTimeZoneType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *TimeWithoutTimeZoneType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *TimeWithoutTimeZoneType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeWithoutTimeZoneType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeWithoutTimeZoneType) SetType(v string)`

SetType sets Type field to given value.


### GetScale

`func (o *TimeWithoutTimeZoneType) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *TimeWithoutTimeZoneType) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *TimeWithoutTimeZoneType) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *TimeWithoutTimeZoneType) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetPrecision

`func (o *TimeWithoutTimeZoneType) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *TimeWithoutTimeZoneType) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *TimeWithoutTimeZoneType) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *TimeWithoutTimeZoneType) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.


### AsDataType

`func (s *TimeWithoutTimeZoneType) AsDataType() DataType`

Convenience method to wrap this instance of TimeWithoutTimeZoneType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


