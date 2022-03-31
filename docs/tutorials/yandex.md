# Setting up ExternalDNS for Services on Yandex Cloud

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on Yandex Cloud.

Make sure to use **> 1.7.1** version of ExternalDNS for this tutorial.

This tutorial uses [Yandex Cloud CLI 0.88.0](https://cloud.yandex.com/en-ru/docs/cli/quickstart) for all
Yandex Cloud commands and assumes that the Kubernetes cluster was created via Yandex Managed Service for Kubernetes and `kubectl` commands
are being run on an orchestration node.

## Creating an Yandex Cloud DNS zone

The Yandex provider for ExternalDNS will find suitable zones for domains it manages; it will
not automatically create zones.

For this tutorial, we will create a Yandex folder named 'externaldns' that can easily be deleted later:

```
$ yc resource-manager folder create \
  --name externaldns
```

Next, create a Yandex DNS zone for "example.com":

```
$ yc dns zone create \
  --name externaldns-examplezone \
  --zone example.com. \
  --folder-name externaldns
```

Substitute a domain you own for "example.com" if desired.

If using your own domain that was registered with a third-party domain registrar, you should point your domain's
name servers to the values in the `NS` record in created dns zone.
Please consult your registrar's documentation on how to do that.

## Permissions to modify DNS zone

External-DNS needs permissions to make changes in the Yandex DNS server. 
These permissions should be granted to Service Account, that External DNS will be used to authorize API requests.

### Create Service Account

There are two options: create service account for External DNS or use service account that associated with any node in kubernetes cluster (service account for nodes).
In this tutorial, we will use dedicated service account, created with following command:

```
$ yc iam service-account create \
  --name externaldns-example-sa \
  --folder-name externaldns
```

### Add access binding

Access binding is the association between role and subject in Yandex Cloud.
A service account with a minimum role `dns.editor` required to manage DNS zone.
You can find service account id with following command:

```
$ yc iam service-account get \
  --name externaldns-example-sa \
  --folder-name externaldns \
  --format json | jq '.id'
```

After that, use service account id to create access binding:

```
$ yc resource-manager folder add-access-binding \
  externaldns \
  --role dns.editor \
  --subject serviceAccount:<service-account-id>
```

### Create Service Account Authorized Key

Create service account [Authorized Key](https://cloud.yandex.com/en-ru/docs/iam/concepts/authorization/key):

```
$ yc iam key create \
  --service-account-name externaldns-example-sa \
  --output authorized-key.json \
  --folder-name externaldns
```

### Create OAuth token

OAuth token can be obtained from [here](https://cloud.yandex.com/en-ru/docs/iam/concepts/authorization/oauth-token)

## Deploy External DNS

This deployment assumes that you will be using nginx-ingress. 
When using nginx-ingress do not deploy it as a Daemon Set. 
This causes nginx-ingress to write the Cluster IP of the backend pods in the ingress status.loadbalancer.ip property which then has external-dns write the Cluster IP(s) in DNS vs. the nginx-ingress service external IP.

Ensure that your nginx-ingress deployment has the following arg: added to it:

```
- --publish-service=namespace/nginx-ingress-controller-svcname
```

For more details see here: [nginx-ingress external-dns](https://github.com/kubernetes-sigs/external-dns/blob/HEAD/docs/faq.md#why-is-externaldns-only-adding-a-single-ip-address-in-route-53-on-aws-when-using-the-nginx-ingress-controller-how-do-i-get-it-to-use-the-fqdn-of-the-elb-assigned-to-my-nginx-ingress-controller-service-instead)

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

### Manifest with Service Account Authorized Key

Create K8S secret from authorized key:

```
$ kubectl create secret generic \
  externaldns-example-sa-auth-key \
  --from-file=key=authorized-key.json
```

Retrieve folder id:

```
$ yc resource-manager folder get \
  --name externaldns \
  --format json
```

Apply manifest:

**Note:** You need to set folder id as value of `--yandex-folder-id` argument

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
        image: k8s.gcr.io/external-dns/external-dns:v1.7.2
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=yandex
        - --yandex-authorization-type=iam-key-file
        - --yandex-authorization-key-file=/var/lib/yandex/auth.json
        - --yandex-folder-id=<insert here the folder id>
        volumeMounts:
        - name: yandex-authorization-key-file
          mountPath: /var/lib/yandex/auth.json
          readOnly: true
          subPath: key
      serviceAccountName: external-dns
      volumes:
      - name: yandex-authorization-key-file
        secret:
          secretName: externaldns-example-sa-auth-key
```

Create the deployment for ExternalDNS:

```
$ kubectl create -f externaldns.yaml
```


### Manifest with OAuth token

Create K8S secret from API Key:

```
$ kubectl create secret generic \
  externaldns-example-sa-oauth-token \
  --from-literal=oauth-token=<insert here obtained token>
```

Retrieve folder id:

```
$ yc resource-manager folder get \
  --name externaldns \
  --format json
```

Apply manifest:

**Note:** You need to set folder id as value of `--yandex-folder-id` argument

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
        image: k8s.gcr.io/external-dns/external-dns:v1.7.2
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=yandex
        - --yandex-authorization-type=iam-token
        - --yandex-folder-id=<insert here the folder id>
        env:
          - name: EXTERNAL_DNS_YANDEX_AUTHORIZATION_OAUTH_TOKEN
            valueFrom:
              secretKeyRef:
                key: oauth-token
                name: externaldns-example-sa-oauth-token
      serviceAccountName: external-dns
```

Create the deployment for ExternalDNS:

```
$ kubectl create -f externaldns.yaml
```

### Manifest with Instance Service Account

This configuration is the same as above, except it required that access binding should be added to node service account.
Also, in this mode, no secrets or tokens required, because API requests will be authorized from [service account](https://cloud.yandex.com/en-ru/docs/managed-kubernetes/operations/kubernetes-cluster/kubernetes-cluster-create#kubernetes-cluster-create) that associated to instance group

Retrieve folder id:

```
$ yc resource-manager folder get \
  --name externaldns \
  --format json
```

Apply manifest:

**Note:** You need to set folder id as value of `--yandex-folder-id` argument

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
        image: k8s.gcr.io/external-dns/external-dns:v1.7.2
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=yandex
        - --yandex-authorization-type=instance-service-account
        - --yandex-folder-id=<insert here the folder id>
      serviceAccountName: external-dns
```

Create the deployment for ExternalDNS:

```
$ kubectl create -f externaldns.yaml
```

## Deploying Nginx Service

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
  name: nginx-svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx
  type: ClusterIP

---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx
  annotations:
    kubernetes.io/ingress.class: nginx
spec:
  rules:
  - host: server.example.com
    http:
      paths:
      - backend:
          service:
            name: nginx-svc
            port: 
              number: 80
        path: /
        pathType: ImplementationSpecific
```

When using external-dns with ingress objects it will automatically create DNS records based on host names specified in ingress objects that match the domain-filter argument in the external-dns deployment manifest. 
When those host names are removed or renamed the corresponding DNS records are also altered.

Create the deployment, service and ingress object:

```
$ kubectl create -f nginx.yaml
```

Since your external IP would have already been assigned to the nginx-ingress service, the DNS records pointing to the IP of the nginx-ingress service should be created within a minute.

## Verifying Yandex DNS records

Run the following command to view the A records for your Yandex DNS zone:

```
$ yc dns zone list-records \
  --name externaldns-examplezone \
  --folder-name externaldns
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.

## Delete Yandex Folder

Now that we have verified that ExternalDNS will automatically manage Yandex DNS records, we can delete the tutorial's
folder:

```
$ yc resource-manager folder delete \
  --name externaldns
```
