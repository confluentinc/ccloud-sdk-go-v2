# NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GCP Ingress Private Service Connect Gateway Spec kind type. | 
**Region** | **string** | GCP region of the Ingress Private Service Connect Gateway. | 

## Methods

### NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpec

`func NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpec(kind string, region string, ) *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec`

NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpec instantiates a new NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpecWithDefaults

`func NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpecWithDefaults() *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec`

NewNetworkingV1GcpIngressPrivateServiceConnectGatewaySpecWithDefaults instantiates a new NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpIngressPrivateServiceConnectGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


