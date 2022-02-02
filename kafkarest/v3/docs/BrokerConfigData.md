# BrokerConfigData

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
**BrokerId** | **int32** |  | 

## Methods

### NewBrokerConfigData

`func NewBrokerConfigData(kind string, metadata ResourceMetadata, clusterId string, name string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []ConfigSynonymData, brokerId int32, ) *BrokerConfigData`

NewBrokerConfigData instantiates a new BrokerConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerConfigDataWithDefaults

`func NewBrokerConfigDataWithDefaults() *BrokerConfigData`

NewBrokerConfigDataWithDefaults instantiates a new BrokerConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerConfigData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerConfigData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerConfigData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerConfigData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerConfigData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerConfigData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BrokerConfigData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerConfigData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerConfigData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *BrokerConfigData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrokerConfigData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrokerConfigData) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *BrokerConfigData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BrokerConfigData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BrokerConfigData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BrokerConfigData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BrokerConfigData) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BrokerConfigData) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *BrokerConfigData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BrokerConfigData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BrokerConfigData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsReadOnly

`func (o *BrokerConfigData) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *BrokerConfigData) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *BrokerConfigData) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsSensitive

`func (o *BrokerConfigData) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *BrokerConfigData) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *BrokerConfigData) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetSource

`func (o *BrokerConfigData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BrokerConfigData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BrokerConfigData) SetSource(v string)`

SetSource sets Source field to given value.


### GetSynonyms

`func (o *BrokerConfigData) GetSynonyms() []ConfigSynonymData`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *BrokerConfigData) GetSynonymsOk() (*[]ConfigSynonymData, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *BrokerConfigData) SetSynonyms(v []ConfigSynonymData)`

SetSynonyms sets Synonyms field to given value.


### GetBrokerId

`func (o *BrokerConfigData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerConfigData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerConfigData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


