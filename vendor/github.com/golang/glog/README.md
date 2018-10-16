# glog-logrus

This packages is a replacement for [glog](github.com/golang/glog)
in projects that use the [logrus](https://godoc.org/github.com/sirupsen/logrus).

It is inspired by istio's glog package for zap:
https://github.com/istio/glog

## Usage

Override the official glog package with this one.
This simply replaces the code in `vendor/golang/glog` with the code of this package.

In your `Gopkg.toml`:
```toml
[[override]]
  name = "github.com/golang/glog"
  source = "github.com/kubermatic/glog-logrus"
```

In your `main.go`:
```go
// Import the package like it is original glog
import (
  "github.com/golang/glog"
  "github.com/sirupsen/logrus"
)

// Create logrus logger in your main.go
logger := logrus.New()
logger.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}

// Overriding the default glog with our logrus glog implementation.
// Thus we need to pass it our logrus logger object.
glog.SetLogger(logger.WithField("foo", "bar"))
```

Setting the logger to the glog package **MUST** happen before using glog in any package.

The functionality of logging the filename and line number is not preserved at this time.

## Function Levels

|     glog     | logrus |
| ------------ | ------ |
| Info         | Debug  |
| InfoDepth    | Debug  |
| Infof        | Debug  |
| Infoln       | Debug  |
| Warning      | Warn   |
| WarningDepth | Warn   |
| Warningf     | Warn   |
| Warningln    | Warn   |
| Error        | Error  |
| ErrorDepth   | Error  |
| Errorf       | Error  |
| Errorln      | Error  |
| Exit         | Fatal  |
| ExitDepth    | Fatal  |
| Exitf        | Fatal  |
| Exitln       | Fatal  |
| Fatal        | Fatal  |
| FatalDepth   | Fatal  |
| Fatalf       | Fatal  |
| Fatalln      | Fatal  |

This table is rather opinionated and build for use with the Kubernetes' [Go client](https://github.com/kubernetes/client-go).
