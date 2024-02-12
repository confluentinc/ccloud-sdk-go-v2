# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Fields** | [**[]StructField**](StructField.md) |  | 
**SchemaId** | Pointer to **int32** |  | [optional] [readonly] 
**IdentifierFieldIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSchema

`func NewSchema(type_ string, fields []StructField, ) *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Schema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Schema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Schema) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *Schema) GetFields() []StructField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Schema) GetFieldsOk() (*[]StructField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Schema) SetFields(v []StructField)`

SetFields sets Fields field to given value.


### GetSchemaId

`func (o *Schema) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *Schema) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *Schema) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *Schema) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetIdentifierFieldIds

`func (o *Schema) GetIdentifierFieldIds() []int32`

GetIdentifierFieldIds returns the IdentifierFieldIds field if non-nil, zero value otherwise.

### GetIdentifierFieldIdsOk

`func (o *Schema) GetIdentifierFieldIdsOk() (*[]int32, bool)`

GetIdentifierFieldIdsOk returns a tuple with the IdentifierFieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierFieldIds

`func (o *Schema) SetIdentifierFieldIds(v []int32)`

SetIdentifierFieldIds sets IdentifierFieldIds field to given value.

### HasIdentifierFieldIds

`func (o *Schema) HasIdentifierFieldIds() bool`

HasIdentifierFieldIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


