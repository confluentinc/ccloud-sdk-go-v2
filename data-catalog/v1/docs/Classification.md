# Classification

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

## Methods

### NewClassification

`func NewClassification() *Classification`

NewClassification instantiates a new Classification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationWithDefaults

`func NewClassificationWithDefaults() *Classification`

NewClassificationWithDefaults instantiates a new Classification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *Classification) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Classification) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Classification) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *Classification) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *Classification) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Classification) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Classification) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Classification) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntityGuid

`func (o *Classification) GetEntityGuid() string`

GetEntityGuid returns the EntityGuid field if non-nil, zero value otherwise.

### GetEntityGuidOk

`func (o *Classification) GetEntityGuidOk() (*string, bool)`

GetEntityGuidOk returns a tuple with the EntityGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityGuid

`func (o *Classification) SetEntityGuid(v string)`

SetEntityGuid sets EntityGuid field to given value.

### HasEntityGuid

`func (o *Classification) HasEntityGuid() bool`

HasEntityGuid returns a boolean if a field has been set.

### GetEntityStatus

`func (o *Classification) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *Classification) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *Classification) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *Classification) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetPropagate

`func (o *Classification) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *Classification) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *Classification) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *Classification) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetValidityPeriods

`func (o *Classification) GetValidityPeriods() []TimeBoundary`

GetValidityPeriods returns the ValidityPeriods field if non-nil, zero value otherwise.

### GetValidityPeriodsOk

`func (o *Classification) GetValidityPeriodsOk() (*[]TimeBoundary, bool)`

GetValidityPeriodsOk returns a tuple with the ValidityPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriods

`func (o *Classification) SetValidityPeriods(v []TimeBoundary)`

SetValidityPeriods sets ValidityPeriods field to given value.

### HasValidityPeriods

`func (o *Classification) HasValidityPeriods() bool`

HasValidityPeriods returns a boolean if a field has been set.

### GetRemovePropagationsOnEntityDelete

`func (o *Classification) GetRemovePropagationsOnEntityDelete() bool`

GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field if non-nil, zero value otherwise.

### GetRemovePropagationsOnEntityDeleteOk

`func (o *Classification) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool)`

GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePropagationsOnEntityDelete

`func (o *Classification) SetRemovePropagationsOnEntityDelete(v bool)`

SetRemovePropagationsOnEntityDelete sets RemovePropagationsOnEntityDelete field to given value.

### HasRemovePropagationsOnEntityDelete

`func (o *Classification) HasRemovePropagationsOnEntityDelete() bool`

HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


