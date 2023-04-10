# DecimalTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The data type of the column. | [optional] 
**Precision** | **int32** | The precision of the decimal type (i.e. the number of digits in the number). | 
**Scale** | **int32** | The scale of the decimal type (i.e. the number of digits to the right of the decimal point). | 

## Methods

### NewDecimalTypeAllOf

`func NewDecimalTypeAllOf(precision int32, scale int32, ) *DecimalTypeAllOf`

NewDecimalTypeAllOf instantiates a new DecimalTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecimalTypeAllOfWithDefaults

`func NewDecimalTypeAllOfWithDefaults() *DecimalTypeAllOf`

NewDecimalTypeAllOfWithDefaults instantiates a new DecimalTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DecimalTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecimalTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecimalTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DecimalTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrecision

`func (o *DecimalTypeAllOf) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *DecimalTypeAllOf) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *DecimalTypeAllOf) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.


### GetScale

`func (o *DecimalTypeAllOf) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *DecimalTypeAllOf) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *DecimalTypeAllOf) SetScale(v int32)`

SetScale sets Scale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


