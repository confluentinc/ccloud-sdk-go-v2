# IamV2IpFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**FilterName** | Pointer to **string** | A human readable name for an IP Filter. Can contain any unicode letter or number, the ASCII space character, or any of the following special characters: &#x60;[&#x60;, &#x60;]&#x60;, &#x60;|&#x60;, &#x60;&amp;&#x60;, &#x60;+&#x60;, &#x60;-&#x60;, &#x60;_&#x60;, &#x60;/&#x60;, &#x60;.&#x60;, &#x60;,&#x60;.  | [optional] 
**ResourceGroup** | Pointer to **string** | Scope of resources covered by this IP filter. The only resource_group currently available is \&quot;management\&quot;.  | [optional] 
**OperationGroups** | Pointer to **[]string** | Scope of resources covered by this IP filter.  | [optional] 
**IpGroups** | Pointer to [**[]GlobalObjectReference**](GlobalObjectReference.md) | A list of IP Groups. | [optional] 

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

### GetFilterName

`func (o *IamV2IpFilter) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *IamV2IpFilter) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *IamV2IpFilter) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *IamV2IpFilter) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

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

### GetOperationGroups

`func (o *IamV2IpFilter) GetOperationGroups() []string`

GetOperationGroups returns the OperationGroups field if non-nil, zero value otherwise.

### GetOperationGroupsOk

`func (o *IamV2IpFilter) GetOperationGroupsOk() (*[]string, bool)`

GetOperationGroupsOk returns a tuple with the OperationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationGroups

`func (o *IamV2IpFilter) SetOperationGroups(v []string)`

SetOperationGroups sets OperationGroups field to given value.

### HasOperationGroups

`func (o *IamV2IpFilter) HasOperationGroups() bool`

HasOperationGroups returns a boolean if a field has been set.

### GetIpGroups

`func (o *IamV2IpFilter) GetIpGroups() []GlobalObjectReference`

GetIpGroups returns the IpGroups field if non-nil, zero value otherwise.

### GetIpGroupsOk

`func (o *IamV2IpFilter) GetIpGroupsOk() (*[]GlobalObjectReference, bool)`

GetIpGroupsOk returns a tuple with the IpGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroups

`func (o *IamV2IpFilter) SetIpGroups(v []GlobalObjectReference)`

SetIpGroups sets IpGroups field to given value.

### HasIpGroups

`func (o *IamV2IpFilter) HasIpGroups() bool`

HasIpGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


