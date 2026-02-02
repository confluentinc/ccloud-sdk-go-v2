# NetworkingV1AwsIngressPrivateLinkEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsIngressPrivateLinkEndpointStatus kind. | 
**VpcEndpointServiceName** | **string** | ID of the Confluent Cloud VPC Endpoint service used for PrivateLink. | [readonly] 
**VpcEndpointId** | **string** | ID of the VPC Endpoint used for connecting to the VPC Endpoint service. | [readonly] 
**DnsDomain** | Pointer to **string** | DNS domain name used to configure the Private Hosted Zone for the Access Point. | [optional] [readonly] 

## Methods

### NewNetworkingV1AwsIngressPrivateLinkEndpointStatus

`func NewNetworkingV1AwsIngressPrivateLinkEndpointStatus(kind string, vpcEndpointServiceName string, vpcEndpointId string, ) *NetworkingV1AwsIngressPrivateLinkEndpointStatus`

NewNetworkingV1AwsIngressPrivateLinkEndpointStatus instantiates a new NetworkingV1AwsIngressPrivateLinkEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsIngressPrivateLinkEndpointStatusWithDefaults

`func NewNetworkingV1AwsIngressPrivateLinkEndpointStatusWithDefaults() *NetworkingV1AwsIngressPrivateLinkEndpointStatus`

NewNetworkingV1AwsIngressPrivateLinkEndpointStatusWithDefaults instantiates a new NetworkingV1AwsIngressPrivateLinkEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVpcEndpointServiceName

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetVpcEndpointServiceName() string`

GetVpcEndpointServiceName returns the VpcEndpointServiceName field if non-nil, zero value otherwise.

### GetVpcEndpointServiceNameOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetVpcEndpointServiceNameOk() (*string, bool)`

GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointServiceName

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) SetVpcEndpointServiceName(v string)`

SetVpcEndpointServiceName sets VpcEndpointServiceName field to given value.


### GetVpcEndpointId

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetVpcEndpointId() string`

GetVpcEndpointId returns the VpcEndpointId field if non-nil, zero value otherwise.

### GetVpcEndpointIdOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetVpcEndpointIdOk() (*string, bool)`

GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointId

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) SetVpcEndpointId(v string)`

SetVpcEndpointId sets VpcEndpointId field to given value.


### GetDnsDomain

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *NetworkingV1AwsIngressPrivateLinkEndpointStatus) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.


### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1AwsIngressPrivateLinkEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressPrivateLinkEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


