# QueryV1QueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultFormat** | **string** | Encoding of &#x60;data&#x60;, echoing the &#x60;options.result_format&#x60; requested on the statement.  | [readonly] 
**Schema** | [**QueryV1ResultSchema**](QueryV1ResultSchema.md) | Column metadata describing the shape of the returned rows. | 
**Data** | [**QueryV1QueryResultDataOneOf**](QueryV1QueryResultDataOneOf.md) | Result rows in the encoding named by &#x60;result_format&#x60;: an array of rows for &#x60;JSON&#x60;, where each row is a positional array of column values ordered to match &#x60;schema.columns&#x60;; or a base64-encoded Arrow IPC stream for &#x60;ARROW_STREAM&#x60;, carrying native column types rather than stringified values. Inline Arrow payloads are size-capped by the service; larger results are delivered through the asynchronous 202 flow.  Under &#x60;JSON&#x60;, individual values — including those of complex &#x60;ROW&#x60;, &#x60;ARRAY&#x60;, &#x60;MAP&#x60; and &#x60;MULTISET&#x60; columns — are encoded as described by &#x60;ResultValue&#x60;.  | 

## Methods

### NewQueryV1QueryResult

`func NewQueryV1QueryResult(resultFormat string, schema QueryV1ResultSchema, data QueryV1QueryResultDataOneOf, ) *QueryV1QueryResult`

NewQueryV1QueryResult instantiates a new QueryV1QueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1QueryResultWithDefaults

`func NewQueryV1QueryResultWithDefaults() *QueryV1QueryResult`

NewQueryV1QueryResultWithDefaults instantiates a new QueryV1QueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultFormat

`func (o *QueryV1QueryResult) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *QueryV1QueryResult) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *QueryV1QueryResult) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.


### GetSchema

`func (o *QueryV1QueryResult) GetSchema() QueryV1ResultSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *QueryV1QueryResult) GetSchemaOk() (*QueryV1ResultSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *QueryV1QueryResult) SetSchema(v QueryV1ResultSchema)`

SetSchema sets Schema field to given value.


### GetData

`func (o *QueryV1QueryResult) GetData() QueryV1QueryResultDataOneOf`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryV1QueryResult) GetDataOk() (*QueryV1QueryResultDataOneOf, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryV1QueryResult) SetData(v QueryV1QueryResultDataOneOf)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


