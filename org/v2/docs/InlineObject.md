# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireDurationMins** | Pointer to **int32** | The duration in minutes after which the token expires. Defaults to 6 months (259200 minutes) if not specified. Minimum: 1 month (43200 minutes). Maximum: 2 years (1051200 minutes).  | [optional] [default to 259200]

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

### GetExpireDurationMins

`func (o *InlineObject) GetExpireDurationMins() int32`

GetExpireDurationMins returns the ExpireDurationMins field if non-nil, zero value otherwise.

### GetExpireDurationMinsOk

`func (o *InlineObject) GetExpireDurationMinsOk() (*int32, bool)`

GetExpireDurationMinsOk returns a tuple with the ExpireDurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDurationMins

`func (o *InlineObject) SetExpireDurationMins(v int32)`

SetExpireDurationMins sets ExpireDurationMins field to given value.

### HasExpireDurationMins

`func (o *InlineObject) HasExpireDurationMins() bool`

HasExpireDurationMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


