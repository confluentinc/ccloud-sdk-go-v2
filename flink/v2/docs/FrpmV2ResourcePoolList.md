# FrpmV2ResourcePoolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]FrpmV2ResourcePool**](FrpmV2ResourcePool.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewFrpmV2ResourcePoolList

`func NewFrpmV2ResourcePoolList(apiVersion string, kind string, metadata ListMeta, data []FrpmV2ResourcePool, ) *FrpmV2ResourcePoolList`

NewFrpmV2ResourcePoolList instantiates a new FrpmV2ResourcePoolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrpmV2ResourcePoolListWithDefaults

`func NewFrpmV2ResourcePoolListWithDefaults() *FrpmV2ResourcePoolList`

NewFrpmV2ResourcePoolListWithDefaults instantiates a new FrpmV2ResourcePoolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FrpmV2ResourcePoolList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FrpmV2ResourcePoolList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FrpmV2ResourcePoolList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FrpmV2ResourcePoolList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FrpmV2ResourcePoolList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FrpmV2ResourcePoolList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FrpmV2ResourcePoolList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FrpmV2ResourcePoolList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FrpmV2ResourcePoolList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *FrpmV2ResourcePoolList) GetData() []FrpmV2ResourcePool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FrpmV2ResourcePoolList) GetDataOk() (*[]FrpmV2ResourcePool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FrpmV2ResourcePoolList) SetData(v []FrpmV2ResourcePool)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


