# TableflowV1TableFlowTopicConfigsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableCompaction** | Pointer to **bool** | This flag determines whether to enable compaction for the Tableflow enabled topic. | [optional] [readonly] 
**EnablePartitioning** | Pointer to **bool** | This flag determines whether to enable partitioning for the Tableflow enabled topic. | [optional] [readonly] 
**RetentionMs** | Pointer to **string** | The maximum age, in milliseconds, of snapshots (for Iceberg) or versions (for Delta) to retain in the table for the Tableflow-enabled topic (snapshot/version expiration).  The default value is \&quot;604800000\&quot; milliseconds (equivalent to 7 days).  The minimum allowed value is \&quot;86400000\&quot; milliseconds (equivalent to 24 hours).  | [optional] 
**RecordFailureStrategy** | Pointer to **string** | The strategy to handle record failures in the Tableflow enabled topic during materialization.  For &#x60;SKIP&#x60;, we skip the bad records and move to the next record,  and for &#x60;SUSPEND&#x60;, we suspend the materialization of the topic.  | [optional] [default to "SUSPEND"]
**ErrorHandling** | Pointer to [**TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf**](TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf.md) | The error mode to handle record failures in the Tableflow enabled topic during materialization.  for &#x60;SKIP&#x60;, we skip the bad records and move to the next record,  for &#x60;SUSPEND&#x60;, we suspend the materialization of the topic,  and for &#x60;LOG&#x60;, we log the bad records to the DLQ and continue processing the rest of the records.  | [optional] 

## Methods

### NewTableflowV1TableFlowTopicConfigsSpec

`func NewTableflowV1TableFlowTopicConfigsSpec() *TableflowV1TableFlowTopicConfigsSpec`

NewTableflowV1TableFlowTopicConfigsSpec instantiates a new TableflowV1TableFlowTopicConfigsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableFlowTopicConfigsSpecWithDefaults

`func NewTableflowV1TableFlowTopicConfigsSpecWithDefaults() *TableflowV1TableFlowTopicConfigsSpec`

NewTableflowV1TableFlowTopicConfigsSpecWithDefaults instantiates a new TableflowV1TableFlowTopicConfigsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableCompaction

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetEnableCompaction() bool`

GetEnableCompaction returns the EnableCompaction field if non-nil, zero value otherwise.

### GetEnableCompactionOk

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetEnableCompactionOk() (*bool, bool)`

GetEnableCompactionOk returns a tuple with the EnableCompaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompaction

`func (o *TableflowV1TableFlowTopicConfigsSpec) SetEnableCompaction(v bool)`

SetEnableCompaction sets EnableCompaction field to given value.

### HasEnableCompaction

`func (o *TableflowV1TableFlowTopicConfigsSpec) HasEnableCompaction() bool`

HasEnableCompaction returns a boolean if a field has been set.

### GetEnablePartitioning

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetEnablePartitioning() bool`

GetEnablePartitioning returns the EnablePartitioning field if non-nil, zero value otherwise.

### GetEnablePartitioningOk

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetEnablePartitioningOk() (*bool, bool)`

GetEnablePartitioningOk returns a tuple with the EnablePartitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitioning

`func (o *TableflowV1TableFlowTopicConfigsSpec) SetEnablePartitioning(v bool)`

SetEnablePartitioning sets EnablePartitioning field to given value.

### HasEnablePartitioning

`func (o *TableflowV1TableFlowTopicConfigsSpec) HasEnablePartitioning() bool`

HasEnablePartitioning returns a boolean if a field has been set.

### GetRetentionMs

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetRetentionMs() string`

GetRetentionMs returns the RetentionMs field if non-nil, zero value otherwise.

### GetRetentionMsOk

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetRetentionMsOk() (*string, bool)`

GetRetentionMsOk returns a tuple with the RetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMs

`func (o *TableflowV1TableFlowTopicConfigsSpec) SetRetentionMs(v string)`

SetRetentionMs sets RetentionMs field to given value.

### HasRetentionMs

`func (o *TableflowV1TableFlowTopicConfigsSpec) HasRetentionMs() bool`

HasRetentionMs returns a boolean if a field has been set.

### GetRecordFailureStrategy

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetRecordFailureStrategy() string`

GetRecordFailureStrategy returns the RecordFailureStrategy field if non-nil, zero value otherwise.

### GetRecordFailureStrategyOk

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetRecordFailureStrategyOk() (*string, bool)`

GetRecordFailureStrategyOk returns a tuple with the RecordFailureStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFailureStrategy

`func (o *TableflowV1TableFlowTopicConfigsSpec) SetRecordFailureStrategy(v string)`

SetRecordFailureStrategy sets RecordFailureStrategy field to given value.

### HasRecordFailureStrategy

`func (o *TableflowV1TableFlowTopicConfigsSpec) HasRecordFailureStrategy() bool`

HasRecordFailureStrategy returns a boolean if a field has been set.

### GetErrorHandling

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetErrorHandling() TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf`

GetErrorHandling returns the ErrorHandling field if non-nil, zero value otherwise.

### GetErrorHandlingOk

`func (o *TableflowV1TableFlowTopicConfigsSpec) GetErrorHandlingOk() (*TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf, bool)`

GetErrorHandlingOk returns a tuple with the ErrorHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHandling

`func (o *TableflowV1TableFlowTopicConfigsSpec) SetErrorHandling(v TableflowV1TableFlowTopicConfigsSpecErrorHandlingOneOf)`

SetErrorHandling sets ErrorHandling field to given value.

### HasErrorHandling

`func (o *TableflowV1TableFlowTopicConfigsSpec) HasErrorHandling() bool`

HasErrorHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


