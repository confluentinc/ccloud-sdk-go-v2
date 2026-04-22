# SqlV1MaterializedTableList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | [readonly] 
**Kind** | **string** |  | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]SqlV1MaterializedTable**](SqlV1MaterializedTable.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewSqlV1MaterializedTableList

`func NewSqlV1MaterializedTableList(apiVersion string, kind string, metadata ListMeta, data []SqlV1MaterializedTable, ) *SqlV1MaterializedTableList`

NewSqlV1MaterializedTableList instantiates a new SqlV1MaterializedTableList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MaterializedTableListWithDefaults

`func NewSqlV1MaterializedTableListWithDefaults() *SqlV1MaterializedTableList`

NewSqlV1MaterializedTableListWithDefaults instantiates a new SqlV1MaterializedTableList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1MaterializedTableList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1MaterializedTableList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1MaterializedTableList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1MaterializedTableList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MaterializedTableList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MaterializedTableList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1MaterializedTableList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1MaterializedTableList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1MaterializedTableList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *SqlV1MaterializedTableList) GetData() []SqlV1MaterializedTable`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1MaterializedTableList) GetDataOk() (*[]SqlV1MaterializedTable, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1MaterializedTableList) SetData(v []SqlV1MaterializedTable)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


