# Setting up ExternalDNS for Services on AWS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on AWS. Make sure to use **>=0.4** version of ExternalDNS for this tutorial

## IAM Permissions

```json
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Effect": "Allow",
     "Action": [
       "route53:ChangeResourceRecordSets"
     ],
     "Resource": [
       "arn:aws:route53:::hostedzone/*"
     ]
   },
   {
     "Effect": "Allow",
     "Action": [
       "route53:ListHostedZones",
       "route53:ListResourceRecordSets"
     ],
     "Resource": [
       "*"
     ]
   }
 ]
}
```

When running on AWS, you need to make sure that your nodes (on which External DNS runs) have the IAM instance profile with the above IAM role assigned (either directly or via something like [kube2iam](https://github.com/jtblin/kube2iam)).

## Set up a hosted zone

*If you prefer to try-out ExternalDNS in one of the existing hosted-zones you can skip this step*

Create a DNS zone which will contain the managed DNS records.

```console
$ aws route53 create-hosted-zone --name "external-dns-test.my-org.com." --caller-reference "external-dns-test-$(date +%s)"
```


Make a note of the ID of the hosted zone you just created, which will serve as the value for my-hostedzone-identifier.

```console
$ aws route53 list-hosted-zones-by-name --output json --dns-name "external-dns-test.my-org.com." | jq -r '.HostedZones[0].Id'
/hostedzone/ZEWFWZ4R16P7IB
```

Make a note of the nameservers that were assigned to your new zone.

```console
$ aws route53 list-resource-record-sets --output json --hosted-zone-id "/hostedzone/ZEWFWZ4R16P7IB" \
    --query "ResourceRecordSets[?Type == 'NS']" | jq -r '.[0].ResourceRecords[].Value'
ns-5514.awsdns-53.org.
...
```

In this case it's the ones shown above but your's will differ.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS. You can check if your cluster has RBAC by `kubectl api-versions | grep rbac.authorization.k8s.io`.

### Manifest (for clusters without RBAC enabled)
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
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-hostedzone-identifier
```

### Manifest (for clusters with RBAC enabled)

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"] 
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list","watch"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: external-dns-viewer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
- kind: ServiceAccount
  name: external-dns
  namespace: default
---
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
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-hostedzone-identifier
```



## Arguments

This list is not the full list, but a few arguments that where chosen.

### aws-zone-type

`aws-zone-type` allows filtering for private and public zones

## Annotations

Annotations which are specific to AWS.

### alias

`external-dns.alpha.kubernetes.io/alias` if set to `true` on an ingress, it will create an ALIAS record when the target is an ALIAS as well. To make the target an alias, the ingress needs to be configured correctly as described in [the docs](./nginx-ingress.md#with-a-separate-tcp-load-balancer).

## Verify ExternalDNS works (Ingress example)

Create an ingress resource manifest file.

> For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object.

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  annotations:
    kubernetes.io/ingress.class: "nginx" # use the one that corresponds to your ingress controller.
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - backend:
          serviceName: foo
          servicePort: 80
```

## Verify ExternalDNS works (Service example)

Create the following sample application to test that ExternalDNS works.

> For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

> If you want to give multiple names to service, you can set it to external-dns.alpha.kubernetes.io/hostname with a comma separator.

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

After roughly two minutes check that a corresponding DNS record for your service was created.

```console
$ aws route53 list-resource-record-sets --output json --hosted-zone-id "/hostedzone/ZEWFWZ4R16P7IB" \
    --query "ResourceRecordSets[?Name == 'nginx.external-dns-test.my-org.com.']|[?Type == 'A']"
[
    {
      "AliasTarget": {
          "HostedZoneId": "ZEWFWZ4R16P7IB",
          "DNSName": "ae11c2360188411e7951602725593fd1-1224345803.eu-central-1.elb.amazonaws.com.",
          "EvaluateTargetHealth": true
      },
      "Name": "external-dns-test.my-org.com.",
      "Type": "A"
    },
    {
      "Name": "external-dns-test.my-org.com",
      "TTL": 300,
      "ResourceRecords": [
          {
              "Value": "\"heritage=external-dns,external-dns/owner=my-hostedzone-identifier\""
          }
      ],
      "Type": "TXT"
    }
]
```

Note created TXT record alongside ALIAS record. TXT record signifies that the corresponding ALIAS record is managed by ExternalDNS. This makes ExternalDNS safe for running in environments where there are other records managed via other means.

Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.

```console
$ dig +short @ns-5514.awsdns-53.org. nginx.external-dns-test.my-org.com.
ae11c2360188411e7951602725593fd1-1224345803.eu-central-1.elb.amazonaws.com.
```

If you hooked up your DNS zone with its parent zone correctly you can use `curl` to access your site.

```console
$ curl nginx.external-dns-test.my-org.com.
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
...
</head>
<body>
...
</body>
</html>
```

Ingress objects on AWS require a separately deployed Ingress controller which we'll describe in another tutorial.

## Custom TTL

The default DNS record TTL (Time-To-Live) is 300 seconds. You can customize this value by setting the annotation `external-dns.alpha.kubernetes.io/ttl`.
e.g., modify the service manifest YAML file above:

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

This will set the DNS record's TTL to 60 seconds.

## Routing policies

Route53 offers [different routing policies](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html). The routing policy for a record can be controlled with the following annotations:

* `external-dns.alpha.kubernetes.io/set-identifier`: this **needs** to be set to use any of the following routing policies

For any given DNS name, only **one** of the following routing policies can be used:

* Weighted records: `external-dns.alpha.kubernetes.io/aws-weight`
* Latency-based routing: `external-dns.alpha.kubernetes.io/aws-region`
* Failover:`external-dns.alpha.kubernetes.io/aws-failover`
* Geolocation-based routing:
  * `external-dns.alpha.kubernetes.io/aws-geolocation-continent-code`
  * `external-dns.alpha.kubernetes.io/aws-geolocation-country-code`
  * `external-dns.alpha.kubernetes.io/aws-geolocation-subdivision-code`
* Multi-value answer:`external-dns.alpha.kubernetes.io/aws-multi-value-answer`

## Clean up

Make sure to delete all Service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the hosted zone if you created one for the testing purpose.

```console
$ aws route53 delete-hosted-zone --id /hostedzone/ZEWFWZ4R16P7IB
```
