# KafkaQuotasV1ClientQuotaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]KafkaQuotasV1ClientQuota**](KafkaQuotasV1ClientQuota.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewKafkaQuotasV1ClientQuotaList

`func NewKafkaQuotasV1ClientQuotaList(apiVersion string, kind string, metadata ListMeta, data []KafkaQuotasV1ClientQuota, ) *KafkaQuotasV1ClientQuotaList`

NewKafkaQuotasV1ClientQuotaList instantiates a new KafkaQuotasV1ClientQuotaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaListWithDefaults

`func NewKafkaQuotasV1ClientQuotaListWithDefaults() *KafkaQuotasV1ClientQuotaList`

NewKafkaQuotasV1ClientQuotaListWithDefaults instantiates a new KafkaQuotasV1ClientQuotaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KafkaQuotasV1ClientQuotaList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaQuotasV1ClientQuotaList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaQuotasV1ClientQuotaList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KafkaQuotasV1ClientQuotaList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaQuotasV1ClientQuotaList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaQuotasV1ClientQuotaList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KafkaQuotasV1ClientQuotaList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaQuotasV1ClientQuotaList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaQuotasV1ClientQuotaList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *KafkaQuotasV1ClientQuotaList) GetData() []KafkaQuotasV1ClientQuota`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KafkaQuotasV1ClientQuotaList) GetDataOk() (*[]KafkaQuotasV1ClientQuota, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KafkaQuotasV1ClientQuotaList) SetData(v []KafkaQuotasV1ClientQuota)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


