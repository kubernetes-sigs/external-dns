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

- glide - widely used in kubernetes projects for vendoring

### Dependencies 

#### Logging 
  - logrus - simple logging library with supported formatter and log levels 

#### Flags
 - spf13/pflag - simple flag library which should suit our needs. We are unlikely to have subcommands and binary will be invoked in a manner to similar shown above

#### Clients
 - k8s.io/client-go
 - aws-sdk-go
 - google.golang.org/api/dns/v1

### Build
  - Makefile - build should be relatively simple process and Makefile should suit our needs

### CI/CD

 - Travis CI - https://github.com/kubernetes-incubator/external-dns/issues/9


