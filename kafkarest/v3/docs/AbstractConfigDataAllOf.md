# AbstractConfigDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | **bool** |  | 
**IsReadOnly** | **bool** |  | 
**IsSensitive** | **bool** |  | 
**Source** | **string** |  | 
**Synonyms** | [**[]ConfigSynonymData**](ConfigSynonymData.md) |  | 

## Methods

### NewAbstractConfigDataAllOf

`func NewAbstractConfigDataAllOf(clusterId string, name string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []ConfigSynonymData, ) *AbstractConfigDataAllOf`

NewAbstractConfigDataAllOf instantiates a new AbstractConfigDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractConfigDataAllOfWithDefaults

`func NewAbstractConfigDataAllOfWithDefaults() *AbstractConfigDataAllOf`

NewAbstractConfigDataAllOfWithDefaults instantiates a new AbstractConfigDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *AbstractConfigDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AbstractConfigDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AbstractConfigDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *AbstractConfigDataAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractConfigDataAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractConfigDataAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AbstractConfigDataAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AbstractConfigDataAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AbstractConfigDataAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AbstractConfigDataAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AbstractConfigDataAllOf) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AbstractConfigDataAllOf) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *AbstractConfigDataAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AbstractConfigDataAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AbstractConfigDataAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsReadOnly

`func (o *AbstractConfigDataAllOf) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *AbstractConfigDataAllOf) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *AbstractConfigDataAllOf) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsSensitive

`func (o *AbstractConfigDataAllOf) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *AbstractConfigDataAllOf) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *AbstractConfigDataAllOf) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetSource

`func (o *AbstractConfigDataAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AbstractConfigDataAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AbstractConfigDataAllOf) SetSource(v string)`

SetSource sets Source field to given value.


### GetSynonyms

`func (o *AbstractConfigDataAllOf) GetSynonyms() []ConfigSynonymData`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *AbstractConfigDataAllOf) GetSynonymsOk() (*[]ConfigSynonymData, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *AbstractConfigDataAllOf) SetSynonyms(v []ConfigSynonymData)`

SetSynonyms sets Synonyms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


