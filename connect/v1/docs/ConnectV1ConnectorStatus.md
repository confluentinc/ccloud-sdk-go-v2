# ConnectV1ConnectorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. | 
**Type** | **string** | Type of connector, sink or source. | 
**Connector** | [**InlineResponse2001Connector**](InlineResponse2001Connector.md) |  | 
**Tasks** | Pointer to [**[]InlineResponse2001Tasks**](InlineResponse2001Tasks.md) | The map containing the task status. | [optional] 
**PluginLifecycle** | Pointer to **string** | The lifecycle status of the plugin used by the connector. | [optional] 

## Methods

### NewConnectV1ConnectorStatus

`func NewConnectV1ConnectorStatus(name string, type_ string, connector InlineResponse2001Connector, ) *ConnectV1ConnectorStatus`

NewConnectV1ConnectorStatus instantiates a new ConnectV1ConnectorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorStatusWithDefaults

`func NewConnectV1ConnectorStatusWithDefaults() *ConnectV1ConnectorStatus`

NewConnectV1ConnectorStatusWithDefaults instantiates a new ConnectV1ConnectorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectV1ConnectorStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorStatus) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ConnectV1ConnectorStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectV1ConnectorStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectV1ConnectorStatus) SetType(v string)`

SetType sets Type field to given value.


### GetConnector

`func (o *ConnectV1ConnectorStatus) GetConnector() InlineResponse2001Connector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectV1ConnectorStatus) GetConnectorOk() (*InlineResponse2001Connector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectV1ConnectorStatus) SetConnector(v InlineResponse2001Connector)`

SetConnector sets Connector field to given value.


### GetTasks

`func (o *ConnectV1ConnectorStatus) GetTasks() []InlineResponse2001Tasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ConnectV1ConnectorStatus) GetTasksOk() (*[]InlineResponse2001Tasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ConnectV1ConnectorStatus) SetTasks(v []InlineResponse2001Tasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ConnectV1ConnectorStatus) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPluginLifecycle

`func (o *ConnectV1ConnectorStatus) GetPluginLifecycle() string`

GetPluginLifecycle returns the PluginLifecycle field if non-nil, zero value otherwise.

### GetPluginLifecycleOk

`func (o *ConnectV1ConnectorStatus) GetPluginLifecycleOk() (*string, bool)`

GetPluginLifecycleOk returns a tuple with the PluginLifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginLifecycle

`func (o *ConnectV1ConnectorStatus) SetPluginLifecycle(v string)`

SetPluginLifecycle sets PluginLifecycle field to given value.

### HasPluginLifecycle

`func (o *ConnectV1ConnectorStatus) HasPluginLifecycle() bool`

HasPluginLifecycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


