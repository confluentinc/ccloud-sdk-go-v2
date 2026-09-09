# SwitchoverV1SwitchoverPairListMetadataPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int32** | The user-requested page size for this pagination. | [optional] 
**TotalSize** | Pointer to **int32** | The total number of items available in the full result set. | [optional] 
**NextPageToken** | Pointer to **string** | An opaque pagination token for requesting the next page of results. | [optional] 
**Expiration** | Pointer to **int32** | The number of seconds until &#x60;next_page_token&#x60; expires. | [optional] 

## Methods

### NewSwitchoverV1SwitchoverPairListMetadataPagination

`func NewSwitchoverV1SwitchoverPairListMetadataPagination() *SwitchoverV1SwitchoverPairListMetadataPagination`

NewSwitchoverV1SwitchoverPairListMetadataPagination instantiates a new SwitchoverV1SwitchoverPairListMetadataPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairListMetadataPaginationWithDefaults

`func NewSwitchoverV1SwitchoverPairListMetadataPaginationWithDefaults() *SwitchoverV1SwitchoverPairListMetadataPagination`

NewSwitchoverV1SwitchoverPairListMetadataPaginationWithDefaults instantiates a new SwitchoverV1SwitchoverPairListMetadataPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetNextPageToken

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetExpiration

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *SwitchoverV1SwitchoverPairListMetadataPagination) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


