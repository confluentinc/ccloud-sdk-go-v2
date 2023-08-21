# SqlV1beta1StatementResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | 
**Kind** | **string** | Kind defines the object this REST resource represents. | 
**Metadata** | [**ResultListMeta**](ResultListMeta.md) |  | 
**Results** | Pointer to [**SqlV1beta1StatementResultResults**](sql.v1beta1.StatementResultResults.md) |  | [optional] 

## Methods

### NewSqlV1beta1StatementResult

`func NewSqlV1beta1StatementResult(apiVersion string, kind string, metadata ResultListMeta, ) *SqlV1beta1StatementResult`

NewSqlV1beta1StatementResult instantiates a new SqlV1beta1StatementResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1beta1StatementResultWithDefaults

`func NewSqlV1beta1StatementResultWithDefaults() *SqlV1beta1StatementResult`

NewSqlV1beta1StatementResultWithDefaults instantiates a new SqlV1beta1StatementResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1beta1StatementResult) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1beta1StatementResult) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1beta1StatementResult) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1beta1StatementResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1beta1StatementResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1beta1StatementResult) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1beta1StatementResult) GetMetadata() ResultListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1beta1StatementResult) GetMetadataOk() (*ResultListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1beta1StatementResult) SetMetadata(v ResultListMeta)`

SetMetadata sets Metadata field to given value.


### GetResults

`func (o *SqlV1beta1StatementResult) GetResults() SqlV1beta1StatementResultResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SqlV1beta1StatementResult) GetResultsOk() (*SqlV1beta1StatementResultResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SqlV1beta1StatementResult) SetResults(v SqlV1beta1StatementResultResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *SqlV1beta1StatementResult) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


