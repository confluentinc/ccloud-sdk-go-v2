# NetworkingV1IpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**IpPrefix** | Pointer to **string** | The IP Address range. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider in which the address exists. | [optional] 
**Region** | Pointer to **string** | The region/location where the IP Address is in use. | [optional] 
**Services** | Pointer to [**NetworkingV1Services**](networking.v1.Services.md) |  | [optional] 
**AddressType** | Pointer to **string** | Whether the address is used for egress or ingress. | [optional] 

## Methods

### NewNetworkingV1IpAddress

`func NewNetworkingV1IpAddress() *NetworkingV1IpAddress`

NewNetworkingV1IpAddress instantiates a new NetworkingV1IpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1IpAddressWithDefaults

`func NewNetworkingV1IpAddressWithDefaults() *NetworkingV1IpAddress`

NewNetworkingV1IpAddressWithDefaults instantiates a new NetworkingV1IpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1IpAddress) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1IpAddress) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1IpAddress) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1IpAddress) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1IpAddress) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1IpAddress) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1IpAddress) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1IpAddress) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetIpPrefix

`func (o *NetworkingV1IpAddress) GetIpPrefix() string`

GetIpPrefix returns the IpPrefix field if non-nil, zero value otherwise.

### GetIpPrefixOk

`func (o *NetworkingV1IpAddress) GetIpPrefixOk() (*string, bool)`

GetIpPrefixOk returns a tuple with the IpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefix

`func (o *NetworkingV1IpAddress) SetIpPrefix(v string)`

SetIpPrefix sets IpPrefix field to given value.

### HasIpPrefix

`func (o *NetworkingV1IpAddress) HasIpPrefix() bool`

HasIpPrefix returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1IpAddress) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1IpAddress) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1IpAddress) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1IpAddress) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkingV1IpAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1IpAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1IpAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkingV1IpAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *NetworkingV1IpAddress) GetServices() NetworkingV1Services`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NetworkingV1IpAddress) GetServicesOk() (*NetworkingV1Services, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NetworkingV1IpAddress) SetServices(v NetworkingV1Services)`

SetServices sets Services field to given value.

### HasServices

`func (o *NetworkingV1IpAddress) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetAddressType

`func (o *NetworkingV1IpAddress) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *NetworkingV1IpAddress) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *NetworkingV1IpAddress) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *NetworkingV1IpAddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


