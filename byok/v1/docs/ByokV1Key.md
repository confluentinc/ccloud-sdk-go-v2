# ByokV1Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Key** | Pointer to [**ByokV1KeyKeyOneOf**](ByokV1KeyKeyOneOf.md) | The cloud-specific key details.  For AWS please provide the corresponding &#x60;key_arn&#x60;. For Azure please provide the corresponding &#x60;key_id&#x60;.  | [optional] 
**Provider** | Pointer to **string** | The cloud provider of the Key. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the key:   AVAILABLE: key can be used for a Kafka cluster provisioning   IN_USE: key is already in use by a Kafka cluster provisioning  | [optional] [readonly] 

## Methods

### NewByokV1Key

`func NewByokV1Key() *ByokV1Key`

NewByokV1Key instantiates a new ByokV1Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyWithDefaults

`func NewByokV1KeyWithDefaults() *ByokV1Key`

NewByokV1KeyWithDefaults instantiates a new ByokV1Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ByokV1Key) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ByokV1Key) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ByokV1Key) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ByokV1Key) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1Key) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1Key) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1Key) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ByokV1Key) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ByokV1Key) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ByokV1Key) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ByokV1Key) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ByokV1Key) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ByokV1Key) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ByokV1Key) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ByokV1Key) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ByokV1Key) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *ByokV1Key) GetKey() ByokV1KeyKeyOneOf`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ByokV1Key) GetKeyOk() (*ByokV1KeyKeyOneOf, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ByokV1Key) SetKey(v ByokV1KeyKeyOneOf)`

SetKey sets Key field to given value.

### HasKey

`func (o *ByokV1Key) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProvider

`func (o *ByokV1Key) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ByokV1Key) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ByokV1Key) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ByokV1Key) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *ByokV1Key) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ByokV1Key) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ByokV1Key) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ByokV1Key) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


