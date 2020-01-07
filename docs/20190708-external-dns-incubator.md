# Move ExternalDNS out of Kubernetes incubator

<!-- TOC depthFrom:1 depthTo:6 withLinks:1 updateOnSave:1 orderedList:0 -->

- [Move ExternalDNS out of Kubernetes incubator](#move-externaldns-out-of-kubernetes-sigs)
	- [Summary](#summary)
	- [Motivation](#motivation)
		- [Goals](#goals)
	- [Proposal](#proposal)
	- [Details](#details)
		- [Graduation Criteria](#graduation-criteria)
			- [Maintainers](#maintainers)
		- [Release process, artifacts](#release-process-artifacts)
		- [Risks and Mitigations](#risks-and-mitigations)

<!-- /TOC -->

## Summary

[ExternalDNS](https://github.com/kubernetes-sigs/external-dns) is a project that synchronizes Kubernetes’ Services, Ingresses and other Kubernetes resources to DNS backends for several DNS providers.

The projects was started as a Kubernetes Incubator project in February 2017 and being the Kubernetes incubation initiative officially over, the maintainers want to propose the project to be moved to the kubernetes GitHub organization or to kubernetes-sigs, under the sponsorship of sig-network.

## Motivation

ExternalDNS started as a community project with the goal of unifying several existing projects that were trying to solve the same problem: create DNS records for Kubernetes resources on several DNS backends.

When the project was proposed (see the [original discussion](https://github.com/kubernetes/kubernetes/issues/28525#issuecomment-270766227)), there were at least 3 existing implementations of the same functionality:

* Mate - [https://github.com/linki/mate](https://github.com/linki/mate)

* DNS-controller from kops - [https://github.com/kubernetes/kops/tree/master/dns-controller](https://github.com/kubernetes/kops/tree/master/dns-controller)

* Route53-kubernetes - [https://github.com/wearemolecule/route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

ExternalDNS’ goal from the beginning was to provide an officially supported solution to those problems.

After two years of development, the project is still in the kubernetes-sigs.

The incubation has been officially discontinued and to quote @thockin "Incubator projects should either become real projects in Kubernetes, shut themselves down, or move elsewhere" (see original thread [here](https://groups.google.com/forum/#!topic/kubernetes-sig-network/fvpDC_nxtEM)).

This KEP proposes to move ExternalDNS to the main Kubernetes organization or kubernetes-sigs. The "Proposal" section details the reasons behind it.

### Goals

The only goal of this KEP is to establish consensus regarding the future of the ExternalDNS project and determine where it belongs.

## Proposal

This KEP is about moving External DNS out of the Kubernetes incubator. This section will cover the reasons why External DNS is useful and what the community would miss in case the project would be discontinued or moved under another organization.

External DNS...

* Is the de facto solution to create DNS records for several Kubernetes resources.

* Is a vital component to achieve an experience close to a PaaS that many Kubernetes users try to replicate on top of Kubernetes, by allowing to automatically create DNS records for web applications.

* Supports already 18 different DNS providers including all major public clouds (AWS, Azure, GCP).

Given that the kubernetes-sigs organization will eventually be shut down, the possible alternatives to moving to be an official Kubernetes project are the following:

* Shut down the project

* Move the project elsewhere

We believe that those alternatives would result in a worse outcome for the community compared to moving the project to the any of the other official Kubernetes organizations.
In fact, shutting down ExternalDNS can cause:

* The community to rebuild the same solution as already happened multiple times before the project was launched. Currently ExternalDNS is easy to be found, referenced in many articles/tutorials and for that reason not exposed to that risk.

* Existing users of the projects to be left without a future proof working solution.

Moving the ExternalDNS project outside of Kubernetes projects would cause:

* Problems (re-)establishing user trust which could eventually lead to fragmentation and duplication.

* It would be hard to establish in which organization the project should be moved to. The most natural would be Zalando’s organization, being the company that put most of the work on the project. While it is possible to assume Zalando’s commitment to open-source, that would be a strategic mistake for the project community and for the Kubernetes ecosystem due to the obvious lack of neutrality.

* Lack of resources to test, lack of issue management via automation.

For those reasons, we propose to move ExternalDNS out of the Kubernetes incubator, to live either under the kubernetes or kubernetes-sigs organization to keep being a vital part of the Kubernetes ecosystem.


## Details

### Graduation Criteria

ExternalDNS is a two years old project widely used in production by many companies. The implementation for the three major cloud providers (AWS, Azure, GCP) is stable, not changing its logic and the project is being used in production by many company using Kubernetes.

We have evidence that many companies are using ExternalDNS in production, but it is out of scope for this proposal to collect a comprehensive list of companies.

The project was quoted by a number of tutorials on the web, including the [official tutorials from AWS](https://aws.amazon.com/blogs/opensource/unified-service-discovery-ecs-kubernetes/).

ExternalDNS can’t be consider to be "done": while the core functionality has been implemented, there is lack of integration testing and structural changes that are needed.

Those are identified in the project roadmap, which is roughly made of the following items:

* Decoupling of the providers

    * Implementation proposal

    * Development

* Bug fixing and performance optimization (i.e. rate limiting on cloud providers)

* Integration testing suite, to be implemented at least for the "stable" providers

For those reasons, we consider ExternalDNS to be in Beta state as a project. We believe that once the items mentioned above will be implemented, the project can reach a declared GA status.

There are a number of other factors that need to be covered to fully describe the state of the project, including who are the maintainers, the way we release and manage the project and so on.

#### Maintainers

The project has the following maintainers:

* hjacobs

* Raffo

* linki

* njuettner

The list of maintainers shrunk over time as people moved out of the original development team (all the team members were working at Zalando at the time of project creation) and the project required less work.

The high number of providers contributed to the project pose a maintainability challenge: it is hard to bring the providers forward in terms of functionalities or even test them. The maintainers believe that the plan to transform the current Provider interface from a Go interface to an API will allow for enough decoupling and to hand over the maintenance of those plugins to the contributors themselves, see the risk and mitigations section for further details.

### Release process, artifacts

The project uses the free quota of TravisCI to run tests for the project.

The release pipeline for the project is currently fully owned by Zalando. It runs on the internal system of the company (closed source) which external maintainers/users can’t access and that pushes images to the publicly accessible docker registry available at the URL `registry.opensource.zalan.do`.

The docker registry service is provided as best effort with no sort of SLA and the maintainers team openly suggests the users to build and maintain their own docker image based on the provided Dockerfiles.

Providing a vanity URL for the docker images was consider a non goal till now, but the community seems to be wanting official images from a GCR domain, similarly to what is available for other parts of official Kubernetes projects.

ExternalDNS does not follow a specific release cycle. Releases are made often when there are major contributions (i.e. new providers) or important bug fixes. That said, the master is considered stable and can be used as well to build images.

### Risks and Mitigations

The following are risks that were identified:

* Low number of maintainers: we are currently facing issues keeping up with the number of pull requests and issues giving the low number of maintainers. The list of maintainers already shrunk from 8 maintainers to 4.

* Issues maintaining community contributed providers: we often lack access to external providers (i.e. InfoBlox, etc.) and this means that we cannot verify the implementations and/or run regression tests that go beyond unit testing.

* Somewhat low quality of releases due to lack of integration testing.

We think that the following actions will constitute appropriate mitigations:

* Decoupling the providers via an API will allow us to resolve the problem of the providers. Being the project already more than 2 years old and given that there are 18 providers implemented, we possess enough informations to define an API that we can be stable in a short timeframe. Once this is stable, the problem of testing the providers can be deferred to be a provider’s responsibility. This will also reduce the scope of External DNS core code, which means that there will be no need for a further increase of the maintaining team.

* We added integration testing for the main cloud providers to the roadmap for the 1.0 release to make sure that we cover the mostly used ones. We believe that this item should be tackled independently from the decoupling of providers as it would be capable of generating value independently from the result of the decoupling efforts.

* With the move to the Kubernetes incubation, we hope that we will be able to access the testing resources of the Kubernetes project. In this way, we hope to decouple the project from the dependency on Zalando’s internal CI tool. This will help open up the possibility to increase the visibility on the project from external contributors, which currently would be blocked by the lack of access to the software used for the whole release pipeline.
