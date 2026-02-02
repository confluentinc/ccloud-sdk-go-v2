# NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | AWS Private Network Interface Gateway Spec kind type. | 
**Region** | **string** | AWS region of the Private Network Interface Gateway. | 
**Zones** | **[]string** | AWS availability zone ids of the Private Network Interface Gateway. | 

## Methods

### NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpec

`func NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpec(kind string, region string, zones []string, ) *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec`

NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpec instantiates a new NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpecWithDefaults

`func NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpecWithDefaults() *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec`

NewNetworkingV1AwsPrivateNetworkInterfaceGatewaySpecWithDefaults instantiates a new NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegion

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetZones

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) SetZones(v []string)`

SetZones sets Zones field to given value.



### AsNetworkingV1GatewaySpecConfigOneOf

`func (s *NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec) AsNetworkingV1GatewaySpecConfigOneOf() NetworkingV1GatewaySpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPrivateNetworkInterfaceGatewaySpec in NetworkingV1GatewaySpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


