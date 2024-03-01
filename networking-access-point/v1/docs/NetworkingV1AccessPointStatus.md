# NetworkingV1AccessPointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the access point:    PROVISIONING: Access point provisioning is in progress;    PENDING_ACCEPT: Access point connection request is pending acceptance by the customer;    READY:  Access point is ready;    FAILED: Access point is in a failed state;    DEPROVISIONING: Access point deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if access point is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if access point is in a failed state. | [optional] [readonly] 
**Config** | Pointer to [**NetworkingV1AccessPointStatusConfigOneOf**](NetworkingV1AccessPointStatusConfigOneOf.md) | Cloud specific status of the access point. | [optional] [readonly] 

## Methods

### NewNetworkingV1AccessPointStatus

`func NewNetworkingV1AccessPointStatus(phase string, ) *NetworkingV1AccessPointStatus`

NewNetworkingV1AccessPointStatus instantiates a new NetworkingV1AccessPointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AccessPointStatusWithDefaults

`func NewNetworkingV1AccessPointStatusWithDefaults() *NetworkingV1AccessPointStatus`

NewNetworkingV1AccessPointStatusWithDefaults instantiates a new NetworkingV1AccessPointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1AccessPointStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1AccessPointStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1AccessPointStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1AccessPointStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1AccessPointStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1AccessPointStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1AccessPointStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1AccessPointStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1AccessPointStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1AccessPointStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1AccessPointStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1AccessPointStatus) GetConfig() NetworkingV1AccessPointStatusConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1AccessPointStatus) GetConfigOk() (*NetworkingV1AccessPointStatusConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1AccessPointStatus) SetConfig(v NetworkingV1AccessPointStatusConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1AccessPointStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


