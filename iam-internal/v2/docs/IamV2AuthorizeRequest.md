# IamV2AuthorizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**TargetPrincipalContext** | Pointer to [**IamV2AuthorizeRequestTargetPrincipalContext**](IamV2AuthorizeRequestTargetPrincipalContext.md) |  | [optional] 
**RequestContext** | Pointer to [**IamV2AuthorizeRequestRequestContext**](IamV2AuthorizeRequestRequestContext.md) |  | [optional] 
**Actions** | Pointer to [**[]IamV2AuthorizeRequestActions**](IamV2AuthorizeRequestActions.md) |  | [optional] 

## Methods

### NewIamV2AuthorizeRequest

`func NewIamV2AuthorizeRequest() *IamV2AuthorizeRequest`

NewIamV2AuthorizeRequest instantiates a new IamV2AuthorizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2AuthorizeRequestWithDefaults

`func NewIamV2AuthorizeRequestWithDefaults() *IamV2AuthorizeRequest`

NewIamV2AuthorizeRequestWithDefaults instantiates a new IamV2AuthorizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2AuthorizeRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2AuthorizeRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2AuthorizeRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2AuthorizeRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2AuthorizeRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2AuthorizeRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2AuthorizeRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2AuthorizeRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2AuthorizeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2AuthorizeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2AuthorizeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamV2AuthorizeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamV2AuthorizeRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2AuthorizeRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2AuthorizeRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamV2AuthorizeRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTargetPrincipalContext

`func (o *IamV2AuthorizeRequest) GetTargetPrincipalContext() IamV2AuthorizeRequestTargetPrincipalContext`

GetTargetPrincipalContext returns the TargetPrincipalContext field if non-nil, zero value otherwise.

### GetTargetPrincipalContextOk

`func (o *IamV2AuthorizeRequest) GetTargetPrincipalContextOk() (*IamV2AuthorizeRequestTargetPrincipalContext, bool)`

GetTargetPrincipalContextOk returns a tuple with the TargetPrincipalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipalContext

`func (o *IamV2AuthorizeRequest) SetTargetPrincipalContext(v IamV2AuthorizeRequestTargetPrincipalContext)`

SetTargetPrincipalContext sets TargetPrincipalContext field to given value.

### HasTargetPrincipalContext

`func (o *IamV2AuthorizeRequest) HasTargetPrincipalContext() bool`

HasTargetPrincipalContext returns a boolean if a field has been set.

### GetRequestContext

`func (o *IamV2AuthorizeRequest) GetRequestContext() IamV2AuthorizeRequestRequestContext`

GetRequestContext returns the RequestContext field if non-nil, zero value otherwise.

### GetRequestContextOk

`func (o *IamV2AuthorizeRequest) GetRequestContextOk() (*IamV2AuthorizeRequestRequestContext, bool)`

GetRequestContextOk returns a tuple with the RequestContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContext

`func (o *IamV2AuthorizeRequest) SetRequestContext(v IamV2AuthorizeRequestRequestContext)`

SetRequestContext sets RequestContext field to given value.

### HasRequestContext

`func (o *IamV2AuthorizeRequest) HasRequestContext() bool`

HasRequestContext returns a boolean if a field has been set.

### GetActions

`func (o *IamV2AuthorizeRequest) GetActions() []IamV2AuthorizeRequestActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IamV2AuthorizeRequest) GetActionsOk() (*[]IamV2AuthorizeRequestActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IamV2AuthorizeRequest) SetActions(v []IamV2AuthorizeRequestActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IamV2AuthorizeRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


