# StreamCatalogV2CollibraDomainsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Url** | Pointer to **string** | The URL provided is specific to Confluent&#39;s Collibra account. | [optional] 
**Username** | Pointer to **string** | The username for authentication. | [optional] 
**Password** | Pointer to **string** | The password for authentication. | [optional] 

## Methods

### NewStreamCatalogV2CollibraDomainsRequest

`func NewStreamCatalogV2CollibraDomainsRequest() *StreamCatalogV2CollibraDomainsRequest`

NewStreamCatalogV2CollibraDomainsRequest instantiates a new StreamCatalogV2CollibraDomainsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCatalogV2CollibraDomainsRequestWithDefaults

`func NewStreamCatalogV2CollibraDomainsRequestWithDefaults() *StreamCatalogV2CollibraDomainsRequest`

NewStreamCatalogV2CollibraDomainsRequestWithDefaults instantiates a new StreamCatalogV2CollibraDomainsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StreamCatalogV2CollibraDomainsRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamCatalogV2CollibraDomainsRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StreamCatalogV2CollibraDomainsRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StreamCatalogV2CollibraDomainsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamCatalogV2CollibraDomainsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StreamCatalogV2CollibraDomainsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *StreamCatalogV2CollibraDomainsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamCatalogV2CollibraDomainsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamCatalogV2CollibraDomainsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamCatalogV2CollibraDomainsRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamCatalogV2CollibraDomainsRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamCatalogV2CollibraDomainsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUrl

`func (o *StreamCatalogV2CollibraDomainsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StreamCatalogV2CollibraDomainsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StreamCatalogV2CollibraDomainsRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *StreamCatalogV2CollibraDomainsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StreamCatalogV2CollibraDomainsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StreamCatalogV2CollibraDomainsRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *StreamCatalogV2CollibraDomainsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StreamCatalogV2CollibraDomainsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StreamCatalogV2CollibraDomainsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StreamCatalogV2CollibraDomainsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


