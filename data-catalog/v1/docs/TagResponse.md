# TagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The tag name | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | The tag attributes | [optional] 
**EntityGuid** | Pointer to **string** | The internal entity guid | [optional] 
**EntityStatus** | Pointer to **string** | The entity status | [optional] 
**Propagate** | Pointer to **bool** | Whether to propagate the tag | [optional] 
**ValidityPeriods** | Pointer to [**[]TimeBoundary**](TimeBoundary.md) | The validity periods | [optional] 
**RemovePropagationsOnEntityDelete** | Pointer to **bool** | Whether to remove propagations on entity delete | [optional] 
**EntityType** | Pointer to **string** | The entity type | [optional] 
**EntityName** | Pointer to **string** | The qualified name of the entity | [optional] 
**Error** | Pointer to [**ErrorMessage**](ErrorMessage.md) |  | [optional] 

## Methods

### NewTagResponse

`func NewTagResponse() *TagResponse`

NewTagResponse instantiates a new TagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResponseWithDefaults

`func NewTagResponseWithDefaults() *TagResponse`

NewTagResponseWithDefaults instantiates a new TagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *TagResponse) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *TagResponse) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *TagResponse) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *TagResponse) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *TagResponse) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TagResponse) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TagResponse) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TagResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityGuid

`func (o *TagResponse) GetEntityGuid() string`

GetEntityGuid returns the EntityGuid field if non-nil, zero value otherwise.

### GetEntityGuidOk

`func (o *TagResponse) GetEntityGuidOk() (*string, bool)`

GetEntityGuidOk returns a tuple with the EntityGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityGuid

`func (o *TagResponse) SetEntityGuid(v string)`

SetEntityGuid sets EntityGuid field to given value.

### HasEntityGuid

`func (o *TagResponse) HasEntityGuid() bool`

HasEntityGuid returns a boolean if a field has been set.

### GetEntityStatus

`func (o *TagResponse) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *TagResponse) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *TagResponse) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *TagResponse) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetPropagate

`func (o *TagResponse) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *TagResponse) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *TagResponse) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *TagResponse) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetValidityPeriods

`func (o *TagResponse) GetValidityPeriods() []TimeBoundary`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *TagResponse) GetValidityPeriodsOk() (*[]TimeBoundary, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *TagResponse) SetValidityPeriods(v []TimeBoundary)`

SetValidityPeriods sets ValidityPeriods field to given value.

### HasValidityPeriods

`func (o *TagResponse) HasValidityPeriods() bool`

HasValidityPeriods returns a boolean if a field has been set.

### GetRemovePropagationsOnEntityDelete

`func (o *TagResponse) GetRemovePropagationsOnEntityDelete() bool`

GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field if non-nil, zero value otherwise.

### GetRemovePropagationsOnEntityDeleteOk

`func (o *TagResponse) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool)`

GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePropagationsOnEntityDelete

`func (o *TagResponse) SetRemovePropagationsOnEntityDelete(v bool)`

SetRemovePropagationsOnEntityDelete sets RemovePropagationsOnEntityDelete field to given value.

### HasRemovePropagationsOnEntityDelete

`func (o *TagResponse) HasRemovePropagationsOnEntityDelete() bool`

HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.

### GetEntityType

`func (o *TagResponse) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TagResponse) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TagResponse) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TagResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityName

`func (o *TagResponse) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *TagResponse) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *TagResponse) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *TagResponse) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetError

`func (o *TagResponse) GetError() ErrorMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TagResponse) GetErrorOk() (*ErrorMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TagResponse) SetError(v ErrorMessage)`

SetError sets Error field to given value.

### HasError

`func (o *TagResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


