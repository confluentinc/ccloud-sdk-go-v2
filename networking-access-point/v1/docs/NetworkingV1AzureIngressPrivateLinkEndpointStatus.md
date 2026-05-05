# NetworkingV1AzureIngressPrivateLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AzureIngressPrivateLinkEndpointStatus kind. | 
**PrivateLinkServiceAlias** | **string** | Alias of the Confluent Cloud Private Link Service. | [readonly] 
**PrivateLinkServiceResourceId** | **string** | Resource ID of the Confluent Cloud Private Link Service. | [readonly] 
**PrivateEndpointResourceId** | **string** | Resource ID of the Private Endpoint used for connecting to the Private Link Service. | [readonly] 
**DnsDomain** | Pointer to **string** | DNS domain name used to configure the Private DNS Zone for the Access Point. | [optional] [readonly] 

## Methods

### NewNetworkingV1AzureIngressPrivateLinkEndpointStatus

`func NewNetworkingV1AzureIngressPrivateLinkEndpointStatus(kind string, privateLinkServiceAlias string, privateLinkServiceResourceId string, privateEndpointResourceId string, ) *NetworkingV1AzureIngressPrivateLinkEndpointStatus`

NewNetworkingV1AzureIngressPrivateLinkEndpointStatus instantiates a new NetworkingV1AzureIngressPrivateLinkEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureIngressPrivateLinkEndpointStatusWithDefaults

`func NewNetworkingV1AzureIngressPrivateLinkEndpointStatusWithDefaults() *NetworkingV1AzureIngressPrivateLinkEndpointStatus`

NewNetworkingV1AzureIngressPrivateLinkEndpointStatusWithDefaults instantiates a new NetworkingV1AzureIngressPrivateLinkEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkServiceAlias

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateLinkServiceAlias() string`

GetPrivateLinkServiceAlias returns the PrivateLinkServiceAlias field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateLinkServiceAliasOk() (*string, bool)`

GetPrivateLinkServiceAliasOk returns a tuple with the PrivateLinkServiceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAlias

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) SetPrivateLinkServiceAlias(v string)`

SetPrivateLinkServiceAlias sets PrivateLinkServiceAlias field to given value.


### GetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.


### GetPrivateEndpointResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.


### GetDnsDomain

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *NetworkingV1AzureIngressPrivateLinkEndpointStatus) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.


### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1AzureIngressPrivateLinkEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureIngressPrivateLinkEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


