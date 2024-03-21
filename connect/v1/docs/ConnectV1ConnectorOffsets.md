# ConnectV1ConnectorOffsets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the connector. | [optional] 
**Id** | Pointer to **string** | The ID of the connector. | [optional] 
**Offsets** | Pointer to [**ConnectV1Offsets**](ConnectV1Offsets.md) |  | [optional] 
**Metadata** | Pointer to [**ConnectV1ConnectorOffsetsMetadata**](ConnectV1ConnectorOffsetsMetadata.md) |  | [optional] 

## Methods

### NewConnectV1ConnectorOffsets

`func NewConnectV1ConnectorOffsets() *ConnectV1ConnectorOffsets`

NewConnectV1ConnectorOffsets instantiates a new ConnectV1ConnectorOffsets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorOffsetsWithDefaults

`func NewConnectV1ConnectorOffsetsWithDefaults() *ConnectV1ConnectorOffsets`

NewConnectV1ConnectorOffsetsWithDefaults instantiates a new ConnectV1ConnectorOffsets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectV1ConnectorOffsets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorOffsets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorOffsets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectV1ConnectorOffsets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ConnectV1ConnectorOffsets) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectV1ConnectorOffsets) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectV1ConnectorOffsets) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectV1ConnectorOffsets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOffsets

`func (o *ConnectV1ConnectorOffsets) GetOffsets() ConnectV1Offsets`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *ConnectV1ConnectorOffsets) GetOffsetsOk() (*ConnectV1Offsets, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *ConnectV1ConnectorOffsets) SetOffsets(v ConnectV1Offsets)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *ConnectV1ConnectorOffsets) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.

### GetMetadata

`func (o *ConnectV1ConnectorOffsets) GetMetadata() ConnectV1ConnectorOffsetsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectV1ConnectorOffsets) GetMetadataOk() (*ConnectV1ConnectorOffsetsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectV1ConnectorOffsets) SetMetadata(v ConnectV1ConnectorOffsetsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectV1ConnectorOffsets) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


