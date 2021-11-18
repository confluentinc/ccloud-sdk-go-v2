# V2ServiceAccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]V2ServiceAccount**](V2ServiceAccount.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewV2ServiceAccountList

`func NewV2ServiceAccountList(apiVersion string, kind string, metadata ListMeta, data []V2ServiceAccount, ) *V2ServiceAccountList`

NewV2ServiceAccountList instantiates a new V2ServiceAccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ServiceAccountListWithDefaults

`func NewV2ServiceAccountListWithDefaults() *V2ServiceAccountList`

NewV2ServiceAccountListWithDefaults instantiates a new V2ServiceAccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V2ServiceAccountList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V2ServiceAccountList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V2ServiceAccountList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *V2ServiceAccountList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V2ServiceAccountList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V2ServiceAccountList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *V2ServiceAccountList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V2ServiceAccountList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V2ServiceAccountList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *V2ServiceAccountList) GetData() []V2ServiceAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2ServiceAccountList) GetDataOk() (*[]V2ServiceAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2ServiceAccountList) SetData(v []V2ServiceAccount)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


