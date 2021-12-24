# NetworkingV1AzurePrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Subscription** | **string** | Azure subscription to allow for PrivateLink access. | 

## Methods

### NewNetworkingV1AzurePrivateLinkAccess

`func NewNetworkingV1AzurePrivateLinkAccess(kind string, subscription string, ) *NetworkingV1AzurePrivateLinkAccess`

NewNetworkingV1AzurePrivateLinkAccess instantiates a new NetworkingV1AzurePrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzurePrivateLinkAccessWithDefaults

`func NewNetworkingV1AzurePrivateLinkAccessWithDefaults() *NetworkingV1AzurePrivateLinkAccess`

NewNetworkingV1AzurePrivateLinkAccessWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzurePrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzurePrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzurePrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubscription

`func (o *NetworkingV1AzurePrivateLinkAccess) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingV1AzurePrivateLinkAccess) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingV1AzurePrivateLinkAccess) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.



### AsNetworkingV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingV1AzurePrivateLinkAccess) AsNetworkingV1PrivateLinkAccessSpecCloudOneOf() NetworkingV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AzurePrivateLinkAccess in NetworkingV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


