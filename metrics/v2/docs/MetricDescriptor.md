# MetricDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the metric. | 
**Labels** | [**[]LabelDescriptor**](LabelDescriptor.md) | Labels for filtering and aggregating this metric. For example, you can filter the &#x60;io.confluent.kafka.server/received_bytes&#x60; metric by the &#x60;topic&#x60; label.  | 
**Name** | **string** | The unique name of the metric, for example, &#x60;io.confluent.kafka.server/received_bytes&#x60;.  | 
**Unit** | **string** | The unit that the metric value is reported in. Units follow the format described in [the Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html). For example, &#x60;By&#x60; for bytes and &#x60;By/s&#x60; for bytes per second.  | 
**Type** | **string** | The type of the measurement. The metric type follows the [OpenTelemetry](https://opentelemetry.io/) exposition format. * &#x60;GAUGE_(INT64|DOUBLE)&#x60;: An instantaneous measurement of a value.   Gauge metrics are implicitly averaged when aggregating over time. * &#x60;COUNTER_(INT64|DOUBLE)&#x60;: The count of occurrences in a _single (one minute) sampling   interval_ (unless otherwise stated in the metric description).   Counter metrics are implicitly summed when aggregating over time.  | 
**LifecycleStage** | **string** | The support lifecycle stage for this metric: * &#x60;PREVIEW&#x60;: May change at any time * &#x60;GENERAL_AVAILABILITY&#x60;: Fully released and stable. Will not change without notice. * &#x60;DEPRECATED&#x60;: No longer supported. Will be removed in the future at the annouced date.  | 
**Exportable** | **bool** | Is this metric included in the &#x60;/export&#x60; endpoint response? | 
**Resources** | **[]string** | The resource types to which this metric pertains. | 

## Methods

### NewMetricDescriptor

`func NewMetricDescriptor(description string, labels []LabelDescriptor, name string, unit string, type_ string, lifecycleStage string, exportable bool, resources []string, ) *MetricDescriptor`

NewMetricDescriptor instantiates a new MetricDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDescriptorWithDefaults

`func NewMetricDescriptorWithDefaults() *MetricDescriptor`

NewMetricDescriptorWithDefaults instantiates a new MetricDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabels

`func (o *MetricDescriptor) GetLabels() []LabelDescriptor`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MetricDescriptor) GetLabelsOk() (*[]LabelDescriptor, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MetricDescriptor) SetLabels(v []LabelDescriptor)`

SetLabels sets Labels field to given value.


### GetName

`func (o *MetricDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetUnit

`func (o *MetricDescriptor) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricDescriptor) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricDescriptor) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetType

`func (o *MetricDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDescriptor) SetType(v string)`

SetType sets Type field to given value.


### GetLifecycleStage

`func (o *MetricDescriptor) GetLifecycleStage() string`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *MetricDescriptor) GetLifecycleStageOk() (*string, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *MetricDescriptor) SetLifecycleStage(v string)`

SetLifecycleStage sets LifecycleStage field to given value.


### GetExportable

`func (o *MetricDescriptor) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *MetricDescriptor) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *MetricDescriptor) SetExportable(v bool)`

SetExportable sets Exportable field to given value.


### GetResources

`func (o *MetricDescriptor) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MetricDescriptor) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MetricDescriptor) SetResources(v []string)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


