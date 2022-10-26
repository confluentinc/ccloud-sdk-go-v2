# StreamGovernanceV2Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**StreamGovernanceV2RegionSpec**](StreamGovernanceV2RegionSpec.md) |  | [optional] 

## Methods

### NewStreamGovernanceV2Region

`func NewStreamGovernanceV2Region() *StreamGovernanceV2Region`

NewStreamGovernanceV2Region instantiates a new StreamGovernanceV2Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV2RegionWithDefaults

`func NewStreamGovernanceV2RegionWithDefaults() *StreamGovernanceV2Region`

NewStreamGovernanceV2RegionWithDefaults instantiates a new StreamGovernanceV2Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StreamGovernanceV2Region) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamGovernanceV2Region) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamGovernanceV2Region) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StreamGovernanceV2Region) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StreamGovernanceV2Region) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamGovernanceV2Region) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamGovernanceV2Region) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StreamGovernanceV2Region) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *StreamGovernanceV2Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamGovernanceV2Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamGovernanceV2Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamGovernanceV2Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamGovernanceV2Region) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamGovernanceV2Region) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamGovernanceV2Region) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamGovernanceV2Region) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *StreamGovernanceV2Region) GetSpec() StreamGovernanceV2RegionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StreamGovernanceV2Region) GetSpecOk() (*StreamGovernanceV2RegionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StreamGovernanceV2Region) SetSpec(v StreamGovernanceV2RegionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *StreamGovernanceV2Region) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


