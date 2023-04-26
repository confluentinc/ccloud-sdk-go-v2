# CidrValidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcCidr** | Pointer to **string** | CIDR in string format | [optional] 
**ZonalCidrs** | Pointer to **[]string** | CIDR in string format | [optional] 
**ReservedCidr** | Pointer to **string** | CIDR in string format | [optional] 
**InternalVpcCidr** | Pointer to **string** | CIDR in string format | [optional] 
**Cloud** | Pointer to **string** | cloud provider | [optional] 

## Methods

### NewCidrValidate

`func NewCidrValidate() *CidrValidate`

NewCidrValidate instantiates a new CidrValidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCidrValidateWithDefaults

`func NewCidrValidateWithDefaults() *CidrValidate`

NewCidrValidateWithDefaults instantiates a new CidrValidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcCidr

`func (o *CidrValidate) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *CidrValidate) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *CidrValidate) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *CidrValidate) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetZonalCidrs

`func (o *CidrValidate) GetZonalCidrs() []string`

GetZonalCidrs returns the ZonalCidrs field if non-nil, zero value otherwise.

### GetZonalCidrsOk

`func (o *CidrValidate) GetZonalCidrsOk() (*[]string, bool)`

GetZonalCidrsOk returns a tuple with the ZonalCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalCidrs

`func (o *CidrValidate) SetZonalCidrs(v []string)`

SetZonalCidrs sets ZonalCidrs field to given value.

### HasZonalCidrs

`func (o *CidrValidate) HasZonalCidrs() bool`

HasZonalCidrs returns a boolean if a field has been set.

### GetReservedCidr

`func (o *CidrValidate) GetReservedCidr() string`

GetReservedCidr returns the ReservedCidr field if non-nil, zero value otherwise.

### GetReservedCidrOk

`func (o *CidrValidate) GetReservedCidrOk() (*string, bool)`

GetReservedCidrOk returns a tuple with the ReservedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidr

`func (o *CidrValidate) SetReservedCidr(v string)`

SetReservedCidr sets ReservedCidr field to given value.

### HasReservedCidr

`func (o *CidrValidate) HasReservedCidr() bool`

HasReservedCidr returns a boolean if a field has been set.

### GetInternalVpcCidr

`func (o *CidrValidate) GetInternalVpcCidr() string`

GetInternalVpcCidr returns the InternalVpcCidr field if non-nil, zero value otherwise.

### GetInternalVpcCidrOk

`func (o *CidrValidate) GetInternalVpcCidrOk() (*string, bool)`

GetInternalVpcCidrOk returns a tuple with the InternalVpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVpcCidr

`func (o *CidrValidate) SetInternalVpcCidr(v string)`

SetInternalVpcCidr sets InternalVpcCidr field to given value.

### HasInternalVpcCidr

`func (o *CidrValidate) HasInternalVpcCidr() bool`

HasInternalVpcCidr returns a boolean if a field has been set.

### GetCloud

`func (o *CidrValidate) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CidrValidate) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CidrValidate) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CidrValidate) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


