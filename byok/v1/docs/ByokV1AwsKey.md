# ByokV1AwsKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyArn** | **string** | The Amazon Resource Name (ARN) of an AWS KMS key.  | 
**Roles** | Pointer to **[]string** | The Amazon Resource Names (ARNs) of IAM Roles created for this key-environment combination.  | [optional] [readonly] 
**Kind** | **string** | BYOK kind type.  | 

## Methods

### NewByokV1AwsKey

`func NewByokV1AwsKey(keyArn string, kind string, ) *ByokV1AwsKey`

NewByokV1AwsKey instantiates a new ByokV1AwsKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1AwsKeyWithDefaults

`func NewByokV1AwsKeyWithDefaults() *ByokV1AwsKey`

NewByokV1AwsKeyWithDefaults instantiates a new ByokV1AwsKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyArn

`func (o *ByokV1AwsKey) GetKeyArn() string`

GetKeyArn returns the KeyArn field if non-nil, zero value otherwise.

### GetKeyArnOk

`func (o *ByokV1AwsKey) GetKeyArnOk() (*string, bool)`

GetKeyArnOk returns a tuple with the KeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArn

`func (o *ByokV1AwsKey) SetKeyArn(v string)`

SetKeyArn sets KeyArn field to given value.


### GetRoles

`func (o *ByokV1AwsKey) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ByokV1AwsKey) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ByokV1AwsKey) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ByokV1AwsKey) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1AwsKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1AwsKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1AwsKey) SetKind(v string)`

SetKind sets Kind field to given value.



### AsByokV1KeyKeyOneOf

`func (s *ByokV1AwsKey) AsByokV1KeyKeyOneOf() ByokV1KeyKeyOneOf`

Convenience method to wrap this instance of ByokV1AwsKey in ByokV1KeyKeyOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


