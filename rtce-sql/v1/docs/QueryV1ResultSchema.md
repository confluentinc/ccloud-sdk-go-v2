# QueryV1ResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]QueryV1ColumnDetails**](QueryV1ColumnDetails.md) | Column definitions in result order, matching the order of values within each row.  | 

## Methods

### NewQueryV1ResultSchema

`func NewQueryV1ResultSchema(columns []QueryV1ColumnDetails, ) *QueryV1ResultSchema`

NewQueryV1ResultSchema instantiates a new QueryV1ResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1ResultSchemaWithDefaults

`func NewQueryV1ResultSchemaWithDefaults() *QueryV1ResultSchema`

NewQueryV1ResultSchemaWithDefaults instantiates a new QueryV1ResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *QueryV1ResultSchema) GetColumns() []QueryV1ColumnDetails`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryV1ResultSchema) GetColumnsOk() (*[]QueryV1ColumnDetails, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryV1ResultSchema) SetColumns(v []QueryV1ColumnDetails)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


