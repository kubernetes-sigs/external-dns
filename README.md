# External DNS
[![Build Status](https://travis-ci.org/kubernetes-incubator/external-dns.svg?branch=master)](https://travis-ci.org/kubernetes-incubator/external-dns)
[![Coverage Status](https://coveralls.io/repos/github/kubernetes-incubator/external-dns/badge.svg?branch=master)](https://coveralls.io/github/kubernetes-incubator/external-dns?branch=master)

ExternalDNS synchronizes exposed Services and Ingresses with cloud DNS providers.

# Motivation

Inspired by Kubernetes' cluster-internal [DNS server](https://github.com/kubernetes/dns) ExternalDNS intends to make Kubernetes resources discoverable via public DNS servers. Similarly to KubeDNS it retrieves a list of resources from the Kubernetes API, such as Services and Ingresses, to determine a desired list of DNS records. However, unlike KubeDNS it's not a DNS server itself but merely configures other DNS providers accordingly, e.g. AWS Route53 or Google CloudDNS.

In a broader sense, it allows you to control DNS records dynamically via Kubernetes resources in a DNS provider agnostic way.

# Quickstart

ExternalDNS' current release is `v0.1.0-beta.0` which allows to keep a managed zone in Google's [CloudDNS](https://cloud.google.com/dns/docs/) service synchronized with Services of `type=LoadBalancer` in your cluster.

In this release ExternalDNS is limited to a single managed zone and takes full ownership of it. That means if you have any existing records in that zone they will be removed. We encourage you to try out ExternalDNS in its own zone first to see if that model works for you. However, ExternalDNS, by default, runs in dryRun mode and won't make any changes to your infrastructure. So, as long as you don't change that flag, you're safe.

Make sure you meet the following prerequisites:
* You have a local Go 1.7+ development environment.
* You have access to a Google project with the DNS API enabled.
* You have access to a Kubernetes cluster that supports exposing Services, e.g. GKE.
* You have a properly setup, **unused** and **empty** hosted zone in Google CloudDNS.

First, get ExternalDNS.

```console
$ go get -u github.com/kubernetes-incubator/external-dns
```

Run an application and expose it via a Kubernetes Service.

```console
$ kubectl run nginx --image=nginx --replicas=1 --port=80
$ kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
```

Annotate the Service with your desired external DNS name. Make sure to change `example.org` to your domain and that it includes the trailing dot.

```console
$ kubectl annotate service nginx "external-dns.alpha.kubernetes.io/hostname=nginx.example.org."
```

Run a single sync loop of ExternalDNS locally. Make sure to change the Google project to one you control and the zone identifier to an **unused** and **empty** hosted zone in that project's Google CloudDNS.

```console
$ external-dns --zone example-org --google-project example-project --once
```

This should output the DNS records it's going to modify to match the managed zone with the DNS records you desire.

Once you're satisfied with the result you can run ExternalDNS like you would run it in your cluster: as a control loop and not in dryRun mode.

```console
$ external-dns --zone example-org --google-project example-project --dry-run=false
```

Check that ExternalDNS created the desired DNS record for your service and that it points to its load balancer's IP. Then try to resolve it.

```console
$ digshort nginx.example.org.
104.155.60.49
```

Now you can experiment and watch how ExternalDNS makes sure that your DNS records are configured as desired. Here are a couple of things you can try out:
* Change the desired hostname by modifying the Service's annotation.
* Recreate the Service and see that the DNS record will be updated to point to the new load balancer IP.
* Add another Service to create more DNS records.
* Remove Services to clean up your managed zone.

The [tutorials](docs/tutorials/gke.md) section contains a more detailed example of how to setup ExternalDNS in your Google Container Engine cluster.

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

The [FAQ](docs/faq.md) contains additional information and goes into more detail. Also have a look at [the milestones](https://github.com/kubernetes-incubator/external-dns/milestones) to the the current state.

### Unclear topics

* Support for CoreDNS
* Support for record weights
* Support for different behavioral policies
* Support for Services with `type=NodePort`
* Support for TPRs
* Support for more advanced DNS record configurations

# Contributing

You can build ExternalDNS for your platform with `make build`. The binary will land at `build/external-dns`.

ExternalDNS's sources of DNS records live in package [source](source). They implement the `Source` interface that has a single method `Endpoints` which returns the represented source's objects converted to `Endpoints`. Endpoints are just a tuple of DNS name and target where target can be an IP or another hostname.

For example, the `ServiceSource` returns all Services converted to `Endpoints` where the hostname is the value of the `external-dns.alpha.kubernetes.io/hostname` annotation and the target is the IP of the load balancer.

This list of endpoints is passed to the [Plan](plan) which determines the difference between the current DNS records and the desired list of `Endpoints`.

Once the difference has been figured out the list of intended changes is passed to a `Provider` which live in the [provider](provider) package. The provider is the adapter to the DNS provider, e.g. Google CloudDNS. It implements two methods: `ApplyChanges` to apply a set of changes and `Records` to retrieve the current list of records from the DNS provider.

The orchestration between the different components is controlled by the [controller](controller).

You can pick which `Source` and `Provider` to use at runtime via the `--source` and `--provider` flags, respectively.

A typical way to start on, e.g. a CoreDNS provider, would be to add a `coredns.go` to the providers package and implement the interface methods. Then you would have to register your provider under a name in `main.go`, e.g. `coredns`, and would be able to trigger it's functions via setting `--provider=coredns`.

Note, how your provider doesn't need to know anything about where the DNS records come from, nor does it have to figure out the difference between the current and the desired state, it merely executes the actions calculated by the plan.

We encourage you to get involved with ExternalDNS, as users as well as contributors. Reach out to us on [Kubernetes slack](https://github.com/kubernetes/community#slack-chat) in the #sig-network channel.

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
