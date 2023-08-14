# ArmV1AccessRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ArmV1AccessRequest**](ArmV1AccessRequest.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewArmV1AccessRequestList

`func NewArmV1AccessRequestList(apiVersion string, kind string, metadata ListMeta, data []ArmV1AccessRequest, ) *ArmV1AccessRequestList`

NewArmV1AccessRequestList instantiates a new ArmV1AccessRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArmV1AccessRequestListWithDefaults

`func NewArmV1AccessRequestListWithDefaults() *ArmV1AccessRequestList`

NewArmV1AccessRequestListWithDefaults instantiates a new ArmV1AccessRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArmV1AccessRequestList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArmV1AccessRequestList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArmV1AccessRequestList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ArmV1AccessRequestList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArmV1AccessRequestList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArmV1AccessRequestList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ArmV1AccessRequestList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArmV1AccessRequestList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArmV1AccessRequestList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ArmV1AccessRequestList) GetData() []ArmV1AccessRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArmV1AccessRequestList) GetDataOk() (*[]ArmV1AccessRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArmV1AccessRequestList) SetData(v []ArmV1AccessRequest)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


