# CmkV2Freight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Freight cluster type.  | 
**MaxEcku** | Pointer to **int32** | The maximum number of Elastic Confluent Kafka Units (eCKUs) that Kafka clusters should auto-scale to. Kafka clusters with &#x60;HIGH&#x60; availability must have at least two eCKUs.  | [optional] 
**Zones** | Pointer to **[]string** | The list of zones the cluster is in.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  | [optional] [readonly] 

## Methods

### NewCmkV2Freight

`func NewCmkV2Freight(kind string, ) *CmkV2Freight`

NewCmkV2Freight instantiates a new CmkV2Freight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2FreightWithDefaults

`func NewCmkV2FreightWithDefaults() *CmkV2Freight`

NewCmkV2FreightWithDefaults instantiates a new CmkV2Freight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CmkV2Freight) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CmkV2Freight) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CmkV2Freight) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaxEcku

`func (o *CmkV2Freight) GetMaxEcku() int32`

GetMaxEcku returns the MaxEcku field if non-nil, zero value otherwise.

### GetMaxEckuOk

`func (o *CmkV2Freight) GetMaxEckuOk() (*int32, bool)`

GetMaxEckuOk returns a tuple with the MaxEcku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEcku

`func (o *CmkV2Freight) SetMaxEcku(v int32)`

SetMaxEcku sets MaxEcku field to given value.

### HasMaxEcku

`func (o *CmkV2Freight) HasMaxEcku() bool`

HasMaxEcku returns a boolean if a field has been set.

### GetZones

`func (o *CmkV2Freight) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CmkV2Freight) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CmkV2Freight) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CmkV2Freight) HasZones() bool`

HasZones returns a boolean if a field has been set.


### AsCmkV2ClusterSpecConfigOneOf

`func (s *CmkV2Freight) AsCmkV2ClusterSpecConfigOneOf() CmkV2ClusterSpecConfigOneOf`

Convenience method to wrap this instance of CmkV2Freight in CmkV2ClusterSpecConfigOneOf

### AsCmkV2ClusterSpecUpdateConfigOneOf

`func (s *CmkV2Freight) AsCmkV2ClusterSpecUpdateConfigOneOf() CmkV2ClusterSpecUpdateConfigOneOf`

Convenience method to wrap this instance of CmkV2Freight in CmkV2ClusterSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


