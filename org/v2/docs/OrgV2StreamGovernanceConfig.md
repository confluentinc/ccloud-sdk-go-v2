# OrgV2StreamGovernanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | **string** | Stream Governance Package. Supported values are ESSENTIALS and ADVANCED. Package comparison can be found [here](https://docs.confluent.io/cloud/current/stream-governance/packages.html#features-by-package-type).  | 

## Methods

### NewOrgV2StreamGovernanceConfig

`func NewOrgV2StreamGovernanceConfig(package_ string, ) *OrgV2StreamGovernanceConfig`

NewOrgV2StreamGovernanceConfig instantiates a new OrgV2StreamGovernanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2StreamGovernanceConfigWithDefaults

`func NewOrgV2StreamGovernanceConfigWithDefaults() *OrgV2StreamGovernanceConfig`

NewOrgV2StreamGovernanceConfigWithDefaults instantiates a new OrgV2StreamGovernanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *OrgV2StreamGovernanceConfig) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *OrgV2StreamGovernanceConfig) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *OrgV2StreamGovernanceConfig) SetPackage(v string)`

SetPackage sets Package field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


