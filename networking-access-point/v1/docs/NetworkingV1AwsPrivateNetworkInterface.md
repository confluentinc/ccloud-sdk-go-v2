# NetworkingV1AwsPrivateNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsPrivateNetworkInterface kind. | 
**NetworkInterfaces** | Pointer to **[]string** | List of the IDs of the Elastic Network Interfaces. | [optional] 
**Account** | Pointer to **string** | The AWS account ID associated with the ENIs you are using for the Confluent Private Network Interface. | [optional] 
**EgressRoutes** | Pointer to **[]string** | List of egress CIDRs (IPv4) for egress PNI. | [optional] 

## Methods

### NewNetworkingV1AwsPrivateNetworkInterface

`func NewNetworkingV1AwsPrivateNetworkInterface(kind string, ) *NetworkingV1AwsPrivateNetworkInterface`

NewNetworkingV1AwsPrivateNetworkInterface instantiates a new NetworkingV1AwsPrivateNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateNetworkInterfaceWithDefaults

`func NewNetworkingV1AwsPrivateNetworkInterfaceWithDefaults() *NetworkingV1AwsPrivateNetworkInterface`

NewNetworkingV1AwsPrivateNetworkInterfaceWithDefaults instantiates a new NetworkingV1AwsPrivateNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateNetworkInterface) SetKind(v string)`

SetKind sets Kind field to given value.


### GetNetworkInterfaces

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetNetworkInterfaces() []string`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetNetworkInterfacesOk() (*[]string, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NetworkingV1AwsPrivateNetworkInterface) SetNetworkInterfaces(v []string)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NetworkingV1AwsPrivateNetworkInterface) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetAccount

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1AwsPrivateNetworkInterface) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkingV1AwsPrivateNetworkInterface) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEgressRoutes

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetEgressRoutes() []string`

GetEgressRoutes returns the EgressRoutes field if non-nil, zero value otherwise.

### GetEgressRoutesOk

`func (o *NetworkingV1AwsPrivateNetworkInterface) GetEgressRoutesOk() (*[]string, bool)`

GetEgressRoutesOk returns a tuple with the EgressRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRoutes

`func (o *NetworkingV1AwsPrivateNetworkInterface) SetEgressRoutes(v []string)`

SetEgressRoutes sets EgressRoutes field to given value.

### HasEgressRoutes

`func (o *NetworkingV1AwsPrivateNetworkInterface) HasEgressRoutes() bool`

HasEgressRoutes returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AwsPrivateNetworkInterface) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateNetworkInterface in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1AwsPrivateNetworkInterface) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateNetworkInterface in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


