# NetworkingV1AwsPrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Account** | **string** | AWS account to allow for PrivateLink access. | 

## Methods

### NewNetworkingV1AwsPrivateLinkAccess

`func NewNetworkingV1AwsPrivateLinkAccess(kind string, account string, ) *NetworkingV1AwsPrivateLinkAccess`

NewNetworkingV1AwsPrivateLinkAccess instantiates a new NetworkingV1AwsPrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateLinkAccessWithDefaults

`func NewNetworkingV1AwsPrivateLinkAccessWithDefaults() *NetworkingV1AwsPrivateLinkAccess`

NewNetworkingV1AwsPrivateLinkAccessWithDefaults instantiates a new NetworkingV1AwsPrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingV1AwsPrivateLinkAccess) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1AwsPrivateLinkAccess) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1AwsPrivateLinkAccess) SetAccount(v string)`

SetAccount sets Account field to given value.



### AsNetworkingV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingV1AwsPrivateLinkAccess) AsNetworkingV1PrivateLinkAccessSpecCloudOneOf() NetworkingV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateLinkAccess in NetworkingV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


