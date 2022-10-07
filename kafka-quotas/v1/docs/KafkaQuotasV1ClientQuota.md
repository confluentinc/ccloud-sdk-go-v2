# KafkaQuotasV1ClientQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**KafkaQuotasV1ClientQuotaSpec**](KafkaQuotasV1ClientQuotaSpec.md) |  | [optional] 

## Methods

### NewKafkaQuotasV1ClientQuota

`func NewKafkaQuotasV1ClientQuota() *KafkaQuotasV1ClientQuota`

NewKafkaQuotasV1ClientQuota instantiates a new KafkaQuotasV1ClientQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaWithDefaults

`func NewKafkaQuotasV1ClientQuotaWithDefaults() *KafkaQuotasV1ClientQuota`

NewKafkaQuotasV1ClientQuotaWithDefaults instantiates a new KafkaQuotasV1ClientQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KafkaQuotasV1ClientQuota) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaQuotasV1ClientQuota) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaQuotasV1ClientQuota) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaQuotasV1ClientQuota) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KafkaQuotasV1ClientQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaQuotasV1ClientQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaQuotasV1ClientQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaQuotasV1ClientQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *KafkaQuotasV1ClientQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaQuotasV1ClientQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaQuotasV1ClientQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KafkaQuotasV1ClientQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *KafkaQuotasV1ClientQuota) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaQuotasV1ClientQuota) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaQuotasV1ClientQuota) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaQuotasV1ClientQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KafkaQuotasV1ClientQuota) GetSpec() KafkaQuotasV1ClientQuotaSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KafkaQuotasV1ClientQuota) GetSpecOk() (*KafkaQuotasV1ClientQuotaSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KafkaQuotasV1ClientQuota) SetSpec(v KafkaQuotasV1ClientQuotaSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *KafkaQuotasV1ClientQuota) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


