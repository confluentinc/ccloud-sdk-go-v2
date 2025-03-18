# TableflowV1TableflowTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**TableflowV1TableflowTopicSpec**](TableflowV1TableflowTopicSpec.md) |  | [optional] 
**Status** | Pointer to [**TableflowV1TableflowTopicStatus**](TableflowV1TableflowTopicStatus.md) |  | [optional] 

## Methods

### NewTableflowV1TableflowTopic

`func NewTableflowV1TableflowTopic() *TableflowV1TableflowTopic`

NewTableflowV1TableflowTopic instantiates a new TableflowV1TableflowTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicWithDefaults

`func NewTableflowV1TableflowTopicWithDefaults() *TableflowV1TableflowTopic`

NewTableflowV1TableflowTopicWithDefaults instantiates a new TableflowV1TableflowTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TableflowV1TableflowTopic) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TableflowV1TableflowTopic) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TableflowV1TableflowTopic) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TableflowV1TableflowTopic) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TableflowV1TableflowTopic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1TableflowTopic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1TableflowTopic) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TableflowV1TableflowTopic) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *TableflowV1TableflowTopic) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TableflowV1TableflowTopic) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TableflowV1TableflowTopic) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TableflowV1TableflowTopic) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *TableflowV1TableflowTopic) GetSpec() TableflowV1TableflowTopicSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TableflowV1TableflowTopic) GetSpecOk() (*TableflowV1TableflowTopicSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TableflowV1TableflowTopic) SetSpec(v TableflowV1TableflowTopicSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TableflowV1TableflowTopic) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *TableflowV1TableflowTopic) GetStatus() TableflowV1TableflowTopicStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TableflowV1TableflowTopic) GetStatusOk() (*TableflowV1TableflowTopicStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TableflowV1TableflowTopic) SetStatus(v TableflowV1TableflowTopicStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TableflowV1TableflowTopic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


