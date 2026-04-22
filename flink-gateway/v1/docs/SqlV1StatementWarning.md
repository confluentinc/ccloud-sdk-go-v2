# SqlV1StatementWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | **string** | Indicates the severity of the warning.  LOW: Indicates a low severity warning and for informing the user.  MODERATE: Indicates a moderate severity warning and may require user action. Could cause degraded statements if certain conditions apply.  CRITICAL: Indicates a critical severity warning and requires user action. It will cause degraded statements eventually.  | 
**CreatedAt** | **time.Time** | The timestamp when the warning was created. It is represented in RFC3339 format and is in UTC. | [readonly] 
**Reason** | **string** | A machine-readable short, upper case summary delimited by underscore. | 
**Message** | **string** | A human-readable string containing the description of the warning. | 

## Methods

### NewSqlV1StatementWarning

`func NewSqlV1StatementWarning(severity string, createdAt time.Time, reason string, message string, ) *SqlV1StatementWarning`

NewSqlV1StatementWarning instantiates a new SqlV1StatementWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementWarningWithDefaults

`func NewSqlV1StatementWarningWithDefaults() *SqlV1StatementWarning`

NewSqlV1StatementWarningWithDefaults instantiates a new SqlV1StatementWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *SqlV1StatementWarning) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SqlV1StatementWarning) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SqlV1StatementWarning) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetCreatedAt

`func (o *SqlV1StatementWarning) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SqlV1StatementWarning) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SqlV1StatementWarning) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReason

`func (o *SqlV1StatementWarning) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SqlV1StatementWarning) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SqlV1StatementWarning) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *SqlV1StatementWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SqlV1StatementWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SqlV1StatementWarning) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


