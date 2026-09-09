# SwitchoverV1SwitchoverEndpointCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The condition type. A &#x60;SwitchoverEndpoint&#x60; reports the single type &#x60;ResourceProvisioningComplete&#x60;. | 
**Status** | **string** | The status of the condition. | 
**Reason** | Pointer to **string** | A machine-readable reason for the condition&#39;s current status. | [optional] 
**Message** | Pointer to **string** | A human-readable message providing additional detail about the condition. | [optional] 

## Methods

### NewSwitchoverV1SwitchoverEndpointCondition

`func NewSwitchoverV1SwitchoverEndpointCondition(type_ string, status string, ) *SwitchoverV1SwitchoverEndpointCondition`

NewSwitchoverV1SwitchoverEndpointCondition instantiates a new SwitchoverV1SwitchoverEndpointCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverEndpointConditionWithDefaults

`func NewSwitchoverV1SwitchoverEndpointConditionWithDefaults() *SwitchoverV1SwitchoverEndpointCondition`

NewSwitchoverV1SwitchoverEndpointConditionWithDefaults instantiates a new SwitchoverV1SwitchoverEndpointCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchoverV1SwitchoverEndpointCondition) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchoverV1SwitchoverEndpointCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SwitchoverV1SwitchoverEndpointCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SwitchoverV1SwitchoverEndpointCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SwitchoverV1SwitchoverEndpointCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SwitchoverV1SwitchoverEndpointCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SwitchoverV1SwitchoverEndpointCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


