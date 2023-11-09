# IamV2IpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**GroupName** | Pointer to **string** | A human readable name for an IP Group. Can contain any unicode letter or number, the ASCII space character, or any of the following special characters: &#x60;[&#x60;, &#x60;]&#x60;, &#x60;|&#x60;, &#x60;&amp;&#x60;, &#x60;+&#x60;, &#x60;-&#x60;, &#x60;_&#x60;, &#x60;/&#x60;, &#x60;.&#x60;, &#x60;,&#x60;.  | [optional] 
**CidrBlocks** | Pointer to **[]string** | A list of CIDRs. | [optional] 

## Methods

### NewIamV2IpGroup

`func NewIamV2IpGroup() *IamV2IpGroup`

NewIamV2IpGroup instantiates a new IamV2IpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IpGroupWithDefaults

`func NewIamV2IpGroupWithDefaults() *IamV2IpGroup`

NewIamV2IpGroupWithDefaults instantiates a new IamV2IpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IpGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IpGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IpGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IpGroup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IpGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IpGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IpGroup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IpGroup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IpGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IpGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IpGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2IpGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2IpGroup) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IpGroup) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IpGroup) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2IpGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetGroupName

`func (o *IamV2IpGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *IamV2IpGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *IamV2IpGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *IamV2IpGroup) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetCidrBlocks

`func (o *IamV2IpGroup) GetCidrBlocks() []string`

GetCidrBlocks returns the CidrBlocks field if non-nil, zero value otherwise.

### GetCidrBlocksOk

`func (o *IamV2IpGroup) GetCidrBlocksOk() (*[]string, bool)`

GetCidrBlocksOk returns a tuple with the CidrBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlocks

`func (o *IamV2IpGroup) SetCidrBlocks(v []string)`

SetCidrBlocks sets CidrBlocks field to given value.

### HasCidrBlocks

`func (o *IamV2IpGroup) HasCidrBlocks() bool`

HasCidrBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


