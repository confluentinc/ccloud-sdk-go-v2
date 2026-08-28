# IamV2ApiKeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | The API key secret. Only provided in &#x60;create&#x60; responses, not in &#x60;get&#x60; or &#x60;list&#x60;. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human readable name for the API key | [optional] 
**Description** | Pointer to **string** | A human readable description for the API key | [optional] 
**ExpiresAt** | Pointer to **string** | The UTC date on which this API key expires, as an ISO 8601 date (YYYY-MM-DD). The key remains valid through the end of this date and expires at 23:59:59Z. Once expired, the key is rejected for authentication; access to some services may continue briefly while the expiration propagates.  | [optional] 
**CreatedBy** | Pointer to **string** | The principal ID that created this API key. | [optional] [readonly] 
**Owner** | Pointer to [**TypedGlobalObjectReference**](TypedGlobalObjectReference.md) | The owner to which this belongs. The owner can be one of iam.v2.User, iam.v2.ServiceAccount. | [optional] 
**Resource** | Pointer to [**NullableTypedEnvScopedObjectReference**](TypedEnvScopedObjectReference.md) | The resource associated with this object. The resource can be one of Kafka Cluster ID (example: lkc-12345), Schema Registry Cluster ID (example: lsrc-12345), ksqlDB Cluster ID (example: lksqlc-12345), or Flink (Environment + Region pair, example: env-abc123.aws.us-east-2). May be null or omitted if not associated with a resource. For creating Cloud API key, resource id should be &#x60;CLOUD&#x60;, for creating Tableflow API key, resource id should be &#x60;TABLEFLOW&#x60;, for creating Global API key, resource id should be &#x60;GLOBAL&#x60;. The resource id is case-insensitive. [Learn more in Authentication](https://docs.confluent.io/cloud/current/api.html#section/Authentication).  Note - Flink is in the [Preview lifecycle stage](https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)  | [optional] 

## Methods

### NewIamV2ApiKeySpec

`func NewIamV2ApiKeySpec() *IamV2ApiKeySpec`

NewIamV2ApiKeySpec instantiates a new IamV2ApiKeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ApiKeySpecWithDefaults

`func NewIamV2ApiKeySpecWithDefaults() *IamV2ApiKeySpec`

NewIamV2ApiKeySpecWithDefaults instantiates a new IamV2ApiKeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *IamV2ApiKeySpec) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IamV2ApiKeySpec) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IamV2ApiKeySpec) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IamV2ApiKeySpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2ApiKeySpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2ApiKeySpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2ApiKeySpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2ApiKeySpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2ApiKeySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2ApiKeySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2ApiKeySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2ApiKeySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *IamV2ApiKeySpec) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IamV2ApiKeySpec) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IamV2ApiKeySpec) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IamV2ApiKeySpec) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IamV2ApiKeySpec) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IamV2ApiKeySpec) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IamV2ApiKeySpec) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IamV2ApiKeySpec) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *IamV2ApiKeySpec) GetOwner() TypedGlobalObjectReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamV2ApiKeySpec) GetOwnerOk() (*TypedGlobalObjectReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamV2ApiKeySpec) SetOwner(v TypedGlobalObjectReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamV2ApiKeySpec) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResource

`func (o *IamV2ApiKeySpec) GetResource() TypedEnvScopedObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamV2ApiKeySpec) GetResourceOk() (*TypedEnvScopedObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamV2ApiKeySpec) SetResource(v TypedEnvScopedObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamV2ApiKeySpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *IamV2ApiKeySpec) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *IamV2ApiKeySpec) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


