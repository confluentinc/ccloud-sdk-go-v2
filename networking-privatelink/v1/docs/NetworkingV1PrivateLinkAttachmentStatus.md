# NetworkingV1PrivateLinkAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the PrivateLink attachment:    PROVISIONING: PrivateLink attachment provisioning is in progress;    WAITING_FOR_CONNECTIONS: PrivateLink attachment is waiting for connections;    READY: PrivateLink attachment is ready;    FAILED: PrivateLink attachment is in a failed state;    EXPIRED: PrivateLink attachment has timed out waiting for connections, can only be deleted;    DEPROVISIONING: PrivateLink attachment deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if PrivateLink attachment is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if PrivateLink attachment is in a failed state. | [optional] [readonly] 
**Cloud** | Pointer to [**NetworkingV1PrivateLinkAttachmentStatusCloudOneOf**](NetworkingV1PrivateLinkAttachmentStatusCloudOneOf.md) | The cloud specific status of the PrivateLink attachment. These will be populated when the PrivateLink attachment reaches the WAITING_FOR_CONNECTIONS state. | [optional] [readonly] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentStatus

`func NewNetworkingV1PrivateLinkAttachmentStatus(phase string, ) *NetworkingV1PrivateLinkAttachmentStatus`

NewNetworkingV1PrivateLinkAttachmentStatus instantiates a new NetworkingV1PrivateLinkAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentStatusWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentStatusWithDefaults() *NetworkingV1PrivateLinkAttachmentStatus`

NewNetworkingV1PrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1PrivateLinkAttachmentStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1PrivateLinkAttachmentStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1PrivateLinkAttachmentStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetCloud() NetworkingV1PrivateLinkAttachmentStatusCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PrivateLinkAttachmentStatus) GetCloudOk() (*NetworkingV1PrivateLinkAttachmentStatusCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PrivateLinkAttachmentStatus) SetCloud(v NetworkingV1PrivateLinkAttachmentStatusCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PrivateLinkAttachmentStatus) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


