# FcpmV2Quickstart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ComputePools** | Pointer to [**[]FcpmV2QuickstartComputePool**](FcpmV2QuickstartComputePool.md) | The list of compute pools to quickstart.  | [optional] 

## Methods

### NewFcpmV2Quickstart

`func NewFcpmV2Quickstart() *FcpmV2Quickstart`

NewFcpmV2Quickstart instantiates a new FcpmV2Quickstart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2QuickstartWithDefaults

`func NewFcpmV2QuickstartWithDefaults() *FcpmV2Quickstart`

NewFcpmV2QuickstartWithDefaults instantiates a new FcpmV2Quickstart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2Quickstart) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2Quickstart) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2Quickstart) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2Quickstart) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2Quickstart) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2Quickstart) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2Quickstart) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2Quickstart) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2Quickstart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2Quickstart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2Quickstart) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2Quickstart) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2Quickstart) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2Quickstart) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2Quickstart) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2Quickstart) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetComputePools

`func (o *FcpmV2Quickstart) GetComputePools() []FcpmV2QuickstartComputePool`

GetComputePools returns the ComputePools field if non-nil, zero value otherwise.

### GetComputePoolsOk

`func (o *FcpmV2Quickstart) GetComputePoolsOk() (*[]FcpmV2QuickstartComputePool, bool)`

GetComputePoolsOk returns a tuple with the ComputePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePools

`func (o *FcpmV2Quickstart) SetComputePools(v []FcpmV2QuickstartComputePool)`

SetComputePools sets ComputePools field to given value.

### HasComputePools

`func (o *FcpmV2Quickstart) HasComputePools() bool`

HasComputePools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


