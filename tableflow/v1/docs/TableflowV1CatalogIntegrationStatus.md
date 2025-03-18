# TableflowV1CatalogIntegrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **string** | The lifecycle phase of the catalog integration:    PENDING: sync to catalog integration is pending;    CONNECTED: catalog integration is connected and syncing;    FAILED: catalog integration failed.  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if catalog integration is in a failed state. | [optional] [readonly] 
**LastSyncAt** | Pointer to **string** | The date and time at which the catalog was last synced. It is represented in RFC3339 format and is in UTC.  | [optional] [readonly] 

## Methods

### NewTableflowV1CatalogIntegrationStatus

`func NewTableflowV1CatalogIntegrationStatus() *TableflowV1CatalogIntegrationStatus`

NewTableflowV1CatalogIntegrationStatus instantiates a new TableflowV1CatalogIntegrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogIntegrationStatusWithDefaults

`func NewTableflowV1CatalogIntegrationStatusWithDefaults() *TableflowV1CatalogIntegrationStatus`

NewTableflowV1CatalogIntegrationStatusWithDefaults instantiates a new TableflowV1CatalogIntegrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *TableflowV1CatalogIntegrationStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TableflowV1CatalogIntegrationStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TableflowV1CatalogIntegrationStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TableflowV1CatalogIntegrationStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TableflowV1CatalogIntegrationStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TableflowV1CatalogIntegrationStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TableflowV1CatalogIntegrationStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TableflowV1CatalogIntegrationStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLastSyncAt

`func (o *TableflowV1CatalogIntegrationStatus) GetLastSyncAt() string`

GetLastSyncAt returns the LastSyncAt field if non-nil, zero value otherwise.

### GetLastSyncAtOk

`func (o *TableflowV1CatalogIntegrationStatus) GetLastSyncAtOk() (*string, bool)`

GetLastSyncAtOk returns a tuple with the LastSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAt

`func (o *TableflowV1CatalogIntegrationStatus) SetLastSyncAt(v string)`

SetLastSyncAt sets LastSyncAt field to given value.

### HasLastSyncAt

`func (o *TableflowV1CatalogIntegrationStatus) HasLastSyncAt() bool`

HasLastSyncAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


