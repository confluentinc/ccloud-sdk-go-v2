# SqlV1alpha1StatementException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the SQL statement exception. | [optional] [readonly] 
**Stacktrace** | Pointer to **string** | Stack trace of the statement exception. | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | The date and time at which the exception occurred. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | The custom metadata in the form of key value pairs attached to a statement exception when a failure occurs. | [optional] [readonly] 

## Methods

### NewSqlV1alpha1StatementException

`func NewSqlV1alpha1StatementException() *SqlV1alpha1StatementException`

NewSqlV1alpha1StatementException instantiates a new SqlV1alpha1StatementException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1alpha1StatementExceptionWithDefaults

`func NewSqlV1alpha1StatementExceptionWithDefaults() *SqlV1alpha1StatementException`

NewSqlV1alpha1StatementExceptionWithDefaults instantiates a new SqlV1alpha1StatementException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1alpha1StatementException) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1alpha1StatementException) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1alpha1StatementException) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1alpha1StatementException) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStacktrace

`func (o *SqlV1alpha1StatementException) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *SqlV1alpha1StatementException) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *SqlV1alpha1StatementException) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *SqlV1alpha1StatementException) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### GetTimestamp

`func (o *SqlV1alpha1StatementException) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SqlV1alpha1StatementException) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SqlV1alpha1StatementException) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SqlV1alpha1StatementException) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *SqlV1alpha1StatementException) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SqlV1alpha1StatementException) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SqlV1alpha1StatementException) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SqlV1alpha1StatementException) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


