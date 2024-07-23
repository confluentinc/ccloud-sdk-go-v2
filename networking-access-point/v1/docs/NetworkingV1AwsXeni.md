# NetworkingV1AwsXeni

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AwsXeni kind. | 
**NetworkInterfaces** | Pointer to **[]string** | List of the IDs of the Elastic Network Interfaces. | [optional] 

## Methods

### NewNetworkingV1AwsXeni

`func NewNetworkingV1AwsXeni(kind string, ) *NetworkingV1AwsXeni`

NewNetworkingV1AwsXeni instantiates a new NetworkingV1AwsXeni object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsXeniWithDefaults

`func NewNetworkingV1AwsXeniWithDefaults() *NetworkingV1AwsXeni`

NewNetworkingV1AwsXeniWithDefaults instantiates a new NetworkingV1AwsXeni object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsXeni) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsXeni) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsXeni) SetKind(v string)`

SetKind sets Kind field to given value.


### GetNetworkInterfaces

`func (o *NetworkingV1AwsXeni) GetNetworkInterfaces() []string`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NetworkingV1AwsXeni) GetNetworkInterfacesOk() (*[]string, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NetworkingV1AwsXeni) SetNetworkInterfaces(v []string)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NetworkingV1AwsXeni) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1AwsXeni) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsXeni in NetworkingV1AccessPointSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


