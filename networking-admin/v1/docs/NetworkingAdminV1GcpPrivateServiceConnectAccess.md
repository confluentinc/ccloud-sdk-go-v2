# NetworkingAdminV1GcpPrivateServiceConnectAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLink kind type. | 
**Project** | **string** | The GCP project ID for the account containing the VPCs that you want to connect from using Private Service Connect. You can find your Google Cloud Project ID under **Project ID** section of your [Google Cloud Console dashboard](https://console.cloud.google.com/home/dashboard).  | 

## Methods

### NewNetworkingAdminV1GcpPrivateServiceConnectAccess

`func NewNetworkingAdminV1GcpPrivateServiceConnectAccess(kind string, project string, ) *NetworkingAdminV1GcpPrivateServiceConnectAccess`

NewNetworkingAdminV1GcpPrivateServiceConnectAccess instantiates a new NetworkingAdminV1GcpPrivateServiceConnectAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1GcpPrivateServiceConnectAccessWithDefaults

`func NewNetworkingAdminV1GcpPrivateServiceConnectAccessWithDefaults() *NetworkingAdminV1GcpPrivateServiceConnectAccess`

NewNetworkingAdminV1GcpPrivateServiceConnectAccessWithDefaults instantiates a new NetworkingAdminV1GcpPrivateServiceConnectAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingAdminV1GcpPrivateServiceConnectAccess) SetProject(v string)`

SetProject sets Project field to given value.



### AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

`func (s *NetworkingAdminV1GcpPrivateServiceConnectAccess) AsNetworkingAdminV1PrivateLinkAccessSpecCloudOneOf() NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1GcpPrivateServiceConnectAccess in NetworkingAdminV1PrivateLinkAccessSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


