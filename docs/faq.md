# Frequently asked questions

### When is this useful to me?

You probably have created many Deployments in the past. Often you would expose your Deployment to the Internet by creating a Service with a type of LoadBalancer. Depending on your environment this usually assigns a random publicly available endpoint to your service at which you can access it from anywhere in the world. On Google Container Engine this is a public IP address.

```console
$ kubectl get svc
NAME      CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
nginx     10.3.249.226   35.187.104.85   80:32281/TCP   1m
```

However, dealing with IPs for service discovery isn't nice so you would probably register this IP with your DNS provider under a better name, probably corresponding to your service name. If that IP were to change you would update the DNS record accordingly.

Those times are over, ExternalDNS takes care of that last part for you and keeps your DNS records synchronized with your external entry points.

It starts to become more clear when you use Ingresses to allow external traffic into your cluster. Via Ingress you can tell Kubernetes to route traffic to different services based on certain HTTP request attributes, e.g. the Host header.

```console
$ kubectl get ing
NAME         HOSTS                                      ADDRESS         PORTS     AGE
entrypoint   frontend.example.org,backend.example.org   35.186.250.78   80        1m
```

But there's nothing that actually makes clients to resolve those hostnames to the Ingress' IP address. Again, you would register each entry with your DNS provider and only if you're lucky you could use a wildcard like in this example.

However, EnternalDNS can solve that for you as well.

### What DNS providers are supported?

There will be support for Google CloudDNS and AWS Route53 with ALIAS records. There's a desire to support CoreDNS as well as Azure DNS. We're open to review and add other providers if the community believes them valuable.

Initial support for Google CloudDNS is targeted for the `v0.1` release. You can already test it with version `v0.1.0-beta.1` onwards.
Initial support for AWS Route53 is targeted for the `v0.2` release. However, you can already test it with CNAME instead of ALIAS using version `v0.1.0`.
There are no plans regarding other providers at the moment.

### What Kubernetes objects are supported?

There will be support for Services exposed via type LoadBalancer and for the hostnames defined in Ingress objects. It also seems useful to expose Services with type NodePort to point to your cluster's nodes directly, but there's no commitment, yet.

### Which Service and Ingress controllers are supported?

Regarding Services we'll support the OSI Layer 4 load balancers that are created by Kubernetes on AWS as well as on Google Container Engine and possibly other clusters running on Google Compute Engine.

Regarding Ingress we'll support:
* Google's Ingress Controller on GKE that integrates with their Layer 7 load balancers (GLBC)
* nginx-ingress-controller v0.9.x with a fronting Service
* Zalando's [AWS Ingress controller](https://github.com/zalando-incubator/kube-ingress-aws-controller) based on AWS ALBs and [Skipper](https://github.com/zalando/skipper)

### What about those other implementations?

ExternalDNS is a joint effort to unify different projects accomplishing the same goals in the past, namely:

* Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* Zalando's [Mate](https://github.com/zalando-incubator/mate)
* Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

We strive to make the migration from these implementations a smooth experience. This means for some time we'll support their annotation semantics in ExternalDNS as well as allow both implementations to run side-by-side. This allows you to migrate incrementally and slowly face out the other implementation.

### How does it work with other implementations and legacy records?

ExternalDNS will allow you to opt-in any Services and Ingresses you want it to consider by an annotation. This way it can co-exist with other implementations running in the same cluster if they support this pattern as well. However, we'll most likely declare ExternalDNS to be the default implementation in the future. This means ExternalDNS will consider Services and Ingresses that don't specifically declare which controller they want to be processed by, similar to the `ingress.class` annotation on GKE.

### I'm afraid you mess up my DNS records!

ExternalDNS will implement the concept of owning DNS records. It means that ExternalDNS keeps track of which records it has control over and will never modify any records of which it doesn't. This is a fundamental requirement to operate ExternalDNS safely when there might be other actors creating DNS records in the same target space.

However, this is a delicate topic and hasn't found its way into ExternalDNS, yet.

### Does anyone use ExternalDNS in production?

No, but ExternalDNS is heavily influenced by Zalando's [Mate](https://github.com/zalando-incubator/mate) which is used in production on AWS. If you want to adopt this approach and need a solution now then try Mate. Otherwise we encourage you to stick with ExternalDNS and help us make it work for you.

### How can we start using ExternalDNS?

ExternalDNS is in an early state and not recommended for production use, yet. However, you can start trying it out on a non-production GKE cluster following [the GKE tutorial](tutorials/gke.md).

Cluster's on AWS that want to make use of Route53 work very similar but a tutorial is still on our TODO list.
