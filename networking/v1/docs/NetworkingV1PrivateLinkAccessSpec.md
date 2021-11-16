# NetworkingV1PrivateLinkAccessSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the PrivateLink access | [optional] 
**Cloud** | Pointer to [**NetworkingV1PrivateLinkAccessSpecCloudOneOf**](NetworkingV1PrivateLinkAccessSpecCloudOneOf.md) | The cloud-specific PrivateLink details. | [optional] 
**Network** | Pointer to **string** | The network for the PrivateLink access | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAccessSpec

`func NewNetworkingV1PrivateLinkAccessSpec() *NetworkingV1PrivateLinkAccessSpec`

NewNetworkingV1PrivateLinkAccessSpec instantiates a new NetworkingV1PrivateLinkAccessSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAccessSpecWithDefaults

`func NewNetworkingV1PrivateLinkAccessSpecWithDefaults() *NetworkingV1PrivateLinkAccessSpec`

NewNetworkingV1PrivateLinkAccessSpecWithDefaults instantiates a new NetworkingV1PrivateLinkAccessSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1PrivateLinkAccessSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1PrivateLinkAccessSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1PrivateLinkAccessSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1PrivateLinkAccessSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1PrivateLinkAccessSpec) GetCloud() NetworkingV1PrivateLinkAccessSpecCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1PrivateLinkAccessSpec) GetCloudOk() (*NetworkingV1PrivateLinkAccessSpecCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1PrivateLinkAccessSpec) SetCloud(v NetworkingV1PrivateLinkAccessSpecCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1PrivateLinkAccessSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1PrivateLinkAccessSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1PrivateLinkAccessSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1PrivateLinkAccessSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1PrivateLinkAccessSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1PrivateLinkAccessSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1PrivateLinkAccessSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1PrivateLinkAccessSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1PrivateLinkAccessSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


