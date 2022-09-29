# NetworkingAdminV1AzureNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**Vnet** | **string** | The resource ID of the Confluent Cloud VNet. | [readonly] 
**Subscription** | **string** | The Azure Subscription ID associated with the Confluent Cloud VPC. | [readonly] 
**PrivateLinkServiceAliases** | Pointer to **map[string]string** | The mapping of zones to Private Link Service Aliases if available. Keys are zones and values are [Azure Private Link Service Aliases](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service).  | [optional] [readonly] 
**PrivateLinkServiceResourceIds** | Pointer to **map[string]string** | The mapping of zones to Private Link Service Resource IDs if available. Keys are zones and values are [Azure Private Link Service Resource IDs](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service).  | [optional] [readonly] 

## Methods

### NewNetworkingAdminV1AzureNetwork

`func NewNetworkingAdminV1AzureNetwork(kind string, vnet string, subscription string, ) *NetworkingAdminV1AzureNetwork`

NewNetworkingAdminV1AzureNetwork instantiates a new NetworkingAdminV1AzureNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AzureNetworkWithDefaults

`func NewNetworkingAdminV1AzureNetworkWithDefaults() *NetworkingAdminV1AzureNetwork`

NewNetworkingAdminV1AzureNetworkWithDefaults instantiates a new NetworkingAdminV1AzureNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AzureNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AzureNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AzureNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVnet

`func (o *NetworkingAdminV1AzureNetwork) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkingAdminV1AzureNetwork) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkingAdminV1AzureNetwork) SetVnet(v string)`

SetVnet sets Vnet field to given value.


### GetSubscription

`func (o *NetworkingAdminV1AzureNetwork) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingAdminV1AzureNetwork) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingAdminV1AzureNetwork) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetPrivateLinkServiceAliases

`func (o *NetworkingAdminV1AzureNetwork) GetPrivateLinkServiceAliases() map[string]string`

GetPrivateLinkServiceAliases returns the PrivateLinkServiceAliases field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasesOk

`func (o *NetworkingAdminV1AzureNetwork) GetPrivateLinkServiceAliasesOk() (*map[string]string, bool)`

GetPrivateLinkServiceAliasesOk returns a tuple with the PrivateLinkServiceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAliases

`func (o *NetworkingAdminV1AzureNetwork) SetPrivateLinkServiceAliases(v map[string]string)`

SetPrivateLinkServiceAliases sets PrivateLinkServiceAliases field to given value.

### HasPrivateLinkServiceAliases

`func (o *NetworkingAdminV1AzureNetwork) HasPrivateLinkServiceAliases() bool`

HasPrivateLinkServiceAliases returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceIds

`func (o *NetworkingAdminV1AzureNetwork) GetPrivateLinkServiceResourceIds() map[string]string`

GetPrivateLinkServiceResourceIds returns the PrivateLinkServiceResourceIds field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdsOk

`func (o *NetworkingAdminV1AzureNetwork) GetPrivateLinkServiceResourceIdsOk() (*map[string]string, bool)`

GetPrivateLinkServiceResourceIdsOk returns a tuple with the PrivateLinkServiceResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceIds

`func (o *NetworkingAdminV1AzureNetwork) SetPrivateLinkServiceResourceIds(v map[string]string)`

SetPrivateLinkServiceResourceIds sets PrivateLinkServiceResourceIds field to given value.

### HasPrivateLinkServiceResourceIds

`func (o *NetworkingAdminV1AzureNetwork) HasPrivateLinkServiceResourceIds() bool`

HasPrivateLinkServiceResourceIds returns a boolean if a field has been set.


### AsNetworkingAdminV1NetworkStatusCloudOneOf

`func (s *NetworkingAdminV1AzureNetwork) AsNetworkingAdminV1NetworkStatusCloudOneOf() NetworkingAdminV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AzureNetwork in NetworkingAdminV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


