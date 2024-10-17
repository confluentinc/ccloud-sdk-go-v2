# CclV1CustomCodeLoggingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DestinationSettings** | Pointer to [**CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf**](CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf.md) | Destination Settings of the Custom Code Logging. | [optional] 

## Methods

### NewCclV1CustomCodeLoggingUpdate

`func NewCclV1CustomCodeLoggingUpdate() *CclV1CustomCodeLoggingUpdate`

NewCclV1CustomCodeLoggingUpdate instantiates a new CclV1CustomCodeLoggingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCclV1CustomCodeLoggingUpdateWithDefaults

`func NewCclV1CustomCodeLoggingUpdateWithDefaults() *CclV1CustomCodeLoggingUpdate`

NewCclV1CustomCodeLoggingUpdateWithDefaults instantiates a new CclV1CustomCodeLoggingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CclV1CustomCodeLoggingUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CclV1CustomCodeLoggingUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CclV1CustomCodeLoggingUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CclV1CustomCodeLoggingUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CclV1CustomCodeLoggingUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CclV1CustomCodeLoggingUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CclV1CustomCodeLoggingUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CclV1CustomCodeLoggingUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CclV1CustomCodeLoggingUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CclV1CustomCodeLoggingUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CclV1CustomCodeLoggingUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CclV1CustomCodeLoggingUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CclV1CustomCodeLoggingUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CclV1CustomCodeLoggingUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CclV1CustomCodeLoggingUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CclV1CustomCodeLoggingUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDestinationSettings

`func (o *CclV1CustomCodeLoggingUpdate) GetDestinationSettings() CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf`

GetDestinationSettings returns the DestinationSettings field if non-nil, zero value otherwise.

### GetDestinationSettingsOk

`func (o *CclV1CustomCodeLoggingUpdate) GetDestinationSettingsOk() (*CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf, bool)`

GetDestinationSettingsOk returns a tuple with the DestinationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSettings

`func (o *CclV1CustomCodeLoggingUpdate) SetDestinationSettings(v CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf)`

SetDestinationSettings sets DestinationSettings field to given value.

### HasDestinationSettings

`func (o *CclV1CustomCodeLoggingUpdate) HasDestinationSettings() bool`

HasDestinationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


