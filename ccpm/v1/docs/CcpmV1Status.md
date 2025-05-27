# CcpmV1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Phase of the Custom Connect Plugin Version. | 
**ErrorMessage** | Pointer to **string** | Displayable error message if version is in a failed state | [optional] 

## Methods

### NewCcpmV1Status

`func NewCcpmV1Status(phase string, ) *CcpmV1Status`

NewCcpmV1Status instantiates a new CcpmV1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1StatusWithDefaults

`func NewCcpmV1StatusWithDefaults() *CcpmV1Status`

NewCcpmV1StatusWithDefaults instantiates a new CcpmV1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *CcpmV1Status) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CcpmV1Status) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CcpmV1Status) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorMessage

`func (o *CcpmV1Status) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CcpmV1Status) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CcpmV1Status) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CcpmV1Status) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


