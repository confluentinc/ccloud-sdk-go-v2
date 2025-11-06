# TableflowV1AzureAdlsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The storage type.  | 
**StorageAccountName** | **string** | Storage Account Name | 
**ContainerName** | **string** | Container name | 
**StorageRegion** | Pointer to **string** | Storage account region | [optional] [readonly] 
**ProviderIntegrationId** | **string** | The provider integration id | 
**TablePath** | Pointer to **string** | The current storage path where the data and metadata is stored for this table | [optional] [readonly] 

## Methods

### NewTableflowV1AzureAdlsSpec

`func NewTableflowV1AzureAdlsSpec(kind string, storageAccountName string, containerName string, providerIntegrationId string, ) *TableflowV1AzureAdlsSpec`

NewTableflowV1AzureAdlsSpec instantiates a new TableflowV1AzureAdlsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1AzureAdlsSpecWithDefaults

`func NewTableflowV1AzureAdlsSpecWithDefaults() *TableflowV1AzureAdlsSpec`

NewTableflowV1AzureAdlsSpecWithDefaults instantiates a new TableflowV1AzureAdlsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1AzureAdlsSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1AzureAdlsSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1AzureAdlsSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStorageAccountName

`func (o *TableflowV1AzureAdlsSpec) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *TableflowV1AzureAdlsSpec) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *TableflowV1AzureAdlsSpec) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### GetContainerName

`func (o *TableflowV1AzureAdlsSpec) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *TableflowV1AzureAdlsSpec) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *TableflowV1AzureAdlsSpec) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetStorageRegion

`func (o *TableflowV1AzureAdlsSpec) GetStorageRegion() string`

GetStorageRegion returns the StorageRegion field if non-nil, zero value otherwise.

### GetStorageRegionOk

`func (o *TableflowV1AzureAdlsSpec) GetStorageRegionOk() (*string, bool)`

GetStorageRegionOk returns a tuple with the StorageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRegion

`func (o *TableflowV1AzureAdlsSpec) SetStorageRegion(v string)`

SetStorageRegion sets StorageRegion field to given value.

### HasStorageRegion

`func (o *TableflowV1AzureAdlsSpec) HasStorageRegion() bool`

HasStorageRegion returns a boolean if a field has been set.

### GetProviderIntegrationId

`func (o *TableflowV1AzureAdlsSpec) GetProviderIntegrationId() string`

GetProviderIntegrationId returns the ProviderIntegrationId field if non-nil, zero value otherwise.

### GetProviderIntegrationIdOk

`func (o *TableflowV1AzureAdlsSpec) GetProviderIntegrationIdOk() (*string, bool)`

GetProviderIntegrationIdOk returns a tuple with the ProviderIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIntegrationId

`func (o *TableflowV1AzureAdlsSpec) SetProviderIntegrationId(v string)`

SetProviderIntegrationId sets ProviderIntegrationId field to given value.


### GetTablePath

`func (o *TableflowV1AzureAdlsSpec) GetTablePath() string`

GetTablePath returns the TablePath field if non-nil, zero value otherwise.

### GetTablePathOk

`func (o *TableflowV1AzureAdlsSpec) GetTablePathOk() (*string, bool)`

GetTablePathOk returns a tuple with the TablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePath

`func (o *TableflowV1AzureAdlsSpec) SetTablePath(v string)`

SetTablePath sets TablePath field to given value.

### HasTablePath

`func (o *TableflowV1AzureAdlsSpec) HasTablePath() bool`

HasTablePath returns a boolean if a field has been set.


### AsTableflowV1TableflowTopicSpecStorageOneOf

`func (s *TableflowV1AzureAdlsSpec) AsTableflowV1TableflowTopicSpecStorageOneOf() TableflowV1TableflowTopicSpecStorageOneOf`

Convenience method to wrap this instance of TableflowV1AzureAdlsSpec in TableflowV1TableflowTopicSpecStorageOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


