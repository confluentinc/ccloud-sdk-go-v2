# FcpmV2IamBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Region** | Pointer to **string** | Flink compute pools in the region provided will be able to use this identity pool. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the region. | [optional] [readonly] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**IdentityPool** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The identity_pool to which this belongs. | [optional] 

## Methods

### NewFcpmV2IamBinding

`func NewFcpmV2IamBinding() *FcpmV2IamBinding`

NewFcpmV2IamBinding instantiates a new FcpmV2IamBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2IamBindingWithDefaults

`func NewFcpmV2IamBindingWithDefaults() *FcpmV2IamBinding`

NewFcpmV2IamBindingWithDefaults instantiates a new FcpmV2IamBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2IamBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2IamBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2IamBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2IamBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2IamBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2IamBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2IamBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2IamBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2IamBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2IamBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2IamBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2IamBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2IamBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2IamBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2IamBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2IamBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *FcpmV2IamBinding) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FcpmV2IamBinding) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FcpmV2IamBinding) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FcpmV2IamBinding) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCloud

`func (o *FcpmV2IamBinding) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FcpmV2IamBinding) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FcpmV2IamBinding) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FcpmV2IamBinding) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *FcpmV2IamBinding) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FcpmV2IamBinding) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FcpmV2IamBinding) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FcpmV2IamBinding) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIdentityPool

`func (o *FcpmV2IamBinding) GetIdentityPool() GlobalObjectReference`

GetIdentityPool returns the IdentityPool field if non-nil, zero value otherwise.

### GetIdentityPoolOk

`func (o *FcpmV2IamBinding) GetIdentityPoolOk() (*GlobalObjectReference, bool)`

GetIdentityPoolOk returns a tuple with the IdentityPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPool

`func (o *FcpmV2IamBinding) SetIdentityPool(v GlobalObjectReference)`

SetIdentityPool sets IdentityPool field to given value.

### HasIdentityPool

`func (o *FcpmV2IamBinding) HasIdentityPool() bool`

HasIdentityPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


