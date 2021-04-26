# ActivatePartnerSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**V2User**](v2.User.md) |  | 
**OrganizationId** | Pointer to **string** | The ID of the organization | 

## Methods

### NewActivatePartnerSignupRequest

`func NewActivatePartnerSignupRequest(user V2User, organizationId string, ) *ActivatePartnerSignupRequest`

NewActivatePartnerSignupRequest instantiates a new ActivatePartnerSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivatePartnerSignupRequestWithDefaults

`func NewActivatePartnerSignupRequestWithDefaults() *ActivatePartnerSignupRequest`

NewActivatePartnerSignupRequestWithDefaults instantiates a new ActivatePartnerSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ActivatePartnerSignupRequest) GetUser() V2User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActivatePartnerSignupRequest) GetUserOk() (*V2User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActivatePartnerSignupRequest) SetUser(v V2User)`

SetUser sets User field to given value.


### GetOrganizationId

`func (o *ActivatePartnerSignupRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ActivatePartnerSignupRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ActivatePartnerSignupRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


