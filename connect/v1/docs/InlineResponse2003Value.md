# InlineResponse2003Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the configuration | [optional] 
**Value** | Pointer to **string** | The value for the configuration | [optional] 
**RecommendedValues** | Pointer to **[]string** | The list of valid values for the configuration | [optional] 
**Errors** | Pointer to **[]string** | Errors, if any, in the configuration value | [optional] 
**Visible** | Pointer to **bool** | The visibility of the configuration. Based on the values of other configuration fields, this visibility boolean value points out if the current field should be visible or not. | [optional] 

## Methods

### NewInlineResponse2003Value

`func NewInlineResponse2003Value() *InlineResponse2003Value`

NewInlineResponse2003Value instantiates a new InlineResponse2003Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003ValueWithDefaults

`func NewInlineResponse2003ValueWithDefaults() *InlineResponse2003Value`

NewInlineResponse2003ValueWithDefaults instantiates a new InlineResponse2003Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2003Value) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003Value) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003Value) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2003Value) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *InlineResponse2003Value) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse2003Value) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse2003Value) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineResponse2003Value) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRecommendedValues

`func (o *InlineResponse2003Value) GetRecommendedValues() []string`

GetRecommendedValues returns the RecommendedValues field if non-nil, zero value otherwise.

### GetRecommendedValuesOk

`func (o *InlineResponse2003Value) GetRecommendedValuesOk() (*[]string, bool)`

GetRecommendedValuesOk returns a tuple with the RecommendedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedValues

`func (o *InlineResponse2003Value) SetRecommendedValues(v []string)`

SetRecommendedValues sets RecommendedValues field to given value.

### HasRecommendedValues

`func (o *InlineResponse2003Value) HasRecommendedValues() bool`

HasRecommendedValues returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse2003Value) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse2003Value) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse2003Value) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse2003Value) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetVisible

`func (o *InlineResponse2003Value) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineResponse2003Value) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineResponse2003Value) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineResponse2003Value) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


