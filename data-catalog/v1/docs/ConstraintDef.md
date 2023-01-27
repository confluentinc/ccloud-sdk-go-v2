# ConstraintDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** | The params | [optional] 

## Methods

### NewConstraintDef

`func NewConstraintDef() *ConstraintDef`

NewConstraintDef instantiates a new ConstraintDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintDefWithDefaults

`func NewConstraintDefWithDefaults() *ConstraintDef`

NewConstraintDefWithDefaults instantiates a new ConstraintDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConstraintDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstraintDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstraintDef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstraintDef) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParams

`func (o *ConstraintDef) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ConstraintDef) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ConstraintDef) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ConstraintDef) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


