# NetworkingV1AzureIngressPrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AzureIngressPrivateLinkEndpoint kind. | 
**PrivateEndpointResourceId** | **string** | Resource ID of a Private Endpoint that will be connected to the Private Link service. | 

## Methods

### NewNetworkingV1AzureIngressPrivateLinkEndpoint

`func NewNetworkingV1AzureIngressPrivateLinkEndpoint(kind string, privateEndpointResourceId string, ) *NetworkingV1AzureIngressPrivateLinkEndpoint`

NewNetworkingV1AzureIngressPrivateLinkEndpoint instantiates a new NetworkingV1AzureIngressPrivateLinkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureIngressPrivateLinkEndpointWithDefaults

`func NewNetworkingV1AzureIngressPrivateLinkEndpointWithDefaults() *NetworkingV1AzureIngressPrivateLinkEndpoint`

NewNetworkingV1AzureIngressPrivateLinkEndpointWithDefaults instantiates a new NetworkingV1AzureIngressPrivateLinkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateEndpointResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkEndpoint) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.



### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AzureIngressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureIngressPrivateLinkEndpoint in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1AzureIngressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureIngressPrivateLinkEndpoint in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


