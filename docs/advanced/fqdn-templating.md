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

| Field         | Description                                      | How to Access                                          |
|:--------------|:-------------------------------------------------|:-------------------------------------------------------|
| `Kind`        | Object kind (e.g., `Service`, `Pod`, `Ingress`)  | `{{ .Kind }}`                                          |
| `APIVersion`  | API version (e.g., `v1`, `networking.k8s.io/v1`) | `{{ .APIVersion }}`                                    |
| `Name`        | Name of the object (e.g., service)               | `{{ .Name }}`                                          |
| `Namespace`   | Namespace of the object                          | `{{ .Namespace }}`                                     |
| `Labels`      | Map of labels applied to the object              | `{{ .Labels.key }}` or `{{ index .Labels "key" }}`     |
| `Annotations` | Map of annotations                               | `{{ index .Annotations "key" }}`                       |
| `Spec`        | Object spec with type-specific fields            | `{{ .Spec.Type }}`, `{{ index .Spec.Selector "app" }}` |
| `Status`      | Object status with type-specific fields          | `{{ .Status.LoadBalancer.Ingress }}`                   |

To explore all available fields for an object type, use `kubectl explain`:

```bash
# View all fields for a Service recursively.
kubectl explain service --api-version=v1 --recursive

# View all fields for a Ingress recursively.
kubectl explain ingress --api-version=networking.k8s.io/v1 --recursive

# View a specific field path. The dot notation is for field path.
kubectl explain service.spec.selector
kubectl explain pod.spec.containers
```

## Supported Sources

<!-- TODO: generate from code -->

| Source                 | Description                                                     | FQDN Supported | FQDN Combine |
|:-----------------------|:----------------------------------------------------------------|:--------------:|:------------:|
| `ambassador-host`      | Queries Ambassador Host resources for endpoints.                |       ❌        |      ❌       |
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

### Using Kind for Conditional Templating

When processing multiple resource types, use `.Kind` to apply templates conditionally:

```yml
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ .Name }}.svc.example.com{{ end }}'

# Only Services will get DNS entries, Pods and other resources will be skipped
```

You can also handle multiple kinds in one template:

```yml
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ .Name }}.svc.example.com{{ end }}{{ if eq .Kind "Pod" }}{{ .Name }}.pod.example.com{{ end }}'
```

### Using Spec Fields

Access type-specific spec fields for advanced filtering:

```yml
# Only ExternalName services
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ if eq .Spec.Type "ExternalName" }}{{ .Name }}.external.example.com{{ end }}{{ end }}'
```

```yml
apiVersion: v1
kind: Service
metadata:
  name: web-frontend
spec:
  selector:
    app: nginx        # This selector will be used in the FQDN
    tier: frontend
  ports:
    - port: 80
---
apiVersion: v1
kind: Service
metadata:
  name: database
spec:
  selector:
    tier: backend     # Won't generate FQDN - no "app" key in selector
  ports:
    - port: 5432
```

```yml
# Services with specific selector
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ if index .Spec.Selector "app" }}{{ .Name }}.{{ index .Spec.Selector "app" }}.example.com{{ end }}{{ end }}'

# Result for web-frontend: web-frontend.nginx.example.com
# Result for database: (no FQDN generated - selector has no "app" key)
```

### Iterating Over Labels with Range

Use `range` to iterate over labels and generate multiple FQDNs:

```yml
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ range $key, $value := .Labels }}{{ if contains $key "app" }}{{ $.Name }}.{{ $value }}.example.com{{ printf "," }}{{ end }}{{ end }}{{ end }}'
```

This generates an FQDN for each label key containing "app". Note:

- `$key` and `$value` are the label key/value pairs
- `$.Name` accesses the root object's Name (use `$` inside `range`)
- `{{ printf "," }}` separates multiple FQDNs

### Working with Annotations

Access a specific annotation:

```yml
args:
  --fqdn-template='{{ index .Annotations "dns.example.com/hostname" }}.example.com'
```

Iterate over annotations and filter by key:

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  annotations:
    dns.example.com/primary: api.example.com
    dns.example.com/secondary: api-backup.example.com
    kubernetes.io/ingress-class: nginx  # Won't match - key doesn't contain "dns.example.com/"
```

```yml
args:
  --fqdn-template='{{ range $key, $value := .Annotations }}{{ if contains $key "dns.example.com/" }}{{ $value }}{{ printf "," }}{{ end }}{{ end }}'

# Captures all annotations with keys containing "dns.example.com/"
# Result: api.example.com, api-backup.example.com
```

Filter annotations by value:

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  annotations:
    custom/hostname: api.example.com
    custom/alias: www.example.com
    custom/internal: internal.local  # Won't match - value doesn't contain ".example.com"
```

```yml
args:
  --fqdn-template='{{ range $key, $value := .Annotations }}{{ if contains $value ".example.com" }}{{ $value }}{{ printf "," }}{{ end }}{{ end }}'

# Captures all annotation values containing ".example.com"
# Result: api.example.com, www.example.com
```

Combine annotation key and value filters:

```yml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  annotations:
    dns/primary: api.example.com
    dns/secondary: api-backup.example.com
    other/hostname: internal.other.org  # Won't match - value doesn't contain "example.com"
    logging/level: debug                # Won't match - key doesn't contain "dns/"
```

```yml
args:
  --fqdn-template='{{ if eq .Kind "Service" }}{{ range $k, $v := .Annotations }}{{ if and (contains $k "dns/") (contains $v "example.com") }}{{ $v }}{{ printf "," }}{{ end }}{{ end }}{{ end }}'

# Result: api.example.com, api-backup.example.com
```

### Combining Kind and Label Filters

Filter by both Kind and label values:

```yml
args:
  --fqdn-template='{{ if eq .Kind "Pod" }}{{ range $k, $v := .Labels }}{{ if and (contains $k "app") (contains $v "my-service-") }}{{ $.Name }}.{{ $v }}.example.com{{ printf "," }}{{ end }}{{ end }}{{ end }}'

# Generates FQDNs only for Pods with labels like app1=my-service-123
# Result: pod-name.my-service-123.example.com
```

### Multi-Variant Domain Support

You can also support regional variants or multi-tenant architectures, where the same service is deployed to different regions or environments:

```yaml
--fqdn-template='{{ .Name }}.{{ .Labels.env }}.{{ .Labels.region }}.example.com, {{ if eq .Labels.env "prod" }}{{ .Name }}.my-company.tld{{ end }}'

# Generates FQDNs for resources with labels env and region
# For a Service named "api" with labels env=prod, region=us-east-1:
# Result: api.prod.us-east-1.example.com, api.my-company.tld

# For a Service named "api" with labels env=staging, region=eu-west-1:
# Result: api.staging.eu-west-1.example.com
```

This is helpful in scenarios such as:

- Blue/green deployments across domains
- Staging vs. production resolution
- Multi-cloud or multi-region failover strategies

## Tips

- If `--fqdn-template` is specified, ExternalDNS ignores any `external-dns.alpha.kubernetes.io/hostname` annotations.
- You must still ensure the resulting FQDN is valid and unique.
- Since Go templates can be error-prone, test your template with simple examples before deploying. Mismatched field names or nil values (e.g., missing labels) will result in errors or skipped entries.

## FAQ

### Can I specify multiple global FQDN templates?

Yes, you can. Pass in a comma separated list to --fqdn-template. Beware this will double (triple, etc) the amount of DNS entries based on how many services, ingresses and so on you have and will get you faster towards the API request limit of your DNS provider.

### Where to find template syntax

- [Go template syntax](https://pkg.go.dev/text/template) - Official reference for template syntax, actions, and pipelines
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
