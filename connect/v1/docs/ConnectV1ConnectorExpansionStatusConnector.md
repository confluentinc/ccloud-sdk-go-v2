# ConnectV1ConnectorExpansionStatusConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the connector. | 
**WorkerId** | **string** | The worker ID of the connector. | 
**Trace** | Pointer to **string** | Exception message in case of an error. | [optional] 

## Methods

### NewConnectV1ConnectorExpansionStatusConnector

`func NewConnectV1ConnectorExpansionStatusConnector(state string, workerId string, ) *ConnectV1ConnectorExpansionStatusConnector`

NewConnectV1ConnectorExpansionStatusConnector instantiates a new ConnectV1ConnectorExpansionStatusConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionStatusConnectorWithDefaults

`func NewConnectV1ConnectorExpansionStatusConnectorWithDefaults() *ConnectV1ConnectorExpansionStatusConnector`

NewConnectV1ConnectorExpansionStatusConnectorWithDefaults instantiates a new ConnectV1ConnectorExpansionStatusConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectV1ConnectorExpansionStatusConnector) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ConnectV1ConnectorExpansionStatusConnector) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetTrace

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ConnectV1ConnectorExpansionStatusConnector) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ConnectV1ConnectorExpansionStatusConnector) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ConnectV1ConnectorExpansionStatusConnector) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


