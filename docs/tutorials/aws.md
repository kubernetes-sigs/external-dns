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

## Set up a hosted zone

*If you prefer to try-out ExternalDNS in one of the existing hosted-zones you can skip this step*

Create a DNS zone which will contain the managed DNS records.

```console
$ aws route53 create-hosted-zone --name "external-dns-test.my-org.com." --caller-reference "external-dns-test-$(date +%s)"
```


Make a note of the ID of the hosted zone you just created.

```console
$ aws route53 list-hosted-zones-by-name --dns-name "external-dns-test.my-org.com." | jq -r '.HostedZones[0].Id'
/hostedzone/ZEWFWZ4R16P7IB
```

Make a note of the nameservers that were assigned to your new zone.

```console
$ aws route53 list-resource-record-sets --hosted-zone-id "/hostedzone/ZEWFWZ4R16P7IB" \
    --query "ResourceRecordSets[?Type == 'NS']" | jq -r '.[0].ResourceRecords[].Value'
ns-5514.awsdns-53.org.
...
```

In this case it's the ones shown above but your's will differ.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
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
        image: registry.opensource.zalan.do/teapot/external-dns:v0.4.8
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
```

## Arguments

This list is not the full list, but a few arguments that where chosen.

### aws-zone-type

`aws-zone-type` allows filtering for private and public zones


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

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com.
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
$ aws route53 list-resource-record-sets --hosted-zone-id "/hostedzone/ZEWFWZ4R16P7IB" \
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
              "Value": "\"heritage=external-dns,external-dns/owner=my-identifier\""
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
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com.
    external-dns.alpha.kubernetes.io/ttl: 60
spec:
    ...
```

This will set the DNS record's TTL to 60 seconds.

## Clean up

Make sure to delete all Service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the hosted zone if you created one for the testing purpose.

```console
$ aws route53 delete-hosted-zone --id /hostedzone/ZEWFWZ4R16P7IB
```
