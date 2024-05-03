# Setting up ExternalDNS for Services on ArvanCloud

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using ArvanCloud DNS.

## Creating ArvanCloud Credentials

After login in ArvanCloud website, doing below steps:

1. Go to **Profile Settings** > **Workspace Management** > **Machine User** and press button `Create New Machine User` (Link: https://panel.arvancloud.ir/profile/machine-user/new)
   * In the new page, fill `Machine User Name` (Also you can fill other fields for adding more described)
   * Copy **Access Key** after creating the user
2. Go to **Profile Settings** > **Workspace Management** > **Resources** and press button `Create New Resource Group` (Link: https://panel.arvancloud.ir/profile/resource-group-create)
   * In the new page, fill `Resource Group Name`
   * After set new name, select `Resource Type` contains your domains
3. Go to **Profile Settings** > **Workspace Management** > **Access Policies** and press button `Create New Access Policies` (Link: https://panel.arvancloud.ir/profile/add-policy)
   * In the new page, according to your requirements, select one of **Workspace Policy** or **Resource Group Policy**
   * From the list of resources, pick the resource you was created in the previous step (step 2)
   * Fill `Policy Name` and `Policy Description`
   * Select members from the list (this user was created in step 1)
   * Then in **Select Role for Policy**, check `DNS administrator` checkbox

### Add new Access Keys to exist machine user

If you have created the machine user and this user exists. Go to **Profile Settings** > **Workspace Management** > **Machine User** and select your user has been created. Then in **Machine User Details** page, press button `New Access Key` and copy your api access key


## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.

Begin by creating a Kubernetes secret to securely store your ArvanCloud API key. This key will enable ExternalDNS to authenticate with ArvanCloud:

```shell
kubectl create secret generic arvancloude-api-key --from-literal=API_KEY="apikey ********-****-****-****-************"
```

Then apply one of the following manifests file to deploy ExternalDNS.

### Using Helm

Create a values.yaml file to configure ExternalDNS to use ArvanCloud as the DNS provider. This file should include the necessary environment variables:

```shell
provider: 
  name: ArvanCloud
env:
  - name: AC_API_TOKEN
    valueFrom:
      secretKeyRef:
        name: arvancloude-api-key
        key: API_KEY
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
         containers:
            - name: external-dns
              image: registry.k8s.io/external-dns/external-dns:v0.14.1
              args:
                 - --source=service # ingress is also possible
                 - --domain-filter=example.com # (optional) limit to only example.com domains
                 - --provider=arvancloud
                 - --arvancloud-proxied # (optional) enable the proxy feature of ArvanCloud
                 - --arvancloud-zone-records-per-page=500 # (optional) configure how many DNS records to fetch per request
              env:
                 - name: AC_API_TOKEN
                   valueFrom:
                      secretKeyRef:
                         name: arvancloude-api-key
                         key: API_KEY
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
     verbs: ["list", "watch"]
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
           image: registry.k8s.io/external-dns/external-dns:v0.14.1
           args:
              - --source=service # ingress is also possible
              - --domain-filter=example.com # (optional) limit to only example.com domains
              - --provider=arvancloud
              - --arvancloud-proxied # (optional) enable the proxy feature of ArvanCloud
              - --arvancloud-zone-records-per-page=500 # (optional) configure how many DNS records to fetch per request
           env:
              - name: AC_API_TOKEN
                valueFrom:
                   secretKeyRef:
                      name: arvancloude-api-key
                      key: API_KEY
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
    external-dns.alpha.kubernetes.io/hostname: example.com
    external-dns.alpha.kubernetes.io/ttl: "120" #optional
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the ArvanCloud DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

By setting the TTL annotation on the service, you have to pass a valid TTL, which must be 120 or above.
This annotation is optional, if you won't set it, it will be 1 (automatic) which is 300.
For ArvanCloud proxied entries, set the TTL annotation to 1 (automatic), or do not set it.

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the ArvanCloud DNS records.