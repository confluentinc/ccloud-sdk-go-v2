# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector to create. | [optional] 
**Config** | Pointer to **map[string]string** | Configuration parameters for the connector. All values should be strings. | [optional] 
**Offsets** | Pointer to **[]map[string]interface{}** | Array of offsets which are categorised into partitions. | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *InlineObject) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineObject) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineObject) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineObject) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOffsets

`func (o *InlineObject) GetOffsets() []map[string]interface{}`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *InlineObject) GetOffsetsOk() (*[]map[string]interface{}, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *InlineObject) SetOffsets(v []map[string]interface{})`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *InlineObject) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


