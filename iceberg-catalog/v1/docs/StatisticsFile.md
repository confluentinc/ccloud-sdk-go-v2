# StatisticsFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **int64** |  | 
**StatisticsPath** | **string** |  | 
**FileSizeInBytes** | **int64** |  | 
**FileFooterSizeInBytes** | **int64** |  | 
**BlobMetadata** | [**[]BlobMetadata**](BlobMetadata.md) |  | 

## Methods

### NewStatisticsFile

`func NewStatisticsFile(snapshotId int64, statisticsPath string, fileSizeInBytes int64, fileFooterSizeInBytes int64, blobMetadata []BlobMetadata, ) *StatisticsFile`

NewStatisticsFile instantiates a new StatisticsFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsFileWithDefaults

`func NewStatisticsFileWithDefaults() *StatisticsFile`

NewStatisticsFileWithDefaults instantiates a new StatisticsFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *StatisticsFile) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *StatisticsFile) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *StatisticsFile) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetStatisticsPath

`func (o *StatisticsFile) GetStatisticsPath() string`

GetStatisticsPath returns the StatisticsPath field if non-nil, zero value otherwise.

### GetStatisticsPathOk

`func (o *StatisticsFile) GetStatisticsPathOk() (*string, bool)`

GetStatisticsPathOk returns a tuple with the StatisticsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPath

`func (o *StatisticsFile) SetStatisticsPath(v string)`

SetStatisticsPath sets StatisticsPath field to given value.


### GetFileSizeInBytes

`func (o *StatisticsFile) GetFileSizeInBytes() int64`

GetFileSizeInBytes returns the FileSizeInBytes field if non-nil, zero value otherwise.

### GetFileSizeInBytesOk

`func (o *StatisticsFile) GetFileSizeInBytesOk() (*int64, bool)`

GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeInBytes

`func (o *StatisticsFile) SetFileSizeInBytes(v int64)`

SetFileSizeInBytes sets FileSizeInBytes field to given value.


### GetFileFooterSizeInBytes

`func (o *StatisticsFile) GetFileFooterSizeInBytes() int64`

GetFileFooterSizeInBytes returns the FileFooterSizeInBytes field if non-nil, zero value otherwise.

### GetFileFooterSizeInBytesOk

`func (o *StatisticsFile) GetFileFooterSizeInBytesOk() (*int64, bool)`

GetFileFooterSizeInBytesOk returns a tuple with the FileFooterSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFooterSizeInBytes

`func (o *StatisticsFile) SetFileFooterSizeInBytes(v int64)`

SetFileFooterSizeInBytes sets FileFooterSizeInBytes field to given value.


### GetBlobMetadata

`func (o *StatisticsFile) GetBlobMetadata() []BlobMetadata`

GetBlobMetadata returns the BlobMetadata field if non-nil, zero value otherwise.

### GetBlobMetadataOk

`func (o *StatisticsFile) GetBlobMetadataOk() (*[]BlobMetadata, bool)`

GetBlobMetadataOk returns a tuple with the BlobMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobMetadata

`func (o *StatisticsFile) SetBlobMetadata(v []BlobMetadata)`

SetBlobMetadata sets BlobMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


