# PartitionLevelTruncationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionId** | **int32** |  | 
**OffsetTruncatedTo** | **string** |  | 
**MessagesTruncated** | **string** |  | 

## Methods

### NewPartitionLevelTruncationData

`func NewPartitionLevelTruncationData(partitionId int32, offsetTruncatedTo string, messagesTruncated string, ) *PartitionLevelTruncationData`

NewPartitionLevelTruncationData instantiates a new PartitionLevelTruncationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionLevelTruncationDataWithDefaults

`func NewPartitionLevelTruncationDataWithDefaults() *PartitionLevelTruncationData`

NewPartitionLevelTruncationDataWithDefaults instantiates a new PartitionLevelTruncationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionId

`func (o *PartitionLevelTruncationData) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *PartitionLevelTruncationData) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *PartitionLevelTruncationData) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.


### GetOffsetTruncatedTo

`func (o *PartitionLevelTruncationData) GetOffsetTruncatedTo() string`

GetOffsetTruncatedTo returns the OffsetTruncatedTo field if non-nil, zero value otherwise.

### GetOffsetTruncatedToOk

`func (o *PartitionLevelTruncationData) GetOffsetTruncatedToOk() (*string, bool)`

GetOffsetTruncatedToOk returns a tuple with the OffsetTruncatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetTruncatedTo

`func (o *PartitionLevelTruncationData) SetOffsetTruncatedTo(v string)`

SetOffsetTruncatedTo sets OffsetTruncatedTo field to given value.


### GetMessagesTruncated

`func (o *PartitionLevelTruncationData) GetMessagesTruncated() string`

GetMessagesTruncated returns the MessagesTruncated field if non-nil, zero value otherwise.

### GetMessagesTruncatedOk

`func (o *PartitionLevelTruncationData) GetMessagesTruncatedOk() (*string, bool)`

GetMessagesTruncatedOk returns a tuple with the MessagesTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesTruncated

`func (o *PartitionLevelTruncationData) SetMessagesTruncated(v string)`

SetMessagesTruncated sets MessagesTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

