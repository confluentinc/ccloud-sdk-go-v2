# AlterConfigBatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlterConfigBatchRequestDataData**](AlterConfigBatchRequestDataData.md) |  | 
**ValidateOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewAlterConfigBatchRequestData

`func NewAlterConfigBatchRequestData(data []AlterConfigBatchRequestDataData, ) *AlterConfigBatchRequestData`

NewAlterConfigBatchRequestData instantiates a new AlterConfigBatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterConfigBatchRequestDataWithDefaults

`func NewAlterConfigBatchRequestDataWithDefaults() *AlterConfigBatchRequestData`

NewAlterConfigBatchRequestDataWithDefaults instantiates a new AlterConfigBatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlterConfigBatchRequestData) GetData() []AlterConfigBatchRequestDataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlterConfigBatchRequestData) GetDataOk() (*[]AlterConfigBatchRequestDataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlterConfigBatchRequestData) SetData(v []AlterConfigBatchRequestDataData)`

SetData sets Data field to given value.


### GetValidateOnly

`func (o *AlterConfigBatchRequestData) GetValidateOnly() bool`

GetValidateOnly returns the ValidateOnly field if non-nil, zero value otherwise.

### GetValidateOnlyOk

`func (o *AlterConfigBatchRequestData) GetValidateOnlyOk() (*bool, bool)`

GetValidateOnlyOk returns a tuple with the ValidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnly

`func (o *AlterConfigBatchRequestData) SetValidateOnly(v bool)`

SetValidateOnly sets ValidateOnly field to given value.

### HasValidateOnly

`func (o *AlterConfigBatchRequestData) HasValidateOnly() bool`

HasValidateOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


