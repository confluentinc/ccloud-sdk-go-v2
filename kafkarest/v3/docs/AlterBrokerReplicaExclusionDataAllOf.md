# AlterBrokerReplicaExclusionDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Exclusion** | **string** |  | 
**Reason** | **string** |  | 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewAlterBrokerReplicaExclusionDataAllOf

`func NewAlterBrokerReplicaExclusionDataAllOf(clusterId string, brokerId int32, exclusion string, reason string, broker Relationship, ) *AlterBrokerReplicaExclusionDataAllOf`

NewAlterBrokerReplicaExclusionDataAllOf instantiates a new AlterBrokerReplicaExclusionDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterBrokerReplicaExclusionDataAllOfWithDefaults

`func NewAlterBrokerReplicaExclusionDataAllOfWithDefaults() *AlterBrokerReplicaExclusionDataAllOf`

NewAlterBrokerReplicaExclusionDataAllOfWithDefaults instantiates a new AlterBrokerReplicaExclusionDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetExclusion

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetExclusion() string`

GetExclusion returns the Exclusion field if non-nil, zero value otherwise.

### GetExclusionOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetExclusionOk() (*string, bool)`

GetExclusionOk returns a tuple with the Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusion

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetExclusion(v string)`

SetExclusion sets Exclusion field to given value.


### GetReason

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetReason(v string)`

SetReason sets Reason field to given value.


### GetErrorCode

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AlterBrokerReplicaExclusionDataAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AlterBrokerReplicaExclusionDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AlterBrokerReplicaExclusionDataAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AlterBrokerReplicaExclusionDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *AlterBrokerReplicaExclusionDataAllOf) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *AlterBrokerReplicaExclusionDataAllOf) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


