# Tag

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

## Methods

### NewTag

`func NewTag() *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *Tag) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Tag) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Tag) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *Tag) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *Tag) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Tag) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Tag) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Tag) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityGuid

`func (o *Tag) GetEntityGuid() string`

GetEntityGuid returns the EntityGuid field if non-nil, zero value otherwise.

### GetEntityGuidOk

`func (o *Tag) GetEntityGuidOk() (*string, bool)`

GetEntityGuidOk returns a tuple with the EntityGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityGuid

`func (o *Tag) SetEntityGuid(v string)`

SetEntityGuid sets EntityGuid field to given value.

### HasEntityGuid

`func (o *Tag) HasEntityGuid() bool`

HasEntityGuid returns a boolean if a field has been set.

### GetEntityStatus

`func (o *Tag) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *Tag) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *Tag) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *Tag) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetPropagate

`func (o *Tag) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *Tag) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *Tag) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *Tag) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetValidityPeriods

`func (o *Tag) GetValidityPeriods() []TimeBoundary`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *Tag) GetValidityPeriodsOk() (*[]TimeBoundary, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *Tag) SetValidityPeriods(v []TimeBoundary)`

SetValidityPeriods sets ValidityPeriods field to given value.

### HasValidityPeriods

`func (o *Tag) HasValidityPeriods() bool`

HasValidityPeriods returns a boolean if a field has been set.

### GetRemovePropagationsOnEntityDelete

`func (o *Tag) GetRemovePropagationsOnEntityDelete() bool`

GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field if non-nil, zero value otherwise.

### GetRemovePropagationsOnEntityDeleteOk

`func (o *Tag) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool)`

GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePropagationsOnEntityDelete

`func (o *Tag) SetRemovePropagationsOnEntityDelete(v bool)`

SetRemovePropagationsOnEntityDelete sets RemovePropagationsOnEntityDelete field to given value.

### HasRemovePropagationsOnEntityDelete

`func (o *Tag) HasRemovePropagationsOnEntityDelete() bool`

HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.

### GetEntityType

`func (o *Tag) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Tag) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Tag) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Tag) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityName

`func (o *Tag) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *Tag) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *Tag) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *Tag) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


