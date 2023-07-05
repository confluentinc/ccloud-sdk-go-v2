# ProduceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**TopicName** | Pointer to **string** |  | [optional] 
**PartitionId** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Key** | Pointer to [**NullableProduceResponseData**](ProduceResponseData.md) |  | [optional] 
**Value** | Pointer to [**NullableProduceResponseData**](ProduceResponseData.md) |  | [optional] 

## Methods

### NewProduceResponse

`func NewProduceResponse(errorCode int32, ) *ProduceResponse`

NewProduceResponse instantiates a new ProduceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProduceResponseWithDefaults

`func NewProduceResponseWithDefaults() *ProduceResponse`

NewProduceResponseWithDefaults instantiates a new ProduceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ProduceResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ProduceResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ProduceResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *ProduceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProduceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProduceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProduceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetClusterId

`func (o *ProduceResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ProduceResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ProduceResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ProduceResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetTopicName

`func (o *ProduceResponse) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *ProduceResponse) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *ProduceResponse) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *ProduceResponse) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetPartitionId

`func (o *ProduceResponse) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ProduceResponse) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ProduceResponse) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *ProduceResponse) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetOffset

`func (o *ProduceResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ProduceResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ProduceResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ProduceResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProduceResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProduceResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProduceResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProduceResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProduceResponse) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProduceResponse) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetKey

`func (o *ProduceResponse) GetKey() ProduceResponseData`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProduceResponse) GetKeyOk() (*ProduceResponseData, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProduceResponse) SetKey(v ProduceResponseData)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProduceResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProduceResponse) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProduceResponse) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ProduceResponse) GetValue() ProduceResponseData`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProduceResponse) GetValueOk() (*ProduceResponseData, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProduceResponse) SetValue(v ProduceResponseData)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProduceResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProduceResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProduceResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


