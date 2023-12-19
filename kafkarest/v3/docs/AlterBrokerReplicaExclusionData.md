# AlterBrokerReplicaExclusionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Exclusion** | **string** |  | 
**Reason** | **string** |  | 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewAlterBrokerReplicaExclusionData

`func NewAlterBrokerReplicaExclusionData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, exclusion string, reason string, broker Relationship, ) *AlterBrokerReplicaExclusionData`

NewAlterBrokerReplicaExclusionData instantiates a new AlterBrokerReplicaExclusionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterBrokerReplicaExclusionDataWithDefaults

`func NewAlterBrokerReplicaExclusionDataWithDefaults() *AlterBrokerReplicaExclusionData`

NewAlterBrokerReplicaExclusionDataWithDefaults instantiates a new AlterBrokerReplicaExclusionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AlterBrokerReplicaExclusionData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlterBrokerReplicaExclusionData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlterBrokerReplicaExclusionData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *AlterBrokerReplicaExclusionData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlterBrokerReplicaExclusionData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlterBrokerReplicaExclusionData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *AlterBrokerReplicaExclusionData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AlterBrokerReplicaExclusionData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AlterBrokerReplicaExclusionData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *AlterBrokerReplicaExclusionData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *AlterBrokerReplicaExclusionData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *AlterBrokerReplicaExclusionData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetExclusion

`func (o *AlterBrokerReplicaExclusionData) GetExclusion() string`

GetExclusion returns the Exclusion field if non-nil, zero value otherwise.

### GetExclusionOk

`func (o *AlterBrokerReplicaExclusionData) GetExclusionOk() (*string, bool)`

GetExclusionOk returns a tuple with the Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusion

`func (o *AlterBrokerReplicaExclusionData) SetExclusion(v string)`

SetExclusion sets Exclusion field to given value.


### GetReason

`func (o *AlterBrokerReplicaExclusionData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlterBrokerReplicaExclusionData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlterBrokerReplicaExclusionData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetErrorCode

`func (o *AlterBrokerReplicaExclusionData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AlterBrokerReplicaExclusionData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AlterBrokerReplicaExclusionData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AlterBrokerReplicaExclusionData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AlterBrokerReplicaExclusionData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AlterBrokerReplicaExclusionData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *AlterBrokerReplicaExclusionData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AlterBrokerReplicaExclusionData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AlterBrokerReplicaExclusionData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AlterBrokerReplicaExclusionData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AlterBrokerReplicaExclusionData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AlterBrokerReplicaExclusionData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetBroker

`func (o *AlterBrokerReplicaExclusionData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *AlterBrokerReplicaExclusionData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *AlterBrokerReplicaExclusionData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


