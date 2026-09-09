# SwitchoverV1EndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A logical name for this endpoint side (e.g. \&quot;west-platt\&quot;), unique within the resource. | 
**Hostname** | Pointer to **string** | The resolved hostname for this endpoint. Set by the Switchover service. | [optional] [readonly] 
**Cloud** | Pointer to **string** | The cloud provider this endpoint resolves in. Output-only, derived server-side. | [optional] [readonly] 
**Region** | Pointer to **string** | The cloud region this endpoint resolves in. Output-only, derived server-side. | [optional] [readonly] 
**ConnectionType** | Pointer to **string** | The connection type of the resolved endpoint (e.g. &#x60;PRIVATE_LINK&#x60;), from cc-endpoint-service&#39;s vocabulary. Output-only; empty when the connection type is unspecified.  | [optional] [readonly] 
**EndpointFilter** | [**SwitchoverV1EndpointFilter**](switchover.v1.EndpointFilter.md) |  | 

## Methods

### NewSwitchoverV1EndpointConfig

`func NewSwitchoverV1EndpointConfig(name string, endpointFilter SwitchoverV1EndpointFilter, ) *SwitchoverV1EndpointConfig`

NewSwitchoverV1EndpointConfig instantiates a new SwitchoverV1EndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1EndpointConfigWithDefaults

`func NewSwitchoverV1EndpointConfigWithDefaults() *SwitchoverV1EndpointConfig`

NewSwitchoverV1EndpointConfigWithDefaults instantiates a new SwitchoverV1EndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwitchoverV1EndpointConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchoverV1EndpointConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchoverV1EndpointConfig) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *SwitchoverV1EndpointConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SwitchoverV1EndpointConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SwitchoverV1EndpointConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SwitchoverV1EndpointConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCloud

`func (o *SwitchoverV1EndpointConfig) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *SwitchoverV1EndpointConfig) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *SwitchoverV1EndpointConfig) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *SwitchoverV1EndpointConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *SwitchoverV1EndpointConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SwitchoverV1EndpointConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SwitchoverV1EndpointConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SwitchoverV1EndpointConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetConnectionType

`func (o *SwitchoverV1EndpointConfig) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SwitchoverV1EndpointConfig) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SwitchoverV1EndpointConfig) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SwitchoverV1EndpointConfig) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndpointFilter

`func (o *SwitchoverV1EndpointConfig) GetEndpointFilter() SwitchoverV1EndpointFilter`

GetEndpointFilter returns the EndpointFilter field if non-nil, zero value otherwise.

### GetEndpointFilterOk

`func (o *SwitchoverV1EndpointConfig) GetEndpointFilterOk() (*SwitchoverV1EndpointFilter, bool)`

GetEndpointFilterOk returns a tuple with the EndpointFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointFilter

`func (o *SwitchoverV1EndpointConfig) SetEndpointFilter(v SwitchoverV1EndpointFilter)`

SetEndpointFilter sets EndpointFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


