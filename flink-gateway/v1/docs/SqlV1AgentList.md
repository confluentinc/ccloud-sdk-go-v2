# SqlV1AgentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | [readonly] 
**Kind** | **string** |  | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]SqlV1Agent**](SqlV1Agent.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewSqlV1AgentList

`func NewSqlV1AgentList(apiVersion string, kind string, metadata ListMeta, data []SqlV1Agent, ) *SqlV1AgentList`

NewSqlV1AgentList instantiates a new SqlV1AgentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1AgentListWithDefaults

`func NewSqlV1AgentListWithDefaults() *SqlV1AgentList`

NewSqlV1AgentListWithDefaults instantiates a new SqlV1AgentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1AgentList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1AgentList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1AgentList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1AgentList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1AgentList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1AgentList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1AgentList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1AgentList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1AgentList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *SqlV1AgentList) GetData() []SqlV1Agent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1AgentList) GetDataOk() (*[]SqlV1Agent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1AgentList) SetData(v []SqlV1Agent)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


