# MirrorLag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | **int32** |  | 
**Lag** | **int64** |  | 
**LastSourceFetchOffset** | **int64** |  | 

## Methods

### NewMirrorLag

`func NewMirrorLag(partition int32, lag int64, lastSourceFetchOffset int64, ) *MirrorLag`

NewMirrorLag instantiates a new MirrorLag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirrorLagWithDefaults

`func NewMirrorLagWithDefaults() *MirrorLag`

NewMirrorLagWithDefaults instantiates a new MirrorLag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *MirrorLag) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *MirrorLag) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *MirrorLag) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetLag

`func (o *MirrorLag) GetLag() int64`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *MirrorLag) GetLagOk() (*int64, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *MirrorLag) SetLag(v int64)`

SetLag sets Lag field to given value.


### GetLastSourceFetchOffset

`func (o *MirrorLag) GetLastSourceFetchOffset() int64`

GetLastSourceFetchOffset returns the LastSourceFetchOffset field if non-nil, zero value otherwise.

### GetLastSourceFetchOffsetOk

`func (o *MirrorLag) GetLastSourceFetchOffsetOk() (*int64, bool)`

GetLastSourceFetchOffsetOk returns a tuple with the LastSourceFetchOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSourceFetchOffset

`func (o *MirrorLag) SetLastSourceFetchOffset(v int64)`

SetLastSourceFetchOffset sets LastSourceFetchOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


