# exo cli

Manage easily your Exoscale infrastructure from the exo command-line


## Installation

We provide many alternatives on the [releases](https://github.com/exoscale/egoscale/releases) page.

### Manual compilation

```
$ go get -u github.com/golang/dep/cmd/dep
$ go get -d github.com/exoscale/egoscale/...

$ cd $GOPATH/src/github.com/exoscale/egoscale/
$ dep ensure -vendor-only

$ cd cmd/exo
$ dep ensure -vendor-only

$ go install
```

## Configuration

The CLI will guide you in the initial configuration.
The configuration file and all assets created by any `exo` command will be saved in the `~/.exoscale/` folder.

You can find your credentials in our [Exoscale Console](https://portal.exoscale.com/account/profile/api)

```shell
$ exo config
```

## Usage

```shell
$ exo --help
```
