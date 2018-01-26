# Frequently asked questions

### When would ExternalDNS become useful to me?

You've probably created many deployments. Typically, you expose your deployment to the Internet by creating a Service with `type=LoadBalancer`. Depending on your environment, this usually assigns a random publicly available endpoint to your service that you can access from anywhere in the world. On Google Container Engine, this is a public IP address:

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

But there's nothing that actually makes clients resolve those hostnames to the Ingress' IP address. Again, you normally have to register each entry with your DNS provider. Only if you're lucky can you use a wildcard, like in the example above.

EnternalDNS can solve this for you as well.

### Which DNS providers are supported?

So far, Google CloudDNS and AWS Route 53 with ALIAS records. There's interest in supporting CoreDNS and Azure DNS. We're open to discussing/adding other providers if the community believes it would be valuable.

Initial support for Google CloudDNS is available since the `v0.1` release. Initial support for AWS Route 53 is available in the `v0.2` release (CNAME based) and ALIAS is targeted for the `v0.3` release.

There are no plans regarding other providers at the moment.

### Which Kubernetes objects are supported?

Services exposed via `type=LoadBalancer` and for the hostnames defined in Ingress objects. It also seems useful to expose Services with `type=NodePort` to point to your cluster's nodes directly, but there's no commitment to doing this yet.

### How do I specify DNS name for my Kubernetes objects?

There are three sources of information for ExternalDNS to decide on DNS name. ExternalDNS will pick one in order as listed below: 

1. For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object. For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

2. If compatibility mode is enabled (e.g. `--compatibility={mate,molecule}` flag), External DNS will parse annotations used by Zalando/Mate, wearemolecule/route53-kubernetes. Compatibility mode with Kops DNS Controller is planned to be added in the future.

3. If `--fqdn-template` flag is specified, e.g. `--fqdn-template={{.Name}}.my-org.com`, ExternalDNS will use service/ingress specifications for the provided template to generate DNS name.

### Which Service and Ingress controllers are supported?

Regarding Services, we'll support the OSI Layer 4 load balancers that Kubernetes creates on AWS and Google Container Engine, and possibly other clusters running on Google Compute Engine.

Regarding Ingress, we'll support:
* Google's Ingress Controller on GKE that integrates with their Layer 7 load balancers (GLBC)
* nginx-ingress-controller v0.9.x with a fronting Service
* Zalando's [AWS Ingress controller](https://github.com/zalando-incubator/kube-ingress-aws-controller), based on AWS ALBs and [Skipper](https://github.com/zalando/skipper)

### Are other Ingress Controllers supported?

For Ingress objects, ExternalDNS will attempt to discover the target hostname of the relevant Ingress Controller automatically. If you are using an Ingress Controller that is not listed above you may have issues with ExternalDNS not discovering Endpoints and consequently not creating any DNS records. As a workaround, it is possible to force create an Endpoint by manually specifying a target host/IP for the records to be created by setting the annotation `external-dns.alpha.kubernetes.io/target` in the Ingress object.

Another reason you may want to override the ingress hostname or IP address is if you have an external mechanism for handling failover across ingress endpoints.  Possible scenarios for this would include using [keepalived-vip](https://github.com/kubernetes/contrib/tree/master/keepalived-vip) to manage failover faster than DNS TTLs might expire.

Note that if you set the target to a hostname, then a CNAME record will be created.  In this case, the hostname specified in the Ingress object's annotation must already exist. (i.e.: You have a Service resource for your Ingress Controller with the `external-dns.alpha.kubernetes.io/hostname` annotation set to the same value.)

### What about those other projects?

ExternalDNS is a joint effort to unify different projects accomplishing the same goals, namely:

* Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/master/dns-controller)
* Zalando's [Mate](https://github.com/linki/mate)
* Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

We strive to make the migration from these implementations a smooth experience. This means that, for some time, we'll support their annotation semantics in ExternalDNS and allow both implementations to run side-by-side. This enables you to migrate incrementally and slowly phase out the other implementation.

### How does it work with other implementations and legacy records?

ExternalDNS will allow you to opt into any Services and Ingresses that you want it to consider, by an annotation. This way, it can co-exist with other implementations running in the same cluster if they also support this pattern. However, we'll most likely declare ExternalDNS to be the default implementation. This means that ExternalDNS will consider Services and Ingresses that don't specifically declare which controller they want to be processed by; this is similar to the `ingress.class` annotation on GKE.

### I'm afraid you will mess up my DNS records!

ExternalDNS since v0.3 implements the concept of owning DNS records. This means that ExternalDNS will keep track of which records it has control over, and will never modify any records over which it doesn't have control. This is a fundamental requirement to operate ExternalDNS safely when there might be other actors creating DNS records in the same target space.

For now ExternalDNS uses TXT records to label owned records, and there might be other alternatives coming in the future releases.

### Does anyone use ExternalDNS in production?

Yes — Zalando replaced [Mate](https://github.com/linki/mate) with ExternalDNS since its v0.3 release, which now runs in production-level clusters. We are planning to document a step-by-step tutorial on how the switch from Mate to ExternalDNS has occurred.

### How can we start using ExternalDNS?

Check out the following descriptive tutorials on how to run ExternalDNS in [GKE](tutorials/gke.md) and [AWS](tutorials/aws.md). 

### Why is ExternalDNS only adding a single IP address in Route 53 on AWS when using the `nginx-ingress-controller`? How do I get it to use the FQDN of the ELB assigned to my `nginx-ingress-controller` Service instead?

By default the `nginx-ingress-controller` assigns a single IP address to an Ingress resource when it's created. ExternalDNS uses what's assigned to the Ingress resource, so it too will use this single IP address when adding the record in Route 53.

In most AWS deployments, you'll instead want the Route 53 entry to be the FQDN of the ELB that is assigned to the `nginx-ingress-controller` Service. To accomplish this, when you create the `nginx-ingress-controller` Deployment, you need to provide the `--publish-service` option to the `/nginx-ingress-controller` executable under `args`. Once this is deployed new Ingress resources will get the ELB's FQDN and ExternalDNS will use the same when creating records in Route 53. 

According to the `nginx-ingress-controller` [docs](https://github.com/kubernetes/ingress/tree/master/controllers/nginx) the value you need to provide `--publish-service` is:

> Service fronting the ingress controllers. Takes the form namespace/name. The controller will set the endpoint records on the ingress objects to reflect those on the service.

For example if your `nginx-ingress-controller` Service's name is `nginx-ingress-controller-svc` and it's in the `default` namespace the start of your resource YAML might look like the following. Note the second to last line.

```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx-ingress-controller
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx-ingress
    spec:
      hostNetwork: false
      containers:
        - name: nginx-ingress-controller
          image: "gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.11"
          imagePullPolicy: "IfNotPresent"
          args:
            - /nginx-ingress-controller
            - --default-backend-service={your-backend-service}
            - --publish-service=default/nginx-ingress-controller-svc
            - --configmap={your-configmap}
```

### I have a Service/Ingress but it's ignored by ExternalDNS. Why?

TODO (https://github.com/kubernetes-incubator/external-dns/issues/267)

### I'm using an ELB with TXT registry but the CNAME record clashes with the TXT record. How to avoid this?

TODO (https://github.com/kubernetes-incubator/external-dns/issues/262)

### Which permissions do I need when running ExternalDNS on a GCE or GKE node.

You need to add either https://www.googleapis.com/auth/ndev.clouddns.readwrite or https://www.googleapis.com/auth/cloud-platform on your instance group's scope.

### How can I run ExternalDNS under a specific GCP Service Account, e.g. to access DNS records in other projects?

Have a look at https://github.com/linki/mate/blob/v0.6.2/examples/google/README.md#permissions

### How do I configure multiple Sources via environment variables? (also applies to domain filters)

Separate the individual values via a line break. The equivalent of `--source=service --source=ingress` would be `service\ningress`. However, it can be tricky do define that depending on your environment. The following examples work (zsh):

Via docker:

```console
$ docker run \
  -e EXTERNAL_DNS_SOURCE=$'service\ningress' \
  -e EXTERNAL_DNS_PROVIDER=google \
  -e EXTERNAL_DNS_DOMAIN_FILTER=$'foo.com\nbar.com' \
  registry.opensource.zalan.do/teapot/external-dns:v0.4.8
time="2017-08-08T14:10:26Z" level=info msg="config: &{Master: KubeConfig: Sources:[service ingress] Namespace: ...
```

Locally:

```console
$ export EXTERNAL_DNS_SOURCE=$'service\ningress'
$ external-dns --provider=google
INFO[0000] config: &{Master: KubeConfig: Sources:[service ingress] Namespace: ...
```

```
$ EXTERNAL_DNS_SOURCE=$'service\ningress' external-dns --provider=google
INFO[0000] config: &{Master: KubeConfig: Sources:[service ingress] Namespace: ...
```

In a Kubernetes manifest:

```yaml
spec:
  containers:
  - name: external-dns
    args:
    - --provider=google
    env:
    - name: EXTERNAL_DNS_SOURCE
      value: "service\ningress"
```

Or preferably:

```yaml
spec:
  containers:
  - name: external-dns
    args:
    - --provider=google
    env:
    - name: EXTERNAL_DNS_SOURCE
      value: |-
        service
        ingress
```


### Running an internal and external dns service

Sometimes you need to run an internal and an external dns service.
The internal one should provision hostnames used on the internal network (perhaps inside a VPC), and the external
one to expose DNS to the internet.

To do this with external-dns you can use the `--annotation-filter` to specifically tie an instance of external-dns to 
an instance of a ingress controller. Let's assume you have two ingress controllers `nginx-internal` and `nginx-external`
then you can start two external-dns providers one with `--annotation-filter=kubernetes.io/ingress.class=nginx-internal` 
and one with `--annotation-filter=kubernetes.io/ingress.class=nginx-external`  