# QueryV1JobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The lifecycle state of the background statement. A job is always at least RUNNING by the time it can be polled — there is no earlier \&quot;queued\&quot; state.  RUNNING: statement execution is in progress;  SUCCEEDED: statement finished successfully; see &#x60;result_urls&#x60; for where to download the result;  FAILED: statement execution failed; see &#x60;errors&#x60; for details.  | [readonly] 
**ResultUrls** | Pointer to **[]string** | Presigned URL(s) the result can be downloaded from. Always present (an empty array if there&#39;s nothing to download yet), not just when &#x60;state&#x60; is &#x60;SUCCEEDED&#x60; — reserved for a future partial-results feature that could populate this incrementally while &#x60;state&#x60; is still &#x60;RUNNING&#x60;; today it&#39;s only ever non-empty once &#x60;state&#x60; is &#x60;SUCCEEDED&#x60;. Empty for &#x60;FAILED&#x60;. Each URL expires after a limited time — re-poll this endpoint to obtain a fresh one if it lapses before download.  | [optional] [readonly] 

## Methods

### NewQueryV1JobStatus

`func NewQueryV1JobStatus(state string, ) *QueryV1JobStatus`

NewQueryV1JobStatus instantiates a new QueryV1JobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1JobStatusWithDefaults

`func NewQueryV1JobStatusWithDefaults() *QueryV1JobStatus`

NewQueryV1JobStatusWithDefaults instantiates a new QueryV1JobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *QueryV1JobStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QueryV1JobStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QueryV1JobStatus) SetState(v string)`

SetState sets State field to given value.


### GetResultUrls

`func (o *QueryV1JobStatus) GetResultUrls() []string`

GetResultUrls returns the ResultUrls field if non-nil, zero value otherwise.

### GetResultUrlsOk

`func (o *QueryV1JobStatus) GetResultUrlsOk() (*[]string, bool)`

GetResultUrlsOk returns a tuple with the ResultUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrls

`func (o *QueryV1JobStatus) SetResultUrls(v []string)`

SetResultUrls sets ResultUrls field to given value.

### HasResultUrls

`func (o *QueryV1JobStatus) HasResultUrls() bool`

HasResultUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


