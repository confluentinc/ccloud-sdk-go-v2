# QueryV1QueryResponseAsync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Result** | [**QueryV1AsyncJobReference**](QueryV1AsyncJobReference.md) | Identifies the background job created for this statement. | 

## Methods

### NewQueryV1QueryResponseAsync

`func NewQueryV1QueryResponseAsync(apiVersion string, kind string, result QueryV1AsyncJobReference, ) *QueryV1QueryResponseAsync`

NewQueryV1QueryResponseAsync instantiates a new QueryV1QueryResponseAsync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1QueryResponseAsyncWithDefaults

`func NewQueryV1QueryResponseAsyncWithDefaults() *QueryV1QueryResponseAsync`

NewQueryV1QueryResponseAsyncWithDefaults instantiates a new QueryV1QueryResponseAsync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QueryV1QueryResponseAsync) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QueryV1QueryResponseAsync) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QueryV1QueryResponseAsync) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *QueryV1QueryResponseAsync) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryV1QueryResponseAsync) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryV1QueryResponseAsync) SetKind(v string)`

SetKind sets Kind field to given value.


### GetResult

`func (o *QueryV1QueryResponseAsync) GetResult() QueryV1AsyncJobReference`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryV1QueryResponseAsync) GetResultOk() (*QueryV1AsyncJobReference, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryV1QueryResponseAsync) SetResult(v QueryV1AsyncJobReference)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


