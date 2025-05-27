# CcpmV1CustomConnectPluginVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**CcpmV1CustomConnectPluginVersionSpec**](CcpmV1CustomConnectPluginVersionSpec.md) |  | [optional] 
**Status** | Pointer to [**CcpmV1CustomConnectPluginVersionStatus**](CcpmV1CustomConnectPluginVersionStatus.md) |  | [optional] 

## Methods

### NewCcpmV1CustomConnectPluginVersion

`func NewCcpmV1CustomConnectPluginVersion() *CcpmV1CustomConnectPluginVersion`

NewCcpmV1CustomConnectPluginVersion instantiates a new CcpmV1CustomConnectPluginVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginVersionWithDefaults

`func NewCcpmV1CustomConnectPluginVersionWithDefaults() *CcpmV1CustomConnectPluginVersion`

NewCcpmV1CustomConnectPluginVersionWithDefaults instantiates a new CcpmV1CustomConnectPluginVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpmV1CustomConnectPluginVersion) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpmV1CustomConnectPluginVersion) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpmV1CustomConnectPluginVersion) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpmV1CustomConnectPluginVersion) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpmV1CustomConnectPluginVersion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpmV1CustomConnectPluginVersion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpmV1CustomConnectPluginVersion) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpmV1CustomConnectPluginVersion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CcpmV1CustomConnectPluginVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcpmV1CustomConnectPluginVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcpmV1CustomConnectPluginVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CcpmV1CustomConnectPluginVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CcpmV1CustomConnectPluginVersion) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpmV1CustomConnectPluginVersion) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpmV1CustomConnectPluginVersion) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CcpmV1CustomConnectPluginVersion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CcpmV1CustomConnectPluginVersion) GetSpec() CcpmV1CustomConnectPluginVersionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CcpmV1CustomConnectPluginVersion) GetSpecOk() (*CcpmV1CustomConnectPluginVersionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CcpmV1CustomConnectPluginVersion) SetSpec(v CcpmV1CustomConnectPluginVersionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CcpmV1CustomConnectPluginVersion) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CcpmV1CustomConnectPluginVersion) GetStatus() CcpmV1CustomConnectPluginVersionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CcpmV1CustomConnectPluginVersion) GetStatusOk() (*CcpmV1CustomConnectPluginVersionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CcpmV1CustomConnectPluginVersion) SetStatus(v CcpmV1CustomConnectPluginVersionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CcpmV1CustomConnectPluginVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


