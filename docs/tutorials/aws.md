# Setting up ExternalDNS for Services on AWS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on AWS. Make sure to use **>=0.4** version of ExternalDNS for this tutorial

## IAM Policy

The following IAM Policy document allows ExternalDNS to update Route53 Resource
Record Sets and Hosted Zones. You'll want to create this Policy in IAM first. In
our example, we'll call the policy AllowExternalDNSUpdates (but you can call
it whatever you prefer).

If you prefer, you may fine-tune the policy to permit updates only to explicit
Hosted Zone IDs.

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

## Create IAM Role

You'll need to create an IAM Role that can be assumed by the ExternalDNS Pod.
Note the role name; you'll need to refer to it in the K8S manifest below.

Attach the AllowExternalDNSUpdates IAM Policy (above) to the role.

The trust relationship associated with the IAM Role will vary depending on how
you've configured your Kubernetes cluster:

### Amazon EKS

If your EKS-managed cluster is >= 1.13 and was created after 2019-09-04, refer
to the [Amazon EKS
documentation](https://docs.aws.amazon.com/eks/latest/userguide/create-service-account-iam-policy-and-role.html)
for instructions on how to create the IAM Role. Otherwise, you will need to use
kiam or kube2iam.

### kiam

If you're using [kiam](https://github.com/uswitch/kiam), follow the
[instructions](https://github.com/uswitch/kiam/blob/HEAD/docs/IAM.md) for
creating the IAM role.

### kube2iam

If you're using [kube2iam](https://github.com/jtblin/kube2iam), follow the
instructions for creating the IAM Role.

### EC2 Instance Role (not recommended)

**:warning: WARNING: This will grant all pods on the node the ability to
manipulate Route 53 Resource Record Sets. If exploited by an attacker, this
could lead to a serious security and/or availability incident. For this reason,
it is not recommended.**

Create an IAM Role for your EC2 instances as described in the [Amazon EC2
documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html).
Then, attach the associated Instance Profile to the EC2 instances that comprise
your K8S cluster.

For this method to work, you must permit your pods the ability to access the EC2
instance metadata service (169.254.169.254). This is allowed by default.

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

For clusters with RBAC enabled, be sure to choose the correct `namespace`.

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
      # If you're using kiam or kube2iam, specify the following annotation.
      # Otherwise, you may safely omit it.
      annotations:
        iam.amazonaws.com/role: arn:aws:iam::ACCOUNT-ID:role/IAM-SERVICE-ROLE-NAME
    spec:
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
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
  # If you're using Amazon EKS with IAM Roles for Service Accounts, specify the following annotation.
  # Otherwise, you may safely omit it.
  annotations:
    # Substitute your account ID and IAM service role name below.
    eks.amazonaws.com/role-arn: arn:aws:iam::ACCOUNT-ID:role/IAM-SERVICE-ROLE-NAME
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
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
      # If you're using kiam or kube2iam, specify the following annotation.
      # Otherwise, you may safely omit it.
      annotations:
        iam.amazonaws.com/role: arn:aws:iam::ACCOUNT-ID:role/IAM-SERVICE-ROLE-NAME
    spec:
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=aws
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --aws-zone-type=public # only look at public hosted zones (valid values are public, private or no value for both)
        - --registry=txt
        - --txt-owner-id=my-hostedzone-identifier
      securityContext:
        fsGroup: 65534 # For ExternalDNS to be able to read Kubernetes and AWS token files
```

## Arguments

This list is not the full list, but a few arguments that where chosen.

### aws-zone-type

`aws-zone-type` allows filtering for private and public zones

## Annotations

Annotations which are specific to AWS.

### alias

`external-dns.alpha.kubernetes.io/alias` if set to `true` on an ingress, it will create an ALIAS record when the target is an ALIAS as well. To make the target an alias, the ingress needs to be configured correctly as described in [the docs](./nginx-ingress.md#with-a-separate-tcp-load-balancer). In particular, the argument `--publish-service=default/nginx-ingress-controller` has to be set on the `nginx-ingress-controller` container. If one uses the `nginx-ingress` Helm chart, this flag can be set with the `controller.publishService.enabled` configuration option.

## Verify ExternalDNS works (Ingress example)

Create an ingress resource manifest file.

> For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object.

```yaml
apiVersion: networking.k8s.io/v1
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

## Associating DNS records with healthchecks

You can configure Route53 to associate DNS records with healthchecks for automated DNS failover using 
`external-dns.alpha.kubernetes.io/aws-health-check-id: <health-check-id>` annotation.

Note: ExternalDNS does not support creating healthchecks, and assumes that `<health-check-id>` already exists.

## Govcloud caveats

Due to the special nature with how Route53 runs in Govcloud, there are a few tweaks in the deployment settings.

* An Environment variable with name of AWS_REGION set to either us-gov-west-1 or us-gov-east-1 is required. Otherwise it tries to lookup a region that does not exist in Govcloud and it errors out.

```yaml
env:
- name: AWS_REGION
  value: us-gov-west-1
```

* Route53 in Govcloud does not allow aliases. Therefore, container args must be set so that it uses CNAMES and a txt-prefix must be set to something. Otherwise, it will try to create a TXT record with the same value than the CNAME itself, which is not allowed.

```yaml
args:
- --aws-prefer-cname
- --txt-prefix={{ YOUR_PREFIX }}
```

* The first two changes are needed if you use Route53 in Govcloud, which only supports private zones. There are also no cross account IAM whatsoever between Govcloud and commerical AWS accounts. If services and ingresses need to make Route 53 entries to an public zone in a commerical account, you will have set env variables of AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY with a key and secret to the commerical account that has the sufficient rights.

```yaml
env:
- name: AWS_ACCESS_KEY_ID
  value: XXXXXXXXX
- name: AWS_SECRET_ACCESS_KEY
  valueFrom:
    secretKeyRef:
      name: {{ YOUR_SECRET_NAME }}
      key: {{ YOUR_SECRET_KEY }}
```

## Clean up

Make sure to delete all Service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the hosted zone if you created one for the testing purpose.

```console
$ aws route53 delete-hosted-zone --id /hostedzone/ZEWFWZ4R16P7IB
```

## Throttling

Route53 has a [5 API requests per second per account hard quota](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html#limits-api-requests-route-53).
Running several fast polling ExternalDNS instances in a given account can easily hit that limit. Some ways to circumvent that issue includes:
* Augment the synchronization interval (`--interval`), at the cost of slower changes propagation.
* If the ExternalDNS managed zones list doesn't change frequently, set `--aws-zones-cache-duration` (zones list cache time-to-live) to a larger value. Note that zones list cache can be disabled with `--aws-zones-cache-duration=0s`.
