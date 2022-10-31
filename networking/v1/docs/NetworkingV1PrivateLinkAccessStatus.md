# NetworkingV1PrivateLinkAccessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the PrivateLink access configuration:    PROVISIONING: PrivateLink access provisioning is in progress;    READY:  PrivateLink access is ready;    FAILED: PrivateLink access is in a failed state;    DEPROVISIONING: PrivateLink access deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if PrivateLink access is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if PrivateLink access is in a failed state | [optional] [readonly] 

## Methods

### NewNetworkingV1PrivateLinkAccessStatus

`func NewNetworkingV1PrivateLinkAccessStatus(phase string, ) *NetworkingV1PrivateLinkAccessStatus`

NewNetworkingV1PrivateLinkAccessStatus instantiates a new NetworkingV1PrivateLinkAccessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAccessStatusWithDefaults

`func NewNetworkingV1PrivateLinkAccessStatusWithDefaults() *NetworkingV1PrivateLinkAccessStatus`

NewNetworkingV1PrivateLinkAccessStatusWithDefaults instantiates a new NetworkingV1PrivateLinkAccessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1PrivateLinkAccessStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1PrivateLinkAccessStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1PrivateLinkAccessStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1PrivateLinkAccessStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1PrivateLinkAccessStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1PrivateLinkAccessStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1PrivateLinkAccessStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1PrivateLinkAccessStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1PrivateLinkAccessStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1PrivateLinkAccessStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1PrivateLinkAccessStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


