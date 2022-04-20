# StreamGovernanceV1ClusterSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** | The billing package. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewStreamGovernanceV1ClusterSpecUpdate

`func NewStreamGovernanceV1ClusterSpecUpdate() *StreamGovernanceV1ClusterSpecUpdate`

NewStreamGovernanceV1ClusterSpecUpdate instantiates a new StreamGovernanceV1ClusterSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGovernanceV1ClusterSpecUpdateWithDefaults

`func NewStreamGovernanceV1ClusterSpecUpdateWithDefaults() *StreamGovernanceV1ClusterSpecUpdate`

NewStreamGovernanceV1ClusterSpecUpdateWithDefaults instantiates a new StreamGovernanceV1ClusterSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *StreamGovernanceV1ClusterSpecUpdate) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *StreamGovernanceV1ClusterSpecUpdate) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *StreamGovernanceV1ClusterSpecUpdate) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *StreamGovernanceV1ClusterSpecUpdate) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetEnvironment

`func (o *StreamGovernanceV1ClusterSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StreamGovernanceV1ClusterSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StreamGovernanceV1ClusterSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StreamGovernanceV1ClusterSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


