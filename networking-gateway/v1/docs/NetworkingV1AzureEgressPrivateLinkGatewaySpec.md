# NetworkingV1AzureEgressPrivateLinkGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Azure Egress Private Link Gateway Spec kind type. | 
**Region** | **string** | Azure region of the Egress Private Link Gateway. | 

## Methods

### NewNetworkingV1AzureEgressPrivateLinkGatewaySpec

`func NewNetworkingV1AzureEgressPrivateLinkGatewaySpec(kind string, region string, ) *NetworkingV1AzureEgressPrivateLinkGatewaySpec`

NewNetworkingV1AzureEgressPrivateLinkGatewaySpec instantiates a new NetworkingV1AzureEgressPrivateLinkGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureEgressPrivateLinkGatewaySpecWithDefaults

`func NewNetworkingV1AzureEgressPrivateLinkGatewaySpecWithDefaults() *NetworkingV1AzureEgressPrivateLinkGatewaySpec`

NewNetworkingV1AzureEgressPrivateLinkGatewaySpecWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1AzureEgressPrivateLinkGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


