# \MetadataApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**InfosVersionGet**](MetadataApi.md#InfosVersionGet) | **Get** /infos/version | Get the current API version|
|[**InfosVersionsGet**](MetadataApi.md#InfosVersionsGet) | **Get** /infos/versions | Fetch all API versions|
|[**VersionsGet**](MetadataApi.md#VersionsGet) | **Get** /versions | PostgreSQL versions for new clusters|



## InfosVersionGet

```go
var result APIVersion = InfosVersionGet(ctx)
                      .Execute()
```

Get the current API version



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.MetadataApi.InfosVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InfosVersionGet`: APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiInfosVersionGetRequest struct via the builder pattern


### Return type

[**APIVersion**](../models/APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"MetadataApiService.InfosVersionGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "MetadataApiService.InfosVersionGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "MetadataApiService.InfosVersionGet": {
    "port": "8443",
},
})
```


## InfosVersionsGet

```go
var result []APIVersion = InfosVersionsGet(ctx)
                      .Execute()
```

Fetch all API versions



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.MetadataApi.InfosVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InfosVersionsGet`: []APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiInfosVersionsGetRequest struct via the builder pattern


### Return type

[**[]APIVersion**](../models/APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"MetadataApiService.InfosVersionsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "MetadataApiService.InfosVersionsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "MetadataApiService.InfosVersionsGet": {
    "port": "8443",
},
})
```


## VersionsGet

```go
var result PostgresVersionList = VersionsGet(ctx)
                      .Execute()
```

PostgreSQL versions for new clusters



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.MetadataApi.VersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.VersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VersionsGet`: PostgresVersionList
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.VersionsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiVersionsGetRequest struct via the builder pattern


### Return type

[**PostgresVersionList**](../models/PostgresVersionList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"MetadataApiService.VersionsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "MetadataApiService.VersionsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "MetadataApiService.VersionsGet": {
    "port": "8443",
},
})
```

