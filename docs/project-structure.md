## Development specifications

Proposal regarding the project structure and related tools

### How to run

```
external-dns --outside-cluster --dnsprovider=aws --source=ingress --source=service
```

### Vendoring tool 

- glide 
- alternatives: govendor, godep

### Project structure

```
./main.go
./config.go - store configurations, flag parsing
./controller - main controlling loop
    controller.go 
./plan/
    record.go - dns provider neutral struct for records
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
    services.go
    source.go - interface
```

### Dependencies 

#### Logging 
  - logrus
  - alternatives: uber-go/zap, glog

#### Build
  - Makefile
  - alternatives: bazel.io 

### CI/CD

 - Travis CI - https://github.com/kubernetes-incubator/external-dns/issues/9

### Flags

 - spf13/pflag
 - alternatives - kingpin, jessevdk/go-flags
 Depends what kind of cmd line requirements we have
