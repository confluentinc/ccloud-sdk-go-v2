# SqlV1ResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]ColumnDetails**](ColumnDetails.md) | The properties of each SQL column in the schema. | [optional] 

## Methods

### NewSqlV1ResultSchema

`func NewSqlV1ResultSchema() *SqlV1ResultSchema`

NewSqlV1ResultSchema instantiates a new SqlV1ResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ResultSchemaWithDefaults

`func NewSqlV1ResultSchemaWithDefaults() *SqlV1ResultSchema`

NewSqlV1ResultSchemaWithDefaults instantiates a new SqlV1ResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *SqlV1ResultSchema) GetColumns() []ColumnDetails`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SqlV1ResultSchema) GetColumnsOk() (*[]ColumnDetails, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SqlV1ResultSchema) SetColumns(v []ColumnDetails)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SqlV1ResultSchema) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


