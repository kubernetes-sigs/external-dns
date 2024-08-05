# Civogo - The Golang client library for Civo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/civo/civogo?tab=doc)
[![Build Status](https://github.com/civo/civogo/workflows/Test/badge.svg)](https://github.com/civo/civogo/actions)
[![Lint](https://github.com/civo/civogo/workflows/Lint/badge.svg)](https://github.com/civo/civogo/actions)

Civogo is a Go client library for accessing the Civo cloud API.

You can view the client API docs at [https://pkg.go.dev/github.com/civo/civogo](https://pkg.go.dev/github.com/civo/civogo) and view the API documentation at [https://api.civo.com](https://api.civo.com)


## Install

```sh
go get github.com/civo/civogo
```

## Usage

```go
import "github.com/civo/civogo"
```

From there you create a Civo client specifying your API key and a region. Then you can use public methods to interact with Civo's API.

### Authentication

You will need both an API key and a region code to create a new client.

Your API key is listed within the [Civo control panel's security page](https://www.civo.com/account/security). You can also reset the token there, for example, if accidentally put it in source code and found it had been leaked.

For the region code, use any region you know exists, e.g. `LON1`. See the [API documentation](https://github.com/civo/civogo.git) for details.

```go
package main

import (
	"context"
	"github.com/civo/civogo"
)

const (
    apiKey = "mykeygoeshere"
    regionCode = "LON1"
)

func main() {
  client, err := civogo.NewClient(apiKey, regionCode)
  // ...
}
```

## Examples

To create a new Instance:

```go
config, err := client.NewInstanceConfig()
if err != nil {
  t.Errorf("Failed to create a new config: %s", err)
  return err
}

config.Hostname = "foo.example.com"

instance, err := client.CreateInstance(config)
if err != nil {
  t.Errorf("Failed to create instance: %s", err)
  return err
}
```

To get all Instances:

```go
instances, err := client.ListAllInstances()
if err != nil {
  t.Errorf("Failed to create instance: %s", err)
  return err
}

for _, i := range instances {
    fmt.Println(i.Hostname)
}
```

### Pagination

If a list of objects is paginated by the API, you must request pages individually. For example, to fetch all instances without using the `ListAllInstances` method:

```go
func MyListAllInstances(client *civogo.Client) ([]civogo.Instance, error) {
    list := []civogo.Instance{}

    pageOfItems, err := client.ListInstances(1, 50)
    if err != nil {
        return []civogo.Instance{}, err
    }

    if pageOfItems.Pages == 1 {
        return pageOfItems.Items, nil
    }

    for page := 2;  page<=pageOfItems.Pages; page++ {
        pageOfItems, err := client.ListInstances(1, 50)
        if err != nil {
            return []civogo.Instance{}, err
        }

        list = append(list, pageOfItems.Items)
    }

    return list, nil
}
```

## Error handler
​
In the latest version of the library we have added a new way to handle errors.
Below are some examples of how to use the new error handler, and the complete list of errors is [here](errors.go).
​
This is an example of how to make use of the new errors, suppose we want to create a new Kubernetes cluster, and do it this way but choose a name that already exists within the clusters that we have:
​
```go
// kubernetes config
configK8s := &civogo.KubernetesClusterConfig{
    NumTargetNodes: 5,
    Name: "existent-name",
}
// Send to create the cluster
resp, err := client.NewKubernetesClusters(configK8s)
if err != nil {
     if errors.Is(err, civogo.DatabaseKubernetesClusterDuplicateError) {
     // add some actions
     }
}
```
The following lines are new:
​
```go
if err != nil {
     if errors.Is(err, civogo.DatabaseKubernetesClusterDuplicateError) {
     // add some actions
     }
}
```
In this way. we can make decisions faster based on known errors, and we know what to expect. There is also the option of being able to say this to account for some errors but not others:
​
```go
if err != nil {
     if errors.Is(err, civogo.DatabaseKubernetesClusterDuplicateError) {
     // add some actions
     }
     if errors.Is(err, civogo.UnknownError) {
         // exit with error
     }
}
```
We can use `UnknownError` for errors that are not defined.

## Contributing

If you want to get involved, we'd love to receive a pull request - or an offer to help over our KUBE100 Slack channel. Please see the [contribution guidelines](CONTRIBUTING.md).
