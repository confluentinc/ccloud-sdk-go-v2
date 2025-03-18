# TableflowV1CatalogIntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]TableflowV1CatalogIntegration**](TableflowV1CatalogIntegration.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewTableflowV1CatalogIntegrationList

`func NewTableflowV1CatalogIntegrationList(apiVersion string, kind string, metadata ListMeta, data []TableflowV1CatalogIntegration, ) *TableflowV1CatalogIntegrationList`

NewTableflowV1CatalogIntegrationList instantiates a new TableflowV1CatalogIntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationListWithDefaults

`func NewTableflowV1CatalogIntegrationListWithDefaults() *TableflowV1CatalogIntegrationList`

NewTableflowV1CatalogIntegrationListWithDefaults instantiates a new TableflowV1CatalogIntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TableflowV1CatalogIntegrationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TableflowV1CatalogIntegrationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TableflowV1CatalogIntegrationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *TableflowV1CatalogIntegrationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TableflowV1CatalogIntegrationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TableflowV1CatalogIntegrationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TableflowV1CatalogIntegrationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *TableflowV1CatalogIntegrationList) GetData() []TableflowV1CatalogIntegration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TableflowV1CatalogIntegrationList) GetDataOk() (*[]TableflowV1CatalogIntegration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TableflowV1CatalogIntegrationList) SetData(v []TableflowV1CatalogIntegration)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


