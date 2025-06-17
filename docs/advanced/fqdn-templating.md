# FQDN Templating Guide

## What is FQDN Templating?

**FQDN templating** is a feature that allows to dynamically construct Fully Qualified Domain Names (FQDNs) using a Go templating engine.
Instead of relying solely on annotations or static names, you can use metadata from Kubernetes objects—such as service names, namespaces, and labels—to generate DNS records programmatically and dynamically.

This is useful for:

- Creating consistent naming conventions across environments.
- Reducing boilerplate annotations.
- Supporting multi-tenant or dynamic environments.
- Migrating from one DNS scheme to another
- Supporting multiple variants, such as a regional one and then one that doesn't or similar.

## How It Works

ExternalDNS has a flag: `--fqdn-template`, which defines a Go template for rendering the desired DNS names.

The template uses the following data from the source object (e.g., a `Service` or `Ingress`):

| Field         | Description                                                       |
|:--------------|:------------------------------------------------------------------|
| `Name`        | Name of the object (e.g., service)                                |
| `Namespace`   | Namespace of the object                                           |
| `Labels`      | Map of labels applied to the object                               |
| `Annotations` | Map of annotations                                                |
| `TargetName`  | For `Service`, it's the service name; for `Ingress`, the hostname |
| `Endpoint`    | Contains more contextual endpoint info, such as IP/target         |
| `Controller`  | Controller type (optional)                                        |

## Supported Sources

<!-- TODO: generate from code -->

| Source                 | Description                                                     | FQDN Supported | FQDN Combine |
|:-----------------------|:----------------------------------------------------------------|:--------------:|:------------:|
| `ambassador-host`      | Queries Ambassador Host resources for endpoints.                |       ❌        |      ❌       |
| `cloudfoundry`         | Queries Cloud Foundry resources for endpoints.                  |       ❌        |      ❌       |
| `connector`            | Queries a custom connector source for endpoints.                |       ❌        |      ❌       |
| `contour-httpproxy`    | Queries Contour HTTPProxy resources for endpoints.              |       ✅        |      ✅       |
| `crd`                  | Queries Custom Resource Definitions (CRDs) for endpoints.       |       ❌        |      ❌       |
| `empty`                | Uses an empty source, typically for testing or no-op scenarios. |       ❌        |      ❌       |
| `f5-transportserver`   | Queries F5 TransportServer resources for endpoints.             |       ❌        |      ❌       |
| `f5-virtualserver`     | Queries F5 VirtualServer resources for endpoints.               |       ❌        |      ❌       |
| `fake`                 | Uses a fake source for testing purposes.                        |       ❌        |      ❌       |
| `gateway-grpcroute`    | Queries GRPCRoute resources from the Gateway API.               |       ✅        |      ❌       |
| `gateway-httproute`    | Queries HTTPRoute resources from the Gateway API.               |       ✅        |      ❌       |
| `gateway-tcproute`     | Queries TCPRoute resources from the Gateway API.                |       ✅        |      ❌       |
| `gateway-tlsroute`     | Queries TLSRoute resources from the Gateway API.                |       ❌        |      ❌       |
| `gateway-udproute`     | Queries UDPRoute resources from the Gateway API.                |       ❌        |      ❌       |
| `gloo-proxy`           | Queries Gloo Proxy resources for endpoints.                     |       ❌        |      ❌       |
| `ingress`              | Queries Kubernetes Ingress resources for endpoints.             |       ✅        |      ✅       |
| `istio-gateway`        | Queries Istio Gateway resources for endpoints.                  |       ✅        |      ✅       |
| `istio-virtualservice` | Queries Istio VirtualService resources for endpoints.           |       ✅        |      ✅       |
| `kong-tcpingress`      | Queries Kong TCPIngress resources for endpoints.                |       ❌        |      ❌       |
| `node`                 | Queries Kubernetes Node resources for endpoints.                |       ✅        |      ✅       |
| `openshift-route`      | Queries OpenShift Route resources for endpoints.                |       ✅        |      ✅       |
| `pod`                  | Queries Kubernetes Pod resources for endpoints.                 |       ✅        |      ✅       |
| `service`              | Queries Kubernetes Service resources for endpoints.             |       ✅        |      ✅       |
| `skipper-routegroup`   | Queries Skipper RouteGroup resources for endpoints.             |       ✅        |      ✅       |
| `traefik-proxy`        | Queries Traefik IngressRoute resources for endpoints.                  |       ❌        |      ❌       |

## Custom Functions

<!-- TODO: generate from code -->

| Function     | Description                                           | Example                                                                          |
|:-------------|:------------------------------------------------------|:---------------------------------------------------------------------------------|
| `contains`   | Check if `substr` is in `string`                      | `{{ contains "hello" "ell" }} → true`                                            |
| `isIPv4`     | Validate an IPv4 address                              | `{{ isIPv4 "192.168.1.1" }} → true`                                              |
| `isIPv6`     | Validate an IPv6 address (including IPv4-mapped IPv6) | `{{ isIPv6 "2001:db8::1" }} → true`<br/>`{{ isIPv6 "::FFFF:192.168.1.1" }}→true` |
| `replace`    | Replace `old` with `new`                              | `{{ replace "hello" "l" "w" }} → hewwo`                                          |
| `trim`       | Remove leading and trailing spaces                    | `{{ trim "  hello  " }} → hello`                                                 |
| `toLower`    | Convert to lowercase                                  | `{{ toLower "HELLO" }} → hello`                                                  |
| `trimPrefix` | Remove the leading `prefix`                           | `{{ trimPrefix "pre" "prefix" }} → fix`                                          |
| `trimSuffix` | Remove the trailing `suffix`                          | `{{ trimSuffix "fix" "suffix" }} → suf`                                          |

---

## Example Usage

> These examples should provide a solid foundation for implementing FQDN templating in your ExternalDNS setup.
> If you have specific requirements or encounter issues, feel free to explore the issues or update this guide.

### Basic Usage

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: my-namespace
```

```sh
external-dns \
  --provider=aws \
  --source=service \
  --fqdn-template="{{ .Name }}.example.com,{{ .Name }}.{{ .Namespace }}.example.tld"

# This will result in DNS entries like
>route53> my-service.example.com
>route53> my-service.my-namespace.example.tld
```

### With Namespace

```yml
---
apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: default
---
apiVersion: v1
kind: Service
metadata:
  name: other-service
  namespace: kube-system
```

```yml
args:
  --fqdn-template="{{.Name}}.{{.Namespace}}.example.com"

# This will result in DNS entries like
# route53> my-service.default.example.com
# route53> other-service.kube-system.example.com
```

### Using Labels  in Templates

You can also utilize labels in your FQDN templates to create more dynamic DNS entries. Assuming your service has:

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  labels:
    environment: staging
```

```yml
args:
  --fqdn-template="{{ .Labels.environment }}.{{ .Name }}.example.com"

# This will result in DNS entries like
# route53> staging.my-service.example.com
```

### Multiple FQDN Templates

ExternalDNS allows specifying multiple FQDN templates, which can be useful when you want to create multiple DNS entries for a single service or ingress.

> Be cautious, as this will create multiple DNS records per resource, potentially increasing the number of API calls to your DNS provider.

```yml
args:
  --fqdn-template={{.Name}}.example.com,{{.Name}}.svc.example.com
```

### Conditional Templating combined with Annotations processing

In scenarios where you want to conditionally generate FQDNs based on annotations, you can use Go template functions like or to provide defaults.

```yml
args:
  - --combine-fqdn-annotation # this is required to combine FQDN templating and annotation processing
  - --fqdn-template={{ or .Annotations.dns "invalid" }}.example.com
  - --exclude-domains=invalid.example.com
```

### Using Annotations for FQDN Templating

This example demonstrates how to use annotations in Kubernetes objects to dynamically generate Fully Qualified Domain Names (FQDNs) using the --fqdn-template flag in ExternalDNS.

The Service object includes an annotation dns.company.com/label with the value my-org-tld-v2. This annotation is used as part of the FQDN template to construct the DNS name.

```yml
apiVersion: v1
kind: Service
metadata:
  name: nginx-v2
  namespace: my-namespace
  annotations:
    dns.company.com/label: my-org-tld-v2
spec:
  type: ClusterIP
  clusterIP: None
```

The --fqdn-template flag is configured to use the annotation value (dns.company.com/label) and append the namespace and a custom domain (company.local) to generate the FQDN.

```yml
args:
  --source=service
  --fqdn-template='{{ index .ObjectMeta.Annotations "dns.company.com/label" }}.{{ .Namespace }}.company.local'

# For the given Service object, the resulting FQDN will be:
# route53> my-org-tld-v2.my-namespace.company.local
```

### DNS Scheme Migration

If you're transitioning from one naming convention to another (e.g., from svc.cluster.local to svc.example.com), --fqdn-template allows you to generate the new records alongside or in place of the old ones — without requiring changes to your Kubernetes manifests.

```yml
args:
- --fqdn-template='{{.Name}}.new-dns.example.com'
```

This helps automate DNS record migration while maintaining service continuity.

### Multi-Variant Domain Support

You can also support regional variants or multi-tenant architectures, where the same service is deployed to different regions or environments:

```yaml
--fqdn-template='{{ .Name }}.{{ .Labels.env }}.{{ .Labels.region }}.example.com, {{ if eq .Labels.env "prod" }}{{ .Name }}.my-company.tld{{ end }}'
```

With additional context (e.g., annotations), this can produce FQDNs like:

```yml
api.prod.us-east-1.example.com
api.my-company.tld
```

This is helpful in scenarios such as:

- Blue/green deployments across domains
- Staging vs. production resolution
- Multi-cloud or multi-region failover strategies

## Tips

- If `--fqdn-template` is specified, ExternalDNS ignores any `external-dns.alpha.kubernetes.io/hostname` annotations.
- You must still ensure the resulting FQDN is valid and unique.
- Since Go templates can be error-prone, test your template with simple examples before deploying. Mismatched field names or nil values (e.g., missing labels) will result in errors or skipped entries.

## FaQ

### Can I specify multiple global FQDN templates?

Yes, you can. Pass in a comma separated list to --fqdn-template. Beware this will double (triple, etc) the amount of DNS entries based on how many services, ingresses and so on you have and will get you faster towards the API request limit of your DNS provider.

### Where to find template syntax

- [Go template syntax](https://pkg.go.dev/text/template)
- [Go func builtins](https://github.com/golang/go/blob/master/src/text/template/funcs.go#L39-L63)

### FQDN Templating, Helm and improper templating syntax

The user encountered errors due to improper templating syntax:

```yml
extraArgs:
  - --fqdn-template={{name}}.uat.example.com
```

The correct syntax should include a dot prefix: `{{ .Name }}`.
Additionally, when using Helm's `tpl` function, it's necessary to escape the braces to prevent premature evaluation:

```yml
extraArgs:
  - --fqdn-template={{ `{{ .Name }}.uat.example.com` }}
```

### Handling Subdomain-Only Hostnames

In [Issue #1872](https://github.com/kubernetes-sigs/external-dns/issues/1872), it was observed that ExternalDNS ignores the `--fqdn-template` when the ingress host field is set to a subdomain (e.g., foo) without a full domain.
The expectation was that the template would still apply, generating entries like `foo.bar.example.com.`
This highlights a limitation to be aware of when designing FQDN templates.

> :warning: This is currently not supported ! User would expect external-dns to generate a dns record according to the fqdnTemplate
> e.g. if the ingress name: foo and host: foo is created while fqdnTemplate={{.Name}}.bar.example.com then a dns record foo.bar.example.com should be created

```yml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
spec:
  rules:
  - host: foo
    http:
      paths:
      - backend:
          serviceName: foo
          servicePort: 80
        path: /
```

### Combining FQDN Template with Annotations

In [Issue #3318](https://github.com/kubernetes-sigs/external-dns/issues/3318), a question was raised about the interaction between --fqdn-template and --combine-fqdn-annotation.
The discussion clarified that when both flags are used, ExternalDNS combines the FQDN generated from the template with the annotation value, providing flexibility in DNS name construction.

### Using Annotations for Dynamic FQDNs

In [Issue #2627](https://github.com/kubernetes-sigs/external-dns/issues/2627), a user aimed to generate DNS entries based on ingress annotations:

```yml
args:
  - --fqdn-template={{.Annotations.hostname}}.example.com
  - --combine-fqdn-annotation
  - --domain-filter=example.com
```

By setting the hostname annotation in the ingress resource, ExternalDNS constructs the FQDN accordingly. This approach allows for dynamic DNS entries without hardcoding hostnames.

### Using a Node's Addresses for FQDNs

```yml
args:
  - --fqdn-template="{{range .Status.Addresses}}{{if and (eq .Type \"ExternalIP\") (isIPv4 .Address)}}{{.Address | replace \".\" \"-\"}}{{break}}{{end}}{{end}}.example.com"
```

This is a complex template that iternates through a list of a Node's Addresses and creates a FQDN with public IPv4 addresses.
