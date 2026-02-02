# NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Private Network Interface Gateway Status kind type. | 
**Account** | Pointer to **string** | The AWS account ID associated with the Private Network Interface Gateway. | [optional] [readonly] 

## Methods

### NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatus

`func NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatus(kind string, ) *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus`

NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatus instantiates a new NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatusWithDefaults

`func NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatusWithDefaults() *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus`

NewNetworkingV1AwsPrivateNetworkInterfaceGatewayStatusWithDefaults instantiates a new NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateNetworkInterfaceGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


