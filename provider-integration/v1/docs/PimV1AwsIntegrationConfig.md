# PimV1AwsIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud uses to assume customer IAM role when it accesses resources in your AWS account.  | [optional] [readonly] 
**ExternalId** | Pointer to **string** | Unique external ID that Confluent Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.  | [optional] [readonly] 
**CustomerIamRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud assumes when it accesses resources in your AWS account.  | [optional] 
**Kind** | **string** | Cloud provider specific config to which access is provided through provider integration. | 

## Methods

### NewPimV1AwsIntegrationConfig

`func NewPimV1AwsIntegrationConfig(kind string, ) *PimV1AwsIntegrationConfig`

NewPimV1AwsIntegrationConfig instantiates a new PimV1AwsIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV1AwsIntegrationConfigWithDefaults

`func NewPimV1AwsIntegrationConfigWithDefaults() *PimV1AwsIntegrationConfig`

NewPimV1AwsIntegrationConfigWithDefaults instantiates a new PimV1AwsIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIamRoleArn

`func (o *PimV1AwsIntegrationConfig) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *PimV1AwsIntegrationConfig) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *PimV1AwsIntegrationConfig) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *PimV1AwsIntegrationConfig) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### GetExternalId

`func (o *PimV1AwsIntegrationConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PimV1AwsIntegrationConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PimV1AwsIntegrationConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PimV1AwsIntegrationConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomerIamRoleArn

`func (o *PimV1AwsIntegrationConfig) GetCustomerIamRoleArn() string`

GetCustomerIamRoleArn returns the CustomerIamRoleArn field if non-nil, zero value otherwise.

### GetCustomerIamRoleArnOk

`func (o *PimV1AwsIntegrationConfig) GetCustomerIamRoleArnOk() (*string, bool)`

GetCustomerIamRoleArnOk returns a tuple with the CustomerIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIamRoleArn

`func (o *PimV1AwsIntegrationConfig) SetCustomerIamRoleArn(v string)`

SetCustomerIamRoleArn sets CustomerIamRoleArn field to given value.

### HasCustomerIamRoleArn

`func (o *PimV1AwsIntegrationConfig) HasCustomerIamRoleArn() bool`

HasCustomerIamRoleArn returns a boolean if a field has been set.

### GetKind

`func (o *PimV1AwsIntegrationConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV1AwsIntegrationConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV1AwsIntegrationConfig) SetKind(v string)`

SetKind sets Kind field to given value.



### AsPimV1IntegrationConfigOneOf

`func (s *PimV1AwsIntegrationConfig) AsPimV1IntegrationConfigOneOf() PimV1IntegrationConfigOneOf`

Convenience method to wrap this instance of PimV1AwsIntegrationConfig in PimV1IntegrationConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


