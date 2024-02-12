# TableIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | [**Namespace**](Namespace.md) |  | 
**Name** | **string** |  | 

## Methods

### NewTableIdentifier

`func NewTableIdentifier(namespace Namespace, name string, ) *TableIdentifier`

NewTableIdentifier instantiates a new TableIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableIdentifierWithDefaults

`func NewTableIdentifierWithDefaults() *TableIdentifier`

NewTableIdentifierWithDefaults instantiates a new TableIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *TableIdentifier) GetNamespace() Namespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TableIdentifier) GetNamespaceOk() (*Namespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TableIdentifier) SetNamespace(v Namespace)`

SetNamespace sets Namespace field to given value.


### GetName

`func (o *TableIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableIdentifier) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


