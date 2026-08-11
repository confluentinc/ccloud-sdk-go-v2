# NotificationsV1RecommendedActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | Schema version of the &#x60;recommended_actions&#x60; payload. Increment when the payload shape changes in a non-backward-compatible way.  | 
**Content** | **string** | Human-readable body text describing the recommended actions. Rendered as Markdown for &#x60;version: 1&#x60;.  | 

## Methods

### NewNotificationsV1RecommendedActions

`func NewNotificationsV1RecommendedActions(version int32, content string, ) *NotificationsV1RecommendedActions`

NewNotificationsV1RecommendedActions instantiates a new NotificationsV1RecommendedActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1RecommendedActionsWithDefaults

`func NewNotificationsV1RecommendedActionsWithDefaults() *NotificationsV1RecommendedActions`

NewNotificationsV1RecommendedActionsWithDefaults instantiates a new NotificationsV1RecommendedActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *NotificationsV1RecommendedActions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NotificationsV1RecommendedActions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NotificationsV1RecommendedActions) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetContent

`func (o *NotificationsV1RecommendedActions) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationsV1RecommendedActions) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationsV1RecommendedActions) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


