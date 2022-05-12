# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The class name of the connector plugin. | [optional] 
**Groups** | Pointer to **[]string** | The list of groups used in configuration definitions. | [optional] 
**ErrorCount** | Pointer to **int32** | The total number of errors encountered during configuration validation. | [optional] 
**Configs** | Pointer to [**[]InlineResponse2003Configs**](InlineResponse2003Configs.md) |  | [optional] 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003() *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2003) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2003) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroups

`func (o *InlineResponse2003) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *InlineResponse2003) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *InlineResponse2003) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *InlineResponse2003) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetErrorCount

`func (o *InlineResponse2003) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *InlineResponse2003) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *InlineResponse2003) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *InlineResponse2003) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetConfigs

`func (o *InlineResponse2003) GetConfigs() []InlineResponse2003Configs`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *InlineResponse2003) GetConfigsOk() (*[]InlineResponse2003Configs, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *InlineResponse2003) SetConfigs(v []InlineResponse2003Configs)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *InlineResponse2003) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


