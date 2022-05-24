# NetworkingAdminV1AwsPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Peering kind type. | 
**Account** | **string** | AWS account for VPC to peer with the network | 
**Vpc** | **string** | The id of the AWS VPC to peer with | 
**Routes** | **[]string** | List of routes for the peering | 
**CustomerRegion** | **string** | Region of customer VPC | 

## Methods

### NewNetworkingAdminV1AwsPeering

`func NewNetworkingAdminV1AwsPeering(kind string, account string, vpc string, routes []string, customerRegion string, ) *NetworkingAdminV1AwsPeering`

NewNetworkingAdminV1AwsPeering instantiates a new NetworkingAdminV1AwsPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1AwsPeeringWithDefaults

`func NewNetworkingAdminV1AwsPeeringWithDefaults() *NetworkingAdminV1AwsPeering`

NewNetworkingAdminV1AwsPeeringWithDefaults instantiates a new NetworkingAdminV1AwsPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1AwsPeering) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1AwsPeering) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1AwsPeering) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingAdminV1AwsPeering) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingAdminV1AwsPeering) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingAdminV1AwsPeering) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetVpc

`func (o *NetworkingAdminV1AwsPeering) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingAdminV1AwsPeering) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingAdminV1AwsPeering) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetRoutes

`func (o *NetworkingAdminV1AwsPeering) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingAdminV1AwsPeering) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingAdminV1AwsPeering) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.


### GetCustomerRegion

`func (o *NetworkingAdminV1AwsPeering) GetCustomerRegion() string`

GetCustomerRegion returns the CustomerRegion field if non-nil, zero value otherwise.

### GetCustomerRegionOk

`func (o *NetworkingAdminV1AwsPeering) GetCustomerRegionOk() (*string, bool)`

GetCustomerRegionOk returns a tuple with the CustomerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRegion

`func (o *NetworkingAdminV1AwsPeering) SetCustomerRegion(v string)`

SetCustomerRegion sets CustomerRegion field to given value.



### AsNetworkingAdminV1PeeringSpecCloudOneOf

`func (s *NetworkingAdminV1AwsPeering) AsNetworkingAdminV1PeeringSpecCloudOneOf() NetworkingAdminV1PeeringSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1AwsPeering in NetworkingAdminV1PeeringSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


