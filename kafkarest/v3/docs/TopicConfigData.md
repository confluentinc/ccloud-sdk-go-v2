# TopicConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Name** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | **bool** |  | 
**IsReadOnly** | **bool** |  | 
**IsSensitive** | **bool** |  | 
**Source** | **string** |  | 
**Synonyms** | [**[]ConfigSynonymData**](ConfigSynonymData.md) |  | 
**TopicName** | **string** |  | 

## Methods

### NewTopicConfigData

`func NewTopicConfigData(kind string, metadata ResourceMetadata, clusterId string, name string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []ConfigSynonymData, topicName string, ) *TopicConfigData`

NewTopicConfigData instantiates a new TopicConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicConfigDataWithDefaults

`func NewTopicConfigDataWithDefaults() *TopicConfigData`

NewTopicConfigDataWithDefaults instantiates a new TopicConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TopicConfigData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TopicConfigData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TopicConfigData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TopicConfigData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TopicConfigData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TopicConfigData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *TopicConfigData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TopicConfigData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TopicConfigData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *TopicConfigData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopicConfigData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopicConfigData) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *TopicConfigData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TopicConfigData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TopicConfigData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TopicConfigData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TopicConfigData) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TopicConfigData) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *TopicConfigData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TopicConfigData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TopicConfigData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsReadOnly

`func (o *TopicConfigData) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *TopicConfigData) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *TopicConfigData) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsSensitive

`func (o *TopicConfigData) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *TopicConfigData) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *TopicConfigData) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetSource

`func (o *TopicConfigData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopicConfigData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopicConfigData) SetSource(v string)`

SetSource sets Source field to given value.


### GetSynonyms

`func (o *TopicConfigData) GetSynonyms() []ConfigSynonymData`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *TopicConfigData) GetSynonymsOk() (*[]ConfigSynonymData, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *TopicConfigData) SetSynonyms(v []ConfigSynonymData)`

SetSynonyms sets Synonyms field to given value.


### GetTopicName

`func (o *TopicConfigData) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *TopicConfigData) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *TopicConfigData) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


