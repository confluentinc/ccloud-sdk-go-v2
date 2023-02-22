# NetworkingV1PrivateLinkAttachmentConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the PrivateLink attachment connection. | [optional] 
**Cloud** | Pointer to [**NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf**](NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf.md) | The cloud-specific PrivateLink attachment connection details. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**PrivateLinkAttachment** | Pointer to [**ObjectReference**](ObjectReference.md) | The private_link_attachment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentConnectionSpec

`func NewNetworkingV1PrivateLinkAttachmentConnectionSpec() *NetworkingV1PrivateLinkAttachmentConnectionSpec`

NewNetworkingV1PrivateLinkAttachmentConnectionSpec instantiates a new NetworkingV1PrivateLinkAttachmentConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentConnectionSpecWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentConnectionSpecWithDefaults() *NetworkingV1PrivateLinkAttachmentConnectionSpec`

NewNetworkingV1PrivateLinkAttachmentConnectionSpecWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetCloud() NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetCloudOk() (*NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) SetCloud(v NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPrivateLinkAttachment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetPrivateLinkAttachment() ObjectReference`

GetPrivateLinkAttachment returns the PrivateLinkAttachment field if non-nil, zero value otherwise.

### GetPrivateLinkAttachmentOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) GetPrivateLinkAttachmentOk() (*ObjectReference, bool)`

GetPrivateLinkAttachmentOk returns a tuple with the PrivateLinkAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkAttachment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) SetPrivateLinkAttachment(v ObjectReference)`

SetPrivateLinkAttachment sets PrivateLinkAttachment field to given value.

### HasPrivateLinkAttachment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpec) HasPrivateLinkAttachment() bool`

HasPrivateLinkAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


