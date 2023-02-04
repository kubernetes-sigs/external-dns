# Setting up ExternalDNS for AdguardHome

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using AdguardHome.

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
  replicas: 1
  selector:
    matchLabels:
      app: external-dns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.13.2
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=adguardhome
        env:
        - name: ADGUARD_HOME_URL
          value: "YOUR_ADGUARD_HOME_URL" # Note: URL should be in the format of http://adguard.home:3000/control/
        - name: ADGUARD_HOME_PASS
          value: "YOUR_ADGUARD_HOME_PASSWORD"
        - name: ADGUARD_HOME_USER
          value: "YOUR_ADGUARD_HOME_USER"

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
  replicas: 1
  selector:
    matchLabels:
      app: external-dns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      containers:
        - name: external-dns
          image: registry.k8s.io/external-dns/external-dns:v0.13.2
          args:
            - --source=service # ingress is also possible
            - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
            - --provider=adguardhome
          env:
            - name: ADGUARD_HOME_URL
              value: "YOUR_ADGUARD_HOME_URL" # Note: URL should be in the format of http://adguard.home:3000/control/
            - name: ADGUARD_HOME_PASS
              value: "YOUR_ADGUARD_HOME_PASSWORD"
            - name: ADGUARD_HOME_USER
              value: "YOUR_ADGUARD_HOME_USER"
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

ExternalDNS uses `external-dns.alpha.kubernetes.io/hostname` annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```console
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize the AdguardHome DNS records.

## Verifying AdguardHome DNS records

Check your AdguardHome DNS records to see if the new record was created.

Click on the zone for the one created above if a different domain was used.

Click on "Filters" and then "Custom filtering rules"

This should show the external IP address of the service as the "Answer" for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage AdguardHome DNS records, we can delete the tutorial's example:

```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
```
