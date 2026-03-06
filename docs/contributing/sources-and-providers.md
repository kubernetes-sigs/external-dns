---
tags:
  - sources
  - providers
  - contributing
  - informers
---

# Sources and Providers

ExternalDNS supports swapping out endpoint **sources** and DNS **providers** and both sides are pluggable. There currently exist multiple sources for different provider implementations.

## Sources

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
* `CRDSource`: returns a list of Endpoint objects sourced from the spec of CRD objects. For more details refer to [CRD source](../sources/crd.md) documentation.
* `EmptySource`: returns an empty list of Endpoint objects for the purpose of testing and cleaning out entries.

### Adding New Sources

When creating a new source, add the following annotations above the source struct definition:

```go
// myNewSource is an implementation of Source for MyResource objects.
//
// +externaldns:source:name=my-new-source
// +externaldns:source:category=Kubernetes Core
// +externaldns:source:description=Creates DNS entries from MyResource objects
// +externaldns:source:resources=MyResource<Kind.apigroup.subdomain.domain>
// +externaldns:source:filters=
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=false|true
type myNewSource struct {
    // ... fields
}
```

**Annotation Reference:**

* `+externaldns:source:name` - The CLI name used with `--source` flag (required)
* `+externaldns:source:category` - Category for documentation grouping (required)
* `+externaldns:source:description` - Short description of what the source does (required)
* `+externaldns:source:resources` - Kubernetes resources watched (comma-separated). Convention `Kind.apigroup.subdomain.domain`
* `+externaldns:source:filters` - Supported filter types (annotation, label)
* `+externaldns:source:namespace` - Namespace support: comma-separated values (all, single, multiple)
* `+externaldns:source:fqdn-template` - FQDN template support (true, false)
* `+externaldns:source:events` - Kubernetes [`events`](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_events/) support  (true, false)

After adding annotations, run `make generate-sources-documentation` to update sources file.

### Informer Patterns

When adding a new source that watches Kubernetes resources via informers, apply
the following patterns from `source/informers` to keep memory usage low and
objects self-describing in the cache.

#### Transformers

Call `SetTransform` on every informer **before** `factory.Start()`. Use
`TransformerWithOptions[T]` with the options appropriate for the resource:

```go
if err = myInformer.Informer().SetTransform(informers.TransformerWithOptions[*v1.MyResource](
    informers.TransformRemoveManagedFields(),     // always — can be megabytes per object
    informers.TransformRemoveLastAppliedConfig(), // always — full JSON snapshot stored as annotation
    informers.TransformRemoveStatusConditions(),  // when Status.Conditions is not used by the source
    informers.TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"), // when only specific annotations are needed
)); err != nil {
    return nil, err
}
```

`TransformRemoveManagedFields` and `TransformRemoveLastAppliedConfig` should be
applied to every informer. `TransformRemoveStatusConditions` can be omitted only
if the source reads `Status.Conditions` for endpoint generation.

The transformer also restores `TypeMeta` (Kind/APIVersion) that informers strip,
making cached objects self-describing for FQDN templates and logging.

#### Indexers

Use `IndexerWithOptions[T]` when the source needs to look up objects by
annotation/label filters, or by a label value that references another resource:

```go
// Filter by annotation/label selectors (index key = object's own namespace/name)
if err = myInformer.Informer().AddIndexers(informers.IndexerWithOptions[*v1.MyResource](
    informers.IndexSelectorWithAnnotationFilter(annotationFilter),
    informers.IndexSelectorWithLabelSelector(labelSelector),
)); err != nil {
    return nil, err
}

// Index by a label's value rather than the object's own name
// (e.g. EndpointSlices indexed by the Service they belong to)
if err = myInformer.Informer().AddIndexers(informers.IndexerWithOptions[*discoveryv1.EndpointSlice](
    informers.IndexSelectorWithLabelKey(discoveryv1.LabelServiceName),
)); err != nil {
    return nil, err
}
```

All indexed objects are stored under the `informers.IndexWithSelectors` key.

##### When to add a new `IndexSelectorWith*` option

Add a new option only when **all three** conditions hold:

1. **No existing option expresses the needed logic** — the current set covers
   annotation filter, label selector match, and label-value-as-key. If none
   of those express your filter or key derivation, a new option is warranted.
2. **The pattern is reusable across more than one source or resource type** —
   if the logic is tightly coupled to a single type's fields, write a dedicated
   standalone `cache.Indexers` function instead of extending the shared options
   struct.
3. **The logic operates solely on `metav1.Object`** — `IndexerWithOptions`
   works generically via labels, annotations, name, and namespace. If you need
   type-specific fields (e.g. `Spec.Selector` on a Service), the logic requires
   type assertions that break the generic design; write a dedicated indexer
   instead.

#### Ordering

Always configure the informer in this order before starting the factory:

```go
if err = myInformer.Informer().SetTransform(...); err != nil { return nil, err } // 1. how objects are stored
if err = myInformer.Informer().AddIndexers(...); err != nil { return nil, err }  // 2. how objects are looked up
_, _ = myInformer.Informer().AddEventHandler(...)                                // 3. who is notified
factory.Start(ctx.Done())                                                        // 4. start — SetTransform errors after this
```

## Usage

You can choose any combination of sources and providers on the command line.
Given a cluster on AWS you would most likely want to use the Service and Ingress Source in combination with the AWS provider.
`Service` + `InMemory` is useful for testing your service collecting functionality, whereas `Fake` + `Google` is useful for testing that the Google provider behaves correctly, etc.

## Providers

Providers are an abstraction over any kind of sink for desired Endpoints, e.g.:

* storing them in Google Cloud DNS
* printing them to stdout for testing purposes
* fanning out to multiple nested providers

The `Provider` interface has two methods: `Records` and `ApplyChanges`.
`Records` should return all currently existing DNS records converted to Endpoint objects as a flat list.
Upon receiving a change set (via an object of `plan.Changes`), `ApplyChanges` should translate these to the provider specific actions in order to persist them in the provider's storage.

```go
type Provider interface {
 Records() ([]*endpoint.Endpoint, error)
 ApplyChanges(changes *plan.Changes) error
}
```

The interface tries to be generic and assumes a flat list of records for both functions. However, many providers scope records into zones.
Therefore, the provider implementation has to do some extra work to return that flat list. For instance, the AWS provider fetches the list of all hosted zones before it can return or apply the list of records.
If the provider has no concept of zones or if it makes sense to cache the list of hosted zones it is happily allowed to do so.
Furthermore, the provider should respect the `--domain-filter` flag to limit the affected records by a domain suffix. For instance, the AWS provider filters out all hosted zones that doesn't match that domain filter.

All providers live in package `provider`.

* `GoogleProvider`: returns and creates DNS records in Google Cloud DNS
* `AWSProvider`: returns and creates DNS records in AWS Route 53
* `AzureProvider`: returns and creates DNS records in Azure DNS
* `InMemoryProvider`: Keeps a list of records in local memory
