# StreamCatalogV2CollibraError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**StatusCode** | Pointer to **int32** | The HTTP status code applicable to this error. | [optional] 
**UserMessage** | Pointer to **string** | A human-readable message aimed at the end user, providing more details about the error. | [optional] 

## Methods

### NewStreamCatalogV2CollibraError

`func NewStreamCatalogV2CollibraError() *StreamCatalogV2CollibraError`

NewStreamCatalogV2CollibraError instantiates a new StreamCatalogV2CollibraError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCatalogV2CollibraErrorWithDefaults

`func NewStreamCatalogV2CollibraErrorWithDefaults() *StreamCatalogV2CollibraError`

NewStreamCatalogV2CollibraErrorWithDefaults instantiates a new StreamCatalogV2CollibraError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StreamCatalogV2CollibraError) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamCatalogV2CollibraError) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamCatalogV2CollibraError) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StreamCatalogV2CollibraError) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StreamCatalogV2CollibraError) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamCatalogV2CollibraError) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamCatalogV2CollibraError) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StreamCatalogV2CollibraError) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *StreamCatalogV2CollibraError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamCatalogV2CollibraError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamCatalogV2CollibraError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamCatalogV2CollibraError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamCatalogV2CollibraError) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamCatalogV2CollibraError) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamCatalogV2CollibraError) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamCatalogV2CollibraError) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatusCode

`func (o *StreamCatalogV2CollibraError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *StreamCatalogV2CollibraError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *StreamCatalogV2CollibraError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *StreamCatalogV2CollibraError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetUserMessage

`func (o *StreamCatalogV2CollibraError) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *StreamCatalogV2CollibraError) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *StreamCatalogV2CollibraError) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *StreamCatalogV2CollibraError) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


