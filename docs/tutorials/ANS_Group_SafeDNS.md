# Setting up ExternalDNS for Services on ANS Group's SafeDNS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using SafeDNS.

Make sure to use **>=0.11.0** version of ExternalDNS for this tutorial.

## Managing DNS with SafeDNS

If you want to learn about how to use the SafeDNS service read the following tutorials:
To learn more about the use of SafeDNS in general, see the following page:

[ANS Group's SafeDNS documentation](https://docs.ukfast.co.uk/domains/safedns/index.html).

## Creating SafeDNS credentials

Generate a fresh API token for use with ExternalDNS, following the instructions
at the ANS Group developer [Getting-Started](https://developers.ukfast.io/getting-started)
page. You will need to grant read/write access to the SafeDNS API. No access to
any other ANS Group service is required.

The environment variable `SAFEDNS_TOKEN` must have a value of this token to run
ExternalDNS with SafeDNS integration.

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
        # You will need to check what the latest version is yourself:
        # https://github.com/kubernetes-sigs/external-dns/releases
        image: k8s.gcr.io/external-dns/external-dns:vX.Y.Z
        args:
        - --source=service # ingress is also possible
        # (optional) limit to only example.com domains; change to match the
        # zone created above.
        - --domain-filter=example.com
        - --provider=safedns
        env:
        - name: SAFEDNS_TOKEN
          value: "SAFEDNSTOKENSAFEDNSTOKEN"
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
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.11.0
        args:
        - --source=service # ingress is also possible
        # (optional) limit to only example.com domains; change to match the
        # zone created above.
        - --domain-filter=example.com
        - --provider=safedns
        env:
        - name: SAFEDNS_TOKEN
          value: "SAFEDNSTOKENSAFEDNSTOKEN"
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
    external-dns.alpha.kubernetes.io/hostname: my-app.example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use a hostname that matches the domain
filter specified above.

ExternalDNS uses this annotation to determine what services should be registered
with DNS. Removing the annotation will cause ExternalDNS to remove the
corresponding DNS records.

Create the deployment and service:

```console
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud
provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new
service IP address and synchronize the SafeDNS records.

## Verifying SafeDNS records

Check your [SafeDNS UI](https://my.ukfast.co.uk/safedns/index.php) and select
the appropriate domain to view the records for your SafeDNS zone.

This should show the external IP address of the service as the A record for your
domain.

Alternatively, you can perform a DNS lookup for the hostname specified:
```console
$ dig +short my-app.example.com
an.ip.addr.ess
```

## Cleanup

Now that we have verified that ExternalDNS will automatically manage SafeDNS
records, we can delete the tutorial's example:

```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
```
