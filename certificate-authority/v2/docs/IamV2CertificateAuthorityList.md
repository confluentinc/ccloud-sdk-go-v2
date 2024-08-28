# IamV2CertificateAuthorityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]IamV2CertificateAuthority**](IamV2CertificateAuthority.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewIamV2CertificateAuthorityList

`func NewIamV2CertificateAuthorityList(apiVersion string, kind string, metadata ListMeta, data []IamV2CertificateAuthority, ) *IamV2CertificateAuthorityList`

NewIamV2CertificateAuthorityList instantiates a new IamV2CertificateAuthorityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2CertificateAuthorityListWithDefaults

`func NewIamV2CertificateAuthorityListWithDefaults() *IamV2CertificateAuthorityList`

NewIamV2CertificateAuthorityListWithDefaults instantiates a new IamV2CertificateAuthorityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2CertificateAuthorityList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2CertificateAuthorityList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2CertificateAuthorityList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *IamV2CertificateAuthorityList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2CertificateAuthorityList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2CertificateAuthorityList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *IamV2CertificateAuthorityList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2CertificateAuthorityList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2CertificateAuthorityList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *IamV2CertificateAuthorityList) GetData() []IamV2CertificateAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamV2CertificateAuthorityList) GetDataOk() (*[]IamV2CertificateAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamV2CertificateAuthorityList) SetData(v []IamV2CertificateAuthority)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


