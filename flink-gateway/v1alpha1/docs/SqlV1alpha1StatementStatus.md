# SqlV1alpha1StatementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the submitted SQL statement: PENDING: SQL statement is pending execution; RUNNING: SQL statement execution is in progress; COMPLETED: SQL statement is completed; DELETING: SQL statement deletion is in progress; FAILING: SQL statement is failing; FAILED: SQL statement execution has failed;  | [readonly] 
**ResultSchema** | Pointer to [**SqlV1alpha1ResultSchema**](SqlV1alpha1ResultSchema.md) |  | [optional] 
**Detail** | Pointer to **string** | Description of a SQL statement phase. | [optional] [readonly] 

## Methods

### NewSqlV1alpha1StatementStatus

`func NewSqlV1alpha1StatementStatus(phase string, ) *SqlV1alpha1StatementStatus`

NewSqlV1alpha1StatementStatus instantiates a new SqlV1alpha1StatementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1alpha1StatementStatusWithDefaults

`func NewSqlV1alpha1StatementStatusWithDefaults() *SqlV1alpha1StatementStatus`

NewSqlV1alpha1StatementStatusWithDefaults instantiates a new SqlV1alpha1StatementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SqlV1alpha1StatementStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SqlV1alpha1StatementStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SqlV1alpha1StatementStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetResultSchema

`func (o *SqlV1alpha1StatementStatus) GetResultSchema() SqlV1alpha1ResultSchema`

GetResultSchema returns the ResultSchema field if non-nil, zero value otherwise.

### GetResultSchemaOk

`func (o *SqlV1alpha1StatementStatus) GetResultSchemaOk() (*SqlV1alpha1ResultSchema, bool)`

GetResultSchemaOk returns a tuple with the ResultSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSchema

`func (o *SqlV1alpha1StatementStatus) SetResultSchema(v SqlV1alpha1ResultSchema)`

SetResultSchema sets ResultSchema field to given value.

### HasResultSchema

`func (o *SqlV1alpha1StatementStatus) HasResultSchema() bool`

HasResultSchema returns a boolean if a field has been set.

### GetDetail

`func (o *SqlV1alpha1StatementStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1alpha1StatementStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1alpha1StatementStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1alpha1StatementStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


