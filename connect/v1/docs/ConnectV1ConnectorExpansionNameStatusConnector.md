# ConnectV1ConnectorExpansionNameStatusConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the connector. | 
**WorkerId** | **string** | The worker ID of the connector. | 
**Trace** | Pointer to **string** | Exception message in case of an error. | [optional] 

## Methods

### NewConnectV1ConnectorExpansionNameStatusConnector

`func NewConnectV1ConnectorExpansionNameStatusConnector(state string, workerId string, ) *ConnectV1ConnectorExpansionNameStatusConnector`

NewConnectV1ConnectorExpansionNameStatusConnector instantiates a new ConnectV1ConnectorExpansionNameStatusConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionNameStatusConnectorWithDefaults

`func NewConnectV1ConnectorExpansionNameStatusConnectorWithDefaults() *ConnectV1ConnectorExpansionNameStatusConnector`

NewConnectV1ConnectorExpansionNameStatusConnectorWithDefaults instantiates a new ConnectV1ConnectorExpansionNameStatusConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetTrace

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ConnectV1ConnectorExpansionNameStatusConnector) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


