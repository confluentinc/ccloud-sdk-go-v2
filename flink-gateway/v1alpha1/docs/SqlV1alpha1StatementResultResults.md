# SqlV1alpha1StatementResultResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]interface{}** | A data property that contains an array of results. Each entry in the array is a separate result.  The value of &#x60;op&#x60; attribute (if present) represents the kind of change that a row can describe in a changelog:  &#x60;0&#x60;: represents &#x60;INSERT&#x60; (&#x60;+I&#x60;), i.e. insertion operation;  &#x60;1&#x60;: represents &#x60;UPDATE_BEFORE&#x60; (&#x60;-U&#x60;), i.e. update operation with the previous content of the updated row. This kind should occur together with &#x60;UPDATE_AFTER&#x60; for modelling an update that needs to retract the previous row first. It is useful in cases of a non-idempotent update, i.e., an update of a row that is not  uniquely identifiable by a key;  &#x60;2&#x60;: represents &#x60;UPDATE_AFTER&#x60; (&#x60;+U&#x60;), i.e. update operation with new content of the updated row; This kind CAN occur together with &#x60;UPDATE_BEFORE&#x60; for modelling an update that needs to retract the previous row first or it describes an idempotent update, i.e., an update of a row that is uniquely identifiable by a key;  &#x60;3&#x60;: represents &#x60;DELETE&#x60; (&#x60;-D&#x60;), i.e. deletion operation;  Defaults to &#x60;0&#x60;.  | [optional] 

## Methods

### NewSqlV1alpha1StatementResultResults

`func NewSqlV1alpha1StatementResultResults() *SqlV1alpha1StatementResultResults`

NewSqlV1alpha1StatementResultResults instantiates a new SqlV1alpha1StatementResultResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1alpha1StatementResultResultsWithDefaults

`func NewSqlV1alpha1StatementResultResultsWithDefaults() *SqlV1alpha1StatementResultResults`

NewSqlV1alpha1StatementResultResultsWithDefaults instantiates a new SqlV1alpha1StatementResultResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SqlV1alpha1StatementResultResults) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1alpha1StatementResultResults) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1alpha1StatementResultResults) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SqlV1alpha1StatementResultResults) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


