# NetworkingV1AwsEgressPrivateLinkGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Egress Private Link Gateway Spec kind type. | 
**Region** | **string** | AWS region of the Egress Private Link Gateway. | 

## Methods

### NewNetworkingV1AwsEgressPrivateLinkGatewaySpec

`func NewNetworkingV1AwsEgressPrivateLinkGatewaySpec(kind string, region string, ) *NetworkingV1AwsEgressPrivateLinkGatewaySpec`

NewNetworkingV1AwsEgressPrivateLinkGatewaySpec instantiates a new NetworkingV1AwsEgressPrivateLinkGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsEgressPrivateLinkGatewaySpecWithDefaults

`func NewNetworkingV1AwsEgressPrivateLinkGatewaySpecWithDefaults() *NetworkingV1AwsEgressPrivateLinkGatewaySpec`

NewNetworkingV1AwsEgressPrivateLinkGatewaySpecWithDefaults instantiates a new NetworkingV1AwsEgressPrivateLinkGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1AwsEgressPrivateLinkGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1AwsEgressPrivateLinkGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsEgressPrivateLinkGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


