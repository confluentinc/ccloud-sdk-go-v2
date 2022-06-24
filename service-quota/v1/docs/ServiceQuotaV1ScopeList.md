# ServiceQuotaV1ScopeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ServiceQuotaV1Scope**](ServiceQuotaV1Scope.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewServiceQuotaV1ScopeList

`func NewServiceQuotaV1ScopeList(apiVersion string, kind string, metadata ListMeta, data []ServiceQuotaV1Scope, ) *ServiceQuotaV1ScopeList`

NewServiceQuotaV1ScopeList instantiates a new ServiceQuotaV1ScopeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQuotaV1ScopeListWithDefaults

`func NewServiceQuotaV1ScopeListWithDefaults() *ServiceQuotaV1ScopeList`

NewServiceQuotaV1ScopeListWithDefaults instantiates a new ServiceQuotaV1ScopeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ServiceQuotaV1ScopeList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ServiceQuotaV1ScopeList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ServiceQuotaV1ScopeList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ServiceQuotaV1ScopeList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceQuotaV1ScopeList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceQuotaV1ScopeList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ServiceQuotaV1ScopeList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceQuotaV1ScopeList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceQuotaV1ScopeList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ServiceQuotaV1ScopeList) GetData() []ServiceQuotaV1Scope`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceQuotaV1ScopeList) GetDataOk() (*[]ServiceQuotaV1Scope, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceQuotaV1ScopeList) SetData(v []ServiceQuotaV1Scope)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


