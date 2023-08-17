# CcpV1PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**PluginType** | Pointer to **string** | Custom plugin type. | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of custom plugin. | [optional] 

## Methods

### NewCcpV1PresignedUrlRequest

`func NewCcpV1PresignedUrlRequest() *CcpV1PresignedUrlRequest`

NewCcpV1PresignedUrlRequest instantiates a new CcpV1PresignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1PresignedUrlRequestWithDefaults

`func NewCcpV1PresignedUrlRequestWithDefaults() *CcpV1PresignedUrlRequest`

NewCcpV1PresignedUrlRequestWithDefaults instantiates a new CcpV1PresignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CcpV1PresignedUrlRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CcpV1PresignedUrlRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CcpV1PresignedUrlRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CcpV1PresignedUrlRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CcpV1PresignedUrlRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CcpV1PresignedUrlRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CcpV1PresignedUrlRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CcpV1PresignedUrlRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CcpV1PresignedUrlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CcpV1PresignedUrlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CcpV1PresignedUrlRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CcpV1PresignedUrlRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CcpV1PresignedUrlRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CcpV1PresignedUrlRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CcpV1PresignedUrlRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CcpV1PresignedUrlRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPluginType

`func (o *CcpV1PresignedUrlRequest) GetPluginType() string`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *CcpV1PresignedUrlRequest) GetPluginTypeOk() (*string, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *CcpV1PresignedUrlRequest) SetPluginType(v string)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *CcpV1PresignedUrlRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetContentFormat

`func (o *CcpV1PresignedUrlRequest) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *CcpV1PresignedUrlRequest) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *CcpV1PresignedUrlRequest) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *CcpV1PresignedUrlRequest) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


