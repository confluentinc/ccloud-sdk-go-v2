# ConfigurationcontrolV1PolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]ConfigurationcontrolV1Policy**](ConfigurationcontrolV1Policy.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewConfigurationcontrolV1PolicyList

`func NewConfigurationcontrolV1PolicyList(apiVersion string, kind string, metadata ListMeta, data []ConfigurationcontrolV1Policy, ) *ConfigurationcontrolV1PolicyList`

NewConfigurationcontrolV1PolicyList instantiates a new ConfigurationcontrolV1PolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationcontrolV1PolicyListWithDefaults

`func NewConfigurationcontrolV1PolicyListWithDefaults() *ConfigurationcontrolV1PolicyList`

NewConfigurationcontrolV1PolicyListWithDefaults instantiates a new ConfigurationcontrolV1PolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConfigurationcontrolV1PolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConfigurationcontrolV1PolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConfigurationcontrolV1PolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ConfigurationcontrolV1PolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConfigurationcontrolV1PolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConfigurationcontrolV1PolicyList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConfigurationcontrolV1PolicyList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConfigurationcontrolV1PolicyList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConfigurationcontrolV1PolicyList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ConfigurationcontrolV1PolicyList) GetData() []ConfigurationcontrolV1Policy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigurationcontrolV1PolicyList) GetDataOk() (*[]ConfigurationcontrolV1Policy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigurationcontrolV1PolicyList) SetData(v []ConfigurationcontrolV1Policy)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


