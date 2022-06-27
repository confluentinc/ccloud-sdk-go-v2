# CdxV1CreateShareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Environment** | Pointer to **string** | Id of ccloud environment | [optional] 
**KafkaCluster** | Pointer to **string** | Id of the Kafka cluster | [optional] 
**DeliveryMethod** | Pointer to **string** | Method by which the invite will be delivered | [optional] 
**ConsumerRestriction** | Pointer to [**CdxV1CreateShareRequestConsumerRestrictionOneOf**](CdxV1CreateShareRequestConsumerRestrictionOneOf.md) | Restrictions on the consumer that can redeem this token | [optional] 
**Resources** | Pointer to **[]string** | List of resource crns to be shared | [optional] 

## Methods

### NewCdxV1CreateShareRequest

`func NewCdxV1CreateShareRequest() *CdxV1CreateShareRequest`

NewCdxV1CreateShareRequest instantiates a new CdxV1CreateShareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1CreateShareRequestWithDefaults

`func NewCdxV1CreateShareRequestWithDefaults() *CdxV1CreateShareRequest`

NewCdxV1CreateShareRequestWithDefaults instantiates a new CdxV1CreateShareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1CreateShareRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1CreateShareRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1CreateShareRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1CreateShareRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1CreateShareRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1CreateShareRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1CreateShareRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1CreateShareRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1CreateShareRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1CreateShareRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1CreateShareRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1CreateShareRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1CreateShareRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1CreateShareRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1CreateShareRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1CreateShareRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEnvironment

`func (o *CdxV1CreateShareRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CdxV1CreateShareRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CdxV1CreateShareRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CdxV1CreateShareRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKafkaCluster

`func (o *CdxV1CreateShareRequest) GetKafkaCluster() string`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *CdxV1CreateShareRequest) GetKafkaClusterOk() (*string, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *CdxV1CreateShareRequest) SetKafkaCluster(v string)`

SetKafkaCluster sets KafkaCluster field to given value.

### HasKafkaCluster

`func (o *CdxV1CreateShareRequest) HasKafkaCluster() bool`

HasKafkaCluster returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *CdxV1CreateShareRequest) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *CdxV1CreateShareRequest) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *CdxV1CreateShareRequest) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *CdxV1CreateShareRequest) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetConsumerRestriction

`func (o *CdxV1CreateShareRequest) GetConsumerRestriction() CdxV1CreateShareRequestConsumerRestrictionOneOf`

GetConsumerRestriction returns the ConsumerRestriction field if non-nil, zero value otherwise.

### GetConsumerRestrictionOk

`func (o *CdxV1CreateShareRequest) GetConsumerRestrictionOk() (*CdxV1CreateShareRequestConsumerRestrictionOneOf, bool)`

GetConsumerRestrictionOk returns a tuple with the ConsumerRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRestriction

`func (o *CdxV1CreateShareRequest) SetConsumerRestriction(v CdxV1CreateShareRequestConsumerRestrictionOneOf)`

SetConsumerRestriction sets ConsumerRestriction field to given value.

### HasConsumerRestriction

`func (o *CdxV1CreateShareRequest) HasConsumerRestriction() bool`

HasConsumerRestriction returns a boolean if a field has been set.

### GetResources

`func (o *CdxV1CreateShareRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CdxV1CreateShareRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CdxV1CreateShareRequest) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CdxV1CreateShareRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


