# SchemaEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityPath** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 

## Methods

### NewSchemaEntity

`func NewSchemaEntity() *SchemaEntity`

NewSchemaEntity instantiates a new SchemaEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaEntityWithDefaults

`func NewSchemaEntityWithDefaults() *SchemaEntity`

NewSchemaEntityWithDefaults instantiates a new SchemaEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityPath

`func (o *SchemaEntity) GetEntityPath() string`

GetEntityPath returns the EntityPath field if non-nil, zero value otherwise.

### GetEntityPathOk

`func (o *SchemaEntity) GetEntityPathOk() (*string, bool)`

GetEntityPathOk returns a tuple with the EntityPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityPath

`func (o *SchemaEntity) SetEntityPath(v string)`

SetEntityPath sets EntityPath field to given value.

### HasEntityPath

`func (o *SchemaEntity) HasEntityPath() bool`

HasEntityPath returns a boolean if a field has been set.

### GetEntityType

`func (o *SchemaEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SchemaEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SchemaEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *SchemaEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


