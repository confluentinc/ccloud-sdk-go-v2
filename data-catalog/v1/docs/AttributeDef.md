# AttributeDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name | [optional] 
**TypeName** | Pointer to **string** | The type name | [optional] 
**IsOptional** | Pointer to **bool** | Whether is optional | [optional] 
**Cardinality** | Pointer to **string** | The cardinality | [optional] 
**ValuesMinCount** | Pointer to **int32** | The values min count | [optional] 
**ValuesMaxCount** | Pointer to **int32** | The values max count | [optional] 
**IsUnique** | Pointer to **bool** | Whether is unique | [optional] 
**IsIndexable** | Pointer to **bool** | Whether is indexable | [optional] 
**IncludeInNotification** | Pointer to **bool** | Whether to include in notifications | [optional] 
**DefaultValue** | Pointer to **string** | The default value | [optional] 
**Description** | Pointer to **string** | The description | [optional] 
**SearchWeight** | Pointer to **int32** | The search weight | [optional] 
**IndexType** | Pointer to **string** | The index type | [optional] 
**Constraints** | Pointer to [**[]ConstraintDef**](ConstraintDef.md) | The constraints | [optional] 
**Options** | Pointer to **map[string]string** | The options | [optional] 
**DisplayName** | Pointer to **string** | The display name | [optional] 

## Methods

### NewAttributeDef

`func NewAttributeDef() *AttributeDef`

NewAttributeDef instantiates a new AttributeDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeDefWithDefaults

`func NewAttributeDefWithDefaults() *AttributeDef`

NewAttributeDefWithDefaults instantiates a new AttributeDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttributeDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeName

`func (o *AttributeDef) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *AttributeDef) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *AttributeDef) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *AttributeDef) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetIsOptional

`func (o *AttributeDef) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *AttributeDef) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *AttributeDef) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *AttributeDef) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetCardinality

`func (o *AttributeDef) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *AttributeDef) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *AttributeDef) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *AttributeDef) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetValuesMinCount

`func (o *AttributeDef) GetValuesMinCount() int32`

GetValuesMinCount returns the ValuesMinCount field if non-nil, zero value otherwise.

### GetValuesMinCountOk

`func (o *AttributeDef) GetValuesMinCountOk() (*int32, bool)`

GetValuesMinCountOk returns a tuple with the ValuesMinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesMinCount

`func (o *AttributeDef) SetValuesMinCount(v int32)`

SetValuesMinCount sets ValuesMinCount field to given value.

### HasValuesMinCount

`func (o *AttributeDef) HasValuesMinCount() bool`

HasValuesMinCount returns a boolean if a field has been set.

### GetValuesMaxCount

`func (o *AttributeDef) GetValuesMaxCount() int32`

GetValuesMaxCount returns the ValuesMaxCount field if non-nil, zero value otherwise.

### GetValuesMaxCountOk

`func (o *AttributeDef) GetValuesMaxCountOk() (*int32, bool)`

GetValuesMaxCountOk returns a tuple with the ValuesMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesMaxCount

`func (o *AttributeDef) SetValuesMaxCount(v int32)`

SetValuesMaxCount sets ValuesMaxCount field to given value.

### HasValuesMaxCount

`func (o *AttributeDef) HasValuesMaxCount() bool`

HasValuesMaxCount returns a boolean if a field has been set.

### GetIsUnique

`func (o *AttributeDef) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *AttributeDef) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *AttributeDef) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *AttributeDef) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsIndexable

`func (o *AttributeDef) GetIsIndexable() bool`

GetIsIndexable returns the IsIndexable field if non-nil, zero value otherwise.

### GetIsIndexableOk

`func (o *AttributeDef) GetIsIndexableOk() (*bool, bool)`

GetIsIndexableOk returns a tuple with the IsIndexable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndexable

`func (o *AttributeDef) SetIsIndexable(v bool)`

SetIsIndexable sets IsIndexable field to given value.

### HasIsIndexable

`func (o *AttributeDef) HasIsIndexable() bool`

HasIsIndexable returns a boolean if a field has been set.

### GetIncludeInNotification

`func (o *AttributeDef) GetIncludeInNotification() bool`

GetIncludeInNotification returns the IncludeInNotification field if non-nil, zero value otherwise.

### GetIncludeInNotificationOk

`func (o *AttributeDef) GetIncludeInNotificationOk() (*bool, bool)`

GetIncludeInNotificationOk returns a tuple with the IncludeInNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInNotification

`func (o *AttributeDef) SetIncludeInNotification(v bool)`

SetIncludeInNotification sets IncludeInNotification field to given value.

### HasIncludeInNotification

`func (o *AttributeDef) HasIncludeInNotification() bool`

HasIncludeInNotification returns a boolean if a field has been set.

### GetDefaultValue

`func (o *AttributeDef) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AttributeDef) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AttributeDef) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AttributeDef) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *AttributeDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSearchWeight

`func (o *AttributeDef) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *AttributeDef) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *AttributeDef) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *AttributeDef) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetIndexType

`func (o *AttributeDef) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *AttributeDef) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *AttributeDef) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *AttributeDef) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetConstraints

`func (o *AttributeDef) GetConstraints() []ConstraintDef`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AttributeDef) GetConstraintsOk() (*[]ConstraintDef, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AttributeDef) SetConstraints(v []ConstraintDef)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AttributeDef) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetOptions

`func (o *AttributeDef) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AttributeDef) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AttributeDef) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AttributeDef) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisplayName

`func (o *AttributeDef) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AttributeDef) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AttributeDef) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AttributeDef) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


