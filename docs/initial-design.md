# Proposal: Design of External DNS

## Background

[Project proposal](https://groups.google.com/forum/#!searchin/kubernetes-dev/external$20dns%7Csort:relevance/kubernetes-dev/2wGQUB0fUuE/9OXz01i2BgAJ)

[Initial discussion](https://docs.google.com/document/d/1ML_q3OppUtQKXan6Q42xIq2jelSoIivuXI8zExbc6ec/edit#heading=h.1pgkuagjhm4p)

This document describes the initial design proposal.

External DNS is purposed to fill the existing gap of creating DNS records for Kubernetes resources. While there exist alternative solutions, this project is meant to be a standard way of managing DNS records for Kubernetes. The current project is a fusion of the following projects and driven by its maintainers:

1. [Kops DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
2. [Mate](https://github.com/linki/mate)
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
2. Setting `spec.tls.hosts` on an ingress object.
3. Adding the annotation `external-dns.alpha.kubernetes.io/hostname` on an ingress object.
4. Adding the annotation `external-dns.alpha.kubernetes.io/hostname` on a `type=LoadBalancer` service object.

### Annotations

Record configuration should occur via resource annotations. Supported annotations:

|   Annotations |   |
|---|---|
|Tag   |external-dns.alpha.kubernetes.io/controller   |
|Description   |  Tells a DNS controller to process this service. This is useful when running different DNS controllers at the same time (or different versions of the same controller). The v1 implementation of dns-controller would look for service annotations `dns-controller` and `dns-controller/v1` but not for `mate/v1` or `dns-controller/v2` |
|Default   | dns-controller  |
|Example|dns-controller/v1|
|Required| false |
|---|---|
|Tag   |external-dns.alpha.kubernetes.io/hostname   |
|Description   |  Fully qualified name of the desired record |
|Default| none |
|Example|foo.example.org|
|Required| Only for services. Ingress hostname is retrieved from `spec.rules.host` meta data on ingress |

### Compatibility

External DNS should be compatible with annotations used by three above mentioned projects. The idea is that resources created and tagged with annotations for other projects should continue to be valid and now managed by External DNS. 

**Mate**

Mate does not require services/ingress to be tagged. Therefore, it is not safe to run both Mate and External-DNS simultaneously. The idea is that initial release (?) of External DNS will support Mate annotations, which indicates the hostname to be created. Therefore the switch should be simple. 

|Annotations |  |
|---|---|
|Tag   |zalando.org/dnsname  |
|Description   |  Hostname to be registered |
|Default   | Empty(falls back to template based approach) |
|Example|foo.example.org|
|Required| false|

**route53-kubernetes**

It should be safe to run both `route53-kubernetes` and `external-dns` simultaneously. Since `route53-kubernetes` only looks at services with the label `dns=route53` and does not support ingress there should be no collisions between annotations. If users desire to switch to `external-dns` they can run both controllers and migrate services over as they are able.


### Ownership

External DNS should be *responsible* for the created records. Which means that the records should be tagged and only tagged records are viable for future deletion/update. It should not mess with pre-existing records created via other means.

#### Ownership via TXT records

Each record managed by External DNS is accompanied with a TXT record with a specific value to indicate that corresponding DNS record is managed by External DNS and it can be updated/deleted respectively. TXT records are limited to lifetimes of service/ingress objects and are created/deleted once k8s resources are created/deleted. 
