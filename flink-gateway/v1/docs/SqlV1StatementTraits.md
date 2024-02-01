# SqlV1StatementTraits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqlKind** | Pointer to **string** | Categorizes the SQL statement. The result is Confluent-specific but inspired by SQL. It uses underscores for separating concepts e.g. \&quot;CREATE_TABLE\&quot;. | [optional] 
**IsBounded** | Pointer to **bool** | Indicates the special case where results of a statement are bounded. | [optional] 
**IsAppendOnly** | Pointer to **bool** | Indicates the special case where results of a statement are insert/append only. | [optional] 
**UpsertColumns** | Pointer to **[]int32** | Defines the column indices clients can use as upsert keys. | [optional] 
**Schema** | Pointer to [**SqlV1ResultSchema**](SqlV1ResultSchema.md) |  | [optional] 

## Methods

### NewSqlV1StatementTraits

`func NewSqlV1StatementTraits() *SqlV1StatementTraits`

NewSqlV1StatementTraits instantiates a new SqlV1StatementTraits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementTraitsWithDefaults

`func NewSqlV1StatementTraitsWithDefaults() *SqlV1StatementTraits`

NewSqlV1StatementTraitsWithDefaults instantiates a new SqlV1StatementTraits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqlKind

`func (o *SqlV1StatementTraits) GetSqlKind() string`

GetSqlKind returns the SqlKind field if non-nil, zero value otherwise.

### GetSqlKindOk

`func (o *SqlV1StatementTraits) GetSqlKindOk() (*string, bool)`

GetSqlKindOk returns a tuple with the SqlKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlKind

`func (o *SqlV1StatementTraits) SetSqlKind(v string)`

SetSqlKind sets SqlKind field to given value.

### HasSqlKind

`func (o *SqlV1StatementTraits) HasSqlKind() bool`

HasSqlKind returns a boolean if a field has been set.

### GetIsBounded

`func (o *SqlV1StatementTraits) GetIsBounded() bool`

GetIsBounded returns the IsBounded field if non-nil, zero value otherwise.

### GetIsBoundedOk

`func (o *SqlV1StatementTraits) GetIsBoundedOk() (*bool, bool)`

GetIsBoundedOk returns a tuple with the IsBounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounded

`func (o *SqlV1StatementTraits) SetIsBounded(v bool)`

SetIsBounded sets IsBounded field to given value.

### HasIsBounded

`func (o *SqlV1StatementTraits) HasIsBounded() bool`

HasIsBounded returns a boolean if a field has been set.

### GetIsAppendOnly

`func (o *SqlV1StatementTraits) GetIsAppendOnly() bool`

GetIsAppendOnly returns the IsAppendOnly field if non-nil, zero value otherwise.

### GetIsAppendOnlyOk

`func (o *SqlV1StatementTraits) GetIsAppendOnlyOk() (*bool, bool)`

GetIsAppendOnlyOk returns a tuple with the IsAppendOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppendOnly

`func (o *SqlV1StatementTraits) SetIsAppendOnly(v bool)`

SetIsAppendOnly sets IsAppendOnly field to given value.

### HasIsAppendOnly

`func (o *SqlV1StatementTraits) HasIsAppendOnly() bool`

HasIsAppendOnly returns a boolean if a field has been set.

### GetUpsertColumns

`func (o *SqlV1StatementTraits) GetUpsertColumns() []int32`

GetUpsertColumns returns the UpsertColumns field if non-nil, zero value otherwise.

### GetUpsertColumnsOk

`func (o *SqlV1StatementTraits) GetUpsertColumnsOk() (*[]int32, bool)`

GetUpsertColumnsOk returns a tuple with the UpsertColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertColumns

`func (o *SqlV1StatementTraits) SetUpsertColumns(v []int32)`

SetUpsertColumns sets UpsertColumns field to given value.

### HasUpsertColumns

`func (o *SqlV1StatementTraits) HasUpsertColumns() bool`

HasUpsertColumns returns a boolean if a field has been set.

### GetSchema

`func (o *SqlV1StatementTraits) GetSchema() SqlV1ResultSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SqlV1StatementTraits) GetSchemaOk() (*SqlV1ResultSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SqlV1StatementTraits) SetSchema(v SqlV1ResultSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SqlV1StatementTraits) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


