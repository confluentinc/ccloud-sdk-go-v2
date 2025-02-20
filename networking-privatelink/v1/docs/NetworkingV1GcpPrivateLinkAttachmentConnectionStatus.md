# NetworkingV1GcpPrivateLinkAttachmentConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnectionStatus kind. | 
**PrivateServiceConnectServiceAttachment** | **string** | GCP Private Service Connect ServiceAttachment. | [readonly] 
**PrivateServiceConnectConnectionId** | **string** | Id of the Private Service connection. | [readonly] 

## Methods

### NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatus

`func NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatus(kind string, privateServiceConnectServiceAttachment string, privateServiceConnectConnectionId string, ) *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus`

NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatus instantiates a new NetworkingV1GcpPrivateLinkAttachmentConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatusWithDefaults

`func NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatusWithDefaults() *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus`

NewNetworkingV1GcpPrivateLinkAttachmentConnectionStatusWithDefaults instantiates a new NetworkingV1GcpPrivateLinkAttachmentConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetPrivateServiceConnectServiceAttachment() string`

GetPrivateServiceConnectServiceAttachment returns the PrivateServiceConnectServiceAttachment field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetPrivateServiceConnectServiceAttachmentOk() (*string, bool)`

GetPrivateServiceConnectServiceAttachmentOk returns a tuple with the PrivateServiceConnectServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) SetPrivateServiceConnectServiceAttachment(v string)`

SetPrivateServiceConnectServiceAttachment sets PrivateServiceConnectServiceAttachment field to given value.


### GetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetPrivateServiceConnectConnectionId() string`

GetPrivateServiceConnectConnectionId returns the PrivateServiceConnectConnectionId field if non-nil, zero value otherwise.

### GetPrivateServiceConnectConnectionIdOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) GetPrivateServiceConnectConnectionIdOk() (*string, bool)`

GetPrivateServiceConnectConnectionIdOk returns a tuple with the PrivateServiceConnectConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) SetPrivateServiceConnectConnectionId(v string)`

SetPrivateServiceConnectConnectionId sets PrivateServiceConnectConnectionId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

`func (s *NetworkingV1GcpPrivateLinkAttachmentConnectionStatus) AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPrivateLinkAttachmentConnectionStatus in NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


