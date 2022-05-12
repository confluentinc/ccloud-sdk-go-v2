# InlineResponse2001Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the connector. | 
**WorkerId** | **string** | The worker ID of the connector. | 
**Trace** | Pointer to **string** | The exception name in case of error. | [optional] 

## Methods

### NewInlineResponse2001Connector

`func NewInlineResponse2001Connector(state string, workerId string, ) *InlineResponse2001Connector`

NewInlineResponse2001Connector instantiates a new InlineResponse2001Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001ConnectorWithDefaults

`func NewInlineResponse2001ConnectorWithDefaults() *InlineResponse2001Connector`

NewInlineResponse2001ConnectorWithDefaults instantiates a new InlineResponse2001Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *InlineResponse2001Connector) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse2001Connector) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse2001Connector) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *InlineResponse2001Connector) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *InlineResponse2001Connector) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *InlineResponse2001Connector) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetTrace

`func (o *InlineResponse2001Connector) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *InlineResponse2001Connector) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *InlineResponse2001Connector) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *InlineResponse2001Connector) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


