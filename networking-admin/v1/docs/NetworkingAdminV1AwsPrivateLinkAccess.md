# NetworkingAdminV1AwsPrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Account** | **string** | The AWS account ID for the account containing the VPCs you want to connect from using AWS PrivateLink. You can find your AWS account ID [here] (https://console.aws.amazon.com/billing/home?#/account) under **My Account** in your AWS Management Console. Must be a **12 character string**.  | 

## Methods

### NewNetworkingAdminV1AwsPrivateLinkAccess

`func NewNetworkingAdminV1AwsPrivateLinkAccess(kind string, account string, ) *NetworkingAdminV1AwsPrivateLinkAccess`

NewNetworkingAdminV1AwsPrivateLinkAccess instantiates a new NetworkingAdminV1AwsPrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AwsPrivateLinkAccessWithDefaults

`func NewNetworkingAdminV1AwsPrivateLinkAccessWithDefaults() *NetworkingAdminV1AwsPrivateLinkAccess`

NewNetworkingAdminV1AwsPrivateLinkAccessWithDefaults instantiates a new NetworkingAdminV1AwsPrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingAdminV1AwsPrivateLinkAccess) SetAccount(v string)`

SetAccount sets Account field to given value.



### AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingAdminV1AwsPrivateLinkAccess) AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf() NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AwsPrivateLinkAccess in NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


