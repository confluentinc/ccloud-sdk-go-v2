# NetworkingV1AwsPrivateLinkAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentStatus kind. | [readonly] 
**VpcEndpointService** | [**NetworkingV1AwsVpcEndpointService**](networking.v1.AwsVpcEndpointService.md) | AWS VPC Endpoint Service that can be used to establish connections for all zones.  | [readonly] 

## Methods

### NewNetworkingV1AwsPrivateLinkAttachmentStatus

`func NewNetworkingV1AwsPrivateLinkAttachmentStatus(kind string, vpcEndpointService NetworkingV1AwsVpcEndpointService, ) *NetworkingV1AwsPrivateLinkAttachmentStatus`

NewNetworkingV1AwsPrivateLinkAttachmentStatus instantiates a new NetworkingV1AwsPrivateLinkAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateLinkAttachmentStatusWithDefaults

`func NewNetworkingV1AwsPrivateLinkAttachmentStatusWithDefaults() *NetworkingV1AwsPrivateLinkAttachmentStatus`

NewNetworkingV1AwsPrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1AwsPrivateLinkAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointService

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) GetVpcEndpointService() NetworkingV1AwsVpcEndpointService`

GetVpcEndpointService returns the VpcEndpointService field if non-nil, zero value otherwise.

### GetVpcEndpointServiceOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) GetVpcEndpointServiceOk() (*NetworkingV1AwsVpcEndpointService, bool)`

GetVpcEndpointServiceOk returns a tuple with the VpcEndpointService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointService

`func (o *NetworkingV1AwsPrivateLinkAttachmentStatus) SetVpcEndpointService(v NetworkingV1AwsVpcEndpointService)`

SetVpcEndpointService sets VpcEndpointService field to given value.



### AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf

`func (s *NetworkingV1AwsPrivateLinkAttachmentStatus) AsNetworkingV1PrivateLinkAttachmentStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateLinkAttachmentStatus in NetworkingV1PrivateLinkAttachmentStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


