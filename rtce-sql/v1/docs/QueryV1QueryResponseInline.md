# QueryV1QueryResponseInline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Result** | [**QueryV1QueryResult**](QueryV1QueryResult.md) | The schema and data produced by the statement. | 

## Methods

### NewQueryV1QueryResponseInline

`func NewQueryV1QueryResponseInline(apiVersion string, kind string, result QueryV1QueryResult, ) *QueryV1QueryResponseInline`

NewQueryV1QueryResponseInline instantiates a new QueryV1QueryResponseInline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1QueryResponseInlineWithDefaults

`func NewQueryV1QueryResponseInlineWithDefaults() *QueryV1QueryResponseInline`

NewQueryV1QueryResponseInlineWithDefaults instantiates a new QueryV1QueryResponseInline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QueryV1QueryResponseInline) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QueryV1QueryResponseInline) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QueryV1QueryResponseInline) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *QueryV1QueryResponseInline) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryV1QueryResponseInline) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryV1QueryResponseInline) SetKind(v string)`

SetKind sets Kind field to given value.


### GetResult

`func (o *QueryV1QueryResponseInline) GetResult() QueryV1QueryResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryV1QueryResponseInline) GetResultOk() (*QueryV1QueryResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryV1QueryResponseInline) SetResult(v QueryV1QueryResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


