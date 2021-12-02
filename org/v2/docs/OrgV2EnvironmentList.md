# OrgV2EnvironmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]OrgV2Environment**](OrgV2Environment.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewOrgV2EnvironmentList

`func NewOrgV2EnvironmentList(apiVersion string, kind string, metadata ListMeta, data []OrgV2Environment, ) *OrgV2EnvironmentList`

NewOrgV2EnvironmentList instantiates a new OrgV2EnvironmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2EnvironmentListWithDefaults

`func NewOrgV2EnvironmentListWithDefaults() *OrgV2EnvironmentList`

NewOrgV2EnvironmentListWithDefaults instantiates a new OrgV2EnvironmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrgV2EnvironmentList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrgV2EnvironmentList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrgV2EnvironmentList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *OrgV2EnvironmentList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgV2EnvironmentList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgV2EnvironmentList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *OrgV2EnvironmentList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrgV2EnvironmentList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrgV2EnvironmentList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *OrgV2EnvironmentList) GetData() []OrgV2Environment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrgV2EnvironmentList) GetDataOk() (*[]OrgV2Environment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrgV2EnvironmentList) SetData(v []OrgV2Environment)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


