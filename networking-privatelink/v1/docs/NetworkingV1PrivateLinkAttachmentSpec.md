# NetworkingV1PrivateLinkAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the PrivateLink attachment. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the resources to access with the PrivateLink attachment.  | [optional] 
**Region** | Pointer to **string** | The cloud service provider region where the resources to be accessed using the PrivateLink attachment are located.  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAttachmentSpec

`func NewNetworkingV1PrivateLinkAttachmentSpec() *NetworkingV1PrivateLinkAttachmentSpec`

NewNetworkingV1PrivateLinkAttachmentSpec instantiates a new NetworkingV1PrivateLinkAttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAttachmentSpecWithDefaults

`func NewNetworkingV1PrivateLinkAttachmentSpecWithDefaults() *NetworkingV1PrivateLinkAttachmentSpec`

NewNetworkingV1PrivateLinkAttachmentSpecWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1PrivateLinkAttachmentSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PrivateLinkAttachmentSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PrivateLinkAttachmentSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1PrivateLinkAttachmentSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkingV1PrivateLinkAttachmentSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1PrivateLinkAttachmentSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1PrivateLinkAttachmentSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


