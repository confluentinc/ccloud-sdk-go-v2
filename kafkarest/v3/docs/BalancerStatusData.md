# BalancerStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**Status** | **string** |  | 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**AnyUnevenLoad** | [**Relationship**](Relationship.md) |  | 
**BrokerTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBalancerStatusData

`func NewBalancerStatusData(kind string, metadata ResourceMetadata, clusterId string, status string, anyUnevenLoad Relationship, brokerTasks Relationship, ) *BalancerStatusData`

NewBalancerStatusData instantiates a new BalancerStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancerStatusDataWithDefaults

`func NewBalancerStatusDataWithDefaults() *BalancerStatusData`

NewBalancerStatusDataWithDefaults instantiates a new BalancerStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BalancerStatusData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BalancerStatusData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BalancerStatusData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BalancerStatusData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BalancerStatusData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BalancerStatusData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BalancerStatusData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BalancerStatusData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BalancerStatusData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetStatus

`func (o *BalancerStatusData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalancerStatusData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalancerStatusData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorCode

`func (o *BalancerStatusData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BalancerStatusData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BalancerStatusData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BalancerStatusData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BalancerStatusData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BalancerStatusData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BalancerStatusData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BalancerStatusData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BalancerStatusData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BalancerStatusData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BalancerStatusData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BalancerStatusData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetAnyUnevenLoad

`func (o *BalancerStatusData) GetAnyUnevenLoad() Relationship`

GetAnyUnevenLoad returns the AnyUnevenLoad field if non-nil, zero value otherwise.

### GetAnyUnevenLoadOk

`func (o *BalancerStatusData) GetAnyUnevenLoadOk() (*Relationship, bool)`

GetAnyUnevenLoadOk returns a tuple with the AnyUnevenLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUnevenLoad

`func (o *BalancerStatusData) SetAnyUnevenLoad(v Relationship)`

SetAnyUnevenLoad sets AnyUnevenLoad field to given value.


### GetBrokerTasks

`func (o *BalancerStatusData) GetBrokerTasks() Relationship`

GetBrokerTasks returns the BrokerTasks field if non-nil, zero value otherwise.

### GetBrokerTasksOk

`func (o *BalancerStatusData) GetBrokerTasksOk() (*Relationship, bool)`

GetBrokerTasksOk returns a tuple with the BrokerTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTasks

`func (o *BalancerStatusData) SetBrokerTasks(v Relationship)`

SetBrokerTasks sets BrokerTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


