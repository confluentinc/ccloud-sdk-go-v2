# NetworkingV1AzureEgressPrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AzureEgressPrivateLinkEndpoint kind. | 
**PrivateLinkServiceResourceId** | **string** | Resource ID of the Azure Private Link service. | 
**PrivateLinkSubresourceName** | Pointer to **string** | Name of the subresource for the Private Endpoint to connect to. | [optional] 
**TargetSystem** | Pointer to **string** | [Used by the Confluent Cloud Console] The target system or service that the PrivateLink Endpoint connects to (e.g. \&quot;MONGODB\&quot; or \&quot;SNOWFLAKE\&quot;). | [optional] 

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


### GetPrivateLinkSubresourceName

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkSubresourceName() string`

GetPrivateLinkSubresourceName returns the PrivateLinkSubresourceName field if non-nil, zero value otherwise.

### GetPrivateLinkSubresourceNameOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkSubresourceNameOk() (*string, bool)`

GetPrivateLinkSubresourceNameOk returns a tuple with the PrivateLinkSubresourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkSubresourceName

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetPrivateLinkSubresourceName(v string)`

SetPrivateLinkSubresourceName sets PrivateLinkSubresourceName field to given value.

### HasPrivateLinkSubresourceName

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) HasPrivateLinkSubresourceName() bool`

HasPrivateLinkSubresourceName returns a boolean if a field has been set.

### GetTargetSystem

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetTargetSystem() string`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetTargetSystemOk() (*string, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetTargetSystem(v string)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkEndpoint in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkEndpoint in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


