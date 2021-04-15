# InlineResponse2003Definition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the configuration | [optional] 
**Type** | Pointer to **string** | The config types | [optional] 
**Required** | Pointer to **bool** | Whether this configuration is required | [optional] 
**DefaultValue** | Pointer to **string** | Default value for this configuration | [optional] 
**Importance** | Pointer to **string** | The importance level for a configuration | [optional] 
**Documentation** | Pointer to **string** | The documentation for the configuration | [optional] 
**Group** | Pointer to **string** | The UI group to which the configuration belongs to | [optional] 
**Width** | Pointer to **string** | The width of a configuration value | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Dependents** | Pointer to **[]string** | Other configurations on which this configuration is dependent | [optional] 
**Order** | Pointer to **int32** | The order of configuration in specified group | [optional] 
**Alias** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse2003Definition

`func NewInlineResponse2003Definition() *InlineResponse2003Definition`

NewInlineResponse2003Definition instantiates a new InlineResponse2003Definition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003DefinitionWithDefaults

`func NewInlineResponse2003DefinitionWithDefaults() *InlineResponse2003Definition`

NewInlineResponse2003DefinitionWithDefaults instantiates a new InlineResponse2003Definition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2003Definition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003Definition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003Definition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2003Definition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse2003Definition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse2003Definition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse2003Definition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse2003Definition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *InlineResponse2003Definition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *InlineResponse2003Definition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *InlineResponse2003Definition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *InlineResponse2003Definition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *InlineResponse2003Definition) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *InlineResponse2003Definition) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *InlineResponse2003Definition) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *InlineResponse2003Definition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetImportance

`func (o *InlineResponse2003Definition) GetImportance() string`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *InlineResponse2003Definition) GetImportanceOk() (*string, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *InlineResponse2003Definition) SetImportance(v string)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *InlineResponse2003Definition) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### GetDocumentation

`func (o *InlineResponse2003Definition) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *InlineResponse2003Definition) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *InlineResponse2003Definition) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *InlineResponse2003Definition) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse2003Definition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse2003Definition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse2003Definition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse2003Definition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetWidth

`func (o *InlineResponse2003Definition) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InlineResponse2003Definition) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InlineResponse2003Definition) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InlineResponse2003Definition) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineResponse2003Definition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineResponse2003Definition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineResponse2003Definition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineResponse2003Definition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDependents

`func (o *InlineResponse2003Definition) GetDependents() []string`

GetDependents returns the Dependents field if non-nil, zero value otherwise.

### GetDependentsOk

`func (o *InlineResponse2003Definition) GetDependentsOk() (*[]string, bool)`

GetDependentsOk returns a tuple with the Dependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependents

`func (o *InlineResponse2003Definition) SetDependents(v []string)`

SetDependents sets Dependents field to given value.

### HasDependents

`func (o *InlineResponse2003Definition) HasDependents() bool`

HasDependents returns a boolean if a field has been set.

### GetOrder

`func (o *InlineResponse2003Definition) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InlineResponse2003Definition) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InlineResponse2003Definition) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InlineResponse2003Definition) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAlias

`func (o *InlineResponse2003Definition) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *InlineResponse2003Definition) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *InlineResponse2003Definition) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *InlineResponse2003Definition) HasAlias() bool`

HasAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


