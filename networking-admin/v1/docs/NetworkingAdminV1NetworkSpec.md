# NetworkingAdminV1NetworkSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider in which the network exists. | [optional] 
**Region** | Pointer to **string** | The cloud service provider region in which the network exists. | [optional] 
**ConnectionTypes** | Pointer to [**NetworkingAdminV1ConnectionTypes**](networking-admin.v1.ConnectionTypes.md) |  | [optional] 
**Cidr** | Pointer to **string** | The IPv4 [CIDR block](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) to used for this network. Must be &#x60;/16&#x60;. Required for VPC peering and AWS TransitGateway.  | [optional] 
**Zones** | Pointer to **[]string** | The 3 availability zones for this network. They can optionally be specified for AWS networks used with PrivateLink. Otherwise, they are automatically chosen by Confluent Cloud.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  On Azure, zones are Confluent-chosen names (e.g. 1, 2, 3) since Azure does not  have universal zone identifiers.  | [optional] 
**LegacyPublicDns** | Pointer to **bool** | Enable legacy public DNS for the network. This only applies to Private Link connection type. If enabled, clusters in this network will include &#39;glb&#39; in domain name and require both public  and private DNS to resolve. If disabled, clusters in this network will not include &#39;glb&#39; in domain name and only require private  DNS to resolve  | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingAdminV1NetworkSpec

`func NewNetworkingAdminV1NetworkSpec() *NetworkingAdminV1NetworkSpec`

NewNetworkingAdminV1NetworkSpec instantiates a new NetworkingAdminV1NetworkSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1NetworkSpecWithDefaults

`func NewNetworkingAdminV1NetworkSpecWithDefaults() *NetworkingAdminV1NetworkSpec`

NewNetworkingAdminV1NetworkSpecWithDefaults instantiates a new NetworkingAdminV1NetworkSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingAdminV1NetworkSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingAdminV1NetworkSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingAdminV1NetworkSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingAdminV1NetworkSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingAdminV1NetworkSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingAdminV1NetworkSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingAdminV1NetworkSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingAdminV1NetworkSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkingAdminV1NetworkSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingAdminV1NetworkSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingAdminV1NetworkSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkingAdminV1NetworkSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetConnectionTypes

`func (o *NetworkingAdminV1NetworkSpec) GetConnectionTypes() NetworkingAdminV1ConnectionTypes`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *NetworkingAdminV1NetworkSpec) GetConnectionTypesOk() (*NetworkingAdminV1ConnectionTypes, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTypes

`func (o *NetworkingAdminV1NetworkSpec) SetConnectionTypes(v NetworkingAdminV1ConnectionTypes)`

SetConnectionTypes sets ConnectionTypes field to given value.

### HasConnectionTypes

`func (o *NetworkingAdminV1NetworkSpec) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### GetCidr

`func (o *NetworkingAdminV1NetworkSpec) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkingAdminV1NetworkSpec) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkingAdminV1NetworkSpec) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkingAdminV1NetworkSpec) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetZones

`func (o *NetworkingAdminV1NetworkSpec) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *NetworkingAdminV1NetworkSpec) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *NetworkingAdminV1NetworkSpec) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *NetworkingAdminV1NetworkSpec) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetLegacyPublicDns

`func (o *NetworkingAdminV1NetworkSpec) GetLegacyPublicDns() bool`

GetLegacyPublicDns returns the LegacyPublicDns field if non-nil, zero value otherwise.

### GetLegacyPublicDnsOk

`func (o *NetworkingAdminV1NetworkSpec) GetLegacyPublicDnsOk() (*bool, bool)`

GetLegacyPublicDnsOk returns a tuple with the LegacyPublicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyPublicDns

`func (o *NetworkingAdminV1NetworkSpec) SetLegacyPublicDns(v bool)`

SetLegacyPublicDns sets LegacyPublicDns field to given value.

### HasLegacyPublicDns

`func (o *NetworkingAdminV1NetworkSpec) HasLegacyPublicDns() bool`

HasLegacyPublicDns returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingAdminV1NetworkSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingAdminV1NetworkSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingAdminV1NetworkSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingAdminV1NetworkSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

