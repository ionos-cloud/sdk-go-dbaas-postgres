# \LogsApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterLogsGet**](LogsApi.md#ClusterLogsGet) | **Get** /clusters/{clusterId}/logs | Get logs of your cluster|



## ClusterLogsGet

```go
var result ClusterLogs = ClusterLogsGet(ctx, clusterId)
                      .Start(start)
                      .End(end)
                      .Direction(direction)
                      .Limit(limit)
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
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.
    start := time.Now() // time.Time | The start time for the query in RFC3339 format. Must not be more than 30 days ago but before the end parameter. The default is 30 days ago. (optional)
    end := time.Now() // time.Time | The end time for the query in RFC3339 format. Must not be greater than now. The default is the current timestamp. (optional)
    direction := "direction_example" // string | The direction in which to scan through the logs. The logs are returned in order of the direction. (optional) (default to "BACKWARD")
    limit := int32(56) // int32 | The maximal number of log lines to return.  If the limit is reached then log lines will be cut at the end (respecting the scan direction). (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.LogsApi.ClusterLogsGet(context.Background(), clusterId).Start(start).End(end).Direction(direction).Limit(limit).Execute()
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
| **start** | **time.Time** | The start time for the query in RFC3339 format. Must not be more than 30 days ago but before the end parameter. The default is 30 days ago. | |
| **end** | **time.Time** | The end time for the query in RFC3339 format. Must not be greater than now. The default is the current timestamp. | |
| **direction** | **string** | The direction in which to scan through the logs. The logs are returned in order of the direction. | [default to &quot;BACKWARD&quot;]|
| **limit** | **int32** | The maximal number of log lines to return.  If the limit is reached then log lines will be cut at the end (respecting the scan direction). | [default to 100]|

### Return type

[**ClusterLogs**](../models/ClusterLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"LogsApiService.ClusterLogsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "LogsApiService.ClusterLogsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "LogsApiService.ClusterLogsGet": {
    "port": "8443",
},
})
```

