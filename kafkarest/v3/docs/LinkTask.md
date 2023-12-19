# LinkTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** |  | 
**State** | **string** |  | 
**Errors** | [**[]LinkTaskError**](LinkTaskError.md) |  | 

## Methods

### NewLinkTask

`func NewLinkTask(taskName string, state string, errors []LinkTaskError, ) *LinkTask`

NewLinkTask instantiates a new LinkTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTaskWithDefaults

`func NewLinkTaskWithDefaults() *LinkTask`

NewLinkTaskWithDefaults instantiates a new LinkTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *LinkTask) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *LinkTask) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *LinkTask) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetState

`func (o *LinkTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkTask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinkTask) SetState(v string)`

SetState sets State field to given value.


### GetErrors

`func (o *LinkTask) GetErrors() []LinkTaskError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LinkTask) GetErrorsOk() (*[]LinkTaskError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LinkTask) SetErrors(v []LinkTaskError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


