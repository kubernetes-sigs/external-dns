# GoVultr

[![Automatic Releaser](https://github.com/vultr/govultr/actions/workflows/releaser.yml/badge.svg)](https://github.com/vultr/govultr/actions/workflows/releaser.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/vultr/govultr/v2)](https://pkg.go.dev/github.com/vultr/govultr/v2)
[![Unit/Coverage Tests](https://github.com/vultr/govultr/actions/workflows/coverage.yml/badge.svg)](https://github.com/vultr/govultr/actions/workflows/coverage.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/vultr/govultr)](https://goreportcard.com/report/github.com/vultr/govultr)

The official Vultr Go client - GoVultr allows you to interact with the Vultr V2 API.

GoVultr V1 that interacts with Vultr V1 API is now on the [v1 branch](https://github.com/vultr/govultr/tree/v1)

## Installation

```sh
go get -u github.com/vultr/govultr/v2
```

## Usage

Vultr uses a PAT (Personal Access token) to interact/authenticate with the APIs. Generate an API Key from the [API menu](https://my.vultr.com/settings/#settingsapi) in the Vultr Customer Portal.

To instantiate a GoVultr client, invoke `NewClient()`. You must pass your `PAT` to an `oauth2` library to create the `*http.Client`, which configures the `Authorization` header with your PAT as the `bearer api-key`.

The client has three optional parameters:

- BaseUrl: Change the Vultr default base URL
- UserAgent: Change the Vultr default UserAgent
- RateLimit: Set a delay between calls. Vultr limits the rate of back-to-back calls. Use this parameter to avoid rate-limit errors.

### Example Client Setup

```go
package main

import (
  "context"
  "os"

  "github.com/vultr/govultr/v2"
  "golang.org/x/oauth2"
)

func main() {
  apiKey := os.Getenv("VultrAPIKey")

  config := &oauth2.Config{}
  ctx := context.Background()
  ts := config.TokenSource(ctx, &oauth2.Token{AccessToken: apiKey})
  vultrClient := govultr.NewClient(oauth2.NewClient(ctx, ts))

  // Optional changes
  _ = vultrClient.SetBaseURL("https://api.vultr.com")
  vultrClient.SetUserAgent("mycool-app")
  vultrClient.SetRateLimit(500)
}
```

### Example Usage

Create a VPS

```go
instanceOptions := &govultr.InstanceCreateReq{
  Label:                "awesome-go-app",
  Hostname:             "awesome-go.com",
  Backups:              "enabled",
  EnableIPv6:           BoolToBoolPtr(false),
  OsID:                 362,
  Plan:                 "vc2-1c-2gb",   
  Region:               "ewr",
}

res, err := vultrClient.Instance.Create(context.Background(), instanceOptions)

if err != nil {
  fmt.Println(err)
}
```

## Pagination

GoVultr v2 introduces pagination for all list calls. Each list call returns a `meta` struct containing the total amount of items in the list and next/previous links to navigate the paging.

```go
// Meta represents the available pagination information
type Meta struct {
  Total int `json:"total"`
  Links *Links
}

// Links represent the next/previous cursor in your pagination calls
type Links struct {
  Next string `json:"next"`
  Prev string `json:"prev"`
}

```
Pass a `per_page` value to the `list_options` struct to adjust the number of items returned per call. The default is 100 items per page and max is 500 items per page. 

This example demonstrates how to retrieve all of your instances, with one instance per page.

```go
listOptions := &govultr.ListOptions{PerPage: 1}
for {
    i, meta, err := client.Instance.List(ctx, listOptions)
    if err != nil {
        return nil, err
    }
    for _, v := range i {
        fmt.Println(v)
    }

    if meta.Links.Next == "" {
        break
    } else {
        listOptions.Cursor = meta.Links.Next
        continue
    }
}    
```
## Versioning

This project follows [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/vultr/govultr/tags).

## Documentation

See our documentation for [detailed information about API v2](https://www.vultr.com/api/).

See our [GoDoc](https://pkg.go.dev/github.com/vultr/govultr/v2) documentation for more details about this client's functionality.

## Contributing

Feel free to send pull requests our way! Please see the [contributing guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details.
