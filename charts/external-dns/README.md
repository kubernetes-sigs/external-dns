# external-dns

![Version: 1.14.3](https://img.shields.io/badge/Version-1.14.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.14.0](https://img.shields.io/badge/AppVersion-0.14.0-informational?style=flat-square)

ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.

**Homepage:** <https://github.com/kubernetes-sigs/external-dns/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| stevehipwell | <steve.hipwell@gmail.com> |  |

## Source Code

* <https://github.com/kubernetes-sigs/external-dns/>

## Installing the Chart

Before you can install the chart you will need to add the `external-dns` repo to [Helm](https://helm.sh/).

```shell
helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
```

After you've installed the repo you can install the chart.

```shell
helm upgrade --install external-dns external-dns/external-dns --version 1.14.3
```

## Providers

Configuring the _ExternalDNS_ provider should be done via the `provider.name` value with provider specific configuration being set via the `provider.<name>.<key>` values, where supported, and the `extraArgs` value. For legacy support `provider` can be set to the name of the provider with all additional configuration being set via the `extraArgs` value.
See [documentation](https://kubernetes-sigs.github.io/external-dns/#new-providers) for more info on available providers and tutorials.

### Providers with Specific Configuration Support

| Provider               | Supported  |
|------------------------|------------|
| `webhook`              | ✅         |

## Namespaced Scoped Installation

external-dns supports running on a namespaced only scope, too.
If `namespaced=true` is defined, the helm chart will setup `Roles` and `RoleBindings` instead `ClusterRoles` and `ClusterRoleBindings`.

### Limited Supported

Not all sources are supported in namespaced scope, since some sources depends on cluster-wide resources.
For example: Source `node` isn't supported, since `kind: Node` has scope `Cluster`.
Sources like `istio-virtualservice` only work, if all resources like `Gateway` and `VirtualService` are present in the same
namespaces as `external-dns`.

The annotation `external-dns.alpha.kubernetes.io/endpoints-type: NodeExternalIP` is not supported.

If `namespaced` is set to `true`, please ensure that `sources` my only contains supported sources (Default: `service,ingress`).

### Support Matrix

| Source                 | Supported  | Infos                  |
|------------------------|------------|------------------------|
| `ingress`              | ✅         |                        |
| `istio-gateway`        | ✅         |                        |
| `istio-virtualservice` | ✅         |                        |
| `crd`                  | ✅         |                        |
| `kong-tcpingress`      | ✅         |                        |
| `openshift-route`      | ✅         |                        |
| `skipper-routegroup`   | ✅         |                        |
| `gloo-proxy`           | ✅         |                        |
| `contour-httpproxy`    | ✅         |                        |
| `service`              | ⚠️️         | NodePort not supported |
| `node`                 | ❌         |                        |
| `pod`                  | ❌         |                        |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity settings for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). If an explicit label selector is not provided for pod affinity or pod anti-affinity one will be created from the pod selector labels. |
| automountServiceAccountToken | bool | `nil` | Set this to `false` to [opt out of API credential automounting](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the `Pod`. |
| commonLabels | object | `{}` | Labels to add to all chart resources. |
| deploymentAnnotations | object | `{}` | Annotations to add to the `Deployment`. |
| deploymentStrategy | object | `{"type":"Recreate"}` | [Deployment Strategy](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy). |
| dnsPolicy | string | `nil` | [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy) for the pod, if not set the default will be used. |
| domainFilters | list | `[]` |  |
| env | list | `[]` | [Environment variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for the `external-dns` container. |
| extraArgs | list | `[]` | Extra arguments to provide to _ExternalDNS_. |
| extraVolumeMounts | list | `[]` | Extra [volume mounts](https://kubernetes.io/docs/concepts/storage/volumes/) for the `external-dns` container. |
| extraVolumes | list | `[]` | Extra [volumes](https://kubernetes.io/docs/concepts/storage/volumes/) for the `Pod`. |
| fullnameOverride | string | `nil` | Override the full name of the chart. |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for the `external-dns` container. |
| image.repository | string | `"registry.k8s.io/external-dns/external-dns"` | Image repository for the `external-dns` container. |
| image.tag | string | `nil` | Image tag for the `external-dns` container, this will default to `.Chart.AppVersion` if not set. |
| imagePullSecrets | list | `[]` | Image pull secrets. |
| initContainers | list | `[]` | [Init containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) to add to the `Pod` definition. |
| interval | string | `"1m"` | Interval for DNS updates. |
| livenessProbe | object | See _values.yaml_ | [Liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| logFormat | string | `"text"` | Log format. |
| logLevel | string | `"info"` | Log level. |
| nameOverride | string | `nil` | Override the name of the chart. |
| namespaced | bool | `false` | if `true`, _ExternalDNS_ will run in a namespaced scope (`Role`` and `Rolebinding`` will be namespaced too). |
| nodeSelector | object | `{}` | Node labels to match for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). |
| podAnnotations | object | `{}` | Annotations to add to the `Pod`. |
| podLabels | object | `{}` | Labels to add to the `Pod`. |
| podSecurityContext | object | See _values.yaml_ | [Pod security context](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#podsecuritycontext-v1-core), this supports full customisation. |
| policy | string | `"upsert-only"` | How DNS records are synchronized between sources and providers; available values are `sync` & `upsert-only`. |
| priorityClassName | string | `nil` | Priority class name for the `Pod`. |
| provider.name | string | `"aws"` | _ExternalDNS_ provider name; for the available providers and how to configure them see [README](https://github.com/kubernetes-sigs/external-dns/blob/master/charts/external-dns/README.md#providers). |
| provider.webhook.args | list | `[]` | Extra arguments to provide for the `webhook` container. |
| provider.webhook.env | list | `[]` | [Environment variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for the `webhook` container. |
| provider.webhook.extraVolumeMounts | list | `[]` | Extra [volume mounts](https://kubernetes.io/docs/concepts/storage/volumes/) for the `webhook` container. |
| provider.webhook.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for the `webhook` container. |
| provider.webhook.image.repository | string | `nil` | Image repository for the `webhook` container. |
| provider.webhook.image.tag | string | `nil` | Image tag for the `webhook` container. |
| provider.webhook.livenessProbe | object | See _values.yaml_ | [Liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| provider.webhook.readinessProbe | object | See _values.yaml_ | [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `webhook` container. |
| provider.webhook.resources | object | `{}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for the `webhook` container. |
| provider.webhook.securityContext | object | See _values.yaml_ | [Pod security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) for the `webhook` container. |
| provider.webhook.serviceMonitor | object | See _values.yaml_ | Optional [Service Monitor](https://prometheus-operator.dev/docs/operator/design/#servicemonitor) configuration for the `webhook` container. |
| rbac.additionalPermissions | list | `[]` | Additional rules to add to the `ClusterRole`. |
| rbac.create | bool | `true` | If `true`, create a `ClusterRole` & `ClusterRoleBinding` with access to the Kubernetes API. |
| readinessProbe | object | See _values.yaml_ | [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| registry | string | `"txt"` | Specify the registry for storing ownership and labels. Valid values are `txt`, `aws-sd`, `dynamodb` & `noop`. |
| resources | object | `{}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for the `external-dns` container. |
| revisionHistoryLimit | int | `nil` | Specify the number of old `ReplicaSets` to retain to allow rollback of the `Deployment``. |
| secretConfiguration.data | object | `{}` | `Secret` data. |
| secretConfiguration.enabled | bool | `false` | If `true`, create a `Secret` to store sensitive provider configuration (**DEPRECATED**). |
| secretConfiguration.mountPath | string | `nil` | Mount path for the `Secret`, this can be templated. |
| secretConfiguration.subPath | string | `nil` | Sub-path for mounting the `Secret`, this can be templated. |
| securityContext | object | See _values.yaml_ | [Security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) for the `external-dns` container. |
| service.annotations | object | `{}` | Service annotations. |
| service.ipFamilies | list | `[]` | Service IP families. |
| service.ipFamilyPolicy | string | `nil` | Service IP family policy. |
| service.port | int | `7979` | Service HTTP port. |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account. |
| serviceAccount.automountServiceAccountToken | string | `nil` | Set this to `false` to [opt out of API credential automounting](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the `ServiceAccount`. |
| serviceAccount.create | bool | `true` | If `true`, create a new `ServiceAccount`. |
| serviceAccount.labels | object | `{}` | Labels to add to the service account. |
| serviceAccount.name | string | `nil` | If this is set and `serviceAccount.create` is `true` this will be used for the created `ServiceAccount` name, if set and `serviceAccount.create` is `false` then this will define an existing `ServiceAccount` to use. |
| serviceMonitor.additionalLabels | object | `{}` | Additional labels for the `ServiceMonitor`. |
| serviceMonitor.annotations | object | `{}` | Annotations to add to the `ServiceMonitor`. |
| serviceMonitor.bearerTokenFile | string | `nil` | Provide a bearer token file for the `ServiceMonitor`. |
| serviceMonitor.enabled | bool | `false` | If `true`, create a `ServiceMonitor` resource to support the _Prometheus Operator_. |
| serviceMonitor.interval | string | `nil` | If set override the _Prometheus_ default interval. |
| serviceMonitor.metricRelabelings | list | `[]` | [Metric relabel configs](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#metric_relabel_configs) to apply to samples before ingestion. |
| serviceMonitor.namespace | string | `nil` | If set create the `ServiceMonitor` in an alternate namespace. |
| serviceMonitor.relabelings | list | `[]` | [Relabel configs](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config) to apply to samples before ingestion. |
| serviceMonitor.scheme | string | `nil` | If set overrides the _Prometheus_ default scheme. |
| serviceMonitor.scrapeTimeout | string | `nil` | If set override the _Prometheus_ default scrape timeout. |
| serviceMonitor.targetLabels | list | `[]` | Provide target labels for the `ServiceMonitor`. |
| serviceMonitor.tlsConfig | object | `{}` | Configure the `ServiceMonitor` [TLS config](https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig). |
| shareProcessNamespace | bool | `false` | If `true`, the `Pod` will have [process namespace sharing](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/) enabled. |
| sources | list | `["service","ingress"]` | _Kubernetes_ resources to monitor for DNS entries. |
| terminationGracePeriodSeconds | int | `nil` | Termination grace period for the `Pod` in seconds. |
| tolerations | list | `[]` | Node taints which will be tolerated for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). |
| topologySpreadConstraints | list | `[]` | Topology spread constraints for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). If an explicit label selector is not provided one will be created from the pod selector labels. |
| triggerLoopOnEvent | bool | `false` | If `true`, triggers run loop on create/update/delete events in addition of regular interval. |
| txtOwnerId | string | `nil` | Specify an identifier for this instance of _ExternalDNS_ wWhen using a registry other than `noop`. |
| txtPrefix | string | `nil` | Specify a prefix for the domain names of TXT records created for the `txt` registry. Mutually exclusive with `txtSuffix`. |
| txtSuffix | string | `nil` | Specify a suffix for the domain names of TXT records created for the `txt` registry. Mutually exclusive with `txtPrefix`. |

----------------------------------------------

Autogenerated from chart metadata using [helm-docs](https://github.com/norwoodj/helm-docs/).
