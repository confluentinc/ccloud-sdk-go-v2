# NetworkingV1AzurePrivateLinkAttachmentConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnection kind. | 
**PrivateEndpointResourceId** | **string** | Resource Id of the PrivateEndpoint that is connected to the PrivateLink service.  | 

## Methods

### NewNetworkingV1AzurePrivateLinkAttachmentConnection

`func NewNetworkingV1AzurePrivateLinkAttachmentConnection(kind string, privateEndpointResourceId string, ) *NetworkingV1AzurePrivateLinkAttachmentConnection`

NewNetworkingV1AzurePrivateLinkAttachmentConnection instantiates a new NetworkingV1AzurePrivateLinkAttachmentConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzurePrivateLinkAttachmentConnectionWithDefaults

`func NewNetworkingV1AzurePrivateLinkAttachmentConnectionWithDefaults() *NetworkingV1AzurePrivateLinkAttachmentConnection`

NewNetworkingV1AzurePrivateLinkAttachmentConnectionWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAttachmentConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateEndpointResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *NetworkingV1AzurePrivateLinkAttachmentConnection) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

`func (s *NetworkingV1AzurePrivateLinkAttachmentConnection) AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AzurePrivateLinkAttachmentConnection in NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


