# CdxV1RedeemShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Apikey** | Pointer to **string** | The api key | [optional] 
**Secret** | Pointer to **string** | The api key secret | [optional] 
**KafkaBootstrapUrl** | Pointer to **string** | The cluster connection url | [optional] 
**Resources** | Pointer to [**[]CdxV1RedeemShareResourcesOneOf**](CdxV1RedeemShareResourcesOneOf.md) | List of shared resources | [optional] 

## Methods

### NewCdxV1RedeemShare

`func NewCdxV1RedeemShare() *CdxV1RedeemShare`

NewCdxV1RedeemShare instantiates a new CdxV1RedeemShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1RedeemShareWithDefaults

`func NewCdxV1RedeemShareWithDefaults() *CdxV1RedeemShare`

NewCdxV1RedeemShareWithDefaults instantiates a new CdxV1RedeemShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1RedeemShare) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1RedeemShare) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1RedeemShare) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1RedeemShare) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1RedeemShare) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1RedeemShare) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1RedeemShare) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1RedeemShare) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1RedeemShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1RedeemShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1RedeemShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1RedeemShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1RedeemShare) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1RedeemShare) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1RedeemShare) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1RedeemShare) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetApikey

`func (o *CdxV1RedeemShare) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CdxV1RedeemShare) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CdxV1RedeemShare) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *CdxV1RedeemShare) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetSecret

`func (o *CdxV1RedeemShare) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CdxV1RedeemShare) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CdxV1RedeemShare) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CdxV1RedeemShare) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetKafkaBootstrapUrl

`func (o *CdxV1RedeemShare) GetKafkaBootstrapUrl() string`

GetKafkaBootstrapUrl returns the KafkaBootstrapUrl field if non-nil, zero value otherwise.

### GetKafkaBootstrapUrlOk

`func (o *CdxV1RedeemShare) GetKafkaBootstrapUrlOk() (*string, bool)`

GetKafkaBootstrapUrlOk returns a tuple with the KafkaBootstrapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBootstrapUrl

`func (o *CdxV1RedeemShare) SetKafkaBootstrapUrl(v string)`

SetKafkaBootstrapUrl sets KafkaBootstrapUrl field to given value.

### HasKafkaBootstrapUrl

`func (o *CdxV1RedeemShare) HasKafkaBootstrapUrl() bool`

HasKafkaBootstrapUrl returns a boolean if a field has been set.

### GetResources

`func (o *CdxV1RedeemShare) GetResources() []CdxV1RedeemShareResourcesOneOf`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1RedeemShare) GetResourcesOk() (*[]CdxV1RedeemShareResourcesOneOf, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1RedeemShare) SetResources(v []CdxV1RedeemShareResourcesOneOf)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1RedeemShare) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


