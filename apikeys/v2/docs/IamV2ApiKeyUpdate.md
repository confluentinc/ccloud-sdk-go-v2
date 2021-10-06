# IamV2ApiKeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Secret** | Pointer to **string** | The API key secret. Only provided in &#x60;create&#x60; responses, not in &#x60;get&#x60; or &#x60;list&#x60;. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human readable name for the API key | [optional] 
**Description** | Pointer to **string** | A human readable description for the API key | [optional] 

## Methods

### NewIamV2ApiKeyUpdate

`func NewIamV2ApiKeyUpdate() *IamV2ApiKeyUpdate`

NewIamV2ApiKeyUpdate instantiates a new IamV2ApiKeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ApiKeyUpdateWithDefaults

`func NewIamV2ApiKeyUpdateWithDefaults() *IamV2ApiKeyUpdate`

NewIamV2ApiKeyUpdateWithDefaults instantiates a new IamV2ApiKeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2ApiKeyUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2ApiKeyUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2ApiKeyUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2ApiKeyUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2ApiKeyUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2ApiKeyUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2ApiKeyUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2ApiKeyUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2ApiKeyUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2ApiKeyUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2ApiKeyUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2ApiKeyUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2ApiKeyUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2ApiKeyUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2ApiKeyUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2ApiKeyUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecret

`func (o *IamV2ApiKeyUpdate) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IamV2ApiKeyUpdate) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IamV2ApiKeyUpdate) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IamV2ApiKeyUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2ApiKeyUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2ApiKeyUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2ApiKeyUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2ApiKeyUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2ApiKeyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2ApiKeyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2ApiKeyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2ApiKeyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


