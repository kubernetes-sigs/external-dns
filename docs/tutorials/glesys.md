# Setting up ExternalDNS for Services on GleSYS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using GleSYS DNS.

Make sure to use **?????** version of ExternalDNS for this tutorial.


## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

### Manifest
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: glesys-external-dns
stringData:
  project: "cl12345"
  accesskey: "???????????????????????????"
---
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
#  namespace: kube-system
spec:
  replicas: 1
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
      tolerations:
        - key: "node-role.kubernetes.io/master"
          effect: NoSchedule
        - key: "node-role.kubernetes.io/control-plane"
          effect: NoSchedule
        - key: "node.kubernetes.io/not-ready"
          effect: NoSchedule
      containers:
      - name: external-dns
        image: source.svenils.se:5050/kubernetes/external-dns:46b526f2-dirty
        args:
        - --log-level=debug
        - --source=service
        - --source=ingress
        - --provider=glesys
        env:
          - name: GLESYS_PROJECT
            valueFrom:
              secretKeyRef:
                name: glesys-external-dns
                key: project
          - name: GLESYS_ACCESS_KEY
            valueFrom:
              secretKeyRef:
                name: glesys-external-dns
                key: accesskey


```


## Deploying an Nginx Service

Create a service file called 'nginx.yaml' with the following contents:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
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
    external-dns.alpha.kubernetes.io/internal-hostname: nginx.internal.exampleab.com
    external-dns.alpha.kubernetes.io/hostname: nginx.exampleab.com

spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the GleSYS DNS zone created above.

ExternalDNS uses this annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```console
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize the GleSYS DNS records.

## Verifying GleSYS DNS records

Check your [GleSYS Cloud UI](https://cloud.glesys.com/) to view the records for your GleSYS DNS zone.

Click on the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.
