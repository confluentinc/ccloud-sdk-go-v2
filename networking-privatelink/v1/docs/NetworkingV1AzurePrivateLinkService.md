# NetworkingV1AzurePrivateLinkService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | Availability zone associated with the Azure PrivateLink service. | [readonly] 
**PrivateLinkServiceAlias** | **string** | Azure PrivateLink service alias for the availability zone. | [readonly] 
**PrivateLinkServiceResourceId** | **string** | Azure PrivateLink service resource id for the availability zone. | [readonly] 

## Methods

### NewNetworkingV1AzurePrivateLinkService

`func NewNetworkingV1AzurePrivateLinkService(zone string, privateLinkServiceAlias string, privateLinkServiceResourceId string, ) *NetworkingV1AzurePrivateLinkService`

NewNetworkingV1AzurePrivateLinkService instantiates a new NetworkingV1AzurePrivateLinkService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AzurePrivateLinkServiceWithDefaults

`func NewNetworkingV1AzurePrivateLinkServiceWithDefaults() *NetworkingV1AzurePrivateLinkService`

NewNetworkingV1AzurePrivateLinkServiceWithDefaults instantiates a new NetworkingV1AzurePrivateLinkService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *NetworkingV1AzurePrivateLinkService) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NetworkingV1AzurePrivateLinkService) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NetworkingV1AzurePrivateLinkService) SetZone(v string)`

SetZone sets Zone field to given value.


### GetPrivateLinkServiceAlias

`func (o *NetworkingV1AzurePrivateLinkService) GetPrivateLinkServiceAlias() string`

GetPrivateLinkServiceAlias returns the PrivateLinkServiceAlias field if non-nil, zero value otherwise.

### GetPrivateLinkServiceAliasOk

`func (o *NetworkingV1AzurePrivateLinkService) GetPrivateLinkServiceAliasOk() (*string, bool)`

GetPrivateLinkServiceAliasOk returns a tuple with the PrivateLinkServiceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceAlias

`func (o *NetworkingV1AzurePrivateLinkService) SetPrivateLinkServiceAlias(v string)`

SetPrivateLinkServiceAlias sets PrivateLinkServiceAlias field to given value.


### GetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzurePrivateLinkService) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *NetworkingV1AzurePrivateLinkService) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *NetworkingV1AzurePrivateLinkService) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


