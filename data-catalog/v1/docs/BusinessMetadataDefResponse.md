# BusinessMetadataDefResponse

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
**Error** | Pointer to [**ErrorMessage**](ErrorMessage.md) |  | [optional] 

## Methods

### NewBusinessMetadataDefResponse

`func NewBusinessMetadataDefResponse() *BusinessMetadataDefResponse`

NewBusinessMetadataDefResponse instantiates a new BusinessMetadataDefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessMetadataDefResponseWithDefaults

`func NewBusinessMetadataDefResponseWithDefaults() *BusinessMetadataDefResponse`

NewBusinessMetadataDefResponseWithDefaults instantiates a new BusinessMetadataDefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BusinessMetadataDefResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BusinessMetadataDefResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BusinessMetadataDefResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BusinessMetadataDefResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *BusinessMetadataDefResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BusinessMetadataDefResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BusinessMetadataDefResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *BusinessMetadataDefResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BusinessMetadataDefResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BusinessMetadataDefResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BusinessMetadataDefResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BusinessMetadataDefResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *BusinessMetadataDefResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *BusinessMetadataDefResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *BusinessMetadataDefResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *BusinessMetadataDefResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *BusinessMetadataDefResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *BusinessMetadataDefResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *BusinessMetadataDefResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *BusinessMetadataDefResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *BusinessMetadataDefResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *BusinessMetadataDefResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *BusinessMetadataDefResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *BusinessMetadataDefResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *BusinessMetadataDefResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BusinessMetadataDefResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BusinessMetadataDefResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BusinessMetadataDefResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *BusinessMetadataDefResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessMetadataDefResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessMetadataDefResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessMetadataDefResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BusinessMetadataDefResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BusinessMetadataDefResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BusinessMetadataDefResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BusinessMetadataDefResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTypeVersion

`func (o *BusinessMetadataDefResponse) GetTypeVersion() string`

GetTypeVersion returns the TypeVersion field if non-nil, zero value otherwise.

### GetTypeVersionOk

`func (o *BusinessMetadataDefResponse) GetTypeVersionOk() (*string, bool)`

GetTypeVersionOk returns a tuple with the TypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVersion

`func (o *BusinessMetadataDefResponse) SetTypeVersion(v string)`

SetTypeVersion sets TypeVersion field to given value.

### HasTypeVersion

`func (o *BusinessMetadataDefResponse) HasTypeVersion() bool`

HasTypeVersion returns a boolean if a field has been set.

### GetServiceType

`func (o *BusinessMetadataDefResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *BusinessMetadataDefResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *BusinessMetadataDefResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *BusinessMetadataDefResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOptions

`func (o *BusinessMetadataDefResponse) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BusinessMetadataDefResponse) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BusinessMetadataDefResponse) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BusinessMetadataDefResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAttributeDefs

`func (o *BusinessMetadataDefResponse) GetAttributeDefs() []AttributeDef`

GetAttributeDefs returns the AttributeDefs field if non-nil, zero value otherwise.

### GetAttributeDefsOk

`func (o *BusinessMetadataDefResponse) GetAttributeDefsOk() (*[]AttributeDef, bool)`

GetAttributeDefsOk returns a tuple with the AttributeDefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDefs

`func (o *BusinessMetadataDefResponse) SetAttributeDefs(v []AttributeDef)`

SetAttributeDefs sets AttributeDefs field to given value.

### HasAttributeDefs

`func (o *BusinessMetadataDefResponse) HasAttributeDefs() bool`

HasAttributeDefs returns a boolean if a field has been set.

### GetError

`func (o *BusinessMetadataDefResponse) GetError() ErrorMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BusinessMetadataDefResponse) GetErrorOk() (*ErrorMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BusinessMetadataDefResponse) SetError(v ErrorMessage)`

SetError sets Error field to given value.

### HasError

`func (o *BusinessMetadataDefResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


