# CmkV2Dedicated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Dedicated cluster type.  | 
**Cku** | **int32** | The number of Confluent Kafka Units (CKUs) for Dedicated cluster types. MULTI_ZONE dedicated clusters must have at least two CKUs.  | 

## Methods

### NewCmkV2Dedicated

`func NewCmkV2Dedicated(kind string, cku int32, ) *CmkV2Dedicated`

NewCmkV2Dedicated instantiates a new CmkV2Dedicated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2DedicatedWithDefaults

`func NewCmkV2DedicatedWithDefaults() *CmkV2Dedicated`

NewCmkV2DedicatedWithDefaults instantiates a new CmkV2Dedicated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CmkV2Dedicated) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CmkV2Dedicated) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CmkV2Dedicated) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCku

`func (o *CmkV2Dedicated) GetCku() int32`

GetCku returns the Cku field if non-nil, zero value otherwise.

### GetCkuOk

`func (o *CmkV2Dedicated) GetCkuOk() (*int32, bool)`

GetCkuOk returns a tuple with the Cku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCku

`func (o *CmkV2Dedicated) SetCku(v int32)`

SetCku sets Cku field to given value.



### AsCmkV2ClusterSpecConfigOneOf

`func (s *CmkV2Dedicated) AsCmkV2ClusterSpecConfigOneOf() CmkV2ClusterSpecConfigOneOf`

Convenience method to wrap this instance of CmkV2Dedicated in CmkV2ClusterSpecConfigOneOf

### AsCmkV2ClusterSpecUpdateConfigOneOf

`func (s *CmkV2Dedicated) AsCmkV2ClusterSpecUpdateConfigOneOf() CmkV2ClusterSpecUpdateConfigOneOf`

Convenience method to wrap this instance of CmkV2Dedicated in CmkV2ClusterSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


