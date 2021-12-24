# NetworkingV1AzureNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Vnet** | **string** | The Azure Virtual Network. | [readonly] 
**Subscription** | **string** | The Azure subscription. | [readonly] 
**PrivateLinkServiceAliases** | Pointer to **map[string]string** | The mapping of zones to PrivateLink Service Aliases if available.  Keys are zones and values are [Azure PrivateLink Service Aliases](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service)  | [optional] [readonly] 

## Methods

### NewNetworkingV1AzureNetwork

`func NewNetworkingV1AzureNetwork(kind string, vnet string, subscription string, ) *NetworkingV1AzureNetwork`

NewNetworkingV1AzureNetwork instantiates a new NetworkingV1AzureNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureNetworkWithDefaults

`func NewNetworkingV1AzureNetworkWithDefaults() *NetworkingV1AzureNetwork`

NewNetworkingV1AzureNetworkWithDefaults instantiates a new NetworkingV1AzureNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVnet

`func (o *NetworkingV1AzureNetwork) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkingV1AzureNetwork) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkingV1AzureNetwork) SetVnet(v string)`

SetVnet sets Vnet field to given value.


### GetSubscription

`func (o *NetworkingV1AzureNetwork) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingV1AzureNetwork) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingV1AzureNetwork) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetPrivateLinkServiceAliases

`func (o *NetworkingV1AzureNetwork) GetPrivateLinkServiceAliases() map[string]string`

GetPrivateLinkServiceAliases returns the PrivateLinkServiceAliases field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasesOk

`func (o *NetworkingV1AzureNetwork) GetPrivateLinkServiceAliasesOk() (*map[string]string, bool)`

GetPrivateLinkServiceAliasesOk returns a tuple with the PrivateLinkServiceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAliases

`func (o *NetworkingV1AzureNetwork) SetPrivateLinkServiceAliases(v map[string]string)`

SetPrivateLinkServiceAliases sets PrivateLinkServiceAliases field to given value.

### HasPrivateLinkServiceAliases

`func (o *NetworkingV1AzureNetwork) HasPrivateLinkServiceAliases() bool`

HasPrivateLinkServiceAliases returns a boolean if a field has been set.


### AsNetworkingV1NetworkStatusCloudOneOf

`func (s *NetworkingV1AzureNetwork) AsNetworkingV1NetworkStatusCloudOneOf() NetworkingV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AzureNetwork in NetworkingV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


