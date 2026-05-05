# NetworkingV1AzureIngressPrivateLinkGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Azure Ingress Private Link Gateway Status kind type. | 
**PrivateLinkServiceAlias** | Pointer to **string** | Alias of the Confluent Cloud Private Link Service. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Resource ID of the Confluent Cloud Private Link Service. | [optional] [readonly] 

## Methods

### NewNetworkingV1AzureIngressPrivateLinkGatewayStatus

`func NewNetworkingV1AzureIngressPrivateLinkGatewayStatus(kind string, ) *NetworkingV1AzureIngressPrivateLinkGatewayStatus`

NewNetworkingV1AzureIngressPrivateLinkGatewayStatus instantiates a new NetworkingV1AzureIngressPrivateLinkGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzureIngressPrivateLinkGatewayStatusWithDefaults

`func NewNetworkingV1AzureIngressPrivateLinkGatewayStatusWithDefaults() *NetworkingV1AzureIngressPrivateLinkGatewayStatus`

NewNetworkingV1AzureIngressPrivateLinkGatewayStatusWithDefaults instantiates a new NetworkingV1AzureIngressPrivateLinkGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkServiceAlias

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetPrivateLinkServiceAlias() string`

GetPrivateLinkServiceAlias returns the PrivateLinkServiceAlias field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasOk

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetPrivateLinkServiceAliasOk() (*string, bool)`

GetPrivateLinkServiceAliasOk returns a tuple with the PrivateLinkServiceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAlias

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) SetPrivateLinkServiceAlias(v string)`

SetPrivateLinkServiceAlias sets PrivateLinkServiceAlias field to given value.

### HasPrivateLinkServiceAlias

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) HasPrivateLinkServiceAlias() bool`

HasPrivateLinkServiceAlias returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *NetworkingV1AzureIngressPrivateLinkGatewayStatus) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1AzureIngressPrivateLinkGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1AzureIngressPrivateLinkGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


