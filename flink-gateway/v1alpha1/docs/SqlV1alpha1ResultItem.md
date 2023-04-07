# SqlV1alpha1ResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **int32** | The kind of change that a row can describe in a changelog:  &#x60;0&#x60;: represents &#x60;INSERT&#x60; (&#x60;+I&#x60;), i.e. insertion operation;  &#x60;1&#x60;: represents &#x60;UPDATE_BEFORE&#x60; (&#x60;-U&#x60;), i.e. update operation with the previous content of the updated row. This kind should occur together with &#x60;UPDATE_AFTER&#x60; for modelling an update that needs to retract the previous row first. It is useful in cases of a non-idempotent update, i.e., an update of a row that is not  uniquely identifiable by a key;  &#x60;2&#x60;: represents &#x60;UPDATE_AFTER&#x60; (&#x60;+U&#x60;), i.e. update operation with new content of the updated row; This kind CAN occur together with &#x60;UPDATE_BEFORE&#x60; for modelling an update that needs to retract the previous row first or it describes an idempotent update, i.e., an update of a row that is uniquely identifiable by a key;  &#x60;3&#x60;: represents &#x60;DELETE&#x60; (&#x60;-D&#x60;), i.e. deletion operation;  Defaults to &#x60;0&#x60;.  | [optional] [default to 0]
**Row** | [**SqlV1alpha1ResultItemRow**](SqlV1alpha1ResultItemRow.md) |  | 

## Methods

### NewSqlV1alpha1ResultItem

`func NewSqlV1alpha1ResultItem(row SqlV1alpha1ResultItemRow, ) *SqlV1alpha1ResultItem`

NewSqlV1alpha1ResultItem instantiates a new SqlV1alpha1ResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1alpha1ResultItemWithDefaults

`func NewSqlV1alpha1ResultItemWithDefaults() *SqlV1alpha1ResultItem`

NewSqlV1alpha1ResultItemWithDefaults instantiates a new SqlV1alpha1ResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *SqlV1alpha1ResultItem) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SqlV1alpha1ResultItem) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SqlV1alpha1ResultItem) SetOp(v int32)`

SetOp sets Op field to given value.

### HasOp

`func (o *SqlV1alpha1ResultItem) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRow

`func (o *SqlV1alpha1ResultItem) GetRow() SqlV1alpha1ResultItemRow`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *SqlV1alpha1ResultItem) GetRowOk() (*SqlV1alpha1ResultItemRow, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *SqlV1alpha1ResultItem) SetRow(v SqlV1alpha1ResultItemRow)`

SetRow sets Row field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


