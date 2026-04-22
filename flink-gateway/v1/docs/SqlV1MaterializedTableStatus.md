# SqlV1MaterializedTableStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **string** | The lifecycle phase of the materialized table. | [optional] 
**Detail** | Pointer to **string** | Optional. Human-readable description of phase. | [optional] 
**Warnings** | Pointer to [**[]SqlV1MaterializedTableWarning**](SqlV1MaterializedTableWarning.md) | List of warnings encountered during materialized table execution. | [optional] 
**CreationStatement** | Pointer to **string** | Entire Materialized Table statement as submitted by user e.g CREATE OR ALTER MATERIALIZED TABLE ... | [optional] 
**ScalingStatus** | Pointer to [**SqlV1ScalingStatus**](SqlV1ScalingStatus.md) |  | [optional] 
**Version** | Pointer to **int32** | Represents the evolution history of the Materialized Table. The current value indicates the latest version. | [optional] 
**LatestVersion** | Pointer to **int32** | Represents the latest submitted version of the Materialized Table. When a query evolution is accepted, &#x60;latest_version&#x60; is incremented immediately and will be greater than &#x60;version&#x60; until the new query is fully activated. | [optional] 

## Methods

### NewSqlV1MaterializedTableStatus

`func NewSqlV1MaterializedTableStatus() *SqlV1MaterializedTableStatus`

NewSqlV1MaterializedTableStatus instantiates a new SqlV1MaterializedTableStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableStatusWithDefaults

`func NewSqlV1MaterializedTableStatusWithDefaults() *SqlV1MaterializedTableStatus`

NewSqlV1MaterializedTableStatusWithDefaults instantiates a new SqlV1MaterializedTableStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SqlV1MaterializedTableStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SqlV1MaterializedTableStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SqlV1MaterializedTableStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SqlV1MaterializedTableStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetDetail

`func (o *SqlV1MaterializedTableStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1MaterializedTableStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1MaterializedTableStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1MaterializedTableStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetWarnings

`func (o *SqlV1MaterializedTableStatus) GetWarnings() []SqlV1MaterializedTableWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SqlV1MaterializedTableStatus) GetWarningsOk() (*[]SqlV1MaterializedTableWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SqlV1MaterializedTableStatus) SetWarnings(v []SqlV1MaterializedTableWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SqlV1MaterializedTableStatus) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetCreationStatement

`func (o *SqlV1MaterializedTableStatus) GetCreationStatement() string`

GetCreationStatement returns the CreationStatement field if non-nil, zero value otherwise.

### GetCreationStatementOk

`func (o *SqlV1MaterializedTableStatus) GetCreationStatementOk() (*string, bool)`

GetCreationStatementOk returns a tuple with the CreationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatement

`func (o *SqlV1MaterializedTableStatus) SetCreationStatement(v string)`

SetCreationStatement sets CreationStatement field to given value.

### HasCreationStatement

`func (o *SqlV1MaterializedTableStatus) HasCreationStatement() bool`

HasCreationStatement returns a boolean if a field has been set.

### GetScalingStatus

`func (o *SqlV1MaterializedTableStatus) GetScalingStatus() SqlV1ScalingStatus`

GetScalingStatus returns the ScalingStatus field if non-nil, zero value otherwise.

### GetScalingStatusOk

`func (o *SqlV1MaterializedTableStatus) GetScalingStatusOk() (*SqlV1ScalingStatus, bool)`

GetScalingStatusOk returns a tuple with the ScalingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingStatus

`func (o *SqlV1MaterializedTableStatus) SetScalingStatus(v SqlV1ScalingStatus)`

SetScalingStatus sets ScalingStatus field to given value.

### HasScalingStatus

`func (o *SqlV1MaterializedTableStatus) HasScalingStatus() bool`

HasScalingStatus returns a boolean if a field has been set.

### GetVersion

`func (o *SqlV1MaterializedTableStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SqlV1MaterializedTableStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SqlV1MaterializedTableStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SqlV1MaterializedTableStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLatestVersion

`func (o *SqlV1MaterializedTableStatus) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *SqlV1MaterializedTableStatus) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *SqlV1MaterializedTableStatus) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *SqlV1MaterializedTableStatus) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


