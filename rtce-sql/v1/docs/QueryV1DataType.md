# QueryV1DataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The Flink logical type name of the column or field. | 
**Nullable** | **bool** | Whether values of this column or field can be null. | 
**Length** | Pointer to **int32** | Declared length of &#x60;CHAR&#x60;, &#x60;VARCHAR&#x60;, &#x60;BINARY&#x60; and &#x60;VARBINARY&#x60;. Unbounded &#x60;VARCHAR&#x60; and &#x60;VARBINARY&#x60; report 2147483647.  | [optional] 
**Precision** | Pointer to **int32** | Declared precision of &#x60;DECIMAL&#x60;, the &#x60;TIME&#x60;/&#x60;TIMESTAMP&#x60; types and the &#x60;INTERVAL_*&#x60; types.  | [optional] 
**Scale** | Pointer to **int32** | Declared scale of &#x60;DECIMAL&#x60;. | [optional] 
**KeyType** | Pointer to [**QueryV1DataType**](QueryV1DataType.md) | Key type of a &#x60;MAP&#x60;. | [optional] 
**ValueType** | Pointer to [**QueryV1DataType**](QueryV1DataType.md) | Value type of a &#x60;MAP&#x60;. | [optional] 
**ElementType** | Pointer to [**QueryV1DataType**](QueryV1DataType.md) | Element type of an &#x60;ARRAY&#x60; or &#x60;MULTISET&#x60;. | [optional] 
**Fields** | Pointer to [**[]QueryV1RowFieldType**](QueryV1RowFieldType.md) | Fields of a &#x60;ROW&#x60;, in declaration order. | [optional] 
**Resolution** | Pointer to **string** | Interval resolution, for example &#x60;YEAR_TO_MONTH&#x60; for &#x60;INTERVAL_YEAR_MONTH&#x60; or &#x60;DAY_TO_SECOND&#x60; for &#x60;INTERVAL_DAY_TIME&#x60;.  | [optional] 
**FractionalPrecision** | Pointer to **int32** | Fractional-second precision of &#x60;INTERVAL_DAY_TIME&#x60;. | [optional] 
**ClassName** | Pointer to **string** | Class name of a structured type. Present in the Flink type model; the engine does not currently emit structured types.  | [optional] 

## Methods

### NewQueryV1DataType

`func NewQueryV1DataType(type_ string, nullable bool, ) *QueryV1DataType`

NewQueryV1DataType instantiates a new QueryV1DataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1DataTypeWithDefaults

`func NewQueryV1DataTypeWithDefaults() *QueryV1DataType`

NewQueryV1DataTypeWithDefaults instantiates a new QueryV1DataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryV1DataType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryV1DataType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryV1DataType) SetType(v string)`

SetType sets Type field to given value.


### GetNullable

`func (o *QueryV1DataType) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *QueryV1DataType) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *QueryV1DataType) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetLength

`func (o *QueryV1DataType) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *QueryV1DataType) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *QueryV1DataType) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *QueryV1DataType) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPrecision

`func (o *QueryV1DataType) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *QueryV1DataType) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *QueryV1DataType) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *QueryV1DataType) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetScale

`func (o *QueryV1DataType) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *QueryV1DataType) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *QueryV1DataType) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *QueryV1DataType) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetKeyType

`func (o *QueryV1DataType) GetKeyType() QueryV1DataType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *QueryV1DataType) GetKeyTypeOk() (*QueryV1DataType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *QueryV1DataType) SetKeyType(v QueryV1DataType)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *QueryV1DataType) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetValueType

`func (o *QueryV1DataType) GetValueType() QueryV1DataType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *QueryV1DataType) GetValueTypeOk() (*QueryV1DataType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *QueryV1DataType) SetValueType(v QueryV1DataType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *QueryV1DataType) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetElementType

`func (o *QueryV1DataType) GetElementType() QueryV1DataType`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *QueryV1DataType) GetElementTypeOk() (*QueryV1DataType, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *QueryV1DataType) SetElementType(v QueryV1DataType)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *QueryV1DataType) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetFields

`func (o *QueryV1DataType) GetFields() []QueryV1RowFieldType`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *QueryV1DataType) GetFieldsOk() (*[]QueryV1RowFieldType, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *QueryV1DataType) SetFields(v []QueryV1RowFieldType)`

SetFields sets Fields field to given value.

### HasFields

`func (o *QueryV1DataType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetResolution

`func (o *QueryV1DataType) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *QueryV1DataType) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *QueryV1DataType) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *QueryV1DataType) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetFractionalPrecision

`func (o *QueryV1DataType) GetFractionalPrecision() int32`

GetFractionalPrecision returns the FractionalPrecision field if non-nil, zero value otherwise.

### GetFractionalPrecisionOk

`func (o *QueryV1DataType) GetFractionalPrecisionOk() (*int32, bool)`

GetFractionalPrecisionOk returns a tuple with the FractionalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionalPrecision

`func (o *QueryV1DataType) SetFractionalPrecision(v int32)`

SetFractionalPrecision sets FractionalPrecision field to given value.

### HasFractionalPrecision

`func (o *QueryV1DataType) HasFractionalPrecision() bool`

HasFractionalPrecision returns a boolean if a field has been set.

### GetClassName

`func (o *QueryV1DataType) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *QueryV1DataType) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *QueryV1DataType) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *QueryV1DataType) HasClassName() bool`

HasClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


