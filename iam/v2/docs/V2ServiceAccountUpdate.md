# V2ServiceAccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Description** | Pointer to **string** | A free-form description of the Service Account | [optional] 

## Methods

### NewV2ServiceAccountUpdate

`func NewV2ServiceAccountUpdate() *V2ServiceAccountUpdate`

NewV2ServiceAccountUpdate instantiates a new V2ServiceAccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ServiceAccountUpdateWithDefaults

`func NewV2ServiceAccountUpdateWithDefaults() *V2ServiceAccountUpdate`

NewV2ServiceAccountUpdateWithDefaults instantiates a new V2ServiceAccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V2ServiceAccountUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V2ServiceAccountUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V2ServiceAccountUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V2ServiceAccountUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V2ServiceAccountUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V2ServiceAccountUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V2ServiceAccountUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V2ServiceAccountUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *V2ServiceAccountUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ServiceAccountUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ServiceAccountUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2ServiceAccountUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *V2ServiceAccountUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V2ServiceAccountUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V2ServiceAccountUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V2ServiceAccountUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *V2ServiceAccountUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2ServiceAccountUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2ServiceAccountUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2ServiceAccountUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


