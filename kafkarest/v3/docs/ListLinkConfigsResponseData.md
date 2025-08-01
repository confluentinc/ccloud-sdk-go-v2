# ListLinkConfigsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsDefault** | **bool** |  | 
**IsReadOnly** | **bool** |  | 
**IsSensitive** | **bool** |  | 
**Source** | **string** |  | 
**Synonyms** | **[]string** |  | 
**LinkName** | **string** |  | 

## Methods

### NewListLinkConfigsResponseData

`func NewListLinkConfigsResponseData(kind string, metadata ResourceMetadata, clusterId string, name string, value string, isDefault bool, isReadOnly bool, isSensitive bool, source string, synonyms []string, linkName string, ) *ListLinkConfigsResponseData`

NewListLinkConfigsResponseData instantiates a new ListLinkConfigsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinkConfigsResponseDataWithDefaults

`func NewListLinkConfigsResponseDataWithDefaults() *ListLinkConfigsResponseData`

NewListLinkConfigsResponseDataWithDefaults instantiates a new ListLinkConfigsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListLinkConfigsResponseData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListLinkConfigsResponseData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListLinkConfigsResponseData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListLinkConfigsResponseData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListLinkConfigsResponseData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListLinkConfigsResponseData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ListLinkConfigsResponseData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ListLinkConfigsResponseData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ListLinkConfigsResponseData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *ListLinkConfigsResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLinkConfigsResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLinkConfigsResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ListLinkConfigsResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListLinkConfigsResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListLinkConfigsResponseData) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsDefault

`func (o *ListLinkConfigsResponseData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ListLinkConfigsResponseData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ListLinkConfigsResponseData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsReadOnly

`func (o *ListLinkConfigsResponseData) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *ListLinkConfigsResponseData) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *ListLinkConfigsResponseData) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsSensitive

`func (o *ListLinkConfigsResponseData) GetIsSensitive() bool`

GetIsSensitive returns the IsSensitive field if non-nil, zero value otherwise.

### GetIsSensitiveOk

`func (o *ListLinkConfigsResponseData) GetIsSensitiveOk() (*bool, bool)`

GetIsSensitiveOk returns a tuple with the IsSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSensitive

`func (o *ListLinkConfigsResponseData) SetIsSensitive(v bool)`

SetIsSensitive sets IsSensitive field to given value.


### GetSource

`func (o *ListLinkConfigsResponseData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListLinkConfigsResponseData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListLinkConfigsResponseData) SetSource(v string)`

SetSource sets Source field to given value.


### GetSynonyms

`func (o *ListLinkConfigsResponseData) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *ListLinkConfigsResponseData) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *ListLinkConfigsResponseData) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.


### GetLinkName

`func (o *ListLinkConfigsResponseData) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ListLinkConfigsResponseData) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ListLinkConfigsResponseData) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


