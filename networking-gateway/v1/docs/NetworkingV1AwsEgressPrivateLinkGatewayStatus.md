# NetworkingV1AwsEgressPrivateLinkGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Egress Private Link Gateway Status kind type. | 
**PrincipalArn** | Pointer to **string** | The principal ARN used by the AWS Egress Private Link Gateway. | [optional] [readonly] 

## Methods

### NewNetworkingV1AwsEgressPrivateLinkGatewayStatus

`func NewNetworkingV1AwsEgressPrivateLinkGatewayStatus(kind string, ) *NetworkingV1AwsEgressPrivateLinkGatewayStatus`

NewNetworkingV1AwsEgressPrivateLinkGatewayStatus instantiates a new NetworkingV1AwsEgressPrivateLinkGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsEgressPrivateLinkGatewayStatusWithDefaults

`func NewNetworkingV1AwsEgressPrivateLinkGatewayStatusWithDefaults() *NetworkingV1AwsEgressPrivateLinkGatewayStatus`

NewNetworkingV1AwsEgressPrivateLinkGatewayStatusWithDefaults instantiates a new NetworkingV1AwsEgressPrivateLinkGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrincipalArn

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) GetPrincipalArn() string`

GetPrincipalArn returns the PrincipalArn field if non-nil, zero value otherwise.

### GetPrincipalArnOk

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) GetPrincipalArnOk() (*string, bool)`

GetPrincipalArnOk returns a tuple with the PrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalArn

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) SetPrincipalArn(v string)`

SetPrincipalArn sets PrincipalArn field to given value.

### HasPrincipalArn

`func (o *NetworkingV1AwsEgressPrivateLinkGatewayStatus) HasPrincipalArn() bool`

HasPrincipalArn returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1AwsEgressPrivateLinkGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1AwsEgressPrivateLinkGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


