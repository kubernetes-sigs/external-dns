# Nomad Service Source

This tutorial explains how to configure ExternalDNS to use Nomad services as a source.

The Nomad service source reads metadata from services registered with the Nomad agent and uses it to determine DNS configuration.

By using `--source=nomad-service`, ExternalDNS can discover DNS endpoints from applications registered in a [HashiCorp Nomad](https://www.nomadproject.io/) cluster.
This allows DNS records to be dynamically created and updated based on the state of Nomad services, similar to how ExternalDNS works with Kubernetes resources.

## CLI Flags

This source respects these external-dns CLI flags:

```bash
--namespace=""                     # Limit resources queried for endpoints to a specific namespace (default: all namespaces)
--fqdn-template=""                 # A templated string that's used to generate DNS names from sources that don't define a hostname themselves
--[no-]combine-fqdn-annotation     # Combine FQDN template and Annotations instead of overwriting
--[no-]ignore-hostname-annotation  # Ignore hostname annotation when generating DNS names, valid only when --fqdn-template is set
```

The following flags are available for customizing Nomad integration (all optional):

```bash
--nomad-address=""         # Nomad API address. Defaults to $NOMAD_ADDR or http://127.0.0.1:4646
--nomad-region=""          # Nomad region. Defaults to agent's configured region
--nomad-token=NOMAD-TOKEN  # Nomad ACL token for authentication
--nomad-wait-time=0s       # API blocking timeout (Watch WaitTime)
```

Also, Nomad agent connection and authentication options can be configured via environment variables (e.g., `NOMAD_ADDR`, `NOMAD_REGION`, `NOMAD_TOKEN`, etc.)

### Example: Run ExternalDNS with Nomad source

You can run ExternalDNS with the Nomad service source and any supported DNS provider:

```bash
external-dns \
  --source=nomad-service \
  --provider=inmemory \
  --nomad-address=http://127.0.0.1:4646
```

## Nomad Service Tags

The `nomad-service` source uses Nomad [service tags](https://developer.hashicorp.com/nomad/docs/job-specification/service#tags) to define configuration that is typically specified using Kubernetes annotations in other sources.

Unlike Kubernetes annotations, which are a key-value map, Nomad tags are represented as a flat array of strings. To work around this difference, the implementation expects tags to follow a `key=value` format under the `external-dns.` prefix.

For example:

```hcl
tags = [
  "external-dns.hostname=example.nomad.internal.",
  "external-dns.ttl=300",
  "external-dns.controller=dns-controller",
  "external-dns.set-identifier=my-id"
]
```

These tags are interpreted by ExternalDNS in the same way it would interpret Kubernetes annotations:

```yaml
annotations:
  external-dns.alpha.kubernetes.io/hostname: "example.nomad.internal."
  external-dns.alpha.kubernetes.io/ttl: "300"
  external-dns.alpha.kubernetes.io/controller: "dns-controller"
  external-dns.alpha.kubernetes.io/set-identifier: "my-id"
```

To reduce verbosity and improve clarity, the Nomad tag keys are shortened versions of their corresponding Kubernetes annotations. This source supports next ExternalDNS annotation equivalents set via service tags:

| Nomad Tag | Kubernetes Annotation Equivalent |
|-----------|----------------------------------|
| external-dns.hostname | external-dns.alpha.kubernetes.io/hostname |
| external-dns.target | external-dns.alpha.kubernetes.io/target |
| external-dns.ttl | external-dns.alpha.kubernetes.io/ttl |
| external-dns.controller | external-dns.alpha.kubernetes.io/controller |
| external-dns.set-identifier | external-dns.alpha.kubernetes.io/set-identifier |

Additionally, provider-specific configuration is also supported via tags. For example:

```hcl
tags = [
  "external-dns.hostname=app.example.org.",
  "external-dns.aws-weight=100",
  "external-dns.cloudflare-proxied=true"
]
```

### Example service block in a Nomad job

```hcl
service {
  name = "whoami-demo"
  port = "http"
  provider = "nomad"
  tags = [
    "external-dns.hostname=whoami.example.org.",
    "external-dns.target=${attr.unique.network.ip-address}",
  ]
}
```

This configuration will result in an A record for whoami.example.org. pointing to the IP of the service port.

## Example Nomad Job

Here's a complete job spec to demonstrate usage:

```hcl
job "whoami" {
  group "demo" {
    network {
      mode = "host"

      port "http" {
        static = 80
      }
    }

    service {
      name = "whoami-demo"
      port = "http"
      provider = "nomad"
      tags = [
        "external-dns.alpha.kubernetes.io/hostname=whoami.example.org.",
        "external-dns.alpha.kubernetes.io/ttl=60"
      ]
    }

    task "server" {
      driver = "docker"

      config {
        image = "traefik/whoami"
        ports = ["http"]
      }

      env {
        WHOAMI_PORT_NUMBER = "${NOMAD_PORT_http}"
      }
    }
  }
}
```

The ExternalDNS itself can be deployed to a Nomad cluster as a regular job:

```hcl
job "external-dns" {
  group "external-dns" {
    network {
      mode = "bridge"

      port "http" {
        to = 7979
      }
    }

    task "controller" {
      driver = "docker"

      config {
        image = "registry.k8s.io/external-dns/external-dns:latest"
        ports = ["http"]

        args = [
          "--source=nomad-service",
          "--nomad-address=http://${attr.unique.network.ip-address}:4646",
        ]
      }

      resources {
        cpu    = 50
        memory = 32
      }

      service {
        provider = "nomad"
        port = "http"

        check {
          type = "http"
          path = "/healthz"
          interval = "10s"
          timeout = "3s"
        }
      }
    }
  }
}
```

## Notes

* This source does not require Kubernetes; it works directly with the Nomad HTTP API.
* Integration with DNS providers (e.g., AWS Route 53, Cloudflare, etc.) is handled the same as with other ExternalDNS sources.
