# ByokV1GcpKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | The Google Cloud Platform key ID.  | 
**SecurityGroup** | Pointer to **string** | The Google security group created for this key.  | [optional] [readonly] 
**Kind** | **string** | BYOK kind type.  | 

## Methods

### NewByokV1GcpKey

`func NewByokV1GcpKey(keyId string, kind string, ) *ByokV1GcpKey`

NewByokV1GcpKey instantiates a new ByokV1GcpKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1GcpKeyWithDefaults

`func NewByokV1GcpKeyWithDefaults() *ByokV1GcpKey`

NewByokV1GcpKeyWithDefaults instantiates a new ByokV1GcpKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ByokV1GcpKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ByokV1GcpKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ByokV1GcpKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetSecurityGroup

`func (o *ByokV1GcpKey) GetSecurityGroup() string`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *ByokV1GcpKey) GetSecurityGroupOk() (*string, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *ByokV1GcpKey) SetSecurityGroup(v string)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *ByokV1GcpKey) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1GcpKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1GcpKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1GcpKey) SetKind(v string)`

SetKind sets Kind field to given value.



### AsByokV1KeyKeyOneOf

`func (s *ByokV1GcpKey) AsByokV1KeyKeyOneOf() ByokV1KeyKeyOneOf`

Convenience method to wrap this instance of ByokV1GcpKey in ByokV1KeyKeyOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


