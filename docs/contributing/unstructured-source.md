# Unstructured Source

Unstructured source provides a generic mechanism to manage DNS records in your favourite DNS provider supported by external-dns by watching for a user specified Kubernetes API to extract hostnames, targets, TTL, and record type information from a combination of well-known annotations and evaluating JSONPath expressions against the resource.

* [Details](#details)
* [Usage](#usage)
* [RBAC Configuration](#rbac-configuration)
* [Demo](#demo)
  * [Setup](#setup-steps)
  * [Examples](#examples):
    * `ConfigMap`
      * [DNS A-record with multiple targets](#dns-a-record-with-multiple-targets-from-a-configmap)
      * [Multiple DNS A-records](#multiple-dns-a-records-from-a-configmap)
    * `Deployment`
      * [DNS CNAME-record from annotations](#dns-cname-record-from-annotation-on-a-deployment)
    * `Workload` (_namespace-scoped CRD_)
      * [DNS A-record](#dns-a-record-from-a-namespace-scoped-crd)
      * [Multiple DNS A-records](#multiple-dns-a-records-from-a-namespace-scoped-crd)
      * [DNS A-record with multiple targets](#dns-a-record-with-multiple-targets-from-a-namespace-scoped-crd)
    * `ClusterWorkload` (_cluster-scoped CRD_)
      * [DNS A-record](#dns-a-record-a-cluster-scoped-crd)
  * [Cleanup](#cleanup-steps)

## Details

The unstructured source is largely designed for API resources that have one hostname and one target (likely an IP address), but it is flexible enough to construct DNS entries if there are multiple hostnames and/or targets per API resource. The high-level summary is:

* one DNS entry is created per hostname from a resource
* each DNS entry contains all of the targets from a resource

For instance, let's say one API resource evaluates to having the following information:

* 2 hostnames
* 1 target

Two DNS entries will be created, one for each hostname, both pointing to the same target. Alternatively, there may be an API resource that evaluates to:

* 2 hostnames
* 3 targets

In this case, two DNS entries will be created, one of each hostname, and both entries will still point to the same, three targets.

## Usage

One can use Unstructured source by specifying `--source` flag with `unstructured` and specifying:

| Flag | Description |
|------|-------------|
| `--unstructured-source-apiversion` | The API Version of the API resource. |
| `--unstructured-source-kind` | The Kind of the API resource. |
| `--unstructured-source-target-json-path` | A JSONPath expression executed against the API resource that must evaluate to a comma or space delimited string with one or more targets. This path must conform to the [JSONPath format](https://kubernetes.io/docs/reference/kubectl/jsonpath/). If this flag is omitted then the hostname value is derived from the annotation, `external-dns.alpha.kubernetes.io/target` |
| `--unstructured-source-hostname-json-path` | An optional, JSONPath expression executed against the API resource that must evaluate to a comma or space delimited string with one or more hostnames. This path must conform to the [JSONPath format](https://kubernetes.io/docs/reference/kubectl/jsonpath/). If this flag is omitted then the hostname value is derived from the annotation, `external-dns.alpha.kubernetes.io/hostname` |

The TTL and record type are always derived from annotations:

| Annotation | Default |
|------------|---------|
| `external-dns.alpha.kubernetes.io/ttl` | `0` |
| `external-dns.alpha.kubernetes.io/record-type` | `A` |

For example:

```
$ build/external-dns \
  --source unstructured \
  --unstructured-source-apiversion v1 \
  --unstructured-source-kind Pod \
  --unstructured-source-target-json-path '{.status.podIP}' \
  --provider inmemory \
  --once \
  --dry-run
```

## RBAC Configuration

If the Kubernetes cluster uses RBAC, the `external-dns` ClusterRole requires access to `get`, `watch`, and `list` the API resource configured with the Unstructured source. For example, if the Unstructured source is configured for `Pod` resources then the following RBAC is required:

```
- apiGroups: ["v1"]
  resources: ["pods"]
  verbs: ["get","watch","list"]
```

## Demo

This section provides a nice little demo of the unstructured source with several examples.

### Setup Steps

1. Build the `external-dns` binary:

    ```shell
    make build
    ```

1. Use [Docker](https://www.docker.com) and [Kind](https://kind.sigs.k8s.io) to create a local, Kubernetes cluster:

    ```shell
    kind create cluster
    ```

1. Update the kubeconfig context to point to the Kind cluster so that the External DNS binary can access the cluster:

    ```shell
    kubectl config set-context kind-kind
    ```

1. Apply all of the CRDs required by the examples below:

    ```shell
    $ kubectl apply -f "docs/contributing/unstructured-source/*-manifest.yaml"
    customresourcedefinition.apiextensions.k8s.io/clusterworkloads.example.com created
    customresourcedefinition.apiextensions.k8s.io/workloads.example.com created
    ```

1. Apply the example resources:

    ```shell
    $ kubectl apply -f "docs/contributing/unstructured-source/*-example.yaml"
    clusterworkload.example.com/my-workload-1 created
    configmap/my-workload-1 created
    configmap/my-workload-2 created
    deployment.apps/my-workload-1 created
    workload.example.com/my-workload-1 created
    workload.example.com/my-workload-2 created
    workload.example.com/my-workload-3 created
    ```

1. Several of the examples require patching a `status` sub-resource, which is not supported by `kubectl`. The following commands persist the access information for the Kind cluster to files may be used by `curl` in order to patch these `status` sub-resources:

    1. Save the API endpoint:

        ```shell
        kubectl config view --raw \
          -o jsonpath='{.clusters[?(@.name == "kind-kind")].cluster.server}' \
          >url.txt
        ```

    1. Save the cluster's certification authority (CA):

        ```shell
        kubectl config view --raw \
          -o jsonpath='{.clusters[?(@.name == "kind-kind")].cluster.certificate-authority-data}' | \
          { base64 -d 2>/dev/null || base64 -D; } \
          >ca.crt
        ```

    1. Save the client's public certificate:

        ```shell
        kubectl config view --raw \
          -o jsonpath='{.users[?(@.name == "kind-kind")].user.client-certificate-data}' | \
          { base64 -d 2>/dev/null || base64 -D; } \
          >client.crt
        ```

    1. Save the client's private key:

        ```shell
        kubectl config view --raw \
          -o jsonpath='{.users[?(@.name == "kind-kind")].user.client-key-data}' | \
          { base64 -d 2>/dev/null || base64 -D; } \
          >client.key
        ```

After running all of the desired examples the command `kind delete cluster` may be used to clean up the local Kubernetes cluster.

### Examples

#### DNS A-Record with Multiple Targets from a ConfigMap

This example realizes a single DNS A-record from a `ConfigMap` resource:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-workload-1
  annotations:
    example: my-workload-1
    description: |
      This example results in the creation of a single endpoint with
      one three, A-record targets -- one IP4 addr and two IP6 addrs.
      The endpoint's DNS name and TTL are derived from the annotations
      below, and since no record type is specified, the default type,
      A-record, is used.
    external-dns.alpha.kubernetes.io/hostname: my-workload-1.example.com
    external-dns.alpha.kubernetes.io/ttl: 10m
data:
  ip4-addrs: 1.2.3.4
  ip6-addrs: "2001:db8:0:1:1:1:1:1,2001:db8:0:1:1:1:1:2"
```

Run External DNS:

```shell
build/external-dns \
  --annotation-filter="example=my-workload-1" \
  --source unstructured \
  --unstructured-source-apiversion v1 \
  --unstructured-source-kind ConfigMap \
  --unstructured-source-target-json-path '{.data.ip4-addrs} {.data.ip6-addrs}' \
  --provider inmemory \
  --once \
  --dry-run
```

The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

```shell
INFO[0000] Unstructured source configured for namespace-scoped resource with kind "ConfigMap" in apiVersion "v1" in namespace "" 
INFO[0000] resource="my-workload-1", hostnames=[my-workload-1.example.com], targets=[1.2.3.4 2001:db8:0:1:1:1:1:1 2001:db8:0:1:1:1:1:2], ttl=600, recordType="A" 
INFO[0000] CREATE: my-workload-1.example.com 600 IN A  1.2.3.4;2001:db8:0:1:1:1:1:1;2001:db8:0:1:1:1:1:2 [] 
INFO[0000] CREATE: my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
INFO[0000] CREATE: a-my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
```

#### Multiple DNS A-Records from a ConfigMap

This example realizes two DNS A-records from a `ConfigMap` resource:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-workload-2
  annotations:
    example: my-workload-2
    description: |
      This example results in the creation of two endpoints with
      one, A-record targets, an IP4 address. An endpoint is created
      for each of the DNS names specified in the annotation below.
      Since no TTL or record type are specified, the default values
      for each are used, a TTL of 0 and an A-record.
    external-dns.alpha.kubernetes.io/hostname: my-workload-1.example.com, my-workload-2.example.com
data:
  ip4-addrs: 1.2.3.4
```

Run External DNS:

```shell
build/external-dns \
  --annotation-filter="example=my-workload-2" \
  --source unstructured \
  --unstructured-source-apiversion v1 \
  --unstructured-source-kind ConfigMap \
  --unstructured-source-target-json-path '{.data.ip4-addrs}' \
  --provider inmemory \
  --once \
  --dry-run
```

The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

```shell
INFO[0000] Unstructured source configured for namespace-scoped resource with kind "ConfigMap" in apiVersion "v1" in namespace "" 
INFO[0000] resource="my-workload-2", hostnames=[my-workload-1.example.com my-workload-2.example.com], targets=[1.2.3.4], ttl=0, recordType="A" 
INFO[0000] CREATE: my-workload-1.example.com 0 IN A  1.2.3.4 [] 
INFO[0000] CREATE: my-workload-2.example.com 0 IN A  1.2.3.4 [] 
INFO[0000] CREATE: my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
INFO[0000] CREATE: a-my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
INFO[0000] CREATE: my-workload-2.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
INFO[0000] CREATE: a-my-workload-2.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
```

### DNS CNAME-Record from Annotation on a Deployment

This example realizes two DNS CNAME-records from a `Deployment` resource:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-workload-1
  labels:
    app: my-workload-1
  annotations:
    example: my-workload-1
    description: |
      This example results in the creation of two endpoints with
      one, CNAME-record targets, a hostname. An endpoint is created
      for each of the DNS names specified in the annotation below.
      Since no TTL is specified, the default value, a TTL of 0, is
      used.
    external-dns.alpha.kubernetes.io/hostname: my-workload-1a.example.com my-workload-1b.example.com
    external-dns.alpha.kubernetes.io/record-type: CNAME
    external-dns.alpha.kubernetes.io/target: my-workload-1.example.com
spec:
  replicas: 0
  selector:
    matchLabels:
      app: my-workload-1
  template:
    metadata:
      labels:
        app: my-workload-1
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
```

Run External DNS:

```shell
build/external-dns \
  --annotation-filter="example=my-workload-1" \
  --source unstructured \
  --unstructured-source-apiversion "apps/v1" \
  --unstructured-source-kind Deployment \
  --provider inmemory \
  --once \
  --dry-run
```

The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

```shell
INFO[0000] Unstructured source configured for namespace-scoped resource with kind "Deployment" in apiVersion "apps/v1" in namespace "" 
INFO[0000] resource="my-workload-1", hostnames=[my-workload-1a.example.com my-workload-1b.example.com], targets=[my-workload-1.example.com], ttl=0, recordType="CNAME" 
INFO[0000] CREATE: my-workload-1a.example.com 0 IN CNAME  my-workload-1.example.com [] 
INFO[0000] CREATE: my-workload-1b.example.com 0 IN CNAME  my-workload-1.example.com [] 
INFO[0000] CREATE: my-workload-1a.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
INFO[0000] CREATE: cname-my-workload-1a.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
INFO[0000] CREATE: my-workload-1b.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
INFO[0000] CREATE: cname-my-workload-1b.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
```

#### DNS A-Record from a Namespace-Scoped CRD

This example realizes a single DNS A-record from a namespace-scoped, CRD resource:

```yaml
apiVersion: example.com/v1alpha1
kind: Workload
metadata:
  name: my-workload-1
  annotations:
    example: my-workload-1
    description: |
      This example results in the creation of a single endpoint with
      one, A-record target, an IP4 address. The endpoint's DNS name
      and TTL are derived from the annotations below, and since no
      record type is specified, the default type, A-record, is used.
    external-dns.alpha.kubernetes.io/hostname: my-workload-1.example.com
    external-dns.alpha.kubernetes.io/ttl: 10m
spec: {}
# kubectl is disallowed from editing the status sub-resource. it is included
# here for completeness, but the example actually used curl to set the status
status:
  addr: 1.2.3.4
```

1. Patch the `status` sub-resource with `curl`:

    ```shell
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XGET -H 'Content-Type: application/json' -H 'Accept: application/json' \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-1" | \
    jq '.status.addr="1.2.3.4"' | \
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XPUT -H 'Content-Type: application/json' -H 'Accept: application/json' -d @- \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-1/status"
    ```

1. Run External DNS:

    ```shell
    build/external-dns \
      --annotation-filter="example=my-workload-1" \
      --source unstructured \
      --unstructured-source-apiversion "example.com/v1alpha1" \
      --unstructured-source-kind Workload \
      --unstructured-source-target-json-path '{.status.addr}' \
      --provider inmemory \
      --once \
      --dry-run
    ```

    The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

    ```shell
    INFO[0000] Unstructured source configured for namespace-scoped resource with kind "Workload" in apiVersion "example.com/v1alpha1" in namespace "" 
    INFO[0000] resource="my-workload-1", hostnames=[my-workload-1.example.com], targets=[1.2.3.4], ttl=600, recordType="A" 
    INFO[0000] CREATE: my-workload-1.example.com 600 IN A  1.2.3.4 [] 
    INFO[0000] CREATE: my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
    INFO[0000] CREATE: a-my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-1" [] 
    ```

#### Multiple DNS A-Records from a Namespace-Scoped CRD

This example realizes three DNS A-record from a namespace-scoped, CRD resource:

```yaml
apiVersion: example.com/v1alpha1
kind: Workload
metadata:
  name: my-workload-2
  annotations:
    example: my-workload-2
    description: |
      This example results in the creation of three endpoints with
      one, A-record target, an IP4 address. An endpoint is created
      for each of the DNS names specified in the spec below.
      Since no TTL or record type are specified, the default values
      for each are used, a TTL of 0 and an A-record.
spec:
  hostname: my-workload-2.example.com
  additionalHostnames:
  - my-workload-2a.example.com
  - my-workload-2b.example.com
# kubectl is disallowed from editing the status sub-resource. it is included
# here for completeness, but the example actually used curl to set the status
status:
  addr: 1.2.3.4
```

1. Patch the `status` sub-resource with `curl`:

    ```shell
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XGET -H 'Content-Type: application/json' -H 'Accept: application/json' \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-2" | \
    jq '.status.addr="1.2.3.4"' | \
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XPUT -H 'Content-Type: application/json' -H 'Accept: application/json' -d @- \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-2/status"
    ```

1. Run External DNS:

    ```shell
    build/external-dns \
      --annotation-filter="example=my-workload-2" \
      --source unstructured \
      --unstructured-source-apiversion "example.com/v1alpha1" \
      --unstructured-source-kind Workload \
      --unstructured-source-target-json-path '{.status.addr}' \
      --unstructured-source-hostname-json-path '{.spec.hostname} {.spec.additionalHostnames[*]}' \
      --provider inmemory \
      --once \
      --dry-run
    ```

    The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

    ```shell
    INFO[0000] Unstructured source configured for namespace-scoped resource with kind "Workload" in apiVersion "example.com/v1alpha1" in namespace "" 
    INFO[0000] resource="my-workload-2", hostnames=[my-workload-2.example.com my-workload-2a.example.com my-workload-2b.example.com], targets=[1.2.3.4], ttl=0, recordType="A" 
    INFO[0000] CREATE: my-workload-2a.example.com 0 IN A  1.2.3.4 [] 
    INFO[0000] CREATE: my-workload-2b.example.com 0 IN A  1.2.3.4 [] 
    INFO[0000] CREATE: my-workload-2.example.com 0 IN A  1.2.3.4 [] 
    INFO[0000] CREATE: my-workload-2a.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    INFO[0000] CREATE: a-my-workload-2a.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    INFO[0000] CREATE: my-workload-2b.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    INFO[0000] CREATE: a-my-workload-2b.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    INFO[0000] CREATE: my-workload-2.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    INFO[0000] CREATE: a-my-workload-2.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-2" [] 
    ```

#### DNS A-Record with Multiple Targets from a Namespace-Scoped CRD

This example realizes a DNS A-record from a namespace-scoped, CRD resource:

```yaml
apiVersion: example.com/v1alpha1
kind: Workload
metadata:
  name: my-workload-3
  annotations:
    example: my-workload-3
    description: |
      This example results in the creation of a single endpoint with
      one four, A-record targets -- two IP4 addrs and two IP6 addrs.
      The endpoint's DNS is derived from the spec below. Since no
      TTL or record type are specified, the default values for each
      are used, a TTL of 0 and an A-record.
spec:
  hostname: my-workload-3.example.com
# kubectl is disallowed from editing the status sub-resource. it is included
# here for completeness, but the example actually used curl to set the status
status:
  addr: 1.2.3.4
  additionalAddrs:
  - 5.6.7.8
  - 2001:db8:0:1:1:1:1:1
  - 2001:db8:0:1:1:1:1:2
```

1. Patch the `status` sub-resource with `curl`:

    ```shell
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XGET -H 'Content-Type: application/json' -H 'Accept: application/json' \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-3" | \
    jq '.status.addr="1.2.3.4"' | jq '.status.additionalAddrs=["5.6.7.8","2001:db8:0:1:1:1:1:1","2001:db8:0:1:1:1:1:2"]' | \
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XPUT -H 'Content-Type: application/json' -H 'Accept: application/json' -d @- \
      "$(cat url.txt)/apis/example.com/v1alpha1/namespaces/default/workloads/my-workload-3/status"
    ```

1. Run External DNS:

    ```shell
    build/external-dns \
      --annotation-filter="example=my-workload-3" \
      --source unstructured \
      --unstructured-source-apiversion "example.com/v1alpha1" \
      --unstructured-source-kind Workload \
      --unstructured-source-target-json-path '{.status.addr} {.status.additionalAddrs[*]}' \
      --unstructured-source-hostname-json-path '{.spec.hostname}' \
      --provider inmemory \
      --once \
      --dry-run
    ```

    The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

    ```shell
    INFO[0000] Unstructured source configured for namespace-scoped resource with kind "Workload" in apiVersion "example.com/v1alpha1" in namespace "" 
    INFO[0000] resource="my-workload-3", hostnames=[my-workload-3.example.com], targets=[1.2.3.4 5.6.7.8 2001:db8:0:1:1:1:1:1 2001:db8:0:1:1:1:1:2], ttl=0, recordType="A" 
    INFO[0000] CREATE: my-workload-3.example.com 0 IN A  1.2.3.4;5.6.7.8;2001:db8:0:1:1:1:1:1;2001:db8:0:1:1:1:1:2 [] 
    INFO[0000] CREATE: my-workload-3.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-3" [] 
    INFO[0000] CREATE: a-my-workload-3.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/default/my-workload-3" [] 
    ```

### DNS A-Record a Cluster-Scoped CRD

This example realizes a single DNS A-record from a cluster-scoped, CRD resource:

```yaml
apiVersion: example.com/v1alpha1
kind: ClusterWorkload
metadata:
  name: my-workload-1
  annotations:
    example: my-workload-1
    description: |
      This example results in the creation of a single endpoint with
      one, A-record targets, an IP4 address. The endpoint's DNS name
      and TTL are derived from the spec and annotations below, and
      since no record type is specified, the default type, A-record,
      is used.

      This example highlights that the unstructured source may be used
      with cluster-scoped resources as well.
    external-dns.alpha.kubernetes.io/ttl: 10m
spec:
  hostname: my-workload-1.example.com
# kubectl is disallowed from editing the status sub-resource. it is included
# here for completeness, but the example actually used curl to set the status
status:
  addr: 1.2.3.4
```

1. Patch the `status` sub-resource with `curl`:

    ```shell
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XGET -H 'Content-Type: application/json' -H 'Accept: application/json' \
      "$(cat url.txt)/apis/example.com/v1alpha1/clusterworkloads/my-workload-1" | \
    jq '.status.addr="1.2.3.4"' | \
    curl --cacert ca.crt --cert client.crt --key client.key \
      --silent --show-error \
      -XPUT -H 'Content-Type: application/json' -H 'Accept: application/json' -d @- \
      "$(cat url.txt)/apis/example.com/v1alpha1/clusterworkloads/my-workload-1/status"
    ```

1. Run External DNS:

    ```shell
    build/external-dns \
      --annotation-filter="example=my-workload-1" \
      --source unstructured \
      --unstructured-source-apiversion "example.com/v1alpha1" \
      --unstructured-source-kind ClusterWorkload \
      --unstructured-source-target-json-path '{.status.addr}' \
      --unstructured-source-hostname-json-path '{.spec.hostname}' \
      --provider inmemory \
      --once \
      --dry-run
    ```

    The following lines at the end of the External DNS output illustrate the successful creation of the DNS record(s):

    ```shell
    INFO[0000] Unstructured source configured for cluster-scoped resource with kind "ClusterWorkload" in apiVersion "example.com/v1alpha1" 
    INFO[0000] resource="my-workload-1", hostnames=[my-workload-1.example.com], targets=[1.2.3.4], ttl=600, recordType="A" 
    INFO[0000] CREATE: my-workload-1.example.com 600 IN A  1.2.3.4 [] 
    INFO[0000] CREATE: my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/my-workload-1" [] 
    INFO[0000] CREATE: a-my-workload-1.example.com 0 IN TXT  "heritage=external-dns,external-dns/owner=default,external-dns/resource=unstructured/my-workload-1" []
    ```

### Cleanup Steps

1. Delete the Kind cluster:

    ```shell
    kind delete cluster
    ```

1. Cleanup the files created to use `curl` with the cluster:

    ```shell
    rm url.txt ca.crt client.crt client.key
    ```