# External DNS
[![Build Status](https://travis-ci.org/kubernetes-incubator/external-dns.svg?branch=master)](https://travis-ci.org/kubernetes-incubator/external-dns)
[![Coverage Status](https://coveralls.io/repos/github/kubernetes-incubator/external-dns/badge.svg?branch=master)](https://coveralls.io/github/kubernetes-incubator/external-dns?branch=master)
[![GitHub release](https://img.shields.io/github/release/kubernetes-incubator/external-dns.svg)](https://github.com/kubernetes-incubator/external-dns/releases)

ExternalDNS synchronizes exposed Services and Ingresses with cloud DNS providers.

# Motivation

Inspired by Kubernetes' cluster-internal [DNS server](https://github.com/kubernetes/dns) ExternalDNS intends to make Kubernetes resources discoverable via public DNS servers. Similarly to KubeDNS it retrieves a list of resources from the Kubernetes API, such as Services and Ingresses, to determine a desired list of DNS records. However, unlike KubeDNS it's not a DNS server itself but merely configures other DNS providers accordingly, e.g. AWS Route53 or Google CloudDNS.

In a broader sense, it allows you to control DNS records dynamically via Kubernetes resources in a DNS provider agnostic way.

# Getting started

The [First Run](docs/tutorials/first-run.md) section is the fastest way to try out ExternalDNS on a Google Container Engine cluster. The [tutorials](docs/tutorials) section contains more detailed examples of how to setup ExternalDNS in different environments, such as other cloud providers and alternative ingress controllers.

# Roadmap

ExternalDNS was built with extensibility in mind. Adding and experimenting with new DNS providers and sources of desired DNS records should be as easy as possible. In addition, it should also be possible to modify how ExternalDNS behaves, e.g. whether it should add but must never delete records.

Furthermore, we're working on an ownership system that allows ExternalDNS to keep track of the records it created and will allow it to never modify records it doesn't have control over.

Here's a rough outline on what is to come:

### v0.1

* Support for Google CloudDNS
* Support for Kubernetes Services

### v0.2

* Support for AWS Route53
* Support for Kubernetes Ingresses

### v0.3

* Support for AWS Route53 via ALIAS
* Support for multiple zones
* Ownership System

### v1.0

* Ability to replace Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* Ability to replace Zalando's [Mate](https://github.com/zalando-incubator/mate)
* Ability to replace Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

The [FAQ](docs/faq.md) contains additional information and addresses several questions about key concepts of ExternalDNS.

Also have a look at [the milestones](https://github.com/kubernetes-incubator/external-dns/milestones) to get an idea of where we currently stand.

### Yet to be defined

* Support for CoreDNS and Azure DNS
* Support for record weights
* Support for different behavioral policies
* Support for Services with `type=NodePort`
* Support for TPRs
* Support for more advanced DNS record configurations

# Contributing

We encourage you to get involved with ExternalDNS, as users as well as contributors. Read the [contributing guidelines](CONTRIBUTING.md) and have a look at [the contributing docs](docs/contributing/getting-started.md) to learn about building the project, the project structure and the purpose of each package.

Feel free to reach out to us on the [Kubernetes slack](http://slack.k8s.io) in the #sig-network channel.

# Heritage

ExternalDNS is an effort to unify the following similar projects in order to bring the Kubernetes community an easy and predictable way of managing DNS records across cloud providers based on their Kubernetes resources.

* Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* Zalando's [Mate](https://github.com/zalando-incubator/mate)
* Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

## Kubernetes Incubator

This is a [Kubernetes Incubator project](https://github.com/kubernetes/community/blob/master/incubator.md).
The project was established 2017-Feb-9 (initial announcement [here](https://groups.google.com/forum/#!searchin/kubernetes-dev/external$20dns%7Csort:relevance/kubernetes-dev/2wGQUB0fUuE/9OXz01i2BgAJ)).
The incubator team for the project is:

* Sponsor: sig-network
* Champion: Tim Hockin (@thockin)
* SIG: sig-network

For more information about sig-network such as meeting times and agenda, check out the [community site](https://github.com/kubernetes/community/tree/master/sig-network).
