# TableflowV1CatalogIntegrationAwsGlueUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of the catalog integration. | 
**CustomDatabase** | Pointer to **string** | The custom database name to use in AWS Glue. | [optional] 

## Methods

### NewTableflowV1CatalogIntegrationAwsGlueUpdateSpec

`func NewTableflowV1CatalogIntegrationAwsGlueUpdateSpec(kind string, ) *TableflowV1CatalogIntegrationAwsGlueUpdateSpec`

NewTableflowV1CatalogIntegrationAwsGlueUpdateSpec instantiates a new TableflowV1CatalogIntegrationAwsGlueUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationAwsGlueUpdateSpecWithDefaults

`func NewTableflowV1CatalogIntegrationAwsGlueUpdateSpecWithDefaults() *TableflowV1CatalogIntegrationAwsGlueUpdateSpec`

NewTableflowV1CatalogIntegrationAwsGlueUpdateSpecWithDefaults instantiates a new TableflowV1CatalogIntegrationAwsGlueUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCustomDatabase

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) GetCustomDatabase() string`

GetCustomDatabase returns the CustomDatabase field if non-nil, zero value otherwise.

### GetCustomDatabaseOk

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) GetCustomDatabaseOk() (*string, bool)`

GetCustomDatabaseOk returns a tuple with the CustomDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatabase

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) SetCustomDatabase(v string)`

SetCustomDatabase sets CustomDatabase field to given value.

### HasCustomDatabase

`func (o *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) HasCustomDatabase() bool`

HasCustomDatabase returns a boolean if a field has been set.


### AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf

`func (s *TableflowV1CatalogIntegrationAwsGlueUpdateSpec) AsTableflowV1CatalogIntegrationUpdateSpecConfigOneOf() TableflowV1CatalogIntegrationUpdateSpecConfigOneOf`

Convenience method to wrap this instance of TableflowV1CatalogIntegrationAwsGlueUpdateSpec in TableflowV1CatalogIntegrationUpdateSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


