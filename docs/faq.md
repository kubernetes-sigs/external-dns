# Frequently asked questions

### When would ExternalDNS become useful to me?

You've probably created many deployments. Typically, you expose your deployment to the Internet by creating a service with a type of load balancer. Depending on your environment, this usually assigns a random publicly available endpoint to your service that you can access from anywhere in the world. On Google Container Engine, this is a public IP address:

```console
$ kubectl get svc
NAME      CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
nginx     10.3.249.226   35.187.104.85   80:32281/TCP   1m
```

But dealing with IPs for service discovery isn't nice, so you register this IP with your DNS provider under a better name—most likely, one that corresponds to your service name. If the IP changes, you update the DNS record accordingly.

Those times are over! ExternalDNS takes care of that last step for you by keeping your DNS records synchronized with your external entry points.

ExternalDNS' usefulness also becomes clear when you use Ingresses to allow external traffic into your cluster. Via Ingress, you can tell Kubernetes to route traffic to different services based on certain HTTP request attributes, e.g. the Host header:

```console
$ kubectl get ing
NAME         HOSTS                                      ADDRESS         PORTS     AGE
entrypoint   frontend.example.org,backend.example.org   35.186.250.78   80        1m
```

But there's nothing that actually makes clients resolve those hostnames with the Ingress' IP address. Again, you normally have to register each entry with your DNS provider. Only if you're lucky can you use a wildcard, like in the example above.

EnternalDNS can solve this for you as well.

### Which DNS providers are supported?

So far, Google CloudDNS and AWS Route 53 with ALIAS records. There's interest in supporting CoreDNS and Azure DNS. We're open to discussing/adding other providers if the community believes it would be valuable.

Initial support for Google CloudDNS is available in the `v0.1` release. Initial support for AWS Route 53 is targeted for the `v0.2` release. However, you can already test it with CNAME instead of ALIAS, using version `v0.1.0`.

There are no plans regarding other providers at the moment.

### Which Kubernetes objects are supported?

Services exposed via type LoadBalancer and for the hostnames defined in Ingress objects. It also seems useful to expose Services with type NodePort to point to your cluster's nodes directly, but there's no commitment to doing this yet.

### Which Service and Ingress controllers are supported?

Regarding Services, we'll support the OSI Layer 4 load balancers that Kubernetes creates on AWS and Google Container Engine, and possibly other clusters running on Google Compute Engine.

Regarding Ingress, we'll support:
* Google's Ingress Controller on GKE that integrates with their Layer 7 load balancers (GLBC)
* nginx-ingress-controller v0.9.x with a fronting Service
* Zalando's [AWS Ingress controller](https://github.com/zalando-incubator/kube-ingress-aws-controller), based on AWS ALBs and [Skipper](https://github.com/zalando/skipper)

### What about those other implementations?

ExternalDNS is a joint effort to unify different projects accomplishing the same goals, namely:

* Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* Zalando's [Mate](https://github.com/zalando-incubator/mate)
* Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

We strive to make the migration from these implementations a smooth experience. This means that, for some time, we'll support their annotation semantics in ExternalDNS and allow both implementations to run side-by-side. This enables you to migrate incrementally and slowly phase out the other implementation.

### How does it work with other implementations and legacy records?

ExternalDNS will allow you to opt into any Services and Ingresses that you want it to consider, by an annotation. This way, it can co-exist with other implementations running in the same cluster if they also support this pattern. However, we'll most likely declare ExternalDNS to be the default implementation. This means that ExternalDNS will consider Services and Ingresses that don't specifically declare which controller they want to be processed by; this is similar to the `ingress.class` annotation on GKE.

### I'm afraid you will mess up my DNS records!

ExternalDNS will implement the concept of owning DNS records. This means that ExternalDNS will keep track of which records it has control over, and will never modify any records over which it doesn't have control. This is a fundamental requirement to operate ExternalDNS safely when there might be other actors creating DNS records in the same target space.

However, this is a delicate topic and hasn't yet found its way into ExternalDNS.

### Does anyone use ExternalDNS in production?

No — but ExternalDNS is heavily influenced by Zalando's [Mate](https://github.com/zalando-incubator/mate), which is used in production on AWS. If you want to adopt this approach and need a solution now, then try Mate. Otherwise, we encourage you to stick with ExternalDNS and help us make it work for you.

### How can we start using ExternalDNS?

ExternalDNS is in an early state and not yet recommended for production use. However, you can start trying it out on a non-production GKE cluster following [the GKE tutorial](tutorials/gke.md).

Clusters on AWS that want to make use of Route 53 work very similar, but a tutorial is still on our TODO list.
