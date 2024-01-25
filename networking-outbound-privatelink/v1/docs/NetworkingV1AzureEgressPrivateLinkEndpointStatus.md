# NetworkingV1AzureEgressPrivateLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | AzureEgressPrivateLinkEndpointStatus kind. | [optional] 
**PrivateEndpointResourceId** | **string** | Resource Id of the Private Endpoint (if any) that is connected to the PrivateLink service | [readonly] 

## Methods

### NewNetworkingV1AzureEgressPrivateLinkEndpointStatus

`func NewNetworkingV1AzureEgressPrivateLinkEndpointStatus(privateEndpointResourceId string, ) *NetworkingV1AzureEgressPrivateLinkEndpointStatus`

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

### HasKind

`func (o *NetworkingV1AzureEgressPrivateLinkEndpointStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

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



### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


