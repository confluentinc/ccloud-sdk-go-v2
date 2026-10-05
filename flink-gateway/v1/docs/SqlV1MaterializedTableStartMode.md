# SqlV1MaterializedTableStartMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The start mode strategy.  - &#x60;FROM_BEGINNING&#x60; reprocesses all available data from the source(s), starting from the beginning of history. - &#x60;FROM_NOW&#x60; consumes from the latest stream record, or from a past point obtained by subtracting &#x60;time_interval&#x60; from the current time. - &#x60;FROM_TIMESTAMP&#x60; begins processing from the absolute point in time given by &#x60;timestamp&#x60;. &#x60;timestamp&#x60; is required. - &#x60;RESUME_OR_FROM_BEGINNING&#x60; resumes from the previous job&#39;s final source offsets, falling back to &#x60;FROM_BEGINNING&#x60; if resume information is unavailable. - &#x60;RESUME_OR_FROM_NOW&#x60; resumes from the previous job&#39;s final source offsets, falling back to &#x60;FROM_NOW&#x60; (with optional &#x60;time_interval&#x60;) if resume information is unavailable. - &#x60;RESUME_OR_FROM_TIMESTAMP&#x60; resumes from the previous job&#39;s final source offsets, falling back to the supplied &#x60;timestamp&#x60; if resume information is unavailable. &#x60;timestamp&#x60; is required.  | 
**Timestamp** | Pointer to **time.Time** | Absolute point in time to start processing from. Required when &#x60;kind&#x60; is &#x60;FROM_TIMESTAMP&#x60; or &#x60;RESUME_OR_FROM_TIMESTAMP&#x60;; ignored otherwise. Must be an RFC 3339 timestamp that includes a time offset (either &#x60;Z&#x60; for UTC or an explicit offset such as &#x60;+02:00&#x60; or &#x60;-07:00&#x60;); the offset alone determines the exact timezone.  | [optional] 
**TimeInterval** | Pointer to [**SqlV1IntervalExpression**](SqlV1IntervalExpression.md) | Lookback interval applied to the &#x60;FROM_NOW&#x60; semantics. Only meaningful when &#x60;kind&#x60; is &#x60;FROM_NOW&#x60; or &#x60;RESUME_OR_FROM_NOW&#x60;. When set, consumption starts from a past point obtained by subtracting this interval from the current time (e.g. &#x60;1 HOURS&#x60; means \&quot;start from one hour ago\&quot;). When omitted with &#x60;FROM_NOW&#x60;, consumption starts from the latest stream record.  | [optional] 

## Methods

### NewSqlV1MaterializedTableStartMode

`func NewSqlV1MaterializedTableStartMode(kind string, ) *SqlV1MaterializedTableStartMode`

NewSqlV1MaterializedTableStartMode instantiates a new SqlV1MaterializedTableStartMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableStartModeWithDefaults

`func NewSqlV1MaterializedTableStartModeWithDefaults() *SqlV1MaterializedTableStartMode`

NewSqlV1MaterializedTableStartModeWithDefaults instantiates a new SqlV1MaterializedTableStartMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1MaterializedTableStartMode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MaterializedTableStartMode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MaterializedTableStartMode) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTimestamp

`func (o *SqlV1MaterializedTableStartMode) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SqlV1MaterializedTableStartMode) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SqlV1MaterializedTableStartMode) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SqlV1MaterializedTableStartMode) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimeInterval

`func (o *SqlV1MaterializedTableStartMode) GetTimeInterval() SqlV1IntervalExpression`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *SqlV1MaterializedTableStartMode) GetTimeIntervalOk() (*SqlV1IntervalExpression, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *SqlV1MaterializedTableStartMode) SetTimeInterval(v SqlV1IntervalExpression)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *SqlV1MaterializedTableStartMode) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


