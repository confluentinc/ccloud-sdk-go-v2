# ClusterConfigData

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
**ConfigType** | **string** |  | 

## Methods

### NewClusterConfigData

`func NewClusterConfigData(kind string, metadata ResourceMetadata, clusterId string, name string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []ConfigSynonymData, configType string, ) *ClusterConfigData`

NewClusterConfigData instantiates a new ClusterConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigDataWithDefaults

`func NewClusterConfigDataWithDefaults() *ClusterConfigData`

NewClusterConfigDataWithDefaults instantiates a new ClusterConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterConfigData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterConfigData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterConfigData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterConfigData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterConfigData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterConfigData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ClusterConfigData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterConfigData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterConfigData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *ClusterConfigData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterConfigData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterConfigData) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ClusterConfigData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterConfigData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterConfigData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterConfigData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClusterConfigData) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterConfigData) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *ClusterConfigData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ClusterConfigData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ClusterConfigData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsReadOnly

`func (o *ClusterConfigData) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *ClusterConfigData) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *ClusterConfigData) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsSensitive

`func (o *ClusterConfigData) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *ClusterConfigData) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *ClusterConfigData) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetSource

`func (o *ClusterConfigData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterConfigData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterConfigData) SetSource(v string)`

SetSource sets Source field to given value.


### GetSynonyms

`func (o *ClusterConfigData) GetSynonyms() []ConfigSynonymData`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *ClusterConfigData) GetSynonymsOk() (*[]ConfigSynonymData, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *ClusterConfigData) SetSynonyms(v []ConfigSynonymData)`

SetSynonyms sets Synonyms field to given value.


### GetConfigType

`func (o *ClusterConfigData) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *ClusterConfigData) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *ClusterConfigData) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


