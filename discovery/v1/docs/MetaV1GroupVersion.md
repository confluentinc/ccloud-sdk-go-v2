# MetaV1GroupVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the API group and version in the form \&quot;group/version\&quot; | 
**Group** | **string** | Specifies the API group name. This removes the need for API clients to split group_version manually. | 
**Version** | **string** | Specifies the API version. This removes the need for API clients to split group_version manually. | 

## Methods

### NewMetaV1GroupVersion

`func NewMetaV1GroupVersion(apiVersion string, group string, version string, ) *MetaV1GroupVersion`

NewMetaV1GroupVersion instantiates a new MetaV1GroupVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaV1GroupVersionWithDefaults

`func NewMetaV1GroupVersionWithDefaults() *MetaV1GroupVersion`

NewMetaV1GroupVersionWithDefaults instantiates a new MetaV1GroupVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MetaV1GroupVersion) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MetaV1GroupVersion) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MetaV1GroupVersion) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetGroup

`func (o *MetaV1GroupVersion) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MetaV1GroupVersion) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MetaV1GroupVersion) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetVersion

`func (o *MetaV1GroupVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetaV1GroupVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetaV1GroupVersion) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


