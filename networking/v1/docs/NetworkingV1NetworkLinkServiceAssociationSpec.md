# NetworkingV1NetworkLinkServiceAssociationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network link endpoint | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the network link endpoint | [optional] [readonly] 
**NetworkLinkEndpoint** | Pointer to **string** | ID of the Network link endpoint. | [optional] [readonly] 
**NetworkLinkService** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network_link_service to which this belongs. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1NetworkLinkServiceAssociationSpec

`func NewNetworkingV1NetworkLinkServiceAssociationSpec() *NetworkingV1NetworkLinkServiceAssociationSpec`

NewNetworkingV1NetworkLinkServiceAssociationSpec instantiates a new NetworkingV1NetworkLinkServiceAssociationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceAssociationSpecWithDefaults

`func NewNetworkingV1NetworkLinkServiceAssociationSpecWithDefaults() *NetworkingV1NetworkLinkServiceAssociationSpec`

NewNetworkingV1NetworkLinkServiceAssociationSpecWithDefaults instantiates a new NetworkingV1NetworkLinkServiceAssociationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworkLinkEndpoint

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkEndpoint() string`

GetNetworkLinkEndpoint returns the NetworkLinkEndpoint field if non-nil, zero value otherwise.

### GetNetworkLinkEndpointOk

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkEndpointOk() (*string, bool)`

GetNetworkLinkEndpointOk returns a tuple with the NetworkLinkEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkEndpoint

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetNetworkLinkEndpoint(v string)`

SetNetworkLinkEndpoint sets NetworkLinkEndpoint field to given value.

### HasNetworkLinkEndpoint

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasNetworkLinkEndpoint() bool`

HasNetworkLinkEndpoint returns a boolean if a field has been set.

### GetNetworkLinkService

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkService() EnvScopedObjectReference`

GetNetworkLinkService returns the NetworkLinkService field if non-nil, zero value otherwise.

### GetNetworkLinkServiceOk

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetNetworkLinkServiceOk() (*EnvScopedObjectReference, bool)`

GetNetworkLinkServiceOk returns a tuple with the NetworkLinkService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkService

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetNetworkLinkService(v EnvScopedObjectReference)`

SetNetworkLinkService sets NetworkLinkService field to given value.

### HasNetworkLinkService

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasNetworkLinkService() bool`

HasNetworkLinkService returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkServiceAssociationSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


