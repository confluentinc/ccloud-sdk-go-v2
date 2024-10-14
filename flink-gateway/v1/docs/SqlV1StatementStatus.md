# SqlV1StatementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the submitted SQL statement:  PENDING: SQL statement is pending execution;  RUNNING: SQL statement execution is in progress;  COMPLETED: SQL statement is completed;  DELETING: SQL statement deletion is in progress;  FAILING: SQL statement is failing;  FAILED: SQL statement execution has failed;  STOPPED: SQL statement execution has successfully been stopped;  | [readonly] 
**ScalingStatus** | Pointer to [**SqlV1ScalingStatus**](SqlV1ScalingStatus.md) |  | [optional] 
**Detail** | Pointer to **string** | Details about the execution status of this statement. | [optional] [readonly] 
**Traits** | Pointer to [**SqlV1StatementTraits**](SqlV1StatementTraits.md) |  | [optional] 
**NetworkKind** | Pointer to **string** | The networking type used by the submitted SQL statement:  PUBLIC: SQL statement is using public networking;  PRIVATE: SQL statement is using private networking;  | [optional] [readonly] 
**LatestOffsets** | Pointer to **map[string]string** | The last Kafka offsets that a statement has processed. Represented by a mapping from Kafka topic to a string representation of partitions mapped to offsets.  | [optional] [readonly] 
**LatestOffsetsTimestamp** | Pointer to **time.Time** | The date and time at which the Kafka topic offsets were added to the statement status. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

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

### GetTraits

`func (o *SqlV1StatementStatus) GetTraits() SqlV1StatementTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *SqlV1StatementStatus) GetTraitsOk() (*SqlV1StatementTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *SqlV1StatementStatus) SetTraits(v SqlV1StatementTraits)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *SqlV1StatementStatus) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetNetworkKind

`func (o *SqlV1StatementStatus) GetNetworkKind() string`

GetNetworkKind returns the NetworkKind field if non-nil, zero value otherwise.

### GetNetworkKindOk

`func (o *SqlV1StatementStatus) GetNetworkKindOk() (*string, bool)`

GetNetworkKindOk returns a tuple with the NetworkKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkKind

`func (o *SqlV1StatementStatus) SetNetworkKind(v string)`

SetNetworkKind sets NetworkKind field to given value.

### HasNetworkKind

`func (o *SqlV1StatementStatus) HasNetworkKind() bool`

HasNetworkKind returns a boolean if a field has been set.

### GetLatestOffsets

`func (o *SqlV1StatementStatus) GetLatestOffsets() map[string]string`

GetLatestOffsets returns the LatestOffsets field if non-nil, zero value otherwise.

### GetLatestOffsetsOk

`func (o *SqlV1StatementStatus) GetLatestOffsetsOk() (*map[string]string, bool)`

GetLatestOffsetsOk returns a tuple with the LatestOffsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffsets

`func (o *SqlV1StatementStatus) SetLatestOffsets(v map[string]string)`

SetLatestOffsets sets LatestOffsets field to given value.

### HasLatestOffsets

`func (o *SqlV1StatementStatus) HasLatestOffsets() bool`

HasLatestOffsets returns a boolean if a field has been set.

### GetLatestOffsetsTimestamp

`func (o *SqlV1StatementStatus) GetLatestOffsetsTimestamp() time.Time`

GetLatestOffsetsTimestamp returns the LatestOffsetsTimestamp field if non-nil, zero value otherwise.

### GetLatestOffsetsTimestampOk

`func (o *SqlV1StatementStatus) GetLatestOffsetsTimestampOk() (*time.Time, bool)`

GetLatestOffsetsTimestampOk returns a tuple with the LatestOffsetsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffsetsTimestamp

`func (o *SqlV1StatementStatus) SetLatestOffsetsTimestamp(v time.Time)`

SetLatestOffsetsTimestamp sets LatestOffsetsTimestamp field to given value.

### HasLatestOffsetsTimestamp

`func (o *SqlV1StatementStatus) HasLatestOffsetsTimestamp() bool`

HasLatestOffsetsTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


