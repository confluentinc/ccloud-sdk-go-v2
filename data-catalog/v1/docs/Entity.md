# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The type name | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The type attributes | [optional] 
**Guid** | Pointer to **string** | The internal guid | [optional] 
**HomeId** | Pointer to **string** | The home id | [optional] 
**IsProxy** | Pointer to **bool** | Whether is a proxy | [optional] 
**IsIncomplete** | Pointer to **bool** | Whether is incomplete | [optional] 
**ProvenanceType** | Pointer to **int32** | The provenance type | [optional] 
**Status** | Pointer to **string** | The status | [optional] 
**CreatedBy** | Pointer to **string** | The creator | [optional] 
**UpdatedBy** | Pointer to **string** | The updater | [optional] 
**CreateTime** | Pointer to **int64** | The create time | [optional] 
**UpdateTime** | Pointer to **int64** | The update time | [optional] 
**Version** | Pointer to **int32** | The version | [optional] 
**RelationshipAttributes** | Pointer to **map[string]interface{}** | The relationship attributes | [optional] 
**Classifications** | Pointer to [**[]Classification**](Classification.md) | The classifications (tags) | [optional] 
**Meanings** | Pointer to [**[]TermAssignmentHeader**](TermAssignmentHeader.md) | The meanings | [optional] 
**CustomAttributes** | Pointer to **map[string]string** | The custom attributes | [optional] 
**BusinessAttributes** | Pointer to **map[string]map[string]interface{}** | The business attributes | [optional] 
**Labels** | Pointer to **[]string** | The labels | [optional] 
**Proxy** | Pointer to **bool** | Whether is a proxy | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *Entity) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Entity) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Entity) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *Entity) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *Entity) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Entity) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Entity) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Entity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGuid

`func (o *Entity) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Entity) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Entity) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Entity) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHomeId

`func (o *Entity) GetHomeId() string`

GetHomeId returns the HomeId field if non-nil, zero value otherwise.

### GetHomeIdOk

`func (o *Entity) GetHomeIdOk() (*string, bool)`

GetHomeIdOk returns a tuple with the HomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeId

`func (o *Entity) SetHomeId(v string)`

SetHomeId sets HomeId field to given value.

### HasHomeId

`func (o *Entity) HasHomeId() bool`

HasHomeId returns a boolean if a field has been set.

### GetIsProxy

`func (o *Entity) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *Entity) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *Entity) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *Entity) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetIsIncomplete

`func (o *Entity) GetIsIncomplete() bool`

GetIsIncomplete returns the IsIncomplete field if non-nil, zero value otherwise.

### GetIsIncompleteOk

`func (o *Entity) GetIsIncompleteOk() (*bool, bool)`

GetIsIncompleteOk returns a tuple with the IsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncomplete

`func (o *Entity) SetIsIncomplete(v bool)`

SetIsIncomplete sets IsIncomplete field to given value.

### HasIsIncomplete

`func (o *Entity) HasIsIncomplete() bool`

HasIsIncomplete returns a boolean if a field has been set.

### GetProvenanceType

`func (o *Entity) GetProvenanceType() int32`

GetProvenanceType returns the ProvenanceType field if non-nil, zero value otherwise.

### GetProvenanceTypeOk

`func (o *Entity) GetProvenanceTypeOk() (*int32, bool)`

GetProvenanceTypeOk returns a tuple with the ProvenanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceType

`func (o *Entity) SetProvenanceType(v int32)`

SetProvenanceType sets ProvenanceType field to given value.

### HasProvenanceType

`func (o *Entity) HasProvenanceType() bool`

HasProvenanceType returns a boolean if a field has been set.

### GetStatus

`func (o *Entity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Entity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Entity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Entity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Entity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Entity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Entity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Entity) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Entity) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Entity) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Entity) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Entity) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *Entity) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Entity) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Entity) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Entity) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Entity) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Entity) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Entity) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Entity) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *Entity) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Entity) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Entity) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Entity) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelationshipAttributes

`func (o *Entity) GetRelationshipAttributes() map[string]interface{}`

GetRelationshipAttributes returns the RelationshipAttributes field if non-nil, zero value otherwise.

### GetRelationshipAttributesOk

`func (o *Entity) GetRelationshipAttributesOk() (*map[string]interface{}, bool)`

GetRelationshipAttributesOk returns a tuple with the RelationshipAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipAttributes

`func (o *Entity) SetRelationshipAttributes(v map[string]interface{})`

SetRelationshipAttributes sets RelationshipAttributes field to given value.

### HasRelationshipAttributes

`func (o *Entity) HasRelationshipAttributes() bool`

HasRelationshipAttributes returns a boolean if a field has been set.

### GetClassifications

`func (o *Entity) GetClassifications() []Classification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *Entity) GetClassificationsOk() (*[]Classification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *Entity) SetClassifications(v []Classification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *Entity) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetMeanings

`func (o *Entity) GetMeanings() []TermAssignmentHeader`

GetMeanings returns the Meanings field if non-nil, zero value otherwise.

### GetMeaningsOk

`func (o *Entity) GetMeaningsOk() (*[]TermAssignmentHeader, bool)`

GetMeaningsOk returns a tuple with the Meanings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanings

`func (o *Entity) SetMeanings(v []TermAssignmentHeader)`

SetMeanings sets Meanings field to given value.

### HasMeanings

`func (o *Entity) HasMeanings() bool`

HasMeanings returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *Entity) GetCustomAttributes() map[string]string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *Entity) GetCustomAttributesOk() (*map[string]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *Entity) SetCustomAttributes(v map[string]string)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *Entity) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetBusinessAttributes

`func (o *Entity) GetBusinessAttributes() map[string]map[string]interface{}`

GetBusinessAttributes returns the BusinessAttributes field if non-nil, zero value otherwise.

### GetBusinessAttributesOk

`func (o *Entity) GetBusinessAttributesOk() (*map[string]map[string]interface{}, bool)`

GetBusinessAttributesOk returns a tuple with the BusinessAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAttributes

`func (o *Entity) SetBusinessAttributes(v map[string]map[string]interface{})`

SetBusinessAttributes sets BusinessAttributes field to given value.

### HasBusinessAttributes

`func (o *Entity) HasBusinessAttributes() bool`

HasBusinessAttributes returns a boolean if a field has been set.

### GetLabels

`func (o *Entity) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Entity) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Entity) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Entity) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProxy

`func (o *Entity) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Entity) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Entity) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Entity) HasProxy() bool`

HasProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


