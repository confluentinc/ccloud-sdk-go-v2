# IamV2JwksStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwksStatus** | Pointer to **string** | The actual state of the public key data | [optional] 
**JwksLastRefreshAt** | Pointer to **time.Time** | The last successful refresh time for the public key data | [optional] 

## Methods

### NewIamV2JwksStatus

`func NewIamV2JwksStatus() *IamV2JwksStatus`

NewIamV2JwksStatus instantiates a new IamV2JwksStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2JwksStatusWithDefaults

`func NewIamV2JwksStatusWithDefaults() *IamV2JwksStatus`

NewIamV2JwksStatusWithDefaults instantiates a new IamV2JwksStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwksStatus

`func (o *IamV2JwksStatus) GetJwksStatus() string`

GetJwksStatus returns the JwksStatus field if non-nil, zero value otherwise.

### GetJwksStatusOk

`func (o *IamV2JwksStatus) GetJwksStatusOk() (*string, bool)`

GetJwksStatusOk returns a tuple with the JwksStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksStatus

`func (o *IamV2JwksStatus) SetJwksStatus(v string)`

SetJwksStatus sets JwksStatus field to given value.

### HasJwksStatus

`func (o *IamV2JwksStatus) HasJwksStatus() bool`

HasJwksStatus returns a boolean if a field has been set.

### GetJwksLastRefreshAt

`func (o *IamV2JwksStatus) GetJwksLastRefreshAt() time.Time`

GetJwksLastRefreshAt returns the JwksLastRefreshAt field if non-nil, zero value otherwise.

### GetJwksLastRefreshAtOk

`func (o *IamV2JwksStatus) GetJwksLastRefreshAtOk() (*time.Time, bool)`

GetJwksLastRefreshAtOk returns a tuple with the JwksLastRefreshAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksLastRefreshAt

`func (o *IamV2JwksStatus) SetJwksLastRefreshAt(v time.Time)`

SetJwksLastRefreshAt sets JwksLastRefreshAt field to given value.

### HasJwksLastRefreshAt

`func (o *IamV2JwksStatus) HasJwksLastRefreshAt() bool`

HasJwksLastRefreshAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


