# TableflowV1CatalogIntegrationUnitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**WorkspaceEndpoint** | **string** | The Databricks workspace URL associated with the Unity Catalog. | 
**CatalogName** | **string** | The name of the catalog within Unity Catalog. | 
**ClientId** | **string** | The OAuth client ID used to authenticate with the Unity Catalog. | 
**ClientSecret** | **string** | The OAuth client secret used for authentication with the Unity Catalog. | 

## Methods

### NewTableflowV1CatalogIntegrationUnitySpec

`func NewTableflowV1CatalogIntegrationUnitySpec(kind string, workspaceEndpoint string, catalogName string, clientId string, clientSecret string, ) *TableflowV1CatalogIntegrationUnitySpec`

NewTableflowV1CatalogIntegrationUnitySpec instantiates a new TableflowV1CatalogIntegrationUnitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationUnitySpecWithDefaults

`func NewTableflowV1CatalogIntegrationUnitySpecWithDefaults() *TableflowV1CatalogIntegrationUnitySpec`

NewTableflowV1CatalogIntegrationUnitySpecWithDefaults instantiates a new TableflowV1CatalogIntegrationUnitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationUnitySpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetWorkspaceEndpoint

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetWorkspaceEndpoint() string`

GetWorkspaceEndpoint returns the WorkspaceEndpoint field if non-nil, zero value otherwise.

### GetWorkspaceEndpointOk

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetWorkspaceEndpointOk() (*string, bool)`

GetWorkspaceEndpointOk returns a tuple with the WorkspaceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceEndpoint

`func (o *TableflowV1CatalogIntegrationUnitySpec) SetWorkspaceEndpoint(v string)`

SetWorkspaceEndpoint sets WorkspaceEndpoint field to given value.


### GetCatalogName

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *TableflowV1CatalogIntegrationUnitySpec) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetClientId

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TableflowV1CatalogIntegrationUnitySpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TableflowV1CatalogIntegrationUnitySpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TableflowV1CatalogIntegrationUnitySpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



### AsTableflowV1CatalogIntegrationSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationUnitySpec) AsTableflowV1CatalogIntegrationSpecConfigOneOf() TableflowV1CatalogIntegrationSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationUnitySpec in TableflowV1CatalogIntegrationSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


