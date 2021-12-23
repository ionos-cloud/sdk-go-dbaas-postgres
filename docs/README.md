# Introduction

## Overview

The IONOS SDK DBaaS Postgres provides you with access to the IONOS Database as a Service Postgres. The client library supports both simple and complex requests. It is designed for developers who are building applications in GO. The SDK for GO wraps the DBaaS Postgres API. All API operations are performed over SSL and authenticated using your IONOS Cloud portal credentials. The API can be accessed within an instance running in IONOS Cloud or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Getting Started

An IONOS account is required for access to the Cloud API; credentials from your registration are used to authenticate against the IONOS Cloud API.

### Installation

Install the Go language from the official [Go installation](https://golang.org/doc/install) guide.

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you will have to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```text
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following `go` command will download `sdk-go-dbaas-postgres` to your configured `GOPATH`:

```go
go get "github.com/ionos-cloud/sdk-go-dbaas-postgres"
```

The source code of the package will be located here:

```text
$GOPATH/src/github.com/ionos-cloud/sdk-go-dbaas-postgres
```

Create main package file _example.go_:

```go
package main

import (
    "fmt"
)

func main() {
}
```

Include the IONOS Cloud SDK for Go under the list of imports.

```go
import(
    "fmt"

    "github.com/ionos-cloud/sdk-go-dbaas-postgres"
)
```


### Authentication

The username and password or the authentication token can be manually specified when initializing the SDK client:

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfiguration(username, password, token))
```

Environment variables can also be used; the SDK uses the following variables:

* IONOS\_USERNAME - to specify the username used to login
* IONOS\_PASSWORD - to specify the password
* IONOS\_TOKEN - if an authentication token is being used

In this case, the client configuration must be initialized using `NewConfigurationFromEnv()`

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfigurationFromEnv())
```

{% hint style="danger" %}
**Warning**: Make sure to follow the Information Security Best Practices when using credentials within your code or storing them in a file.
{% endhint %}


## Feature Reference

The IONOS Cloud SDK for GO DBaaS Postgres aims to offer access to all resources in the IONOS Cloud API, and has additional features to make integration easier:

* Authentication for API calls
* Asynchronous request handling 

## FAQ

1. How can I open a bug report/feature request? 

Bug reports and feature requests can be opened in the Issues repository: [https://github.com/ionos-cloud/sdk-go-dbaas-postgres/issues/new/choose](https://github.com/ionos-cloud/sdk-go-dbaas-postgres/issues/new/choose)

2. Can I contribute to the GO DBaaS Postgres SDK?

Pure SDKs are automatically generated using OpenAPI Generator and don’t support manual changes. If you require changes, please open an issue and we will try to address it.
