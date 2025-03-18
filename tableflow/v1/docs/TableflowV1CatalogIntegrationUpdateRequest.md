# TableflowV1CatalogIntegrationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**TableflowV1CatalogIntegrationUpdateSpec**](tableflow.v1.CatalogIntegrationUpdateSpec.md) |  | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationUpdateRequest

`func NewTableflowV1CatalogIntegrationUpdateRequest() *TableflowV1CatalogIntegrationUpdateRequest`

NewTableflowV1CatalogIntegrationUpdateRequest instantiates a new TableflowV1CatalogIntegrationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationUpdateRequestWithDefaults

`func NewTableflowV1CatalogIntegrationUpdateRequestWithDefaults() *TableflowV1CatalogIntegrationUpdateRequest`

NewTableflowV1CatalogIntegrationUpdateRequestWithDefaults instantiates a new TableflowV1CatalogIntegrationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TableflowV1CatalogIntegrationUpdateRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TableflowV1CatalogIntegrationUpdateRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationUpdateRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TableflowV1CatalogIntegrationUpdateRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TableflowV1CatalogIntegrationUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TableflowV1CatalogIntegrationUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TableflowV1CatalogIntegrationUpdateRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TableflowV1CatalogIntegrationUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetSpec() TableflowV1CatalogIntegrationUpdateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TableflowV1CatalogIntegrationUpdateRequest) GetSpecOk() (*TableflowV1CatalogIntegrationUpdateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TableflowV1CatalogIntegrationUpdateRequest) SetSpec(v TableflowV1CatalogIntegrationUpdateSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TableflowV1CatalogIntegrationUpdateRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


