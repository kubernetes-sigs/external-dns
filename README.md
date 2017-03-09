# External DNS
[![Build Status](https://travis-ci.org/kubernetes-incubator/external-dns.svg?branch=master)](https://travis-ci.org/kubernetes-incubator/external-dns)
[![Coverage Status](https://coveralls.io/repos/github/kubernetes-incubator/external-dns/badge.svg?branch=master)](https://coveralls.io/github/kubernetes-incubator/external-dns?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubernetes-incubator/external-dns)](https://goreportcard.com/report/github.com/kubernetes-incubator/external-dns)

`external-dns` synchronizes external DNS servers (e.g. Google CloudDNS) with exposed Kubernetes Services and Ingresses.

This is a new Kubernetes Incubator project and will incorporate features from the following existing projects:

* [Kops DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* [Mate](https://github.com/zalando-incubator/mate)
* [wearemolecule/route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

## Current Status

The project just started and isn't in a usable state as of now. The current roadmap looks like the following.

* Create an MVP that allows managing DNS names for Service resources via Google CloudDNS with the official annotation set. (Done)
* Add support for Ingress and Node resources as well as AWS Route53.
* Add support for the annotation semantics of the three parent projects so that `external-dns` becomes a drop-in replacement for them.
* Switch from regular sync-only to watch and other advanced topics.

Please have a look at [the milestones](https://github.com/kubernetes-incubator/external-dns/milestones) to find corresponding issues.

## Features

* External DNS should be able to create/update/delete records on multiple cloud providers
* The used cloud provider should be configurable at runtime
* External DNS should take the ownership of the records created by it
* It should support Kubernetes Services with `type=Loadbalancer` and `type=NodePort`, Ingresses and Nodes
* Allow to customize external name via annotations
* It should be fault tolerance to individual pod failures
* Support weighted records annotations - allow different resources share same hostname, and respective weighted records should be created.
* Support multiple hosted zones - therefore External DNS should be able to create records as long as there exist a hosted zone matching the desired hostname

## Nice to have

* Should do smart cloud provider updates, i.e. Cloud Provider API should be called only when necessary
* High Availability - should be possible to run multiple instances of External DNS
* Should be able to monitor resource changes via K8S API for quick updates
* New DNS record sources (e.g. TPRs) and targets (e.g. Azure DNS) should be pluggable and easy to add

## Example

The following gives a rough view of a barely working `external-dns`. There are a couple of requirements to make the following work for you.
* Have a local Go 1.7+ development environment.
* Have access to a Google project with the DNS API enabled.
* Have access to a Kubernetes cluster that supports exposing Services, e.g. GKE.
* Have a properly setup, **unused** and **empty** hosted zone in Google.

Build the binary.

```console
$ mkdir -p $GOPATH/src/github.com/kubernetes-incubator/external-dns
$ cd $GOPATH/src/github.com/kubernetes-incubator/external-dns
$ git clone https://github.com/kubernetes-incubator/external-dns.git .
$ go build -o build/external-dns .
```

Run an application and expose it via a Kubernetes Service.

```console
$ kubectl run nginx --image=nginx --replicas=1 --port=80
$ kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
```

Annotate the service with your desired external DNS name (change `example.org` to your domain).

```console
$ kubectl annotate service nginx "external-dns.alpha.kubernetes.io/hostname=nginx.example.org."
```

Run a single sync loop of `external-dns`. In a real setup this would run constantly in your cluster. Change the Google project and zone identifier to an **unused** and **empty** hosted zone in Google. `external-dns` keeps the entire zone in sync with the desired records, which means that it will remove any records it doesn't know about. However, this will change in the future so that it tolerates and doesn't mess with existing records.

```console
$ build/external-dns --google-project example-project --google-zone example-org --once --dry-run=false
```

Check your cloud provider and see that the DNS record was created with the value of your load balancer IP.
Give DNS some time to propagate, then check that it resolves to your service IP.

```console
$ dig +short nginx.example.org.
1.2.3.4
```

Remove the annotation, delete or re-create the service, run `external-dns` again and watch it synchronize the DNS record accordingly.

When you're done testing remove the DNS record from the hosted zone via the UI and delete the example deployment.

```console
$ kubectl delete service nginx
$ kubectl delete deployment nginx
```

## Getting involved!

Want to contribute to External DNS? We would love the extra help from the community.

Reach out to us on [Kubernetes slack](https://github.com/kubernetes/community#slack-chat).

## Kubernetes Incubator

This is a [Kubernetes Incubator project](https://github.com/kubernetes/community/blob/master/incubator.md).
The project was established 2017-Feb-9 (initial announcement [here](https://groups.google.com/forum/#!searchin/kubernetes-dev/external$20dns%7Csort:relevance/kubernetes-dev/2wGQUB0fUuE/9OXz01i2BgAJ)).
The incubator team for the project is:

* Sponsor: sig-network
* Champion: Tim Hockin (@thockin)
* SIG: sig-network


For more information about sig-network such as meeting times and agenda, check out the community site.
