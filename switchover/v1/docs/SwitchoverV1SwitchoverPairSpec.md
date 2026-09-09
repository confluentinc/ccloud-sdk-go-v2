# SwitchoverV1SwitchoverPairSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentCrn** | Pointer to **string** | The CRN of the environment this resource belongs to.  Every reference in a request body is a CRN (see &#x60;member_crn&#x60;), so this is one too rather than an &#x60;{id}&#x60; relationship object — callers learn a single convention, and the organization always travels with the reference. A CRN that names no environment is rejected; it is not defaulted.  The &#x60;?environment&#x3D;&#x60; query parameter on the other operations remains a bare environment ID: it is the standard Confluent collection filter, not a reference to a stored resource.  | [optional] 
**DisplayName** | Pointer to **string** | A human-readable name for the switchover pair. | [optional] 
**Members** | Pointer to [**[]SwitchoverV1SwitchoverPairMember**](SwitchoverV1SwitchoverPairMember.md) | The two clusters participating in this switchover pair. Must contain exactly 2 members.  | [optional] 
**ActiveMember** | Pointer to **string** | The name of the member that is currently active. On create, this selects which member starts as active; it must match one of the &#x60;members[].name&#x60; values. Use the &#x60;:failover&#x60; operation to change the active member after creation.  | [optional] 
**FirstActive** | Pointer to **string** | The member that was active when the pair was created. Output-only and immutable; set implicitly at create time and never accepted as input. Comparing it with &#x60;active_member&#x60; distinguishes a pair that has failed over (they differ) from one still on its original active member (they match).  | [optional] [readonly] 
**FailoverType** | Pointer to **string** | The failover semantics most recently applied to this pair via the &#x60;:failover&#x60; operation. Not settable directly; empty until a failover has been triggered.  | [optional] [readonly] 

## Methods

### NewSwitchoverV1SwitchoverPairSpec

`func NewSwitchoverV1SwitchoverPairSpec() *SwitchoverV1SwitchoverPairSpec`

NewSwitchoverV1SwitchoverPairSpec instantiates a new SwitchoverV1SwitchoverPairSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairSpecWithDefaults

`func NewSwitchoverV1SwitchoverPairSpecWithDefaults() *SwitchoverV1SwitchoverPairSpec`

NewSwitchoverV1SwitchoverPairSpecWithDefaults instantiates a new SwitchoverV1SwitchoverPairSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverPairSpec) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverPairSpec) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.

### HasEnvironmentCrn

`func (o *SwitchoverV1SwitchoverPairSpec) HasEnvironmentCrn() bool`

HasEnvironmentCrn returns a boolean if a field has been set.

### GetDisplayName

`func (o *SwitchoverV1SwitchoverPairSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SwitchoverV1SwitchoverPairSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SwitchoverV1SwitchoverPairSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMembers

`func (o *SwitchoverV1SwitchoverPairSpec) GetMembers() []SwitchoverV1SwitchoverPairMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetMembersOk() (*[]SwitchoverV1SwitchoverPairMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SwitchoverV1SwitchoverPairSpec) SetMembers(v []SwitchoverV1SwitchoverPairMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SwitchoverV1SwitchoverPairSpec) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetActiveMember

`func (o *SwitchoverV1SwitchoverPairSpec) GetActiveMember() string`

GetActiveMember returns the ActiveMember field if non-nil, zero value otherwise.

### GetActiveMemberOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetActiveMemberOk() (*string, bool)`

GetActiveMemberOk returns a tuple with the ActiveMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMember

`func (o *SwitchoverV1SwitchoverPairSpec) SetActiveMember(v string)`

SetActiveMember sets ActiveMember field to given value.

### HasActiveMember

`func (o *SwitchoverV1SwitchoverPairSpec) HasActiveMember() bool`

HasActiveMember returns a boolean if a field has been set.

### GetFirstActive

`func (o *SwitchoverV1SwitchoverPairSpec) GetFirstActive() string`

GetFirstActive returns the FirstActive field if non-nil, zero value otherwise.

### GetFirstActiveOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetFirstActiveOk() (*string, bool)`

GetFirstActiveOk returns a tuple with the FirstActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstActive

`func (o *SwitchoverV1SwitchoverPairSpec) SetFirstActive(v string)`

SetFirstActive sets FirstActive field to given value.

### HasFirstActive

`func (o *SwitchoverV1SwitchoverPairSpec) HasFirstActive() bool`

HasFirstActive returns a boolean if a field has been set.

### GetFailoverType

`func (o *SwitchoverV1SwitchoverPairSpec) GetFailoverType() string`

GetFailoverType returns the FailoverType field if non-nil, zero value otherwise.

### GetFailoverTypeOk

`func (o *SwitchoverV1SwitchoverPairSpec) GetFailoverTypeOk() (*string, bool)`

GetFailoverTypeOk returns a tuple with the FailoverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverType

`func (o *SwitchoverV1SwitchoverPairSpec) SetFailoverType(v string)`

SetFailoverType sets FailoverType field to given value.

### HasFailoverType

`func (o *SwitchoverV1SwitchoverPairSpec) HasFailoverType() bool`

HasFailoverType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


