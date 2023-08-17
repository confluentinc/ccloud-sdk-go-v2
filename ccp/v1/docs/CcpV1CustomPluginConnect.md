# CcpV1CustomPluginConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginType** | **string** | [immutable] Plugin Type.  | 
**ConnectorClass** | **string** | [immutable] Java class or alias for connector. You can get connector class from connector documentation provided by developer. | 
**ConnectorType** | **string** | [immutable] Custom Connector type.  | 
**SensitiveConfigProperties** | Pointer to **[]string** | A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector. | [optional] 

## Methods

### NewCcpV1CustomPluginConnect

`func NewCcpV1CustomPluginConnect(pluginType string, connectorClass string, connectorType string, ) *CcpV1CustomPluginConnect`

NewCcpV1CustomPluginConnect instantiates a new CcpV1CustomPluginConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcpV1CustomPluginConnectWithDefaults

`func NewCcpV1CustomPluginConnectWithDefaults() *CcpV1CustomPluginConnect`

NewCcpV1CustomPluginConnectWithDefaults instantiates a new CcpV1CustomPluginConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginType

`func (o *CcpV1CustomPluginConnect) GetPluginType() string`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *CcpV1CustomPluginConnect) GetPluginTypeOk() (*string, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *CcpV1CustomPluginConnect) SetPluginType(v string)`

SetPluginType sets PluginType field to given value.


### GetConnectorClass

`func (o *CcpV1CustomPluginConnect) GetConnectorClass() string`

GetConnectorClass returns the ConnectorClass field if non-nil, zero value otherwise.

### GetConnectorClassOk

`func (o *CcpV1CustomPluginConnect) GetConnectorClassOk() (*string, bool)`

GetConnectorClassOk returns a tuple with the ConnectorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorClass

`func (o *CcpV1CustomPluginConnect) SetConnectorClass(v string)`

SetConnectorClass sets ConnectorClass field to given value.


### GetConnectorType

`func (o *CcpV1CustomPluginConnect) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CcpV1CustomPluginConnect) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CcpV1CustomPluginConnect) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.


### GetSensitiveConfigProperties

`func (o *CcpV1CustomPluginConnect) GetSensitiveConfigProperties() []string`

GetSensitiveConfigProperties returns the SensitiveConfigProperties field if non-nil, zero value otherwise.

### GetSensitiveConfigPropertiesOk

`func (o *CcpV1CustomPluginConnect) GetSensitiveConfigPropertiesOk() (*[]string, bool)`

GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveConfigProperties

`func (o *CcpV1CustomPluginConnect) SetSensitiveConfigProperties(v []string)`

SetSensitiveConfigProperties sets SensitiveConfigProperties field to given value.

### HasSensitiveConfigProperties

`func (o *CcpV1CustomPluginConnect) HasSensitiveConfigProperties() bool`

HasSensitiveConfigProperties returns a boolean if a field has been set.


### AsCcpV1CustomPluginConfigOneOf

`func (s *CcpV1CustomPluginConnect) AsCcpV1CustomPluginConfigOneOf() CcpV1CustomPluginConfigOneOf`

Convenience method to wrap this instance of CcpV1CustomPluginConnect in CcpV1CustomPluginConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


