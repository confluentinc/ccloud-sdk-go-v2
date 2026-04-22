# SqlV1StateLimitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateLimitState** | Pointer to **string** | OK: The statement is within state limits.  APPROACHING_SOFT_LIMIT: The statement is approaching soft state limits.  EXCEEDING_SOFT_LIMIT: The statement is exceeding soft state limits.  APPROACHING_HARD_LIMIT: The statement is approaching hard state limits.  EXCEEDING_HARD_LIMIT: The statement is exceeding hard state limits.  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The last time the state limit status was updated. | [optional] [readonly] 
**Detail** | Pointer to **string** | Details about why state limit status is in its current state. | [optional] [readonly] 

## Methods

### NewSqlV1StateLimitStatus

`func NewSqlV1StateLimitStatus() *SqlV1StateLimitStatus`

NewSqlV1StateLimitStatus instantiates a new SqlV1StateLimitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StateLimitStatusWithDefaults

`func NewSqlV1StateLimitStatusWithDefaults() *SqlV1StateLimitStatus`

NewSqlV1StateLimitStatusWithDefaults instantiates a new SqlV1StateLimitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateLimitState

`func (o *SqlV1StateLimitStatus) GetStateLimitState() string`

GetStateLimitState returns the StateLimitState field if non-nil, zero value otherwise.

### GetStateLimitStateOk

`func (o *SqlV1StateLimitStatus) GetStateLimitStateOk() (*string, bool)`

GetStateLimitStateOk returns a tuple with the StateLimitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateLimitState

`func (o *SqlV1StateLimitStatus) SetStateLimitState(v string)`

SetStateLimitState sets StateLimitState field to given value.

### HasStateLimitState

`func (o *SqlV1StateLimitStatus) HasStateLimitState() bool`

HasStateLimitState returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SqlV1StateLimitStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SqlV1StateLimitStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SqlV1StateLimitStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SqlV1StateLimitStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDetail

`func (o *SqlV1StateLimitStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1StateLimitStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1StateLimitStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1StateLimitStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


