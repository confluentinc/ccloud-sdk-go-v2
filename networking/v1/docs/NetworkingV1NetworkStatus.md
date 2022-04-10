# NetworkingV1NetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecyle phase of the network:  PROVISIONING:  network provisioning is in progress;  READY:  network is ready;  FAILED: provisioning failed  | [readonly] 
**SupportedConnectionTypes** | [**NetworkingV1StatusConnectionTypes**](networking.v1.StatusConnectionTypes.md) |  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if network is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if network is in a failed state | [optional] [readonly] 
**DnsDomain** | Pointer to **string** | The root DNS domain for the network if applicable. Present on networks that support PrivateLink. | [optional] [readonly] 
**ZonalSubdomains** | Pointer to **map[string]string** | The DNS subdomain for each zone. Present on networks that support PrivateLink. Keys are zones and values are DNS domains.  | [optional] [readonly] 
**Cloud** | Pointer to [**NetworkingV1NetworkStatusCloudOneOf**](NetworkingV1NetworkStatusCloudOneOf.md) | The cloud-specific network details. These will be populated when the network reaches the READY state. | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkStatus

`func NewNetworkingV1NetworkStatus(phase string, supportedConnectionTypes NetworkingV1StatusConnectionTypes, ) *NetworkingV1NetworkStatus`

NewNetworkingV1NetworkStatus instantiates a new NetworkingV1NetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkStatusWithDefaults

`func NewNetworkingV1NetworkStatusWithDefaults() *NetworkingV1NetworkStatus`

NewNetworkingV1NetworkStatusWithDefaults instantiates a new NetworkingV1NetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1NetworkStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1NetworkStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1NetworkStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetSupportedConnectionTypes

`func (o *NetworkingV1NetworkStatus) GetSupportedConnectionTypes() NetworkingV1StatusConnectionTypes`

GetSupportedConnectionTypes returns the SupportedConnectionTypes field if non-nil, zero value otherwise.

### GetSupportedConnectionTypesOk

`func (o *NetworkingV1NetworkStatus) GetSupportedConnectionTypesOk() (*NetworkingV1StatusConnectionTypes, bool)`

GetSupportedConnectionTypesOk returns a tuple with the SupportedConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedConnectionTypes

`func (o *NetworkingV1NetworkStatus) SetSupportedConnectionTypes(v NetworkingV1StatusConnectionTypes)`

SetSupportedConnectionTypes sets SupportedConnectionTypes field to given value.


### GetErrorCode

`func (o *NetworkingV1NetworkStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1NetworkStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1NetworkStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1NetworkStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1NetworkStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1NetworkStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1NetworkStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1NetworkStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDnsDomain

`func (o *NetworkingV1NetworkStatus) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *NetworkingV1NetworkStatus) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *NetworkingV1NetworkStatus) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *NetworkingV1NetworkStatus) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetZonalSubdomains

`func (o *NetworkingV1NetworkStatus) GetZonalSubdomains() map[string]string`

GetZonalSubdomains returns the ZonalSubdomains field if non-nil, zero value otherwise.

### GetZonalSubdomainsOk

`func (o *NetworkingV1NetworkStatus) GetZonalSubdomainsOk() (*map[string]string, bool)`

GetZonalSubdomainsOk returns a tuple with the ZonalSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalSubdomains

`func (o *NetworkingV1NetworkStatus) SetZonalSubdomains(v map[string]string)`

SetZonalSubdomains sets ZonalSubdomains field to given value.

### HasZonalSubdomains

`func (o *NetworkingV1NetworkStatus) HasZonalSubdomains() bool`

HasZonalSubdomains returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1NetworkStatus) GetCloud() NetworkingV1NetworkStatusCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1NetworkStatus) GetCloudOk() (*NetworkingV1NetworkStatusCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1NetworkStatus) SetCloud(v NetworkingV1NetworkStatusCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1NetworkStatus) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


