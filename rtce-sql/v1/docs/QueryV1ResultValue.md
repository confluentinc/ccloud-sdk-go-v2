# QueryV1ResultValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayOfQueryV1ResultValue** | Pointer to [**[]QueryV1ResultValue**](QueryV1ResultValue.md) | Set when the value is a JSON array: a &#x60;ROW&#x60;, &#x60;ARRAY&#x60;, &#x60;MAP&#x60; or &#x60;MULTISET&#x60; value. | [optional] 
**String** | Pointer to **string** | Set when the value is a JSON string: a scalar value. | [optional] 

At most one field is set. Both are nil for JSON &#x60;null&#x60; (SQL NULL), and a value with neither set marshals to &#x60;null&#x60;.

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


