# IamV2IpFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ResourceGroup** | Pointer to **string** | Container for IP Group. The only resource_group currently available is \&quot;management\&quot;. | [optional] 
**IpGroup** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | An IP Group. | [optional] 

## Methods

### NewIamV2IpFilter

`func NewIamV2IpFilter() *IamV2IpFilter`

NewIamV2IpFilter instantiates a new IamV2IpFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IpFilterWithDefaults

`func NewIamV2IpFilterWithDefaults() *IamV2IpFilter`

NewIamV2IpFilterWithDefaults instantiates a new IamV2IpFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IpFilter) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IpFilter) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IpFilter) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IpFilter) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IpFilter) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IpFilter) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IpFilter) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IpFilter) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IpFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IpFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IpFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2IpFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2IpFilter) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IpFilter) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IpFilter) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2IpFilter) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceGroup

`func (o *IamV2IpFilter) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *IamV2IpFilter) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *IamV2IpFilter) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *IamV2IpFilter) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetIpGroup

`func (o *IamV2IpFilter) GetIpGroup() GlobalObjectReference`

GetIpGroup returns the IpGroup field if non-nil, zero value otherwise.

### GetIpGroupOk

`func (o *IamV2IpFilter) GetIpGroupOk() (*GlobalObjectReference, bool)`

GetIpGroupOk returns a tuple with the IpGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroup

`func (o *IamV2IpFilter) SetIpGroup(v GlobalObjectReference)`

SetIpGroup sets IpGroup field to given value.

### HasIpGroup

`func (o *IamV2IpFilter) HasIpGroup() bool`

HasIpGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


