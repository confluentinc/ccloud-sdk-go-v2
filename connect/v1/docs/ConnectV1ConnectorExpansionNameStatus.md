# ConnectV1ConnectorExpansionNameStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. | 
**Type** | **string** | Type of connector, sink or source. | 
**Connector** | [**ConnectV1ConnectorExpansionNameStatusConnector**](ConnectV1ConnectorExpansionNameStatusConnector.md) |  | 
**Tasks** | Pointer to [**[]ConnectV1ConnectorExpansionNameStatusTasks**](ConnectV1ConnectorExpansionNameStatusTasks.md) | A map containing the task status. | [optional] 

## Methods

### NewConnectV1ConnectorExpansionNameStatus

`func NewConnectV1ConnectorExpansionNameStatus(name string, type_ string, connector ConnectV1ConnectorExpansionNameStatusConnector, ) *ConnectV1ConnectorExpansionNameStatus`

NewConnectV1ConnectorExpansionNameStatus instantiates a new ConnectV1ConnectorExpansionNameStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionNameStatusWithDefaults

`func NewConnectV1ConnectorExpansionNameStatusWithDefaults() *ConnectV1ConnectorExpansionNameStatus`

NewConnectV1ConnectorExpansionNameStatusWithDefaults instantiates a new ConnectV1ConnectorExpansionNameStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectV1ConnectorExpansionNameStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorExpansionNameStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorExpansionNameStatus) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ConnectV1ConnectorExpansionNameStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectV1ConnectorExpansionNameStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectV1ConnectorExpansionNameStatus) SetType(v string)`

SetType sets Type field to given value.


### GetConnector

`func (o *ConnectV1ConnectorExpansionNameStatus) GetConnector() ConnectV1ConnectorExpansionNameStatusConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectV1ConnectorExpansionNameStatus) GetConnectorOk() (*ConnectV1ConnectorExpansionNameStatusConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectV1ConnectorExpansionNameStatus) SetConnector(v ConnectV1ConnectorExpansionNameStatusConnector)`

SetConnector sets Connector field to given value.


### GetTasks

`func (o *ConnectV1ConnectorExpansionNameStatus) GetTasks() []ConnectV1ConnectorExpansionNameStatusTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ConnectV1ConnectorExpansionNameStatus) GetTasksOk() (*[]ConnectV1ConnectorExpansionNameStatusTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ConnectV1ConnectorExpansionNameStatus) SetTasks(v []ConnectV1ConnectorExpansionNameStatusTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ConnectV1ConnectorExpansionNameStatus) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


