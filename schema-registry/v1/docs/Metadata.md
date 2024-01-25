# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string][]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Sensitive** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Metadata) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Metadata) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Metadata) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Metadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProperties

`func (o *Metadata) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Metadata) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Metadata) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Metadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSensitive

`func (o *Metadata) GetSensitive() []string`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *Metadata) GetSensitiveOk() (*[]string, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *Metadata) SetSensitive(v []string)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *Metadata) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


