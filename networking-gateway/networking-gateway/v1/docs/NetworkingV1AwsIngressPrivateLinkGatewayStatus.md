# NetworkingV1AwsIngressPrivateLinkGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Ingress Private Link Gateway Status kind type. | 
**VpcEndpointServiceName** | Pointer to **string** | The ID of the AWS VPC Endpoint Service that can be used to establish connections for all zones. | [optional] [readonly] 

## Methods

### NewNetworkingV1AwsIngressPrivateLinkGatewayStatus

`func NewNetworkingV1AwsIngressPrivateLinkGatewayStatus(kind string, ) *NetworkingV1AwsIngressPrivateLinkGatewayStatus`

NewNetworkingV1AwsIngressPrivateLinkGatewayStatus instantiates a new NetworkingV1AwsIngressPrivateLinkGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsIngressPrivateLinkGatewayStatusWithDefaults

`func NewNetworkingV1AwsIngressPrivateLinkGatewayStatusWithDefaults() *NetworkingV1AwsIngressPrivateLinkGatewayStatus`

NewNetworkingV1AwsIngressPrivateLinkGatewayStatusWithDefaults instantiates a new NetworkingV1AwsIngressPrivateLinkGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointServiceName

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) GetVpcEndpointServiceName() string`

GetVpcEndpointServiceName returns the VpcEndpointServiceName field if non-nil, zero value otherwise.

### GetVpcEndpointServiceNameOk

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) GetVpcEndpointServiceNameOk() (*string, bool)`

GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointServiceName

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) SetVpcEndpointServiceName(v string)`

SetVpcEndpointServiceName sets VpcEndpointServiceName field to given value.

### HasVpcEndpointServiceName

`func (o *NetworkingV1AwsIngressPrivateLinkGatewayStatus) HasVpcEndpointServiceName() bool`

HasVpcEndpointServiceName returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1AwsIngressPrivateLinkGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressPrivateLinkGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


