# NetworkingV1GcpPeeringGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GCP Peering Gateway Status kind type. | 
**PrincipalArn** | Pointer to **string** | The principal ARN used by the GCP Peering Gateway. | [optional] [readonly] 

## Methods

### NewNetworkingV1GcpPeeringGatewayStatus

`func NewNetworkingV1GcpPeeringGatewayStatus(kind string, ) *NetworkingV1GcpPeeringGatewayStatus`

NewNetworkingV1GcpPeeringGatewayStatus instantiates a new NetworkingV1GcpPeeringGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPeeringGatewayStatusWithDefaults

`func NewNetworkingV1GcpPeeringGatewayStatusWithDefaults() *NetworkingV1GcpPeeringGatewayStatus`

NewNetworkingV1GcpPeeringGatewayStatusWithDefaults instantiates a new NetworkingV1GcpPeeringGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPeeringGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPeeringGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPeeringGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrincipalArn

`func (o *NetworkingV1GcpPeeringGatewayStatus) GetPrincipalArn() string`

GetPrincipalArn returns the PrincipalArn field if non-nil, zero value otherwise.

### GetPrincipalArnOk

`func (o *NetworkingV1GcpPeeringGatewayStatus) GetPrincipalArnOk() (*string, bool)`

GetPrincipalArnOk returns a tuple with the PrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalArn

`func (o *NetworkingV1GcpPeeringGatewayStatus) SetPrincipalArn(v string)`

SetPrincipalArn sets PrincipalArn field to given value.

### HasPrincipalArn

`func (o *NetworkingV1GcpPeeringGatewayStatus) HasPrincipalArn() bool`

HasPrincipalArn returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1GcpPeeringGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPeeringGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

