# ConnectV1ConnectorExpansionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. | 
**Type** | **string** | Type of connector, sink or source. | 
**Connector** | [**ConnectV1ConnectorExpansionStatusConnector**](ConnectV1ConnectorExpansionStatusConnector.md) |  | 
**Tasks** | Pointer to [**[]ConnectV1ConnectorExpansionStatusTasks**](ConnectV1ConnectorExpansionStatusTasks.md) | A map containing the task status. | [optional] 

## Methods

### NewConnectV1ConnectorExpansionStatus

`func NewConnectV1ConnectorExpansionStatus(name string, type_ string, connector ConnectV1ConnectorExpansionStatusConnector, ) *ConnectV1ConnectorExpansionStatus`

NewConnectV1ConnectorExpansionStatus instantiates a new ConnectV1ConnectorExpansionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionStatusWithDefaults

`func NewConnectV1ConnectorExpansionStatusWithDefaults() *ConnectV1ConnectorExpansionStatus`

NewConnectV1ConnectorExpansionStatusWithDefaults instantiates a new ConnectV1ConnectorExpansionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectV1ConnectorExpansionStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorExpansionStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorExpansionStatus) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ConnectV1ConnectorExpansionStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectV1ConnectorExpansionStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectV1ConnectorExpansionStatus) SetType(v string)`

SetType sets Type field to given value.


### GetConnector

`func (o *ConnectV1ConnectorExpansionStatus) GetConnector() ConnectV1ConnectorExpansionStatusConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectV1ConnectorExpansionStatus) GetConnectorOk() (*ConnectV1ConnectorExpansionStatusConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectV1ConnectorExpansionStatus) SetConnector(v ConnectV1ConnectorExpansionStatusConnector)`

SetConnector sets Connector field to given value.


### GetTasks

`func (o *ConnectV1ConnectorExpansionStatus) GetTasks() []ConnectV1ConnectorExpansionStatusTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ConnectV1ConnectorExpansionStatus) GetTasksOk() (*[]ConnectV1ConnectorExpansionStatusTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ConnectV1ConnectorExpansionStatus) SetTasks(v []ConnectV1ConnectorExpansionStatusTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ConnectV1ConnectorExpansionStatus) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


