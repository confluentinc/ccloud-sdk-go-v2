# BalancerStatusDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | **string** |  | 
**ErrorCode** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**AnyUnevenLoad** | [**Relationship**](Relationship.md) |  | 
**BrokerTasks** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBalancerStatusDataAllOf

`func NewBalancerStatusDataAllOf(clusterId string, status string, anyUnevenLoad Relationship, brokerTasks Relationship, ) *BalancerStatusDataAllOf`

NewBalancerStatusDataAllOf instantiates a new BalancerStatusDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancerStatusDataAllOfWithDefaults

`func NewBalancerStatusDataAllOfWithDefaults() *BalancerStatusDataAllOf`

NewBalancerStatusDataAllOfWithDefaults instantiates a new BalancerStatusDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BalancerStatusDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BalancerStatusDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BalancerStatusDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetStatus

`func (o *BalancerStatusDataAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalancerStatusDataAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalancerStatusDataAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorCode

`func (o *BalancerStatusDataAllOf) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BalancerStatusDataAllOf) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BalancerStatusDataAllOf) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BalancerStatusDataAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BalancerStatusDataAllOf) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BalancerStatusDataAllOf) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BalancerStatusDataAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BalancerStatusDataAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BalancerStatusDataAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BalancerStatusDataAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BalancerStatusDataAllOf) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BalancerStatusDataAllOf) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetAnyUnevenLoad

`func (o *BalancerStatusDataAllOf) GetAnyUnevenLoad() Relationship`

GetAnyUnevenLoad returns the AnyUnevenLoad field if non-nil, zero value otherwise.

### GetAnyUnevenLoadOk

`func (o *BalancerStatusDataAllOf) GetAnyUnevenLoadOk() (*Relationship, bool)`

GetAnyUnevenLoadOk returns a tuple with the AnyUnevenLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUnevenLoad

`func (o *BalancerStatusDataAllOf) SetAnyUnevenLoad(v Relationship)`

SetAnyUnevenLoad sets AnyUnevenLoad field to given value.


### GetBrokerTasks

`func (o *BalancerStatusDataAllOf) GetBrokerTasks() Relationship`

GetBrokerTasks returns the BrokerTasks field if non-nil, zero value otherwise.

### GetBrokerTasksOk

`func (o *BalancerStatusDataAllOf) GetBrokerTasksOk() (*Relationship, bool)`

GetBrokerTasksOk returns a tuple with the BrokerTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTasks

`func (o *BalancerStatusDataAllOf) SetBrokerTasks(v Relationship)`

SetBrokerTasks sets BrokerTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


