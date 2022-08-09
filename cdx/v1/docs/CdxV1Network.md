# CdxV1Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**DnsDomain** | Pointer to **string** | The root DNS domain for the network if applicable. | [optional] [readonly] 
**ZonalSubdomains** | Pointer to **map[string]string** | The DNS subdomain for each zone. Present on networks that support PrivateLink. Keys are zones and values are DNS domains.  | [optional] [readonly] 
**Cloud** | Pointer to [**CdxV1NetworkCloudOneOf**](CdxV1NetworkCloudOneOf.md) | The cloud-specific network details. These will be populated when the network reaches the READY state. | [optional] [readonly] 

## Methods

### NewCdxV1Network

`func NewCdxV1Network() *CdxV1Network`

NewCdxV1Network instantiates a new CdxV1Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1NetworkWithDefaults

`func NewCdxV1NetworkWithDefaults() *CdxV1Network`

NewCdxV1NetworkWithDefaults instantiates a new CdxV1Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1Network) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1Network) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1Network) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1Network) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1Network) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1Network) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1Network) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1Network) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1Network) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1Network) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1Network) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1Network) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDnsDomain

`func (o *CdxV1Network) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *CdxV1Network) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *CdxV1Network) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *CdxV1Network) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetZonalSubdomains

`func (o *CdxV1Network) GetZonalSubdomains() map[string]string`

GetZonalSubdomains returns the ZonalSubdomains field if non-nil, zero value otherwise.

### GetZonalSubdomainsOk

`func (o *CdxV1Network) GetZonalSubdomainsOk() (*map[string]string, bool)`

GetZonalSubdomainsOk returns a tuple with the ZonalSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalSubdomains

`func (o *CdxV1Network) SetZonalSubdomains(v map[string]string)`

SetZonalSubdomains sets ZonalSubdomains field to given value.

### HasZonalSubdomains

`func (o *CdxV1Network) HasZonalSubdomains() bool`

HasZonalSubdomains returns a boolean if a field has been set.

### GetCloud

`func (o *CdxV1Network) GetCloud() CdxV1NetworkCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CdxV1Network) GetCloudOk() (*CdxV1NetworkCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CdxV1Network) SetCloud(v CdxV1NetworkCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CdxV1Network) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


