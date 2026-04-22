# SqlV1AgentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the agent. | [optional] 
**Model** | Pointer to **string** | The name of the model the agent uses for inferencing. | [optional] 
**Prompt** | Pointer to **string** | The instruction prompt that guides the agent&#39;s behavior. | [optional] 
**Tools** | Pointer to **[]string** | The list of tools available to the agent. | [optional] 
**Properties** | Pointer to **map[string]string** | A set of key-value option pairs that configure the agent&#39;s behavior. | [optional] 

## Methods

### NewSqlV1AgentSpec

`func NewSqlV1AgentSpec() *SqlV1AgentSpec`

NewSqlV1AgentSpec instantiates a new SqlV1AgentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1AgentSpecWithDefaults

`func NewSqlV1AgentSpecWithDefaults() *SqlV1AgentSpec`

NewSqlV1AgentSpecWithDefaults instantiates a new SqlV1AgentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SqlV1AgentSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SqlV1AgentSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SqlV1AgentSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SqlV1AgentSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *SqlV1AgentSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SqlV1AgentSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SqlV1AgentSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SqlV1AgentSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *SqlV1AgentSpec) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *SqlV1AgentSpec) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *SqlV1AgentSpec) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *SqlV1AgentSpec) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTools

`func (o *SqlV1AgentSpec) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *SqlV1AgentSpec) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *SqlV1AgentSpec) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *SqlV1AgentSpec) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetProperties

`func (o *SqlV1AgentSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SqlV1AgentSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SqlV1AgentSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SqlV1AgentSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


