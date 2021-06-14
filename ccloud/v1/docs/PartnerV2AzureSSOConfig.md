# PartnerV2AzureSSOConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | 
**TenantId** | Pointer to **string** | The Azure AD tenant ID | 

## Methods

### NewPartnerV2AzureSSOConfig

`func NewPartnerV2AzureSSOConfig(kind string, tenantId string, ) *PartnerV2AzureSSOConfig`

NewPartnerV2AzureSSOConfig instantiates a new PartnerV2AzureSSOConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerV2AzureSSOConfigWithDefaults

`func NewPartnerV2AzureSSOConfigWithDefaults() *PartnerV2AzureSSOConfig`

NewPartnerV2AzureSSOConfigWithDefaults instantiates a new PartnerV2AzureSSOConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PartnerV2AzureSSOConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerV2AzureSSOConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerV2AzureSSOConfig) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTenantId

`func (o *PartnerV2AzureSSOConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PartnerV2AzureSSOConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PartnerV2AzureSSOConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



### AsPartnerV2OrganizationSsoConfigOneOf

`func (s *PartnerV2AzureSSOConfig) AsPartnerV2OrganizationSsoConfigOneOf() PartnerV2OrganizationSsoConfigOneOf`

Convenience method to wrap this instance of PartnerV2AzureSSOConfig in PartnerV2OrganizationSsoConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


