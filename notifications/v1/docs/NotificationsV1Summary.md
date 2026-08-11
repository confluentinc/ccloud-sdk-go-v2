# NotificationsV1Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**UnreadCount** | **int32** | Total number of unread notifications. | 
**Severities** | [**[]NotificationsV1SummarySeverities**](NotificationsV1SummarySeverities.md) | Breakdown of unread notifications by severity level. One entry per severity present in the user&#39;s unread set; severities with a zero count may be omitted. New severity values may be added over time without a breaking change to this schema.  | 

## Methods

### NewNotificationsV1Summary

`func NewNotificationsV1Summary(apiVersion string, kind string, unreadCount int32, severities []NotificationsV1SummarySeverities, ) *NotificationsV1Summary`

NewNotificationsV1Summary instantiates a new NotificationsV1Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SummaryWithDefaults

`func NewNotificationsV1SummaryWithDefaults() *NotificationsV1Summary`

NewNotificationsV1SummaryWithDefaults instantiates a new NotificationsV1Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1Summary) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1Summary) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1Summary) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotificationsV1Summary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1Summary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1Summary) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUnreadCount

`func (o *NotificationsV1Summary) GetUnreadCount() int32`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *NotificationsV1Summary) GetUnreadCountOk() (*int32, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *NotificationsV1Summary) SetUnreadCount(v int32)`

SetUnreadCount sets UnreadCount field to given value.


### GetSeverities

`func (o *NotificationsV1Summary) GetSeverities() []NotificationsV1SummarySeverities`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *NotificationsV1Summary) GetSeveritiesOk() (*[]NotificationsV1SummarySeverities, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *NotificationsV1Summary) SetSeverities(v []NotificationsV1SummarySeverities)`

SetSeverities sets Severities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


