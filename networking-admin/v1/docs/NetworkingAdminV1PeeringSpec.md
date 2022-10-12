# NetworkingAdminV1PeeringSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the peering | [optional] 
**Cloud** | Pointer to [**NetworkingAdminV1PeeringSpecCloudOneOf**](NetworkingAdminV1PeeringSpecCloudOneOf.md) | The cloud-specific peering details. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewNetworkingAdminV1PeeringSpec

`func NewNetworkingAdminV1PeeringSpec() *NetworkingAdminV1PeeringSpec`

NewNetworkingAdminV1PeeringSpec instantiates a new NetworkingAdminV1PeeringSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1PeeringSpecWithDefaults

`func NewNetworkingAdminV1PeeringSpecWithDefaults() *NetworkingAdminV1PeeringSpec`

NewNetworkingAdminV1PeeringSpecWithDefaults instantiates a new NetworkingAdminV1PeeringSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingAdminV1PeeringSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingAdminV1PeeringSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingAdminV1PeeringSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingAdminV1PeeringSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingAdminV1PeeringSpec) GetCloud() NetworkingAdminV1PeeringSpecCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingAdminV1PeeringSpec) GetCloudOk() (*NetworkingAdminV1PeeringSpecCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingAdminV1PeeringSpec) SetCloud(v NetworkingAdminV1PeeringSpecCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingAdminV1PeeringSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingAdminV1PeeringSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingAdminV1PeeringSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingAdminV1PeeringSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingAdminV1PeeringSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingAdminV1PeeringSpec) GetNetwork() EnvScopedObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingAdminV1PeeringSpec) GetNetworkOk() (*EnvScopedObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingAdminV1PeeringSpec) SetNetwork(v EnvScopedObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingAdminV1PeeringSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


