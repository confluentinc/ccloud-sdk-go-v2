# NetworkingV1GcpPrivateLinkAttachmentConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnection kind. | 
**PrivateServiceConnectConnectionId** | **string** | Id of the Private Service connection. | [readonly] 

## Methods

### NewNetworkingV1GcpPrivateLinkAttachmentConnection

`func NewNetworkingV1GcpPrivateLinkAttachmentConnection(kind string, privateServiceConnectConnectionId string, ) *NetworkingV1GcpPrivateLinkAttachmentConnection`

NewNetworkingV1GcpPrivateLinkAttachmentConnection instantiates a new NetworkingV1GcpPrivateLinkAttachmentConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPrivateLinkAttachmentConnectionWithDefaults

`func NewNetworkingV1GcpPrivateLinkAttachmentConnectionWithDefaults() *NetworkingV1GcpPrivateLinkAttachmentConnection`

NewNetworkingV1GcpPrivateLinkAttachmentConnectionWithDefaults instantiates a new NetworkingV1GcpPrivateLinkAttachmentConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) GetPrivateServiceConnectConnectionId() string`

GetPrivateServiceConnectConnectionId returns the PrivateServiceConnectConnectionId field if non-nil, zero value otherwise.

### GetPrivateServiceConnectConnectionIdOk

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) GetPrivateServiceConnectConnectionIdOk() (*string, bool)`

GetPrivateServiceConnectConnectionIdOk returns a tuple with the PrivateServiceConnectConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpPrivateLinkAttachmentConnection) SetPrivateServiceConnectConnectionId(v string)`

SetPrivateServiceConnectConnectionId sets PrivateServiceConnectConnectionId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

`func (s *NetworkingV1GcpPrivateLinkAttachmentConnection) AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPrivateLinkAttachmentConnection in NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


