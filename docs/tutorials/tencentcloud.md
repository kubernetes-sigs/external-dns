# Setting up ExternalDNS for Tencent Cloud

## External Dns Version
* Make sure to use **>=1.7.2** version of ExternalDNS for this tutorial

## Set up PrivateDns or DNSPod

Tencent Cloud DNSPod Service is the domain name resolution and management service for public access.
Tencent Cloud PrivateDNS Service is the domain name resolution and management service for VPC internal access.

* If you want to use internal dns service in Tencent Cloud. 
1. Set up the args `--tencent-cloud-zone-type=private`   
2. Create a DNS domain in PrivateDNS console. DNS domain which will contain the managed DNS records.

* If you want to use public dns service in Tencent Cloud.
1. Set up the args `--tencent-cloud-zone-type=public`   
2. Create a Domain in DnsPod console. DNS domain which will contain the managed DNS records.

## Set up CAM for API Key

In Tencent CAM Console. you may get the secretId and secretKey pair. make sure the key pair has those Policy.
```json
{
    "version": "2.0",
    "statement": [
        {
            "effect": "allow",
            "action": [
                "dnspod:ModifyRecord",
                "dnspod:DeleteRecord",
                "dnspod:CreateRecord",
                "dnspod:DescribeRecordList",
                "dnspod:DescribeDomainList"
            ],
            "resource": [
                "*"
            ]
        },
        {
            "effect": "allow",
            "action": [
                "privatedns:DescribePrivateZoneList",
                "privatedns:DescribePrivateZoneRecordList",
                "privatedns:CreatePrivateZoneRecord",
                "privatedns:DeletePrivateZoneRecord",
                "privatedns:ModifyPrivateZoneRecord"
            ],
            "resource": [
                "*"
            ]
        }
    ]
}
```

# Deploy ExternalDNS

## Manifest (for clusters with RBAC enabled)

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
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
  verbs: ["list"]
---
apiVersion: rbac.authorization.k8s.io/v1
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
apiVersion: v1
kind: ConfigMap
metadata:
  name: external-dns
data:
  tencent-cloud.json: |
    {
      "regionId": "ap-shanghai",
      "secretId": "******",  
      "secretKey": "******",
      "vpcId": "vpc-******"
    }
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
      containers:
      - args:
        - --source=service
        - --source=ingress
        - --domain-filter=external-dns-test.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=tencentcloud
        - --policy=sync # set `upsert-only` would prevent ExternalDNS from deleting any records
        - --tencent-cloud-zone-type=private # only look at private hosted zones. set `public` to use the public dns service.
        - --tencent-cloud-config-file=/etc/kubernetes/tencent-cloud.json
        image: k8s.gcr.io/external-dns/external-dns:v1.7.2
        imagePullPolicy: Always
        name: external-dns
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/kubernetes
          name: config-volume
          readOnly: true
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: external-dns
      serviceAccountName: external-dns
      terminationGracePeriodSeconds: 30
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: tencent-cloud.json
            path: tencent-cloud.json
          name: external-dns
        name: config-volume
```

# Example

## Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.com
    external-dns.alpha.kubernetes.io/internal-hostname: nginx-internal.external-dns-test.com
    external-dns.alpha.kubernetes.io/ttl: "600"
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

`nginx.external-dns-test.com` will record to the Loadbalancer VIP.
`nginx-internal.external-dns-test.com` will record to the ClusterIP.
all of the DNS Record ttl will be 600.

# Attention

This makes ExternalDNS safe for running in environments where there are other records managed via other means.

