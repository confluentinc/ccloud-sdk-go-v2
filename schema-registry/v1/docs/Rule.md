# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rule name | [optional] 
**Doc** | Pointer to **string** | Rule doc | [optional] 
**Kind** | Pointer to **string** | Rule kind | [optional] 
**Mode** | Pointer to **string** | Rule mode | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 
**Tags** | Pointer to **[]string** | The tags to which this rule applies | [optional] 
**Params** | Pointer to **map[string]string** | Optional params for the rule | [optional] 
**Expr** | Pointer to **string** | Rule expression | [optional] 
**OnSuccess** | Pointer to **string** | Rule action on success | [optional] 
**OnFailure** | Pointer to **string** | Rule action on failure | [optional] 
**Disabled** | Pointer to **bool** | Whether the rule is disabled | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Rule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDoc

`func (o *Rule) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *Rule) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *Rule) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *Rule) HasDoc() bool`

HasDoc returns a boolean if a field has been set.

### GetKind

`func (o *Rule) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Rule) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Rule) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Rule) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMode

`func (o *Rule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Rule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Rule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Rule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *Rule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Rule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *Rule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Rule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Rule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Rule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParams

`func (o *Rule) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Rule) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Rule) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *Rule) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetExpr

`func (o *Rule) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *Rule) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *Rule) SetExpr(v string)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *Rule) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### GetOnSuccess

`func (o *Rule) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *Rule) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *Rule) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *Rule) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetOnFailure

`func (o *Rule) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *Rule) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *Rule) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *Rule) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetDisabled

`func (o *Rule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Rule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Rule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Rule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


