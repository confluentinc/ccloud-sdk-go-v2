# RtceV1RtceTopicStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the RtceTopic:    PENDING: RtceTopic is pending initial validation;    PROVISIONING: RtceTopic infrastructure is being provisioned;    ACTIVE: RtceTopic is active and ready for use;    DELETING: RtceTopic is being deleted;    FAILED: RtceTopic provisioning failed;    UNAVAILABLE: RtceTopic is temporarily unavailable.  | [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if RtceTopic is in a failed state. | [optional] [readonly] 

## Methods

### NewRtceV1RtceTopicStatus

`func NewRtceV1RtceTopicStatus(phase string, ) *RtceV1RtceTopicStatus`

NewRtceV1RtceTopicStatus instantiates a new RtceV1RtceTopicStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicStatusWithDefaults

`func NewRtceV1RtceTopicStatusWithDefaults() *RtceV1RtceTopicStatus`

NewRtceV1RtceTopicStatusWithDefaults instantiates a new RtceV1RtceTopicStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *RtceV1RtceTopicStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *RtceV1RtceTopicStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *RtceV1RtceTopicStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorMessage

`func (o *RtceV1RtceTopicStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RtceV1RtceTopicStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RtceV1RtceTopicStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RtceV1RtceTopicStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


