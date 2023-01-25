# NetworkingV1NetworkSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider in which the network exists. | [optional] 
**Region** | Pointer to **string** | The cloud service provider region in which the network exists. | [optional] 
**ConnectionTypes** | Pointer to [**NetworkingV1ConnectionTypes**](networking.v1.ConnectionTypes.md) |  | [optional] 
**Cidr** | Pointer to **string** | The IPv4 [CIDR block](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) to used for this network. Must be &#x60;/16&#x60;. Required for VPC peering and AWS TransitGateway.  | [optional] 
**Zones** | Pointer to **[]string** | The 3 availability zones for this network. They can optionally be specified for AWS networks used with PrivateLink, for GCP networks used with Private Service Connect, and for AWS and GCP networks used with Peering. Otherwise, they are automatically chosen by Confluent Cloud.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  On Azure, zones are Confluent-chosen names (e.g. 1, 2, 3) since Azure does not  have universal zone identifiers.  | [optional] 
**ZonesInfo** | Pointer to [**NetworkingV1ZonesInfo**](networking.v1.ZonesInfo.md) | Each item represents information related to a single zone.  Note - The attribute is in an [Early Access lifecycle stage]   (https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)  | [optional] 
**DnsConfig** | Pointer to [**NetworkingV1DnsConfig**](networking.v1.DnsConfig.md) | DNS config only applies to PrivateLink network connection type.  When resolution is CHASED_PRIVATE, clusters in this network require both public and private DNS  to resolve cluster endpoints.  When resolution is PRIVATE, clusters in this network only require private DNS  to resolve cluster endpoints.  Note - The attribute is in an [Early Access lifecycle stage](https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)  and it&#39;s available only for AWS networks with PRIVATELINK connection type.  | [optional] 
**ReservedCidr** | Pointer to **string** | The reserved CIDR config is used only by AWS networks with connection_types &#x3D; Vpc_Peering or Transit_Gateway  An IPv4 [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)   reserved for Confluent Cloud Network. Must be \\24.   If not specified, Confluent Cloud Network uses 172.20.255.0/24  Note - The attribute is in an [Early Access lifecycle stage](https://docs.confluent.io/cloud/current/api.html#section/Versioning/API-Lifecycle-Policy)  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1NetworkSpec

`func NewNetworkingV1NetworkSpec() *NetworkingV1NetworkSpec`

NewNetworkingV1NetworkSpec instantiates a new NetworkingV1NetworkSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkSpecWithDefaults

`func NewNetworkingV1NetworkSpecWithDefaults() *NetworkingV1NetworkSpec`

NewNetworkingV1NetworkSpecWithDefaults instantiates a new NetworkingV1NetworkSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1NetworkSpec) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1NetworkSpec) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1NetworkSpec) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1NetworkSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkingV1NetworkSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkingV1NetworkSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkingV1NetworkSpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkingV1NetworkSpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetConnectionTypes

`func (o *NetworkingV1NetworkSpec) GetConnectionTypes() NetworkingV1ConnectionTypes`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *NetworkingV1NetworkSpec) GetConnectionTypesOk() (*NetworkingV1ConnectionTypes, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTypes

`func (o *NetworkingV1NetworkSpec) SetConnectionTypes(v NetworkingV1ConnectionTypes)`

SetConnectionTypes sets ConnectionTypes field to given value.

### HasConnectionTypes

`func (o *NetworkingV1NetworkSpec) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### GetCidr

`func (o *NetworkingV1NetworkSpec) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkingV1NetworkSpec) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkingV1NetworkSpec) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkingV1NetworkSpec) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetZones

`func (o *NetworkingV1NetworkSpec) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *NetworkingV1NetworkSpec) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *NetworkingV1NetworkSpec) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *NetworkingV1NetworkSpec) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZonesInfo

`func (o *NetworkingV1NetworkSpec) GetZonesInfo() NetworkingV1ZonesInfo`

GetZonesInfo returns the ZonesInfo field if non-nil, zero value otherwise.

### GetZonesInfoOk

`func (o *NetworkingV1NetworkSpec) GetZonesInfoOk() (*NetworkingV1ZonesInfo, bool)`

GetZonesInfoOk returns a tuple with the ZonesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonesInfo

`func (o *NetworkingV1NetworkSpec) SetZonesInfo(v NetworkingV1ZonesInfo)`

SetZonesInfo sets ZonesInfo field to given value.

### HasZonesInfo

`func (o *NetworkingV1NetworkSpec) HasZonesInfo() bool`

HasZonesInfo returns a boolean if a field has been set.

### GetDnsConfig

`func (o *NetworkingV1NetworkSpec) GetDnsConfig() NetworkingV1DnsConfig`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *NetworkingV1NetworkSpec) GetDnsConfigOk() (*NetworkingV1DnsConfig, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *NetworkingV1NetworkSpec) SetDnsConfig(v NetworkingV1DnsConfig)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *NetworkingV1NetworkSpec) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetReservedCidr

`func (o *NetworkingV1NetworkSpec) GetReservedCidr() string`

GetReservedCidr returns the ReservedCidr field if non-nil, zero value otherwise.

### GetReservedCidrOk

`func (o *NetworkingV1NetworkSpec) GetReservedCidrOk() (*string, bool)`

GetReservedCidrOk returns a tuple with the ReservedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidr

`func (o *NetworkingV1NetworkSpec) SetReservedCidr(v string)`

SetReservedCidr sets ReservedCidr field to given value.

### HasReservedCidr

`func (o *NetworkingV1NetworkSpec) HasReservedCidr() bool`

HasReservedCidr returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


