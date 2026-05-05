# NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GCP Ingress Private Service Connect Gateway Status kind type. | 
**PrivateServiceConnectServiceAttachment** | Pointer to **string** | URI of the Private Service Connect Service Attachment in Confluent Cloud. | [optional] [readonly] 

## Methods

### NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatus

`func NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatus(kind string, ) *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus`

NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatus instantiates a new NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatusWithDefaults

`func NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatusWithDefaults() *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus`

NewNetworkingV1GcpIngressPrivateServiceConnectGatewayStatusWithDefaults instantiates a new NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) GetPrivateServiceConnectServiceAttachment() string`

GetPrivateServiceConnectServiceAttachment returns the PrivateServiceConnectServiceAttachment field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) GetPrivateServiceConnectServiceAttachmentOk() (*string, bool)`

GetPrivateServiceConnectServiceAttachmentOk returns a tuple with the PrivateServiceConnectServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) SetPrivateServiceConnectServiceAttachment(v string)`

SetPrivateServiceConnectServiceAttachment sets PrivateServiceConnectServiceAttachment field to given value.

### HasPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) HasPrivateServiceConnectServiceAttachment() bool`

HasPrivateServiceConnectServiceAttachment returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1GcpIngressPrivateServiceConnectGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


