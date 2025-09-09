# TableflowV1ErrorHandlingSuspend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The error handling mode for the Tableflow enabled topic.  In this mode, the materialization of the topic is suspended in case of record failures.  | 

## Methods

### NewTableflowV1ErrorHandlingSuspend

`func NewTableflowV1ErrorHandlingSuspend(mode string, ) *TableflowV1ErrorHandlingSuspend`

NewTableflowV1ErrorHandlingSuspend instantiates a new TableflowV1ErrorHandlingSuspend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1ErrorHandlingSuspendWithDefaults

`func NewTableflowV1ErrorHandlingSuspendWithDefaults() *TableflowV1ErrorHandlingSuspend`

NewTableflowV1ErrorHandlingSuspendWithDefaults instantiates a new TableflowV1ErrorHandlingSuspend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TableflowV1ErrorHandlingSuspend) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TableflowV1ErrorHandlingSuspend) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TableflowV1ErrorHandlingSuspend) SetMode(v string)`

SetMode sets Mode field to given value.



### AsTableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf

`func (s *TableflowV1ErrorHandlingSuspend) AsTableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf() TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf`

Convenience method to wrap this instance of TableflowV1ErrorHandlingSuspend in TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


