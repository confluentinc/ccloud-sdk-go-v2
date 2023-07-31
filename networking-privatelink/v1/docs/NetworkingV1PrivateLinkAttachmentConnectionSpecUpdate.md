# NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the PrivateLink attachment connection. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdate

`func NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdate() *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate`

NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdate instantiates a new NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdateWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdateWithDefaults() *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate`

NewNetworkingV1PrivateLinkAttachmentConnectionSpecUpdateWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentConnectionSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


