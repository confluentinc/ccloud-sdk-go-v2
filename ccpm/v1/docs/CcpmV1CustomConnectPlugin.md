# CcpmV1CustomConnectPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**CcpmV1CustomConnectPluginSpec**](CcpmV1CustomConnectPluginSpec.md) |  | [optional] 

## Methods

### NewCcpmV1CustomConnectPlugin

`func NewCcpmV1CustomConnectPlugin() *CcpmV1CustomConnectPlugin`

NewCcpmV1CustomConnectPlugin instantiates a new CcpmV1CustomConnectPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1CustomConnectPluginWithDefaults

`func NewCcpmV1CustomConnectPluginWithDefaults() *CcpmV1CustomConnectPlugin`

NewCcpmV1CustomConnectPluginWithDefaults instantiates a new CcpmV1CustomConnectPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpmV1CustomConnectPlugin) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpmV1CustomConnectPlugin) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpmV1CustomConnectPlugin) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpmV1CustomConnectPlugin) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpmV1CustomConnectPlugin) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpmV1CustomConnectPlugin) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpmV1CustomConnectPlugin) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpmV1CustomConnectPlugin) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CcpmV1CustomConnectPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcpmV1CustomConnectPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcpmV1CustomConnectPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CcpmV1CustomConnectPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CcpmV1CustomConnectPlugin) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpmV1CustomConnectPlugin) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpmV1CustomConnectPlugin) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CcpmV1CustomConnectPlugin) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CcpmV1CustomConnectPlugin) GetSpec() CcpmV1CustomConnectPluginSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CcpmV1CustomConnectPlugin) GetSpecOk() (*CcpmV1CustomConnectPluginSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CcpmV1CustomConnectPlugin) SetSpec(v CcpmV1CustomConnectPluginSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CcpmV1CustomConnectPlugin) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


