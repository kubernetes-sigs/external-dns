---
tags:
  - sources
  - providers
  - contributing
  - informers
---

# Sources and Providers

ExternalDNS supports swapping out endpoint **sources** and DNS **providers** and both sides are pluggable. There currently exist multiple sources for different provider implementations.

**Usage**

You can choose any combination of sources and providers on the command line.
Given a cluster on AWS you would most likely want to use the Service and Ingress Source in combination with the AWS provider.
`Service` + `InMemory` is useful for testing your service collecting functionality, whereas `Fake` + `Google` is useful for testing that the Google provider behaves correctly, etc.

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

Call `MustSetTransform` on every informer **before** `factory.Start()`. Use
`TransformerWithOptions[T]` with the options appropriate for the resource:

```go
informers.MustSetTransform(myInformer.Informer(), informers.TransformerWithOptions[*v1.MyResource](
    informers.TransformRemoveManagedFields(),     // always — can be megabytes per object
    informers.TransformRemoveLastAppliedConfig(), // always — full JSON snapshot stored as annotation
    informers.TransformRemoveStatusConditions(),  // when Status.Conditions is not used by the source
    informers.TransformKeepAnnotationPrefix("external-dns.alpha.kubernetes.io/"), // when only specific annotations are needed
))
```

`TransformRemoveManagedFields` and `TransformRemoveLastAppliedConfig` should be
applied to every informer. `TransformRemoveStatusConditions` can be omitted only
if the source reads `Status.Conditions` for endpoint generation.

The transformer also restores `TypeMeta` (Kind/APIVersion) that informers strip,
making cached objects self-describing for FQDN templates and logging. For types
registered in the core k8s scheme (Services, Pods, Nodes, Ingresses, …) the full
Group/Version/Kind is set. For CRD-backed types (e.g. Istio Gateway) the Kind is
derived from the Go struct name via reflection; Group and Version remain empty.

#### Indexers

Use `IndexerWithOptions[T]` when the source needs to look up objects by
annotation/label filters, or by a label value that references another resource:

```go
// Filter by annotation/label selectors (index key = object's own namespace/name)
informers.MustAddIndexers(myInformer.Informer(), informers.IndexerWithOptions[*v1.MyResource](
    informers.IndexSelectorWithAnnotationFilter(annotationFilter),
    informers.IndexSelectorWithLabelSelector(labelSelector),
))

// Index by a label's value rather than the object's own name
// (e.g. EndpointSlices indexed by the Service they belong to)
informers.MustAddIndexers(myInformer.Informer(), informers.IndexerWithOptions[*discoveryv1.EndpointSlice](
    informers.IndexSelectorWithLabelKey(discoveryv1.LabelServiceName),
))
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

#### Event handlers

Use `MustAddEventHandler` instead of the `_, _ = informer.AddEventHandler(...)` idiom.
`AddEventHandler` only returns an error when the informer has already been stopped. The
helper logs a warning rather than panicking, because unlike `SetTransform`/`AddIndexers`,
event handlers are also registered at runtime (from `Source.AddEventHandler`), where a
stopped informer can be a transient shutdown condition rather than a programming error:

```go
informers.MustAddEventHandler(myInformer.Informer(), informers.DefaultEventHandler())
```

#### Ordering

Always configure the informer in this order before starting the factory:

```go
informers.MustSetTransform(myInformer.Informer(), ...)       // 1. how objects are stored
informers.MustAddIndexers(myInformer.Informer(), ...)        // 2. how objects are looked up
informers.MustAddEventHandler(myInformer.Informer(), ...)    // 3. who is notified
factory.Start(ctx.Done())                                    // 4. start — Must* helpers panic after this
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

### Implementing GetDomainFilter

`GetDomainFilter()` is a method on the `Provider` interface. The default implementation in
`BaseProvider` returns an empty filter with no effect. Providers can override it to
contribute an additional domain constraint to the reconcile plan, on top of whatever the
user configured via `--domain-filter`.

#### How the controller uses it

Each reconcile cycle, the controller builds a plan combining two filters:

```go
DomainFilter: endpoint.MatchAllDomainFilters{c.DomainFilter, registryFilter}
```

* `c.DomainFilter` — from the `--domain-filter` CLI flag (user-supplied)
* `registryFilter` — the value returned by `provider.GetDomainFilter()`

`MatchAllDomainFilters` is a logical AND: a record must satisfy both to be included in the
plan. The provider filter acts as an additional, provider-side constraint on top of whatever
the user configured.

#### When to leave the default

If your provider has no concept of zones, domains, or hosted zones — for example, a
provider backed by flat storage like etcd — the `BaseProvider` default is fine. Do not
override it just to echo `config.DomainFilter` back. For example, if the user runs with
`--domain-filter=example.com` and the provider returns the same value, the plan sees:

```go
MatchAllDomainFilters{example.com, example.com}  // same filter twice, no added value
```

This is functionally identical to the default and adds no protection.

#### When and how to override — the dynamic pattern

Override `GetDomainFilter()` when your provider has an authoritative list of zones,
domains, or hosted zones it manages — regardless of what the DNS provider calls them —
and can narrow the scope independently of what the user configured. Two concrete
benefits make this worthwhile:

**Protection without user configuration** — when no `--domain-filter` is set,
`BaseProvider` returns an empty filter and the controller has no domain constraint at all.
A dynamic override builds the constraint from zones the provider actually manages, so the
controller is scoped correctly even if the operator never sets a flag.

**The filter reflects reality, not intent** — `--domain-filter` expresses what the
operator wants to manage. `GetDomainFilter()` expresses what the provider actually manages
at runtime — zones that exist and are accessible with the current credentials. The
intersection of the two is tighter and safer than either alone.

For example, if `--domain-filter=example.com` is set but the provider only has access to
`api.example.com` and `prod.example.com`, a dynamic implementation scopes the plan to
exactly those two zones rather than anything under `example.com`.

The correct approach is to query your zone API at runtime and build the filter from the
zones your provider actually controls. `AWSProvider.GetDomainFilter()` is the canonical
example:

```go
func (p *MyProvider) GetDomainFilter() endpoint.DomainFilterInterface {
    zones, err := p.zones()
    if err != nil {
        return &endpoint.DomainFilter{}
    }
    // Apply your own configured filter to keep only zones this provider manages.
    filteredZones := applyDomainFilter(zones)

    names := make([]string, 0, len(zones))
    for _, z := range filteredZones {
        names = append(names, z.Name, "."+z.Name)
    }
    return endpoint.NewDomainFilter(names)
}
```

Each zone name is added twice — as a bare domain (`example.com`) and with a leading dot
(`.example.com`) — so the filter matches both exact records and subdomains.

For example, suppose the provider manages four zones:

```sh
api.example.com
prod.myapp.io
staging.myapp.io
legacy.internal.net
```

**Without `--domain-filter`** — the provider filter alone constrains the plan:

```go
MatchAllDomainFilters{
    <empty>,                                        // no CLI flag, matches everything
    [api.example.com, .api.example.com,
     prod.myapp.io,   .prod.myapp.io,
     staging.myapp.io, .staging.myapp.io,
     legacy.internal.net, .legacy.internal.net],   // only provider-managed zones
}
```

The controller will only touch records in those four zones. Any other zone in the cluster
is left untouched, even if records pointing to it appear in sources.

**With `--domain-filter=myapp.io`** — the two filters intersect:

```go
MatchAllDomainFilters{
    myapp.io,                                       // CLI flag
    [api.example.com, .api.example.com,
     prod.myapp.io,   .prod.myapp.io,
     staging.myapp.io, .staging.myapp.io,
     legacy.internal.net, .legacy.internal.net],
}
```

Only `prod.myapp.io` and `staging.myapp.io` satisfy both filters and are in scope.
`api.example.com` and `legacy.internal.net` are excluded by the CLI filter.

On error, return an empty `&endpoint.DomainFilter{}`. This has the same effect as the
`BaseProvider` default — the CLI filter becomes the sole authority. If the user specifies
a domain the provider does not manage, reconciliation will proceed against it. This is a
deliberate tradeoff: a temporary API failure should not block all reconciliation.

For example, if the provider manages `a.com` and `b.com` but the user sets
`--domain-filter=c.com`, a dynamic implementation produces an empty intersection —
the controller does nothing:

```go
MatchAllDomainFilters{
    c.com,               // CLI flag
    [a.com, .a.com,      // provider zones — no overlap with c.com
     b.com, .b.com],
}
```

With an empty `GetDomainFilter()` (default or error), only the CLI filter applies and
the controller attempts to reconcile `c.com` against a provider that does not manage it.

#### Zone name formatting

Check the format your provider's API returns for zone names before passing them to
`endpoint.NewDomainFilter`. Some APIs include a trailing dot (`"example.com."`), which
must be stripped first:

```go
// API returns:  "foo.example.com."
// Filter needs: "foo.example.com"
name := strings.TrimSuffix(z.Name, ".")
names = append(names, name, "."+name)
```

#### Summary

| Implementation                        | `--domain-filter` unset                    | `--domain-filter` set                        |
|---------------------------------------|--------------------------------------------|----------------------------------------------|
| `BaseProvider` default                | No additional constraint                   | User filter applied                          |
| Static (echoes `config.DomainFilter`) | No additional constraint (same as default) | Same filter applied twice — redundant        |
| Dynamic (`ListZones` + filter)        | Provider-managed zones constrain the plan  | Intersection of user filter + provider zones |

The dynamic approach is what gives `GetDomainFilter()` its value: when no `--domain-filter`
is set, it prevents the controller from touching records in zones the provider does not
manage.

#### Testing

`GetDomainFilter()` must have a unit test. See `TestAWSProvider_GetDomainFilter` for a
reference. At minimum, test that:

* Zone names are correctly mapped to filter entries (including the leading-dot variant)
* An error from `ListZones` returns an empty `DomainFilter` gracefully

## Provider Blueprints

The `provider/blueprint` package contains reusable building blocks for provider
implementations. Using them keeps providers consistent and avoids reimplementing
solved problems.

### ZoneCache

`ZoneCache[T]` is a generic, thread-safe TTL cache for zone, domain, or hosted zone data.
See `provider/blueprint/zone_cache.go` for the full API and godoc.

**Reduced API pressure** — listing zones, domains, or hosted zones is called on every
reconcile cycle, but they are rarely created or deleted. Caching the result for a
configurable TTL means the provider only hits the API when the cache has expired, rather
than on every loop.

**Consistent behaviour across providers** — thread safety, TTL logic, and the
disable-via-zero behaviour are implemented and tested once in `blueprint`. Providers that
use `ZoneCache` behave the same way, reducing drift between implementations over time.

The typical usage pattern — taken from `AWSProvider.zones()` — is:

```go
// On the provider struct:
zonesCache *blueprint.ZoneCache[map[string]*MyZone]

// In the constructor:
zonesCache: blueprint.NewZoneCache[map[string]*MyZone](config.ZoneCacheDuration),

// In the zone/domain-listing method:
func (p *MyProvider) zones() (map[string]*MyZone, error) {
    if !p.zonesCache.Expired() {
        return p.zonesCache.Get(), nil
    }

    zones, err := p.client.ListZones()
    if err != nil {
        return nil, err
    }

    p.zonesCache.Reset(zones)
    return zones, nil
}
```

Full behaviour is documented in the `ZoneCache` godoc. The key contract to keep in mind
when implementing the pattern: `Get()` returns stale data after expiry rather than a zero
value — callers must check `Expired()` first and decide whether to refresh.

### Configuration flag

`ZoneCache` is controlled by a single shared flag:

| Flag                     | Default | Description                                  |
|--------------------------|---------|----------------------------------------------|
| `--zones-cache-duration` | `0s`    | Zone list cache TTL. Set to `0s` to disable. |

Add a `ZoneCacheDuration time.Duration` field to your provider config struct, wire it to
this flag in `pkg/apis/externaldns/types.go`, and pass it to `NewZoneCache` in the
constructor.
