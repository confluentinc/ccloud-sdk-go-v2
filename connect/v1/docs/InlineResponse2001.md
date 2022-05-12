# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. | 
**Type** | **string** | Type of connector, sink or source. | 
**Connector** | [**InlineResponse2001Connector**](InlineResponse2001Connector.md) |  | 
**Tasks** | Pointer to [**[]InlineResponse2001Tasks**](InlineResponse2001Tasks.md) | The map containing the task status. | [optional] 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001(name string, type_ string, connector InlineResponse2001Connector, ) *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2001) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2001) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2001) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InlineResponse2001) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse2001) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse2001) SetType(v string)`

SetType sets Type field to given value.


### GetConnector

`func (o *InlineResponse2001) GetConnector() InlineResponse2001Connector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *InlineResponse2001) GetConnectorOk() (*InlineResponse2001Connector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *InlineResponse2001) SetConnector(v InlineResponse2001Connector)`

SetConnector sets Connector field to given value.


### GetTasks

`func (o *InlineResponse2001) GetTasks() []InlineResponse2001Tasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InlineResponse2001) GetTasksOk() (*[]InlineResponse2001Tasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InlineResponse2001) SetTasks(v []InlineResponse2001Tasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InlineResponse2001) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


