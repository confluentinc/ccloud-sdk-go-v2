# ColumnDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the SQL table column. | 
**Type** | [**DataType**](DataType.md) | JSON object in TableSchema format; describes the data returned by the results serving API. | 

## Methods

### NewColumnDetails

`func NewColumnDetails(name string, type_ DataType, ) *ColumnDetails`

NewColumnDetails instantiates a new ColumnDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnDetailsWithDefaults

`func NewColumnDetailsWithDefaults() *ColumnDetails`

NewColumnDetailsWithDefaults instantiates a new ColumnDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ColumnDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ColumnDetails) GetType() DataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnDetails) GetTypeOk() (*DataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnDetails) SetType(v DataType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


