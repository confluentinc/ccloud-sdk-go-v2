# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **int64** |  | 
**ParentSnapshotId** | Pointer to **int64** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**TimestampMs** | **int64** |  | 
**ManifestList** | **string** | Location of the snapshot&#39;s manifest list file | 
**Summary** | [**SnapshotSummary**](SnapshotSummary.md) |  | 
**SchemaId** | Pointer to **int32** |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot(snapshotId int64, timestampMs int64, manifestList string, summary SnapshotSummary, ) *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *Snapshot) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Snapshot) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *Snapshot) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetParentSnapshotId

`func (o *Snapshot) GetParentSnapshotId() int64`

GetParentSnapshotId returns the ParentSnapshotId field if non-nil, zero value otherwise.

### GetParentSnapshotIdOk

`func (o *Snapshot) GetParentSnapshotIdOk() (*int64, bool)`

GetParentSnapshotIdOk returns a tuple with the ParentSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshotId

`func (o *Snapshot) SetParentSnapshotId(v int64)`

SetParentSnapshotId sets ParentSnapshotId field to given value.

### HasParentSnapshotId

`func (o *Snapshot) HasParentSnapshotId() bool`

HasParentSnapshotId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Snapshot) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Snapshot) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Snapshot) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Snapshot) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetTimestampMs

`func (o *Snapshot) GetTimestampMs() int64`

GetTimestampMs returns the TimestampMs field if non-nil, zero value otherwise.

### GetTimestampMsOk

`func (o *Snapshot) GetTimestampMsOk() (*int64, bool)`

GetTimestampMsOk returns a tuple with the TimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMs

`func (o *Snapshot) SetTimestampMs(v int64)`

SetTimestampMs sets TimestampMs field to given value.


### GetManifestList

`func (o *Snapshot) GetManifestList() string`

GetManifestList returns the ManifestList field if non-nil, zero value otherwise.

### GetManifestListOk

`func (o *Snapshot) GetManifestListOk() (*string, bool)`

GetManifestListOk returns a tuple with the ManifestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestList

`func (o *Snapshot) SetManifestList(v string)`

SetManifestList sets ManifestList field to given value.


### GetSummary

`func (o *Snapshot) GetSummary() SnapshotSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Snapshot) GetSummaryOk() (*SnapshotSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Snapshot) SetSummary(v SnapshotSummary)`

SetSummary sets Summary field to given value.


### GetSchemaId

`func (o *Snapshot) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *Snapshot) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *Snapshot) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *Snapshot) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


