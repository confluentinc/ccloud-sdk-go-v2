# FcpmV2Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Key** | Pointer to **string** | The key of the customer provided secret. | [optional] 
**Secret** | Pointer to **string** | The value of the customer provided secret. Only required in &#x60;create&#x60;, &#x60;update&#x60; requests, will be ignored in in &#x60;get&#x60; or &#x60;list&#x60;.  | [optional] 
**Description** | Pointer to **string** | A human readable description for the secret  | [optional] 

## Methods

### NewFcpmV2Secret

`func NewFcpmV2Secret() *FcpmV2Secret`

NewFcpmV2Secret instantiates a new FcpmV2Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2SecretWithDefaults

`func NewFcpmV2SecretWithDefaults() *FcpmV2Secret`

NewFcpmV2SecretWithDefaults instantiates a new FcpmV2Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2Secret) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2Secret) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2Secret) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2Secret) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2Secret) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2Secret) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2Secret) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2Secret) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2Secret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2Secret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2Secret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2Secret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2Secret) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2Secret) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2Secret) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2Secret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *FcpmV2Secret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FcpmV2Secret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FcpmV2Secret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FcpmV2Secret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecret

`func (o *FcpmV2Secret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *FcpmV2Secret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *FcpmV2Secret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *FcpmV2Secret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDescription

`func (o *FcpmV2Secret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FcpmV2Secret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FcpmV2Secret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FcpmV2Secret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


