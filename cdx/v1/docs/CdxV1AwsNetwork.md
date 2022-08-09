# CdxV1AwsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**PrivateLinkEndpointService** | Pointer to **string** | The AWS VPC endpoint service for the network (used for PrivateLink) if available. | [optional] [readonly] 

## Methods

### NewCdxV1AwsNetwork

`func NewCdxV1AwsNetwork(kind string, ) *CdxV1AwsNetwork`

NewCdxV1AwsNetwork instantiates a new CdxV1AwsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1AwsNetworkWithDefaults

`func NewCdxV1AwsNetworkWithDefaults() *CdxV1AwsNetwork`

NewCdxV1AwsNetworkWithDefaults instantiates a new CdxV1AwsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CdxV1AwsNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1AwsNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1AwsNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateLinkEndpointService

`func (o *CdxV1AwsNetwork) GetPrivateLinkEndpointService() string`

GetPrivateLinkEndpointService returns the PrivateLinkEndpointService field if non-nil, zero value otherwise.

### GetPrivateLinkEndpointServiceOk

`func (o *CdxV1AwsNetwork) GetPrivateLinkEndpointServiceOk() (*string, bool)`

GetPrivateLinkEndpointServiceOk returns a tuple with the PrivateLinkEndpointService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkEndpointService

`func (o *CdxV1AwsNetwork) SetPrivateLinkEndpointService(v string)`

SetPrivateLinkEndpointService sets PrivateLinkEndpointService field to given value.

### HasPrivateLinkEndpointService

`func (o *CdxV1AwsNetwork) HasPrivateLinkEndpointService() bool`

HasPrivateLinkEndpointService returns a boolean if a field has been set.


### AsCdxV1NetworkCloudOneOf

`func (s *CdxV1AwsNetwork) AsCdxV1NetworkCloudOneOf() CdxV1NetworkCloudOneOf`

Convenience method to wrap this instance of CdxV1AwsNetwork in CdxV1NetworkCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


