# ConnectV1ConnectorExpansionStatusTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of task. | 
**State** | **string** | The state of the task. | 
**WorkerId** | **string** | The worker ID of the task. | 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectV1ConnectorExpansionStatusTasks

`func NewConnectV1ConnectorExpansionStatusTasks(id int32, state string, workerId string, ) *ConnectV1ConnectorExpansionStatusTasks`

NewConnectV1ConnectorExpansionStatusTasks instantiates a new ConnectV1ConnectorExpansionStatusTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionStatusTasksWithDefaults

`func NewConnectV1ConnectorExpansionStatusTasksWithDefaults() *ConnectV1ConnectorExpansionStatusTasks`

NewConnectV1ConnectorExpansionStatusTasksWithDefaults instantiates a new ConnectV1ConnectorExpansionStatusTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1ConnectorExpansionStatusTasks) SetId(v int32)`

SetId sets Id field to given value.


### GetState

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectV1ConnectorExpansionStatusTasks) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ConnectV1ConnectorExpansionStatusTasks) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetMsg

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ConnectV1ConnectorExpansionStatusTasks) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ConnectV1ConnectorExpansionStatusTasks) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ConnectV1ConnectorExpansionStatusTasks) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


