# TableflowV1CatalogIntegrationSnowflakeUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**Endpoint** | Pointer to **string** | The catalog integration connection endpoint for Snowflake Open Catalog. | [optional] 
**ClientId** | Pointer to **string** | The client ID of the catalog integration. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret of the catalog integration. | [optional] 
**Warehouse** | Pointer to **string** | Warehouse name of the Snowflake Open Catalog. | [optional] 
**AllowedScope** | Pointer to **string** | Allowed scope of the Snowflake Open Catalog. | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationSnowflakeUpdateSpec

`func NewTableflowV1CatalogIntegrationSnowflakeUpdateSpec(kind string, ) *TableflowV1CatalogIntegrationSnowflakeUpdateSpec`

NewTableflowV1CatalogIntegrationSnowflakeUpdateSpec instantiates a new TableflowV1CatalogIntegrationSnowflakeUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationSnowflakeUpdateSpecWithDefaults

`func NewTableflowV1CatalogIntegrationSnowflakeUpdateSpecWithDefaults() *TableflowV1CatalogIntegrationSnowflakeUpdateSpec`

NewTableflowV1CatalogIntegrationSnowflakeUpdateSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationSnowflakeUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEndpoint

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetClientId

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetWarehouse

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetAllowedScope

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetAllowedScope() string`

GetAllowedScope returns the AllowedScope field if non-nil, zero value otherwise.

### GetAllowedScopeOk

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) GetAllowedScopeOk() (*string, bool)`

GetAllowedScopeOk returns a tuple with the AllowedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScope

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) SetAllowedScope(v string)`

SetAllowedScope sets AllowedScope field to given value.

### HasAllowedScope

`func (o *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) HasAllowedScope() bool`

HasAllowedScope returns a boolean if a field has been set.


### AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationSnowflakeUpdateSpec) AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf() TableflowV1CatalogIntegrationUpdateSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationSnowflakeUpdateSpec in TableflowV1CatalogIntegrationUpdateSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


