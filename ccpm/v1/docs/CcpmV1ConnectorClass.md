# CcpmV1ConnectorClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** | Java class or alias for connector. You can get connector class from connector documentation provided by developer. | 
**Type** | **string** | Type of the connector class. Should be either &#x60;SOURCE&#x60; or &#x60;SINK&#x60;.  | 

## Methods

### NewCcpmV1ConnectorClass

`func NewCcpmV1ConnectorClass(className string, type_ string, ) *CcpmV1ConnectorClass`

NewCcpmV1ConnectorClass instantiates a new CcpmV1ConnectorClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpmV1ConnectorClassWithDefaults

`func NewCcpmV1ConnectorClassWithDefaults() *CcpmV1ConnectorClass`

NewCcpmV1ConnectorClassWithDefaults instantiates a new CcpmV1ConnectorClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *CcpmV1ConnectorClass) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *CcpmV1ConnectorClass) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *CcpmV1ConnectorClass) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetType

`func (o *CcpmV1ConnectorClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CcpmV1ConnectorClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CcpmV1ConnectorClass) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


