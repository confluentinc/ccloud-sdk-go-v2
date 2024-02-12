# TableMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | **int32** |  | 
**TableUuid** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] 
**LastUpdatedMs** | Pointer to **int64** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Schemas** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**CurrentSchemaId** | Pointer to **int32** |  | [optional] 
**LastColumnId** | Pointer to **int32** |  | [optional] 
**PartitionSpecs** | Pointer to [**[]PartitionSpec**](PartitionSpec.md) |  | [optional] 
**DefaultSpecId** | Pointer to **int32** |  | [optional] 
**LastPartitionId** | Pointer to **int32** |  | [optional] 
**SortOrders** | Pointer to [**[]SortOrder**](SortOrder.md) |  | [optional] 
**DefaultSortOrderId** | Pointer to **int32** |  | [optional] 
**Snapshots** | Pointer to [**[]Snapshot**](Snapshot.md) |  | [optional] 
**Refs** | Pointer to [**SnapshotReferences**](SnapshotReferences.md) |  | [optional] 
**CurrentSnapshotId** | Pointer to **int64** |  | [optional] 
**LastSequenceNumber** | Pointer to **int64** |  | [optional] 
**SnapshotLog** | Pointer to [**SnapshotLog**](SnapshotLog.md) |  | [optional] 
**MetadataLog** | Pointer to [**MetadataLog**](MetadataLog.md) |  | [optional] 
**StatisticsFiles** | Pointer to [**[]StatisticsFile**](StatisticsFile.md) |  | [optional] 
**PartitionStatisticsFiles** | Pointer to [**[]PartitionStatisticsFile**](PartitionStatisticsFile.md) |  | [optional] 

## Methods

### NewTableMetadata

`func NewTableMetadata(formatVersion int32, tableUuid string, ) *TableMetadata`

NewTableMetadata instantiates a new TableMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableMetadataWithDefaults

`func NewTableMetadataWithDefaults() *TableMetadata`

NewTableMetadataWithDefaults instantiates a new TableMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormatVersion

`func (o *TableMetadata) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *TableMetadata) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *TableMetadata) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.


### GetTableUuid

`func (o *TableMetadata) GetTableUuid() string`

GetTableUuid returns the TableUuid field if non-nil, zero value otherwise.

### GetTableUuidOk

`func (o *TableMetadata) GetTableUuidOk() (*string, bool)`

GetTableUuidOk returns a tuple with the TableUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUuid

`func (o *TableMetadata) SetTableUuid(v string)`

SetTableUuid sets TableUuid field to given value.


### GetLocation

`func (o *TableMetadata) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TableMetadata) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TableMetadata) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TableMetadata) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLastUpdatedMs

`func (o *TableMetadata) GetLastUpdatedMs() int64`

GetLastUpdatedMs returns the LastUpdatedMs field if non-nil, zero value otherwise.

### GetLastUpdatedMsOk

`func (o *TableMetadata) GetLastUpdatedMsOk() (*int64, bool)`

GetLastUpdatedMsOk returns a tuple with the LastUpdatedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedMs

`func (o *TableMetadata) SetLastUpdatedMs(v int64)`

SetLastUpdatedMs sets LastUpdatedMs field to given value.

### HasLastUpdatedMs

`func (o *TableMetadata) HasLastUpdatedMs() bool`

HasLastUpdatedMs returns a boolean if a field has been set.

### GetProperties

`func (o *TableMetadata) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TableMetadata) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TableMetadata) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TableMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSchemas

`func (o *TableMetadata) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TableMetadata) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TableMetadata) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *TableMetadata) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetCurrentSchemaId

`func (o *TableMetadata) GetCurrentSchemaId() int32`

GetCurrentSchemaId returns the CurrentSchemaId field if non-nil, zero value otherwise.

### GetCurrentSchemaIdOk

`func (o *TableMetadata) GetCurrentSchemaIdOk() (*int32, bool)`

GetCurrentSchemaIdOk returns a tuple with the CurrentSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSchemaId

`func (o *TableMetadata) SetCurrentSchemaId(v int32)`

SetCurrentSchemaId sets CurrentSchemaId field to given value.

### HasCurrentSchemaId

`func (o *TableMetadata) HasCurrentSchemaId() bool`

HasCurrentSchemaId returns a boolean if a field has been set.

### GetLastColumnId

`func (o *TableMetadata) GetLastColumnId() int32`

GetLastColumnId returns the LastColumnId field if non-nil, zero value otherwise.

### GetLastColumnIdOk

`func (o *TableMetadata) GetLastColumnIdOk() (*int32, bool)`

GetLastColumnIdOk returns a tuple with the LastColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastColumnId

`func (o *TableMetadata) SetLastColumnId(v int32)`

SetLastColumnId sets LastColumnId field to given value.

### HasLastColumnId

`func (o *TableMetadata) HasLastColumnId() bool`

HasLastColumnId returns a boolean if a field has been set.

### GetPartitionSpecs

`func (o *TableMetadata) GetPartitionSpecs() []PartitionSpec`

GetPartitionSpecs returns the PartitionSpecs field if non-nil, zero value otherwise.

### GetPartitionSpecsOk

`func (o *TableMetadata) GetPartitionSpecsOk() (*[]PartitionSpec, bool)`

GetPartitionSpecsOk returns a tuple with the PartitionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionSpecs

`func (o *TableMetadata) SetPartitionSpecs(v []PartitionSpec)`

SetPartitionSpecs sets PartitionSpecs field to given value.

### HasPartitionSpecs

`func (o *TableMetadata) HasPartitionSpecs() bool`

HasPartitionSpecs returns a boolean if a field has been set.

### GetDefaultSpecId

`func (o *TableMetadata) GetDefaultSpecId() int32`

GetDefaultSpecId returns the DefaultSpecId field if non-nil, zero value otherwise.

### GetDefaultSpecIdOk

`func (o *TableMetadata) GetDefaultSpecIdOk() (*int32, bool)`

GetDefaultSpecIdOk returns a tuple with the DefaultSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSpecId

`func (o *TableMetadata) SetDefaultSpecId(v int32)`

SetDefaultSpecId sets DefaultSpecId field to given value.

### HasDefaultSpecId

`func (o *TableMetadata) HasDefaultSpecId() bool`

HasDefaultSpecId returns a boolean if a field has been set.

### GetLastPartitionId

`func (o *TableMetadata) GetLastPartitionId() int32`

GetLastPartitionId returns the LastPartitionId field if non-nil, zero value otherwise.

### GetLastPartitionIdOk

`func (o *TableMetadata) GetLastPartitionIdOk() (*int32, bool)`

GetLastPartitionIdOk returns a tuple with the LastPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPartitionId

`func (o *TableMetadata) SetLastPartitionId(v int32)`

SetLastPartitionId sets LastPartitionId field to given value.

### HasLastPartitionId

`func (o *TableMetadata) HasLastPartitionId() bool`

HasLastPartitionId returns a boolean if a field has been set.

### GetSortOrders

`func (o *TableMetadata) GetSortOrders() []SortOrder`

GetSortOrders returns the SortOrders field if non-nil, zero value otherwise.

### GetSortOrdersOk

`func (o *TableMetadata) GetSortOrdersOk() (*[]SortOrder, bool)`

GetSortOrdersOk returns a tuple with the SortOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrders

`func (o *TableMetadata) SetSortOrders(v []SortOrder)`

SetSortOrders sets SortOrders field to given value.

### HasSortOrders

`func (o *TableMetadata) HasSortOrders() bool`

HasSortOrders returns a boolean if a field has been set.

### GetDefaultSortOrderId

`func (o *TableMetadata) GetDefaultSortOrderId() int32`

GetDefaultSortOrderId returns the DefaultSortOrderId field if non-nil, zero value otherwise.

### GetDefaultSortOrderIdOk

`func (o *TableMetadata) GetDefaultSortOrderIdOk() (*int32, bool)`

GetDefaultSortOrderIdOk returns a tuple with the DefaultSortOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrderId

`func (o *TableMetadata) SetDefaultSortOrderId(v int32)`

SetDefaultSortOrderId sets DefaultSortOrderId field to given value.

### HasDefaultSortOrderId

`func (o *TableMetadata) HasDefaultSortOrderId() bool`

HasDefaultSortOrderId returns a boolean if a field has been set.

### GetSnapshots

`func (o *TableMetadata) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *TableMetadata) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *TableMetadata) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *TableMetadata) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetRefs

`func (o *TableMetadata) GetRefs() SnapshotReferences`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *TableMetadata) GetRefsOk() (*SnapshotReferences, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *TableMetadata) SetRefs(v SnapshotReferences)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *TableMetadata) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### GetCurrentSnapshotId

`func (o *TableMetadata) GetCurrentSnapshotId() int64`

GetCurrentSnapshotId returns the CurrentSnapshotId field if non-nil, zero value otherwise.

### GetCurrentSnapshotIdOk

`func (o *TableMetadata) GetCurrentSnapshotIdOk() (*int64, bool)`

GetCurrentSnapshotIdOk returns a tuple with the CurrentSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshotId

`func (o *TableMetadata) SetCurrentSnapshotId(v int64)`

SetCurrentSnapshotId sets CurrentSnapshotId field to given value.

### HasCurrentSnapshotId

`func (o *TableMetadata) HasCurrentSnapshotId() bool`

HasCurrentSnapshotId returns a boolean if a field has been set.

### GetLastSequenceNumber

`func (o *TableMetadata) GetLastSequenceNumber() int64`

GetLastSequenceNumber returns the LastSequenceNumber field if non-nil, zero value otherwise.

### GetLastSequenceNumberOk

`func (o *TableMetadata) GetLastSequenceNumberOk() (*int64, bool)`

GetLastSequenceNumberOk returns a tuple with the LastSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSequenceNumber

`func (o *TableMetadata) SetLastSequenceNumber(v int64)`

SetLastSequenceNumber sets LastSequenceNumber field to given value.

### HasLastSequenceNumber

`func (o *TableMetadata) HasLastSequenceNumber() bool`

HasLastSequenceNumber returns a boolean if a field has been set.

### GetSnapshotLog

`func (o *TableMetadata) GetSnapshotLog() SnapshotLog`

GetSnapshotLog returns the SnapshotLog field if non-nil, zero value otherwise.

### GetSnapshotLogOk

`func (o *TableMetadata) GetSnapshotLogOk() (*SnapshotLog, bool)`

GetSnapshotLogOk returns a tuple with the SnapshotLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLog

`func (o *TableMetadata) SetSnapshotLog(v SnapshotLog)`

SetSnapshotLog sets SnapshotLog field to given value.

### HasSnapshotLog

`func (o *TableMetadata) HasSnapshotLog() bool`

HasSnapshotLog returns a boolean if a field has been set.

### GetMetadataLog

`func (o *TableMetadata) GetMetadataLog() MetadataLog`

GetMetadataLog returns the MetadataLog field if non-nil, zero value otherwise.

### GetMetadataLogOk

`func (o *TableMetadata) GetMetadataLogOk() (*MetadataLog, bool)`

GetMetadataLogOk returns a tuple with the MetadataLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLog

`func (o *TableMetadata) SetMetadataLog(v MetadataLog)`

SetMetadataLog sets MetadataLog field to given value.

### HasMetadataLog

`func (o *TableMetadata) HasMetadataLog() bool`

HasMetadataLog returns a boolean if a field has been set.

### GetStatisticsFiles

`func (o *TableMetadata) GetStatisticsFiles() []StatisticsFile`

GetStatisticsFiles returns the StatisticsFiles field if non-nil, zero value otherwise.

### GetStatisticsFilesOk

`func (o *TableMetadata) GetStatisticsFilesOk() (*[]StatisticsFile, bool)`

GetStatisticsFilesOk returns a tuple with the StatisticsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsFiles

`func (o *TableMetadata) SetStatisticsFiles(v []StatisticsFile)`

SetStatisticsFiles sets StatisticsFiles field to given value.

### HasStatisticsFiles

`func (o *TableMetadata) HasStatisticsFiles() bool`

HasStatisticsFiles returns a boolean if a field has been set.

### GetPartitionStatisticsFiles

`func (o *TableMetadata) GetPartitionStatisticsFiles() []PartitionStatisticsFile`

GetPartitionStatisticsFiles returns the PartitionStatisticsFiles field if non-nil, zero value otherwise.

### GetPartitionStatisticsFilesOk

`func (o *TableMetadata) GetPartitionStatisticsFilesOk() (*[]PartitionStatisticsFile, bool)`

GetPartitionStatisticsFilesOk returns a tuple with the PartitionStatisticsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionStatisticsFiles

`func (o *TableMetadata) SetPartitionStatisticsFiles(v []PartitionStatisticsFile)`

SetPartitionStatisticsFiles sets PartitionStatisticsFiles field to given value.

### HasPartitionStatisticsFiles

`func (o *TableMetadata) HasPartitionStatisticsFiles() bool`

HasPartitionStatisticsFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


