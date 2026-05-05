# NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GcpIngressPrivateServiceConnectEndpointStatus kind. | 
**PrivateServiceConnectServiceAttachment** | **string** | URI of the Private Service Connect Service Attachment in Confluent Cloud. | [readonly] 
**PrivateServiceConnectConnectionId** | **string** | The ID of the Private Service Connect connection. | [readonly] 
**DnsDomain** | Pointer to **string** | DNS domain name used to configure the DNS Zone for the Access Point. | [optional] [readonly] 

## Methods

### NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatus

`func NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatus(kind string, privateServiceConnectServiceAttachment string, privateServiceConnectConnectionId string, ) *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus`

NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatus instantiates a new NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatusWithDefaults

`func NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatusWithDefaults() *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus`

NewNetworkingV1GcpIngressPrivateServiceConnectEndpointStatusWithDefaults instantiates a new NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectServiceAttachment() string`

GetPrivateServiceConnectServiceAttachment returns the PrivateServiceConnectServiceAttachment field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectServiceAttachmentOk() (*string, bool)`

GetPrivateServiceConnectServiceAttachmentOk returns a tuple with the PrivateServiceConnectServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) SetPrivateServiceConnectServiceAttachment(v string)`

SetPrivateServiceConnectServiceAttachment sets PrivateServiceConnectServiceAttachment field to given value.


### GetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectConnectionId() string`

GetPrivateServiceConnectConnectionId returns the PrivateServiceConnectConnectionId field if non-nil, zero value otherwise.

### GetPrivateServiceConnectConnectionIdOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectConnectionIdOk() (*string, bool)`

GetPrivateServiceConnectConnectionIdOk returns a tuple with the PrivateServiceConnectConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) SetPrivateServiceConnectConnectionId(v string)`

SetPrivateServiceConnectConnectionId sets PrivateServiceConnectConnectionId field to given value.


### GetDnsDomain

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.


### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpIngressPrivateServiceConnectEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


