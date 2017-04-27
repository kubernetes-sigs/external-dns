# Setting up ExternalDNS for Services on AWS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on AWS.

Create a DNS zone which will contain the managed DNS records.

```console
$ aws route53 create-hosted-zone --name "external-dns-test.teapot.zalan.do." --caller-reference "external-dns-test-$(date +%s)"
```

Make a note of the ID of the hosted zone you just created.

```console
$ aws route53 list-hosted-zones-by-name --dns-name "external-dns-test.teapot.zalan.do." | jq -r '.HostedZones[0].Id'
/hostedzone/Z16P7IEWFWZ4RB
```

Make a note of the nameservers that were assigned to your new zone.

```console
$ aws route53 list-resource-record-sets --hosted-zone-id "/hostedzone/Z16P7IEWFWZ4RB" \
    --query "ResourceRecordSets[?Type == 'NS']" | jq -r '.[0].ResourceRecords[].Value'
ns-1455.awsdns-53.org.
ns-1694.awsdns-19.co.uk.
ns-764.awsdns-31.net.
ns-62.awsdns-07.com.
```

In this case it's the ones shown above but your's will differ.

If you decide not to create a new zone but reuse an existing one, make sure it's currently **unused** and **empty**. This version of ExternalDNS will remove all records it doesn't recognize from the zone.

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
        image: registry.opensource.zalan.do/teapot/external-dns:v0.2.1
        args:
        - --in-cluster
        - --zone=Z16P7IEWFWZ4RB
        - --source=service
        - --provider=aws
        - --dry-run=false
```

Create the following sample application to test that ExternalDNS works.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.teapot.zalan.do.
spec:
  type: LoadBalancer
  ports:
  - port: 80
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
```

After roughly two minutes check that a corresponding DNS record for your service was created.

```console
$ aws route53 list-resource-record-sets --hosted-zone-id "/hostedzone/Z16P7IEWFWZ4RB" \
    --query "ResourceRecordSets[?Name == 'nginx.external-dns-test.teapot.zalan.do.']|[?Type == 'CNAME']"
[
    {
        "ResourceRecords": [
            {
                "Value": "ae11c2360188411e7951602725593fd1-1224345803.eu-central-1.elb.amazonaws.com"
            }
        ],
        "Type": "CNAME",
        "Name": "nginx.external-dns-test.teapot.zalan.do.",
        "TTL": 300
    }
]
```

Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.

```console
$ dig +short @ns-1455.awsdns-53.org. nginx.external-dns-test.teapot.zalan.do.
ae11c2360188411e7951602725593fd1-1224345803.eu-central-1.elb.amazonaws.com.
```

If you hooked up your DNS zone with its parent zone correctly you can use `curl` to access your site.

```console
$ curl nginx.external-dns-test.teapot.zalan.do.
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

## Clean up

Make sure to delete all Service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the hosted zone.

```console
$ aws route53 delete-hosted-zone --id /hostedzone/Z16P7IEWFWZ4RB
```
