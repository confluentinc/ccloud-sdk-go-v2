# BusinessMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The business metadata name | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | The business metadata attributes | [optional] 
**EntityType** | Pointer to **string** | The entity type | [optional] 
**EntityName** | Pointer to **string** | The qualified name of the entity | [optional] 

## Methods

### NewBusinessMetadata

`func NewBusinessMetadata() *BusinessMetadata`

NewBusinessMetadata instantiates a new BusinessMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessMetadataWithDefaults

`func NewBusinessMetadataWithDefaults() *BusinessMetadata`

NewBusinessMetadataWithDefaults instantiates a new BusinessMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *BusinessMetadata) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BusinessMetadata) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BusinessMetadata) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BusinessMetadata) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *BusinessMetadata) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BusinessMetadata) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BusinessMetadata) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BusinessMetadata) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityType

`func (o *BusinessMetadata) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BusinessMetadata) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BusinessMetadata) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BusinessMetadata) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityName

`func (o *BusinessMetadata) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *BusinessMetadata) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *BusinessMetadata) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *BusinessMetadata) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


