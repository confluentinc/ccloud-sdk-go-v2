# SqlV1Watermark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 

## Methods

### NewSqlV1Watermark

`func NewSqlV1Watermark() *SqlV1Watermark`

NewSqlV1Watermark instantiates a new SqlV1Watermark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1WatermarkWithDefaults

`func NewSqlV1WatermarkWithDefaults() *SqlV1Watermark`

NewSqlV1WatermarkWithDefaults instantiates a new SqlV1Watermark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *SqlV1Watermark) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *SqlV1Watermark) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *SqlV1Watermark) SetColumn(v string)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *SqlV1Watermark) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetExpression

`func (o *SqlV1Watermark) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SqlV1Watermark) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SqlV1Watermark) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SqlV1Watermark) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


