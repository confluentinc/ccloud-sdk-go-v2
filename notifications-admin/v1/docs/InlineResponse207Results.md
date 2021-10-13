# InlineResponse207Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | Status code of message sent to specified integration. | [optional] 
**OrgId** | Pointer to **int32** | Organization Id for which the notification is sent. | [optional] 
**IntegrationKind** | Pointer to **string** | The type of the target system where notification is delivered. | [optional] 

## Methods

### NewInlineResponse207Results

`func NewInlineResponse207Results() *InlineResponse207Results`

NewInlineResponse207Results instantiates a new InlineResponse207Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse207ResultsWithDefaults

`func NewInlineResponse207ResultsWithDefaults() *InlineResponse207Results`

NewInlineResponse207ResultsWithDefaults instantiates a new InlineResponse207Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse207Results) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse207Results) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse207Results) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse207Results) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrgId

`func (o *InlineResponse207Results) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InlineResponse207Results) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InlineResponse207Results) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InlineResponse207Results) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetIntegrationKind

`func (o *InlineResponse207Results) GetIntegrationKind() string`

GetIntegrationKind returns the IntegrationKind field if non-nil, zero value otherwise.

### GetIntegrationKindOk

`func (o *InlineResponse207Results) GetIntegrationKindOk() (*string, bool)`

GetIntegrationKindOk returns a tuple with the IntegrationKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKind

`func (o *InlineResponse207Results) SetIntegrationKind(v string)`

SetIntegrationKind sets IntegrationKind field to given value.

### HasIntegrationKind

`func (o *InlineResponse207Results) HasIntegrationKind() bool`

HasIntegrationKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


