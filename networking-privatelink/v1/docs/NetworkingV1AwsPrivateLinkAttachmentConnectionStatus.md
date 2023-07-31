# NetworkingV1AwsPrivateLinkAttachmentConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAttachmentConnectionStatus kind. | 
**VpcEndpointServiceName** | **string** | Id of the VPC Endpoint service used for PrivateLink. | [readonly] 
**VpcEndpointId** | **string** | Id of the VPC Endpoint (if any) that is connected to the VPC Endpoint service. | [readonly] 

## Methods

### NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatus

`func NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatus(kind string, vpcEndpointServiceName string, vpcEndpointId string, ) *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus`

NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatus instantiates a new NetworkingV1AwsPrivateLinkAttachmentConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatusWithDefaults

`func NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatusWithDefaults() *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus`

NewNetworkingV1AwsPrivateLinkAttachmentConnectionStatusWithDefaults instantiates a new NetworkingV1AwsPrivateLinkAttachmentConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointServiceName

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetVpcEndpointServiceName() string`

GetVpcEndpointServiceName returns the VpcEndpointServiceName field if non-nil, zero value otherwise.

### GetVpcEndpointServiceNameOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetVpcEndpointServiceNameOk() (*string, bool)`

GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointServiceName

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) SetVpcEndpointServiceName(v string)`

SetVpcEndpointServiceName sets VpcEndpointServiceName field to given value.


### GetVpcEndpointId

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.



### AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

`func (s *NetworkingV1AwsPrivateLinkAttachmentConnectionStatus) AsNetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf() NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateLinkAttachmentConnectionStatus in NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


