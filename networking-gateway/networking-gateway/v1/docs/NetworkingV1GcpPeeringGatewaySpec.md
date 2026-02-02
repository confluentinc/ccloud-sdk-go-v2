# NetworkingV1GcpPeeringGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GCP Peering Gateway Spec kind type. | 
**Region** | **string** | GCP region of the Peering Gateway. | 

## Methods

### NewNetworkingV1GcpPeeringGatewaySpec

`func NewNetworkingV1GcpPeeringGatewaySpec(kind string, region string, ) *NetworkingV1GcpPeeringGatewaySpec`

NewNetworkingV1GcpPeeringGatewaySpec instantiates a new NetworkingV1GcpPeeringGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPeeringGatewaySpecWithDefaults

`func NewNetworkingV1GcpPeeringGatewaySpecWithDefaults() *NetworkingV1GcpPeeringGatewaySpec`

NewNetworkingV1GcpPeeringGatewaySpecWithDefaults instantiates a new NetworkingV1GcpPeeringGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPeeringGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPeeringGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPeeringGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1GcpPeeringGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1GcpPeeringGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1GcpPeeringGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1GcpPeeringGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPeeringGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


