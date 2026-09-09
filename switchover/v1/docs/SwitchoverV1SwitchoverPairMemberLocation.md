# SwitchoverV1SwitchoverPairMemberLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to **string** | The cloud provider hosting this member (e.g. &#x60;aws&#x60;, &#x60;gcp&#x60;, &#x60;azure&#x60;). | [optional] 
**Region** | Pointer to **string** | The cloud region hosting this member (e.g. &#x60;us-west-2&#x60;). | [optional] 

## Methods

### NewSwitchoverV1SwitchoverPairMemberLocation

`func NewSwitchoverV1SwitchoverPairMemberLocation() *SwitchoverV1SwitchoverPairMemberLocation`

NewSwitchoverV1SwitchoverPairMemberLocation instantiates a new SwitchoverV1SwitchoverPairMemberLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairMemberLocationWithDefaults

`func NewSwitchoverV1SwitchoverPairMemberLocationWithDefaults() *SwitchoverV1SwitchoverPairMemberLocation`

NewSwitchoverV1SwitchoverPairMemberLocationWithDefaults instantiates a new SwitchoverV1SwitchoverPairMemberLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *SwitchoverV1SwitchoverPairMemberLocation) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *SwitchoverV1SwitchoverPairMemberLocation) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *SwitchoverV1SwitchoverPairMemberLocation) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *SwitchoverV1SwitchoverPairMemberLocation) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *SwitchoverV1SwitchoverPairMemberLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SwitchoverV1SwitchoverPairMemberLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SwitchoverV1SwitchoverPairMemberLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SwitchoverV1SwitchoverPairMemberLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


