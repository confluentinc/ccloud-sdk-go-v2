# NetworkingAdminV1PrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkingAdminV1PrivateLinkAccessSpec**](NetworkingAdminV1PrivateLinkAccessSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkingAdminV1PrivateLinkAccessStatus**](NetworkingAdminV1PrivateLinkAccessStatus.md) |  | [optional] 

## Methods

### NewNetworkingAdminV1PrivateLinkAccess

`func NewNetworkingAdminV1PrivateLinkAccess() *NetworkingAdminV1PrivateLinkAccess`

NewNetworkingAdminV1PrivateLinkAccess instantiates a new NetworkingAdminV1PrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1PrivateLinkAccessWithDefaults

`func NewNetworkingAdminV1PrivateLinkAccessWithDefaults() *NetworkingAdminV1PrivateLinkAccess`

NewNetworkingAdminV1PrivateLinkAccessWithDefaults instantiates a new NetworkingAdminV1PrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkingAdminV1PrivateLinkAccess) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkingAdminV1PrivateLinkAccess) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkingAdminV1PrivateLinkAccess) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkingAdminV1PrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1PrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkingAdminV1PrivateLinkAccess) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NetworkingAdminV1PrivateLinkAccess) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkingAdminV1PrivateLinkAccess) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkingAdminV1PrivateLinkAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkingAdminV1PrivateLinkAccess) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkingAdminV1PrivateLinkAccess) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkingAdminV1PrivateLinkAccess) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkingAdminV1PrivateLinkAccess) GetSpec() NetworkingAdminV1PrivateLinkAccessSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetSpecOk() (*NetworkingAdminV1PrivateLinkAccessSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkingAdminV1PrivateLinkAccess) SetSpec(v NetworkingAdminV1PrivateLinkAccessSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkingAdminV1PrivateLinkAccess) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkingAdminV1PrivateLinkAccess) GetStatus() NetworkingAdminV1PrivateLinkAccessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkingAdminV1PrivateLinkAccess) GetStatusOk() (*NetworkingAdminV1PrivateLinkAccessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkingAdminV1PrivateLinkAccess) SetStatus(v NetworkingAdminV1PrivateLinkAccessStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkingAdminV1PrivateLinkAccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


