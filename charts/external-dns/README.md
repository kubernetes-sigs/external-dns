# external-dns

![Version: 1.13.1](https://img.shields.io/badge/Version-1.13.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.13.6](https://img.shields.io/badge/AppVersion-0.13.6-informational?style=flat-square)

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
helm upgrade --install external-dns external-dns/external-dns --version 1.13.1
```

## Providers

Configuring the _ExternalDNS_ provider should be done via the `provider.name` value with provider specific configuration being set via the
`provider.<name>.<key>` values, where supported, and the `extraArgs` value. For legacy support `provider` can be set to the name of the
provider with all additional configuration being set via the `extraArgs` value.

### Providers with Specific Configuration Support
| Parameter                                     | Description                                                                                                                                                                                                                                                                                                           | Default                                     |
|-----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| `image.repository`                            | Image repository.                                                                                                                                                                                                                                                                                                     | `registry.k8s.io/external-dns/external-dns` |
| `image.tag`                                   | Image tag, will override the default tag derived from the chart app version.                                                                                                                                                                                                                                          | `""`                                        |
| `image.pullPolicy`                            | Image pull policy.                                                                                                                                                                                                                                                                                                    | `IfNotPresent`                              |
| `imagePullSecrets`                            | Image pull secrets.                                                                                                                                                                                                                                                                                                   | `[]`                                        |
| `nameOverride`                                | Override the `name` of the chart.                                                                                                                                                                                                                                                                                     | `""`                                        |
| `fullnameOverride`                            | Override the `fullname` of the chart.                                                                                                                                                                                                                                                                                 | `""`                                        |
| `serviceAccount.create`                       | If `true`, create a new `serviceaccount`.                                                                                                                                                                                                                                                                             | `true`                                      |
| `serviceAccount.annotations`                  | Annotations to add to the service account.                                                                                                                                                                                                                                                                            | `{}`                                        |
| `serviceAccount.labels`                       | Labels to add to the service account.                                                                                                                                                                                                                                                                                 | `{}`                                        |
| `serviceAccount.name`                         | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the full name template.                                                                                                                                                                                       | `""`                                        |
| `serviceAccount.automountServiceAccountToken` | Opt out of the [service account token automounting feature](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the service account                                                                                                       | `null`                                      |
| `rbac.create`                                 | If `true`, create the RBAC resources.                                                                                                                                                                                                                                                                                 | `true`                                      |
| `rbac.additionalPermissions`                  | Additional permissions to be added to the cluster role.                                                                                                                                                                                                                                                               | `{}`                                        |
| `initContainers`                              | Add init containers to the pod.                                                                                                                                                                                                                                                                                       | `[]`                                        |
| `deploymentAnnotations`                       | Annotations to add to the Deployment.                                                                                                                                                                                                                                                                                 | `{}`                                        |
| `podLabels`                                   | Labels to add to the pod.                                                                                                                                                                                                                                                                                             | `{}`                                        |
| `podAnnotations`                              | Annotations to add to the pod.                                                                                                                                                                                                                                                                                        | `{}`                                        |
| `podSecurityContext`                          | Security context for the pod, this supports the full [PodSecurityContext](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#podsecuritycontext-v1-core) API.                                                                                                                                       | _see values.yaml_                           |
| `shareProcessNamespace`                       | If `true` enable [Process Namespace Sharing](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/)                                                                                                                                                                                       | `false`                                     |
| `securityContext`                             | Security context for the _external-dns_ container, this supports the full [SecurityContext](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#securitycontext-v1-core) API.                                                                                                                        | _see values.yaml_                           |
| `priorityClassName`                           | Priority class name to use for the pod.                                                                                                                                                                                                                                                                               | `""`                                        |
| `terminationGracePeriodSeconds`               | Termination grace period for the pod.                                                                                                                                                                                                                                                                                 | `null`                                      |
| `serviceMonitor.enabled`                      | If `true`, create a _Prometheus_ service monitor.                                                                                                                                                                                                                                                                     | `false`                                     |
| `serviceMonitor.namespace`                    | Forced namespace for ServiceMonitor.                                                                                                                                                                                                                                                                                  | `null`                                      |
| `serviceMonitor.annotations`                  | Annotations to be set on the ServiceMonitor.                                                                                                                                                                                                                                                                          | `{}`                                        |
| `serviceMonitor.additionalLabels`             | Additional labels to be set on the ServiceMonitor.                                                                                                                                                                                                                                                                    | `{}`                                        |
| `serviceMonitor.interval`                     | _Prometheus_ scrape frequency.                                                                                                                                                                                                                                                                                        | `null`                                      |
| `serviceMonitor.scrapeTimeout`                | _Prometheus_ scrape timeout.                                                                                                                                                                                                                                                                                          | `null`                                      |
| `serviceMonitor.scheme`                       | _Prometheus_ scrape scheme.                                                                                                                                                                                                                                                                                           | `null`                                      |
| `serviceMonitor.tlsConfig`                    | _Prometheus_ scrape tlsConfig.                                                                                                                                                                                                                                                                                        | `{}`                                        |
| `serviceMonitor.metricRelabelings`            | _Prometheus_ scrape metricRelabelings.                                                                                                                                                                                                                                                                                | `[]`                                        |
| `serviceMonitor.relabelings`                  | _Prometheus_ scrape relabelings.                                                                                                                                                                                                                                                                                      | `[]`                                        |
| `serviceMonitor.targetLabels`                 | _Prometheus_ scrape targetLabels.                                                                                                                                                                                                                                                                                     | `[]`                                        |
| `env`                                         | [Environment variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for the _external-dns_ container, this supports the full [EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#envvar-v1-core) API including secrets and configmaps. | `[]`                                        |
| `livenessProbe`                               | [Liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) for the _external-dns_ container, this supports the full [Probe](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#probe-v1-core) API.                                     | See _values.yaml_                           |
| `readinessProbe`                              | [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) for the _external-dns_ container, this supports the full [Probe](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#probe-v1-core) API.                                    | See _values.yaml_                           |
| `service.annotations`                         | Annotations to add to the service.                                                                                                                                                                                                                                                                                    | `{}`                                        |
| `service.port`                                | Port to expose via the service.                                                                                                                                                                                                                                                                                       | `7979`                                      |
| `extraVolumes`                                | Additional volumes for the pod, this supports the full [VolumeDevice](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#volumedevice-v1-core) API.                                                                                                                                                 | `[]`                                        |
| `extraVolumeMounts`                           | Additional volume mounts for the _external-dns_ container, this supports the full [VolumeMount](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#volumemount-v1-core) API.                                                                                                                        | `[]`                                        |
| `resources`                                   | Resource requests and limits for the _external-dns_ container, this supports the full [ResourceRequirements](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#resourcerequirements-v1-core) API.                                                                                                  | `{}`                                        |
| `nodeSelector`                                | Node labels for pod assignment.                                                                                                                                                                                                                                                                                       | `{}`                                        |
| `tolerations`                                 | Tolerations for pod assignment, this supports the full [Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#toleration-v1-core) API.                                                                                                                                                     | `[]`                                        |
| `affinity`                                    | Affinity settings for pod assignment, this supports the full [Affinity](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#affinity-v1-core) API.                                                                                                                                                   | `{}`                                        |
| `topologySpreadConstraints`                   | TopologySpreadConstraint settings for pod assignment, this supports the full [TopologySpreadConstraints](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#topologyspreadconstraint-v1-core) API.                                                                                                  | `[]`                                        |
| `logLevel`                                    | Verbosity of the logs, available values are: `panic`, `debug`, `info`, `warning`, `error`, `fatal`.                                                                                                                                                                                                                   | `info`                                      |
| `logFormat`                                   | Formats of the logs, available values are: `text`, `json`.                                                                                                                                                                                                                                                            | `text`                                      |
| `interval`                                    | The interval for DNS updates.                                                                                                                                                                                                                                                                                         | `1m`                                        |
| `triggerLoopOnEvent`                          | When enabled, triggers run loop on create/update/delete events in addition of regular interval.                                                                                                                                                                                                                       | `false`                                     |
| `namespaced`                                  | When enabled, external-dns runs on namespace scope. Additionally, Role and Rolebinding will be namespaced, too.                                                                                                                                                                                                       | `false`                                     |
| `sources`                                     | K8s resources type to be observed for new DNS entries.                                                                                                                                                                                                                                                                | See _values.yaml_                           |
| `policy`                                      | How DNS records are synchronized between sources and providers, available values are: `sync`, `upsert-only`.                                                                                                                                                                                                          | `upsert-only`                               |
| `registry`                                    | Registry Type, available types are: `txt`, `noop`.                                                                                                                                                                                                                                                                    | `txt`                                       |
| `txtOwnerId`                                  | TXT registry identifier.                                                                                                                                                                                                                                                                                              | `""`                                        |
| `txtPrefix`                                   | Prefix to create a TXT record with a name following the pattern `prefix.<CNAME record>`.                                                                                                                                                                                                                              | `""`                                        |
| `domainFilters`                               | Limit possible target zones by domain suffixes.                                                                                                                                                                                                                                                                       | `[]`                                        |
| `provider.name`                               | DNS provider where the DNS records will be created. [In-tree](https://github.com/kubernetes-sigs/external-dns#deploying-to-a-cluster) or [webhook](https://github.com/kubernetes-sigs/external-dns#new-providers).                                                                                                    | `aws`                                       |
| `provider.version`                            | On webhook provider: tagged version to use                                  | `latest`                                    |
| `provider.env`                                | On webhook provider: env variable on the sidecar                            | `[]`                                        |
| `provider.args`                               | On webhook provider: args on the sidecar                                    | `[]`                                        |
| `provider.resources`                          | On webhook provider: resources limits and requests                          | `{}`                                        |
| `provider.securityContext`                    | On webhook provider: securityContext to use                                 | `{}`                                        |

| `extraArgs`                                   | Extra arguments to pass to the _external-dns_ container, these are needed for provider specific arguments (these can be templated).                                                                                                                                                                                   | `[]`                                        |
| `deploymentStrategy`                          | .spec.strategy of the external-dns Deployment. Defaults to 'Recreate' since multiple external-dns pods may conflict with each other.                                                                                                                                                                                  | `{type: Recreate}`                          |
| `secretConfiguration.enabled`                 | Enable additional secret configuration.                                                                                                                                                                                                                                                                               | `false`                                     |
| `secretConfiguration.mountPath`               | Mount path of secret configuration secret (this can be templated).                                                                                                                                                                                                                                                    | `""`                                        |
| `secretConfiguration.data`                    | Secret configuration secret data. Could be used to store DNS provider credentials.                                                                                                                                                                                                                                    | `{}`                                        |
| `secretConfiguration.subPath`                 | Sub-path of secret configuration secret (this can be templated).                                                                                                                                                                                                                                                      | `""`                                        |
| `automountServiceAccountToken`                | Opt out of the [service account token automounting feature](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the pod                                                                                                                   | `null`                                      |
| `revisionHistoryLimit`                        | Optional field that specifies the number of old ReplicaSets to retain to allow rollback with the Deployment.                                                                                                                                                                                                          | `null`                                      |

| Provider               | Supported  |
|------------------------|------------|
| `webhook`              | ❌         |

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
| provider.name | string | `"aws"` | _ExternalDNS_ provider name; for the available providers and how to configure them see the [README](https://github.com/kubernetes-sigs/external-dns#deploying-to-a-cluster). |
| rbac.additionalPermissions | list | `[]` | Additional rules to add to the `ClusterRole`. |
| rbac.create | bool | `true` | If `true`, create a `ClusterRole` & `ClusterRoleBinding` with access to the Kubernetes API. |
| readinessProbe | object | See _values.yaml_ | Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| registry | string | `"txt"` | Specify the registry for storing ownership and labels. Valid values are `txt`, `aws-sd`, `dynamodb` & `noop`. |
| resources | object | `{}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for the `external-dns` container. |
| revisionHistoryLimit | int | `nil` | Specify the number of old `ReplicaSets` to retain to allow rollback of the `Deployment``. |
| secretConfiguration.data | object | `{}` | `Secret` data. |
| secretConfiguration.enabled | bool | `false` | If `true`, create a `Secret` to store sensitive provider configuration. |
| secretConfiguration.mountPath | string | `nil` | Mount path for the `Secret`, this can be templated. |
| secretConfiguration.subPath | string | `nil` | Sub-path for mounting the `Secret`, this can be templated. |
| securityContext | object | See _values.yaml_ | [Security context](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#securitycontext-v1-core) for the `external-dns` container. |
| service.annotations | object | `{}` | Service annotations. |
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
