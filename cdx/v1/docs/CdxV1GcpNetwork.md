# CdxV1GcpNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**PrivateServiceConnectServiceAttachments** | Pointer to **map[string]string** | The mapping of zones to Private Service Connect Service Attachments if available. Keys are zones and values are [GCP Private Service Connect Service Attachment](https://cloud.google.com/vpc/docs/configure-private-service-connect-producer#api_7)  | [optional] [readonly] 

## Methods

### NewCdxV1GcpNetwork

`func NewCdxV1GcpNetwork(kind string, ) *CdxV1GcpNetwork`

NewCdxV1GcpNetwork instantiates a new CdxV1GcpNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1GcpNetworkWithDefaults

`func NewCdxV1GcpNetworkWithDefaults() *CdxV1GcpNetwork`

NewCdxV1GcpNetworkWithDefaults instantiates a new CdxV1GcpNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CdxV1GcpNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1GcpNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1GcpNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectServiceAttachments

`func (o *CdxV1GcpNetwork) GetPrivateServiceConnectServiceAttachments() map[string]string`

GetPrivateServiceConnectServiceAttachments returns the PrivateServiceConnectServiceAttachments field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentsOk

`func (o *CdxV1GcpNetwork) GetPrivateServiceConnectServiceAttachmentsOk() (*map[string]string, bool)`

GetPrivateServiceConnectServiceAttachmentsOk returns a tuple with the PrivateServiceConnectServiceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachments

`func (o *CdxV1GcpNetwork) SetPrivateServiceConnectServiceAttachments(v map[string]string)`

SetPrivateServiceConnectServiceAttachments sets PrivateServiceConnectServiceAttachments field to given value.

### HasPrivateServiceConnectServiceAttachments

`func (o *CdxV1GcpNetwork) HasPrivateServiceConnectServiceAttachments() bool`

HasPrivateServiceConnectServiceAttachments returns a boolean if a field has been set.


### AsCdxV1NetworkCloudOneOf

`func (s *CdxV1GcpNetwork) AsCdxV1NetworkCloudOneOf() CdxV1NetworkCloudOneOf`

Convenience method to wrap this instance of CdxV1GcpNetwork in CdxV1NetworkCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


