# SqlV1beta1ScalingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScalingState** | Pointer to **string** | OK: The statement runs at the right scale. PENDING_SCALE_DOWN: The statement requires less resources, and will be scaled down in the near future. PENDING_SCALE_UP: The statement requires more resources, and will be scaled up in the near future. POOL_EXHAUSTED: The statement requires more resources, but not enough resources are available.  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The last time the scaling status was updated. | [optional] [readonly] 

## Methods

### NewSqlV1beta1ScalingStatus

`func NewSqlV1beta1ScalingStatus() *SqlV1beta1ScalingStatus`

NewSqlV1beta1ScalingStatus instantiates a new SqlV1beta1ScalingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1beta1ScalingStatusWithDefaults

`func NewSqlV1beta1ScalingStatusWithDefaults() *SqlV1beta1ScalingStatus`

NewSqlV1beta1ScalingStatusWithDefaults instantiates a new SqlV1beta1ScalingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScalingState

`func (o *SqlV1beta1ScalingStatus) GetScalingState() string`

GetScalingState returns the ScalingState field if non-nil, zero value otherwise.

### GetScalingStateOk

`func (o *SqlV1beta1ScalingStatus) GetScalingStateOk() (*string, bool)`

GetScalingStateOk returns a tuple with the ScalingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingState

`func (o *SqlV1beta1ScalingStatus) SetScalingState(v string)`

SetScalingState sets ScalingState field to given value.

### HasScalingState

`func (o *SqlV1beta1ScalingStatus) HasScalingState() bool`

HasScalingState returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SqlV1beta1ScalingStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SqlV1beta1ScalingStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SqlV1beta1ScalingStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SqlV1beta1ScalingStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


