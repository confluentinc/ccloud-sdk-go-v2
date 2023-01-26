# NetworkingV1ZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **string** | Cloud provider zone id | [optional] 
**Cidr** | Pointer to **string** | The IPv4 [CIDR block](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) to used for this network. Must be a &#x60;/27&#x60;. Required for VPC peering and AWS TransitGateway.  | [optional] 

## Methods

### NewNetworkingV1ZoneInfo

`func NewNetworkingV1ZoneInfo() *NetworkingV1ZoneInfo`

NewNetworkingV1ZoneInfo instantiates a new NetworkingV1ZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1ZoneInfoWithDefaults

`func NewNetworkingV1ZoneInfoWithDefaults() *NetworkingV1ZoneInfo`

NewNetworkingV1ZoneInfoWithDefaults instantiates a new NetworkingV1ZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *NetworkingV1ZoneInfo) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *NetworkingV1ZoneInfo) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *NetworkingV1ZoneInfo) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *NetworkingV1ZoneInfo) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetCidr

`func (o *NetworkingV1ZoneInfo) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkingV1ZoneInfo) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkingV1ZoneInfo) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkingV1ZoneInfo) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


