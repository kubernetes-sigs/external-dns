# Gateway sources

The gateway-grpcroute, gateway-httproute, gateway-tcproute, gateway-tlsroute, and gateway-udproute
sources create DNS entries based on their respective `gateway.networking.k8s.io` resources.

## Filtering the Routes considered

These sources support the `--label-filter` flag, which filters \*Route resources
by a set of labels.

## Domain names

To calculate the Domain names created from a *Route, this source first collects a set
of [domain names from the *Route](#domain-names-from-route).

It then iterates over each of the `status.parents` with
a [matching Gateway](#matching-gateways) and at least one [matching listener](#matching-listeners).
For each matching listener, if the
listener has a `hostname`, it narrows the set of domain names from the \*Route to the portion
that overlaps the `hostname`. If a matching listener does not have a `hostname`, it uses
the un-narrowed set of domain names.

### Domain names from Route

The set of domain names from a \*Route is sourced from the following places:

- If the \*Route is a GRPCRoute, HTTPRoute, or TLSRoute, adds each of the`spec.hostnames`.

- Adds the hostnames from any `external-dns.kubernetes.io/hostname` annotation on the \*Route.
  This behavior is suppressed if the `--ignore-hostname-annotation` flag was specified.

- If no endpoints were produced by the previous steps
  or the `--combine-fqdn-annotation` flag was specified, then adds hostnames
  generated from any`--fqdn-template` flag.

- If no endpoints were produced by the previous steps, each
  attached Gateway listener will use its `hostname`, if present.

### Matching Gateways

Matching Gateways are discovered by iterating over the \*Route's `status.parents`:

- Ignores parents with a `parentRef.group` other than
  `gateway.networking.k8s.io` or a `parentRef.kind` other than `Gateway`.

- If the `--gateway-name` flag was specified, ignores parents with a `parentRef.name` other than the
  specified value.

  For example, given the following HTTPRoute:

    ```yaml
    apiVersion: gateway.networking.k8s.io/v1
    kind: HTTPRoute
    metadata:
      name: echo
    spec:
      hostnames:
        - echoserver.example.org
      parentRefs:
        - group: networking.k8s.io
          kind: Gateway
          name: internal
    ---
    apiVersion: gateway.networking.k8s.io/v1
    kind: HTTPRoute
    metadata:
      name: echo2
    spec:
      hostnames:
        - echoserver2.example.org
      parentRefs:
        - group: networking.k8s.io
          kind: Gateway
          name: external
    ```

  And using the `--gateway-name=external` flag, only the `echo2` HTTPRoute will be considered for DNS entries.

- If the `--gateway-namespace` flag was specified, ignores parents with a `parentRef.namespace` other
  than the specified value.

- If the `--gateway-label-filter` flag was specified, ignores parents whose Gateway does not match the
  specified label filter.

- Ignores parents whose Gateway either does not exist or has not accepted the route.

### Matching listeners

Iterates over all listeners for the parent's `parentRef.sectionName`:

- Ignores listeners whose `protocol` field does not match the kind of the \*Route per the following table:

| kind      | protocols   |
| --------- | ----------- |
| GRPCRoute | HTTP, HTTPS |
| HTTPRoute | HTTP, HTTPS |
| TCPRoute  | TCP         |
| TLSRoute  | TLS         |
| UDPRoute  | UDP         |

- If the parent's `parentRef.port` port is specified, ignores listeners without a matching `port`.

- Ignores listeners which specify an `allowedRoutes` which does not allow the route.

### API group defaulting

`parentRef.group`, and the `group` of each entry in a listener's `allowedRoutes.kinds`, are optional
and default to `gateway.networking.k8s.io` when they are omitted. An explicit empty string is not
the same thing: it names the core API group, which holds no Gateway and no \*Route kind.

| written as                         | resolves to                 |
| ---------------------------------- | --------------------------- |
| `group` omitted                    | `gateway.networking.k8s.io` |
| `group: gateway.networking.k8s.io` | `gateway.networking.k8s.io` |
| `group: ""`                        | the core API group          |

Earlier versions read an explicit `group: ""` as `gateway.networking.k8s.io`. A `parentRef` written
that way resolved to the same-named Gateway, and a listener whose `allowedRoutes.kinds` was written
that way accepted Gateway API \*Routes. Both now follow the API, so records that existed only
because of the old reading are no longer generated.

With `--policy=sync` and a TXT registry, those records and the ownership TXT records that go with
them are deleted on the next reconcile. Before upgrading, look for `parentRefs` and
`allowedRoutes.kinds` entries with an explicit empty group, and either set the group to
`gateway.networking.k8s.io` or drop the field wherever a Gateway API object was meant.

Do not change `group: ""` everywhere. It is the correct way to name an object in the core API
group, such as the Service in the `backendRef` of the example below. Only references that were
meant to name a Gateway or a \*Route kind are affected.

`--policy=upsert-only` suppresses the deletion during a rollout, but it does not clean the stale
records up afterwards and it still allows existing record sets to be updated, so it is a way to
stage the change rather than a guarantee that nothing moves.

## Targets

The targets of the DNS entries created from a \*Route are sourced from the following places:

1. If a matching parent Gateway has an `external-dns.kubernetes.io/target` annotation, uses
   the values from that.

2. Otherwise, iterates over that parent Gateway's `status.addresses`,
   adding each address's `value`.

The targets from each parent Gateway matching the \*Route are then combined and de-duplicated.

## Dualstack Routes

Gateway resources may be served from an external-loadbalancer which may support
both IPv4 and "dualstack" (both IPv4 and IPv6) interfaces. When using the AWS
Route53 provider, External DNS Controller will always create both A and AAAA
alias DNS records by default, regardless of whether the load balancer is dual
stack or not.

## Example

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: echo
spec:
  hostnames:
    - echoserver.example.org
  rules:
    - backendRefs:
        - group: ""
          kind: Service
          name: echo
          port: 1027
          weight: 1
      matches:
        - path:
            type: PathPrefix
            value: /echo
```
