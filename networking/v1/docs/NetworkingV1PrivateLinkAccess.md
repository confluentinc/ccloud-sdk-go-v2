# NetworkingV1PrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingV1PrivateLinkAccessSpec**](NetworkingV1PrivateLinkAccessSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkingV1PrivateLinkAccessStatus**](NetworkingV1PrivateLinkAccessStatus.md) |  | [optional] 

## Methods

### NewNetworkingV1PrivateLinkAccess

`func NewNetworkingV1PrivateLinkAccess() *NetworkingV1PrivateLinkAccess`

NewNetworkingV1PrivateLinkAccess instantiates a new NetworkingV1PrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAccessWithDefaults

`func NewNetworkingV1PrivateLinkAccessWithDefaults() *NetworkingV1PrivateLinkAccess`

NewNetworkingV1PrivateLinkAccessWithDefaults instantiates a new NetworkingV1PrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingV1PrivateLinkAccess) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingV1PrivateLinkAccess) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingV1PrivateLinkAccess) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingV1PrivateLinkAccess) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingV1PrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingV1PrivateLinkAccess) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingV1PrivateLinkAccess) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingV1PrivateLinkAccess) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingV1PrivateLinkAccess) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingV1PrivateLinkAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingV1PrivateLinkAccess) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingV1PrivateLinkAccess) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingV1PrivateLinkAccess) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingV1PrivateLinkAccess) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingV1PrivateLinkAccess) GetSpec() NetworkingV1PrivateLinkAccessSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingV1PrivateLinkAccess) GetSpecOk() (*NetworkingV1PrivateLinkAccessSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingV1PrivateLinkAccess) SetSpec(v NetworkingV1PrivateLinkAccessSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingV1PrivateLinkAccess) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingV1PrivateLinkAccess) GetStatus() NetworkingV1PrivateLinkAccessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingV1PrivateLinkAccess) GetStatusOk() (*NetworkingV1PrivateLinkAccessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingV1PrivateLinkAccess) SetStatus(v NetworkingV1PrivateLinkAccessStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingV1PrivateLinkAccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


