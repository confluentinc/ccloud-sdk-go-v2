# SwitchoverV1SwitchoverEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SwitchoverV1SwitchoverEndpointSpec**](SwitchoverV1SwitchoverEndpointSpec.md) |  | [optional] 
**Status** | Pointer to [**SwitchoverV1SwitchoverEndpointStatus**](SwitchoverV1SwitchoverEndpointStatus.md) |  | [optional] 

## Methods

### NewSwitchoverV1SwitchoverEndpoint

`func NewSwitchoverV1SwitchoverEndpoint() *SwitchoverV1SwitchoverEndpoint`

NewSwitchoverV1SwitchoverEndpoint instantiates a new SwitchoverV1SwitchoverEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverEndpointWithDefaults

`func NewSwitchoverV1SwitchoverEndpointWithDefaults() *SwitchoverV1SwitchoverEndpoint`

NewSwitchoverV1SwitchoverEndpointWithDefaults instantiates a new SwitchoverV1SwitchoverEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SwitchoverV1SwitchoverEndpoint) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SwitchoverV1SwitchoverEndpoint) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SwitchoverV1SwitchoverEndpoint) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SwitchoverV1SwitchoverEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SwitchoverV1SwitchoverEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SwitchoverV1SwitchoverEndpoint) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SwitchoverV1SwitchoverEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwitchoverV1SwitchoverEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SwitchoverV1SwitchoverEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SwitchoverV1SwitchoverEndpoint) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SwitchoverV1SwitchoverEndpoint) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SwitchoverV1SwitchoverEndpoint) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SwitchoverV1SwitchoverEndpoint) GetSpec() SwitchoverV1SwitchoverEndpointSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetSpecOk() (*SwitchoverV1SwitchoverEndpointSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SwitchoverV1SwitchoverEndpoint) SetSpec(v SwitchoverV1SwitchoverEndpointSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SwitchoverV1SwitchoverEndpoint) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SwitchoverV1SwitchoverEndpoint) GetStatus() SwitchoverV1SwitchoverEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchoverV1SwitchoverEndpoint) GetStatusOk() (*SwitchoverV1SwitchoverEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchoverV1SwitchoverEndpoint) SetStatus(v SwitchoverV1SwitchoverEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwitchoverV1SwitchoverEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


