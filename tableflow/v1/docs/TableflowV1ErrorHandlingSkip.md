# TableflowV1ErrorHandlingSkip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The error handling mode for the Tableflow enabled topic.  In this mode, the bad records are skipped and the materialization continues with the next record.  | 

## Methods

### NewTableflowV1ErrorHandlingSkip

`func NewTableflowV1ErrorHandlingSkip(mode string, ) *TableflowV1ErrorHandlingSkip`

NewTableflowV1ErrorHandlingSkip instantiates a new TableflowV1ErrorHandlingSkip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1ErrorHandlingSkipWithDefaults

`func NewTableflowV1ErrorHandlingSkipWithDefaults() *TableflowV1ErrorHandlingSkip`

NewTableflowV1ErrorHandlingSkipWithDefaults instantiates a new TableflowV1ErrorHandlingSkip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TableflowV1ErrorHandlingSkip) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TableflowV1ErrorHandlingSkip) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TableflowV1ErrorHandlingSkip) SetMode(v string)`

SetMode sets Mode field to given value.



### AsTableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf

`func (s *TableflowV1ErrorHandlingSkip) AsTableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf() TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf`

Convenience method to wrap this instance of TableflowV1ErrorHandlingSkip in TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


