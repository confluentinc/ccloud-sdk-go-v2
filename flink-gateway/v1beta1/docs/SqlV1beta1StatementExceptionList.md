# SqlV1beta1StatementExceptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | 
**Kind** | **string** | Kind defines the object this REST resource represents. | 
**Metadata** | [**ExceptionListMeta**](ExceptionListMeta.md) |  | 
**Data** | [**[]SqlV1beta1StatementException**](SqlV1beta1StatementException.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewSqlV1beta1StatementExceptionList

`func NewSqlV1beta1StatementExceptionList(apiVersion string, kind string, metadata ExceptionListMeta, data []SqlV1beta1StatementException, ) *SqlV1beta1StatementExceptionList`

NewSqlV1beta1StatementExceptionList instantiates a new SqlV1beta1StatementExceptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1beta1StatementExceptionListWithDefaults

`func NewSqlV1beta1StatementExceptionListWithDefaults() *SqlV1beta1StatementExceptionList`

NewSqlV1beta1StatementExceptionListWithDefaults instantiates a new SqlV1beta1StatementExceptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1beta1StatementExceptionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1beta1StatementExceptionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1beta1StatementExceptionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlV1beta1StatementExceptionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1beta1StatementExceptionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1beta1StatementExceptionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlV1beta1StatementExceptionList) GetMetadata() ExceptionListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1beta1StatementExceptionList) GetMetadataOk() (*ExceptionListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1beta1StatementExceptionList) SetMetadata(v ExceptionListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *SqlV1beta1StatementExceptionList) GetData() []SqlV1beta1StatementException`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1beta1StatementExceptionList) GetDataOk() (*[]SqlV1beta1StatementException, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1beta1StatementExceptionList) SetData(v []SqlV1beta1StatementException)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


