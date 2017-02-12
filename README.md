# External DNS

Configure external DNS servers for Kubernetes clusters.

This is a new Kubernetes Incubator project and will incorporate features from the following existing projects:

* [Kops DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* [Mate](https://github.com/zalando-incubator/mate)
* [wearemolecule/route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)


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
