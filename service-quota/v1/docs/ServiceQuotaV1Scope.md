# ServiceQuotaV1Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Description** | Pointer to **string** | the quota scope for listing quotas queries | [optional] 

## Methods

### NewServiceQuotaV1Scope

`func NewServiceQuotaV1Scope() *ServiceQuotaV1Scope`

NewServiceQuotaV1Scope instantiates a new ServiceQuotaV1Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQuotaV1ScopeWithDefaults

`func NewServiceQuotaV1ScopeWithDefaults() *ServiceQuotaV1Scope`

NewServiceQuotaV1ScopeWithDefaults instantiates a new ServiceQuotaV1Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ServiceQuotaV1Scope) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ServiceQuotaV1Scope) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ServiceQuotaV1Scope) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ServiceQuotaV1Scope) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ServiceQuotaV1Scope) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceQuotaV1Scope) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceQuotaV1Scope) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServiceQuotaV1Scope) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ServiceQuotaV1Scope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQuotaV1Scope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQuotaV1Scope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceQuotaV1Scope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceQuotaV1Scope) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceQuotaV1Scope) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceQuotaV1Scope) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceQuotaV1Scope) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceQuotaV1Scope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceQuotaV1Scope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceQuotaV1Scope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceQuotaV1Scope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


