# KsqldbcmV1ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]KsqldbcmV1Cluster**](KsqldbcmV1Cluster.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewKsqldbcmV1ClusterList

`func NewKsqldbcmV1ClusterList(apiVersion string, kind string, metadata ListMeta, data []KsqldbcmV1Cluster, ) *KsqldbcmV1ClusterList`

NewKsqldbcmV1ClusterList instantiates a new KsqldbcmV1ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV1ClusterListWithDefaults

`func NewKsqldbcmV1ClusterListWithDefaults() *KsqldbcmV1ClusterList`

NewKsqldbcmV1ClusterListWithDefaults instantiates a new KsqldbcmV1ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KsqldbcmV1ClusterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KsqldbcmV1ClusterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KsqldbcmV1ClusterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KsqldbcmV1ClusterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KsqldbcmV1ClusterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KsqldbcmV1ClusterList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KsqldbcmV1ClusterList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KsqldbcmV1ClusterList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KsqldbcmV1ClusterList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *KsqldbcmV1ClusterList) GetData() []KsqldbcmV1Cluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KsqldbcmV1ClusterList) GetDataOk() (*[]KsqldbcmV1Cluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KsqldbcmV1ClusterList) SetData(v []KsqldbcmV1Cluster)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


