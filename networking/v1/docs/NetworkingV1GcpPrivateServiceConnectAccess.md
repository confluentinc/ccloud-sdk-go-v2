# NetworkingV1GcpPrivateServiceConnectAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Project** | **string** | GCP project to allow for Private Service Connect access. | 

## Methods

### NewNetworkingV1GcpPrivateServiceConnectAccess

`func NewNetworkingV1GcpPrivateServiceConnectAccess(kind string, project string, ) *NetworkingV1GcpPrivateServiceConnectAccess`

NewNetworkingV1GcpPrivateServiceConnectAccess instantiates a new NetworkingV1GcpPrivateServiceConnectAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPrivateServiceConnectAccessWithDefaults

`func NewNetworkingV1GcpPrivateServiceConnectAccessWithDefaults() *NetworkingV1GcpPrivateServiceConnectAccess`

NewNetworkingV1GcpPrivateServiceConnectAccessWithDefaults instantiates a new NetworkingV1GcpPrivateServiceConnectAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1GcpPrivateServiceConnectAccess) SetProject(v string)`

SetProject sets Project field to given value.



### AsNetworkingV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingV1GcpPrivateServiceConnectAccess) AsNetworkingV1PrivateLinkAccessSpecCloudOneOf() NetworkingV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPrivateServiceConnectAccess in NetworkingV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


