# PimV1IntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]PimV1Integration**](PimV1Integration.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewPimV1IntegrationList

`func NewPimV1IntegrationList(apiVersion string, kind string, metadata ListMeta, data []PimV1Integration, ) *PimV1IntegrationList`

NewPimV1IntegrationList instantiates a new PimV1IntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV1IntegrationListWithDefaults

`func NewPimV1IntegrationListWithDefaults() *PimV1IntegrationList`

NewPimV1IntegrationListWithDefaults instantiates a new PimV1IntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PimV1IntegrationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PimV1IntegrationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PimV1IntegrationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *PimV1IntegrationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV1IntegrationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV1IntegrationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *PimV1IntegrationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PimV1IntegrationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PimV1IntegrationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *PimV1IntegrationList) GetData() []PimV1Integration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PimV1IntegrationList) GetDataOk() (*[]PimV1Integration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PimV1IntegrationList) SetData(v []PimV1Integration)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


