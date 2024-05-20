# Setting up External-DNS for Services on Akamai Edge DNS

## Prerequisites

External-DNS v0.8.0 or greater.

### Zones

External-DNS manages service endpoints in existing DNS zones. The Akamai provider does not add, remove or configure new zones. The [Akamai Control Center](https://control.akamai.com) or [Akamai DevOps Tools](https://developer.akamai.com/devops), [Akamai CLI](https://developer.akamai.com/cli) and [Akamai Terraform Provider](https://developer.akamai.com/tools/integrations/terraform) can create and manage Edge DNS zones. 

### Akamai Edge DNS Authentication

The Akamai Edge DNS provider requires valid Akamai Edgegrid API authentication credentials to access zones and manage  DNS records. 

Either directly by key or indirectly via a file can set credentials for the provider. The Akamai credential keys and mappings to the Akamai provider utilizing different presentation methods are:

| Edgegrid Auth Key | External-DNS Cmd Line Key | Environment/ConfigMap Key | Description |
| ----------------- | ------------------------- | ------------------------- | ----------- |
| host | akamai-serviceconsumerdomain | EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN | Akamai Edgegrid API server |
| access_token | akamai-access-token | EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN | Akamai Edgegrid API access token |
| client_token | akamai-client-token  | EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN |Akamai Edgegrid API client token |
| client-secret | akamai-client-secret | EXTERNAL_DNS_AKAMAI_CLIENT_SECRET |Akamai Edgegrid API client secret |

In addition to specifying auth credentials individually, an Akamai Edgegrid .edgerc file convention can set credentials.

| External-DNS Cmd Line | Environment/ConfigMap | Description |
| --------------------- | --------------------- | ----------- |
| akamai-edgerc-path | EXTERNAL_DNS_AKAMAI_EDGERC_PATH | Accessible path to Edgegrid credentials file, e.g /home/test/.edgerc |
| akamai-edgerc-section | EXTERNAL_DNS_AKAMAI_EDGERC_SECTION | Section in Edgegrid credentials file containing credentials |

[Akamai API Authentication](https://developer.akamai.com/getting-started/edgegrid) provides an overview and further information about authorization credentials for API base applications and tools.

## Deploy External-DNS

An operational External-DNS deployment consists of an External-DNS container and service. The following sections demonstrate the ConfigMap objects that would make up an example functional external DNS kubernetes configuration utilizing NGINX as the service.

Connect your `kubectl` client to the External-DNS cluster.

Begin by creating a Kubernetes secret to securely store your  Akamai Edge DNS Access Tokens. This key will enable ExternalDNS to authenticate with Akamai Edge DNS:

```shell
kubectl create secret generic AKAMAI-DNS --from-literal=EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN=YOUR_SERVICECONSUMERDOMAIN --from-literal=EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN=YOUR_CLIENT_TOKEN --from-literal=EXTERNAL_DNS_AKAMAI_CLIENT_SECRET=YOUR_CLIENT_SECRET --from-literal=EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN=YOUR_ACCESS_TOKEN
```

Ensure to replace YOUR_SERVICECONSUMERDOMAIN, EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN, YOUR_CLIENT_SECRET and YOUR_ACCESS_TOKEN with your actual Akamai Edge DNS API keys.

Then apply one of the following manifests file to deploy ExternalDNS.

### Using Helm

Create a values.yaml file to configure ExternalDNS to use Akamai Edge DNS as the DNS provider. This file should include the necessary environment variables:

```shell
provider:
  name: akamai
env:
  - name: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
    valueFrom:
      secretKeyRef:
        name: AKAMAI-DNS
        key: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
  - name: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
    valueFrom:
      secretKeyRef:
        name: AKAMAI-DNS
        key: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
  - name: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
    valueFrom:
      secretKeyRef:
        name: AKAMAI-DNS
        key: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
  - name: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
    valueFrom:
      secretKeyRef:
        name: AKAMAI-DNS
        key: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
```

Finally, install the ExternalDNS chart with Helm using the configuration specified in your values.yaml file:

```shell
helm upgrade --install external-dns external-dns/external-dns --values values.yaml
```


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
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=service  # or ingress or both
        - --provider=akamai
        - --domain-filter=example.com
        # zone-id-filter may be specified as well to filter on contract ID
        - --registry=txt
        - --txt-owner-id={{ owner-id-for-this-external-dns }}
        - --txt-prefix={{ prefix label for TXT record }}.
        env:
        - name: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
        - name: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
```

### Manifest (for clusters with RBAC enabled)

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
  verbs: ["watch", "list"]
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
        image: registry.k8s.io/external-dns/external-dns:v0.14.2
        args:
        - --source=service  # or ingress or both
        - --provider=akamai
        - --domain-filter=example.com
        # zone-id-filter may be specified as well to filter on contract ID
        - --registry=txt
        - --txt-owner-id={{ owner-id-for-this-external-dns }}
        - --txt-prefix={{ prefix label for TXT record }}.
        env:
        - name: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN
        - name: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_CLIENT_SECRET
        - name: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
          valueFrom:
            secretKeyRef:
              name: AKAMAI-DNS
              key: EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN
```

Create the deployment for External-DNS:

```
$ kubectl apply -f externaldns.yaml
```

## Deploying an Nginx Service

Create a service file called 'nginx.yaml' with the following contents:

```yaml
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
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.example.com
    external-dns.alpha.kubernetes.io/ttl: "600" #optional
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Create the deployment and service object:

```
$ kubectl apply -f nginx.yaml
```

## Verify Akamai Edge DNS Records

Wait 3-5 minutes before validating the records to allow the record changes to propagate to all the Akamai name servers.

Validate records using the [Akamai Control Center](http://control.akamai.com) or by executing a dig, nslookup or similar DNS command.
 
## Cleanup

Once you successfully configure and verify record management via External-DNS, you can delete the tutorial's examples:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```

## Additional Information

* The Akamai provider allows the administrative user to filter zones by both name (`domain-filter`) and contract Id (`zone-id-filter`). The Edge DNS API will return a '500 Internal Error' for invalid contract Ids.
* The provider will substitute quotes in TXT records with a `` ` `` (back tick) when writing records with the API.
