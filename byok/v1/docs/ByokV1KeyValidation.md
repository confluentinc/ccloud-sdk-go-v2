# ByokV1KeyValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The validation phase of the key:    INITIALIZING: Initial phase for new keys awaiting first successful validation.    VALID: Last validation attempt succeeded.    INVALID: Last validation attempt failed.  | 
**Message** | Pointer to **string** | A message describing validation events.  | [optional] 
**Since** | **time.Time** | The timestamp since which the key is in the current validation phase. Changes to the validation message or phase will update this timestamp.  | 
**Region** | Pointer to **string** | The cloud region where the key is deployed. This value is computed by the API after the key is successfully validated.  | [optional] [readonly] 

## Methods

### NewByokV1KeyValidation

`func NewByokV1KeyValidation(phase string, since time.Time, ) *ByokV1KeyValidation`

NewByokV1KeyValidation instantiates a new ByokV1KeyValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyValidationWithDefaults

`func NewByokV1KeyValidationWithDefaults() *ByokV1KeyValidation`

NewByokV1KeyValidationWithDefaults instantiates a new ByokV1KeyValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *ByokV1KeyValidation) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ByokV1KeyValidation) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ByokV1KeyValidation) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetMessage

`func (o *ByokV1KeyValidation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ByokV1KeyValidation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ByokV1KeyValidation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ByokV1KeyValidation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSince

`func (o *ByokV1KeyValidation) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *ByokV1KeyValidation) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *ByokV1KeyValidation) SetSince(v time.Time)`

SetSince sets Since field to given value.


### GetRegion

`func (o *ByokV1KeyValidation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ByokV1KeyValidation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ByokV1KeyValidation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ByokV1KeyValidation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


