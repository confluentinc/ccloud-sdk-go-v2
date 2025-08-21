# PimV2IntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]PimV2Integration**](PimV2Integration.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewPimV2IntegrationList

`func NewPimV2IntegrationList(apiVersion string, kind string, metadata ListMeta, data []PimV2Integration, ) *PimV2IntegrationList`

NewPimV2IntegrationList instantiates a new PimV2IntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2IntegrationListWithDefaults

`func NewPimV2IntegrationListWithDefaults() *PimV2IntegrationList`

NewPimV2IntegrationListWithDefaults instantiates a new PimV2IntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PimV2IntegrationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PimV2IntegrationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PimV2IntegrationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *PimV2IntegrationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2IntegrationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2IntegrationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *PimV2IntegrationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PimV2IntegrationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PimV2IntegrationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *PimV2IntegrationList) GetData() []PimV2Integration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PimV2IntegrationList) GetDataOk() (*[]PimV2Integration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PimV2IntegrationList) SetData(v []PimV2Integration)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


