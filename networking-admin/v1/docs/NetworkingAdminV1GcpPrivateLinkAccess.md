# NetworkingAdminV1GcpPrivateLinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Project** | **string** | GCP project to allow for PrivateLink access. | 

## Methods

### NewNetworkingAdminV1GcpPrivateLinkAccess

`func NewNetworkingAdminV1GcpPrivateLinkAccess(kind string, project string, ) *NetworkingAdminV1GcpPrivateLinkAccess`

NewNetworkingAdminV1GcpPrivateLinkAccess instantiates a new NetworkingAdminV1GcpPrivateLinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1GcpPrivateLinkAccessWithDefaults

`func NewNetworkingAdminV1GcpPrivateLinkAccessWithDefaults() *NetworkingAdminV1GcpPrivateLinkAccess`

NewNetworkingAdminV1GcpPrivateLinkAccessWithDefaults instantiates a new NetworkingAdminV1GcpPrivateLinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingAdminV1GcpPrivateLinkAccess) SetProject(v string)`

SetProject sets Project field to given value.



### AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingAdminV1GcpPrivateLinkAccess) AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf() NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1GcpPrivateLinkAccess in NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


