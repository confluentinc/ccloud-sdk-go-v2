# LoadTableResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataLocation** | Pointer to **string** | May be null if the table is staged as part of a transaction | [optional] 
**Metadata** | [**TableMetadata**](TableMetadata.md) |  | 
**Config** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLoadTableResult

`func NewLoadTableResult(metadata TableMetadata, ) *LoadTableResult`

NewLoadTableResult instantiates a new LoadTableResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadTableResultWithDefaults

`func NewLoadTableResultWithDefaults() *LoadTableResult`

NewLoadTableResultWithDefaults instantiates a new LoadTableResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataLocation

`func (o *LoadTableResult) GetMetadataLocation() string`

GetMetadataLocation returns the MetadataLocation field if non-nil, zero value otherwise.

### GetMetadataLocationOk

`func (o *LoadTableResult) GetMetadataLocationOk() (*string, bool)`

GetMetadataLocationOk returns a tuple with the MetadataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLocation

`func (o *LoadTableResult) SetMetadataLocation(v string)`

SetMetadataLocation sets MetadataLocation field to given value.

### HasMetadataLocation

`func (o *LoadTableResult) HasMetadataLocation() bool`

HasMetadataLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *LoadTableResult) GetMetadata() TableMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LoadTableResult) GetMetadataOk() (*TableMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LoadTableResult) SetMetadata(v TableMetadata)`

SetMetadata sets Metadata field to given value.


### GetConfig

`func (o *LoadTableResult) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoadTableResult) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoadTableResult) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoadTableResult) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


