# Pod Source

The pod source creates DNS entries based on `Pod` resources.

## Pods not running with host networking

By default, the pod source will not consider the pods that aren't running with host networking enabled. You can override this behavior by using the `--ignore-non-host-network-pods` option.

## Using a default domain for pods

By default, the pod source will look into the pod annotations to find the FQDN associated with a pod. You can also use the option `--pod-source-domain=example.org` to build the FQDN of the pods. The pod named "test-pod" will then be registered as "test-pod.example.org".

## Configuration for registering all pods with their associated PTR record

A use case where combining these options can be pertinent is when you are running on-premise Kubernetes clusters without SNAT enabled for the pod network.
You might want to register all the pods in the DNS with their associated PTR record so that the source of some traffic outside of the cluster can be rapidly associated with a workload using the "nslookup" or "dig" command on the pod IP.
This can be particularly useful if you are running a large number of Kubernetes clusters.

You will then use the following mix of options:

- `--domain-filter=example.org`
- `--domain-filter=10.0.0.in-addr.arpa`
- `--source=pod`
- `--pod-source-domain=example.org`
- `--no-ignore-non-host-network-pods`
- `--rfc2136-create-ptr`
- `--rfc2136-zone=example.org`
- `--rfc2136-zone=10.0.0.in-addr.arpa`
