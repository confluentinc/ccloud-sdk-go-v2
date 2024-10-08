# NetworkingV1AzureEgressPrivateLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AzureEgressPrivateLinkEndpointStatus kind. | 
**PrivateEndpointResourceId** | **string** | Resource ID of the Private Endpoint (if any) that is connected to the Private Link service. | [readonly] 
**PrivateEndpointDomain** | Pointer to **string** | Domain of the Private Endpoint (if any) that is connected to the Private Link service. | [optional] [readonly] 
**PrivateEndpointIpAddress** | **string** | IP address of the Private Endpoint (if any) that is connected to the Private Link service. | [readonly] 
**PrivateEndpointCustomDnsConfigDomains** | Pointer to **[]string** | Domains of the Private Endpoint (if any) based off FQDNs in Azure custom DNS configs, which are required in your private DNS setup. | [optional] [readonly] 

## Methods

### NewNetworkingV1AzureEgressPrivateLinkEndpointStatus

`func NewNetworkingV1AzureEgressPrivateLinkEndpointStatus(kind string, privateEndpointResourceId string, privateEndpointIpAddress string, ) *NetworkingV1AzureEgressPrivateLinkEndpointStatus`

NewNetworkingV1AzureEgressPrivateLinkEndpointStatus instantiates a new NetworkingV1AzureEgressPrivateLinkEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureEgressPrivateLinkEndpointStatusWithDefaults

`func NewNetworkingV1AzureEgressPrivateLinkEndpointStatusWithDefaults() *NetworkingV1AzureEgressPrivateLinkEndpointStatus`

NewNetworkingV1AzureEgressPrivateLinkEndpointStatusWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateEndpointResourceId

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.


### GetPrivateEndpointDomain

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointDomain() string`

GetPrivateEndpointDomain returns the PrivateEndpointDomain field if non-nil, zero value otherwise.

### GetPrivateEndpointDomainOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointDomainOk() (*string, bool)`

GetPrivateEndpointDomainOk returns a tuple with the PrivateEndpointDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointDomain

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointDomain(v string)`

SetPrivateEndpointDomain sets PrivateEndpointDomain field to given value.

### HasPrivateEndpointDomain

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) HasPrivateEndpointDomain() bool`

HasPrivateEndpointDomain returns a boolean if a field has been set.

### GetPrivateEndpointIpAddress

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointIpAddress() string`

GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIpAddressOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointIpAddressOk() (*string, bool)`

GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIpAddress

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointIpAddress(v string)`

SetPrivateEndpointIpAddress sets PrivateEndpointIpAddress field to given value.


### GetPrivateEndpointCustomDnsConfigDomains

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointCustomDnsConfigDomains() []string`

GetPrivateEndpointCustomDnsConfigDomains returns the PrivateEndpointCustomDnsConfigDomains field if non-nil, zero value otherwise.

### GetPrivateEndpointCustomDnsConfigDomainsOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) GetPrivateEndpointCustomDnsConfigDomainsOk() (*[]string, bool)`

GetPrivateEndpointCustomDnsConfigDomainsOk returns a tuple with the PrivateEndpointCustomDnsConfigDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointCustomDnsConfigDomains

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) SetPrivateEndpointCustomDnsConfigDomains(v []string)`

SetPrivateEndpointCustomDnsConfigDomains sets PrivateEndpointCustomDnsConfigDomains field to given value.

### HasPrivateEndpointCustomDnsConfigDomains

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) HasPrivateEndpointCustomDnsConfigDomains() bool`

HasPrivateEndpointCustomDnsConfigDomains returns a boolean if a field has been set.


### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


