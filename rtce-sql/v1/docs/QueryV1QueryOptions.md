# QueryV1QueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResultRows** | Pointer to **int32** | Maximum rows to return inline, up to 200.  | [optional] 
**ResultFormat** | Pointer to **string** | Encoding for the returned rows, echoed back as &#x60;result.result_format&#x60;. &#x60;JSON&#x60; returns &#x60;result.data&#x60; as an array of rows, requiring clients to coerce values using &#x60;result.schema&#x60;. &#x60;ARROW_STREAM&#x60; returns &#x60;result.data&#x60; as a base64-encoded Arrow IPC stream, preserving native column types. Inline &#x60;ARROW_STREAM&#x60; payloads are size-capped; larger results are delivered through the asynchronous 202 flow.  | [optional] [default to "JSON"]

## Methods

### NewQueryV1QueryOptions

`func NewQueryV1QueryOptions() *QueryV1QueryOptions`

NewQueryV1QueryOptions instantiates a new QueryV1QueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1QueryOptionsWithDefaults

`func NewQueryV1QueryOptionsWithDefaults() *QueryV1QueryOptions`

NewQueryV1QueryOptionsWithDefaults instantiates a new QueryV1QueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResultRows

`func (o *QueryV1QueryOptions) GetMaxResultRows() int32`

GetMaxResultRows returns the MaxResultRows field if non-nil, zero value otherwise.

### GetMaxResultRowsOk

`func (o *QueryV1QueryOptions) GetMaxResultRowsOk() (*int32, bool)`

GetMaxResultRowsOk returns a tuple with the MaxResultRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultRows

`func (o *QueryV1QueryOptions) SetMaxResultRows(v int32)`

SetMaxResultRows sets MaxResultRows field to given value.

### HasMaxResultRows

`func (o *QueryV1QueryOptions) HasMaxResultRows() bool`

HasMaxResultRows returns a boolean if a field has been set.

### GetResultFormat

`func (o *QueryV1QueryOptions) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *QueryV1QueryOptions) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *QueryV1QueryOptions) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.

### HasResultFormat

`func (o *QueryV1QueryOptions) HasResultFormat() bool`

HasResultFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


