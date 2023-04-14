# BusinessMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The business metadata name | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | The business metadata attributes | [optional] 
**EntityType** | Pointer to **string** | The entity type | [optional] 
**EntityName** | Pointer to **string** | The qualified name of the entity | [optional] 
**Error** | Pointer to [**ErrorMessage**](ErrorMessage.md) |  | [optional] 

## Methods

### NewBusinessMetadataResponse

`func NewBusinessMetadataResponse() *BusinessMetadataResponse`

NewBusinessMetadataResponse instantiates a new BusinessMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessMetadataResponseWithDefaults

`func NewBusinessMetadataResponseWithDefaults() *BusinessMetadataResponse`

NewBusinessMetadataResponseWithDefaults instantiates a new BusinessMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *BusinessMetadataResponse) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *BusinessMetadataResponse) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *BusinessMetadataResponse) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *BusinessMetadataResponse) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *BusinessMetadataResponse) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BusinessMetadataResponse) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BusinessMetadataResponse) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BusinessMetadataResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityType

`func (o *BusinessMetadataResponse) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BusinessMetadataResponse) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BusinessMetadataResponse) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BusinessMetadataResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityName

`func (o *BusinessMetadataResponse) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *BusinessMetadataResponse) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *BusinessMetadataResponse) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *BusinessMetadataResponse) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetError

`func (o *BusinessMetadataResponse) GetError() ErrorMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BusinessMetadataResponse) GetErrorOk() (*ErrorMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BusinessMetadataResponse) SetError(v ErrorMessage)`

SetError sets Error field to given value.

### HasError

`func (o *BusinessMetadataResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


