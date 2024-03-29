# ConnectV1AlterOffsetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**ConnectV1AlterOffsetRequestInfo**](ConnectV1AlterOffsetRequestInfo.md) |  | 
**Status** | [**ConnectV1AlterOffsetStatusStatus**](ConnectV1AlterOffsetStatusStatus.md) |  | 
**PreviousOffsets** | Pointer to **[]map[string]interface{}** | Array of offsets which are categorised into partitions. | [optional] 
**AppliedAt** | Pointer to **NullableTime** | The time at which the offsets were applied. The time is in UTC, ISO 8601 format. | [optional] [readonly] 

## Methods

### NewConnectV1AlterOffsetStatus

`func NewConnectV1AlterOffsetStatus(request ConnectV1AlterOffsetRequestInfo, status ConnectV1AlterOffsetStatusStatus, ) *ConnectV1AlterOffsetStatus`

NewConnectV1AlterOffsetStatus instantiates a new ConnectV1AlterOffsetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1AlterOffsetStatusWithDefaults

`func NewConnectV1AlterOffsetStatusWithDefaults() *ConnectV1AlterOffsetStatus`

NewConnectV1AlterOffsetStatusWithDefaults instantiates a new ConnectV1AlterOffsetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *ConnectV1AlterOffsetStatus) GetRequest() ConnectV1AlterOffsetRequestInfo`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ConnectV1AlterOffsetStatus) GetRequestOk() (*ConnectV1AlterOffsetRequestInfo, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ConnectV1AlterOffsetStatus) SetRequest(v ConnectV1AlterOffsetRequestInfo)`

SetRequest sets Request field to given value.


### GetStatus

`func (o *ConnectV1AlterOffsetStatus) GetStatus() ConnectV1AlterOffsetStatusStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectV1AlterOffsetStatus) GetStatusOk() (*ConnectV1AlterOffsetStatusStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectV1AlterOffsetStatus) SetStatus(v ConnectV1AlterOffsetStatusStatus)`

SetStatus sets Status field to given value.


### GetPreviousOffsets

`func (o *ConnectV1AlterOffsetStatus) GetPreviousOffsets() []map[string]interface{}`

GetPreviousOffsets returns the PreviousOffsets field if non-nil, zero value otherwise.

### GetPreviousOffsetsOk

`func (o *ConnectV1AlterOffsetStatus) GetPreviousOffsetsOk() (*[]map[string]interface{}, bool)`

GetPreviousOffsetsOk returns a tuple with the PreviousOffsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousOffsets

`func (o *ConnectV1AlterOffsetStatus) SetPreviousOffsets(v []map[string]interface{})`

SetPreviousOffsets sets PreviousOffsets field to given value.

### HasPreviousOffsets

`func (o *ConnectV1AlterOffsetStatus) HasPreviousOffsets() bool`

HasPreviousOffsets returns a boolean if a field has been set.

### GetAppliedAt

`func (o *ConnectV1AlterOffsetStatus) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *ConnectV1AlterOffsetStatus) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *ConnectV1AlterOffsetStatus) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *ConnectV1AlterOffsetStatus) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### SetAppliedAtNil

`func (o *ConnectV1AlterOffsetStatus) SetAppliedAtNil(b bool)`

 SetAppliedAtNil sets the value for AppliedAt to be an explicit nil

### UnsetAppliedAt
`func (o *ConnectV1AlterOffsetStatus) UnsetAppliedAt()`

UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


