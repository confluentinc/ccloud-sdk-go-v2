# CmkV2Dedicated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Dedicated cluster type.  | 
**Cku** | **int32** | The number of Confluent Kafka Units (CKUs) for Dedicated cluster types. MULTI_ZONE dedicated clusters must have at least two CKUs.  | 
**EncryptionKey** | Pointer to **string** | The id of the encryption key that is used to encrypt the data in the Kafka cluster. (e.g. for Amazon Web Services, the Amazon Resource Name of the key).  | [optional] 
**Zones** | Pointer to **[]string** | The list of zones the cluster is in.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  | [optional] [readonly] 

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


### GetEncryptionKey

`func (o *CmkV2Dedicated) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CmkV2Dedicated) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CmkV2Dedicated) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *CmkV2Dedicated) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetZones

`func (o *CmkV2Dedicated) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CmkV2Dedicated) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CmkV2Dedicated) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CmkV2Dedicated) HasZones() bool`

HasZones returns a boolean if a field has been set.


### AsCmkV2ClusterSpecConfigOneOf

`func (s *CmkV2Dedicated) AsCmkV2ClusterSpecConfigOneOf() CmkV2ClusterSpecConfigOneOf`

Convenience method to wrap this instance of CmkV2Dedicated in CmkV2ClusterSpecConfigOneOf

### AsCmkV2ClusterSpecUpdateConfigOneOf

`func (s *CmkV2Dedicated) AsCmkV2ClusterSpecUpdateConfigOneOf() CmkV2ClusterSpecUpdateConfigOneOf`

Convenience method to wrap this instance of CmkV2Dedicated in CmkV2ClusterSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


