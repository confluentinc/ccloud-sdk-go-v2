# CmkV2Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind defines the object this REST resource represents. | 
**Cloud** | **string** | The cloud service provider in which the cluster is running. | 
**Region** | **string** | The cloud service provider region where the cluster is running. | 

## Methods

### NewCmkV2Location

`func NewCmkV2Location(kind string, cloud string, region string, ) *CmkV2Location`

NewCmkV2Location instantiates a new CmkV2Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2LocationWithDefaults

`func NewCmkV2LocationWithDefaults() *CmkV2Location`

NewCmkV2LocationWithDefaults instantiates a new CmkV2Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CmkV2Location) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CmkV2Location) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CmkV2Location) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCloud

`func (o *CmkV2Location) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CmkV2Location) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CmkV2Location) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *CmkV2Location) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CmkV2Location) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CmkV2Location) SetRegion(v string)`

SetRegion sets Region field to given value.



### AsCmkV2KafkaClusterSpecPlacementOneOf

`func (s *CmkV2Location) AsCmkV2KafkaClusterSpecPlacementOneOf() CmkV2KafkaClusterSpecPlacementOneOf`

Convenience method to wrap this instance of CmkV2Location in CmkV2KafkaClusterSpecPlacementOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


