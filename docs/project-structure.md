## Development specifications

### Vendoring tool 

- glide 
- govendor 
- 

### Project structure

```
./main.go
./controller.go - main controlling loop
./plan/
  plan.go - implements the logic for managing records
./kubernetes/
  manager.go - provides watching capabilities + clientset
./cloudprovider/ - cloud providers
  aws.go
  google.go
  fake.go 
./sources - list of sources
  fake.go
  ingress.go
  services.go
```

### Dependencies 

#### Logging 
  - logrus
  - uber-go/zap
  - glog

#### Build
  - Makefile
  - Bazel

### CI/CD

- Travis CI