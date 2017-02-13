# Proposal: Design of External DNS

## Background

**Note: This DOC is WIP**

[Project proposal](https://groups.google.com/forum/#!searchin/kubernetes-dev/external$20dns%7Csort:relevance/kubernetes-dev/2wGQUB0fUuE/9OXz01i2BgAJ)

[Initial discussion](https://docs.google.com/document/d/1ML_q3OppUtQKXan6Q42xIq2jelSoIivuXI8zExbc6ec/edit#heading=h.1pgkuagjhm4p)

This document describes the initial design proposal

External DNS is purposed to fill the existing gap of creating DNS records for Kubernetes resources. While there exist alternative solutions, this project is meant to be a standard way of managing DNS records for Kubernetes. The current project is a fusion of the following projects and driven by its maintainers:

1. [Kops DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
2. [Mate](https://github.com/zalando-incubator/mate)
3. [wearemolecule/route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

## Example use case:

User runs `kubectl create -f ingress.yaml`, this will create an ingress as normal.
Typically the user would then have to manually create a DNS record pointing the ingress endpoint
If the external-dns controller is running on the cluster, it could automatically configure the DNS records instead, by observing the host attribute in the ingress object.

## Goals

1. Support AWS Route53 and Google Cloud DNS providers
2. DNS for Kubernetes services(type=Loadbalancer) and Ingress
3. Create/update/remove records as according to Kubernetes resources state
4. It should address main requirements and support main features of the projects mentioned above

## Design

### Extensibility

New cloud providers should be easily pluggable. Initially only AWS/Google platforms are supported. However, in the future we are planning to incorporate CoreDNS and Azure DNS as possible DNS providers

### Configuration

DNS records will be automatically created in multiple situations:
1. Setting `spec.rules.host` on an ingress object.
2. Specifying two annotations (`external-dns.kubernetes.io/controller` and `external-dns.kubernetes.io/hostname`) on a `type=LoadBalancer` service object.

### Annotations

TODO:*This should probably be placed in a separate file*.

Record configuration should occur via resource annotations. Supported annotations: 

|   Annotations |   |
|---|---|
|Tag   |external-dns.kubernetes.io/controller   |
|Description   |  Tells a DNS controller to process this service. This is useful when running different DNS controllers at the same time (or different versions of the same controller). The v1 implementation of dns-controller would look for service annotations `dns-controller` and `dns-controller/v1` but not for `mate/v1` or `dns-controller/v2` |
|Default   | dns-controller  |
|Example|dns-controller/v1|
|---|---|
|Tag   |external-dns.kubernetes.io/hostname   |
|Description   |  Fully qualified name of the desired record. Only required for services(Loadbalancer)  |
|Default| none |
|Example|foo.example.org|

### Compatibility

External DNS should be compatible with annotations used by three above mentioned projects. The idea is that resources created and tagged with annotations for other projects should continue to be valid and now managed by External DNS. 

TODO:*Add complete list here*

**Mate**

|Annotations |  |
|---|---|
|Tag   |zalando.org/dnsname  |
|Description   |  Hostname to be registered |
|Default   | Empty(falls back to template based approach) |
|Example|foo.example.org|


### Ownership

External DNS should be *responsible* for the created records. Which means that the records should be tagged (TODO:*describe how this supposed to work?*) and only tagged records are viable for future deletion/update. It should not mess with pre-existing records created via other means

