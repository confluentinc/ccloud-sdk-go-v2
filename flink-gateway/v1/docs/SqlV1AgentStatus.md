# SqlV1AgentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **string** | Describes the status of the agent:  READY: The Agent is created;  RUNNING: The Agent is created and running in a query;  | [optional] 

## Methods

### NewSqlV1AgentStatus

`func NewSqlV1AgentStatus() *SqlV1AgentStatus`

NewSqlV1AgentStatus instantiates a new SqlV1AgentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1AgentStatusWithDefaults

`func NewSqlV1AgentStatusWithDefaults() *SqlV1AgentStatus`

NewSqlV1AgentStatusWithDefaults instantiates a new SqlV1AgentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SqlV1AgentStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SqlV1AgentStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SqlV1AgentStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SqlV1AgentStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


