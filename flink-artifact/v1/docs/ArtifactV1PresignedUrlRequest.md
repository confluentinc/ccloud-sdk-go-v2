# ArtifactV1PresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ContentFormat** | Pointer to **string** | Archive format of the Flink Artifact. | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Flink Artifact archive is uploaded. | [optional] 
**Region** | Pointer to **string** | The Cloud provider region the Flink Artifact archive is uploaded. | [optional] 
**Environment** | Pointer to **string** | The Environment the uploaded Flink Artifact belongs to. | [optional] 

## Methods

### NewArtifactV1PresignedUrlRequest

`func NewArtifactV1PresignedUrlRequest() *ArtifactV1PresignedUrlRequest`

NewArtifactV1PresignedUrlRequest instantiates a new ArtifactV1PresignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1PresignedUrlRequestWithDefaults

`func NewArtifactV1PresignedUrlRequestWithDefaults() *ArtifactV1PresignedUrlRequest`

NewArtifactV1PresignedUrlRequestWithDefaults instantiates a new ArtifactV1PresignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArtifactV1PresignedUrlRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArtifactV1PresignedUrlRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArtifactV1PresignedUrlRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArtifactV1PresignedUrlRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ArtifactV1PresignedUrlRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArtifactV1PresignedUrlRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArtifactV1PresignedUrlRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArtifactV1PresignedUrlRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *ArtifactV1PresignedUrlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactV1PresignedUrlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactV1PresignedUrlRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactV1PresignedUrlRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactV1PresignedUrlRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactV1PresignedUrlRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactV1PresignedUrlRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactV1PresignedUrlRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetContentFormat

`func (o *ArtifactV1PresignedUrlRequest) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *ArtifactV1PresignedUrlRequest) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *ArtifactV1PresignedUrlRequest) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *ArtifactV1PresignedUrlRequest) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetCloud

`func (o *ArtifactV1PresignedUrlRequest) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ArtifactV1PresignedUrlRequest) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ArtifactV1PresignedUrlRequest) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ArtifactV1PresignedUrlRequest) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *ArtifactV1PresignedUrlRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArtifactV1PresignedUrlRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArtifactV1PresignedUrlRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArtifactV1PresignedUrlRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironment

`func (o *ArtifactV1PresignedUrlRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ArtifactV1PresignedUrlRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ArtifactV1PresignedUrlRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ArtifactV1PresignedUrlRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


