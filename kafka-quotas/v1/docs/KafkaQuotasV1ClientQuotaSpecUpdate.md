# KafkaQuotasV1ClientQuotaSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the client quota. | [optional] 
**Description** | Pointer to **string** | A human readable description for the client quota. | [optional] 
**Throughput** | Pointer to [**KafkaQuotasV1Throughput**](kafka-quotas.v1.Throughput.md) | Throughput for the client quota. | [optional] 
**Principals** | Pointer to [**[]GlobalObjectReference**](GlobalObjectReference.md) | A list of principals to apply a client quota to. Use \&quot;&lt;default&gt;\&quot; to apply a client quota to all service accounts (see [Control application usage with Client Quotas](https://docs.confluent.io/cloud/current/clusters/client-quotas.html#control-application-usage-with-client-quotas) for more details).  | [optional] 

## Methods

### NewKafkaQuotasV1ClientQuotaSpecUpdate

`func NewKafkaQuotasV1ClientQuotaSpecUpdate() *KafkaQuotasV1ClientQuotaSpecUpdate`

NewKafkaQuotasV1ClientQuotaSpecUpdate instantiates a new KafkaQuotasV1ClientQuotaSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaQuotasV1ClientQuotaSpecUpdateWithDefaults

`func NewKafkaQuotasV1ClientQuotaSpecUpdateWithDefaults() *KafkaQuotasV1ClientQuotaSpecUpdate`

NewKafkaQuotasV1ClientQuotaSpecUpdateWithDefaults instantiates a new KafkaQuotasV1ClientQuotaSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetThroughput

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetThroughput() KafkaQuotasV1Throughput`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetThroughputOk() (*KafkaQuotasV1Throughput, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) SetThroughput(v KafkaQuotasV1Throughput)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetPrincipals() []GlobalObjectReference`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) GetPrincipalsOk() (*[]GlobalObjectReference, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) SetPrincipals(v []GlobalObjectReference)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *KafkaQuotasV1ClientQuotaSpecUpdate) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


