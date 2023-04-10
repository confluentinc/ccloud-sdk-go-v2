# RowType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullable** | **bool** | Indicates whether values in this column can be null. | 
**Type** | **string** | The data type of the column. | 
**Fields** | [**[]RowFieldType**](RowFieldType.md) | The fields of the row. Can be of type [INTEGER, DECIMAL, CHARACTER, ROW, ARRAY, TIMESTAMP, MAP&lt;INT, STRING&gt;] | 

## Methods

### NewRowType

`func NewRowType(nullable bool, type_ string, fields []RowFieldType, ) *RowType`

NewRowType instantiates a new RowType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowTypeWithDefaults

`func NewRowTypeWithDefaults() *RowType`

NewRowTypeWithDefaults instantiates a new RowType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullable

`func (o *RowType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *RowType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *RowType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetType

`func (o *RowType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RowType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RowType) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *RowType) GetFields() []RowFieldType`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RowType) GetFieldsOk() (*[]RowFieldType, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RowType) SetFields(v []RowFieldType)`

SetFields sets Fields field to given value.



### AsDataType

`func (s *RowType) AsDataType() DataType`

Convenience method to wrap this instance of RowType in DataType

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


