# Set up ExternalDNS for Azure Private DNS

This tutorial describes how to set up ExternalDNS for managing records in Azure Private DNS.  

It comprises of the following steps:
1) Install NGINX Ingress Controller 
2) Provision Azure Private DNS
3) Configure service principal for managing the zone
4) Deploy ExternalDNS  

Everything will be deployed on Kubernetes.   
Therefore, please see the subsequent prerequisites.

## Prerequisites
- Azure Kubernetes Service is deployed and ready
- [Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) and `kubectl` installed on the box to execute the subsequent steps 

## Install NGINX Ingress Controller 

Helm is used to deploy the ingress controller. 

We employ the popular chart [stable/nginx-ingress](https://github.com/helm/charts/tree/master/stable/nginx-ingress).

```
$ helm install stable/nginx-ingress \
     --name nginx-ingress \
     --set controller.publishService.enabled=true
```
  
The parameter `controller.publishService.enabled` needs to be set to `true.`  

It will make the ingress controller update the endpoint records of ingress-resources to contain the external-ip of the loadbalancer serving the ingress-controller. 
This is crucial as ExternalDNS reads those endpoints records when creating DNS-Records from ingress-resources.  
In the subsequent parameter we will make use of this. If you don't want to work with ingress-resources in your later use, you can leave the parameter out.

Verify the correct propagation of the loadbalancer's ip by listing the ingresses.
```
$ kubectl get ingress
```
The address column should contain the ip for each ingress. ExternalDNS will pick up exactly this piece of information.
```
NAME     HOSTS             ADDRESS          PORTS   AGE
nginx1   sample1.aks.com   52.167.195.110   80      6d22h
nginx2   sample2.aks.com   52.167.195.110   80      6d21h
```


If you do not want to deploy the ingress controller with Helm, ensure to pass the following cmdline-flags to it through the mechanism of your choice:

```
flags:
--publish-service=<namespace of ingress-controller >/<svcname of ingress-controller>
--update-status=true (default-value)

example:
./nginx-ingress-controller --publish-service=default/nginx-ingress-controller
``` 

## Provision Azure Private DNS

The provider will find suitable zones for domains it manages. It will
not automatically create zones.

For this tutorial, we will create a Azure resource group named 'externaldns' that can easily be deleted later.

```
$ az group create -n externaldns -l westeurope
```

Substitute a more suitable location for the resource group if desired.

As a prerequisite for Azure Private DNS to resolve records is to define links with VNETs.  
Thus, first create a VNET.

```
$ az network vnet create \
  --name myvnet \
  --resource-group externaldns \
  --location westeurope \
  --address-prefix 10.2.0.0/16 \
  --subnet-name mysubnet \
  --subnet-prefixes 10.2.0.0/24
```

Next, create a Azure Private DNS zone for "example.com":

```
$ az network private-dns zone create -g externaldns -n example.com
```

Substitute a domain you own for "example.com" if desired.

Finally, create the mentioned link with the VNET.

```
$ az network private-dns link vnet create -g externaldns -n mylink \
   -z example.com -v myvnet --registration-enabled false
```

## Configure service principal for managing the zone
ExternalDNS needs permissions to make changes in Azure Private DNS.  
These permissions are roles assigned to the service principal used by ExternalDNS.

A service principal with a minimum access level of `Private DNS Zone Contributor` to the Private DNS zone(s) and `Reader` to the resource group containing the Azure Private DNS zone(s) is necessary.
More powerful role-assignments like `Owner` or assignments on subscription-level work too. 

Start off by **creating the service principal** without role-assignments.
```
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

```
# find out the resource ids of the resource group where the dns zone is deployed, and the dns zone itself
$ az group show --name externaldns --query id -o tsv
/subscriptions/id/resourceGroups/externaldns

$ az network private-dns zone show --name example.com -g externaldns --query id -o tsv
/subscriptions/.../resourceGroups/externaldns/providers/Microsoft.Network/privateDnsZones/example.com
```
Now, **create role assignments**.
```
# 1. as a reader to the resource group
$ az role assignment create --role "Reader" --assignee <appId GUID> --scope <resource group resource id>  

# 2. as a contributor to DNS Zone itself
$ az role assignment create --role "Private DNS Zone Contributor" --assignee <appId GUID> --scope <dns zone resource id>  
```

## Deploy ExternalDNS
Configure `kubectl` to be able to communicate and authenticate with your cluster.   
This is per default done through the file `~/.kube/config`.

For general background information on this see [kubernetes-docs](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/).  
Azure-CLI features functionality for automatically maintaining this file for AKS-Clusters. See [Azure-Docs](https://docs.microsoft.com/de-de/cli/azure/aks?view=azure-cli-latest#az-aks-get-credentials).

Then apply one of the following manifests depending on whether you use RBAC or not.

The credentials of the service principal are provided to ExternalDNS as environment-variables.

### Manifest (for clusters without RBAC enabled)
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: externaldns
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: externaldns
    spec:
      containers:
      - name: externaldns
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        env:
        - name: AZURE_TENANT_ID
          value: "<use the tenantId discovered during creation of service principal>"
        - name: AZURE_CLIENT_ID
          value: "<use the aadClientId discovered during creation of service principal>"
        - name: AZURE_CLIENT_SECRET
          value: "<use the aadClientSecret discovered during creation of service principal>"
```

### Manifest (for clusters with RBAC enabled, cluster access)
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: externaldns
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: externaldns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"] 
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
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
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: externaldns
spec:
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        env:
        - name: AZURE_TENANT_ID
          value: "<use the tenantId discovered during creation of service principal>"
        - name: AZURE_CLIENT_ID
          value: "<use the aadClientId discovered during creation of service principal>"
        - name: AZURE_CLIENT_SECRET
          value: "<use the aadClientSecret discovered during creation of service principal>"
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
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: externaldns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
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
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: externaldns
spec:
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com
        - --provider=azure-private-dns
        - --azure-resource-group=externaldns
        - --azure-subscription-id=<use the id of your subscription>
        env:
        - name: AZURE_TENANT_ID
          value: "<use the tenantId discovered during creation of service principal>"
        - name: AZURE_CLIENT_ID
          value: "<use the aadClientId discovered during creation of service principal>"
        - name: AZURE_CLIENT_SECRET
          value: "<use the aadClientSecret discovered during creation of service principal>"
```

Create the deployment for ExternalDNS:

```
$ kubectl create -f externaldns.yaml
```

## Deploying sample service

Create a service file called 'nginx.yaml' with the following contents:

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
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
apiVersion: extensions/v1beta1
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
          serviceName: nginx-svc
          servicePort: 80
        path: /
```

When using ExternalDNS with ingress objects it will automatically create DNS records based on host names specified in ingress objects that match the domain-filter argument in the externaldns deployment manifest. When those host names are removed or renamed the corresponding DNS records are also altered.

Create the deployment, service and ingress object:

```
$ kubectl create -f nginx.yaml
```

Since your external IP would have already been assigned to the nginx-ingress service, the DNS records pointing to the IP of the nginx-ingress service should be created within a minute. 

## Verify created records

Run the following command to view the A records for your Azure Private DNS zone:

```
$ az network private-dns record-set a list -g externaldns -z example.com
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain ('@' indicates the record is for the zone itself).
