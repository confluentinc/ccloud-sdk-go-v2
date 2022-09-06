# CdxV1CreateProviderShareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DeliveryMethod** | Pointer to **string** | Method by which the invite will be delivered | [optional] 
**ConsumerRestriction** | Pointer to [**CdxV1CreateProviderShareRequestConsumerRestrictionOneOf**](CdxV1CreateProviderShareRequestConsumerRestrictionOneOf.md) | Restrictions on the consumer that can redeem this token | [optional] 
**Resources** | Pointer to **[]string** | List of resource crns to be shared | [optional] 

## Methods

### NewCdxV1CreateProviderShareRequest

`func NewCdxV1CreateProviderShareRequest() *CdxV1CreateProviderShareRequest`

NewCdxV1CreateProviderShareRequest instantiates a new CdxV1CreateProviderShareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1CreateProviderShareRequestWithDefaults

`func NewCdxV1CreateProviderShareRequestWithDefaults() *CdxV1CreateProviderShareRequest`

NewCdxV1CreateProviderShareRequestWithDefaults instantiates a new CdxV1CreateProviderShareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1CreateProviderShareRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1CreateProviderShareRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1CreateProviderShareRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1CreateProviderShareRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1CreateProviderShareRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1CreateProviderShareRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1CreateProviderShareRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1CreateProviderShareRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1CreateProviderShareRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1CreateProviderShareRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1CreateProviderShareRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1CreateProviderShareRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1CreateProviderShareRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1CreateProviderShareRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1CreateProviderShareRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1CreateProviderShareRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *CdxV1CreateProviderShareRequest) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *CdxV1CreateProviderShareRequest) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *CdxV1CreateProviderShareRequest) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *CdxV1CreateProviderShareRequest) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetConsumerRestriction

`func (o *CdxV1CreateProviderShareRequest) GetConsumerRestriction() CdxV1CreateProviderShareRequestConsumerRestrictionOneOf`

GetConsumerRestriction returns the ConsumerRestriction field if non-nil, zero value otherwise.

### GetConsumerRestrictionOk

`func (o *CdxV1CreateProviderShareRequest) GetConsumerRestrictionOk() (*CdxV1CreateProviderShareRequestConsumerRestrictionOneOf, bool)`

GetConsumerRestrictionOk returns a tuple with the ConsumerRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRestriction

`func (o *CdxV1CreateProviderShareRequest) SetConsumerRestriction(v CdxV1CreateProviderShareRequestConsumerRestrictionOneOf)`

SetConsumerRestriction sets ConsumerRestriction field to given value.

### HasConsumerRestriction

`func (o *CdxV1CreateProviderShareRequest) HasConsumerRestriction() bool`

HasConsumerRestriction returns a boolean if a field has been set.

### GetResources

`func (o *CdxV1CreateProviderShareRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1CreateProviderShareRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1CreateProviderShareRequest) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1CreateProviderShareRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


