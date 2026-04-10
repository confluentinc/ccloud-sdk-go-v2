# \StreamsGroupV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaStreamsGroup**](StreamsGroupV3Api.md#GetKafkaStreamsGroup) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id} | Get Streams Group
[**GetKafkaStreamsGroupMember**](StreamsGroupV3Api.md#GetKafkaStreamsGroupMember) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id} | Get Streams Group Member
[**GetKafkaStreamsGroupMemberAssignmentTaskPartitions**](StreamsGroupV3Api.md#GetKafkaStreamsGroupMemberAssignmentTaskPartitions) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/assignments/{assignments_type}/subtopologies/{subtopology_id} | List Streams Group Assignments Task Partitions of a Specific Type and Subtopology
[**GetKafkaStreamsGroupMemberAssignments**](StreamsGroupV3Api.md#GetKafkaStreamsGroupMemberAssignments) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/assignments | Get Streams Group Member Assignments
[**GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions**](StreamsGroupV3Api.md#GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/target-assignments/{assignments_type}/subtopologies/{subtopology_id} | List Streams Group Target Assignments Task Partitions of a Specific Type and Subtopology
[**GetKafkaStreamsGroupMemberTargetAssignments**](StreamsGroupV3Api.md#GetKafkaStreamsGroupMemberTargetAssignments) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/target-assignments | Get Streams Group Member Target Assignments
[**GetKafkaStreamsGroupSubtopology**](StreamsGroupV3Api.md#GetKafkaStreamsGroupSubtopology) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/subtopologies/{subtopology_id} | Get Streams Group Subtopology
[**ListKafkaStreamsGroupMemberAssignmentTasks**](StreamsGroupV3Api.md#ListKafkaStreamsGroupMemberAssignmentTasks) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/assignments/{assignments_type} | List Streams Group Assignments of a Specific Type
[**ListKafkaStreamsGroupMemberTargetAssignmentTasks**](StreamsGroupV3Api.md#ListKafkaStreamsGroupMemberTargetAssignmentTasks) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members/{member_id}/target-assignments/{assignments_type} | List Streams Group Target Assignments of a Specific Type
[**ListKafkaStreamsGroupMembers**](StreamsGroupV3Api.md#ListKafkaStreamsGroupMembers) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/members | List Streams Group Members
[**ListKafkaStreamsGroupSubtopologies**](StreamsGroupV3Api.md#ListKafkaStreamsGroupSubtopologies) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups/{group_id}/subtopologies | List Streams Group Subtopologies
[**ListKafkaStreamsGroups**](StreamsGroupV3Api.md#ListKafkaStreamsGroups) | **Get** /kafka/v3/clusters/{cluster_id}/streams-groups | List Streams Groups



## GetKafkaStreamsGroup

> StreamsGroupData GetKafkaStreamsGroup(ctx, clusterId, groupId).Execute()

Get Streams Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroup(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroup`: StreamsGroupData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamsGroupData**](StreamsGroupData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupMember

> StreamsGroupMemberData GetKafkaStreamsGroupMember(ctx, clusterId, groupId, memberId).Execute()

Get Streams Group Member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupMember(context.Background(), clusterId, groupId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupMember`: StreamsGroupMemberData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsGroupMemberData**](StreamsGroupMemberData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupMemberAssignmentTaskPartitions

> StreamsTaskData GetKafkaStreamsGroupMemberAssignmentTaskPartitions(ctx, clusterId, groupId, memberId, assignmentsType, subtopologyId).Execute()

List Streams Group Assignments Task Partitions of a Specific Type and Subtopology



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.
    assignmentsType := "active" // string | The streams member Assignment type.
    subtopologyId := "subtopology-1" // string | The streams subtopology ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignmentTaskPartitions(context.Background(), clusterId, groupId, memberId, assignmentsType, subtopologyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignmentTaskPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupMemberAssignmentTaskPartitions`: StreamsTaskData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignmentTaskPartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 
**assignmentsType** | **string** | The streams member Assignment type. | 
**subtopologyId** | **string** | The streams subtopology ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupMemberAssignmentTaskPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**StreamsTaskData**](StreamsTaskData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupMemberAssignments

> StreamsGroupMemberAssignmentData GetKafkaStreamsGroupMemberAssignments(ctx, clusterId, groupId, memberId).Execute()

Get Streams Group Member Assignments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignments(context.Background(), clusterId, groupId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupMemberAssignments`: StreamsGroupMemberAssignmentData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupMemberAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupMemberAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsGroupMemberAssignmentData**](StreamsGroupMemberAssignmentData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions

> StreamsTaskData GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions(ctx, clusterId, groupId, memberId, assignmentsType, subtopologyId).Execute()

List Streams Group Target Assignments Task Partitions of a Specific Type and Subtopology



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.
    assignmentsType := "active" // string | The streams member Assignment type.
    subtopologyId := "subtopology-1" // string | The streams subtopology ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions(context.Background(), clusterId, groupId, memberId, assignmentsType, subtopologyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions`: StreamsTaskData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignmentTaskPartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 
**assignmentsType** | **string** | The streams member Assignment type. | 
**subtopologyId** | **string** | The streams subtopology ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupMemberTargetAssignmentTaskPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**StreamsTaskData**](StreamsTaskData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupMemberTargetAssignments

> StreamsGroupMemberAssignmentData GetKafkaStreamsGroupMemberTargetAssignments(ctx, clusterId, groupId, memberId).Execute()

Get Streams Group Member Target Assignments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignments(context.Background(), clusterId, groupId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupMemberTargetAssignments`: StreamsGroupMemberAssignmentData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupMemberTargetAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupMemberTargetAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsGroupMemberAssignmentData**](StreamsGroupMemberAssignmentData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaStreamsGroupSubtopology

> StreamsGroupSubtopologyData GetKafkaStreamsGroupSubtopology(ctx, clusterId, groupId, subtopologyId).Execute()

Get Streams Group Subtopology



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    subtopologyId := "subtopology-1" // string | The streams subtopology ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.GetKafkaStreamsGroupSubtopology(context.Background(), clusterId, groupId, subtopologyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.GetKafkaStreamsGroupSubtopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaStreamsGroupSubtopology`: StreamsGroupSubtopologyData
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.GetKafkaStreamsGroupSubtopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**subtopologyId** | **string** | The streams subtopology ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaStreamsGroupSubtopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StreamsGroupSubtopologyData**](StreamsGroupSubtopologyData.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaStreamsGroupMemberAssignmentTasks

> StreamsTaskDataList ListKafkaStreamsGroupMemberAssignmentTasks(ctx, clusterId, groupId, memberId, assignmentsType).Execute()

List Streams Group Assignments of a Specific Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.
    assignmentsType := "active" // string | The streams member Assignment type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.ListKafkaStreamsGroupMemberAssignmentTasks(context.Background(), clusterId, groupId, memberId, assignmentsType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.ListKafkaStreamsGroupMemberAssignmentTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaStreamsGroupMemberAssignmentTasks`: StreamsTaskDataList
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.ListKafkaStreamsGroupMemberAssignmentTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 
**assignmentsType** | **string** | The streams member Assignment type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaStreamsGroupMemberAssignmentTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StreamsTaskDataList**](StreamsTaskDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaStreamsGroupMemberTargetAssignmentTasks

> StreamsTaskDataList ListKafkaStreamsGroupMemberTargetAssignmentTasks(ctx, clusterId, groupId, memberId, assignmentsType).Execute()

List Streams Group Target Assignments of a Specific Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.
    memberId := "member-1" // string | The streams member ID.
    assignmentsType := "active" // string | The streams member Assignment type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.ListKafkaStreamsGroupMemberTargetAssignmentTasks(context.Background(), clusterId, groupId, memberId, assignmentsType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.ListKafkaStreamsGroupMemberTargetAssignmentTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaStreamsGroupMemberTargetAssignmentTasks`: StreamsTaskDataList
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.ListKafkaStreamsGroupMemberTargetAssignmentTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 
**memberId** | **string** | The streams member ID. | 
**assignmentsType** | **string** | The streams member Assignment type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaStreamsGroupMemberTargetAssignmentTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StreamsTaskDataList**](StreamsTaskDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaStreamsGroupMembers

> StreamsGroupMemberDataList ListKafkaStreamsGroupMembers(ctx, clusterId, groupId).Execute()

List Streams Group Members



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.ListKafkaStreamsGroupMembers(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.ListKafkaStreamsGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaStreamsGroupMembers`: StreamsGroupMemberDataList
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.ListKafkaStreamsGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaStreamsGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamsGroupMemberDataList**](StreamsGroupMemberDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaStreamsGroupSubtopologies

> StreamsGroupSubtopologyDataList ListKafkaStreamsGroupSubtopologies(ctx, clusterId, groupId).Execute()

List Streams Group Subtopologies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    groupId := "group-1" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.ListKafkaStreamsGroupSubtopologies(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.ListKafkaStreamsGroupSubtopologies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaStreamsGroupSubtopologies`: StreamsGroupSubtopologyDataList
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.ListKafkaStreamsGroupSubtopologies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaStreamsGroupSubtopologiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamsGroupSubtopologyDataList**](StreamsGroupSubtopologyDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaStreamsGroups

> StreamsGroupDataList ListKafkaStreamsGroups(ctx, clusterId).Execute()

List Streams Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "cluster-1" // string | The Kafka cluster ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamsGroupV3Api.ListKafkaStreamsGroups(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsGroupV3Api.ListKafkaStreamsGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkaStreamsGroups`: StreamsGroupDataList
    fmt.Fprintf(os.Stdout, "Response from `StreamsGroupV3Api.ListKafkaStreamsGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKafkaStreamsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamsGroupDataList**](StreamsGroupDataList.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

