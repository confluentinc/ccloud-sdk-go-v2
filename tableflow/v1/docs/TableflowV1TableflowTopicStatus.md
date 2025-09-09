# TableflowV1TableflowTopicStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **string** | The lifecycle phase of the Tableflow:    PENDING: Tableflow setup is pending;    RUNNING: Tableflow is currently running;    FAILED: Tableflow failed  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if Tableflow topic is in an error state | [optional] [readonly] 
**CatalogSyncStatuses** | Pointer to [**[]TableflowV1CatalogSyncStatus**](TableflowV1CatalogSyncStatus.md) | List of associated catalogs and their synchronization statuses for this Tableflow topic.  | [optional] [readonly] 
**FailingTableFormats** | Pointer to [**[]TableflowV1TableflowTopicStatusFailingTableFormats**](TableflowV1TableflowTopicStatusFailingTableFormats.md) | List of failing table formats for the Tableflow-enabled topic, including error details.  | [optional] [readonly] 
**WriteMode** | **string** | The write mode for the Tableflow-enabled topic, determining how data is written to the table.  | [readonly] 

## Methods

### NewTableflowV1TableflowTopicStatus

`func NewTableflowV1TableflowTopicStatus(writeMode string, ) *TableflowV1TableflowTopicStatus`

NewTableflowV1TableflowTopicStatus instantiates a new TableflowV1TableflowTopicStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicStatusWithDefaults

`func NewTableflowV1TableflowTopicStatusWithDefaults() *TableflowV1TableflowTopicStatus`

NewTableflowV1TableflowTopicStatusWithDefaults instantiates a new TableflowV1TableflowTopicStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *TableflowV1TableflowTopicStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TableflowV1TableflowTopicStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TableflowV1TableflowTopicStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TableflowV1TableflowTopicStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TableflowV1TableflowTopicStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TableflowV1TableflowTopicStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TableflowV1TableflowTopicStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TableflowV1TableflowTopicStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCatalogSyncStatuses

`func (o *TableflowV1TableflowTopicStatus) GetCatalogSyncStatuses() []TableflowV1CatalogSyncStatus`

GetCatalogSyncStatuses returns the CatalogSyncStatuses field if non-nil, zero value otherwise.

### GetCatalogSyncStatusesOk

`func (o *TableflowV1TableflowTopicStatus) GetCatalogSyncStatusesOk() (*[]TableflowV1CatalogSyncStatus, bool)`

GetCatalogSyncStatusesOk returns a tuple with the CatalogSyncStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogSyncStatuses

`func (o *TableflowV1TableflowTopicStatus) SetCatalogSyncStatuses(v []TableflowV1CatalogSyncStatus)`

SetCatalogSyncStatuses sets CatalogSyncStatuses field to given value.

### HasCatalogSyncStatuses

`func (o *TableflowV1TableflowTopicStatus) HasCatalogSyncStatuses() bool`

HasCatalogSyncStatuses returns a boolean if a field has been set.

### GetFailingTableFormats

`func (o *TableflowV1TableflowTopicStatus) GetFailingTableFormats() []TableflowV1TableflowTopicStatusFailingTableFormats`

GetFailingTableFormats returns the FailingTableFormats field if non-nil, zero value otherwise.

### GetFailingTableFormatsOk

`func (o *TableflowV1TableflowTopicStatus) GetFailingTableFormatsOk() (*[]TableflowV1TableflowTopicStatusFailingTableFormats, bool)`

GetFailingTableFormatsOk returns a tuple with the FailingTableFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingTableFormats

`func (o *TableflowV1TableflowTopicStatus) SetFailingTableFormats(v []TableflowV1TableflowTopicStatusFailingTableFormats)`

SetFailingTableFormats sets FailingTableFormats field to given value.

### HasFailingTableFormats

`func (o *TableflowV1TableflowTopicStatus) HasFailingTableFormats() bool`

HasFailingTableFormats returns a boolean if a field has been set.

### GetWriteMode

`func (o *TableflowV1TableflowTopicStatus) GetWriteMode() string`

GetWriteMode returns the WriteMode field if non-nil, zero value otherwise.

### GetWriteModeOk

`func (o *TableflowV1TableflowTopicStatus) GetWriteModeOk() (*string, bool)`

GetWriteModeOk returns a tuple with the WriteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMode

`func (o *TableflowV1TableflowTopicStatus) SetWriteMode(v string)`

SetWriteMode sets WriteMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


