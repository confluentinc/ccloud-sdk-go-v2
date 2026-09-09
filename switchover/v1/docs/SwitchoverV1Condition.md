# SwitchoverV1Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | The &#x60;SwitchoverPair&#x60; member this condition belongs to. Conditions are recorded per member and flattened into a single list, so the same condition &#x60;type&#x60; can appear once per member; this field disambiguates them. Empty for pair-level conditions that have no member dimension.  | [optional] 
**Type** | **string** | The condition type — one of &#x60;ResourceValidationComplete&#x60;, &#x60;ResourcePlannedFailoverComplete&#x60;, &#x60;ResourceUnplannedFailoverComplete&#x60;, or &#x60;ResourceRestoreComplete&#x60;.  | 
**Status** | **string** | The status of the condition. | 
**Reason** | Pointer to **string** | A machine-readable reason for the condition&#39;s current status. | [optional] 
**Message** | Pointer to **string** | A human-readable message providing additional detail about the condition. | [optional] 

## Methods

### NewSwitchoverV1Condition

`func NewSwitchoverV1Condition(type_ string, status string, ) *SwitchoverV1Condition`

NewSwitchoverV1Condition instantiates a new SwitchoverV1Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1ConditionWithDefaults

`func NewSwitchoverV1ConditionWithDefaults() *SwitchoverV1Condition`

NewSwitchoverV1ConditionWithDefaults instantiates a new SwitchoverV1Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *SwitchoverV1Condition) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *SwitchoverV1Condition) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *SwitchoverV1Condition) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *SwitchoverV1Condition) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetType

`func (o *SwitchoverV1Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchoverV1Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchoverV1Condition) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SwitchoverV1Condition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchoverV1Condition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchoverV1Condition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *SwitchoverV1Condition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SwitchoverV1Condition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SwitchoverV1Condition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SwitchoverV1Condition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *SwitchoverV1Condition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SwitchoverV1Condition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SwitchoverV1Condition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SwitchoverV1Condition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


