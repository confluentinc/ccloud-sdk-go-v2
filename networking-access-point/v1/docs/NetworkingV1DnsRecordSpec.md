# NetworkingV1DnsRecordSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the DNS record. | [optional] 
**Domain** | Pointer to **string** | The fully qualified domain name of the DNS record. | [optional] 
**Config** | Pointer to [**NetworkingV1DnsRecordSpecConfigOneOf**](NetworkingV1DnsRecordSpecConfigOneOf.md) | The config of the DNS record. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Gateway** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The gateway to which this belongs. | [optional] 

## Methods

### NewNetworkingV1DnsRecordSpec

`func NewNetworkingV1DnsRecordSpec() *NetworkingV1DnsRecordSpec`

NewNetworkingV1DnsRecordSpec instantiates a new NetworkingV1DnsRecordSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1DnsRecordSpecWithDefaults

`func NewNetworkingV1DnsRecordSpecWithDefaults() *NetworkingV1DnsRecordSpec`

NewNetworkingV1DnsRecordSpecWithDefaults instantiates a new NetworkingV1DnsRecordSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1DnsRecordSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1DnsRecordSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1DnsRecordSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1DnsRecordSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *NetworkingV1DnsRecordSpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NetworkingV1DnsRecordSpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NetworkingV1DnsRecordSpec) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NetworkingV1DnsRecordSpec) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1DnsRecordSpec) GetConfig() NetworkingV1DnsRecordSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1DnsRecordSpec) GetConfigOk() (*NetworkingV1DnsRecordSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1DnsRecordSpec) SetConfig(v NetworkingV1DnsRecordSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1DnsRecordSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1DnsRecordSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1DnsRecordSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1DnsRecordSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1DnsRecordSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkingV1DnsRecordSpec) GetGateway() EnvScopedObjectReference`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkingV1DnsRecordSpec) GetGatewayOk() (*EnvScopedObjectReference, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkingV1DnsRecordSpec) SetGateway(v EnvScopedObjectReference)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkingV1DnsRecordSpec) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


