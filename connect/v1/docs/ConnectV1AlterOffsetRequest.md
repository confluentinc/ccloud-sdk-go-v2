# ConnectV1AlterOffsetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConnectV1AlterOffsetRequestType**](ConnectV1AlterOffsetRequestType.md) |  | 
**Offsets** | Pointer to **[]map[string]interface{}** | Array of offsets which are categorised into partitions. | [optional] 

## Methods

### NewConnectV1AlterOffsetRequest

`func NewConnectV1AlterOffsetRequest(type_ ConnectV1AlterOffsetRequestType, ) *ConnectV1AlterOffsetRequest`

NewConnectV1AlterOffsetRequest instantiates a new ConnectV1AlterOffsetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1AlterOffsetRequestWithDefaults

`func NewConnectV1AlterOffsetRequestWithDefaults() *ConnectV1AlterOffsetRequest`

NewConnectV1AlterOffsetRequestWithDefaults instantiates a new ConnectV1AlterOffsetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectV1AlterOffsetRequest) GetType() ConnectV1AlterOffsetRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectV1AlterOffsetRequest) GetTypeOk() (*ConnectV1AlterOffsetRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectV1AlterOffsetRequest) SetType(v ConnectV1AlterOffsetRequestType)`

SetType sets Type field to given value.


### GetOffsets

`func (o *ConnectV1AlterOffsetRequest) GetOffsets() []map[string]interface{}`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *ConnectV1AlterOffsetRequest) GetOffsetsOk() (*[]map[string]interface{}, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *ConnectV1AlterOffsetRequest) SetOffsets(v []map[string]interface{})`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *ConnectV1AlterOffsetRequest) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


