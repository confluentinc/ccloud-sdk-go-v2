# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Message to be shown in the notification to the customer. | 
**OrgIds** | **[]int32** |  | 
**DryRun** | **bool** | Indicates the intent of caller whether to send the notification to actual customers (if it&#39;s false) or to just see the output shown to the customers (if it&#39;s true).  | 

## Methods

### NewInlineObject

`func NewInlineObject(message string, orgIds []int32, dryRun bool, ) *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgIds

`func (o *InlineObject) GetOrgIds() []int32`

GetOrgIds returns the OrgIds field if non-nil, zero value otherwise.

### GetOrgIdsOk

`func (o *InlineObject) GetOrgIdsOk() (*[]int32, bool)`

GetOrgIdsOk returns a tuple with the OrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIds

`func (o *InlineObject) SetOrgIds(v []int32)`

SetOrgIds sets OrgIds field to given value.


### GetDryRun

`func (o *InlineObject) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *InlineObject) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *InlineObject) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


