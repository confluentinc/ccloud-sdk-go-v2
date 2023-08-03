# EntityHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | Pointer to **string** | The type name | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | The attributes | [optional] 
**Guid** | Pointer to **string** | The internal guid | [optional] 
**Status** | Pointer to **string** | The status | [optional] 
**DisplayText** | Pointer to **string** | The display text | [optional] 
**ClassificationNames** | Pointer to **[]string** | The classification (tag) names | [optional] 
**Classifications** | Pointer to [**[]Classification**](Classification.md) | The classifications (tags) | [optional] 
**MeaningNames** | Pointer to **[]string** | The meaning names | [optional] 
**Meanings** | Pointer to [**[]TermAssignmentHeader**](TermAssignmentHeader.md) | The meanings | [optional] 
**IsIncomplete** | Pointer to **bool** | Whether is incomplete | [optional] 
**Labels** | Pointer to **[]string** | The labels | [optional] 

## Methods

### NewEntityHeader

`func NewEntityHeader() *EntityHeader`

NewEntityHeader instantiates a new EntityHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityHeaderWithDefaults

`func NewEntityHeaderWithDefaults() *EntityHeader`

NewEntityHeaderWithDefaults instantiates a new EntityHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *EntityHeader) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EntityHeader) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EntityHeader) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *EntityHeader) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetAttributes

`func (o *EntityHeader) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntityHeader) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntityHeader) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntityHeader) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGuid

`func (o *EntityHeader) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *EntityHeader) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *EntityHeader) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *EntityHeader) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetStatus

`func (o *EntityHeader) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityHeader) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityHeader) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntityHeader) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisplayText

`func (o *EntityHeader) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *EntityHeader) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *EntityHeader) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *EntityHeader) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetClassificationNames

`func (o *EntityHeader) GetClassificationNames() []string`

GetClassificationNames returns the ClassificationNames field if non-nil, zero value otherwise.

### GetClassificationNamesOk

`func (o *EntityHeader) GetClassificationNamesOk() (*[]string, bool)`

GetClassificationNamesOk returns a tuple with the ClassificationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationNames

`func (o *EntityHeader) SetClassificationNames(v []string)`

SetClassificationNames sets ClassificationNames field to given value.

### HasClassificationNames

`func (o *EntityHeader) HasClassificationNames() bool`

HasClassificationNames returns a boolean if a field has been set.

### GetClassifications

`func (o *EntityHeader) GetClassifications() []Classification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *EntityHeader) GetClassificationsOk() (*[]Classification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *EntityHeader) SetClassifications(v []Classification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *EntityHeader) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetMeaningNames

`func (o *EntityHeader) GetMeaningNames() []string`

GetMeaningNames returns the MeaningNames field if non-nil, zero value otherwise.

### GetMeaningNamesOk

`func (o *EntityHeader) GetMeaningNamesOk() (*[]string, bool)`

GetMeaningNamesOk returns a tuple with the MeaningNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeaningNames

`func (o *EntityHeader) SetMeaningNames(v []string)`

SetMeaningNames sets MeaningNames field to given value.

### HasMeaningNames

`func (o *EntityHeader) HasMeaningNames() bool`

HasMeaningNames returns a boolean if a field has been set.

### GetMeanings

`func (o *EntityHeader) GetMeanings() []TermAssignmentHeader`

GetMeanings returns the Meanings field if non-nil, zero value otherwise.

### GetMeaningsOk

`func (o *EntityHeader) GetMeaningsOk() (*[]TermAssignmentHeader, bool)`

GetMeaningsOk returns a tuple with the Meanings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanings

`func (o *EntityHeader) SetMeanings(v []TermAssignmentHeader)`

SetMeanings sets Meanings field to given value.

### HasMeanings

`func (o *EntityHeader) HasMeanings() bool`

HasMeanings returns a boolean if a field has been set.

### GetIsIncomplete

`func (o *EntityHeader) GetIsIncomplete() bool`

GetIsIncomplete returns the IsIncomplete field if non-nil, zero value otherwise.

### GetIsIncompleteOk

`func (o *EntityHeader) GetIsIncompleteOk() (*bool, bool)`

GetIsIncompleteOk returns a tuple with the IsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncomplete

`func (o *EntityHeader) SetIsIncomplete(v bool)`

SetIsIncomplete sets IsIncomplete field to given value.

### HasIsIncomplete

`func (o *EntityHeader) HasIsIncomplete() bool`

HasIsIncomplete returns a boolean if a field has been set.

### GetLabels

`func (o *EntityHeader) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EntityHeader) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EntityHeader) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EntityHeader) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


