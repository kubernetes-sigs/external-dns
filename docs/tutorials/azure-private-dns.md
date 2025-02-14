# Azure Private DNS

This tutorial describes how to set up ExternalDNS for managing records in Azure Private DNS.

It comprises of the following steps:

1) Provision Azure Private DNS
2) Configure service principal for managing the zone
3) Deploy ExternalDNS
4) Expose an NGINX service with a LoadBalancer and annotate it with the desired DNS name
5) Install NGINX Ingress Controller (Optional)
6) Expose an nginx service with an ingress (Optional)
7) Verify the DNS records

Everything will be deployed on Kubernetes.
Therefore, please see the subsequent prerequisites.

## Prerequisites

- Azure Kubernetes Service is deployed and ready
- [Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) and `kubectl` installed on the box to execute the subsequent steps

## Provision Azure Private DNS

The provider will find suitable zones for domains it manages. It will
not automatically create zones.

For this tutorial, we will create a Azure resource group named 'externaldns' that can easily be deleted later.

```sh
az group create -n externaldns -l westeurope
```

Substitute a more suitable location for the resource group if desired.

As a prerequisite for Azure Private DNS to resolve records is to define links with VNETs.
Thus, first create a VNET.

```sh
$ az network vnet create \
  --name myvnet \
  --resource-group externaldns \
  --location westeurope \
  --address-prefix 10.2.0.0/16 \
  --subnet-name mysubnet \
  --subnet-prefixes 10.2.0.0/24
```

Next, create a Azure Private DNS zone for "example.com":

```sh
az network private-dns zone create -g externaldns -n example.com
```

Substitute a domain you own for "example.com" if desired.

Finally, create the mentioned link with the VNET.

```sh
$ az network private-dns link vnet create -g externaldns -n mylink \
   -z example.com -v myvnet --registration-enabled false
```

## Configure service principal for managing the zone

ExternalDNS needs permissions to make changes in Azure Private DNS.
These permissions are roles assigned to the service principal used by ExternalDNS.

A service principal with a minimum access level of `Private DNS Zone Contributor` to the Private DNS zone(s) and `Reader` to the resource group containing the Azure Private DNS zone(s) is necessary.
More powerful role-assignments like `Owner` or assignments on subscription-level work too.

Start off by **creating the service principal** without role-assignments.

```sh
$ az ad sp create-for-rbac --skip-assignment -n http://externaldns-sp
{
  "appId": "appId GUID",  <-- aadClientId value
  ...
  "password": "password",  <-- aadClientSecret value
  "tenant": "AzureAD Tenant Id"  <-- tenantId value
}
```

> Note: Alternatively, you can issue `az account show --query "tenantId"` to retrieve the id of your AAD Tenant too.

Next, assign the roles to the service principal.
But first **retrieve the ID's** of the objects to assign roles on.

```sh
# find out the resource ids of the resource group where the dns zone is deployed, and the dns zone itself
$ az group show --name externaldns --query id -o tsv
/subscriptions/id/resourceGroups/externaldns

$ az network private-dns zone show --name example.com -g externaldns --query id -o tsv
/subscriptions/.../resourceGroups/externaldns/providers/Microsoft.Network/privateDnsZones/example.com
```

Now, **create role assignments**.

```sh
# 1. as a reader to the resource group
$ az role assignment create --role "Reader" --assignee <appId GUID> --scope <resource group resource id>

# 2. as a contributor to DNS Zone itself
$ az role assignment create --role "Private DNS Zone Contributor" --assignee <appId GUID> --scope <dns zone resource id>
```

## Throttling

When the ExternalDNS managed zones list doesn't change frequently, one can set `--azure-zones-cache-duration` (zones list cache time-to-live). The zones list cache is disabled by default, with a value of 0s.

## Deploy ExternalDNS

Configure `kubectl` to be able to communicate and authenticate with your cluster.
This is per default done through the file `~/.kube/config`.

For general background information on this see [kubernetes-docs](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/).
Azure-CLI features functionality for automatically maintaining this file for AKS-Clusters. See [Azure-Docs](https://docs.microsoft.com/de-de/cli/azure/aks?view=azure-cli-latest#az-aks-get-credentials).

Follow the steps for [azure-dns provider](./azure.md#creating-configuration-file) to create a configuration file.

Then apply one of the following manifests depending on whether you use RBAC or not.

The credentials of the service principal are provided to ExternalDNS as environment-variables.

### Manifest (for clusters without RBAC enabled)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: externaldns
spec:
  selector:
    matchLabels:
      app: externaldns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: externaldns
    spec:
      containers:
      - name: externaldns
        image: registry.k8s.io/external-dns/external-dns:v0.15.1
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        volumeMounts:
        - name: azure-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: azure-config-file
        secret:
          secretName: azure-config-file
```

### Manifest (for clusters with RBAC enabled, cluster access)

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: externaldns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: externaldns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: externaldns-viewer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: externaldns
subjects:
- kind: ServiceAccount
  name: externaldns
  namespace: default
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: externaldns
spec:
  selector:
    matchLabels:
      app: externaldns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: externaldns
    spec:
      serviceAccountName: externaldns
      containers:
      - name: externaldns
        image: registry.k8s.io/external-dns/external-dns:v0.15.1
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        volumeMounts:
        - name: azure-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: azure-config-file
        secret:
          secretName: azure-config-file
```

### Manifest (for clusters with RBAC enabled, namespace access)

This configuration is the same as above, except it only requires privileges for the current namespace, not for the whole cluster.
However, access to [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/) requires cluster access, so when using this manifest,
services with type `NodePort` will be skipped!

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: externaldns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: externaldns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions","networking.k8s.io"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: externaldns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: externaldns
subjects:
- kind: ServiceAccount
  name: externaldns
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: externaldns
spec:
  selector:
    matchLabels:
      app: externaldns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: externaldns
    spec:
      serviceAccountName: externaldns
      containers:
      - name: externaldns
        image: registry.k8s.io/external-dns/external-dns:v0.15.1
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        volumeMounts:
        - name: azure-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: azure-config-file
        secret:
          secretName: azure-config-file
```

Create the deployment for ExternalDNS:

```sh
kubectl create -f externaldns.yaml
```

## Create an nginx deployment

This step creates a demo workload in your cluster. Apply the following manifest to create a deployment that we are going to expose later in this tutorial in multiple ways:

```yaml
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
```

## Expose the nginx deployment with a load balancer

Apply the following manifest to create a service of type `LoadBalancer`. This will create a public load balancer in Azure that will forward traffic to the nginx pods.

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-svc
  annotations:
    service.beta.kubernetes.io/azure-load-balancer-internal: "true"
    external-dns.alpha.kubernetes.io/hostname: server.example.com
    external-dns.alpha.kubernetes.io/internal-hostname: server-clusterip.example.com
spec:
  ports:
    - port: 80
      protocol: TCP
      targetPort: 80
  selector:
    app: nginx
  type: LoadBalancer
```

In the service we used multiple annotations.
The annotation `service.beta.kubernetes.io/azure-load-balancer-internal` is used to create an internal load balancer.
The annotation `external-dns.alpha.kubernetes.io/hostname` is used to create a DNS record for the load balancer that will point to the internal IP address in the VNET allocated by the internal load balancer.
The annotation `external-dns.alpha.kubernetes.io/internal-hostname` is used to create a private DNS record for the load balancer that will point to the cluster IP.

## Install NGINX Ingress Controller (Optional)

Helm is used to deploy the ingress controller.

We employ the popular chart [ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx).

```sh
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm repo update
$ helm install [RELEASE_NAME] ingress-nginx/ingress-nginx
     --set controller.publishService.enabled=true
```

The parameter `controller.publishService.enabled` needs to be set to `true.`

It will make the ingress controller update the endpoint records of ingress-resources to contain the external-ip of the loadbalancer serving the ingress-controller.
This is crucial as ExternalDNS reads those endpoints records when creating DNS-Records from ingress-resources.
In the subsequent parameter we will make use of this. If you don't want to work with ingress-resources in your later use, you can leave the parameter out.

Verify the correct propagation of the loadbalancer's ip by listing the ingresses.

```sh
kubectl get ingress
```

The address column should contain the ip for each ingress. ExternalDNS will pick up exactly this piece of information.

```sh
NAME     HOSTS             ADDRESS          PORTS   AGE
nginx1   sample1.aks.com   52.167.195.110   80      6d22h
nginx2   sample2.aks.com   52.167.195.110   80      6d21h
```

If you do not want to deploy the ingress controller with Helm, ensure to pass the following cmdline-flags to it through the mechanism of your choice:

```sh
flags:
--publish-service=<namespace of ingress-controller >/<svcname of ingress-controller>
--update-status=true (default-value)

example:
./nginx-ingress-controller --publish-service=default/nginx-ingress-controller
```

## Expose the nginx deployment with the ingress (Optional)

Apply the following manifest to create an ingress resource that will expose the nginx deployment. The ingress resource backend points to a `ClusterIP` service that is needed to select the pods that will receive the traffic.

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-svc-clusterip
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
spec:
  ingressClassName: nginx
  rules:
  - host: server.example.com
    http:
      paths:
      - backend:
          service:
            name: nginx-svc-clusterip
            port:
              number: 80
        pathType: Prefix
```

When you use ExternalDNS with Ingress resources, it automatically creates DNS records based on the hostnames listed in those Ingress objects.
Those hostnames must match the filters that you defined (if any):

- By default, `--domain-filter` filters Azure Private DNS zone.
- If you use `--domain-filter` together with `--zone-name-filter`, the behavior changes: `--domain-filter` then filters Ingress domains, not the Azure Private DNS zone name.

When those hostnames are removed or renamed the corresponding DNS records are also altered.

Create the deployment, service and ingress object:

```sh
kubectl create -f nginx.yaml
```

Since your external IP would have already been assigned to the nginx-ingress service, the DNS records pointing to the IP of the nginx-ingress service should be created within a minute.

## Verify created records

Run the following command to view the A records for your Azure Private DNS zone:

```sh
az network private-dns record-set a list -g externaldns -z example.com
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain ('@' indicates the record is for the zone itself).
