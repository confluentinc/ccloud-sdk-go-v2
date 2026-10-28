# QueryV1QueryResultDataOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayOfQueryV1ResultRow** | Pointer to [**[]QueryV1ResultRow**](QueryV1ResultRow.md) | Set when &#x60;result_format&#x60; is &#x60;JSON&#x60;: the result rows. | [optional] 
**String** | Pointer to **string** | Set when &#x60;result_format&#x60; is &#x60;ARROW_STREAM&#x60;: the base64-encoded Arrow IPC stream. | [optional] 

At most one field is set. Both are nil when &#x60;data&#x60; is absent or JSON &#x60;null&#x60;, and a value with neither set marshals to &#x60;null&#x60;.

## Methods

### ArrayOfQueryV1ResultRowAsQueryV1QueryResultDataOneOf

`func ArrayOfQueryV1ResultRowAsQueryV1QueryResultDataOneOf(v *[]QueryV1ResultRow) QueryV1QueryResultDataOneOf`

ArrayOfQueryV1ResultRowAsQueryV1QueryResultDataOneOf returns the rows wrapped in QueryV1QueryResultDataOneOf.

### StringAsQueryV1QueryResultDataOneOf

`func StringAsQueryV1QueryResultDataOneOf(v *string) QueryV1QueryResultDataOneOf`

StringAsQueryV1QueryResultDataOneOf returns the Arrow stream wrapped in QueryV1QueryResultDataOneOf.

### GetActualInstance

`func (obj *QueryV1QueryResultDataOneOf) GetActualInstance() interface{}`

GetActualInstance returns the field that is set, or nil if neither is.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


