# BusinessMetadataDef

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

## Methods

### NewBusinessMetadataDef

`func NewBusinessMetadataDef() *BusinessMetadataDef`

NewBusinessMetadataDef instantiates a new BusinessMetadataDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessMetadataDefWithDefaults

`func NewBusinessMetadataDefWithDefaults() *BusinessMetadataDef`

NewBusinessMetadataDefWithDefaults instantiates a new BusinessMetadataDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BusinessMetadataDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BusinessMetadataDef) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BusinessMetadataDef) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BusinessMetadataDef) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *BusinessMetadataDef) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BusinessMetadataDef) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BusinessMetadataDef) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *BusinessMetadataDef) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BusinessMetadataDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BusinessMetadataDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BusinessMetadataDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BusinessMetadataDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *BusinessMetadataDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *BusinessMetadataDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *BusinessMetadataDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *BusinessMetadataDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *BusinessMetadataDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *BusinessMetadataDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *BusinessMetadataDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *BusinessMetadataDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *BusinessMetadataDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *BusinessMetadataDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *BusinessMetadataDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *BusinessMetadataDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *BusinessMetadataDef) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BusinessMetadataDef) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BusinessMetadataDef) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BusinessMetadataDef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *BusinessMetadataDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessMetadataDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessMetadataDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessMetadataDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BusinessMetadataDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BusinessMetadataDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BusinessMetadataDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BusinessMetadataDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTypeVersion

`func (o *BusinessMetadataDef) GetTypeVersion() string`

GetTypeVersion returns the TypeVersion field if non-nil, zero value otherwise.

### GetTypeVersionOk

`func (o *BusinessMetadataDef) GetTypeVersionOk() (*string, bool)`

GetTypeVersionOk returns a tuple with the TypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVersion

`func (o *BusinessMetadataDef) SetTypeVersion(v string)`

SetTypeVersion sets TypeVersion field to given value.

### HasTypeVersion

`func (o *BusinessMetadataDef) HasTypeVersion() bool`

HasTypeVersion returns a boolean if a field has been set.

### GetServiceType

`func (o *BusinessMetadataDef) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *BusinessMetadataDef) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *BusinessMetadataDef) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *BusinessMetadataDef) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOptions

`func (o *BusinessMetadataDef) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BusinessMetadataDef) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BusinessMetadataDef) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BusinessMetadataDef) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAttributeDefs

`func (o *BusinessMetadataDef) GetAttributeDefs() []AttributeDef`

GetAttributeDefs returns the AttributeDefs field if non-nil, zero value otherwise.

### GetAttributeDefsOk

`func (o *BusinessMetadataDef) GetAttributeDefsOk() (*[]AttributeDef, bool)`

GetAttributeDefsOk returns a tuple with the AttributeDefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDefs

`func (o *BusinessMetadataDef) SetAttributeDefs(v []AttributeDef)`

SetAttributeDefs sets AttributeDefs field to given value.

### HasAttributeDefs

`func (o *BusinessMetadataDef) HasAttributeDefs() bool`

HasAttributeDefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


