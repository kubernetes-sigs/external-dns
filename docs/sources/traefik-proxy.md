# Traefik Proxy Source

- [Traefik Documentation](https://doc.traefik.io/traefik/)
- [Traefik Helm Chart](https://github.com/traefik/traefik-helm-chart)

This tutorial describes how to configure ExternalDNS to use the Traefik Proxy source.
It is meant to supplement the other provider-specific setup tutorials.

## Manifest (for clusters without RBAC enabled)

```yaml
[[% include 'traefik-proxy/without-rbac.yaml' %]]
```

## Manifest (for clusters with RBAC enabled)

```yaml
[[% include 'traefik-proxy/with-cluster-rbac.yaml' %]]
```

## Deploying a Traefik IngressRoute

Create an IngressRoute file called 'ingress-route-default' with the following contents:

```yaml
[[% include 'traefik-proxy/ingress-route-default.yaml' %]]
```

Note the annotation on the IngressRoute (`external-dns.alpha.kubernetes.io/target`); use the same hostname as the traefik DNS.

ExternalDNS uses this annotation to determine what services should be registered with DNS.

Create the IngressRoute:

```sh
kubectl create -f docs/snippets/traefik-proxy/ingress-route-default.yaml
```

Depending where you run your IngressRoute it can take a little while for ExternalDNS synchronize the DNS record.

## Support private and public routing

To create a more robust and manageable Kubernetes environment, leverage separate Ingress classes to finely control public and private routing's security, performance, and operational policies. Similar approach could work in multi-tenant environments.

For this we are going to need two instances of `traefik` (public and private) as well as two instances of `external-dns`.

The `traefik` configuration should contain (for more detailed configured validate with the vendor)

```yaml
[[% include 'traefik-proxy/traefik-public-private-config.yaml' %]]
```

Create a IngressRoutes files with the following contents:

```yaml
[[% include 'traefik-proxy/ingress-route-public-private.yaml' %]]
```

And the arguments for `external-dns` instances should looks like

```yaml
---
args:
  - --source=traefik-proxy
  - --annotation-filter="kubernetes.io/ingress.class=traefik-public"
---
args:
  - --source=traefik-proxy
  - --annotation-filter="kubernetes.io/ingress.class=traefik-private"
```

## Cleanup

Now that we have verified that ExternalDNS will automatically manage Traefik DNS records, we can delete the tutorial's example:

```sh
kubectl delete -f docs/snippets/traefik-proxy/ingress-route-default.yaml
kubectl delete -f externaldns.yaml
```

## Additional Flags

| Flag                    | Description                                             |
|-------------------------|---------------------------------------------------------|
| --traefik-enable-legacy | Enable listeners on Resources under traefik.containo.us |
| --traefik-disable-new   | Disable listeners on Resources under traefik.io         |

### Resource Listeners

Traefik has deprecated the legacy API group, _traefik.containo.us_, in favor of _traefik.io_. By default the `traefik-proxy` source listen for resources under traefik.io API groups.

If needed, you can enable legacy listener with `--traefik-enable-legacy` and also disable new listener with `--traefik-disable-new`.
