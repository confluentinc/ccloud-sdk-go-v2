# RtceV1Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the region. | [optional] [readonly] 
**Region** | Pointer to **string** | The cloud service provider region. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The human-readable display name for the region. | [optional] [readonly] 

## Methods

### NewRtceV1Region

`func NewRtceV1Region() *RtceV1Region`

NewRtceV1Region instantiates a new RtceV1Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RegionWithDefaults

`func NewRtceV1RegionWithDefaults() *RtceV1Region`

NewRtceV1RegionWithDefaults instantiates a new RtceV1Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RtceV1Region) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RtceV1Region) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RtceV1Region) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RtceV1Region) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RtceV1Region) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RtceV1Region) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RtceV1Region) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RtceV1Region) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *RtceV1Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RtceV1Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RtceV1Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RtceV1Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *RtceV1Region) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RtceV1Region) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RtceV1Region) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RtceV1Region) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCloud

`func (o *RtceV1Region) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *RtceV1Region) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *RtceV1Region) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *RtceV1Region) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *RtceV1Region) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RtceV1Region) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RtceV1Region) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RtceV1Region) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDisplayName

`func (o *RtceV1Region) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RtceV1Region) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RtceV1Region) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RtceV1Region) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


