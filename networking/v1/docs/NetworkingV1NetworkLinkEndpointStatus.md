# NetworkingV1NetworkLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the network link endpoint:    PROVISIONING: network link endpoint provisioning is in progress;    PENDING_ACCEPT: network link endpoint request is pending acceptance by the the owner of the target;    READY:  network link endpoint is ready;    FAILED: network link endpoint is in a failed state;    DEPROVISIONING: network link endpoint deprovisioning is in progress;    EXPIRED: network link endpoint request is expired, can only be deleted;    DISCONNECTED: network link endpoint is in a disconnected state, target owner has removed the permissions;    DISCONNECTING: network link endpoint disconnection is in progress;    INACTIVE: network link endpoint is created, but not active since there are no clusters in the network;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if network link is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if network link is in a failed state | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The date and time when the request expires if it is not accepted by the target network admin. | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkLinkEndpointStatus

`func NewNetworkingV1NetworkLinkEndpointStatus(phase string, ) *NetworkingV1NetworkLinkEndpointStatus`

NewNetworkingV1NetworkLinkEndpointStatus instantiates a new NetworkingV1NetworkLinkEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkEndpointStatusWithDefaults

`func NewNetworkingV1NetworkLinkEndpointStatusWithDefaults() *NetworkingV1NetworkLinkEndpointStatus`

NewNetworkingV1NetworkLinkEndpointStatusWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1NetworkLinkEndpointStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1NetworkLinkEndpointStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1NetworkLinkEndpointStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1NetworkLinkEndpointStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1NetworkLinkEndpointStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExpiresAt

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NetworkingV1NetworkLinkEndpointStatus) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NetworkingV1NetworkLinkEndpointStatus) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *NetworkingV1NetworkLinkEndpointStatus) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


