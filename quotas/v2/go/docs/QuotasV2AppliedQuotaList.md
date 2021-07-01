# QuotasV2AppliedQuotaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]QuotasV2AppliedQuota**](QuotasV2AppliedQuota.md) |  | 

## Methods

### NewQuotasV2AppliedQuotaList

`func NewQuotasV2AppliedQuotaList(apiVersion string, kind string, metadata ListMeta, data []QuotasV2AppliedQuota, ) *QuotasV2AppliedQuotaList`

NewQuotasV2AppliedQuotaList instantiates a new QuotasV2AppliedQuotaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotasV2AppliedQuotaListWithDefaults

`func NewQuotasV2AppliedQuotaListWithDefaults() *QuotasV2AppliedQuotaList`

NewQuotasV2AppliedQuotaListWithDefaults instantiates a new QuotasV2AppliedQuotaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QuotasV2AppliedQuotaList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QuotasV2AppliedQuotaList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QuotasV2AppliedQuotaList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *QuotasV2AppliedQuotaList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotasV2AppliedQuotaList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotasV2AppliedQuotaList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *QuotasV2AppliedQuotaList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuotasV2AppliedQuotaList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuotasV2AppliedQuotaList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *QuotasV2AppliedQuotaList) GetData() []QuotasV2AppliedQuota`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuotasV2AppliedQuotaList) GetDataOk() (*[]QuotasV2AppliedQuota, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuotasV2AppliedQuotaList) SetData(v []QuotasV2AppliedQuota)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


