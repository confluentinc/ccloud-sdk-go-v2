# NetworkingAdminV1AwsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**Vpc** | **string** | The AWS VPC id for the network. | [readonly] 
**Account** | **string** | The AWS account id for the network. | [readonly] 
**PrivateLinkEndpointService** | Pointer to **string** | The AWS VPC endpoint service for the network (used for PrivateLink) if available. | [optional] [readonly] 

## Methods

### NewNetworkingAdminV1AwsNetwork

`func NewNetworkingAdminV1AwsNetwork(kind string, vpc string, account string, ) *NetworkingAdminV1AwsNetwork`

NewNetworkingAdminV1AwsNetwork instantiates a new NetworkingAdminV1AwsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AwsNetworkWithDefaults

`func NewNetworkingAdminV1AwsNetworkWithDefaults() *NetworkingAdminV1AwsNetwork`

NewNetworkingAdminV1AwsNetworkWithDefaults instantiates a new NetworkingAdminV1AwsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AwsNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AwsNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AwsNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpc

`func (o *NetworkingAdminV1AwsNetwork) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingAdminV1AwsNetwork) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingAdminV1AwsNetwork) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetAccount

`func (o *NetworkingAdminV1AwsNetwork) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingAdminV1AwsNetwork) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingAdminV1AwsNetwork) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPrivateLinkEndpointService

`func (o *NetworkingAdminV1AwsNetwork) GetPrivateLinkEndpointService() string`

GetPrivateLinkEndpointService returns the PrivateLinkEndpointService field if non-nil, zero value otherwise.

### GetPrivateLinkEndpointServiceOk

`func (o *NetworkingAdminV1AwsNetwork) GetPrivateLinkEndpointServiceOk() (*string, bool)`

GetPrivateLinkEndpointServiceOk returns a tuple with the PrivateLinkEndpointService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkEndpointService

`func (o *NetworkingAdminV1AwsNetwork) SetPrivateLinkEndpointService(v string)`

SetPrivateLinkEndpointService sets PrivateLinkEndpointService field to given value.

### HasPrivateLinkEndpointService

`func (o *NetworkingAdminV1AwsNetwork) HasPrivateLinkEndpointService() bool`

HasPrivateLinkEndpointService returns a boolean if a field has been set.


### AsNetworkingAdminV1NetworkStatusCloudOneOf

`func (s *NetworkingAdminV1AwsNetwork) AsNetworkingAdminV1NetworkStatusCloudOneOf() NetworkingAdminV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AwsNetwork in NetworkingAdminV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


