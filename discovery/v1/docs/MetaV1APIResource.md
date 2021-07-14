# MetaV1APIResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ResourceKind** | Pointer to **string** | The resource kind | [optional] 
**ResourceVersion** | Pointer to [**MetaV1GroupVersion**](meta.v1.GroupVersion.md) |  | [optional] 
**PluralName** | Pointer to **string** | The resource name | [optional] 
**SingularName** | Pointer to **string** | The resource name | [optional] 
**Verbs** | Pointer to **[]string** | A list of supported operations for this resource | [optional] 
**ShortNames** | Pointer to **[]string** | A list of suggested short names for the resource | [optional] 

## Methods

### NewMetaV1APIResource

`func NewMetaV1APIResource() *MetaV1APIResource`

NewMetaV1APIResource instantiates a new MetaV1APIResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaV1APIResourceWithDefaults

`func NewMetaV1APIResourceWithDefaults() *MetaV1APIResource`

NewMetaV1APIResourceWithDefaults instantiates a new MetaV1APIResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MetaV1APIResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MetaV1APIResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MetaV1APIResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MetaV1APIResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MetaV1APIResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetaV1APIResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetaV1APIResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetaV1APIResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *MetaV1APIResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaV1APIResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaV1APIResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetaV1APIResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *MetaV1APIResource) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaV1APIResource) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaV1APIResource) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetaV1APIResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceKind

`func (o *MetaV1APIResource) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *MetaV1APIResource) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *MetaV1APIResource) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *MetaV1APIResource) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetResourceVersion

`func (o *MetaV1APIResource) GetResourceVersion() MetaV1GroupVersion`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *MetaV1APIResource) GetResourceVersionOk() (*MetaV1GroupVersion, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *MetaV1APIResource) SetResourceVersion(v MetaV1GroupVersion)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *MetaV1APIResource) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetPluralName

`func (o *MetaV1APIResource) GetPluralName() string`

GetPluralName returns the PluralName field if non-nil, zero value otherwise.

### GetPluralNameOk

`func (o *MetaV1APIResource) GetPluralNameOk() (*string, bool)`

GetPluralNameOk returns a tuple with the PluralName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralName

`func (o *MetaV1APIResource) SetPluralName(v string)`

SetPluralName sets PluralName field to given value.

### HasPluralName

`func (o *MetaV1APIResource) HasPluralName() bool`

HasPluralName returns a boolean if a field has been set.

### GetSingularName

`func (o *MetaV1APIResource) GetSingularName() string`

GetSingularName returns the SingularName field if non-nil, zero value otherwise.

### GetSingularNameOk

`func (o *MetaV1APIResource) GetSingularNameOk() (*string, bool)`

GetSingularNameOk returns a tuple with the SingularName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularName

`func (o *MetaV1APIResource) SetSingularName(v string)`

SetSingularName sets SingularName field to given value.

### HasSingularName

`func (o *MetaV1APIResource) HasSingularName() bool`

HasSingularName returns a boolean if a field has been set.

### GetVerbs

`func (o *MetaV1APIResource) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *MetaV1APIResource) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *MetaV1APIResource) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.

### HasVerbs

`func (o *MetaV1APIResource) HasVerbs() bool`

HasVerbs returns a boolean if a field has been set.

### GetShortNames

`func (o *MetaV1APIResource) GetShortNames() []string`

GetShortNames returns the ShortNames field if non-nil, zero value otherwise.

### GetShortNamesOk

`func (o *MetaV1APIResource) GetShortNamesOk() (*[]string, bool)`

GetShortNamesOk returns a tuple with the ShortNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNames

`func (o *MetaV1APIResource) SetShortNames(v []string)`

SetShortNames sets ShortNames field to given value.

### HasShortNames

`func (o *MetaV1APIResource) HasShortNames() bool`

HasShortNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


