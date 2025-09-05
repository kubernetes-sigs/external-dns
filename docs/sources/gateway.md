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

- Adds the hostnames from any `external-dns.alpha.kubernetes.io/hostname` annotation on the \*Route.
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

## Targets

Targets are derived from a combination of Route annotations, Gateway annotations, and
the Gateway's `status.addresses`. How these sources are combined is controlled (per Route)
by the optional strategy annotation:

```yaml
external-dns.alpha.kubernetes.io/target-strategy: <strategy>
```

Supported strategies (current behavior for Gateway *Route sources):

| Strategy         | Behavior | Fallback Order | Publishes Multiple Target Sets? |
|------------------|----------|----------------|----------------------------------|
| `route-preferred` (default when omitted or unrecognized) | Use Route targets if present; otherwise Gateway targets if present; otherwise addresses | Route → Gateway → Addresses | No (at most one set) |
| `route-only`     | Use Route targets only if present; otherwise fall back directly to addresses (Gateway targets are ignored even if set) | Route → Addresses | No |
| `gateway-only`   | Use Gateway targets only if present; otherwise fall back to addresses (Route targets are ignored) | Gateway → Addresses | No |
| `merge`          | Combine Route and Gateway targets (deduped). If neither annotation supplies targets, fall back to addresses | (Route ∪ Gateway) → Addresses | Yes |

Where the individual sources come from:

1. Route targets: values from a non-empty `external-dns.alpha.kubernetes.io/target` annotation on the Route.
2. Gateway targets: values from a non-empty `external-dns.alpha.kubernetes.io/target` annotation on the matching parent Gateway.
3. Addresses: each `value` in the parent Gateway's `status.addresses` field.

Notes & nuances:

- The empty string (`external-dns.alpha.kubernetes.io/target: ""`) on a Route or Gateway does not disable the other annotation; it simply contributes no targets.
- Under `route-only`, any Gateway target annotation is intentionally ignored even if populated.
- Under `gateway-only`, any Route target annotation is intentionally ignored.
- Under `merge`, duplicate targets between Route and Gateway are removed before creating DNS records.
- If (after applying the strategy rules) no annotation targets are selected, Gateway `status.addresses` are always used as a safety fallback to avoid producing zero endpoints for an otherwise valid attachment.

Example usages on a *Route:

```yaml
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target-strategy: route-only
```

Publishes the hostname(s) from the route using the Gateway addresses, ignoring target annotation on the Gateway.

```yaml
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target: canary.lb.example.net
    external-dns.alpha.kubernetes.io/target-strategy: route-preferred
```

Publishes only `canary.lb.example.net` (even if the Gateway also has a target annotation).

```yaml
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target-strategy: merge
```

Route has no target annotation; if the Gateway has `external-dns.alpha.kubernetes.io/target: edge.lb.example.net` the record will include that value; otherwise the Gateway addresses are used.

```yaml
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/target: direct.lb.example.net
    external-dns.alpha.kubernetes.io/target-strategy: merge
```

If the Gateway also has `external-dns.alpha.kubernetes.io/target: edge.lb.example.net` both targets are published (deduped if identical).

The combined targets from each matching parent Gateway are gathered per hostname and de-duplicated before generating DNS records.

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
