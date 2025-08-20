# PimV2AwsIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud uses to assume customer IAM role when it accesses resources in your AWS account.  | [optional] [readonly] 
**ExternalId** | Pointer to **string** | Unique external ID that Confluent Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.  | [optional] [readonly] 
**CustomerIamRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud assumes when it accesses resources in your AWS account.  | [optional] 
**Kind** | **string** | Cloud provider specific config to which access is provided through provider integration. | 

## Methods

### NewPimV2AwsIntegrationConfig

`func NewPimV2AwsIntegrationConfig(kind string, ) *PimV2AwsIntegrationConfig`

NewPimV2AwsIntegrationConfig instantiates a new PimV2AwsIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2AwsIntegrationConfigWithDefaults

`func NewPimV2AwsIntegrationConfigWithDefaults() *PimV2AwsIntegrationConfig`

NewPimV2AwsIntegrationConfigWithDefaults instantiates a new PimV2AwsIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIamRoleArn

`func (o *PimV2AwsIntegrationConfig) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *PimV2AwsIntegrationConfig) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *PimV2AwsIntegrationConfig) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *PimV2AwsIntegrationConfig) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### GetExternalId

`func (o *PimV2AwsIntegrationConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PimV2AwsIntegrationConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PimV2AwsIntegrationConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PimV2AwsIntegrationConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomerIamRoleArn

`func (o *PimV2AwsIntegrationConfig) GetCustomerIamRoleArn() string`

GetCustomerIamRoleArn returns the CustomerIamRoleArn field if non-nil, zero value otherwise.

### GetCustomerIamRoleArnOk

`func (o *PimV2AwsIntegrationConfig) GetCustomerIamRoleArnOk() (*string, bool)`

GetCustomerIamRoleArnOk returns a tuple with the CustomerIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIamRoleArn

`func (o *PimV2AwsIntegrationConfig) SetCustomerIamRoleArn(v string)`

SetCustomerIamRoleArn sets CustomerIamRoleArn field to given value.

### HasCustomerIamRoleArn

`func (o *PimV2AwsIntegrationConfig) HasCustomerIamRoleArn() bool`

HasCustomerIamRoleArn returns a boolean if a field has been set.

### GetKind

`func (o *PimV2AwsIntegrationConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2AwsIntegrationConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2AwsIntegrationConfig) SetKind(v string)`

SetKind sets Kind field to given value.



### AsPimV2IntegrationConfigOneOf

`func (s *PimV2AwsIntegrationConfig) AsPimV2IntegrationConfigOneOf() PimV2IntegrationConfigOneOf`

Convenience method to wrap this instance of PimV2AwsIntegrationConfig in PimV2IntegrationConfigOneOf

### AsPimV2IntegrationUpdateConfigOneOf

`func (s *PimV2AwsIntegrationConfig) AsPimV2IntegrationUpdateConfigOneOf() PimV2IntegrationUpdateConfigOneOf`

Convenience method to wrap this instance of PimV2AwsIntegrationConfig in PimV2IntegrationUpdateConfigOneOf

### AsPimV2IntegrationValidateRequestConfigOneOf

`func (s *PimV2AwsIntegrationConfig) AsPimV2IntegrationValidateRequestConfigOneOf() PimV2IntegrationValidateRequestConfigOneOf`

Convenience method to wrap this instance of PimV2AwsIntegrationConfig in PimV2IntegrationValidateRequestConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


