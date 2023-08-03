# PartialUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The type name | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The attributes | [optional] 
**Guid** | Pointer to **string** | The internal guid | [optional] 
**Status** | Pointer to **string** | The status | [optional] 
**ClassificationNames** | Pointer to **[]string** | The classification (tag) names | [optional] 
**Classifications** | Pointer to [**[]ClassificationHeader**](ClassificationHeader.md) | The classifications (tags) | [optional] 
**IsIncomplete** | Pointer to **bool** | Whether is incomplete | [optional] 

## Methods

### NewPartialUpdateParams

`func NewPartialUpdateParams() *PartialUpdateParams`

NewPartialUpdateParams instantiates a new PartialUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialUpdateParamsWithDefaults

`func NewPartialUpdateParamsWithDefaults() *PartialUpdateParams`

NewPartialUpdateParamsWithDefaults instantiates a new PartialUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *PartialUpdateParams) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *PartialUpdateParams) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *PartialUpdateParams) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *PartialUpdateParams) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *PartialUpdateParams) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PartialUpdateParams) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PartialUpdateParams) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PartialUpdateParams) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGuid

`func (o *PartialUpdateParams) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PartialUpdateParams) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PartialUpdateParams) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PartialUpdateParams) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetStatus

`func (o *PartialUpdateParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartialUpdateParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartialUpdateParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartialUpdateParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClassificationNames

`func (o *PartialUpdateParams) GetClassificationNames() []string`

GetClassificationNames returns the ClassificationNames field if non-nil, zero value otherwise.

### GetClassificationNamesOk

`func (o *PartialUpdateParams) GetClassificationNamesOk() (*[]string, bool)`

GetClassificationNamesOk returns a tuple with the ClassificationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationNames

`func (o *PartialUpdateParams) SetClassificationNames(v []string)`

SetClassificationNames sets ClassificationNames field to given value.

### HasClassificationNames

`func (o *PartialUpdateParams) HasClassificationNames() bool`

HasClassificationNames returns a boolean if a field has been set.

### GetClassifications

`func (o *PartialUpdateParams) GetClassifications() []ClassificationHeader`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *PartialUpdateParams) GetClassificationsOk() (*[]ClassificationHeader, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *PartialUpdateParams) SetClassifications(v []ClassificationHeader)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *PartialUpdateParams) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetIsIncomplete

`func (o *PartialUpdateParams) GetIsIncomplete() bool`

GetIsIncomplete returns the IsIncomplete field if non-nil, zero value otherwise.

### GetIsIncompleteOk

`func (o *PartialUpdateParams) GetIsIncompleteOk() (*bool, bool)`

GetIsIncompleteOk returns a tuple with the IsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncomplete

`func (o *PartialUpdateParams) SetIsIncomplete(v bool)`

SetIsIncomplete sets IsIncomplete field to given value.

### HasIsIncomplete

`func (o *PartialUpdateParams) HasIsIncomplete() bool`

HasIsIncomplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


