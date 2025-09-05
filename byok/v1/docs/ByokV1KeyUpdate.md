# ByokV1KeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The human-readable name of the key object.  | [optional] 
**Provider** | Pointer to **string** | The cloud provider of the Key. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the key:    AVAILABLE: key can be used for a Kafka cluster provisioning.    IN_USE: key is already in use by a Kafka cluster provisioning.  | [optional] [readonly] 
**Validation** | Pointer to [**ByokV1KeyValidation**](byok.v1.KeyValidation.md) | The validation details of the key.  | [optional] [readonly] 

## Methods

### NewByokV1KeyUpdate

`func NewByokV1KeyUpdate() *ByokV1KeyUpdate`

NewByokV1KeyUpdate instantiates a new ByokV1KeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyUpdateWithDefaults

`func NewByokV1KeyUpdateWithDefaults() *ByokV1KeyUpdate`

NewByokV1KeyUpdateWithDefaults instantiates a new ByokV1KeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ByokV1KeyUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ByokV1KeyUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ByokV1KeyUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ByokV1KeyUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1KeyUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1KeyUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1KeyUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ByokV1KeyUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ByokV1KeyUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ByokV1KeyUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ByokV1KeyUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ByokV1KeyUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ByokV1KeyUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ByokV1KeyUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ByokV1KeyUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ByokV1KeyUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ByokV1KeyUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ByokV1KeyUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ByokV1KeyUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ByokV1KeyUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProvider

`func (o *ByokV1KeyUpdate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ByokV1KeyUpdate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ByokV1KeyUpdate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ByokV1KeyUpdate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *ByokV1KeyUpdate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ByokV1KeyUpdate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ByokV1KeyUpdate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ByokV1KeyUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValidation

`func (o *ByokV1KeyUpdate) GetValidation() ByokV1KeyValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ByokV1KeyUpdate) GetValidationOk() (*ByokV1KeyValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ByokV1KeyUpdate) SetValidation(v ByokV1KeyValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *ByokV1KeyUpdate) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


