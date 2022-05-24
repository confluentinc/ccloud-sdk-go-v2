# NetworkingAdminV1AzurePrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Subscription** | **string** | Azure subscription to allow for PrivateLink access. | 

## Methods

### NewNetworkingAdminV1AzurePrivateLinkAccess

`func NewNetworkingAdminV1AzurePrivateLinkAccess(kind string, subscription string, ) *NetworkingAdminV1AzurePrivateLinkAccess`

NewNetworkingAdminV1AzurePrivateLinkAccess instantiates a new NetworkingAdminV1AzurePrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AzurePrivateLinkAccessWithDefaults

`func NewNetworkingAdminV1AzurePrivateLinkAccessWithDefaults() *NetworkingAdminV1AzurePrivateLinkAccess`

NewNetworkingAdminV1AzurePrivateLinkAccessWithDefaults instantiates a new NetworkingAdminV1AzurePrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubscription

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingAdminV1AzurePrivateLinkAccess) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.



### AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingAdminV1AzurePrivateLinkAccess) AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf() NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AzurePrivateLinkAccess in NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


