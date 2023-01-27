# TermAssignmentHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermGuid** | Pointer to **string** | The term guid | [optional] 
**RelationGuid** | Pointer to **string** | The relation guid | [optional] 
**Description** | Pointer to **string** | The description | [optional] 
**DisplayText** | Pointer to **string** | The display text | [optional] 
**Expression** | Pointer to **string** | The expression | [optional] 
**CreatedBy** | Pointer to **string** | The creator | [optional] 
**Steward** | Pointer to **string** | The steward | [optional] 
**Source** | Pointer to **string** | The source | [optional] 
**Confidence** | Pointer to **int32** | The confidence | [optional] 
**Status** | Pointer to **string** | The status | [optional] 

## Methods

### NewTermAssignmentHeader

`func NewTermAssignmentHeader() *TermAssignmentHeader`

NewTermAssignmentHeader instantiates a new TermAssignmentHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermAssignmentHeaderWithDefaults

`func NewTermAssignmentHeaderWithDefaults() *TermAssignmentHeader`

NewTermAssignmentHeaderWithDefaults instantiates a new TermAssignmentHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermGuid

`func (o *TermAssignmentHeader) GetTermGuid() string`

GetTermGuid returns the TermGuid field if non-nil, zero value otherwise.

### GetTermGuidOk

`func (o *TermAssignmentHeader) GetTermGuidOk() (*string, bool)`

GetTermGuidOk returns a tuple with the TermGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermGuid

`func (o *TermAssignmentHeader) SetTermGuid(v string)`

SetTermGuid sets TermGuid field to given value.

### HasTermGuid

`func (o *TermAssignmentHeader) HasTermGuid() bool`

HasTermGuid returns a boolean if a field has been set.

### GetRelationGuid

`func (o *TermAssignmentHeader) GetRelationGuid() string`

GetRelationGuid returns the RelationGuid field if non-nil, zero value otherwise.

### GetRelationGuidOk

`func (o *TermAssignmentHeader) GetRelationGuidOk() (*string, bool)`

GetRelationGuidOk returns a tuple with the RelationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationGuid

`func (o *TermAssignmentHeader) SetRelationGuid(v string)`

SetRelationGuid sets RelationGuid field to given value.

### HasRelationGuid

`func (o *TermAssignmentHeader) HasRelationGuid() bool`

HasRelationGuid returns a boolean if a field has been set.

### GetDescription

`func (o *TermAssignmentHeader) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermAssignmentHeader) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TermAssignmentHeader) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TermAssignmentHeader) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayText

`func (o *TermAssignmentHeader) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *TermAssignmentHeader) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *TermAssignmentHeader) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *TermAssignmentHeader) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetExpression

`func (o *TermAssignmentHeader) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TermAssignmentHeader) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TermAssignmentHeader) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *TermAssignmentHeader) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TermAssignmentHeader) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TermAssignmentHeader) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TermAssignmentHeader) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TermAssignmentHeader) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSteward

`func (o *TermAssignmentHeader) GetSteward() string`

GetSteward returns the Steward field if non-nil, zero value otherwise.

### GetStewardOk

`func (o *TermAssignmentHeader) GetStewardOk() (*string, bool)`

GetStewardOk returns a tuple with the Steward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteward

`func (o *TermAssignmentHeader) SetSteward(v string)`

SetSteward sets Steward field to given value.

### HasSteward

`func (o *TermAssignmentHeader) HasSteward() bool`

HasSteward returns a boolean if a field has been set.

### GetSource

`func (o *TermAssignmentHeader) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TermAssignmentHeader) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TermAssignmentHeader) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TermAssignmentHeader) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetConfidence

`func (o *TermAssignmentHeader) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TermAssignmentHeader) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TermAssignmentHeader) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TermAssignmentHeader) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetStatus

`func (o *TermAssignmentHeader) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TermAssignmentHeader) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TermAssignmentHeader) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TermAssignmentHeader) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


