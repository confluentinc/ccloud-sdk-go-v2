# SqlV1MaterializedTableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaClusterId** | Pointer to **string** | The ID of the Kafka cluster hosting the Materialized Table&#39;s topic. This value must match the &#x60;kafka_cluster_id&#x60; path parameter. It is immutable after creation and is ignored or rejected on update if changed. | [optional] 
**ComputePoolId** | Pointer to **string** | The id associated with the compute pool in context. If not specified, the materialized table will use the default compute pool. The default pool is automatically determined by the system. | [optional] 
**Principal** | Pointer to **string** | The id of a principal this Materialized Table query runs as. | [optional] 
**Stopped** | Pointer to **bool** | Indicates whether the Materialized Table query should be stopped. | [optional] 
**TableOptions** | Pointer to **map[string]string** | Defines configuration properties for the table, equivalent to the SQL &#39;WITH&#39; clause | [optional] 
**SessionOptions** | Pointer to **map[string]string** | Session configurations equivalent to the SQL &#39;SET&#39; statement. Only applicable on creation; ignored on update. | [optional] 
**Columns** | Pointer to [**[]SqlV1ColumnDetails**](SqlV1ColumnDetails.md) | Details of each column in Materialized Table resource. If columns are not specified, we infer from query. If it&#39;s specified it must be compatible with the types in the query. | [optional] 
**Watermark** | Pointer to [**SqlV1Watermark**](SqlV1Watermark.md) |  | [optional] 
**Constraints** | Pointer to [**[]SqlV1Constraint**](SqlV1Constraint.md) | Specify table constraints. | [optional] 
**Distribution** | Pointer to [**SqlV1Distribution**](SqlV1Distribution.md) | Only applicable on creation; ignored on update. | [optional] 
**Query** | Pointer to **string** | Contains the query section (usually starting with a SELECT) of the latest Materialized Table. | [optional] 

## Methods

### NewSqlV1MaterializedTableSpec

`func NewSqlV1MaterializedTableSpec() *SqlV1MaterializedTableSpec`

NewSqlV1MaterializedTableSpec instantiates a new SqlV1MaterializedTableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableSpecWithDefaults

`func NewSqlV1MaterializedTableSpecWithDefaults() *SqlV1MaterializedTableSpec`

NewSqlV1MaterializedTableSpecWithDefaults instantiates a new SqlV1MaterializedTableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaClusterId

`func (o *SqlV1MaterializedTableSpec) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *SqlV1MaterializedTableSpec) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *SqlV1MaterializedTableSpec) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.

### HasKafkaClusterId

`func (o *SqlV1MaterializedTableSpec) HasKafkaClusterId() bool`

HasKafkaClusterId returns a boolean if a field has been set.

### GetComputePoolId

`func (o *SqlV1MaterializedTableSpec) GetComputePoolId() string`

GetComputePoolId returns the ComputePoolId field if non-nil, zero value otherwise.

### GetComputePoolIdOk

`func (o *SqlV1MaterializedTableSpec) GetComputePoolIdOk() (*string, bool)`

GetComputePoolIdOk returns a tuple with the ComputePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolId

`func (o *SqlV1MaterializedTableSpec) SetComputePoolId(v string)`

SetComputePoolId sets ComputePoolId field to given value.

### HasComputePoolId

`func (o *SqlV1MaterializedTableSpec) HasComputePoolId() bool`

HasComputePoolId returns a boolean if a field has been set.

### GetPrincipal

`func (o *SqlV1MaterializedTableSpec) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlV1MaterializedTableSpec) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlV1MaterializedTableSpec) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlV1MaterializedTableSpec) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetStopped

`func (o *SqlV1MaterializedTableSpec) GetStopped() bool`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *SqlV1MaterializedTableSpec) GetStoppedOk() (*bool, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *SqlV1MaterializedTableSpec) SetStopped(v bool)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *SqlV1MaterializedTableSpec) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### GetTableOptions

`func (o *SqlV1MaterializedTableSpec) GetTableOptions() map[string]string`

GetTableOptions returns the TableOptions field if non-nil, zero value otherwise.

### GetTableOptionsOk

`func (o *SqlV1MaterializedTableSpec) GetTableOptionsOk() (*map[string]string, bool)`

GetTableOptionsOk returns a tuple with the TableOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableOptions

`func (o *SqlV1MaterializedTableSpec) SetTableOptions(v map[string]string)`

SetTableOptions sets TableOptions field to given value.

### HasTableOptions

`func (o *SqlV1MaterializedTableSpec) HasTableOptions() bool`

HasTableOptions returns a boolean if a field has been set.

### GetSessionOptions

`func (o *SqlV1MaterializedTableSpec) GetSessionOptions() map[string]string`

GetSessionOptions returns the SessionOptions field if non-nil, zero value otherwise.

### GetSessionOptionsOk

`func (o *SqlV1MaterializedTableSpec) GetSessionOptionsOk() (*map[string]string, bool)`

GetSessionOptionsOk returns a tuple with the SessionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionOptions

`func (o *SqlV1MaterializedTableSpec) SetSessionOptions(v map[string]string)`

SetSessionOptions sets SessionOptions field to given value.

### HasSessionOptions

`func (o *SqlV1MaterializedTableSpec) HasSessionOptions() bool`

HasSessionOptions returns a boolean if a field has been set.

### GetColumns

`func (o *SqlV1MaterializedTableSpec) GetColumns() []SqlV1ColumnDetails`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SqlV1MaterializedTableSpec) GetColumnsOk() (*[]SqlV1ColumnDetails, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SqlV1MaterializedTableSpec) SetColumns(v []SqlV1ColumnDetails)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SqlV1MaterializedTableSpec) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetWatermark

`func (o *SqlV1MaterializedTableSpec) GetWatermark() SqlV1Watermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *SqlV1MaterializedTableSpec) GetWatermarkOk() (*SqlV1Watermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *SqlV1MaterializedTableSpec) SetWatermark(v SqlV1Watermark)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *SqlV1MaterializedTableSpec) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetConstraints

`func (o *SqlV1MaterializedTableSpec) GetConstraints() []SqlV1Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SqlV1MaterializedTableSpec) GetConstraintsOk() (*[]SqlV1Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SqlV1MaterializedTableSpec) SetConstraints(v []SqlV1Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SqlV1MaterializedTableSpec) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDistribution

`func (o *SqlV1MaterializedTableSpec) GetDistribution() SqlV1Distribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *SqlV1MaterializedTableSpec) GetDistributionOk() (*SqlV1Distribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *SqlV1MaterializedTableSpec) SetDistribution(v SqlV1Distribution)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *SqlV1MaterializedTableSpec) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetQuery

`func (o *SqlV1MaterializedTableSpec) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SqlV1MaterializedTableSpec) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SqlV1MaterializedTableSpec) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SqlV1MaterializedTableSpec) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


