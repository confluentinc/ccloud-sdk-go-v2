# MetaV1APIGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The API Group name | [optional] 
**PreferredVersion** | Pointer to [**MetaV1GroupVersion**](meta.v1.GroupVersion.md) |  | [optional] 
**Versions** | Pointer to [**[]MetaV1GroupVersion**](MetaV1GroupVersion.md) |  | [optional] 

## Methods

### NewMetaV1APIGroup

`func NewMetaV1APIGroup() *MetaV1APIGroup`

NewMetaV1APIGroup instantiates a new MetaV1APIGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaV1APIGroupWithDefaults

`func NewMetaV1APIGroupWithDefaults() *MetaV1APIGroup`

NewMetaV1APIGroupWithDefaults instantiates a new MetaV1APIGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MetaV1APIGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MetaV1APIGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MetaV1APIGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MetaV1APIGroup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MetaV1APIGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetaV1APIGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetaV1APIGroup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetaV1APIGroup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *MetaV1APIGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaV1APIGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaV1APIGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetaV1APIGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *MetaV1APIGroup) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaV1APIGroup) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaV1APIGroup) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetaV1APIGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *MetaV1APIGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaV1APIGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaV1APIGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaV1APIGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreferredVersion

`func (o *MetaV1APIGroup) GetPreferredVersion() MetaV1GroupVersion`

GetPreferredVersion returns the PreferredVersion field if non-nil, zero value otherwise.

### GetPreferredVersionOk

`func (o *MetaV1APIGroup) GetPreferredVersionOk() (*MetaV1GroupVersion, bool)`

GetPreferredVersionOk returns a tuple with the PreferredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVersion

`func (o *MetaV1APIGroup) SetPreferredVersion(v MetaV1GroupVersion)`

SetPreferredVersion sets PreferredVersion field to given value.

### HasPreferredVersion

`func (o *MetaV1APIGroup) HasPreferredVersion() bool`

HasPreferredVersion returns a boolean if a field has been set.

### GetVersions

`func (o *MetaV1APIGroup) GetVersions() []MetaV1GroupVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MetaV1APIGroup) GetVersionsOk() (*[]MetaV1GroupVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MetaV1APIGroup) SetVersions(v []MetaV1GroupVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MetaV1APIGroup) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


