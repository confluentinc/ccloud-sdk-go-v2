# TableflowV1CatalogIntegrationSnowflakeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**Endpoint** | **string** | The catalog integration connection endpoint for Snowflake Open Catalog.  | 
**ClientId** | **string** | The client ID of the catalog integration. | 
**ClientSecret** | **string** | The client secret of the catalog integration. | 
**Warehouse** | **string** | Warehouse name of the Snowflake Open Catalog. | 
**AllowedScope** | **string** | Allowed scope of the Snowflake Open Catalog. | 

## Methods

### NewTableflowV1CatalogIntegrationSnowflakeSpec

`func NewTableflowV1CatalogIntegrationSnowflakeSpec(kind string, endpoint string, clientId string, clientSecret string, warehouse string, allowedScope string, ) *TableflowV1CatalogIntegrationSnowflakeSpec`

NewTableflowV1CatalogIntegrationSnowflakeSpec instantiates a new TableflowV1CatalogIntegrationSnowflakeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationSnowflakeSpecWithDefaults

`func NewTableflowV1CatalogIntegrationSnowflakeSpecWithDefaults() *TableflowV1CatalogIntegrationSnowflakeSpec`

NewTableflowV1CatalogIntegrationSnowflakeSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationSnowflakeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEndpoint

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetClientId

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetWarehouse

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.


### GetAllowedScope

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetAllowedScope() string`

GetAllowedScope returns the AllowedScope field if non-nil, zero value otherwise.

### GetAllowedScopeOk

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) GetAllowedScopeOk() (*string, bool)`

GetAllowedScopeOk returns a tuple with the AllowedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScope

`func (o *TableflowV1CatalogIntegrationSnowflakeSpec) SetAllowedScope(v string)`

SetAllowedScope sets AllowedScope field to given value.



### AsTableflowV1CatalogIntegrationSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationSnowflakeSpec) AsTableflowV1CatalogIntegrationSpecConfigOneOf() TableflowV1CatalogIntegrationSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationSnowflakeSpec in TableflowV1CatalogIntegrationSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


