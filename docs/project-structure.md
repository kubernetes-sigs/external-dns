## Development specifications

Proposal regarding the project structure and related tools

### How to run

```
external-dns --in-cluster=false --dnsprovider=aws --source=ingress --source=service
```

### Project structure

```
./main.go
./config - store configurations, flag parsing
    config.go 
./controller - main controlling loop
    controller.go 
./plan/
    plan.go - implements the logic for managing records
./kubernetes/
    manager.go - provides watching capabilities + clientset
./dnsprovider/ - dns providers
    aws.go
    google.go
    fake.go 
    dnsprovider.go - interface
./source/ - list of sources
    fake.go
    ingress.go
    service.go
    source.go - interface
```

### Vendoring tool 

- glide - **to be used**
- alternatives: govendor, godep

### Dependencies 

#### Logging 
  - logrus - **to be used**
  - alternatives: uber-go/zap, glog

#### Flags

 - spf13/pflag - **to be used**
 - alternatives - kingpin, jessevdk/go-flags

#### Clients
 - k8s.io/client-go
 - aws-sdk-go
 - google.golang.org/api/dns/v1

### Build
  - Makefile - **to be used**
  - alternatives: bazel.io 

### CI/CD

 - Travis CI - https://github.com/kubernetes-incubator/external-dns/issues/9


