# ProduceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionId** | Pointer to **NullableInt32** |  | [optional] 
**Headers** | Pointer to [**[]ProduceRequestHeader**](ProduceRequestHeader.md) |  | [optional] 
**Key** | Pointer to [**NullableProduceRequestData**](ProduceRequestData.md) |  | [optional] 
**Value** | Pointer to [**NullableProduceRequestData**](ProduceRequestData.md) |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewProduceRequest

`func NewProduceRequest() *ProduceRequest`

NewProduceRequest instantiates a new ProduceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProduceRequestWithDefaults

`func NewProduceRequestWithDefaults() *ProduceRequest`

NewProduceRequestWithDefaults instantiates a new ProduceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionId

`func (o *ProduceRequest) GetPartitionId() int32`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ProduceRequest) GetPartitionIdOk() (*int32, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ProduceRequest) SetPartitionId(v int32)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *ProduceRequest) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### SetPartitionIdNil

`func (o *ProduceRequest) SetPartitionIdNil(b bool)`

 SetPartitionIdNil sets the value for PartitionId to be an explicit nil

### UnsetPartitionId
`func (o *ProduceRequest) UnsetPartitionId()`

UnsetPartitionId ensures that no value is present for PartitionId, not even an explicit nil
### GetHeaders

`func (o *ProduceRequest) GetHeaders() []ProduceRequestHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProduceRequest) GetHeadersOk() (*[]ProduceRequestHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProduceRequest) SetHeaders(v []ProduceRequestHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProduceRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetKey

`func (o *ProduceRequest) GetKey() ProduceRequestData`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProduceRequest) GetKeyOk() (*ProduceRequestData, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProduceRequest) SetKey(v ProduceRequestData)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProduceRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProduceRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProduceRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ProduceRequest) GetValue() ProduceRequestData`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProduceRequest) GetValueOk() (*ProduceRequestData, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProduceRequest) SetValue(v ProduceRequestData)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProduceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProduceRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProduceRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTimestamp

`func (o *ProduceRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProduceRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProduceRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProduceRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProduceRequest) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProduceRequest) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


