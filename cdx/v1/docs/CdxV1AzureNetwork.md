# CdxV1AzureNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**PrivateLinkServiceAliases** | Pointer to **map[string]string** | The mapping of zones to PrivateLink Service Aliases if available.  Keys are zones and values are [Azure PrivateLink Service Aliases](https://docs.microsoft.com/en-us/azure/private-link/private-link-service-overview#share-your-service)  | [optional] [readonly] 

## Methods

### NewCdxV1AzureNetwork

`func NewCdxV1AzureNetwork(kind string, ) *CdxV1AzureNetwork`

NewCdxV1AzureNetwork instantiates a new CdxV1AzureNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1AzureNetworkWithDefaults

`func NewCdxV1AzureNetworkWithDefaults() *CdxV1AzureNetwork`

NewCdxV1AzureNetworkWithDefaults instantiates a new CdxV1AzureNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CdxV1AzureNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1AzureNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1AzureNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkServiceAliases

`func (o *CdxV1AzureNetwork) GetPrivateLinkServiceAliases() map[string]string`

GetPrivateLinkServiceAliases returns the PrivateLinkServiceAliases field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasesOk

`func (o *CdxV1AzureNetwork) GetPrivateLinkServiceAliasesOk() (*map[string]string, bool)`

GetPrivateLinkServiceAliasesOk returns a tuple with the PrivateLinkServiceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAliases

`func (o *CdxV1AzureNetwork) SetPrivateLinkServiceAliases(v map[string]string)`

SetPrivateLinkServiceAliases sets PrivateLinkServiceAliases field to given value.

### HasPrivateLinkServiceAliases

`func (o *CdxV1AzureNetwork) HasPrivateLinkServiceAliases() bool`

HasPrivateLinkServiceAliases returns a boolean if a field has been set.


### AsCdxV1NetworkCloudOneOf

`func (s *CdxV1AzureNetwork) AsCdxV1NetworkCloudOneOf() CdxV1NetworkCloudOneOf`

Convenience method to wrap this instance of CdxV1AzureNetwork in CdxV1NetworkCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


