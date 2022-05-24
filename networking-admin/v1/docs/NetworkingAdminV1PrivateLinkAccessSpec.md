# NetworkingAdminV1PrivateLinkAccessSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the PrivateLink access | [optional] 
**Cloud** | Pointer to [**NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf**](NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf.md) | The cloud-specific PrivateLink details. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**ObjectReference**](ObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewNetworkingAdminV1PrivateLinkAccessSpec

`func NewNetworkingAdminV1PrivateLinkAccessSpec() *NetworkingAdminV1PrivateLinkAccessSpec`

NewNetworkingAdminV1PrivateLinkAccessSpec instantiates a new NetworkingAdminV1PrivateLinkAccessSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1PrivateLinkAccessSpecWithDefaults

`func NewNetworkingAdminV1PrivateLinkAccessSpecWithDefaults() *NetworkingAdminV1PrivateLinkAccessSpec`

NewNetworkingAdminV1PrivateLinkAccessSpecWithDefaults instantiates a new NetworkingAdminV1PrivateLinkAccessSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetCloud() NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetCloudOk() (*NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) SetCloud(v NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetNetwork() ObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) GetNetworkOk() (*ObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) SetNetwork(v ObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingAdminV1PrivateLinkAccessSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


