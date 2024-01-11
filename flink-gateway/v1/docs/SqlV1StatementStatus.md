# SqlV1StatementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the submitted SQL statement: PENDING: SQL statement is pending execution; RUNNING: SQL statement execution is in progress; COMPLETED: SQL statement is completed; DELETING: SQL statement deletion is in progress; FAILING: SQL statement is failing; FAILED: SQL statement execution has failed; STOPPED: SQL statement execution has successfully been stopped;  | [readonly] 
**ScalingStatus** | Pointer to [**SqlV1ScalingStatus**](SqlV1ScalingStatus.md) |  | [optional] 
**ResultSchema** | Pointer to [**SqlV1ResultSchema**](SqlV1ResultSchema.md) |  | [optional] 
**Detail** | Pointer to **string** | Description of a SQL statement phase. | [optional] [readonly] 

## Methods

### NewSqlV1StatementStatus

`func NewSqlV1StatementStatus(phase string, ) *SqlV1StatementStatus`

NewSqlV1StatementStatus instantiates a new SqlV1StatementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementStatusWithDefaults

`func NewSqlV1StatementStatusWithDefaults() *SqlV1StatementStatus`

NewSqlV1StatementStatusWithDefaults instantiates a new SqlV1StatementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SqlV1StatementStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SqlV1StatementStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SqlV1StatementStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetScalingStatus

`func (o *SqlV1StatementStatus) GetScalingStatus() SqlV1ScalingStatus`

GetScalingStatus returns the ScalingStatus field if non-nil, zero value otherwise.

### GetScalingStatusOk

`func (o *SqlV1StatementStatus) GetScalingStatusOk() (*SqlV1ScalingStatus, bool)`

GetScalingStatusOk returns a tuple with the ScalingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingStatus

`func (o *SqlV1StatementStatus) SetScalingStatus(v SqlV1ScalingStatus)`

SetScalingStatus sets ScalingStatus field to given value.

### HasScalingStatus

`func (o *SqlV1StatementStatus) HasScalingStatus() bool`

HasScalingStatus returns a boolean if a field has been set.

### GetResultSchema

`func (o *SqlV1StatementStatus) GetResultSchema() SqlV1ResultSchema`

GetResultSchema returns the ResultSchema field if non-nil, zero value otherwise.

### GetResultSchemaOk

`func (o *SqlV1StatementStatus) GetResultSchemaOk() (*SqlV1ResultSchema, bool)`

GetResultSchemaOk returns a tuple with the ResultSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSchema

`func (o *SqlV1StatementStatus) SetResultSchema(v SqlV1ResultSchema)`

SetResultSchema sets ResultSchema field to given value.

### HasResultSchema

`func (o *SqlV1StatementStatus) HasResultSchema() bool`

HasResultSchema returns a boolean if a field has been set.

### GetDetail

`func (o *SqlV1StatementStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1StatementStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1StatementStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1StatementStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

