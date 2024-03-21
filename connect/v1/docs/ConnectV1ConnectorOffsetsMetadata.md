# ConnectV1ConnectorOffsetsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservedAt** | Pointer to **time.Time** | The time at which the offsets were observed. The time is in UTC, ISO 8601 format. | [optional] [readonly] 

## Methods

### NewConnectV1ConnectorOffsetsMetadata

`func NewConnectV1ConnectorOffsetsMetadata() *ConnectV1ConnectorOffsetsMetadata`

NewConnectV1ConnectorOffsetsMetadata instantiates a new ConnectV1ConnectorOffsetsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorOffsetsMetadataWithDefaults

`func NewConnectV1ConnectorOffsetsMetadataWithDefaults() *ConnectV1ConnectorOffsetsMetadata`

NewConnectV1ConnectorOffsetsMetadataWithDefaults instantiates a new ConnectV1ConnectorOffsetsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservedAt

`func (o *ConnectV1ConnectorOffsetsMetadata) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ConnectV1ConnectorOffsetsMetadata) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ConnectV1ConnectorOffsetsMetadata) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *ConnectV1ConnectorOffsetsMetadata) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


