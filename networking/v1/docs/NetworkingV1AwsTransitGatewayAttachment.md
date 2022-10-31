# NetworkingV1AwsTransitGatewayAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Transit Gateway Attachment kind type. | 
**RamShareArn** | **string** | The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud to be attached to. | 
**TransitGatewayId** | **string** | The ID of the AWS Transit Gateway that you want Confluent CLoud to be attached to. | 
**EnableCustomRoutes** | Pointer to **bool** | Enable custom destination routes in Confluent Cloud. | [optional] [default to false]
**Routes** | Pointer to **[]string** | List of destination routes. | [optional] 

## Methods

### NewNetworkingV1AwsTransitGatewayAttachment

`func NewNetworkingV1AwsTransitGatewayAttachment(kind string, ramShareArn string, transitGatewayId string, ) *NetworkingV1AwsTransitGatewayAttachment`

NewNetworkingV1AwsTransitGatewayAttachment instantiates a new NetworkingV1AwsTransitGatewayAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsTransitGatewayAttachmentWithDefaults

`func NewNetworkingV1AwsTransitGatewayAttachmentWithDefaults() *NetworkingV1AwsTransitGatewayAttachment`

NewNetworkingV1AwsTransitGatewayAttachmentWithDefaults instantiates a new NetworkingV1AwsTransitGatewayAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsTransitGatewayAttachment) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRamShareArn

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetRamShareArn() string`

GetRamShareArn returns the RamShareArn field if non-nil, zero value otherwise.

### GetRamShareArnOk

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetRamShareArnOk() (*string, bool)`

GetRamShareArnOk returns a tuple with the RamShareArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamShareArn

`func (o *NetworkingV1AwsTransitGatewayAttachment) SetRamShareArn(v string)`

SetRamShareArn sets RamShareArn field to given value.


### GetTransitGatewayId

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetTransitGatewayId() string`

GetTransitGatewayId returns the TransitGatewayId field if non-nil, zero value otherwise.

### GetTransitGatewayIdOk

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetTransitGatewayIdOk() (*string, bool)`

GetTransitGatewayIdOk returns a tuple with the TransitGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayId

`func (o *NetworkingV1AwsTransitGatewayAttachment) SetTransitGatewayId(v string)`

SetTransitGatewayId sets TransitGatewayId field to given value.


### GetEnableCustomRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetEnableCustomRoutes() bool`

GetEnableCustomRoutes returns the EnableCustomRoutes field if non-nil, zero value otherwise.

### GetEnableCustomRoutesOk

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetEnableCustomRoutesOk() (*bool, bool)`

GetEnableCustomRoutesOk returns a tuple with the EnableCustomRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) SetEnableCustomRoutes(v bool)`

SetEnableCustomRoutes sets EnableCustomRoutes field to given value.

### HasEnableCustomRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) HasEnableCustomRoutes() bool`

HasEnableCustomRoutes returns a boolean if a field has been set.

### GetRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingV1AwsTransitGatewayAttachment) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NetworkingV1AwsTransitGatewayAttachment) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


### AsNetworkingV1TransitGatewayAttachmentSpecCloudOneOf

`func (s *NetworkingV1AwsTransitGatewayAttachment) AsNetworkingV1TransitGatewayAttachmentSpecCloudOneOf() NetworkingV1TransitGatewayAttachmentSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsTransitGatewayAttachment in NetworkingV1TransitGatewayAttachmentSpecCloudOneOf

### AsNetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf

`func (s *NetworkingV1AwsTransitGatewayAttachment) AsNetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf() NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsTransitGatewayAttachment in NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


