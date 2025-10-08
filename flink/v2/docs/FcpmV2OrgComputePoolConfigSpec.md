# FcpmV2OrgComputePoolConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPoolEnabled** | Pointer to **bool** | Whether default compute pools are enabled for the organization. When enabled, environments can have default compute pools created automatically.  | [optional] [default to true]
**DefaultPoolMaxCfu** | Pointer to **int32** | Maximum number of Confluent Flink Units (CFUs) that default compute pools in this organization should auto-scale to.  | [optional] 

## Methods

### NewFcpmV2OrgComputePoolConfigSpec

`func NewFcpmV2OrgComputePoolConfigSpec() *FcpmV2OrgComputePoolConfigSpec`

NewFcpmV2OrgComputePoolConfigSpec instantiates a new FcpmV2OrgComputePoolConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2OrgComputePoolConfigSpecWithDefaults

`func NewFcpmV2OrgComputePoolConfigSpecWithDefaults() *FcpmV2OrgComputePoolConfigSpec`

NewFcpmV2OrgComputePoolConfigSpecWithDefaults instantiates a new FcpmV2OrgComputePoolConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPoolEnabled

`func (o *FcpmV2OrgComputePoolConfigSpec) GetDefaultPoolEnabled() bool`

GetDefaultPoolEnabled returns the DefaultPoolEnabled field if non-nil, zero value otherwise.

### GetDefaultPoolEnabledOk

`func (o *FcpmV2OrgComputePoolConfigSpec) GetDefaultPoolEnabledOk() (*bool, bool)`

GetDefaultPoolEnabledOk returns a tuple with the DefaultPoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolEnabled

`func (o *FcpmV2OrgComputePoolConfigSpec) SetDefaultPoolEnabled(v bool)`

SetDefaultPoolEnabled sets DefaultPoolEnabled field to given value.

### HasDefaultPoolEnabled

`func (o *FcpmV2OrgComputePoolConfigSpec) HasDefaultPoolEnabled() bool`

HasDefaultPoolEnabled returns a boolean if a field has been set.

### GetDefaultPoolMaxCfu

`func (o *FcpmV2OrgComputePoolConfigSpec) GetDefaultPoolMaxCfu() int32`

GetDefaultPoolMaxCfu returns the DefaultPoolMaxCfu field if non-nil, zero value otherwise.

### GetDefaultPoolMaxCfuOk

`func (o *FcpmV2OrgComputePoolConfigSpec) GetDefaultPoolMaxCfuOk() (*int32, bool)`

GetDefaultPoolMaxCfuOk returns a tuple with the DefaultPoolMaxCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolMaxCfu

`func (o *FcpmV2OrgComputePoolConfigSpec) SetDefaultPoolMaxCfu(v int32)`

SetDefaultPoolMaxCfu sets DefaultPoolMaxCfu field to given value.

### HasDefaultPoolMaxCfu

`func (o *FcpmV2OrgComputePoolConfigSpec) HasDefaultPoolMaxCfu() bool`

HasDefaultPoolMaxCfu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


