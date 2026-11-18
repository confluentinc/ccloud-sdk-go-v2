# LoggingV1SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Always &#x60;v1&#x60; on this endpoint. | 
**Kind** | **string** | Always &#x60;Logs&#x60; on this endpoint. | 
**Data** | **[]map[string]interface{}** | Log documents, each limited to the returned fields listed for its resource type in the Resource types section above. The shape differs per resource type. | 
**Metadata** | Pointer to [**LoggingV1SearchResponseMetadata**](LoggingV1SearchResponseMetadata.md) |  | [optional] 
**Aggregations** | Pointer to [**LoggingV1AggregationResponse**](LoggingV1AggregationResponse.md) |  | [optional] 

## Methods

### NewLoggingV1SearchResponse

`func NewLoggingV1SearchResponse(apiVersion string, kind string, data []map[string]interface{}, ) *LoggingV1SearchResponse`

NewLoggingV1SearchResponse instantiates a new LoggingV1SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1SearchResponseWithDefaults

`func NewLoggingV1SearchResponseWithDefaults() *LoggingV1SearchResponse`

NewLoggingV1SearchResponseWithDefaults instantiates a new LoggingV1SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *LoggingV1SearchResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LoggingV1SearchResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LoggingV1SearchResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *LoggingV1SearchResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LoggingV1SearchResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LoggingV1SearchResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetData

`func (o *LoggingV1SearchResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoggingV1SearchResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoggingV1SearchResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.


### GetMetadata

`func (o *LoggingV1SearchResponse) GetMetadata() LoggingV1SearchResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LoggingV1SearchResponse) GetMetadataOk() (*LoggingV1SearchResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LoggingV1SearchResponse) SetMetadata(v LoggingV1SearchResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LoggingV1SearchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAggregations

`func (o *LoggingV1SearchResponse) GetAggregations() LoggingV1AggregationResponse`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *LoggingV1SearchResponse) GetAggregationsOk() (*LoggingV1AggregationResponse, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *LoggingV1SearchResponse) SetAggregations(v LoggingV1AggregationResponse)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *LoggingV1SearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


