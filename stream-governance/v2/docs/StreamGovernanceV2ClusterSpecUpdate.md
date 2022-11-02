# StreamGovernanceV2ClusterSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** | The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS.  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewStreamGovernanceV2ClusterSpecUpdate

`func NewStreamGovernanceV2ClusterSpecUpdate() *StreamGovernanceV2ClusterSpecUpdate`

NewStreamGovernanceV2ClusterSpecUpdate instantiates a new StreamGovernanceV2ClusterSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV2ClusterSpecUpdateWithDefaults

`func NewStreamGovernanceV2ClusterSpecUpdateWithDefaults() *StreamGovernanceV2ClusterSpecUpdate`

NewStreamGovernanceV2ClusterSpecUpdateWithDefaults instantiates a new StreamGovernanceV2ClusterSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *StreamGovernanceV2ClusterSpecUpdate) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *StreamGovernanceV2ClusterSpecUpdate) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *StreamGovernanceV2ClusterSpecUpdate) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *StreamGovernanceV2ClusterSpecUpdate) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetEnvironment

`func (o *StreamGovernanceV2ClusterSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StreamGovernanceV2ClusterSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StreamGovernanceV2ClusterSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StreamGovernanceV2ClusterSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


