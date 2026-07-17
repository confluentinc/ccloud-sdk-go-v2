# SqlV1ScalingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselineCfu** | Pointer to **int32** | The baseline number of Confluent Flink Units (CFUs) targeted for the statement. This is a best-effort target rather than a guarantee, and the statement may auto-scale above it.  | [optional] 

## Methods

### NewSqlV1ScalingSpec

`func NewSqlV1ScalingSpec() *SqlV1ScalingSpec`

NewSqlV1ScalingSpec instantiates a new SqlV1ScalingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ScalingSpecWithDefaults

`func NewSqlV1ScalingSpecWithDefaults() *SqlV1ScalingSpec`

NewSqlV1ScalingSpecWithDefaults instantiates a new SqlV1ScalingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaselineCfu

`func (o *SqlV1ScalingSpec) GetBaselineCfu() int32`

GetBaselineCfu returns the BaselineCfu field if non-nil, zero value otherwise.

### GetBaselineCfuOk

`func (o *SqlV1ScalingSpec) GetBaselineCfuOk() (*int32, bool)`

GetBaselineCfuOk returns a tuple with the BaselineCfu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCfu

`func (o *SqlV1ScalingSpec) SetBaselineCfu(v int32)`

SetBaselineCfu sets BaselineCfu field to given value.

### HasBaselineCfu

`func (o *SqlV1ScalingSpec) HasBaselineCfu() bool`

HasBaselineCfu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


