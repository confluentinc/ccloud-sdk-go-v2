# ExporterStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of exporter. | [optional] 
**State** | Pointer to **string** | State of the exporter. Could be STARTING, RUNNING or PAUSED | [optional] 
**Offset** | Pointer to **int64** | Offset of the exporter | [optional] 
**Ts** | Pointer to **int64** | Timestamp of the exporter | [optional] 
**Trace** | Pointer to **string** | Error trace of the exporter | [optional] 

## Methods

### NewExporterStatusResponse

`func NewExporterStatusResponse() *ExporterStatusResponse`

NewExporterStatusResponse instantiates a new ExporterStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterStatusResponseWithDefaults

`func NewExporterStatusResponseWithDefaults() *ExporterStatusResponse`

NewExporterStatusResponseWithDefaults instantiates a new ExporterStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExporterStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExporterStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExporterStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExporterStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ExporterStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExporterStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExporterStatusResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExporterStatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOffset

`func (o *ExporterStatusResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ExporterStatusResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ExporterStatusResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ExporterStatusResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTs

`func (o *ExporterStatusResponse) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ExporterStatusResponse) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ExporterStatusResponse) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ExporterStatusResponse) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTrace

`func (o *ExporterStatusResponse) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ExporterStatusResponse) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ExporterStatusResponse) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ExporterStatusResponse) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


