# Sources and Providers

ExternalDNS supports swapping out endpoint **sources** and DNS **providers** and both sides are pluggable. There currently exist three sources and four provider implementations.

### Sources

Sources are an abstraction over any kind of source of desired Endpoints, e.g.:
* a list of Service objects from Kubernetes
* a random list for testing purposes
* an aggregated list of multiple nested sources

The `Source` interface has a single method called `Endpoints` that should return all desired Endpoint objects as a flat list.

```go
type Source interface {
	Endpoints() ([]*endpoint.Endpoint, error)
}
```

All sources live in package `source`.

* `ServiceSource`: collects all Services that have an external IP and returns them as Endpoint objects. The desired DNS name corresponds to an annotation set on the Service or is compiled from the Service attributes via the FQDN Go template string.
* `IngressSource`: collects all Ingresses that have an external IP and returns them as Endpoint objects. The desired DNS name corresponds to the host rules defined in the Ingress object.
* `IstioGatewaySource`: collects all Istio Gateways and returns them as Endpoint objects. The desired DNS name corresponds to the hosts listed within the servers spec of each Gateway object.
* `ContourIngressRouteSource`: collects all Contour IngressRoutes and returns them as Endpoint objects. The desired DNS name corresponds to the `virtualhost.fqdn` listed within the spec of each IngressRoute object.
* `FakeSource`: returns a random list of Endpoints for the purpose of testing providers without having access to a Kubernetes cluster.
* `ConnectorSource`: returns a list of Endpoint objects which are served by a tcp server configured through `connector-source-server` flag.
* `CRDSource`: returns a list of Endpoint objects sourced from the spec of CRD objects. For more details refer to [CRD source](../crd-source.md) documentation.
* `EmptySource`: returns an empty list of Endpoint objects for the purpose of testing and cleaning out entries.

### Providers

Providers are an abstraction over any kind of sink for desired Endpoints, e.g.:
* storing them in Google Cloud DNS
* printing them to stdout for testing purposes
* fanning out to multiple nested providers

The `Provider` interface has two methods: `Records` and `ApplyChanges`. `Records` should return all currently existing DNS records converted to Endpoint objects as a flat list. Upon receiving a change set (via an object of `plan.Changes`), `ApplyChanges` should translate these to the provider specific actions in order to persist them in the provider's storage.

```go
type Provider interface {
	Records() ([]*endpoint.Endpoint, error)
	ApplyChanges(changes *plan.Changes) error
}
```

The interface tries to be generic and assumes a flat list of records for both functions. However, many providers scope records into zones. Therefore, the provider implementation has to do some extra work to return that flat list. For instance, the AWS provider fetches the list of all hosted zones before it can return or apply the list of records. If the provider has no concept of zones or if it makes sense to cache the list of hosted zones it is happily allowed to do so. Furthermore, the provider should respect the `--domain-filter` flag to limit the affected records by a domain suffix. For instance, the AWS provider filters out all hosted zones that doesn't match that domain filter.

All providers live in package `provider`.

* `GoogleProvider`: returns and creates DNS records in Google Cloud DNS
* `AWSProvider`: returns and creates DNS records in AWS Route 53
* `AzureProvider`: returns and creates DNS records in Azure DNS
* `InMemoryProvider`: Keeps a list of records in local memory

### Usage

You can choose any combination of sources and providers on the command line. Given a cluster on AWS you would most likely want to use the Service and Ingress Source in combination with the AWS provider. `Service` + `InMemory` is useful for testing your service collecting functionality, whereas `Fake` + `Google` is useful for testing that the Google provider behaves correctly, etc.
