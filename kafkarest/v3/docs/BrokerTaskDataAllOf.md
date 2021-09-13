# BrokerTaskDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**TaskType** | [**BrokerTaskType**](BrokerTaskType.md) |  | 
**TaskStatus** | [**BrokerTaskStatus**](BrokerTaskStatus.md) |  | 
**ShutdownScheduled** | Pointer to **bool** |  | [optional] 
**SubTaskStatuses** | **map[string]string** |  | 
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time at which this task was created. | [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time at which this task was last updated. | [readonly] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Broker** | [**Relationship**](Relationship.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


