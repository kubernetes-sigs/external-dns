## Development specifications

Proposal regarding the project structure and related tools

### Vendoring tool 

- glide 
- alternatives: govendor, godep

### Project structure

```
./main.go
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
./source/ - list of sources
    fake.go
    ingress.go
    services.go
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

