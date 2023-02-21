# OrgV2OrganizationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]OrgV2Organization**](OrgV2Organization.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewOrgV2OrganizationList

`func NewOrgV2OrganizationList(apiVersion string, kind string, metadata ListMeta, data []OrgV2Organization, ) *OrgV2OrganizationList`

NewOrgV2OrganizationList instantiates a new OrgV2OrganizationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2OrganizationListWithDefaults

`func NewOrgV2OrganizationListWithDefaults() *OrgV2OrganizationList`

NewOrgV2OrganizationListWithDefaults instantiates a new OrgV2OrganizationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrgV2OrganizationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrgV2OrganizationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrgV2OrganizationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *OrgV2OrganizationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgV2OrganizationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgV2OrganizationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *OrgV2OrganizationList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrgV2OrganizationList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrgV2OrganizationList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *OrgV2OrganizationList) GetData() []OrgV2Organization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrgV2OrganizationList) GetDataOk() (*[]OrgV2Organization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrgV2OrganizationList) SetData(v []OrgV2Organization)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


