# NetworkingV1AwsIngressPrivateLinkGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Ingress Private Link Gateway Spec kind type. | 
**Region** | **string** | AWS region of the Ingress Private Link Gateway. | 

## Methods

### NewNetworkingV1AwsIngressPrivateLinkGatewaySpec

`func NewNetworkingV1AwsIngressPrivateLinkGatewaySpec(kind string, region string, ) *NetworkingV1AwsIngressPrivateLinkGatewaySpec`

NewNetworkingV1AwsIngressPrivateLinkGatewaySpec instantiates a new NetworkingV1AwsIngressPrivateLinkGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsIngressPrivateLinkGatewaySpecWithDefaults

`func NewNetworkingV1AwsIngressPrivateLinkGatewaySpecWithDefaults() *NetworkingV1AwsIngressPrivateLinkGatewaySpec`

NewNetworkingV1AwsIngressPrivateLinkGatewaySpecWithDefaults instantiates a new NetworkingV1AwsIngressPrivateLinkGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1AwsIngressPrivateLinkGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1AwsIngressPrivateLinkGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressPrivateLinkGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


