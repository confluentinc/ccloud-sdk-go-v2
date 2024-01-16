# \OrgPreferencesAiV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAiV1OrgPreference**](OrgPreferencesAiV1Api.md#GetAiV1OrgPreference) | **Get** /ai/v1/org-preferences | Read the organization&#39;s ai-assistant setting in org-preferences.
[**UpdateAiV1OrgPreference**](OrgPreferencesAiV1Api.md#UpdateAiV1OrgPreference) | **Patch** /ai/v1/org-preferences | Set the organization&#39;s ai-assistant setting in org-preferences.



## GetAiV1OrgPreference

> AiV1OrgPreferences GetAiV1OrgPreference(ctx).Execute()

Read the organization's ai-assistant setting in org-preferences.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgPreferencesAiV1Api.GetAiV1OrgPreference(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgPreferencesAiV1Api.GetAiV1OrgPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiV1OrgPreference`: AiV1OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgPreferencesAiV1Api.GetAiV1OrgPreference`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiV1OrgPreferenceRequest struct via the builder pattern


### Return type

[**AiV1OrgPreferences**](ai.v1.OrgPreferences.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiV1OrgPreference

> AiV1OrgPreferences UpdateAiV1OrgPreference(ctx).AiV1OrgPreferences(aiV1OrgPreferences).Execute()

Set the organization's ai-assistant setting in org-preferences.



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
    aiV1OrgPreferences := *openapiclient.NewAiV1OrgPreferences() // AiV1OrgPreferences |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgPreferencesAiV1Api.UpdateAiV1OrgPreference(context.Background()).AiV1OrgPreferences(aiV1OrgPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgPreferencesAiV1Api.UpdateAiV1OrgPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAiV1OrgPreference`: AiV1OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgPreferencesAiV1Api.UpdateAiV1OrgPreference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiV1OrgPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiV1OrgPreferences** | [**AiV1OrgPreferences**](AiV1OrgPreferences.md) |  | 

### Return type

[**AiV1OrgPreferences**](ai.v1.OrgPreferences.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

