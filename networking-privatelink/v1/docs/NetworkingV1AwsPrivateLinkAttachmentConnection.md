# NetworkingV1AwsPrivateLinkAttachmentConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnection kind. | 
**VpcEndpointId** | **string** | Id of a VPC Endpoint that is connected to the VPC Endpoint service. | 

## Methods

### NewNetworkingV1AwsPrivateLinkAttachmentConnection

`func NewNetworkingV1AwsPrivateLinkAttachmentConnection(kind string, vpcEndpointId string, ) *NetworkingV1AwsPrivateLinkAttachmentConnection`

NewNetworkingV1AwsPrivateLinkAttachmentConnection instantiates a new NetworkingV1AwsPrivateLinkAttachmentConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateLinkAttachmentConnectionWithDefaults

`func NewNetworkingV1AwsPrivateLinkAttachmentConnectionWithDefaults() *NetworkingV1AwsPrivateLinkAttachmentConnection`

NewNetworkingV1AwsPrivateLinkAttachmentConnectionWithDefaults instantiates a new NetworkingV1AwsPrivateLinkAttachmentConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointId

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnection) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

`func (s *NetworkingV1AwsPrivateLinkAttachmentConnection) AsNetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateLinkAttachmentConnection in NetworkingV1PrivateLinkAttachmentConnectionSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


