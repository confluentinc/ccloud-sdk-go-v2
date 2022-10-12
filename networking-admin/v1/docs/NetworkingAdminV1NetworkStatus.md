# NetworkingAdminV1NetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecyle phase of the network:  PROVISIONING:  network provisioning is in progress;  READY:  network is ready;  FAILED: provisioning failed;  DEPROVISIONING: network deprovisioning is in progress;  | [readonly] 
**SupportedConnectionTypes** | [**NetworkingAdminV1SupportedConnectionTypes**](networking-admin.v1.SupportedConnectionTypes.md) |  | [readonly] 
**ActiveConnectionTypes** | [**NetworkingAdminV1ConnectionTypes**](networking-admin.v1.ConnectionTypes.md) |  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if network is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if network is in a failed state | [optional] [readonly] 
**DnsDomain** | Pointer to **string** | The root DNS domain for the network if applicable. Present on networks that support PrivateLink. | [optional] [readonly] 
**ZonalSubdomains** | Pointer to **map[string]string** | The DNS subdomain for each zone. Present on networks that support PrivateLink. Keys are zones and values are DNS domains.  | [optional] [readonly] 
**Cloud** | Pointer to [**NetworkingAdminV1NetworkStatusCloudOneOf**](NetworkingAdminV1NetworkStatusCloudOneOf.md) | The cloud-specific network details. These will be populated when the network reaches the READY state. | [optional] [readonly] 

## Methods

### NewNetworkingAdminV1NetworkStatus

`func NewNetworkingAdminV1NetworkStatus(phase string, supportedConnectionTypes NetworkingAdminV1SupportedConnectionTypes, activeConnectionTypes NetworkingAdminV1ConnectionTypes, ) *NetworkingAdminV1NetworkStatus`

NewNetworkingAdminV1NetworkStatus instantiates a new NetworkingAdminV1NetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1NetworkStatusWithDefaults

`func NewNetworkingAdminV1NetworkStatusWithDefaults() *NetworkingAdminV1NetworkStatus`

NewNetworkingAdminV1NetworkStatusWithDefaults instantiates a new NetworkingAdminV1NetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingAdminV1NetworkStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingAdminV1NetworkStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingAdminV1NetworkStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetSupportedConnectionTypes

`func (o *NetworkingAdminV1NetworkStatus) GetSupportedConnectionTypes() NetworkingAdminV1SupportedConnectionTypes`

GetSupportedConnectionTypes returns the SupportedConnectionTypes field if non-nil, zero value otherwise.

### GetSupportedConnectionTypesOk

`func (o *NetworkingAdminV1NetworkStatus) GetSupportedConnectionTypesOk() (*NetworkingAdminV1SupportedConnectionTypes, bool)`

GetSupportedConnectionTypesOk returns a tuple with the SupportedConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedConnectionTypes

`func (o *NetworkingAdminV1NetworkStatus) SetSupportedConnectionTypes(v NetworkingAdminV1SupportedConnectionTypes)`

SetSupportedConnectionTypes sets SupportedConnectionTypes field to given value.


### GetActiveConnectionTypes

`func (o *NetworkingAdminV1NetworkStatus) GetActiveConnectionTypes() NetworkingAdminV1ConnectionTypes`

GetActiveConnectionTypes returns the ActiveConnectionTypes field if non-nil, zero value otherwise.

### GetActiveConnectionTypesOk

`func (o *NetworkingAdminV1NetworkStatus) GetActiveConnectionTypesOk() (*NetworkingAdminV1ConnectionTypes, bool)`

GetActiveConnectionTypesOk returns a tuple with the ActiveConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConnectionTypes

`func (o *NetworkingAdminV1NetworkStatus) SetActiveConnectionTypes(v NetworkingAdminV1ConnectionTypes)`

SetActiveConnectionTypes sets ActiveConnectionTypes field to given value.


### GetErrorCode

`func (o *NetworkingAdminV1NetworkStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingAdminV1NetworkStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingAdminV1NetworkStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingAdminV1NetworkStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingAdminV1NetworkStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingAdminV1NetworkStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingAdminV1NetworkStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingAdminV1NetworkStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDnsDomain

`func (o *NetworkingAdminV1NetworkStatus) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *NetworkingAdminV1NetworkStatus) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *NetworkingAdminV1NetworkStatus) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *NetworkingAdminV1NetworkStatus) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetZonalSubdomains

`func (o *NetworkingAdminV1NetworkStatus) GetZonalSubdomains() map[string]string`

GetZonalSubdomains returns the ZonalSubdomains field if non-nil, zero value otherwise.

### GetZonalSubdomainsOk

`func (o *NetworkingAdminV1NetworkStatus) GetZonalSubdomainsOk() (*map[string]string, bool)`

GetZonalSubdomainsOk returns a tuple with the ZonalSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonalSubdomains

`func (o *NetworkingAdminV1NetworkStatus) SetZonalSubdomains(v map[string]string)`

SetZonalSubdomains sets ZonalSubdomains field to given value.

### HasZonalSubdomains

`func (o *NetworkingAdminV1NetworkStatus) HasZonalSubdomains() bool`

HasZonalSubdomains returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingAdminV1NetworkStatus) GetCloud() NetworkingAdminV1NetworkStatusCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingAdminV1NetworkStatus) GetCloudOk() (*NetworkingAdminV1NetworkStatusCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingAdminV1NetworkStatus) SetCloud(v NetworkingAdminV1NetworkStatusCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingAdminV1NetworkStatus) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


