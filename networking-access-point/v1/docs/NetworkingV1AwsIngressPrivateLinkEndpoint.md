# NetworkingV1AwsIngressPrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsIngressPrivateLinkEndpoint kind. | 
**VpcEndpointId** | **string** | ID of a VPC Endpoint that will be connected to the VPC Endpoint service. | 

## Methods

### NewNetworkingV1AwsIngressPrivateLinkEndpoint

`func NewNetworkingV1AwsIngressPrivateLinkEndpoint(kind string, vpcEndpointId string, ) *NetworkingV1AwsIngressPrivateLinkEndpoint`

NewNetworkingV1AwsIngressPrivateLinkEndpoint instantiates a new NetworkingV1AwsIngressPrivateLinkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsIngressPrivateLinkEndpointWithDefaults

`func NewNetworkingV1AwsIngressPrivateLinkEndpointWithDefaults() *NetworkingV1AwsIngressPrivateLinkEndpoint`

NewNetworkingV1AwsIngressPrivateLinkEndpointWithDefaults instantiates a new NetworkingV1AwsIngressPrivateLinkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointId

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *NetworkingV1AwsIngressPrivateLinkEndpoint) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.



### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AwsIngressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressPrivateLinkEndpoint in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1AwsIngressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressPrivateLinkEndpoint in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


