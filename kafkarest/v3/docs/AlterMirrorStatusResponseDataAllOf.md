# AlterMirrorStatusResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MirrorTopicName** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**ErrorCode** | **NullableInt32** |  | 
**MirrorLags** | [**MirrorLags**](MirrorLags.md) |  | 
**MessagesTruncated** | **NullableInt64** |  | 
**PartitionLevelTruncationData** | [**PartitionLevelTruncationDataList**](PartitionLevelTruncationDataList.md) |  | 

## Methods

### NewAlterMirrorStatusResponseDataAllOf

`func NewAlterMirrorStatusResponseDataAllOf(mirrorTopicName string, errorMessage NullableString, errorCode NullableInt32, mirrorLags MirrorLags, messagesTruncated NullableInt64, partitionLevelTruncationData PartitionLevelTruncationDataList, ) *AlterMirrorStatusResponseDataAllOf`

NewAlterMirrorStatusResponseDataAllOf instantiates a new AlterMirrorStatusResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterMirrorStatusResponseDataAllOfWithDefaults

`func NewAlterMirrorStatusResponseDataAllOfWithDefaults() *AlterMirrorStatusResponseDataAllOf`

NewAlterMirrorStatusResponseDataAllOfWithDefaults instantiates a new AlterMirrorStatusResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMirrorTopicName

`func (o *AlterMirrorStatusResponseDataAllOf) GetMirrorTopicName() string`

GetMirrorTopicName returns the MirrorTopicName field if non-nil, zero value otherwise.

### GetMirrorTopicNameOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetMirrorTopicNameOk() (*string, bool)`

GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicName

`func (o *AlterMirrorStatusResponseDataAllOf) SetMirrorTopicName(v string)`

SetMirrorTopicName sets MirrorTopicName field to given value.


### GetErrorMessage

`func (o *AlterMirrorStatusResponseDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AlterMirrorStatusResponseDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *AlterMirrorStatusResponseDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AlterMirrorStatusResponseDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorCode

`func (o *AlterMirrorStatusResponseDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AlterMirrorStatusResponseDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### SetErrorCodeNil

`func (o *AlterMirrorStatusResponseDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AlterMirrorStatusResponseDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMirrorLags

`func (o *AlterMirrorStatusResponseDataAllOf) GetMirrorLags() MirrorLags`

GetMirrorLags returns the MirrorLags field if non-nil, zero value otherwise.

### GetMirrorLagsOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetMirrorLagsOk() (*MirrorLags, bool)`

GetMirrorLagsOk returns a tuple with the MirrorLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorLags

`func (o *AlterMirrorStatusResponseDataAllOf) SetMirrorLags(v MirrorLags)`

SetMirrorLags sets MirrorLags field to given value.


### GetMessagesTruncated

`func (o *AlterMirrorStatusResponseDataAllOf) GetMessagesTruncated() int64`

GetMessagesTruncated returns the MessagesTruncated field if non-nil, zero value otherwise.

### GetMessagesTruncatedOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetMessagesTruncatedOk() (*int64, bool)`

GetMessagesTruncatedOk returns a tuple with the MessagesTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesTruncated

`func (o *AlterMirrorStatusResponseDataAllOf) SetMessagesTruncated(v int64)`

SetMessagesTruncated sets MessagesTruncated field to given value.


### SetMessagesTruncatedNil

`func (o *AlterMirrorStatusResponseDataAllOf) SetMessagesTruncatedNil(b bool)`

 SetMessagesTruncatedNil sets the value for MessagesTruncated to be an explicit nil

### UnsetMessagesTruncated
`func (o *AlterMirrorStatusResponseDataAllOf) UnsetMessagesTruncated()`

UnsetMessagesTruncated ensures that no value is present for MessagesTruncated, not even an explicit nil
### GetPartitionLevelTruncationData

`func (o *AlterMirrorStatusResponseDataAllOf) GetPartitionLevelTruncationData() PartitionLevelTruncationDataList`

GetPartitionLevelTruncationData returns the PartitionLevelTruncationData field if non-nil, zero value otherwise.

### GetPartitionLevelTruncationDataOk

`func (o *AlterMirrorStatusResponseDataAllOf) GetPartitionLevelTruncationDataOk() (*PartitionLevelTruncationDataList, bool)`

GetPartitionLevelTruncationDataOk returns a tuple with the PartitionLevelTruncationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionLevelTruncationData

`func (o *AlterMirrorStatusResponseDataAllOf) SetPartitionLevelTruncationData(v PartitionLevelTruncationDataList)`

SetPartitionLevelTruncationData sets PartitionLevelTruncationData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


