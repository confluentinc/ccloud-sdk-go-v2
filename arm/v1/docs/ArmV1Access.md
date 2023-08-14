# ArmV1Access

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crn** | **string** | The requesting resource crn | 
**Roles** | **[]string** | List of desire roles for the resource | 

## Methods

### NewArmV1Access

`func NewArmV1Access(crn string, roles []string, ) *ArmV1Access`

NewArmV1Access instantiates a new ArmV1Access object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmV1AccessWithDefaults

`func NewArmV1AccessWithDefaults() *ArmV1Access`

NewArmV1AccessWithDefaults instantiates a new ArmV1Access object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrn

`func (o *ArmV1Access) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *ArmV1Access) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *ArmV1Access) SetCrn(v string)`

SetCrn sets Crn field to given value.


### GetRoles

`func (o *ArmV1Access) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ArmV1Access) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ArmV1Access) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


