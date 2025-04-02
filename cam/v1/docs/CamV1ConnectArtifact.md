# CamV1ConnectArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**CamV1ConnectArtifactSpec**](cam.v1.ConnectArtifactSpec.md) |  | [optional] 
**Status** | Pointer to [**CamV1ConnectArtifactStatus**](CamV1ConnectArtifactStatus.md) |  | [optional] 

## Methods

### NewCamV1ConnectArtifact

`func NewCamV1ConnectArtifact() *CamV1ConnectArtifact`

NewCamV1ConnectArtifact instantiates a new CamV1ConnectArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCamV1ConnectArtifactWithDefaults

`func NewCamV1ConnectArtifactWithDefaults() *CamV1ConnectArtifact`

NewCamV1ConnectArtifactWithDefaults instantiates a new CamV1ConnectArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CamV1ConnectArtifact) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CamV1ConnectArtifact) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CamV1ConnectArtifact) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CamV1ConnectArtifact) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CamV1ConnectArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CamV1ConnectArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CamV1ConnectArtifact) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CamV1ConnectArtifact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CamV1ConnectArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CamV1ConnectArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CamV1ConnectArtifact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CamV1ConnectArtifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CamV1ConnectArtifact) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CamV1ConnectArtifact) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CamV1ConnectArtifact) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CamV1ConnectArtifact) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CamV1ConnectArtifact) GetSpec() CamV1ConnectArtifactSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CamV1ConnectArtifact) GetSpecOk() (*CamV1ConnectArtifactSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CamV1ConnectArtifact) SetSpec(v CamV1ConnectArtifactSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CamV1ConnectArtifact) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CamV1ConnectArtifact) GetStatus() CamV1ConnectArtifactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CamV1ConnectArtifact) GetStatusOk() (*CamV1ConnectArtifactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CamV1ConnectArtifact) SetStatus(v CamV1ConnectArtifactStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CamV1ConnectArtifact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


