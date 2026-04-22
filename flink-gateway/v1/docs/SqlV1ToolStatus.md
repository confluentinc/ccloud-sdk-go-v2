# SqlV1ToolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Describes the status of the tool:  ACTIVE: The Tool is usable;  INACTIVE: The Tool is not currently active;  ERROR: The Tool encountered an error;  | [readonly] 
**Detail** | Pointer to **string** | Details about why the tool transitioned into a given status. | [optional] [readonly] 

## Methods

### NewSqlV1ToolStatus

`func NewSqlV1ToolStatus(phase string, ) *SqlV1ToolStatus`

NewSqlV1ToolStatus instantiates a new SqlV1ToolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ToolStatusWithDefaults

`func NewSqlV1ToolStatusWithDefaults() *SqlV1ToolStatus`

NewSqlV1ToolStatusWithDefaults instantiates a new SqlV1ToolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *SqlV1ToolStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SqlV1ToolStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SqlV1ToolStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetDetail

`func (o *SqlV1ToolStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1ToolStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1ToolStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1ToolStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


