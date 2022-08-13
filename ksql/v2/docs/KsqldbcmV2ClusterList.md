# KsqldbcmV2ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]KsqldbcmV2Cluster**](KsqldbcmV2Cluster.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewKsqldbcmV2ClusterList

`func NewKsqldbcmV2ClusterList(apiVersion string, kind string, metadata ListMeta, data []KsqldbcmV2Cluster, ) *KsqldbcmV2ClusterList`

NewKsqldbcmV2ClusterList instantiates a new KsqldbcmV2ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV2ClusterListWithDefaults

`func NewKsqldbcmV2ClusterListWithDefaults() *KsqldbcmV2ClusterList`

NewKsqldbcmV2ClusterListWithDefaults instantiates a new KsqldbcmV2ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KsqldbcmV2ClusterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KsqldbcmV2ClusterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KsqldbcmV2ClusterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KsqldbcmV2ClusterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KsqldbcmV2ClusterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KsqldbcmV2ClusterList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KsqldbcmV2ClusterList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KsqldbcmV2ClusterList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KsqldbcmV2ClusterList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *KsqldbcmV2ClusterList) GetData() []KsqldbcmV2Cluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KsqldbcmV2ClusterList) GetDataOk() (*[]KsqldbcmV2Cluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KsqldbcmV2ClusterList) SetData(v []KsqldbcmV2Cluster)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


