# CdxV1RedeemTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ApiKey** | Pointer to **string** | The api key | [optional] [readonly] 
**Secret** | Pointer to **string** | The api key secret | [optional] [readonly] 
**KafkaBootstrapUrl** | Pointer to **string** | The kafka cluster bootstrap url | [optional] [readonly] 
**SchemaRegistryApiKey** | Pointer to **string** | The api key for schema registry | [optional] [readonly] 
**SchemaRegistrySecret** | Pointer to **string** | The api key secret for schema registry | [optional] [readonly] 
**SchemaRegistryUrl** | Pointer to **string** | The schema registry endpoint url | [optional] [readonly] 
**Resources** | Pointer to [**[]CdxV1RedeemTokenResponseResourcesOneOf**](CdxV1RedeemTokenResponseResourcesOneOf.md) | List of shared resources | [optional] 

## Methods

### NewCdxV1RedeemTokenResponse

`func NewCdxV1RedeemTokenResponse() *CdxV1RedeemTokenResponse`

NewCdxV1RedeemTokenResponse instantiates a new CdxV1RedeemTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1RedeemTokenResponseWithDefaults

`func NewCdxV1RedeemTokenResponseWithDefaults() *CdxV1RedeemTokenResponse`

NewCdxV1RedeemTokenResponseWithDefaults instantiates a new CdxV1RedeemTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1RedeemTokenResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1RedeemTokenResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1RedeemTokenResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1RedeemTokenResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1RedeemTokenResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1RedeemTokenResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1RedeemTokenResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1RedeemTokenResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1RedeemTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1RedeemTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1RedeemTokenResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1RedeemTokenResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1RedeemTokenResponse) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1RedeemTokenResponse) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1RedeemTokenResponse) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1RedeemTokenResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetApiKey

`func (o *CdxV1RedeemTokenResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CdxV1RedeemTokenResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CdxV1RedeemTokenResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CdxV1RedeemTokenResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecret

`func (o *CdxV1RedeemTokenResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CdxV1RedeemTokenResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CdxV1RedeemTokenResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CdxV1RedeemTokenResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetKafkaBootstrapUrl

`func (o *CdxV1RedeemTokenResponse) GetKafkaBootstrapUrl() string`

GetKafkaBootstrapUrl returns the KafkaBootstrapUrl field if non-nil, zero value otherwise.

### GetKafkaBootstrapUrlOk

`func (o *CdxV1RedeemTokenResponse) GetKafkaBootstrapUrlOk() (*string, bool)`

GetKafkaBootstrapUrlOk returns a tuple with the KafkaBootstrapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBootstrapUrl

`func (o *CdxV1RedeemTokenResponse) SetKafkaBootstrapUrl(v string)`

SetKafkaBootstrapUrl sets KafkaBootstrapUrl field to given value.

### HasKafkaBootstrapUrl

`func (o *CdxV1RedeemTokenResponse) HasKafkaBootstrapUrl() bool`

HasKafkaBootstrapUrl returns a boolean if a field has been set.

### GetSchemaRegistryApiKey

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryApiKey() string`

GetSchemaRegistryApiKey returns the SchemaRegistryApiKey field if non-nil, zero value otherwise.

### GetSchemaRegistryApiKeyOk

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryApiKeyOk() (*string, bool)`

GetSchemaRegistryApiKeyOk returns a tuple with the SchemaRegistryApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryApiKey

`func (o *CdxV1RedeemTokenResponse) SetSchemaRegistryApiKey(v string)`

SetSchemaRegistryApiKey sets SchemaRegistryApiKey field to given value.

### HasSchemaRegistryApiKey

`func (o *CdxV1RedeemTokenResponse) HasSchemaRegistryApiKey() bool`

HasSchemaRegistryApiKey returns a boolean if a field has been set.

### GetSchemaRegistrySecret

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistrySecret() string`

GetSchemaRegistrySecret returns the SchemaRegistrySecret field if non-nil, zero value otherwise.

### GetSchemaRegistrySecretOk

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistrySecretOk() (*string, bool)`

GetSchemaRegistrySecretOk returns a tuple with the SchemaRegistrySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistrySecret

`func (o *CdxV1RedeemTokenResponse) SetSchemaRegistrySecret(v string)`

SetSchemaRegistrySecret sets SchemaRegistrySecret field to given value.

### HasSchemaRegistrySecret

`func (o *CdxV1RedeemTokenResponse) HasSchemaRegistrySecret() bool`

HasSchemaRegistrySecret returns a boolean if a field has been set.

### GetSchemaRegistryUrl

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryUrl() string`

GetSchemaRegistryUrl returns the SchemaRegistryUrl field if non-nil, zero value otherwise.

### GetSchemaRegistryUrlOk

`func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryUrlOk() (*string, bool)`

GetSchemaRegistryUrlOk returns a tuple with the SchemaRegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryUrl

`func (o *CdxV1RedeemTokenResponse) SetSchemaRegistryUrl(v string)`

SetSchemaRegistryUrl sets SchemaRegistryUrl field to given value.

### HasSchemaRegistryUrl

`func (o *CdxV1RedeemTokenResponse) HasSchemaRegistryUrl() bool`

HasSchemaRegistryUrl returns a boolean if a field has been set.

### GetResources

`func (o *CdxV1RedeemTokenResponse) GetResources() []CdxV1RedeemTokenResponseResourcesOneOf`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1RedeemTokenResponse) GetResourcesOk() (*[]CdxV1RedeemTokenResponseResourcesOneOf, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1RedeemTokenResponse) SetResources(v []CdxV1RedeemTokenResponseResourcesOneOf)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1RedeemTokenResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


