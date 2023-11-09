# IamV2IpFilterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]IamV2IpFilter**](IamV2IpFilter.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewIamV2IpFilterList

`func NewIamV2IpFilterList(apiVersion string, kind string, metadata ListMeta, data []IamV2IpFilter, ) *IamV2IpFilterList`

NewIamV2IpFilterList instantiates a new IamV2IpFilterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IpFilterListWithDefaults

`func NewIamV2IpFilterListWithDefaults() *IamV2IpFilterList`

NewIamV2IpFilterListWithDefaults instantiates a new IamV2IpFilterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IpFilterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IpFilterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IpFilterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *IamV2IpFilterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IpFilterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IpFilterList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *IamV2IpFilterList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IpFilterList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IpFilterList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *IamV2IpFilterList) GetData() []IamV2IpFilter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamV2IpFilterList) GetDataOk() (*[]IamV2IpFilter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamV2IpFilterList) SetData(v []IamV2IpFilter)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


