# ArmV1UpdateAccessRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**AccessRequests** | Pointer to [**[]ArmV1AccessRequestUpdateOperation**](ArmV1AccessRequestUpdateOperation.md) | List of access requests | [optional] 

## Methods

### NewArmV1UpdateAccessRequestRequest

`func NewArmV1UpdateAccessRequestRequest() *ArmV1UpdateAccessRequestRequest`

NewArmV1UpdateAccessRequestRequest instantiates a new ArmV1UpdateAccessRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmV1UpdateAccessRequestRequestWithDefaults

`func NewArmV1UpdateAccessRequestRequestWithDefaults() *ArmV1UpdateAccessRequestRequest`

NewArmV1UpdateAccessRequestRequestWithDefaults instantiates a new ArmV1UpdateAccessRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArmV1UpdateAccessRequestRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArmV1UpdateAccessRequestRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArmV1UpdateAccessRequestRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArmV1UpdateAccessRequestRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArmV1UpdateAccessRequestRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArmV1UpdateAccessRequestRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArmV1UpdateAccessRequestRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArmV1UpdateAccessRequestRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArmV1UpdateAccessRequestRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArmV1UpdateAccessRequestRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArmV1UpdateAccessRequestRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArmV1UpdateAccessRequestRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArmV1UpdateAccessRequestRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArmV1UpdateAccessRequestRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArmV1UpdateAccessRequestRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArmV1UpdateAccessRequestRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccessRequests

`func (o *ArmV1UpdateAccessRequestRequest) GetAccessRequests() []ArmV1AccessRequestUpdateOperation`

GetAccessRequests returns the AccessRequests field if non-nil, zero value otherwise.

### GetAccessRequestsOk

`func (o *ArmV1UpdateAccessRequestRequest) GetAccessRequestsOk() (*[]ArmV1AccessRequestUpdateOperation, bool)`

GetAccessRequestsOk returns a tuple with the AccessRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequests

`func (o *ArmV1UpdateAccessRequestRequest) SetAccessRequests(v []ArmV1AccessRequestUpdateOperation)`

SetAccessRequests sets AccessRequests field to given value.

### HasAccessRequests

`func (o *ArmV1UpdateAccessRequestRequest) HasAccessRequests() bool`

HasAccessRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


