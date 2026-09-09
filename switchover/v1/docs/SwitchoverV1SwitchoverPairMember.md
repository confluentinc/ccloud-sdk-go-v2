# SwitchoverV1SwitchoverPairMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A logical name for this member (e.g. \&quot;west\&quot; or \&quot;east\&quot;), unique within the pair. | 
**MemberCrn** | **string** | The Confluent Resource Name (CRN) of the resource this member represents.  A CRN rather than an ID plus an environment because member types do not share a parent hierarchy: a Kafka cluster is &#x60;environment → cloud-cluster&#x60;, a Flink compute pool is &#x60;environment → flink-region → compute-pool&#x60;, and a Connect connector is &#x60;environment → cloud-cluster → connector&#x60; with a *name* as its leaf, unique only within its cluster. A single &#x60;environment&#x60; field cannot address the latter two unambiguously.  The environment is read from the CRN, so a pair whose two members live in different environments needs no separate per-member environment field.  For a Kafka cluster both the canonical form and the variant carrying the redundant trailing &#x60;/kafka&#x3D;&lt;lkc&gt;&#x60; (which the cmk API returns in &#x60;metadata.resource_name&#x60;) are accepted. The short &#x60;crn://confluent.cloud/kafka&#x3D;&lt;lkc&gt;&#x60; form is rejected: it is a valid CRN but carries no environment. In Terraform the value is available as &#x60;confluent_kafka_cluster.&lt;name&gt;.rbac_crn&#x60;.  | 
**Location** | Pointer to [**SwitchoverV1SwitchoverPairMemberLocation**](SwitchoverV1SwitchoverPairMemberLocation.md) |  | [optional] 

## Methods

### NewSwitchoverV1SwitchoverPairMember

`func NewSwitchoverV1SwitchoverPairMember(name string, memberCrn string, ) *SwitchoverV1SwitchoverPairMember`

NewSwitchoverV1SwitchoverPairMember instantiates a new SwitchoverV1SwitchoverPairMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairMemberWithDefaults

`func NewSwitchoverV1SwitchoverPairMemberWithDefaults() *SwitchoverV1SwitchoverPairMember`

NewSwitchoverV1SwitchoverPairMemberWithDefaults instantiates a new SwitchoverV1SwitchoverPairMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwitchoverV1SwitchoverPairMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchoverV1SwitchoverPairMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchoverV1SwitchoverPairMember) SetName(v string)`

SetName sets Name field to given value.


### GetMemberCrn

`func (o *SwitchoverV1SwitchoverPairMember) GetMemberCrn() string`

GetMemberCrn returns the MemberCrn field if non-nil, zero value otherwise.

### GetMemberCrnOk

`func (o *SwitchoverV1SwitchoverPairMember) GetMemberCrnOk() (*string, bool)`

GetMemberCrnOk returns a tuple with the MemberCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCrn

`func (o *SwitchoverV1SwitchoverPairMember) SetMemberCrn(v string)`

SetMemberCrn sets MemberCrn field to given value.


### GetLocation

`func (o *SwitchoverV1SwitchoverPairMember) GetLocation() SwitchoverV1SwitchoverPairMemberLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SwitchoverV1SwitchoverPairMember) GetLocationOk() (*SwitchoverV1SwitchoverPairMemberLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SwitchoverV1SwitchoverPairMember) SetLocation(v SwitchoverV1SwitchoverPairMemberLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SwitchoverV1SwitchoverPairMember) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


