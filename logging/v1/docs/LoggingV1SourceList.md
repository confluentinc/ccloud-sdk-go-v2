# LoggingV1SourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]LoggingV1Source**](LoggingV1Source.md) | Every searchable source in the resource catalog. | 

## Methods

### NewLoggingV1SourceList

`func NewLoggingV1SourceList(apiVersion string, kind string, metadata ListMeta, data []LoggingV1Source, ) *LoggingV1SourceList`

NewLoggingV1SourceList instantiates a new LoggingV1SourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1SourceListWithDefaults

`func NewLoggingV1SourceListWithDefaults() *LoggingV1SourceList`

NewLoggingV1SourceListWithDefaults instantiates a new LoggingV1SourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *LoggingV1SourceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LoggingV1SourceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LoggingV1SourceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *LoggingV1SourceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LoggingV1SourceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LoggingV1SourceList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *LoggingV1SourceList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LoggingV1SourceList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LoggingV1SourceList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *LoggingV1SourceList) GetData() []LoggingV1Source`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoggingV1SourceList) GetDataOk() (*[]LoggingV1Source, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoggingV1SourceList) SetData(v []LoggingV1Source)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


