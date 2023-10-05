# FcpmV2SecretUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Secret** | Pointer to **string** | The value of the customer provided secret. Only required in &#x60;create&#x60;, &#x60;update&#x60; requests, will be ignored in in &#x60;get&#x60; or &#x60;list&#x60;.  | [optional] 
**Description** | Pointer to **string** | A human readable description for the secret  | [optional] 

## Methods

### NewFcpmV2SecretUpdate

`func NewFcpmV2SecretUpdate() *FcpmV2SecretUpdate`

NewFcpmV2SecretUpdate instantiates a new FcpmV2SecretUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2SecretUpdateWithDefaults

`func NewFcpmV2SecretUpdateWithDefaults() *FcpmV2SecretUpdate`

NewFcpmV2SecretUpdateWithDefaults instantiates a new FcpmV2SecretUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2SecretUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2SecretUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2SecretUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2SecretUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2SecretUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2SecretUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2SecretUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2SecretUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2SecretUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2SecretUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2SecretUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2SecretUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2SecretUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2SecretUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2SecretUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2SecretUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecret

`func (o *FcpmV2SecretUpdate) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *FcpmV2SecretUpdate) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *FcpmV2SecretUpdate) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *FcpmV2SecretUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDescription

`func (o *FcpmV2SecretUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FcpmV2SecretUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FcpmV2SecretUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FcpmV2SecretUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


