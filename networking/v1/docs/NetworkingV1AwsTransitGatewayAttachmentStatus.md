# NetworkingV1AwsTransitGatewayAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | AWS Transit Gateway Attachment Status kind type. | [optional] 
**TransitGatewayAttachmentId** | **string** | The ID of the AWS Transit Gateway VPC Attachment that attaches Confluent VPC to Transit Gateway. | [readonly] 

## Methods

### NewNetworkingV1AwsTransitGatewayAttachmentStatus

`func NewNetworkingV1AwsTransitGatewayAttachmentStatus(transitGatewayAttachmentId string, ) *NetworkingV1AwsTransitGatewayAttachmentStatus`

NewNetworkingV1AwsTransitGatewayAttachmentStatus instantiates a new NetworkingV1AwsTransitGatewayAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsTransitGatewayAttachmentStatusWithDefaults

`func NewNetworkingV1AwsTransitGatewayAttachmentStatusWithDefaults() *NetworkingV1AwsTransitGatewayAttachmentStatus`

NewNetworkingV1AwsTransitGatewayAttachmentStatusWithDefaults instantiates a new NetworkingV1AwsTransitGatewayAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTransitGatewayAttachmentId

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) GetTransitGatewayAttachmentId() string`

GetTransitGatewayAttachmentId returns the TransitGatewayAttachmentId field if non-nil, zero value otherwise.

### GetTransitGatewayAttachmentIdOk

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) GetTransitGatewayAttachmentIdOk() (*string, bool)`

GetTransitGatewayAttachmentIdOk returns a tuple with the TransitGatewayAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayAttachmentId

`func (o *NetworkingV1AwsTransitGatewayAttachmentStatus) SetTransitGatewayAttachmentId(v string)`

SetTransitGatewayAttachmentId sets TransitGatewayAttachmentId field to given value.



### AsNetworkingV1TransitGatewayAttachmentStatusCloudOneOf

`func (s *NetworkingV1AwsTransitGatewayAttachmentStatus) AsNetworkingV1TransitGatewayAttachmentStatusCloudOneOf() NetworkingV1TransitGatewayAttachmentStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsTransitGatewayAttachmentStatus in NetworkingV1TransitGatewayAttachmentStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


