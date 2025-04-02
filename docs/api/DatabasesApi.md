# \DatabasesApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatabasesDelete**](DatabasesApi.md#DatabasesDelete) | **Delete** /clusters/{clusterId}/databases/{databasename} | Delete database|
|[**DatabasesGet**](DatabasesApi.md#DatabasesGet) | **Get** /clusters/{clusterId}/databases/{databasename} | Get database|
|[**DatabasesList**](DatabasesApi.md#DatabasesList) | **Get** /clusters/{clusterId}/databases | List databases|
|[**DatabasesPost**](DatabasesApi.md#DatabasesPost) | **Post** /clusters/{clusterId}/databases | Create a database|



## DatabasesDelete

```go
var result  = DatabasesDelete(ctx, clusterId, databasename)
                      .Execute()
```

Delete database



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
    databasename := "benjamindb" // string | The database name.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.DatabasesApi.DatabasesDelete(context.Background(), clusterId, databasename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DatabasesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**databasename** | **string** | The database name. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatabasesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DatabasesApiService.DatabasesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DatabasesApiService.DatabasesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DatabasesApiService.DatabasesDelete": {
    "port": "8443",
},
})
```


## DatabasesGet

```go
var result DatabaseResource = DatabasesGet(ctx, clusterId, databasename)
                      .Execute()
```

Get database



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
    databasename := "benjamindb" // string | The database name.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DatabasesApi.DatabasesGet(context.Background(), clusterId, databasename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DatabasesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatabasesGet`: DatabaseResource
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DatabasesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**databasename** | **string** | The database name. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatabasesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**DatabaseResource**](../models/DatabaseResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DatabasesApiService.DatabasesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DatabasesApiService.DatabasesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DatabasesApiService.DatabasesGet": {
    "port": "8443",
},
})
```


## DatabasesList

```go
var result DatabaseList = DatabasesList(ctx, clusterId)
                      .Limit(limit)
                      .Offset(offset)
                      .Execute()
```

List databases



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
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DatabasesApi.DatabasesList(context.Background(), clusterId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DatabasesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatabasesList`: DatabaseList
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DatabasesList`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatabasesListRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|

### Return type

[**DatabaseList**](../models/DatabaseList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DatabasesApiService.DatabasesList"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DatabasesApiService.DatabasesList": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DatabasesApiService.DatabasesList": {
    "port": "8443",
},
})
```


## DatabasesPost

```go
var result DatabaseResource = DatabasesPost(ctx, clusterId)
                      .Database(database)
                      .Execute()
```

Create a database



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
    database := *openapiclient.NewDatabase(*openapiclient.NewDatabaseProperties("benjamindb", "benjamin")) // Database | a database to create

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.DatabasesApi.DatabasesPost(context.Background(), clusterId).Database(database).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DatabasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatabasesPost`: DatabaseResource
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DatabasesPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatabasesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **database** | [**Database**](../models/Database.md) | a database to create | |

### Return type

[**DatabaseResource**](../models/DatabaseResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"DatabasesApiService.DatabasesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "DatabasesApiService.DatabasesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "DatabasesApiService.DatabasesPost": {
    "port": "8443",
},
})
```

