# PartnerV2OrganizationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | Pointer to [**ListMeta**](ListMeta.md) |  | 
**Data** | Pointer to [**[]PartnerV2Organization**](partner.v2.Organization.md) |  | 

## Methods

### NewPartnerV2OrganizationList

`func NewPartnerV2OrganizationList(apiVersion string, kind string, metadata ListMeta, data []PartnerV2Organization, ) *PartnerV2OrganizationList`

NewPartnerV2OrganizationList instantiates a new PartnerV2OrganizationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerV2OrganizationListWithDefaults

`func NewPartnerV2OrganizationListWithDefaults() *PartnerV2OrganizationList`

NewPartnerV2OrganizationListWithDefaults instantiates a new PartnerV2OrganizationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerV2OrganizationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerV2OrganizationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerV2OrganizationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *PartnerV2OrganizationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerV2OrganizationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerV2OrganizationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *PartnerV2OrganizationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerV2OrganizationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerV2OrganizationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *PartnerV2OrganizationList) GetData() []PartnerV2Organization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PartnerV2OrganizationList) GetDataOk() (*[]PartnerV2Organization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PartnerV2OrganizationList) SetData(v []PartnerV2Organization)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


