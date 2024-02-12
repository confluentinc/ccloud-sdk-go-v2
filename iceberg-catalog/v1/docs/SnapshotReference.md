# SnapshotReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SnapshotId** | **int64** |  | 
**MaxRefAgeMs** | Pointer to **int64** |  | [optional] 
**MaxSnapshotAgeMs** | Pointer to **int64** |  | [optional] 
**MinSnapshotsToKeep** | Pointer to **int32** |  | [optional] 

## Methods

### NewSnapshotReference

`func NewSnapshotReference(type_ string, snapshotId int64, ) *SnapshotReference`

NewSnapshotReference instantiates a new SnapshotReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotReferenceWithDefaults

`func NewSnapshotReferenceWithDefaults() *SnapshotReference`

NewSnapshotReferenceWithDefaults instantiates a new SnapshotReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SnapshotReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotReference) SetType(v string)`

SetType sets Type field to given value.


### GetSnapshotId

`func (o *SnapshotReference) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotReference) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SnapshotReference) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetMaxRefAgeMs

`func (o *SnapshotReference) GetMaxRefAgeMs() int64`

GetMaxRefAgeMs returns the MaxRefAgeMs field if non-nil, zero value otherwise.

### GetMaxRefAgeMsOk

`func (o *SnapshotReference) GetMaxRefAgeMsOk() (*int64, bool)`

GetMaxRefAgeMsOk returns a tuple with the MaxRefAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRefAgeMs

`func (o *SnapshotReference) SetMaxRefAgeMs(v int64)`

SetMaxRefAgeMs sets MaxRefAgeMs field to given value.

### HasMaxRefAgeMs

`func (o *SnapshotReference) HasMaxRefAgeMs() bool`

HasMaxRefAgeMs returns a boolean if a field has been set.

### GetMaxSnapshotAgeMs

`func (o *SnapshotReference) GetMaxSnapshotAgeMs() int64`

GetMaxSnapshotAgeMs returns the MaxSnapshotAgeMs field if non-nil, zero value otherwise.

### GetMaxSnapshotAgeMsOk

`func (o *SnapshotReference) GetMaxSnapshotAgeMsOk() (*int64, bool)`

GetMaxSnapshotAgeMsOk returns a tuple with the MaxSnapshotAgeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotAgeMs

`func (o *SnapshotReference) SetMaxSnapshotAgeMs(v int64)`

SetMaxSnapshotAgeMs sets MaxSnapshotAgeMs field to given value.

### HasMaxSnapshotAgeMs

`func (o *SnapshotReference) HasMaxSnapshotAgeMs() bool`

HasMaxSnapshotAgeMs returns a boolean if a field has been set.

### GetMinSnapshotsToKeep

`func (o *SnapshotReference) GetMinSnapshotsToKeep() int32`

GetMinSnapshotsToKeep returns the MinSnapshotsToKeep field if non-nil, zero value otherwise.

### GetMinSnapshotsToKeepOk

`func (o *SnapshotReference) GetMinSnapshotsToKeepOk() (*int32, bool)`

GetMinSnapshotsToKeepOk returns a tuple with the MinSnapshotsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSnapshotsToKeep

`func (o *SnapshotReference) SetMinSnapshotsToKeep(v int32)`

SetMinSnapshotsToKeep sets MinSnapshotsToKeep field to given value.

### HasMinSnapshotsToKeep

`func (o *SnapshotReference) HasMinSnapshotsToKeep() bool`

HasMinSnapshotsToKeep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


