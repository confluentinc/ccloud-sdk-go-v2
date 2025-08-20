# PimV2AzureIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfluentMultiTenantAppId** | Pointer to **string** | The ID of the Confluent Multi-Tenant App that Confluent Cloud uses to impersonate customer Azure App when it accesses resources in your Azure subscription.  | [optional] [readonly] 
**CustomerAzureTenantId** | Pointer to **string** | The ID of the customer&#39;s Azure Active Directory (Azure AD) tenant  | [optional] 
**Kind** | **string** | Cloud provider specific config to which access is provided through provider integration. | 

## Methods

### NewPimV2AzureIntegrationConfig

`func NewPimV2AzureIntegrationConfig(kind string, ) *PimV2AzureIntegrationConfig`

NewPimV2AzureIntegrationConfig instantiates a new PimV2AzureIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2AzureIntegrationConfigWithDefaults

`func NewPimV2AzureIntegrationConfigWithDefaults() *PimV2AzureIntegrationConfig`

NewPimV2AzureIntegrationConfigWithDefaults instantiates a new PimV2AzureIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfluentMultiTenantAppId

`func (o *PimV2AzureIntegrationConfig) GetConfluentMultiTenantAppId() string`

GetConfluentMultiTenantAppId returns the ConfluentMultiTenantAppId field if non-nil, zero value otherwise.

### GetConfluentMultiTenantAppIdOk

`func (o *PimV2AzureIntegrationConfig) GetConfluentMultiTenantAppIdOk() (*string, bool)`

GetConfluentMultiTenantAppIdOk returns a tuple with the ConfluentMultiTenantAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfluentMultiTenantAppId

`func (o *PimV2AzureIntegrationConfig) SetConfluentMultiTenantAppId(v string)`

SetConfluentMultiTenantAppId sets ConfluentMultiTenantAppId field to given value.

### HasConfluentMultiTenantAppId

`func (o *PimV2AzureIntegrationConfig) HasConfluentMultiTenantAppId() bool`

HasConfluentMultiTenantAppId returns a boolean if a field has been set.

### GetCustomerAzureTenantId

`func (o *PimV2AzureIntegrationConfig) GetCustomerAzureTenantId() string`

GetCustomerAzureTenantId returns the CustomerAzureTenantId field if non-nil, zero value otherwise.

### GetCustomerAzureTenantIdOk

`func (o *PimV2AzureIntegrationConfig) GetCustomerAzureTenantIdOk() (*string, bool)`

GetCustomerAzureTenantIdOk returns a tuple with the CustomerAzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAzureTenantId

`func (o *PimV2AzureIntegrationConfig) SetCustomerAzureTenantId(v string)`

SetCustomerAzureTenantId sets CustomerAzureTenantId field to given value.

### HasCustomerAzureTenantId

`func (o *PimV2AzureIntegrationConfig) HasCustomerAzureTenantId() bool`

HasCustomerAzureTenantId returns a boolean if a field has been set.

### GetKind

`func (o *PimV2AzureIntegrationConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2AzureIntegrationConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2AzureIntegrationConfig) SetKind(v string)`

SetKind sets Kind field to given value.



### AsPimV2IntegrationConfigOneOf

`func (s *PimV2AzureIntegrationConfig) AsPimV2IntegrationConfigOneOf() PimV2IntegrationConfigOneOf`

Convenience method to wrap this instance of PimV2AzureIntegrationConfig in PimV2IntegrationConfigOneOf

### AsPimV2IntegrationUpdateConfigOneOf

`func (s *PimV2AzureIntegrationConfig) AsPimV2IntegrationUpdateConfigOneOf() PimV2IntegrationUpdateConfigOneOf`

Convenience method to wrap this instance of PimV2AzureIntegrationConfig in PimV2IntegrationUpdateConfigOneOf

### AsPimV2IntegrationValidateRequestConfigOneOf

`func (s *PimV2AzureIntegrationConfig) AsPimV2IntegrationValidateRequestConfigOneOf() PimV2IntegrationValidateRequestConfigOneOf`

Convenience method to wrap this instance of PimV2AzureIntegrationConfig in PimV2IntegrationValidateRequestConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


