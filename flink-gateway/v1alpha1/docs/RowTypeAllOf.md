# RowTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The data type of the column. | [optional] 
**Fields** | [**[]RowFieldType**](RowFieldType.md) | The fields of the row. Can be of type [INTEGER, DECIMAL, CHARACTER, ROW, ARRAY, TIMESTAMP, MAP&lt;INT, STRING&gt;] | 

## Methods

### NewRowTypeAllOf

`func NewRowTypeAllOf(fields []RowFieldType, ) *RowTypeAllOf`

NewRowTypeAllOf instantiates a new RowTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowTypeAllOfWithDefaults

`func NewRowTypeAllOfWithDefaults() *RowTypeAllOf`

NewRowTypeAllOfWithDefaults instantiates a new RowTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RowTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RowTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RowTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RowTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFields

`func (o *RowTypeAllOf) GetFields() []RowFieldType`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RowTypeAllOf) GetFieldsOk() (*[]RowFieldType, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RowTypeAllOf) SetFields(v []RowFieldType)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


