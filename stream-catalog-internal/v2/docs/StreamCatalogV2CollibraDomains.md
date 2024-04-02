# StreamCatalogV2CollibraDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Total** | Pointer to **int32** | The total number of domains available. | [optional] 
**Offset** | Pointer to **int32** | The starting index of the returned data within the total dataset. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of items returned in this response. | [optional] 
**Domains** | Pointer to [**[]StreamCatalogV2Domain**](StreamCatalogV2Domain.md) | The domain details | [optional] 

## Methods

### NewStreamCatalogV2CollibraDomains

`func NewStreamCatalogV2CollibraDomains() *StreamCatalogV2CollibraDomains`

NewStreamCatalogV2CollibraDomains instantiates a new StreamCatalogV2CollibraDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCatalogV2CollibraDomainsWithDefaults

`func NewStreamCatalogV2CollibraDomainsWithDefaults() *StreamCatalogV2CollibraDomains`

NewStreamCatalogV2CollibraDomainsWithDefaults instantiates a new StreamCatalogV2CollibraDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StreamCatalogV2CollibraDomains) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamCatalogV2CollibraDomains) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamCatalogV2CollibraDomains) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StreamCatalogV2CollibraDomains) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StreamCatalogV2CollibraDomains) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StreamCatalogV2CollibraDomains) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StreamCatalogV2CollibraDomains) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StreamCatalogV2CollibraDomains) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *StreamCatalogV2CollibraDomains) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamCatalogV2CollibraDomains) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamCatalogV2CollibraDomains) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamCatalogV2CollibraDomains) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamCatalogV2CollibraDomains) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamCatalogV2CollibraDomains) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamCatalogV2CollibraDomains) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamCatalogV2CollibraDomains) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTotal

`func (o *StreamCatalogV2CollibraDomains) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StreamCatalogV2CollibraDomains) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StreamCatalogV2CollibraDomains) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StreamCatalogV2CollibraDomains) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *StreamCatalogV2CollibraDomains) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StreamCatalogV2CollibraDomains) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StreamCatalogV2CollibraDomains) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StreamCatalogV2CollibraDomains) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *StreamCatalogV2CollibraDomains) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StreamCatalogV2CollibraDomains) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StreamCatalogV2CollibraDomains) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *StreamCatalogV2CollibraDomains) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetDomains

`func (o *StreamCatalogV2CollibraDomains) GetDomains() []StreamCatalogV2Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *StreamCatalogV2CollibraDomains) GetDomainsOk() (*[]StreamCatalogV2Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *StreamCatalogV2CollibraDomains) SetDomains(v []StreamCatalogV2Domain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *StreamCatalogV2CollibraDomains) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


