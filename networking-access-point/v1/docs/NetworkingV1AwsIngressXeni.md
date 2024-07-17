# NetworkingV1AwsIngressXeni

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsIngressXeni kind. | 
**NetworkInterfaces** | Pointer to **[]string** | List of the IDs of the Elastic Network Interfaces. | [optional] 

## Methods

### NewNetworkingV1AwsIngressXeni

`func NewNetworkingV1AwsIngressXeni(kind string, ) *NetworkingV1AwsIngressXeni`

NewNetworkingV1AwsIngressXeni instantiates a new NetworkingV1AwsIngressXeni object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsIngressXeniWithDefaults

`func NewNetworkingV1AwsIngressXeniWithDefaults() *NetworkingV1AwsIngressXeni`

NewNetworkingV1AwsIngressXeniWithDefaults instantiates a new NetworkingV1AwsIngressXeni object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsIngressXeni) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsIngressXeni) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsIngressXeni) SetKind(v string)`

SetKind sets Kind field to given value.


### GetNetworkInterfaces

`func (o *NetworkingV1AwsIngressXeni) GetNetworkInterfaces() []string`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NetworkingV1AwsIngressXeni) GetNetworkInterfacesOk() (*[]string, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NetworkingV1AwsIngressXeni) SetNetworkInterfaces(v []string)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NetworkingV1AwsIngressXeni) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AwsIngressXeni) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsIngressXeni in NetworkingV1AccessPointSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


