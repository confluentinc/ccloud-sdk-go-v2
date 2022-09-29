# NetworkingAdminV1AzurePeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Peering kind type. | 
**Tenant** | **string** | The Azure Tenant ID in which your Azure Subscription exists. Represents an organization in Azure Active Directory. You can find your Azure Tenant ID in the Azure Portal under [Azure Active Directory](https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview). Must be a valid **32 character UUID string**.  | 
**Vnet** | **string** | The resource ID of the VNet that you are peering with Confluent Cloud. You can find the name of your Azure VNet in the [Azure Portal on the Overview tab of your Azure Virtual Network](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Network%2FvirtualNetworks). | 
**CustomerRegion** | **string** | The region of the VNet you are peering with Confluent Cloud network. | 

## Methods

### NewNetworkingAdminV1AzurePeering

`func NewNetworkingAdminV1AzurePeering(kind string, tenant string, vnet string, customerRegion string, ) *NetworkingAdminV1AzurePeering`

NewNetworkingAdminV1AzurePeering instantiates a new NetworkingAdminV1AzurePeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AzurePeeringWithDefaults

`func NewNetworkingAdminV1AzurePeeringWithDefaults() *NetworkingAdminV1AzurePeering`

NewNetworkingAdminV1AzurePeeringWithDefaults instantiates a new NetworkingAdminV1AzurePeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AzurePeering) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AzurePeering) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AzurePeering) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTenant

`func (o *NetworkingAdminV1AzurePeering) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *NetworkingAdminV1AzurePeering) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *NetworkingAdminV1AzurePeering) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetVnet

`func (o *NetworkingAdminV1AzurePeering) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkingAdminV1AzurePeering) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkingAdminV1AzurePeering) SetVnet(v string)`

SetVnet sets Vnet field to given value.


### GetCustomerRegion

`func (o *NetworkingAdminV1AzurePeering) GetCustomerRegion() string`

GetCustomerRegion returns the CustomerRegion field if non-nil, zero value otherwise.

### GetCustomerRegionOk

`func (o *NetworkingAdminV1AzurePeering) GetCustomerRegionOk() (*string, bool)`

GetCustomerRegionOk returns a tuple with the CustomerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRegion

`func (o *NetworkingAdminV1AzurePeering) SetCustomerRegion(v string)`

SetCustomerRegion sets CustomerRegion field to given value.



### AsNetworkingAdminV1PeeringSpecCloudOneOf

`func (s *NetworkingAdminV1AzurePeering) AsNetworkingAdminV1PeeringSpecCloudOneOf() NetworkingAdminV1PeeringSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AzurePeering in NetworkingAdminV1PeeringSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


