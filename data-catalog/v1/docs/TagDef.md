# TagDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category | [optional] 
**Guid** | Pointer to **string** | The internal guid | [optional] 
**CreatedBy** | Pointer to **string** | The creator | [optional] 
**UpdatedBy** | Pointer to **string** | The updater | [optional] 
**CreateTime** | Pointer to **int64** | The create time | [optional] 
**UpdateTime** | Pointer to **int64** | The update time | [optional] 
**Version** | Pointer to **int32** | The version | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Description** | Pointer to **string** | The description | [optional] 
**TypeVersion** | Pointer to **string** | The type version | [optional] 
**ServiceType** | Pointer to **string** | The service type | [optional] 
**Options** | Pointer to **map[string]string** | The options | [optional] 
**AttributeDefs** | Pointer to [**[]AttributeDef**](AttributeDef.md) | The attribute definitions | [optional] 
**SuperTypes** | Pointer to **[]string** | The supertypes | [optional] 
**EntityTypes** | Pointer to **[]string** | The entity types | [optional] 
**SubTypes** | Pointer to **[]string** | The subtypes | [optional] 

## Methods

### NewTagDef

`func NewTagDef() *TagDef`

NewTagDef instantiates a new TagDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDefWithDefaults

`func NewTagDefWithDefaults() *TagDef`

NewTagDefWithDefaults instantiates a new TagDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TagDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TagDef) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TagDef) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TagDef) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *TagDef) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TagDef) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TagDef) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TagDef) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TagDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TagDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TagDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TagDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TagDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TagDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TagDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TagDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *TagDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TagDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TagDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TagDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *TagDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TagDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TagDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TagDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *TagDef) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TagDef) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TagDef) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TagDef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *TagDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTypeVersion

`func (o *TagDef) GetTypeVersion() string`

GetTypeVersion returns the TypeVersion field if non-nil, zero value otherwise.

### GetTypeVersionOk

`func (o *TagDef) GetTypeVersionOk() (*string, bool)`

GetTypeVersionOk returns a tuple with the TypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVersion

`func (o *TagDef) SetTypeVersion(v string)`

SetTypeVersion sets TypeVersion field to given value.

### HasTypeVersion

`func (o *TagDef) HasTypeVersion() bool`

HasTypeVersion returns a boolean if a field has been set.

### GetServiceType

`func (o *TagDef) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TagDef) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TagDef) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *TagDef) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOptions

`func (o *TagDef) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TagDef) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TagDef) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TagDef) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAttributeDefs

`func (o *TagDef) GetAttributeDefs() []AttributeDef`

GetAttributeDefs returns the AttributeDefs field if non-nil, zero value otherwise.

### GetAttributeDefsOk

`func (o *TagDef) GetAttributeDefsOk() (*[]AttributeDef, bool)`

GetAttributeDefsOk returns a tuple with the AttributeDefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDefs

`func (o *TagDef) SetAttributeDefs(v []AttributeDef)`

SetAttributeDefs sets AttributeDefs field to given value.

### HasAttributeDefs

`func (o *TagDef) HasAttributeDefs() bool`

HasAttributeDefs returns a boolean if a field has been set.

### GetSuperTypes

`func (o *TagDef) GetSuperTypes() []string`

GetSuperTypes returns the SuperTypes field if non-nil, zero value otherwise.

### GetSuperTypesOk

`func (o *TagDef) GetSuperTypesOk() (*[]string, bool)`

GetSuperTypesOk returns a tuple with the SuperTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperTypes

`func (o *TagDef) SetSuperTypes(v []string)`

SetSuperTypes sets SuperTypes field to given value.

### HasSuperTypes

`func (o *TagDef) HasSuperTypes() bool`

HasSuperTypes returns a boolean if a field has been set.

### GetEntityTypes

`func (o *TagDef) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *TagDef) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *TagDef) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *TagDef) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetSubTypes

`func (o *TagDef) GetSubTypes() []string`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *TagDef) GetSubTypesOk() (*[]string, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *TagDef) SetSubTypes(v []string)`

SetSubTypes sets SubTypes field to given value.

### HasSubTypes

`func (o *TagDef) HasSubTypes() bool`

HasSubTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


