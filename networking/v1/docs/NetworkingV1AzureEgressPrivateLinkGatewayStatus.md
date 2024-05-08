# NetworkingV1AzureEgressPrivateLinkGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Azure Egress Private Link Gateway Status kind type. | 
**Subscription** | Pointer to **string** | The Azure Subscription ID associated with the Confluent Cloud VPC. | [optional] [readonly] 

## Methods

### NewNetworkingV1AzureEgressPrivateLinkGatewayStatus

`func NewNetworkingV1AzureEgressPrivateLinkGatewayStatus(kind string, ) *NetworkingV1AzureEgressPrivateLinkGatewayStatus`

NewNetworkingV1AzureEgressPrivateLinkGatewayStatus instantiates a new NetworkingV1AzureEgressPrivateLinkGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureEgressPrivateLinkGatewayStatusWithDefaults

`func NewNetworkingV1AzureEgressPrivateLinkGatewayStatusWithDefaults() *NetworkingV1AzureEgressPrivateLinkGatewayStatus`

NewNetworkingV1AzureEgressPrivateLinkGatewayStatusWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubscription

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *NetworkingV1AzureEgressPrivateLinkGatewayStatus) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1AzureEgressPrivateLinkGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1AzureEgressPrivateLinkGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


