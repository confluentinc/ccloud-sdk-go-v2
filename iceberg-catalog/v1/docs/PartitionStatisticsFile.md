# PartitionStatisticsFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **int64** |  | 
**StatisticsPath** | **string** |  | 
**FileSizeInBytes** | **int64** |  | 

## Methods

### NewPartitionStatisticsFile

`func NewPartitionStatisticsFile(snapshotId int64, statisticsPath string, fileSizeInBytes int64, ) *PartitionStatisticsFile`

NewPartitionStatisticsFile instantiates a new PartitionStatisticsFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionStatisticsFileWithDefaults

`func NewPartitionStatisticsFileWithDefaults() *PartitionStatisticsFile`

NewPartitionStatisticsFileWithDefaults instantiates a new PartitionStatisticsFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *PartitionStatisticsFile) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *PartitionStatisticsFile) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *PartitionStatisticsFile) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetStatisticsPath

`func (o *PartitionStatisticsFile) GetStatisticsPath() string`

GetStatisticsPath returns the StatisticsPath field if non-nil, zero value otherwise.

### GetStatisticsPathOk

`func (o *PartitionStatisticsFile) GetStatisticsPathOk() (*string, bool)`

GetStatisticsPathOk returns a tuple with the StatisticsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPath

`func (o *PartitionStatisticsFile) SetStatisticsPath(v string)`

SetStatisticsPath sets StatisticsPath field to given value.


### GetFileSizeInBytes

`func (o *PartitionStatisticsFile) GetFileSizeInBytes() int64`

GetFileSizeInBytes returns the FileSizeInBytes field if non-nil, zero value otherwise.

### GetFileSizeInBytesOk

`func (o *PartitionStatisticsFile) GetFileSizeInBytesOk() (*int64, bool)`

GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeInBytes

`func (o *PartitionStatisticsFile) SetFileSizeInBytes(v int64)`

SetFileSizeInBytes sets FileSizeInBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

