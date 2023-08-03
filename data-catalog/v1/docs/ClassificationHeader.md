# ClassificationHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The tag name | [optional] 
**EntityGuid** | Pointer to **string** | The internal entity guid | [optional] 
**EntityStatus** | Pointer to **string** | The entity status | [optional] 
**Propagate** | Pointer to **bool** | Whether to propagate the tag | [optional] 
**RemovePropagationsOnEntityDelete** | Pointer to **bool** | Whether to remove propagations on entity delete | [optional] 

## Methods

### NewClassificationHeader

`func NewClassificationHeader() *ClassificationHeader`

NewClassificationHeader instantiates a new ClassificationHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationHeaderWithDefaults

`func NewClassificationHeaderWithDefaults() *ClassificationHeader`

NewClassificationHeaderWithDefaults instantiates a new ClassificationHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *ClassificationHeader) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ClassificationHeader) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ClassificationHeader) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ClassificationHeader) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetEntityGuid

`func (o *ClassificationHeader) GetEntityGuid() string`

GetEntityGuid returns the EntityGuid field if non-nil, zero value otherwise.

### GetEntityGuidOk

`func (o *ClassificationHeader) GetEntityGuidOk() (*string, bool)`

GetEntityGuidOk returns a tuple with the EntityGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityGuid

`func (o *ClassificationHeader) SetEntityGuid(v string)`

SetEntityGuid sets EntityGuid field to given value.

### HasEntityGuid

`func (o *ClassificationHeader) HasEntityGuid() bool`

HasEntityGuid returns a boolean if a field has been set.

### GetEntityStatus

`func (o *ClassificationHeader) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *ClassificationHeader) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *ClassificationHeader) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *ClassificationHeader) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetPropagate

`func (o *ClassificationHeader) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *ClassificationHeader) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *ClassificationHeader) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *ClassificationHeader) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetRemovePropagationsOnEntityDelete

`func (o *ClassificationHeader) GetRemovePropagationsOnEntityDelete() bool`

GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field if non-nil, zero value otherwise.

### GetRemovePropagationsOnEntityDeleteOk

`func (o *ClassificationHeader) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool)`

GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePropagationsOnEntityDelete

`func (o *ClassificationHeader) SetRemovePropagationsOnEntityDelete(v bool)`

SetRemovePropagationsOnEntityDelete sets RemovePropagationsOnEntityDelete field to given value.

### HasRemovePropagationsOnEntityDelete

`func (o *ClassificationHeader) HasRemovePropagationsOnEntityDelete() bool`

HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


