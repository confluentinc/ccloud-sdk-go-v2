# BlobMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SnapshotId** | **int64** |  | 
**SequenceNumber** | **int64** |  | 
**Fields** | **[]int32** |  | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBlobMetadata

`func NewBlobMetadata(type_ string, snapshotId int64, sequenceNumber int64, fields []int32, ) *BlobMetadata`

NewBlobMetadata instantiates a new BlobMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobMetadataWithDefaults

`func NewBlobMetadataWithDefaults() *BlobMetadata`

NewBlobMetadataWithDefaults instantiates a new BlobMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlobMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetSnapshotId

`func (o *BlobMetadata) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *BlobMetadata) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *BlobMetadata) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.


### GetSequenceNumber

`func (o *BlobMetadata) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *BlobMetadata) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *BlobMetadata) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetFields

`func (o *BlobMetadata) GetFields() []int32`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BlobMetadata) GetFieldsOk() (*[]int32, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BlobMetadata) SetFields(v []int32)`

SetFields sets Fields field to given value.


### GetProperties

`func (o *BlobMetadata) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BlobMetadata) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BlobMetadata) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BlobMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


