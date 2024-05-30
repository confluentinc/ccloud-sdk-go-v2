# PimV1AwsIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfluentAwsRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud assumes when it accesses resources in your AWS account.  | [optional] [readonly] 
**ConfluentAwsExternalId** | Pointer to **string** | Unique external ID that Confluent Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.  | [optional] [readonly] 
**AwsRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that Confluent Cloud assumes when it accesses resources in your AWS account.  | [optional] 
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

### GetConfluentAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) GetConfluentAwsRoleArn() string`

GetConfluentAwsRoleArn returns the ConfluentAwsRoleArn field if non-nil, zero value otherwise.

### GetConfluentAwsRoleArnOk

`func (o *PimV1AwsIntegrationConfig) GetConfluentAwsRoleArnOk() (*string, bool)`

GetConfluentAwsRoleArnOk returns a tuple with the ConfluentAwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfluentAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) SetConfluentAwsRoleArn(v string)`

SetConfluentAwsRoleArn sets ConfluentAwsRoleArn field to given value.

### HasConfluentAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) HasConfluentAwsRoleArn() bool`

HasConfluentAwsRoleArn returns a boolean if a field has been set.

### GetConfluentAwsExternalId

`func (o *PimV1AwsIntegrationConfig) GetConfluentAwsExternalId() string`

GetConfluentAwsExternalId returns the ConfluentAwsExternalId field if non-nil, zero value otherwise.

### GetConfluentAwsExternalIdOk

`func (o *PimV1AwsIntegrationConfig) GetConfluentAwsExternalIdOk() (*string, bool)`

GetConfluentAwsExternalIdOk returns a tuple with the ConfluentAwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfluentAwsExternalId

`func (o *PimV1AwsIntegrationConfig) SetConfluentAwsExternalId(v string)`

SetConfluentAwsExternalId sets ConfluentAwsExternalId field to given value.

### HasConfluentAwsExternalId

`func (o *PimV1AwsIntegrationConfig) HasConfluentAwsExternalId() bool`

HasConfluentAwsExternalId returns a boolean if a field has been set.

### GetAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *PimV1AwsIntegrationConfig) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *PimV1AwsIntegrationConfig) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

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


