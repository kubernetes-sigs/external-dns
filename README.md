# External DNS

Configure external DNS servers for Kubernetes clusters.

This is a new Kubernetes Incubator project and will incorporate features from the following existing projects:

* [Kops DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* [Mate](https://github.com/zalando-incubator/mate)
* [wearemolecule/route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

## Current Status

The project just started and isn't in a usable state as of now. The current roadmap looks like the following.

* Create an MVP that allows managing DNS names for Service resources via Google CloudDNS with the official annotation set.
* Add support for Ingress and Node resources as well as AWS Route53.
* Add support for the annotation semantics of the three parent projects so that `external-dns` becomes a drop-in replacement for them.
* Switch from regular sync-only to watch and other advanced topics.

Please have a look at [the milestones](https://github.com/kubernetes-incubator/external-dns/milestones) to find corresponding issues.

## Features

* External DNS should be able to create/update/delete records on Cloud Provider DNS server
* Configurable Cloud Provider
* External DNS should take the ownership of the records created by it
* Support Kubernetes Service(type=Loadbalancer), Ingress and NodePorts
* Support DNS naming via annotations
* Fault Tolerance
* Support weighted records annotations - allow different resources share same hostname, and respective weighted records should be created.
* Support multiple hosted zones - therefore External DNS should be able to create records as long as there exist a hosted zone matching the desired hostname

## Minor Features

* Should do smart cloud provider updates, i.e. Cloud Provider API should be called only when necessary
* High Availability - should be possible to run multiple instances of External DNS
* Should be able to monitor resource changes via K8S API for quick updates
* New DNS record sources (e.g. TPRs) and targets (e.g. Azure DNS) should be pluggable and easy to add

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
