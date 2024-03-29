# NetworkingV1AwsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**Vpc** | **string** | The Confluent Cloud VPC ID. | [readonly] 
**Account** | **string** | The AWS account ID associated with the Confluent Cloud VPC. | [readonly] 
**PrivateLinkEndpointService** | Pointer to **string** | The endpoint service of the Confluent Cloud VPC. (used for PrivateLink) if available. | [optional] [readonly] 

## Methods

### NewNetworkingV1AwsNetwork

`func NewNetworkingV1AwsNetwork(kind string, vpc string, account string, ) *NetworkingV1AwsNetwork`

NewNetworkingV1AwsNetwork instantiates a new NetworkingV1AwsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsNetworkWithDefaults

`func NewNetworkingV1AwsNetworkWithDefaults() *NetworkingV1AwsNetwork`

NewNetworkingV1AwsNetworkWithDefaults instantiates a new NetworkingV1AwsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpc

`func (o *NetworkingV1AwsNetwork) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingV1AwsNetwork) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingV1AwsNetwork) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetAccount

`func (o *NetworkingV1AwsNetwork) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1AwsNetwork) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1AwsNetwork) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPrivateLinkEndpointService

`func (o *NetworkingV1AwsNetwork) GetPrivateLinkEndpointService() string`

GetPrivateLinkEndpointService returns the PrivateLinkEndpointService field if non-nil, zero value otherwise.

### GetPrivateLinkEndpointServiceOk

`func (o *NetworkingV1AwsNetwork) GetPrivateLinkEndpointServiceOk() (*string, bool)`

GetPrivateLinkEndpointServiceOk returns a tuple with the PrivateLinkEndpointService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkEndpointService

`func (o *NetworkingV1AwsNetwork) SetPrivateLinkEndpointService(v string)`

SetPrivateLinkEndpointService sets PrivateLinkEndpointService field to given value.

### HasPrivateLinkEndpointService

`func (o *NetworkingV1AwsNetwork) HasPrivateLinkEndpointService() bool`

HasPrivateLinkEndpointService returns a boolean if a field has been set.


### AsNetworkingV1NetworkStatusCloudOneOf

`func (s *NetworkingV1AwsNetwork) AsNetworkingV1NetworkStatusCloudOneOf() NetworkingV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsNetwork in NetworkingV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


