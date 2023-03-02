# FcpmV2Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The display name. | [optional] [readonly] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the region. | [optional] [readonly] 
**RegionName** | Pointer to **string** | The region name. | [optional] [readonly] 

## Methods

### NewFcpmV2Region

`func NewFcpmV2Region() *FcpmV2Region`

NewFcpmV2Region instantiates a new FcpmV2Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2RegionWithDefaults

`func NewFcpmV2RegionWithDefaults() *FcpmV2Region`

NewFcpmV2RegionWithDefaults instantiates a new FcpmV2Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2Region) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2Region) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2Region) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2Region) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2Region) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2Region) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2Region) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2Region) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2Region) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2Region) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2Region) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2Region) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *FcpmV2Region) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FcpmV2Region) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FcpmV2Region) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FcpmV2Region) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *FcpmV2Region) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FcpmV2Region) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FcpmV2Region) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FcpmV2Region) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegionName

`func (o *FcpmV2Region) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *FcpmV2Region) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *FcpmV2Region) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *FcpmV2Region) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


