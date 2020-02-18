# Setting up ExternalDNS for Services on Alibaba Cloud

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on Alibaba Cloud. Make sure to use **>=0.5.6** version of ExternalDNS for this tutorial

## RAM Permissions

```json
{
  "Version": "1",
  "Statement": [
    {
      "Action": "alidns:AddDomainRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "alidns:DeleteDomainRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "alidns:UpdateDomainRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "alidns:DescribeDomainRecords",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "alidns:DescribeDomains",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:AddZoneRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:DeleteZoneRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:UpdateZoneRecord",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:DescribeZoneRecords",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:DescribeZones",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "pvtz:DescribeZoneInfo",
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
```

When running on Alibaba Cloud, you need to make sure that your nodes (on which External DNS runs) have the RAM instance profile with the above RAM role assigned.

## Set up a Alibaba Cloud DNS service or Private Zone service

Alibaba Cloud DNS Service is the domain name resolution and management service for public access. It routes access from end-users to the designated web app.
Alibaba Cloud Private Zone is the domain name resolution and management service for VPC internal access. 

*If you prefer to try-out ExternalDNS in one of the existing domain or zone you can skip this step*

Create a DNS domain which will contain the managed DNS records. For public DNS service, the domain name should be valid and owned by yourself.

```console
$ aliyun alidns AddDomain --DomainName "external-dns-test.com"
```


Make a note of the ID of the hosted zone you just created.

```console
$ aliyun alidns DescribeDomains --KeyWord="external-dns-test.com" | jq -r '.Domains.Domain[0].DomainId'
```

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

### Manifest (for clusters without RBAC enabled)
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: external-dns
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
        - --domain-filter=external-dns-test.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=alibabacloud
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --alibaba-cloud-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
        volumeMounts:
        - mountPath: /usr/share/zoneinfo
          name: hostpath
      volumes:
      - name: hostpath
        hostPath:
          path: /usr/share/zoneinfo
          type: Directory
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
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"] 
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
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
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: external-dns
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
        - --domain-filter=external-dns-test.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=alibabacloud
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --alibaba-cloud-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-identifier
        - --alibaba-cloud-config-file= # enable sts token 
        volumeMounts:
        - mountPath: /usr/share/zoneinfo
          name: hostpath
      volumes:
      - name: hostpath
        hostPath:
          path: /usr/share/zoneinfo
          type: Directory
```



## Arguments

This list is not the full list, but a few arguments that where chosen.

### alibaba-cloud-zone-type

`alibaba-cloud-zone-type` allows filtering for private and public zones

* If value is `public`, it will sync with records in Alibaba Cloud DNS Service
* If value is `private`, it will sync with records in Alibaba Cloud Private Zone Service


## Verify ExternalDNS works (Ingress example)

Create an ingress resource manifest file.

> For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object.

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: foo
  annotations:
    kubernetes.io/ingress.class: "nginx" # use the one that corresponds to your ingress controller.
spec:
  rules:
  - host: foo.external-dns-test.com
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
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.com.
spec:
  type: LoadBalancer
  ports:
  - port: 80
    name: http
    targetPort: 80
  selector:
    app: nginx

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  selector:
    matchLabels:
      app: nginx
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
$ aliyun alidns DescribeDomainRecords --DomainName=external-dns-test.com
{
  "PageNumber": 1,
  "TotalCount": 1,
  "PageSize": 20,
  "RequestId": "1DBEF426-F771-46C7-9802-4989E9C94EE8",
  "DomainRecords": {
    "Record": [
      {
        "RR": "nginx",
        "Status": "ENABLE",
        "Value": "1.2.3.4",
        "Weight": 1,
        "RecordId": "3994015629411328",
        "Type": "A",
        "DomainName": "external-dns-test.com",
        "Locked": false,
        "Line": "default",
        "TTL": 600
      }，
      {
        "RR": "nginx",
        "Status": "ENABLE",
        "Value": "heritage=external-dns;external-dns/owner=my-identifier",
        "Weight": 1,
        "RecordId": "3994015629411329",
        "Type": "TTL",
        "DomainName": "external-dns-test.com",
        "Locked": false,
        "Line": "default",
        "TTL": 600
      }      
    ]
  }
}
```

Note created TXT record alongside ALIAS record. TXT record signifies that the corresponding ALIAS record is managed by ExternalDNS. This makes ExternalDNS safe for running in environments where there are other records managed via other means.

Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.

```console
$ dig nginx.external-dns-test.com.
```

If you hooked up your DNS zone with its parent zone correctly you can use `curl` to access your site.

```console
$ curl nginx.external-dns-test.com.
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

## Custom TTL

The default DNS record TTL (Time-To-Live) is 300 seconds. You can customize this value by setting the annotation `external-dns.alpha.kubernetes.io/ttl`.
e.g., modify the service manifest YAML file above:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.com
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
$ aliyun alidns DeleteDomain --DomainName external-dns-test.com
```

For more info about Alibaba Cloud external dns, please refer this [docs](https://yq.aliyun.com/articles/633412)
