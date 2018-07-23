# Setting up ExternalDNS using AWS Service Discovery API

This tutorial describes how to set up ExternalDNS for usage within a Kubernetes cluster on AWS with [Service Discovery API](https://docs.aws.amazon.com/Route53/latest/APIReference/overview-service-discovery.html).

The **Service Discovery API** is an alternative approach to managing DNS records directly using the Route53 API. It is more suitable for a dynamic environment where service endpoints change frequently. It abstracts away technical details of the DNS protocol and offers a simplified model. Service discovery consists of three main API calls:

* CreatePublicDnsNamespace – automatically creates a DNS hosted zone
* CreateService – creates a new named service inside the specified namespace
* RegisterInstance/DeregisterInstance – can be called multiple times to create a DNS record for the specified *Service*

Learn more about the API in the [Amazon Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_Operations_Amazon_Route_53_Auto_Naming.html).


## IAM Permissions

To use the service discovery API, a user executing the ExternalDNS must have the permissions in the `AmazonRoute53AutoNamingFullAccess` managed policy.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "route53:GetHostedZone",
                "route53:ListHostedZonesByName",
                "route53:CreateHostedZone",
                "route53:DeleteHostedZone",
                "route53:ChangeResourceRecordSets",
                "route53:CreateHealthCheck",
                "route53:GetHealthCheck",
                "route53:DeleteHealthCheck",
                "route53:UpdateHealthCheck",
                "ec2:DescribeVpcs",
                "ec2:DescribeRegions",
                "servicediscovery:*"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
```

## Set up a namespace

Create a DNS namespace using the service discovery API

```console
$ aws servicediscovery create-public-dns-namespace --name "external-dns-test.my-org.com"
```

Verify that the namespace was truly created

```console
$ aws servicediscovery list-namespaces
```

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster that you want to test ExternalDNS with.
Then apply the following manifest file to deploy ExternalDNS.

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      containers:
      - name: external-dns
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=external-dns-test.my-org.com # Makes ExternalDNS see only the namespaces that match the specified domain. Omit the filter if you want to process all available namespaces.
        - --provider=aws-sd
        - --aws-zone-type=public # Only look at public namespaces. Valid values are public, private, or no value for both)
        - --txt-owner-id=my-identifier
```


## Verify that ExternalDNS works (Service example)

Create the following sample application to test that ExternalDNS works.

> For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com
spec:
  type: LoadBalancer
  ports:
  - port: 80
    name: http
    targetPort: 80
  selector:
    app: nginx

---

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
        ports:
        - containerPort: 80
          name: http
```

After one minute check that a corresponding DNS record for your service was created in your hosted zone. We recommended that you use the [Amazon Route53 console](https://console.aws.amazon.com/route53) for that purpose.


## Custom TTL

The default DNS record TTL (time to live) is 300 seconds. You can customize this value by setting the annotation `external-dns.alpha.kubernetes.io/ttl`.
For example, modify the service manifest YAML file above:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com
    external-dns.alpha.kubernetes.io/ttl: 60
spec:
    ...
```

This will set the TTL for the DNS record to 60 seconds.


## Clean up

Delete all service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the remaining service and namespace.

```console
$ aws servicediscovery list-services

{
    "Services": [
        {
            "Id": "srv-6dygt5ywvyzvi3an",
            "Arn": "arn:aws:servicediscovery:us-west-2:861574988794:service/srv-6dygt5ywvyzvi3an",
            "Name": "nginx"
        }
    ]
}
```

```console
$ aws servicediscovery delete-service --id srv-6dygt5ywvyzvi3an
```

```console
$ aws servicediscovery list-namespaces
{
    "Namespaces": [
        {
            "Type": "DNS_PUBLIC",
            "Id": "ns-durf2oxu4gxcgo6z",
            "Arn": "arn:aws:servicediscovery:us-west-2:861574988794:namespace/ns-durf2oxu4gxcgo6z",
            "Name": "external-dns-test.my-org.com"
        }
    ]
}
```

```console
$ aws servicediscovery delete-namespace --id ns-durf2oxu4gxcgo6z
```
