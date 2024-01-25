# NetworkingV1AzureEgressPrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AzureEgressPrivateLinkEndpoint kind. | 
**PrivateLinkServiceResourceId** | **string** | Resource id of the Azure PrivateLink service. | 

## Methods

### NewNetworkingV1AzureEgressPrivateLinkEndpoint

`func NewNetworkingV1AzureEgressPrivateLinkEndpoint(kind string, privateLinkServiceResourceId string, ) *NetworkingV1AzureEgressPrivateLinkEndpoint`

NewNetworkingV1AzureEgressPrivateLinkEndpoint instantiates a new NetworkingV1AzureEgressPrivateLinkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureEgressPrivateLinkEndpointWithDefaults

`func NewNetworkingV1AzureEgressPrivateLinkEndpointWithDefaults() *NetworkingV1AzureEgressPrivateLinkEndpoint`

NewNetworkingV1AzureEgressPrivateLinkEndpointWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.



### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkEndpoint in NetworkingV1AccessPointSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


