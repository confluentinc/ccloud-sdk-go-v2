# NetworkingV1AzurePrivateLinkAttachmentConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnectionStatus kind. | 
**PrivateLinkServiceAlias** | **string** | Azure PrivateLink service alias. | [readonly] 
**PrivateLinkServiceResourceId** | **string** | Azure PrivateLink service resource id. | [readonly] 
**PrivateEndpointResourceId** | **string** | Resource Id of the PrivateEndpoint (if any) that is connected to the PrivateLink service.  | [readonly] 

## Methods

### NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatus

`func NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatus(kind string, privateLinkServiceAlias string, privateLinkServiceResourceId string, privateEndpointResourceId string, ) *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus`

NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatus instantiates a new NetworkingV1AzurePrivateLinkAttachmentConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatusWithDefaults

`func NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatusWithDefaults() *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus`

NewNetworkingV1AzurePrivateLinkAttachmentConnectionStatusWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAttachmentConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkServiceAlias

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateLinkServiceAlias() string`

GetPrivateLinkServiceAlias returns the PrivateLinkServiceAlias field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateLinkServiceAliasOk() (*string, bool)`

GetPrivateLinkServiceAliasOk returns a tuple with the PrivateLinkServiceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAlias

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) SetPrivateLinkServiceAlias(v string)`

SetPrivateLinkServiceAlias sets PrivateLinkServiceAlias field to given value.


### GetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.


### GetPrivateEndpointResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

`func (s *NetworkingV1AzurePrivateLinkAttachmentConnectionStatus) AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AzurePrivateLinkAttachmentConnectionStatus in NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


