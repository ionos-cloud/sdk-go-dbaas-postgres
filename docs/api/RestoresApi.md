# \RestoresApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterRestorePost**](RestoresApi.md#ClusterRestorePost) | **Post** /clusters/{clusterId}/restore | In-place restore of a cluster|



## ClusterRestorePost

```go
var result  = ClusterRestorePost(ctx, clusterId)
                      .CreateRestoreRequest(createRestoreRequest)
                      .Execute()
```

In-place restore of a cluster



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
    createRestoreRequest := *openapiclient.NewCreateRestoreRequest("dcd31531-3ac8-11eb-9feb-046c59cc737e") // CreateRestoreRequest | The restore request to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RestoresApi.ClusterRestorePost(context.Background(), clusterId).CreateRestoreRequest(createRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ClusterRestorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterRestorePostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createRestoreRequest** | [**CreateRestoreRequest**](../models/CreateRestoreRequest.md) | The restore request to create. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RestoresApiService.ClusterRestorePost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RestoresApiService.ClusterRestorePost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RestoresApiService.ClusterRestorePost": {
    "port": "8443",
},
})
```

