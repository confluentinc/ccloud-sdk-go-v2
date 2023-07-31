# NetworkingV1GcpPscServiceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | Zone associated with the PSC Service attachment. | [readonly] 
**PrivateServiceConnectServiceAttachment** | **string** | Id of a Private Service Connect Service Attachment in Confluent Cloud. | [readonly] 

## Methods

### NewNetworkingV1GcpPscServiceAttachment

`func NewNetworkingV1GcpPscServiceAttachment(zone string, privateServiceConnectServiceAttachment string, ) *NetworkingV1GcpPscServiceAttachment`

NewNetworkingV1GcpPscServiceAttachment instantiates a new NetworkingV1GcpPscServiceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPscServiceAttachmentWithDefaults

`func NewNetworkingV1GcpPscServiceAttachmentWithDefaults() *NetworkingV1GcpPscServiceAttachment`

NewNetworkingV1GcpPscServiceAttachmentWithDefaults instantiates a new NetworkingV1GcpPscServiceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *NetworkingV1GcpPscServiceAttachment) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NetworkingV1GcpPscServiceAttachment) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NetworkingV1GcpPscServiceAttachment) SetZone(v string)`

SetZone sets Zone field to given value.


### GetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpPscServiceAttachment) GetPrivateServiceConnectServiceAttachment() string`

GetPrivateServiceConnectServiceAttachment returns the PrivateServiceConnectServiceAttachment field if non-nil, zero value otherwise.

### GetPrivateServiceConnectServiceAttachmentOk

`func (o *NetworkingV1GcpPscServiceAttachment) GetPrivateServiceConnectServiceAttachmentOk() (*string, bool)`

GetPrivateServiceConnectServiceAttachmentOk returns a tuple with the PrivateServiceConnectServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectServiceAttachment

`func (o *NetworkingV1GcpPscServiceAttachment) SetPrivateServiceConnectServiceAttachment(v string)`

SetPrivateServiceConnectServiceAttachment sets PrivateServiceConnectServiceAttachment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


