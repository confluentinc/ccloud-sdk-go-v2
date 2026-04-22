# SqlV1Distribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind of distribution. | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**BucketCount** | Pointer to **int32** | The number of buckets. | [optional] [default to 6]

## Methods

### NewSqlV1Distribution

`func NewSqlV1Distribution() *SqlV1Distribution`

NewSqlV1Distribution instantiates a new SqlV1Distribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1DistributionWithDefaults

`func NewSqlV1DistributionWithDefaults() *SqlV1Distribution`

NewSqlV1DistributionWithDefaults instantiates a new SqlV1Distribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1Distribution) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1Distribution) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1Distribution) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1Distribution) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetKeys

`func (o *SqlV1Distribution) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SqlV1Distribution) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SqlV1Distribution) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SqlV1Distribution) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetBucketCount

`func (o *SqlV1Distribution) GetBucketCount() int32`

GetBucketCount returns the BucketCount field if non-nil, zero value otherwise.

### GetBucketCountOk

`func (o *SqlV1Distribution) GetBucketCountOk() (*int32, bool)`

GetBucketCountOk returns a tuple with the BucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketCount

`func (o *SqlV1Distribution) SetBucketCount(v int32)`

SetBucketCount sets BucketCount field to given value.

### HasBucketCount

`func (o *SqlV1Distribution) HasBucketCount() bool`

HasBucketCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


