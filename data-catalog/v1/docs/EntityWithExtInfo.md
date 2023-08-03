# EntityWithExtInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferredEntities** | Pointer to [**map[string]Entity**](Entity.md) | The referred entities | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 

## Methods

### NewEntityWithExtInfo

`func NewEntityWithExtInfo() *EntityWithExtInfo`

NewEntityWithExtInfo instantiates a new EntityWithExtInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithExtInfoWithDefaults

`func NewEntityWithExtInfoWithDefaults() *EntityWithExtInfo`

NewEntityWithExtInfoWithDefaults instantiates a new EntityWithExtInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferredEntities

`func (o *EntityWithExtInfo) GetReferredEntities() map[string]Entity`

GetReferredEntities returns the ReferredEntities field if non-nil, zero value otherwise.

### GetReferredEntitiesOk

`func (o *EntityWithExtInfo) GetReferredEntitiesOk() (*map[string]Entity, bool)`

GetReferredEntitiesOk returns a tuple with the ReferredEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredEntities

`func (o *EntityWithExtInfo) SetReferredEntities(v map[string]Entity)`

SetReferredEntities sets ReferredEntities field to given value.

### HasReferredEntities

`func (o *EntityWithExtInfo) HasReferredEntities() bool`

HasReferredEntities returns a boolean if a field has been set.

### GetEntity

`func (o *EntityWithExtInfo) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *EntityWithExtInfo) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *EntityWithExtInfo) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *EntityWithExtInfo) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


