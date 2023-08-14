# ArmV1UpdateAccessRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**AccessRequest** | Pointer to [**ArmV1AccessRequest**](arm.v1.AccessRequest.md) | The updated AccessRequest | [optional] 
**Error** | Pointer to [**ArmV1Error**](arm.v1.Error.md) | The optional error if the operation fails | [optional] 

## Methods

### NewArmV1UpdateAccessRequestResponse

`func NewArmV1UpdateAccessRequestResponse() *ArmV1UpdateAccessRequestResponse`

NewArmV1UpdateAccessRequestResponse instantiates a new ArmV1UpdateAccessRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmV1UpdateAccessRequestResponseWithDefaults

`func NewArmV1UpdateAccessRequestResponseWithDefaults() *ArmV1UpdateAccessRequestResponse`

NewArmV1UpdateAccessRequestResponseWithDefaults instantiates a new ArmV1UpdateAccessRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArmV1UpdateAccessRequestResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArmV1UpdateAccessRequestResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArmV1UpdateAccessRequestResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArmV1UpdateAccessRequestResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArmV1UpdateAccessRequestResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArmV1UpdateAccessRequestResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArmV1UpdateAccessRequestResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArmV1UpdateAccessRequestResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArmV1UpdateAccessRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArmV1UpdateAccessRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArmV1UpdateAccessRequestResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArmV1UpdateAccessRequestResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArmV1UpdateAccessRequestResponse) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArmV1UpdateAccessRequestResponse) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArmV1UpdateAccessRequestResponse) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArmV1UpdateAccessRequestResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccessRequest

`func (o *ArmV1UpdateAccessRequestResponse) GetAccessRequest() ArmV1AccessRequest`

GetAccessRequest returns the AccessRequest field if non-nil, zero value otherwise.

### GetAccessRequestOk

`func (o *ArmV1UpdateAccessRequestResponse) GetAccessRequestOk() (*ArmV1AccessRequest, bool)`

GetAccessRequestOk returns a tuple with the AccessRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequest

`func (o *ArmV1UpdateAccessRequestResponse) SetAccessRequest(v ArmV1AccessRequest)`

SetAccessRequest sets AccessRequest field to given value.

### HasAccessRequest

`func (o *ArmV1UpdateAccessRequestResponse) HasAccessRequest() bool`

HasAccessRequest returns a boolean if a field has been set.

### GetError

`func (o *ArmV1UpdateAccessRequestResponse) GetError() ArmV1Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ArmV1UpdateAccessRequestResponse) GetErrorOk() (*ArmV1Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ArmV1UpdateAccessRequestResponse) SetError(v ArmV1Error)`

SetError sets Error field to given value.

### HasError

`func (o *ArmV1UpdateAccessRequestResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


