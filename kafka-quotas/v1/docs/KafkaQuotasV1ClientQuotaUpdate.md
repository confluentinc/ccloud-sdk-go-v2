# KafkaQuotasV1ClientQuotaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the client quota. | [optional] 
**Description** | Pointer to **string** | A human readable description for the client quota. | [optional] 
**Throughput** | Pointer to [**KafkaQuotasV1Throughput**](kafka-quotas.v1.Throughput.md) | Throughput for the client quota. | [optional] 
**Principals** | Pointer to [**[]ObjectReference**](ObjectReference.md) | A list of service accounts. Special name \&quot;default\&quot; can be used to represent the default quota for all users and service accounts.  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewKafkaQuotasV1ClientQuotaUpdate

`func NewKafkaQuotasV1ClientQuotaUpdate() *KafkaQuotasV1ClientQuotaUpdate`

NewKafkaQuotasV1ClientQuotaUpdate instantiates a new KafkaQuotasV1ClientQuotaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaUpdateWithDefaults

`func NewKafkaQuotasV1ClientQuotaUpdateWithDefaults() *KafkaQuotasV1ClientQuotaUpdate`

NewKafkaQuotasV1ClientQuotaUpdateWithDefaults instantiates a new KafkaQuotasV1ClientQuotaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetThroughput

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetThroughput() KafkaQuotasV1Throughput`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetThroughputOk() (*KafkaQuotasV1Throughput, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetThroughput(v KafkaQuotasV1Throughput)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetPrincipals

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetPrincipals() []ObjectReference`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetPrincipalsOk() (*[]ObjectReference, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetPrincipals(v []ObjectReference)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetEnvironment

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KafkaQuotasV1ClientQuotaUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KafkaQuotasV1ClientQuotaUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KafkaQuotasV1ClientQuotaUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


