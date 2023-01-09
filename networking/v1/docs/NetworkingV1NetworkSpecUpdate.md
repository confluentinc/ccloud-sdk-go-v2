# NetworkingV1NetworkSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network | [optional] 
**ZoneInfo** | Pointer to [**NetworkingV1ZoneInfo**](networking.v1.ZoneInfo.md) | Each item represents information related to a single zone.  Note - The attribute is in an [Early Access lifecycle stage]   (https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1NetworkSpecUpdate

`func NewNetworkingV1NetworkSpecUpdate() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdate instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkSpecUpdateWithDefaults

`func NewNetworkingV1NetworkSpecUpdateWithDefaults() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetZoneInfo

`func (o *NetworkingV1NetworkSpecUpdate) GetZoneInfo() NetworkingV1ZoneInfo`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *NetworkingV1NetworkSpecUpdate) GetZoneInfoOk() (*NetworkingV1ZoneInfo, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *NetworkingV1NetworkSpecUpdate) SetZoneInfo(v NetworkingV1ZoneInfo)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *NetworkingV1NetworkSpecUpdate) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


