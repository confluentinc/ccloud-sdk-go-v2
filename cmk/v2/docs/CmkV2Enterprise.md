# CmkV2Enterprise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Enterprise cluster type.  | 
**MaxEcku** | Pointer to **int32** | The maximum number of Elastic Confluent Kafka Units (eCKUs) that Kafka clusters should auto-scale to. Kafka clusters with &#x60;HIGH&#x60; availability must have at least two eCKUs.  | [optional] 

## Methods

### NewCmkV2Enterprise

`func NewCmkV2Enterprise(kind string, ) *CmkV2Enterprise`

NewCmkV2Enterprise instantiates a new CmkV2Enterprise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2EnterpriseWithDefaults

`func NewCmkV2EnterpriseWithDefaults() *CmkV2Enterprise`

NewCmkV2EnterpriseWithDefaults instantiates a new CmkV2Enterprise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CmkV2Enterprise) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CmkV2Enterprise) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CmkV2Enterprise) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaxEcku

`func (o *CmkV2Enterprise) GetMaxEcku() int32`

GetMaxEcku returns the MaxEcku field if non-nil, zero value otherwise.

### GetMaxEckuOk

`func (o *CmkV2Enterprise) GetMaxEckuOk() (*int32, bool)`

GetMaxEckuOk returns a tuple with the MaxEcku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEcku

`func (o *CmkV2Enterprise) SetMaxEcku(v int32)`

SetMaxEcku sets MaxEcku field to given value.

### HasMaxEcku

`func (o *CmkV2Enterprise) HasMaxEcku() bool`

HasMaxEcku returns a boolean if a field has been set.


### AsCmkV2ClusterSpecConfigOneOf

`func (s *CmkV2Enterprise) AsCmkV2ClusterSpecConfigOneOf() CmkV2ClusterSpecConfigOneOf`

Convenience method to wrap this instance of CmkV2Enterprise in CmkV2ClusterSpecConfigOneOf

### AsCmkV2ClusterSpecUpdateConfigOneOf

`func (s *CmkV2Enterprise) AsCmkV2ClusterSpecUpdateConfigOneOf() CmkV2ClusterSpecUpdateConfigOneOf`

Convenience method to wrap this instance of CmkV2Enterprise in CmkV2ClusterSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


