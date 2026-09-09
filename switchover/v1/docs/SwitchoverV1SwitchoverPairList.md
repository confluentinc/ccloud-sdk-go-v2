# SwitchoverV1SwitchoverPairList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**SwitchoverV1SwitchoverPairListMetadata**](SwitchoverV1SwitchoverPairListMetadata.md) |  | 
**Data** | [**[]SwitchoverV1SwitchoverPair**](SwitchoverV1SwitchoverPair.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource.  | 

## Methods

### NewSwitchoverV1SwitchoverPairList

`func NewSwitchoverV1SwitchoverPairList(apiVersion string, kind string, metadata SwitchoverV1SwitchoverPairListMetadata, data []SwitchoverV1SwitchoverPair, ) *SwitchoverV1SwitchoverPairList`

NewSwitchoverV1SwitchoverPairList instantiates a new SwitchoverV1SwitchoverPairList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairListWithDefaults

`func NewSwitchoverV1SwitchoverPairListWithDefaults() *SwitchoverV1SwitchoverPairList`

NewSwitchoverV1SwitchoverPairListWithDefaults instantiates a new SwitchoverV1SwitchoverPairList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SwitchoverV1SwitchoverPairList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SwitchoverV1SwitchoverPairList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SwitchoverV1SwitchoverPairList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SwitchoverV1SwitchoverPairList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SwitchoverV1SwitchoverPairList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SwitchoverV1SwitchoverPairList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SwitchoverV1SwitchoverPairList) GetMetadata() SwitchoverV1SwitchoverPairListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SwitchoverV1SwitchoverPairList) GetMetadataOk() (*SwitchoverV1SwitchoverPairListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SwitchoverV1SwitchoverPairList) SetMetadata(v SwitchoverV1SwitchoverPairListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *SwitchoverV1SwitchoverPairList) GetData() []SwitchoverV1SwitchoverPair`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SwitchoverV1SwitchoverPairList) GetDataOk() (*[]SwitchoverV1SwitchoverPair, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SwitchoverV1SwitchoverPairList) SetData(v []SwitchoverV1SwitchoverPair)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


