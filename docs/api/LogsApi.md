# \LogsApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterLogsGet**](LogsApi.md#ClusterLogsGet) | **Get** /clusters/{clusterId}/logs | Get logs of your cluster|



## ClusterLogsGet

```go
var result ClusterLogs = ClusterLogsGet(ctx, clusterId)
                      .Limit(limit)
                      .Start(start)
                      .End(end)
                      .Execute()
```

Get logs of your cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-postgres"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    limit := int32(56) // int32 | The maximal number of log lines to return. (optional)
    start := time.Now() // time.Time | The start time for the query in RFC3339 format. (optional)
    end := time.Now() // time.Time | The end time for the query in RFC3339 format. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.LogsApi.ClusterLogsGet(context.Background(), clusterId).Limit(limit).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ClusterLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClusterLogsGet`: ClusterLogs
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.ClusterLogsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterLogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximal number of log lines to return. | |
| **start** | **time.Time** | The start time for the query in RFC3339 format. | |
| **end** | **time.Time** | The end time for the query in RFC3339 format. | |

### Return type

[**ClusterLogs**](ClusterLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


