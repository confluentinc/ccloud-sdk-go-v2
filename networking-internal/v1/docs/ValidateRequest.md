# ValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcCidr** | Pointer to **string** | CIDR in string format | [optional] 
**ZonalCidrs** | Pointer to **[]string** | CIDR in string format | [optional] 
**ReservedCidr** | Pointer to **string** | CIDR in string format | [optional] 
**InternalVpcCidr** | Pointer to **string** | CIDR in string format | [optional] 
**Cloud** | Pointer to **string** | cloud provider | [optional] 

## Methods

### NewValidateRequest

`func NewValidateRequest() *ValidateRequest`

NewValidateRequest instantiates a new ValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateRequestWithDefaults

`func NewValidateRequestWithDefaults() *ValidateRequest`

NewValidateRequestWithDefaults instantiates a new ValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcCidr

`func (o *ValidateRequest) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *ValidateRequest) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *ValidateRequest) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *ValidateRequest) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetZonalCidrs

`func (o *ValidateRequest) GetZonalCidrs() []string`

GetZonalCidrs returns the ZonalCidrs field if non-nil, zero value otherwise.

### GetZonalCidrsOk

`func (o *ValidateRequest) GetZonalCidrsOk() (*[]string, bool)`

GetZonalCidrsOk returns a tuple with the ZonalCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalCidrs

`func (o *ValidateRequest) SetZonalCidrs(v []string)`

SetZonalCidrs sets ZonalCidrs field to given value.

### HasZonalCidrs

`func (o *ValidateRequest) HasZonalCidrs() bool`

HasZonalCidrs returns a boolean if a field has been set.

### GetReservedCidr

`func (o *ValidateRequest) GetReservedCidr() string`

GetReservedCidr returns the ReservedCidr field if non-nil, zero value otherwise.

### GetReservedCidrOk

`func (o *ValidateRequest) GetReservedCidrOk() (*string, bool)`

GetReservedCidrOk returns a tuple with the ReservedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidr

`func (o *ValidateRequest) SetReservedCidr(v string)`

SetReservedCidr sets ReservedCidr field to given value.

### HasReservedCidr

`func (o *ValidateRequest) HasReservedCidr() bool`

HasReservedCidr returns a boolean if a field has been set.

### GetInternalVpcCidr

`func (o *ValidateRequest) GetInternalVpcCidr() string`

GetInternalVpcCidr returns the InternalVpcCidr field if non-nil, zero value otherwise.

### GetInternalVpcCidrOk

`func (o *ValidateRequest) GetInternalVpcCidrOk() (*string, bool)`

GetInternalVpcCidrOk returns a tuple with the InternalVpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVpcCidr

`func (o *ValidateRequest) SetInternalVpcCidr(v string)`

SetInternalVpcCidr sets InternalVpcCidr field to given value.

### HasInternalVpcCidr

`func (o *ValidateRequest) HasInternalVpcCidr() bool`

HasInternalVpcCidr returns a boolean if a field has been set.

### GetCloud

`func (o *ValidateRequest) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ValidateRequest) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ValidateRequest) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ValidateRequest) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


