# SqlV1ToolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to **string** | The name of the connection this tool uses. Required for MCP and A2A tools. Mutually exclusive with function. | [optional] 
**Function** | Pointer to **string** | The name of the function this tool wraps. Required for function-based tools. Mutually exclusive with connection. | [optional] 
**Comment** | Pointer to **string** | An optional comment describing the tool. | [optional] 
**Options** | Pointer to **map[string]string** | A set of key-value option pairs that configure the tool&#39;s behavior. Supported options vary by tool type: - MCP tools: type, allowed_tools, request_timeout, max_retries, headers - A2A tools: type, agent_card_path, request_timeout, max_retries - Function tools: type, description | [optional] 

## Methods

### NewSqlV1ToolSpec

`func NewSqlV1ToolSpec() *SqlV1ToolSpec`

NewSqlV1ToolSpec instantiates a new SqlV1ToolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ToolSpecWithDefaults

`func NewSqlV1ToolSpecWithDefaults() *SqlV1ToolSpec`

NewSqlV1ToolSpecWithDefaults instantiates a new SqlV1ToolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *SqlV1ToolSpec) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *SqlV1ToolSpec) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *SqlV1ToolSpec) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *SqlV1ToolSpec) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetFunction

`func (o *SqlV1ToolSpec) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *SqlV1ToolSpec) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *SqlV1ToolSpec) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *SqlV1ToolSpec) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetComment

`func (o *SqlV1ToolSpec) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SqlV1ToolSpec) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SqlV1ToolSpec) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SqlV1ToolSpec) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetOptions

`func (o *SqlV1ToolSpec) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SqlV1ToolSpec) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SqlV1ToolSpec) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SqlV1ToolSpec) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


