# ArmV1AccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource id | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Accesses** | Pointer to [**[]ArmV1Access**](ArmV1Access.md) | The desire accesses, can only support requesting 1 topic. | [optional] 
**Requester** | Pointer to [**ArmV1Requester**](arm.v1.Requester.md) | The user who initiate this request | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the request | [optional] [readonly] 
**Message** | Pointer to **string** | Optional message attached with the request | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this request is created | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The date and time at which this request will expire if not approved or rejected | [optional] [readonly] 

## Methods

### NewArmV1AccessRequest

`func NewArmV1AccessRequest() *ArmV1AccessRequest`

NewArmV1AccessRequest instantiates a new ArmV1AccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmV1AccessRequestWithDefaults

`func NewArmV1AccessRequestWithDefaults() *ArmV1AccessRequest`

NewArmV1AccessRequestWithDefaults instantiates a new ArmV1AccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArmV1AccessRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArmV1AccessRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArmV1AccessRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArmV1AccessRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArmV1AccessRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArmV1AccessRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArmV1AccessRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArmV1AccessRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArmV1AccessRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArmV1AccessRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArmV1AccessRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArmV1AccessRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArmV1AccessRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArmV1AccessRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArmV1AccessRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArmV1AccessRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccesses

`func (o *ArmV1AccessRequest) GetAccesses() []ArmV1Access`

GetAccesses returns the Accesses field if non-nil, zero value otherwise.

### GetAccessesOk

`func (o *ArmV1AccessRequest) GetAccessesOk() (*[]ArmV1Access, bool)`

GetAccessesOk returns a tuple with the Accesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesses

`func (o *ArmV1AccessRequest) SetAccesses(v []ArmV1Access)`

SetAccesses sets Accesses field to given value.

### HasAccesses

`func (o *ArmV1AccessRequest) HasAccesses() bool`

HasAccesses returns a boolean if a field has been set.

### GetRequester

`func (o *ArmV1AccessRequest) GetRequester() ArmV1Requester`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ArmV1AccessRequest) GetRequesterOk() (*ArmV1Requester, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ArmV1AccessRequest) SetRequester(v ArmV1Requester)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ArmV1AccessRequest) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetStatus

`func (o *ArmV1AccessRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArmV1AccessRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArmV1AccessRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArmV1AccessRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ArmV1AccessRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArmV1AccessRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArmV1AccessRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArmV1AccessRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArmV1AccessRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArmV1AccessRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArmV1AccessRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArmV1AccessRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ArmV1AccessRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArmV1AccessRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArmV1AccessRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArmV1AccessRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


