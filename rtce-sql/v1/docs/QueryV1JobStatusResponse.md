# QueryV1JobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Result** | [**QueryV1JobStatus**](QueryV1JobStatus.md) | The status of the background statement. | 
**Errors** | Pointer to [**[]Error**](Error.md) | Failure details for the background statement. Populated only when &#x60;result.state&#x60; is &#x60;FAILED&#x60;; omitted otherwise.  | [optional] 

## Methods

### NewQueryV1JobStatusResponse

`func NewQueryV1JobStatusResponse(apiVersion string, kind string, result QueryV1JobStatus, ) *QueryV1JobStatusResponse`

NewQueryV1JobStatusResponse instantiates a new QueryV1JobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1JobStatusResponseWithDefaults

`func NewQueryV1JobStatusResponseWithDefaults() *QueryV1JobStatusResponse`

NewQueryV1JobStatusResponseWithDefaults instantiates a new QueryV1JobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QueryV1JobStatusResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QueryV1JobStatusResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QueryV1JobStatusResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *QueryV1JobStatusResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryV1JobStatusResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryV1JobStatusResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetResult

`func (o *QueryV1JobStatusResponse) GetResult() QueryV1JobStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryV1JobStatusResponse) GetResultOk() (*QueryV1JobStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryV1JobStatusResponse) SetResult(v QueryV1JobStatus)`

SetResult sets Result field to given value.


### GetErrors

`func (o *QueryV1JobStatusResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *QueryV1JobStatusResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *QueryV1JobStatusResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *QueryV1JobStatusResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


