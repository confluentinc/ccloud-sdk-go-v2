# NetworkingV1PrivateLinkAttachmentConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the PrivateLink attachment:    PROVISIONING: PrivateLink attachment connection provisioning is in progress;    READY: PrivateLink attachment connection is ready;    FAILED: PrivateLink attachment connection is in a failed state;    DEPROVISIONING: PrivateLink attachment connection deprovisioning is in progress;    DISCONNECTED:|     PrivateLink attachment connection is in a disconnected state. This means the     private endpoint associated with this PrivateLink attachment connection has been deleted;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if PrivateLink attachment connection is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if PrivateLink attachment connection is in a failed state. | [optional] [readonly] 
**Cloud** | Pointer to [**NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf**](NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf.md) | The cloud specific status of the PrivateLink attachment connection. | [optional] [readonly] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentConnectionStatus

`func NewNetworkingV1PrivateLinkAttachmentConnectionStatus(phase string, ) *NetworkingV1PrivateLinkAttachmentConnectionStatus`

NewNetworkingV1PrivateLinkAttachmentConnectionStatus instantiates a new NetworkingV1PrivateLinkAttachmentConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentConnectionStatusWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentConnectionStatusWithDefaults() *NetworkingV1PrivateLinkAttachmentConnectionStatus`

NewNetworkingV1PrivateLinkAttachmentConnectionStatusWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetCloud() NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetCloudOk() (*NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetCloud(v NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


