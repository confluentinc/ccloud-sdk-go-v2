# NotificationsV1ResourceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of Confluent Cloud resource this notification relates to. | 
**Crn** | **string** | CRN of the Confluent Cloud resource at delivery time. | 
**DisplayName** | **string** | Human-readable name of the resource captured at notification time. Does not update if the underlying resource is later renamed.  | 

## Methods

### NewNotificationsV1ResourceSnapshot

`func NewNotificationsV1ResourceSnapshot(type_ string, crn string, displayName string, ) *NotificationsV1ResourceSnapshot`

NewNotificationsV1ResourceSnapshot instantiates a new NotificationsV1ResourceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1ResourceSnapshotWithDefaults

`func NewNotificationsV1ResourceSnapshotWithDefaults() *NotificationsV1ResourceSnapshot`

NewNotificationsV1ResourceSnapshotWithDefaults instantiates a new NotificationsV1ResourceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationsV1ResourceSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationsV1ResourceSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationsV1ResourceSnapshot) SetType(v string)`

SetType sets Type field to given value.


### GetCrn

`func (o *NotificationsV1ResourceSnapshot) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *NotificationsV1ResourceSnapshot) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *NotificationsV1ResourceSnapshot) SetCrn(v string)`

SetCrn sets Crn field to given value.


### GetDisplayName

`func (o *NotificationsV1ResourceSnapshot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotificationsV1ResourceSnapshot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NotificationsV1ResourceSnapshot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


