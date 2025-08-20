# PimV2GcpIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleServiceAccount** | Pointer to **string** | The ID of the Google Service Account that Confluent Cloud uses to impersonate customer Google Service Account when it accesses resources in your GCP project.  | [optional] [readonly] 
**CustomerGoogleServiceAccount** | Pointer to **string** | The ID of the Google Service Account that Confluent Cloud impersonates to access resources in your GCP Project.  | [optional] 
**Kind** | **string** | Cloud provider specific config to which access is provided through provider integration. | 

## Methods

### NewPimV2GcpIntegrationConfig

`func NewPimV2GcpIntegrationConfig(kind string, ) *PimV2GcpIntegrationConfig`

NewPimV2GcpIntegrationConfig instantiates a new PimV2GcpIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2GcpIntegrationConfigWithDefaults

`func NewPimV2GcpIntegrationConfigWithDefaults() *PimV2GcpIntegrationConfig`

NewPimV2GcpIntegrationConfigWithDefaults instantiates a new PimV2GcpIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) GetGoogleServiceAccount() string`

GetGoogleServiceAccount returns the GoogleServiceAccount field if non-nil, zero value otherwise.

### GetGoogleServiceAccountOk

`func (o *PimV2GcpIntegrationConfig) GetGoogleServiceAccountOk() (*string, bool)`

GetGoogleServiceAccountOk returns a tuple with the GoogleServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) SetGoogleServiceAccount(v string)`

SetGoogleServiceAccount sets GoogleServiceAccount field to given value.

### HasGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) HasGoogleServiceAccount() bool`

HasGoogleServiceAccount returns a boolean if a field has been set.

### GetCustomerGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) GetCustomerGoogleServiceAccount() string`

GetCustomerGoogleServiceAccount returns the CustomerGoogleServiceAccount field if non-nil, zero value otherwise.

### GetCustomerGoogleServiceAccountOk

`func (o *PimV2GcpIntegrationConfig) GetCustomerGoogleServiceAccountOk() (*string, bool)`

GetCustomerGoogleServiceAccountOk returns a tuple with the CustomerGoogleServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) SetCustomerGoogleServiceAccount(v string)`

SetCustomerGoogleServiceAccount sets CustomerGoogleServiceAccount field to given value.

### HasCustomerGoogleServiceAccount

`func (o *PimV2GcpIntegrationConfig) HasCustomerGoogleServiceAccount() bool`

HasCustomerGoogleServiceAccount returns a boolean if a field has been set.

### GetKind

`func (o *PimV2GcpIntegrationConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2GcpIntegrationConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2GcpIntegrationConfig) SetKind(v string)`

SetKind sets Kind field to given value.



### AsPimV2IntegrationConfigOneOf

`func (s *PimV2GcpIntegrationConfig) AsPimV2IntegrationConfigOneOf() PimV2IntegrationConfigOneOf`

Convenience method to wrap this instance of PimV2GcpIntegrationConfig in PimV2IntegrationConfigOneOf

### AsPimV2IntegrationUpdateConfigOneOf

`func (s *PimV2GcpIntegrationConfig) AsPimV2IntegrationUpdateConfigOneOf() PimV2IntegrationUpdateConfigOneOf`

Convenience method to wrap this instance of PimV2GcpIntegrationConfig in PimV2IntegrationUpdateConfigOneOf

### AsPimV2IntegrationValidateRequestConfigOneOf

`func (s *PimV2GcpIntegrationConfig) AsPimV2IntegrationValidateRequestConfigOneOf() PimV2IntegrationValidateRequestConfigOneOf`

Convenience method to wrap this instance of PimV2GcpIntegrationConfig in PimV2IntegrationValidateRequestConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


