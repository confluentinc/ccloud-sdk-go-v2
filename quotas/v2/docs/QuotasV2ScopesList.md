# QuotasV2ScopesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]QuotasV2Scopes**](QuotasV2Scopes.md) |  | 

## Methods

### NewQuotasV2ScopesList

`func NewQuotasV2ScopesList(apiVersion string, kind string, metadata ListMeta, data []QuotasV2Scopes, ) *QuotasV2ScopesList`

NewQuotasV2ScopesList instantiates a new QuotasV2ScopesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotasV2ScopesListWithDefaults

`func NewQuotasV2ScopesListWithDefaults() *QuotasV2ScopesList`

NewQuotasV2ScopesListWithDefaults instantiates a new QuotasV2ScopesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *QuotasV2ScopesList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *QuotasV2ScopesList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *QuotasV2ScopesList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *QuotasV2ScopesList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotasV2ScopesList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotasV2ScopesList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *QuotasV2ScopesList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuotasV2ScopesList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuotasV2ScopesList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *QuotasV2ScopesList) GetData() []QuotasV2Scopes`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuotasV2ScopesList) GetDataOk() (*[]QuotasV2Scopes, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuotasV2ScopesList) SetData(v []QuotasV2Scopes)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


