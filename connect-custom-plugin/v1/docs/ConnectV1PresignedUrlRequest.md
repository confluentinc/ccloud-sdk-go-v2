# ConnectV1PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of the Custom Connector Plugin. | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Custom Connector Plugin archive is uploaded. | [optional] [default to "AWS"]

## Methods

### NewConnectV1PresignedUrlRequest

`func NewConnectV1PresignedUrlRequest() *ConnectV1PresignedUrlRequest`

NewConnectV1PresignedUrlRequest instantiates a new ConnectV1PresignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1PresignedUrlRequestWithDefaults

`func NewConnectV1PresignedUrlRequestWithDefaults() *ConnectV1PresignedUrlRequest`

NewConnectV1PresignedUrlRequestWithDefaults instantiates a new ConnectV1PresignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConnectV1PresignedUrlRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConnectV1PresignedUrlRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConnectV1PresignedUrlRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConnectV1PresignedUrlRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ConnectV1PresignedUrlRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectV1PresignedUrlRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectV1PresignedUrlRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectV1PresignedUrlRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1PresignedUrlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1PresignedUrlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1PresignedUrlRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1PresignedUrlRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1PresignedUrlRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1PresignedUrlRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1PresignedUrlRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1PresignedUrlRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContentFormat

`func (o *ConnectV1PresignedUrlRequest) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ConnectV1PresignedUrlRequest) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ConnectV1PresignedUrlRequest) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ConnectV1PresignedUrlRequest) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetCloud

`func (o *ConnectV1PresignedUrlRequest) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ConnectV1PresignedUrlRequest) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ConnectV1PresignedUrlRequest) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ConnectV1PresignedUrlRequest) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


