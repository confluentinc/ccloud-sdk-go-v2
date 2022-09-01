# NetworkingV1AwsPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Peering kind type. | 
**Account** | **string** | AWS account for VPC to peer with the network | 
**Vpc** | **string** | The id of the AWS VPC to peer with | 
**Routes** | **[]string** | List of routes for the peering | 
**CustomerRegion** | **string** | Region of customer VPC | 

## Methods

### NewNetworkingV1AwsPeering

`func NewNetworkingV1AwsPeering(kind string, account string, vpc string, routes []string, customerRegion string, ) *NetworkingV1AwsPeering`

NewNetworkingV1AwsPeering instantiates a new NetworkingV1AwsPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AwsPeeringWithDefaults

`func NewNetworkingV1AwsPeeringWithDefaults() *NetworkingV1AwsPeering`

NewNetworkingV1AwsPeeringWithDefaults instantiates a new NetworkingV1AwsPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1AwsPeering) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1AwsPeering) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1AwsPeering) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAccount

`func (o *NetworkingV1AwsPeering) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkingV1AwsPeering) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkingV1AwsPeering) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetVpc

`func (o *NetworkingV1AwsPeering) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *NetworkingV1AwsPeering) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *NetworkingV1AwsPeering) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetRoutes

`func (o *NetworkingV1AwsPeering) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkingV1AwsPeering) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkingV1AwsPeering) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.


### GetCustomerRegion

`func (o *NetworkingV1AwsPeering) GetCustomerRegion() string`

GetCustomerRegion returns the CustomerRegion field if non-nil, zero value otherwise.

### GetCustomerRegionOk

`func (o *NetworkingV1AwsPeering) GetCustomerRegionOk() (*string, bool)`

GetCustomerRegionOk returns a tuple with the CustomerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRegion

`func (o *NetworkingV1AwsPeering) SetCustomerRegion(v string)`

SetCustomerRegion sets CustomerRegion field to given value.



### AsNetworkingV1PeeringSpecCloudOneOf

`func (s *NetworkingV1AwsPeering) AsNetworkingV1PeeringSpecCloudOneOf() NetworkingV1PeeringSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1AwsPeering in NetworkingV1PeeringSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


