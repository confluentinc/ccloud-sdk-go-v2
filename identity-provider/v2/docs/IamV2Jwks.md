# IamV2Jwks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Spec** | Pointer to [**IamV2JwksSpec**](IamV2JwksSpec.md) |  | [optional] 
**Status** | Pointer to [**IamV2JwksStatus**](IamV2JwksStatus.md) |  | [optional] 

## Methods

### NewIamV2Jwks

`func NewIamV2Jwks() *IamV2Jwks`

NewIamV2Jwks instantiates a new IamV2Jwks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2JwksWithDefaults

`func NewIamV2JwksWithDefaults() *IamV2Jwks`

NewIamV2JwksWithDefaults instantiates a new IamV2Jwks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2Jwks) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2Jwks) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2Jwks) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2Jwks) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2Jwks) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2Jwks) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2Jwks) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2Jwks) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSpec

`func (o *IamV2Jwks) GetSpec() IamV2JwksSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IamV2Jwks) GetSpecOk() (*IamV2JwksSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IamV2Jwks) SetSpec(v IamV2JwksSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IamV2Jwks) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IamV2Jwks) GetStatus() IamV2JwksStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamV2Jwks) GetStatusOk() (*IamV2JwksStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamV2Jwks) SetStatus(v IamV2JwksStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamV2Jwks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


