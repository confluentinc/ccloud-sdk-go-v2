# CclV1CustomCodeLoggingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]CclV1CustomCodeLogging**](CclV1CustomCodeLogging.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewCclV1CustomCodeLoggingList

`func NewCclV1CustomCodeLoggingList(apiVersion string, kind string, metadata ListMeta, data []CclV1CustomCodeLogging, ) *CclV1CustomCodeLoggingList`

NewCclV1CustomCodeLoggingList instantiates a new CclV1CustomCodeLoggingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCclV1CustomCodeLoggingListWithDefaults

`func NewCclV1CustomCodeLoggingListWithDefaults() *CclV1CustomCodeLoggingList`

NewCclV1CustomCodeLoggingListWithDefaults instantiates a new CclV1CustomCodeLoggingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CclV1CustomCodeLoggingList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CclV1CustomCodeLoggingList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CclV1CustomCodeLoggingList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *CclV1CustomCodeLoggingList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CclV1CustomCodeLoggingList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CclV1CustomCodeLoggingList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *CclV1CustomCodeLoggingList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CclV1CustomCodeLoggingList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CclV1CustomCodeLoggingList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *CclV1CustomCodeLoggingList) GetData() []CclV1CustomCodeLogging`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CclV1CustomCodeLoggingList) GetDataOk() (*[]CclV1CustomCodeLogging, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CclV1CustomCodeLoggingList) SetData(v []CclV1CustomCodeLogging)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


