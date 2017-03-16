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

The [tutorials](docs/tutorials/gke.md) section contains a detailed example of how to setup `external-dns` on Google Container Engine.

## Building

You need a working Go 1.7+ development environment. Then run `make build` to build `external-dns` for your platform. The binary will land at `build/external-dns`.

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
