# InlineResponse2001Tasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of task. | 
**State** | Pointer to **string** | The state of the connector. | 
**WorkerId** | Pointer to **string** | The worker ID of the connector. | 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse2001Tasks

`func NewInlineResponse2001Tasks(id int32, state string, workerId string, ) *InlineResponse2001Tasks`

NewInlineResponse2001Tasks instantiates a new InlineResponse2001Tasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001TasksWithDefaults

`func NewInlineResponse2001TasksWithDefaults() *InlineResponse2001Tasks`

NewInlineResponse2001TasksWithDefaults instantiates a new InlineResponse2001Tasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2001Tasks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2001Tasks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2001Tasks) SetId(v int32)`

SetId sets Id field to given value.


### GetState

`func (o *InlineResponse2001Tasks) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse2001Tasks) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse2001Tasks) SetState(v string)`

SetState sets State field to given value.


### GetWorkerId

`func (o *InlineResponse2001Tasks) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *InlineResponse2001Tasks) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *InlineResponse2001Tasks) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetMsg

`func (o *InlineResponse2001Tasks) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InlineResponse2001Tasks) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InlineResponse2001Tasks) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *InlineResponse2001Tasks) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


