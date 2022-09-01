# NetworkingV1Peering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingV1PeeringSpec**](NetworkingV1PeeringSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkingV1PeeringStatus**](NetworkingV1PeeringStatus.md) |  | [optional] 

## Methods

### NewNetworkingV1Peering

`func NewNetworkingV1Peering() *NetworkingV1Peering`

NewNetworkingV1Peering instantiates a new NetworkingV1Peering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PeeringWithDefaults

`func NewNetworkingV1PeeringWithDefaults() *NetworkingV1Peering`

NewNetworkingV1PeeringWithDefaults instantiates a new NetworkingV1Peering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1Peering) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1Peering) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1Peering) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1Peering) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1Peering) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1Peering) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1Peering) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1Peering) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1Peering) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1Peering) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1Peering) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingV1Peering) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingV1Peering) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1Peering) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1Peering) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingV1Peering) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingV1Peering) GetSpec() NetworkingV1PeeringSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1Peering) GetSpecOk() (*NetworkingV1PeeringSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1Peering) SetSpec(v NetworkingV1PeeringSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingV1Peering) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingV1Peering) GetStatus() NetworkingV1PeeringStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1Peering) GetStatusOk() (*NetworkingV1PeeringStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1Peering) SetStatus(v NetworkingV1PeeringStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingV1Peering) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


