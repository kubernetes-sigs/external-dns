# Gateway sources

The gateway-grcproute, gateway-httproute, gateway-tcproute, gateway-tlsroute, and gateway-udproute
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

The targets of the DNS entries created from a \*Route are sourced from the following places:

1. If a matching parent Gateway has an `external-dns.alpha.kubernetes.io/target` annotation, uses
   the values from that.

2. Otherwise, iterates over that parent Gateway's `status.addresses`,
   adding each address's `value`.

The targets from each parent Gateway matching the \*Route are then combined and de-duplicated.

## Dualstack Routes

Gateway resources may be served from an external-loadbalancer which may support both IPv4 and "dualstack" (both IPv4 and IPv6) interfaces.
External DNS Controller uses the `external-dns.alpha.kubernetes.io/dualstack` annotation to determine this. If this annotation is
set to `true` then ExternalDNS will create two records (one A record
and one AAAA record) for each hostname associated with the Route resource.

Example:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/dualstack: "true"
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

The above HTTPRoute resource is backed by a dualstack Gateway.
ExternalDNS will create both an A `echoserver.example.org` record and
an AAAA record of the same name, that each are aliases for the same LB.
