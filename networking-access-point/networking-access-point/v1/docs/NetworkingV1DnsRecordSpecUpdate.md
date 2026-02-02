# NetworkingV1DnsRecordSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the DNS record. | [optional] 
**Config** | Pointer to [**NetworkingV1DnsRecordSpecUpdateConfigOneOf**](NetworkingV1DnsRecordSpecUpdateConfigOneOf.md) | The config of the DNS record. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1DnsRecordSpecUpdate

`func NewNetworkingV1DnsRecordSpecUpdate() *NetworkingV1DnsRecordSpecUpdate`

NewNetworkingV1DnsRecordSpecUpdate instantiates a new NetworkingV1DnsRecordSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1DnsRecordSpecUpdateWithDefaults

`func NewNetworkingV1DnsRecordSpecUpdateWithDefaults() *NetworkingV1DnsRecordSpecUpdate`

NewNetworkingV1DnsRecordSpecUpdateWithDefaults instantiates a new NetworkingV1DnsRecordSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1DnsRecordSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1DnsRecordSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1DnsRecordSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1DnsRecordSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1DnsRecordSpecUpdate) GetConfig() NetworkingV1DnsRecordSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1DnsRecordSpecUpdate) GetConfigOk() (*NetworkingV1DnsRecordSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1DnsRecordSpecUpdate) SetConfig(v NetworkingV1DnsRecordSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1DnsRecordSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1DnsRecordSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1DnsRecordSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1DnsRecordSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1DnsRecordSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


