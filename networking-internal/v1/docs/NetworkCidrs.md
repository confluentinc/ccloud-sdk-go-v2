# NetworkCidrs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcCidr** | Pointer to **string** | CIDR in string format | [optional] 
**ZonalCidrs** | Pointer to **[]string** | CIDR in string format | [optional] 
**ReservedCidr** | Pointer to **string** | CIDR in string format | [optional] 
**InternalVpcCidr** | Pointer to **string** | CIDR in string format | [optional] 

## Methods

### NewNetworkCidrs

`func NewNetworkCidrs() *NetworkCidrs`

NewNetworkCidrs instantiates a new NetworkCidrs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkCidrsWithDefaults

`func NewNetworkCidrsWithDefaults() *NetworkCidrs`

NewNetworkCidrsWithDefaults instantiates a new NetworkCidrs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcCidr

`func (o *NetworkCidrs) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *NetworkCidrs) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *NetworkCidrs) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *NetworkCidrs) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetZonalCidrs

`func (o *NetworkCidrs) GetZonalCidrs() []string`

GetZonalCidrs returns the ZonalCidrs field if non-nil, zero value otherwise.

### GetZonalCidrsOk

`func (o *NetworkCidrs) GetZonalCidrsOk() (*[]string, bool)`

GetZonalCidrsOk returns a tuple with the ZonalCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalCidrs

`func (o *NetworkCidrs) SetZonalCidrs(v []string)`

SetZonalCidrs sets ZonalCidrs field to given value.

### HasZonalCidrs

`func (o *NetworkCidrs) HasZonalCidrs() bool`

HasZonalCidrs returns a boolean if a field has been set.

### GetReservedCidr

`func (o *NetworkCidrs) GetReservedCidr() string`

GetReservedCidr returns the ReservedCidr field if non-nil, zero value otherwise.

### GetReservedCidrOk

`func (o *NetworkCidrs) GetReservedCidrOk() (*string, bool)`

GetReservedCidrOk returns a tuple with the ReservedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidr

`func (o *NetworkCidrs) SetReservedCidr(v string)`

SetReservedCidr sets ReservedCidr field to given value.

### HasReservedCidr

`func (o *NetworkCidrs) HasReservedCidr() bool`

HasReservedCidr returns a boolean if a field has been set.

### GetInternalVpcCidr

`func (o *NetworkCidrs) GetInternalVpcCidr() string`

GetInternalVpcCidr returns the InternalVpcCidr field if non-nil, zero value otherwise.

### GetInternalVpcCidrOk

`func (o *NetworkCidrs) GetInternalVpcCidrOk() (*string, bool)`

GetInternalVpcCidrOk returns a tuple with the InternalVpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVpcCidr

`func (o *NetworkCidrs) SetInternalVpcCidr(v string)`

SetInternalVpcCidr sets InternalVpcCidr field to given value.

### HasInternalVpcCidr

`func (o *NetworkCidrs) HasInternalVpcCidr() bool`

HasInternalVpcCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


