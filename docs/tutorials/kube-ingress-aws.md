# Using ExternalDNS with kube-ingress-aws-controller

This tutorial describes how to use ExternalDNS with the [kube-ingress-aws-controller][1].

[1]: https://github.com/zalando-incubator/kube-ingress-aws-controller

## Setting up ExternalDNS and kube-ingress-aws-controller

Follow the [AWS tutorial](aws.md) to setup ExternalDNS for use in Kubernetes clusters
running in AWS. Specify the `source=ingress` argument so that ExternalDNS will look
for hostnames in Ingress objects. In addition, you may wish to limit which Ingress
objects are used as an ExternalDNS source via the `ingress-class` argument, but
this is not required.

For help setting up the Kubernetes Ingress AWS Controller, that can
create ALBs and NLBs, follow the [Setup Guide][2].

[2]: https://github.com/zalando-incubator/kube-ingress-aws-controller/tree/HEAD/deploy


### Optional RouteGroup

[RouteGroup][3] is a CRD, that enables you to do complex routing with
[Skipper][4].

First, you have to apply the RouteGroup CRD to your cluster:

```
kubectl apply -f https://github.com/zalando/skipper/blob/HEAD/dataclients/kubernetes/deploy/apply/routegroups_crd.yaml
```

You have to grant all controllers: [Skipper][4],
[kube-ingress-aws-controller][1] and external-dns to read the routegroup resource and
kube-ingress-aws-controller to update the status field of a routegroup.
This depends on your RBAC policies, in case you use RBAC, you can use
this for all 3 controllers:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: kube-ingress-aws-controller
rules:
- apiGroups:
  - extensions
  - networking.k8s.io
  resources:
  - ingresses
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - extensions
  - networking.k8s.io
  resources:
  - ingresses/status
  verbs:
  - patch
  - update
- apiGroups:
  - zalando.org
  resources:
  - routegroups
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - zalando.org
  resources:
  - routegroups/status
  verbs:
  - patch
  - update
```

See also current RBAC yaml files:
- [kube-ingress-aws-controller](https://github.com/zalando-incubator/kubernetes-on-aws/blob/dev/cluster/manifests/ingress-controller/01-rbac.yaml)
- [skipper](https://github.com/zalando-incubator/kubernetes-on-aws/blob/dev/cluster/manifests/skipper/rbac.yaml)
- [external-dns](https://github.com/zalando-incubator/kubernetes-on-aws/blob/dev/cluster/manifests/external-dns/01-rbac.yaml)

[3]: https://opensource.zalando.com/skipper/kubernetes/routegroups/#routegroups
[4]: https://opensource.zalando.com/skipper


## Deploy an example application

Create the following sample "echoserver" application to demonstrate how
ExternalDNS works with ingress objects, that were created by [kube-ingress-aws-controller][1].

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echoserver
spec:
  replicas: 1
  selector:
    matchLabels:
      app: echoserver
  template:
    metadata:
      labels:
        app: echoserver
    spec:
      containers:
      - image: gcr.io/google_containers/echoserver:1.4
        imagePullPolicy: Always
        name: echoserver
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: echoserver
spec:
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
  type: ClusterIP
  selector:
    app: echoserver
```

Note that the Service object is of type `ClusterIP`, because we will
target [Skipper][4] and do the HTTP routing in Skipper. We don't need
a Service of type `LoadBalancer` here, since we will be using a shared
skipper-ingress for all Ingress. Skipper use `hostNetwork` to be able
to get traffic from AWS LoadBalancers EC2 network. ALBs or NLBs, will
be created based on need and will be shared across all ingress as
default.

## Ingress examples

Create the following Ingress to expose the echoserver application to the Internet.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: skipper
  name: echoserver
spec:
  ingressClassName: skipper
  rules:
  - host: echoserver.mycluster.example.org
    http: &echoserver_root
      paths:
      - path: /
        backend:
          service:
            name: echoserver
            port:
              number: 80
        pathType: Prefix
  - host: echoserver.example.org
    http: *echoserver_root
```

The above should result in the creation of an (ipv4) ALB in AWS which will forward
traffic to skipper which will forward to the echoserver application.

If the `--source=ingress` argument is specified, then ExternalDNS will create DNS
records based on the hosts specified in ingress objects. The above example would
result in two alias records being created, `echoserver.mycluster.example.org` and
`echoserver.example.org`, which both alias the ALB that is associated with the
Ingress object.

Note that the above example makes use of the YAML anchor feature to avoid having
to repeat the http section for multiple hosts that use the exact same paths. If
this Ingress object will only be fronting one backend Service, we might instead
create the following:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: echoserver.mycluster.example.org, echoserver.example.org
    kubernetes.io/ingress.class: skipper
  name: echoserver
spec:
  ingressClassName: skipper
  rules:
  - http:
      paths:
      - path: /
        backend:
          service:
            name: echoserver
            port:
              number: 80
        pathType: Prefix
```

In the above example we create a default path that works for any hostname, and
make use of the `external-dns.alpha.kubernetes.io/hostname` annotation to create
multiple aliases for the resulting ALB.

## Dualstack ALBs

AWS [supports](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#ip-address-type) both IPv4 and "dualstack" (both IPv4 and IPv6) interfaces for ALBs.
The Kubernetes Ingress AWS controller supports the `alb.ingress.kubernetes.io/ip-address-type`
annotation (which defaults to `ipv4`) to determine this. If this annotation is
set to `dualstack` then ExternalDNS will create two alias records (one A record
and one AAAA record) for each hostname associated with the Ingress object.


Example:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/ip-address-type: dualstack
    kubernetes.io/ingress.class: skipper
  name: echoserver
spec:
  ingressClassName: skipper
  rules:
  - host: echoserver.example.org
    http:
      paths:
      - path: /
        backend:
          service:
            name: echoserver
            port:
              number: 80
        pathType: Prefix
```

The above Ingress object will result in the creation of an ALB with a dualstack
interface. ExternalDNS will create both an A `echoserver.example.org` record and
an AAAA record of the same name, that each are aliases for the same ALB.

## NLBs

AWS has
[NLBs](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html)
and [kube-ingress-aws-controller][1] is able to create NLBs instead of ALBs.
The Kubernetes Ingress AWS controller supports the `zalando.org/aws-load-balancer-type`
annotation (which defaults to `alb`) to determine this. If this annotation is
set to `nlb` then ExternalDNS will create an NLB instead of an ALB.

Example:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    zalando.org/aws-load-balancer-type: nlb
    kubernetes.io/ingress.class: skipper
  name: echoserver
spec:
  ingressClassName: skipper
  rules:
  - host: echoserver.example.org
    http:
      paths:
      - path: /
        backend:
          service:
            name: echoserver
            port:
              number: 80
        pathType: Prefix
```

The above Ingress object will result in the creation of an NLB. A
successful create, you can observe in the ingress `status` field, that is
written by [kube-ingress-aws-controller][1]:

```yaml
status:
  loadBalancer:
    ingress:
    - hostname: kube-ing-lb-atedkrlml7iu-1681027139.$region.elb.amazonaws.com
```

ExternalDNS will create a A-records `echoserver.example.org`, that
use AWS ALIAS record to automatically maintain IP addresses of the NLB.

## RouteGroup (optional)

[Kube-ingress-aws-controller][1], [Skipper][4] and external-dns
support [RouteGroups][3]. External-dns needs to be started with
`--source=skipper-routegroup` parameter in order to work on RouteGroup objects.

Here we can not show [all RouteGroup
capabilities](https://opensource.zalando.com/skipper/kubernetes/routegroups/),
but we show one simple example with an application and a custom https
redirect.

```yaml
apiVersion: zalando.org/v1
kind: RouteGroup
metadata:
  name: my-route-group
spec:
  backends:
  - name: my-backend
    type: service
    serviceName: my-service
    servicePort: 80
  - name: redirectShunt
    type: shunt
  defaultBackends:
  - backendName: my-service
  routes:
  - pathSubtree: /
  - pathSubtree: /
    predicates:
    - Header("X-Forwarded-Proto", "http")
    filters:
    - redirectTo(302, "https:")
    backends:
    - redirectShunt
```
