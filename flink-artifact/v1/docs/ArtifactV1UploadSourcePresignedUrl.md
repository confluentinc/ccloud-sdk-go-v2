# ArtifactV1UploadSourcePresignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Location** | Pointer to **string** | Location of the Flink Artifact source.  | [optional] 
**UploadId** | Pointer to **string** | Upload ID returned by the &#x60;/presigned-upload-url&#x60; API. This field returns an empty string in all responses. | [optional] 

## Methods

### NewArtifactV1UploadSourcePresignedUrl

`func NewArtifactV1UploadSourcePresignedUrl() *ArtifactV1UploadSourcePresignedUrl`

NewArtifactV1UploadSourcePresignedUrl instantiates a new ArtifactV1UploadSourcePresignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1UploadSourcePresignedUrlWithDefaults

`func NewArtifactV1UploadSourcePresignedUrlWithDefaults() *ArtifactV1UploadSourcePresignedUrl`

NewArtifactV1UploadSourcePresignedUrlWithDefaults instantiates a new ArtifactV1UploadSourcePresignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1UploadSourcePresignedUrl) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1UploadSourcePresignedUrl) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArtifactV1UploadSourcePresignedUrl) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArtifactV1UploadSourcePresignedUrl) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1UploadSourcePresignedUrl) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArtifactV1UploadSourcePresignedUrl) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArtifactV1UploadSourcePresignedUrl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactV1UploadSourcePresignedUrl) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactV1UploadSourcePresignedUrl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactV1UploadSourcePresignedUrl) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactV1UploadSourcePresignedUrl) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactV1UploadSourcePresignedUrl) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLocation

`func (o *ArtifactV1UploadSourcePresignedUrl) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ArtifactV1UploadSourcePresignedUrl) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ArtifactV1UploadSourcePresignedUrl) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUploadId

`func (o *ArtifactV1UploadSourcePresignedUrl) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *ArtifactV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *ArtifactV1UploadSourcePresignedUrl) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *ArtifactV1UploadSourcePresignedUrl) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.


### AsArtifactV1FlinkArtifactVersionUploadSourceOneOf

`func (s *ArtifactV1UploadSourcePresignedUrl) AsArtifactV1FlinkArtifactVersionUploadSourceOneOf() ArtifactV1FlinkArtifactVersionUploadSourceOneOf`

Convenience method to wrap this instance of ArtifactV1UploadSourcePresignedUrl in ArtifactV1FlinkArtifactVersionUploadSourceOneOf

### AsInlineObjectUploadSourceOneOf

`func (s *ArtifactV1UploadSourcePresignedUrl) AsInlineObjectUploadSourceOneOf() InlineObjectUploadSourceOneOf`

Convenience method to wrap this instance of ArtifactV1UploadSourcePresignedUrl in InlineObjectUploadSourceOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


