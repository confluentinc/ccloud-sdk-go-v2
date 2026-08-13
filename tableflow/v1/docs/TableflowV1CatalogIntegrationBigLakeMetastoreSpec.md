# TableflowV1CatalogIntegrationBigLakeMetastoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**ProviderIntegrationId** | **string** | The provider integration id. | 
**GcpProjectId** | **string** | The GCP project id that hosts the BigLake Metastore catalog. | 
**CatalogName** | **string** | The name of the catalog within BigLake Metastore. | 
**CustomNamespace** | Pointer to **string** | The custom namespace to use in BigLake Metastore. | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationBigLakeMetastoreSpec

`func NewTableflowV1CatalogIntegrationBigLakeMetastoreSpec(kind string, providerIntegrationId string, gcpProjectId string, catalogName string, ) *TableflowV1CatalogIntegrationBigLakeMetastoreSpec`

NewTableflowV1CatalogIntegrationBigLakeMetastoreSpec instantiates a new TableflowV1CatalogIntegrationBigLakeMetastoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationBigLakeMetastoreSpecWithDefaults

`func NewTableflowV1CatalogIntegrationBigLakeMetastoreSpecWithDefaults() *TableflowV1CatalogIntegrationBigLakeMetastoreSpec`

NewTableflowV1CatalogIntegrationBigLakeMetastoreSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationBigLakeMetastoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProviderIntegrationId

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetProviderIntegrationId() string`

GetProviderIntegrationId returns the ProviderIntegrationId field if non-nil, zero value otherwise.

### GetProviderIntegrationIdOk

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetProviderIntegrationIdOk() (*string, bool)`

GetProviderIntegrationIdOk returns a tuple with the ProviderIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIntegrationId

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) SetProviderIntegrationId(v string)`

SetProviderIntegrationId sets ProviderIntegrationId field to given value.


### GetGcpProjectId

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.


### GetCatalogName

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetCustomNamespace

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetCustomNamespace() string`

GetCustomNamespace returns the CustomNamespace field if non-nil, zero value otherwise.

### GetCustomNamespaceOk

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) GetCustomNamespaceOk() (*string, bool)`

GetCustomNamespaceOk returns a tuple with the CustomNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNamespace

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) SetCustomNamespace(v string)`

SetCustomNamespace sets CustomNamespace field to given value.

### HasCustomNamespace

`func (o *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) HasCustomNamespace() bool`

HasCustomNamespace returns a boolean if a field has been set.


### AsTableflowV1CatalogIntegrationSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationBigLakeMetastoreSpec) AsTableflowV1CatalogIntegrationSpecConfigOneOf() TableflowV1CatalogIntegrationSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationBigLakeMetastoreSpec in TableflowV1CatalogIntegrationSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


