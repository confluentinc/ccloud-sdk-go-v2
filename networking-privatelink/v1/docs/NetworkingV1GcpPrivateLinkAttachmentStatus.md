# NetworkingV1GcpPrivateLinkAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentStatus kind. | [readonly] 
**ServiceAttachments** | [**[]NetworkingV1GcpPscServiceAttachment**](NetworkingV1GcpPscServiceAttachment.md) | Array of GCP PSC Service attachments that can be used to connect PSC Endpoints for each zone.  | [readonly] 

## Methods

### NewNetworkingV1GcpPrivateLinkAttachmentStatus

`func NewNetworkingV1GcpPrivateLinkAttachmentStatus(kind string, serviceAttachments []NetworkingV1GcpPscServiceAttachment, ) *NetworkingV1GcpPrivateLinkAttachmentStatus`

NewNetworkingV1GcpPrivateLinkAttachmentStatus instantiates a new NetworkingV1GcpPrivateLinkAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPrivateLinkAttachmentStatusWithDefaults

`func NewNetworkingV1GcpPrivateLinkAttachmentStatusWithDefaults() *NetworkingV1GcpPrivateLinkAttachmentStatus`

NewNetworkingV1GcpPrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1GcpPrivateLinkAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetServiceAttachments

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) GetServiceAttachments() []NetworkingV1GcpPscServiceAttachment`

GetServiceAttachments returns the ServiceAttachments field if non-nil, zero value otherwise.

### GetServiceAttachmentsOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) GetServiceAttachmentsOk() (*[]NetworkingV1GcpPscServiceAttachment, bool)`

GetServiceAttachmentsOk returns a tuple with the ServiceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachments

`func (o *NetworkingV1GcpPrivateLinkAttachmentStatus) SetServiceAttachments(v []NetworkingV1GcpPscServiceAttachment)`

SetServiceAttachments sets ServiceAttachments field to given value.



### AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf

`func (s *NetworkingV1GcpPrivateLinkAttachmentStatus) AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPrivateLinkAttachmentStatus in NetworkingV1PrivateLinkAttachmentStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


