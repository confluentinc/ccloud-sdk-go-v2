# TableflowV1CatalogIntegrationUnityUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**WorkspaceEndpoint** | Pointer to **string** | The Databricks workspace URL associated with the Unity Catalog. | [optional] 
**CatalogName** | Pointer to **string** | The name of the catalog within Unity Catalog. | [optional] 
**ClientId** | Pointer to **string** | The OAuth client ID used to authenticate with the Unity Catalog. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth client secret used for authentication with the Unity Catalog. | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationUnityUpdateSpec

`func NewTableflowV1CatalogIntegrationUnityUpdateSpec(kind string, ) *TableflowV1CatalogIntegrationUnityUpdateSpec`

NewTableflowV1CatalogIntegrationUnityUpdateSpec instantiates a new TableflowV1CatalogIntegrationUnityUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationUnityUpdateSpecWithDefaults

`func NewTableflowV1CatalogIntegrationUnityUpdateSpecWithDefaults() *TableflowV1CatalogIntegrationUnityUpdateSpec`

NewTableflowV1CatalogIntegrationUnityUpdateSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationUnityUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetWorkspaceEndpoint

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetWorkspaceEndpoint() string`

GetWorkspaceEndpoint returns the WorkspaceEndpoint field if non-nil, zero value otherwise.

### GetWorkspaceEndpointOk

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetWorkspaceEndpointOk() (*string, bool)`

GetWorkspaceEndpointOk returns a tuple with the WorkspaceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceEndpoint

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) SetWorkspaceEndpoint(v string)`

SetWorkspaceEndpoint sets WorkspaceEndpoint field to given value.

### HasWorkspaceEndpoint

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) HasWorkspaceEndpoint() bool`

HasWorkspaceEndpoint returns a boolean if a field has been set.

### GetCatalogName

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetClientId

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TableflowV1CatalogIntegrationUnityUpdateSpec) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


### AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationUnityUpdateSpec) AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf() TableflowV1CatalogIntegrationUpdateSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationUnityUpdateSpec in TableflowV1CatalogIntegrationUpdateSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


