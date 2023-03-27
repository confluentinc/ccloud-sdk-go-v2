# FcpmV2IdentityPoolEnvRegionBinding

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
**IdentityPool** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The identity_pool to which this belongs. | [optional] 

## Methods

### NewFcpmV2IdentityPoolEnvRegionBinding

`func NewFcpmV2IdentityPoolEnvRegionBinding() *FcpmV2IdentityPoolEnvRegionBinding`

NewFcpmV2IdentityPoolEnvRegionBinding instantiates a new FcpmV2IdentityPoolEnvRegionBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2IdentityPoolEnvRegionBindingWithDefaults

`func NewFcpmV2IdentityPoolEnvRegionBindingWithDefaults() *FcpmV2IdentityPoolEnvRegionBinding`

NewFcpmV2IdentityPoolEnvRegionBindingWithDefaults instantiates a new FcpmV2IdentityPoolEnvRegionBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCloud

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIdentityPool

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetIdentityPool() EnvScopedObjectReference`

GetIdentityPool returns the IdentityPool field if non-nil, zero value otherwise.

### GetIdentityPoolOk

`func (o *FcpmV2IdentityPoolEnvRegionBinding) GetIdentityPoolOk() (*EnvScopedObjectReference, bool)`

GetIdentityPoolOk returns a tuple with the IdentityPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPool

`func (o *FcpmV2IdentityPoolEnvRegionBinding) SetIdentityPool(v EnvScopedObjectReference)`

SetIdentityPool sets IdentityPool field to given value.

### HasIdentityPool

`func (o *FcpmV2IdentityPoolEnvRegionBinding) HasIdentityPool() bool`

HasIdentityPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


