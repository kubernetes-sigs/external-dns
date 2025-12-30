# AWS Load Balancer Controller

This tutorial describes how to use ExternalDNS with the [aws-load-balancer-controller][1].

[1]: https://kubernetes-sigs.github.io/aws-load-balancer-controller

## Setting up ExternalDNS and aws-load-balancer-controller

Follow the [AWS tutorial](aws.md) to setup ExternalDNS for use in Kubernetes clusters
running in AWS. Specify the `source=ingress` argument so that ExternalDNS will look
for hostnames in Ingress objects. In addition, you may wish to limit which Ingress
objects are used as an ExternalDNS source via the `ingress-class` argument, but
this is not required.

For help setting up the AWS Load Balancer Controller, follow the [Setup Guide][2].

[2]: https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/deploy/installation/

Note that the AWS Load Balancer Controller uses the same tags for [subnet auto-discovery][3]
as Kubernetes does with the AWS cloud provider.

[3]: https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/deploy/subnet_discovery/

In the examples that follow, it is assumed that you configured the ALB Ingress
Controller with the `ingress-class=alb` argument (not to be confused with the
same argument to ExternalDNS) so that the controller will only respect Ingress
objects with the `ingressClassName` field set to "alb".

## Deploy an example application

Create the following sample "echoserver" application to demonstrate how
ExternalDNS works with ALB ingress objects.

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
  type: NodePort
  selector:
    app: echoserver
```

Note that the Service object is of type `NodePort`. We don't need a Service of
type `LoadBalancer` here, since we will be using an Ingress to create an ALB.

## Ingress examples

Create the following Ingress to expose the echoserver application to the Internet.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing
  name: echoserver
spec:
  ingressClassName: alb
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
traffic to the echoserver application.

If the `--source=ingress` argument is specified, then ExternalDNS will create
DNS records based on the hosts specified in ingress objects. The above example
would result in two alias records (A and AAAA) being created for each of the
domains: `echoserver.mycluster.example.org` and `echoserver.example.org`. All
four records alias the ALB that is associated with the Ingress object. As the
ALB is IPv4 only, the AAAA alias records have no effect.

If you would like ExternalDNS to not create AAAA records at all, you can add the
following command line parameter: `--exclude-record-types=AAAA`. Please be
aware, this will disable AAAA record creation even for dualstack enabled load
balancers.

Note that the above example makes use of the YAML anchor feature to avoid having
to repeat the http section for multiple hosts that use the exact same paths. If
this Ingress object will only be fronting one backend Service, we might instead
create the following:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing
    external-dns.alpha.kubernetes.io/hostname: echoserver.mycluster.example.org, echoserver.example.org
  name: echoserver
spec:
  ingressClassName: alb
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

## Dualstack Load Balancers

AWS [supports both IPv4 and "dualstack" (both IPv4 and IPv6) interfaces for ALBs][4]
and [NLBs][5]. The AWS Load Balancer Controller uses the `alb.ingress.kubernetes.io/ip-address-type`
annotation (which defaults to `ipv4`) to determine this. ExternalDNS creates
both A and AAAA alias DNS records by default, regardless of this annotation.
It's possible to create only A records with the following command line
parameter: `--exclude-record-types=AAAA`

[4]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#ip-address-type
[5]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#ip-address-type

Example:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing
    alb.ingress.kubernetes.io/ip-address-type: dualstack
  name: echoserver
spec:
  ingressClassName: alb
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
interface.

## Frontend Network Load Balancer (NLB)

The AWS Load Balancer Controller supports [fronting ALBs with an NLB][6] for improved performance
and static IP addresses. When this feature is enabled, the controller creates both an ALB and an
NLB, resulting in two hostnames in the Ingress status.

[6]: https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/guide/ingress/annotations/#enable-frontend-nlb

### Known Issue with Internal ALBs

When using an internal ALB (`alb.ingress.kubernetes.io/scheme: internal`) with frontend NLB,
ExternalDNS may create DNS records pointing to the ALB instead of the NLB due to alphabetical
ordering:

- Internal ALB hostname: `internal-k8s-myapp-alb.us-east-1.elb.amazonaws.com`
- NLB hostname: `k8s-myapp-nlb-123456789.elb.us-east-1.amazonaws.com`

When multiple targets exist, Route53 selects the first one alphabetically, which incorrectly
selects the internal ALB. See [issue #5661][7] for details.

[7]: https://github.com/kubernetes-sigs/external-dns/issues/5661

### Workarounds

There are several approaches to ensure DNS records point to the correct (NLB) target:

#### Option 1: Combine load balancer naming with target annotation (Recommended)

Use [`alb.ingress.kubernetes.io/load-balancer-name`][8] to create predictable hostnames, then
explicitly reference the NLB using [`external-dns.alpha.kubernetes.io/target`][9]:

[8]: https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/guide/ingress/annotations/#load-balancer-name
[9]: https://kubernetes-sigs.github.io/external-dns/latest/docs/annotations/annotations/#external-dnsalphakubernetesiotarget

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/scheme: internal
    alb.ingress.kubernetes.io/enable-frontend-nlb: "true"
    alb.ingress.kubernetes.io/frontend-nlb-scheme: internal
    alb.ingress.kubernetes.io/load-balancer-name: myapp-alb
    external-dns.alpha.kubernetes.io/target: k8s-myapp-nlb.elb.us-east-1.amazonaws.com
  name: echoserver
spec:
  ingressClassName: alb
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

**Benefits**:

- Predictable, consistent load balancer naming across environments
- Explicit control over which target ExternalDNS uses
- Works reliably with internal ALBs
- No need to lookup auto-generated NLB names

**NLB hostname pattern**: When you set `load-balancer-name: myapp-alb`, the NLB hostname
becomes `k8s-myapp-nlb.elb.<region>.amazonaws.com` (note the `-nlb` suffix).

**ALB internal hostname pattern**: When you set `load-balancer-name: myapp-alb`, the ALB hostname
becomes `internal-myapp-nlb.<region>.elb.amazonaws.com` (note the `-nlb` suffix).
(note the `internal-` suffix)

#### Option 2: Use the target annotation only

If you cannot control the load balancer name, explicitly specify the NLB hostname:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/scheme: internal
    alb.ingress.kubernetes.io/enable-frontend-nlb: "true"
    alb.ingress.kubernetes.io/frontend-nlb-scheme: internal
    external-dns.alpha.kubernetes.io/target: k8s-myapp-nlb-123456789.elb.us-east-1.amazonaws.com
  name: echoserver
spec:
  ingressClassName: alb
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

**Note**: You'll need to lookup the auto-generated NLB hostname after the controller creates it.

#### Option 3: Use a DNSEndpoint resource

Create a `DNSEndpoint` custom resource to explicitly define the DNS record:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: echoserver-dns
spec:
  endpoints:
  - dnsName: echoserver.example.org
    recordType: CNAME
    targets:
    - k8s-myapp-nlb-123456789.elb.us-east-1.amazonaws.com
```

This approach is useful when you want to manage DNS records independently of the Ingress resource.
