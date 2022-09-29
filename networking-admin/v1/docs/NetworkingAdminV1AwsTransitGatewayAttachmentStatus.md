# NetworkingAdminV1AwsTransitGatewayAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | AWS Transit Gateway Attachment Status kind type. | [optional] 
**TransitGatewayAttachmentId** | **string** | The ID of the AWS Transit Gateway VPC Attachment that attaches Confluent VPC to Transit Gateway. | [readonly] 

## Methods

### NewNetworkingAdminV1AwsTransitGatewayAttachmentStatus

`func NewNetworkingAdminV1AwsTransitGatewayAttachmentStatus(transitGatewayAttachmentId string, ) *NetworkingAdminV1AwsTransitGatewayAttachmentStatus`

NewNetworkingAdminV1AwsTransitGatewayAttachmentStatus instantiates a new NetworkingAdminV1AwsTransitGatewayAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AwsTransitGatewayAttachmentStatusWithDefaults

`func NewNetworkingAdminV1AwsTransitGatewayAttachmentStatusWithDefaults() *NetworkingAdminV1AwsTransitGatewayAttachmentStatus`

NewNetworkingAdminV1AwsTransitGatewayAttachmentStatusWithDefaults instantiates a new NetworkingAdminV1AwsTransitGatewayAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTransitGatewayAttachmentId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) GetTransitGatewayAttachmentId() string`

GetTransitGatewayAttachmentId returns the TransitGatewayAttachmentId field if non-nil, zero value otherwise.

### GetTransitGatewayAttachmentIdOk

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) GetTransitGatewayAttachmentIdOk() (*string, bool)`

GetTransitGatewayAttachmentIdOk returns a tuple with the TransitGatewayAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayAttachmentId

`func (o *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) SetTransitGatewayAttachmentId(v string)`

SetTransitGatewayAttachmentId sets TransitGatewayAttachmentId field to given value.



### AsNetworkingAdminV1TransitGatewayAttachmentStatusCloudOneOf

`func (s *NetworkingAdminV1AwsTransitGatewayAttachmentStatus) AsNetworkingAdminV1TransitGatewayAttachmentStatusCloudOneOf() NetworkingAdminV1TransitGatewayAttachmentStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AwsTransitGatewayAttachmentStatus in NetworkingAdminV1TransitGatewayAttachmentStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


