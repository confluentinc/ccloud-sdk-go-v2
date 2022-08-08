# NetworkingAdminV1AwsTransitGatewayAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Transit Gateway Attachment kind type. | 
**RamShareArn** | **string** | The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud attached to. | 
**TransitGatewayId** | **string** | The ID for the AWS Transit Gateway that you want Confluent CLoud attached to. | 
**TransitGatewayAttachmentId** | Pointer to **string** | The ID for the AWS Transit Gateway VPC Attachment that attaches Confluent VPC to Transit Gateway. | [optional] 
**EnableCustomRoutes** | Pointer to **bool** | Enable custom destination routes in Confluent Cloud. | [optional] [default to false]
**Routes** | Pointer to **[]string** | List of destination routes. | [optional] 

## Methods

### NewNetworkingAdminV1AwsTransitGatewayAttachment

`func NewNetworkingAdminV1AwsTransitGatewayAttachment(kind string, ramShareArn string, transitGatewayId string, ) *NetworkingAdminV1AwsTransitGatewayAttachment`

NewNetworkingAdminV1AwsTransitGatewayAttachment instantiates a new NetworkingAdminV1AwsTransitGatewayAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AwsTransitGatewayAttachmentWithDefaults

`func NewNetworkingAdminV1AwsTransitGatewayAttachmentWithDefaults() *NetworkingAdminV1AwsTransitGatewayAttachment`

NewNetworkingAdminV1AwsTransitGatewayAttachmentWithDefaults instantiates a new NetworkingAdminV1AwsTransitGatewayAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRamShareArn

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetRamShareArn() string`

GetRamShareArn returns the RamShareArn field if non-nil, zero value otherwise.

### GetRamShareArnOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetRamShareArnOk() (*string, bool)`

GetRamShareArnOk returns a tuple with the RamShareArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamShareArn

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetRamShareArn(v string)`

SetRamShareArn sets RamShareArn field to given value.


### GetTransitGatewayId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetTransitGatewayId() string`

GetTransitGatewayId returns the TransitGatewayId field if non-nil, zero value otherwise.

### GetTransitGatewayIdOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetTransitGatewayIdOk() (*string, bool)`

GetTransitGatewayIdOk returns a tuple with the TransitGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetTransitGatewayId(v string)`

SetTransitGatewayId sets TransitGatewayId field to given value.


### GetTransitGatewayAttachmentId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetTransitGatewayAttachmentId() string`

GetTransitGatewayAttachmentId returns the TransitGatewayAttachmentId field if non-nil, zero value otherwise.

### GetTransitGatewayAttachmentIdOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetTransitGatewayAttachmentIdOk() (*string, bool)`

GetTransitGatewayAttachmentIdOk returns a tuple with the TransitGatewayAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayAttachmentId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetTransitGatewayAttachmentId(v string)`

SetTransitGatewayAttachmentId sets TransitGatewayAttachmentId field to given value.

### HasTransitGatewayAttachmentId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) HasTransitGatewayAttachmentId() bool`

HasTransitGatewayAttachmentId returns a boolean if a field has been set.

### GetEnableCustomRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetEnableCustomRoutes() bool`

GetEnableCustomRoutes returns the EnableCustomRoutes field if non-nil, zero value otherwise.

### GetEnableCustomRoutesOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetEnableCustomRoutesOk() (*bool, bool)`

GetEnableCustomRoutesOk returns a tuple with the EnableCustomRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetEnableCustomRoutes(v bool)`

SetEnableCustomRoutes sets EnableCustomRoutes field to given value.

### HasEnableCustomRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) HasEnableCustomRoutes() bool`

HasEnableCustomRoutes returns a boolean if a field has been set.

### GetRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NetworkingAdminV1AwsTransitGatewayAttachment) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


### AsNetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf

`func (s *NetworkingAdminV1AwsTransitGatewayAttachment) AsNetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf() NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AwsTransitGatewayAttachment in NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


