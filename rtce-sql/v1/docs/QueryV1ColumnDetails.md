# QueryV1ColumnDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the SQL result column. | 
**Type** | [**QueryV1DataType**](QueryV1DataType.md) | The logical type of the column. | 

## Methods

### NewQueryV1ColumnDetails

`func NewQueryV1ColumnDetails(name string, type_ QueryV1DataType, ) *QueryV1ColumnDetails`

NewQueryV1ColumnDetails instantiates a new QueryV1ColumnDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1ColumnDetailsWithDefaults

`func NewQueryV1ColumnDetailsWithDefaults() *QueryV1ColumnDetails`

NewQueryV1ColumnDetailsWithDefaults instantiates a new QueryV1ColumnDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryV1ColumnDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryV1ColumnDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryV1ColumnDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *QueryV1ColumnDetails) GetType() QueryV1DataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryV1ColumnDetails) GetTypeOk() (*QueryV1DataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryV1ColumnDetails) SetType(v QueryV1DataType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


