# FiamV2FlinkIdentityPoolEnvRegionBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]FiamV2FlinkIdentityPoolEnvRegionBinding**](FiamV2FlinkIdentityPoolEnvRegionBinding.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewFiamV2FlinkIdentityPoolEnvRegionBindingList

`func NewFiamV2FlinkIdentityPoolEnvRegionBindingList(apiVersion string, kind string, metadata ListMeta, data []FiamV2FlinkIdentityPoolEnvRegionBinding, ) *FiamV2FlinkIdentityPoolEnvRegionBindingList`

NewFiamV2FlinkIdentityPoolEnvRegionBindingList instantiates a new FiamV2FlinkIdentityPoolEnvRegionBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiamV2FlinkIdentityPoolEnvRegionBindingListWithDefaults

`func NewFiamV2FlinkIdentityPoolEnvRegionBindingListWithDefaults() *FiamV2FlinkIdentityPoolEnvRegionBindingList`

NewFiamV2FlinkIdentityPoolEnvRegionBindingListWithDefaults instantiates a new FiamV2FlinkIdentityPoolEnvRegionBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetData() []FiamV2FlinkIdentityPoolEnvRegionBinding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) GetDataOk() (*[]FiamV2FlinkIdentityPoolEnvRegionBinding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FiamV2FlinkIdentityPoolEnvRegionBindingList) SetData(v []FiamV2FlinkIdentityPoolEnvRegionBinding)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


