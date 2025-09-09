# TableflowV1CatalogSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogIntegrationId** | Pointer to **string** | The ID of the catalog integration | [optional] 
**CatalogType** | Pointer to **string** | The type of the external catalog | [optional] 
**SyncStatus** | Pointer to **string** | The current synchronization status:    PENDING: sync is pending;    SYNCED: successfully synced;    FAILED: sync failed;    DISCONNECTED: catalog integration is disconnected.  | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Error message if the sync failed. This field is only present when &#x60;sync_status&#x60; is &#x60;FAILED&#x60;.  | [optional] 

## Methods

### NewTableflowV1CatalogSyncStatus

`func NewTableflowV1CatalogSyncStatus() *TableflowV1CatalogSyncStatus`

NewTableflowV1CatalogSyncStatus instantiates a new TableflowV1CatalogSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1CatalogSyncStatusWithDefaults

`func NewTableflowV1CatalogSyncStatusWithDefaults() *TableflowV1CatalogSyncStatus`

NewTableflowV1CatalogSyncStatusWithDefaults instantiates a new TableflowV1CatalogSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogIntegrationId

`func (o *TableflowV1CatalogSyncStatus) GetCatalogIntegrationId() string`

GetCatalogIntegrationId returns the CatalogIntegrationId field if non-nil, zero value otherwise.

### GetCatalogIntegrationIdOk

`func (o *TableflowV1CatalogSyncStatus) GetCatalogIntegrationIdOk() (*string, bool)`

GetCatalogIntegrationIdOk returns a tuple with the CatalogIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIntegrationId

`func (o *TableflowV1CatalogSyncStatus) SetCatalogIntegrationId(v string)`

SetCatalogIntegrationId sets CatalogIntegrationId field to given value.

### HasCatalogIntegrationId

`func (o *TableflowV1CatalogSyncStatus) HasCatalogIntegrationId() bool`

HasCatalogIntegrationId returns a boolean if a field has been set.

### GetCatalogType

`func (o *TableflowV1CatalogSyncStatus) GetCatalogType() string`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *TableflowV1CatalogSyncStatus) GetCatalogTypeOk() (*string, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *TableflowV1CatalogSyncStatus) SetCatalogType(v string)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *TableflowV1CatalogSyncStatus) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.

### GetSyncStatus

`func (o *TableflowV1CatalogSyncStatus) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *TableflowV1CatalogSyncStatus) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *TableflowV1CatalogSyncStatus) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *TableflowV1CatalogSyncStatus) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TableflowV1CatalogSyncStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TableflowV1CatalogSyncStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TableflowV1CatalogSyncStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TableflowV1CatalogSyncStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TableflowV1CatalogSyncStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TableflowV1CatalogSyncStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


