# AlterMirrorsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MirrorTopicNames** | Pointer to **[]string** | The mirror topics specified as a list of topic names. | [optional] 
**MirrorTopicNamePattern** | Pointer to **string** | The mirror topics specified as a pattern. | [optional] 

## Methods

### NewAlterMirrorsRequestData

`func NewAlterMirrorsRequestData() *AlterMirrorsRequestData`

NewAlterMirrorsRequestData instantiates a new AlterMirrorsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterMirrorsRequestDataWithDefaults

`func NewAlterMirrorsRequestDataWithDefaults() *AlterMirrorsRequestData`

NewAlterMirrorsRequestDataWithDefaults instantiates a new AlterMirrorsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMirrorTopicNames

`func (o *AlterMirrorsRequestData) GetMirrorTopicNames() []string`

GetMirrorTopicNames returns the MirrorTopicNames field if non-nil, zero value otherwise.

### GetMirrorTopicNamesOk

`func (o *AlterMirrorsRequestData) GetMirrorTopicNamesOk() (*[]string, bool)`

GetMirrorTopicNamesOk returns a tuple with the MirrorTopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicNames

`func (o *AlterMirrorsRequestData) SetMirrorTopicNames(v []string)`

SetMirrorTopicNames sets MirrorTopicNames field to given value.

### HasMirrorTopicNames

`func (o *AlterMirrorsRequestData) HasMirrorTopicNames() bool`

HasMirrorTopicNames returns a boolean if a field has been set.

### GetMirrorTopicNamePattern

`func (o *AlterMirrorsRequestData) GetMirrorTopicNamePattern() string`

GetMirrorTopicNamePattern returns the MirrorTopicNamePattern field if non-nil, zero value otherwise.

### GetMirrorTopicNamePatternOk

`func (o *AlterMirrorsRequestData) GetMirrorTopicNamePatternOk() (*string, bool)`

GetMirrorTopicNamePatternOk returns a tuple with the MirrorTopicNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicNamePattern

`func (o *AlterMirrorsRequestData) SetMirrorTopicNamePattern(v string)`

SetMirrorTopicNamePattern sets MirrorTopicNamePattern field to given value.

### HasMirrorTopicNamePattern

`func (o *AlterMirrorsRequestData) HasMirrorTopicNamePattern() bool`

HasMirrorTopicNamePattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


