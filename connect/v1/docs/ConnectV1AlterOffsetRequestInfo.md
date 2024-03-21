# ConnectV1AlterOffsetRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the connector. | 
**Name** | **string** | The name of the connector. | 
**Offsets** | Pointer to [**ConnectV1Offsets**](ConnectV1Offsets.md) |  | [optional] 
**RequestedAt** | **time.Time** | The time at which the request was made. The time is in UTC, ISO 8601 format. | [readonly] 
**Type** | [**ConnectV1AlterOffsetRequestType**](ConnectV1AlterOffsetRequestType.md) |  | 

## Methods

### NewConnectV1AlterOffsetRequestInfo

`func NewConnectV1AlterOffsetRequestInfo(id string, name string, requestedAt time.Time, type_ ConnectV1AlterOffsetRequestType, ) *ConnectV1AlterOffsetRequestInfo`

NewConnectV1AlterOffsetRequestInfo instantiates a new ConnectV1AlterOffsetRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1AlterOffsetRequestInfoWithDefaults

`func NewConnectV1AlterOffsetRequestInfoWithDefaults() *ConnectV1AlterOffsetRequestInfo`

NewConnectV1AlterOffsetRequestInfoWithDefaults instantiates a new ConnectV1AlterOffsetRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectV1AlterOffsetRequestInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1AlterOffsetRequestInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1AlterOffsetRequestInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectV1AlterOffsetRequestInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1AlterOffsetRequestInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1AlterOffsetRequestInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOffsets

`func (o *ConnectV1AlterOffsetRequestInfo) GetOffsets() ConnectV1Offsets`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *ConnectV1AlterOffsetRequestInfo) GetOffsetsOk() (*ConnectV1Offsets, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *ConnectV1AlterOffsetRequestInfo) SetOffsets(v ConnectV1Offsets)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *ConnectV1AlterOffsetRequestInfo) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ConnectV1AlterOffsetRequestInfo) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ConnectV1AlterOffsetRequestInfo) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ConnectV1AlterOffsetRequestInfo) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetType

`func (o *ConnectV1AlterOffsetRequestInfo) GetType() ConnectV1AlterOffsetRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectV1AlterOffsetRequestInfo) GetTypeOk() (*ConnectV1AlterOffsetRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectV1AlterOffsetRequestInfo) SetType(v ConnectV1AlterOffsetRequestType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


