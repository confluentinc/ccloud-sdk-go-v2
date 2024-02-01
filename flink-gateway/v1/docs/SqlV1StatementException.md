# SqlV1StatementException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] 
**Name** | Pointer to **string** | Name of the SQL statement exception. | [optional] [readonly] 
**Message** | Pointer to **string** | Error message of the statement exception. | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | The date and time at which the exception occurred. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewSqlV1StatementException

`func NewSqlV1StatementException() *SqlV1StatementException`

NewSqlV1StatementException instantiates a new SqlV1StatementException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementExceptionWithDefaults

`func NewSqlV1StatementExceptionWithDefaults() *SqlV1StatementException`

NewSqlV1StatementExceptionWithDefaults instantiates a new SqlV1StatementException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1StatementException) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1StatementException) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1StatementException) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1StatementException) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *SqlV1StatementException) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1StatementException) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1StatementException) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1StatementException) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *SqlV1StatementException) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SqlV1StatementException) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SqlV1StatementException) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SqlV1StatementException) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *SqlV1StatementException) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SqlV1StatementException) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SqlV1StatementException) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SqlV1StatementException) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


