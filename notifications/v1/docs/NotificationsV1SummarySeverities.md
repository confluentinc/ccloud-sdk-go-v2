# NotificationsV1SummarySeverities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | **string** | The severity level this entry counts. | 
**Count** | **int32** | Number of unread notifications at this severity. | 

## Methods

### NewNotificationsV1SummarySeverities

`func NewNotificationsV1SummarySeverities(severity string, count int32, ) *NotificationsV1SummarySeverities`

NewNotificationsV1SummarySeverities instantiates a new NotificationsV1SummarySeverities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SummarySeveritiesWithDefaults

`func NewNotificationsV1SummarySeveritiesWithDefaults() *NotificationsV1SummarySeverities`

NewNotificationsV1SummarySeveritiesWithDefaults instantiates a new NotificationsV1SummarySeverities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *NotificationsV1SummarySeverities) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationsV1SummarySeverities) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationsV1SummarySeverities) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetCount

`func (o *NotificationsV1SummarySeverities) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NotificationsV1SummarySeverities) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NotificationsV1SummarySeverities) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


