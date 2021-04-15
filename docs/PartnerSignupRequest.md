# PartnerSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**PartnerV2Organization**](partner.v2.Organization.md) |  | 
**User** | Pointer to [**V2User**](v2.User.md) |  | [optional] 
**Entitlement** | Pointer to [**PartnerSignupRequestEntitlementOneOf**](PartnerSignupRequestEntitlementOneOf.md) |  | 

## Methods

### NewPartnerSignupRequest

`func NewPartnerSignupRequest(organization PartnerV2Organization, entitlement PartnerSignupRequestEntitlementOneOf, ) *PartnerSignupRequest`

NewPartnerSignupRequest instantiates a new PartnerSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestWithDefaults

`func NewPartnerSignupRequestWithDefaults() *PartnerSignupRequest`

NewPartnerSignupRequestWithDefaults instantiates a new PartnerSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *PartnerSignupRequest) GetOrganization() PartnerV2Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerSignupRequest) GetOrganizationOk() (*PartnerV2Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerSignupRequest) SetOrganization(v PartnerV2Organization)`

SetOrganization sets Organization field to given value.


### GetUser

`func (o *PartnerSignupRequest) GetUser() V2User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PartnerSignupRequest) GetUserOk() (*V2User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PartnerSignupRequest) SetUser(v V2User)`

SetUser sets User field to given value.

### HasUser

`func (o *PartnerSignupRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEntitlement

`func (o *PartnerSignupRequest) GetEntitlement() PartnerSignupRequestEntitlementOneOf`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *PartnerSignupRequest) GetEntitlementOk() (*PartnerSignupRequestEntitlementOneOf, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *PartnerSignupRequest) SetEntitlement(v PartnerSignupRequestEntitlementOneOf)`

SetEntitlement sets Entitlement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


