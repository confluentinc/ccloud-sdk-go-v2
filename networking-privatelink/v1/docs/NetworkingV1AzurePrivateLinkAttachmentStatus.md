# NetworkingV1AzurePrivateLinkAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentStatus kind. | [readonly] 
**PrivateLinkService** | [**NetworkingV1AzurePrivateLinkService**](networking.v1.AzurePrivateLinkService.md) | Azure PrivateLink service that can be used to connect to a PrivateEndpoint.  | [readonly] 

## Methods

### NewNetworkingV1AzurePrivateLinkAttachmentStatus

`func NewNetworkingV1AzurePrivateLinkAttachmentStatus(kind string, privateLinkService NetworkingV1AzurePrivateLinkService, ) *NetworkingV1AzurePrivateLinkAttachmentStatus`

NewNetworkingV1AzurePrivateLinkAttachmentStatus instantiates a new NetworkingV1AzurePrivateLinkAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzurePrivateLinkAttachmentStatusWithDefaults

`func NewNetworkingV1AzurePrivateLinkAttachmentStatusWithDefaults() *NetworkingV1AzurePrivateLinkAttachmentStatus`

NewNetworkingV1AzurePrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkService

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetPrivateLinkService() NetworkingV1AzurePrivateLinkService`

GetPrivateLinkService returns the PrivateLinkService field if non-nil, zero value otherwise.

### GetPrivateLinkServiceOk

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetPrivateLinkServiceOk() (*NetworkingV1AzurePrivateLinkService, bool)`

GetPrivateLinkServiceOk returns a tuple with the PrivateLinkService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkService

`func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) SetPrivateLinkService(v NetworkingV1AzurePrivateLinkService)`

SetPrivateLinkService sets PrivateLinkService field to given value.



### AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf

`func (s *NetworkingV1AzurePrivateLinkAttachmentStatus) AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AzurePrivateLinkAttachmentStatus in NetworkingV1PrivateLinkAttachmentStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


