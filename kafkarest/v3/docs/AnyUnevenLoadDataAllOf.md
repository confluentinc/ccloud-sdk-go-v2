# AnyUnevenLoadDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | [**AnyUnevenLoadStatus**](AnyUnevenLoadStatus.md) |  | 
**PreviousStatus** | [**AnyUnevenLoadStatus**](AnyUnevenLoadStatus.md) |  | 
**StatusUpdatedAt** | [**time.Time**](time.Time.md) | The date and time at which this task was created. | [readonly] 
**PreviousStatusUpdatedAt** | [**time.Time**](time.Time.md) | The date and time at which this task was created. | [readonly] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**BrokerTasks** | [**Relationship**](Relationship.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


