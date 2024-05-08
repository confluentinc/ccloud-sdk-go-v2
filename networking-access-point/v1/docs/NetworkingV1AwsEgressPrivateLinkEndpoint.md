# NetworkingV1AwsEgressPrivateLinkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsEgressPrivateLinkEndpoint kind. | 
**VpcEndpointServiceName** | **string** | ID of the VPC Endpoint service used for PrivateLink. | 
**EnableHighAvailability** | Pointer to **bool** | Whether a resource should be provisioned with high availability. Endpoints deployed with high availability have network interfaces deployed in multiple AZs. | [optional] 

## Methods

### NewNetworkingV1AwsEgressPrivateLinkEndpoint

`func NewNetworkingV1AwsEgressPrivateLinkEndpoint(kind string, vpcEndpointServiceName string, ) *NetworkingV1AwsEgressPrivateLinkEndpoint`

NewNetworkingV1AwsEgressPrivateLinkEndpoint instantiates a new NetworkingV1AwsEgressPrivateLinkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsEgressPrivateLinkEndpointWithDefaults

`func NewNetworkingV1AwsEgressPrivateLinkEndpointWithDefaults() *NetworkingV1AwsEgressPrivateLinkEndpoint`

NewNetworkingV1AwsEgressPrivateLinkEndpointWithDefaults instantiates a new NetworkingV1AwsEgressPrivateLinkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointServiceName

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetVpcEndpointServiceName() string`

GetVpcEndpointServiceName returns the VpcEndpointServiceName field if non-nil, zero value otherwise.

### GetVpcEndpointServiceNameOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetVpcEndpointServiceNameOk() (*string, bool)`

GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointServiceName

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) SetVpcEndpointServiceName(v string)`

SetVpcEndpointServiceName sets VpcEndpointServiceName field to given value.


### GetEnableHighAvailability

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetEnableHighAvailability() bool`

GetEnableHighAvailability returns the EnableHighAvailability field if non-nil, zero value otherwise.

### GetEnableHighAvailabilityOk

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) GetEnableHighAvailabilityOk() (*bool, bool)`

GetEnableHighAvailabilityOk returns a tuple with the EnableHighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHighAvailability

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) SetEnableHighAvailability(v bool)`

SetEnableHighAvailability sets EnableHighAvailability field to given value.

### HasEnableHighAvailability

`func (o *NetworkingV1AwsEgressPrivateLinkEndpoint) HasEnableHighAvailability() bool`

HasEnableHighAvailability returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AwsEgressPrivateLinkEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsEgressPrivateLinkEndpoint in NetworkingV1AccessPointSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


