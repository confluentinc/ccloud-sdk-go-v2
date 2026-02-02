# NetworkingV1AwsEgressPrivateLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsEgressPrivateLinkEndpointStatus kind. | 
**VpcEndpointId** | **string** | ID of a VPC Endpoint (if any) that is connected to the VPC Endpoint service. | [readonly] 
**VpcEndpointDnsName** | **string** | DNS name of a VPC Endpoint (if any) that is connected to the VPC Endpoint service. | [readonly] 

## Methods

### NewNetworkingV1AwsEgressPrivateLinkEndpointStatus

`func NewNetworkingV1AwsEgressPrivateLinkEndpointStatus(kind string, vpcEndpointId string, vpcEndpointDnsName string, ) *NetworkingV1AwsEgressPrivateLinkEndpointStatus`

NewNetworkingV1AwsEgressPrivateLinkEndpointStatus instantiates a new NetworkingV1AwsEgressPrivateLinkEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsEgressPrivateLinkEndpointStatusWithDefaults

`func NewNetworkingV1AwsEgressPrivateLinkEndpointStatusWithDefaults() *NetworkingV1AwsEgressPrivateLinkEndpointStatus`

NewNetworkingV1AwsEgressPrivateLinkEndpointStatusWithDefaults instantiates a new NetworkingV1AwsEgressPrivateLinkEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointId

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.


### GetVpcEndpointDnsName

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetVpcEndpointDnsName() string`

GetVpcEndpointDnsName returns the VpcEndpointDnsName field if non-nil, zero value otherwise.

### GetVpcEndpointDnsNameOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) GetVpcEndpointDnsNameOk() (*string, bool)`

GetVpcEndpointDnsNameOk returns a tuple with the VpcEndpointDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointDnsName

`func (o *NetworkingV1AwsEgressPrivateLinkEndpointStatus) SetVpcEndpointDnsName(v string)`

SetVpcEndpointDnsName sets VpcEndpointDnsName field to given value.



### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1AwsEgressPrivateLinkEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsEgressPrivateLinkEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


