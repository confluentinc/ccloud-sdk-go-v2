# SqlV1alpha1ErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for this particular occurrence of the problem. | [optional] 
**Status** | Pointer to **string** | The HTTP status code applicable to this problem, expressed as a string value. | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem. It **SHOULD NOT** change from occurrence to occurrence of the problem, except for purposes of localization. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Resolution** | Pointer to **string** | Instructions for the end-user for correcting the error. | [optional] 

## Methods

### NewSqlV1alpha1ErrorDetails

`func NewSqlV1alpha1ErrorDetails() *SqlV1alpha1ErrorDetails`

NewSqlV1alpha1ErrorDetails instantiates a new SqlV1alpha1ErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1alpha1ErrorDetailsWithDefaults

`func NewSqlV1alpha1ErrorDetailsWithDefaults() *SqlV1alpha1ErrorDetails`

NewSqlV1alpha1ErrorDetailsWithDefaults instantiates a new SqlV1alpha1ErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SqlV1alpha1ErrorDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlV1alpha1ErrorDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlV1alpha1ErrorDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SqlV1alpha1ErrorDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SqlV1alpha1ErrorDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1alpha1ErrorDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1alpha1ErrorDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1alpha1ErrorDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *SqlV1alpha1ErrorDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SqlV1alpha1ErrorDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SqlV1alpha1ErrorDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SqlV1alpha1ErrorDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *SqlV1alpha1ErrorDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SqlV1alpha1ErrorDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SqlV1alpha1ErrorDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SqlV1alpha1ErrorDetails) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetResolution

`func (o *SqlV1alpha1ErrorDetails) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *SqlV1alpha1ErrorDetails) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *SqlV1alpha1ErrorDetails) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *SqlV1alpha1ErrorDetails) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


