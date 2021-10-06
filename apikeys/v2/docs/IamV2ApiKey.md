# IamV2ApiKey

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
**Owner** | Pointer to [**ObjectReference**](ObjectReference.md) | The owner to which this belongs. The owner can be one of iam.v2.User, iam.v2.ServiceAccount. | [optional] 
**Resource** | Pointer to [**ObjectReference**](ObjectReference.md) | The resource associated with this object. The resource can be one of cmk.v2.KafkaCluster, apikeys.v2.SchemaRegistry, apikeys.v2.ksqlDB. | [optional] 

## Methods

### NewIamV2ApiKey

`func NewIamV2ApiKey() *IamV2ApiKey`

NewIamV2ApiKey instantiates a new IamV2ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ApiKeyWithDefaults

`func NewIamV2ApiKeyWithDefaults() *IamV2ApiKey`

NewIamV2ApiKeyWithDefaults instantiates a new IamV2ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2ApiKey) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2ApiKey) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2ApiKey) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2ApiKey) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2ApiKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2ApiKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2ApiKey) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2ApiKey) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2ApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2ApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2ApiKey) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2ApiKey) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2ApiKey) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2ApiKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecret

`func (o *IamV2ApiKey) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IamV2ApiKey) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IamV2ApiKey) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IamV2ApiKey) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2ApiKey) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2ApiKey) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2ApiKey) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2ApiKey) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2ApiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2ApiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2ApiKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2ApiKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IamV2ApiKey) GetOwner() ObjectReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamV2ApiKey) GetOwnerOk() (*ObjectReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamV2ApiKey) SetOwner(v ObjectReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamV2ApiKey) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResource

`func (o *IamV2ApiKey) GetResource() ObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamV2ApiKey) GetResourceOk() (*ObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamV2ApiKey) SetResource(v ObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamV2ApiKey) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


