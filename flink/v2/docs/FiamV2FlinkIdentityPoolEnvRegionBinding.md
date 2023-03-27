# FiamV2FlinkIdentityPoolEnvRegionBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Region** | Pointer to **string** | Flink compute pools in the region provider will be able to use this identity pool. | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider that hosts the region. | [optional] [readonly] 
**IdentityPool** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The identity_pool to which this belongs. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewFiamV2FlinkIdentityPoolEnvRegionBinding

`func NewFiamV2FlinkIdentityPoolEnvRegionBinding() *FiamV2FlinkIdentityPoolEnvRegionBinding`

NewFiamV2FlinkIdentityPoolEnvRegionBinding instantiates a new FiamV2FlinkIdentityPoolEnvRegionBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiamV2FlinkIdentityPoolEnvRegionBindingWithDefaults

`func NewFiamV2FlinkIdentityPoolEnvRegionBindingWithDefaults() *FiamV2FlinkIdentityPoolEnvRegionBinding`

NewFiamV2FlinkIdentityPoolEnvRegionBindingWithDefaults instantiates a new FiamV2FlinkIdentityPoolEnvRegionBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCloud

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetIdentityPool

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetIdentityPool() GlobalObjectReference`

GetIdentityPool returns the IdentityPool field if non-nil, zero value otherwise.

### GetIdentityPoolOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetIdentityPoolOk() (*GlobalObjectReference, bool)`

GetIdentityPoolOk returns a tuple with the IdentityPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPool

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetIdentityPool(v GlobalObjectReference)`

SetIdentityPool sets IdentityPool field to given value.

### HasIdentityPool

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasIdentityPool() bool`

HasIdentityPool returns a boolean if a field has been set.

### GetEnvironment

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FiamV2FlinkIdentityPoolEnvRegionBinding) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


