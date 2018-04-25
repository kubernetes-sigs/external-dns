
# Setting up ExternalDNS for Services on Azure

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on Azure.

Make sure to use **>=0.4.2** version of ExternalDNS for this tutorial.

This tutorial uses [Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) for all
Azure commands and assumes that the Kubernetes cluster was created via Azure Container Services and `kubectl` commands
are being run on an orchestration master.

## Creating a Azure DNS zone

The Azure provider for ExternalDNS will find suitable zones for domains it manages; it will
not automatically create zones.

For this tutorial, we will create a Azure resource group named 'externaldns' that can easily be deleted later:

```
$ az group create -n externaldns -l eastus
```

Substitute a more suitable location for the resource group if desired.

Next, create a Azure DNS zone for "example.com":

```
$ az network dns zone create -g externaldns -n example.com
```

Substitute a domain you own for "example.com" if desired.

If using your own domain that was registered with a third-party domain registrar, you should point your domain's
name servers to the values in the `nameServers` field from the JSON data returned by the `az network dns zone create` command.
Please consult your registrar's documentation on how to do that.

## Creating Azure Credentials Secret
The Azure DNS provider expects, by default, that the configuration file is at `/etc/kubernetes/azure.json`.  This can be overridden with
the `--azure-config-file` option when starting ExternalDNS.

### Azure Container Services
When your Kubernetes cluster is created by ACS, a file named `/etc/kubernetes/azure.json` is created to store
the Azure credentials for API access.  Kubernetes uses this file for the Azure cloud provider.

For ExternalDNS to access the Azure API, it also needs access to this file.  However, we will be deploying ExternalDNS inside of
the Kubernetes cluster so we will need to use a Kubernetes secret.

To create the secret:

```
$ kubectl create secret generic azure-config-file --from-file=/etc/kubernetes/azure.json
```

### Azure Kubernetes Services (aka AKS)
When your cluster is created, unlike ACS there are no Azure credentials stored and you must create an azure.json object manually like with other hosting providers. In order to create the azure.json you must first create an Azure AD service principal in the Azure AD tenant linked to your Azure subscription that is hosting your DNS zone.

#### Create service principal
A Service Principal with a minimum access level of contribute to the resource group containing the Azure DNS zone(s) is necessary for ExternalDNS to be able to edit DNS records. This is an Azure CLI example on how to query the Azure API for the information required for the Resource Group and DNS zone you would have already created in previous steps.

```
>az login
...
# find the relevant subscription and set the az context. id = subscriptionId value in the azure.json.
>az account list
{
    "cloudName": "AzureCloud",
    "id": "<subscriptionId GUID>",
    "isDefault": false,
    "name": "My Subscription",
    "state": "Enabled",
    "tenantId": "AzureAD tenant ID",
    "user": {
      "name": "name",
      "type": "user"
}
>az account set -s id
...
>az group show --name externaldns
{
  "id": "/subscriptions/id/resourceGroups/externaldns",
  ...
}

# use the id from the previous step in the scopes argument
>az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/id/resourceGroups/externaldns" -n ExternalDnsServicePrincipal
{
  "appId": "appId GUID",  <-- aadClientId value
  ...
  "password": "password",  <-- aadClientSecret value
  "tenant": "AzureAD Tenant Id"  <-- tenantId value
}
...

```
### Other hosting providers
If the Kubernetes cluster is not hosted by Azure Container Services and you still want to use Azure DNS, you need to create the secret manually. The secret should contain an object named azure.json with content similar to this:
```
{
  "tenantId": "AzureAD tenant Id",
  "subscriptionId": "Id",
  "aadClientId": "Service Principal AppId",
  "aadClientSecret": "Service Principal Password",
  "resourceGroup": "MyDnsResourceGroup",
}
```
If you have all the information necessary: create a file called azure.json containing the json structure above and substitute the values. Otherwise create a service principal as previously shown before creating the Kubernetes secret.

Then add the secret to the Kubernetes cluster before continuing:
```
kubectl create secret generic azure-config-file --from-file=azure.json
```



## Deploy ExternalDNS

This deployment assumes that you will be using nginx-ingress. When using nginx-ingress do not deploy it as a Daemon Set. This causes nginx-ingress to write the Cluster IP of the backend pods in the ingress status.loadbalancer.ip property which then has external-dns write the Cluster IP(s) in DNS vs. the nginx-ingress service external IP.

Ensure that your nginx-ingress deployment has the following arg: added to it:

```
- --publish-service=namespace/nginx-ingress-controller-svcname
```

For more details see here: [nginx-ingress external-dns](https://github.com/kubernetes-incubator/external-dns/blob/master/docs/faq.md#why-is-externaldns-only-adding-a-single-ip-address-in-route-53-on-aws-when-using-the-nginx-ingress-controller-how-do-i-get-it-to-use-the-fqdn-of-the-elb-assigned-to-my-nginx-ingress-controller-service-instead)

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS.

### Manifest (for clusters without RBAC enabled)
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      containers:
      - name: external-dns
        image: registry.opensource.zalan.do/teapot/external-dns:v0.5.0
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=azure
        - --azure-resource-group=externaldns # (optional) use the DNS zones from the tutorial's resource group
        volumeMounts:
        - name: azure-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: azure-config-file
        secret:
          secretName: azure-config-file
```

### Manifest (for clusters with RBAC enabled)
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"] 
  resources: ["ingresses"] 
  verbs: ["get","watch","list"]
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
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.opensource.zalan.do/teapot/external-dns:v0.5.0
        args:
        - --source=service
        - --source=ingress
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=azure
        - --azure-resource-group=externaldns # (optional) use the DNS zones from the tutorial's resource group
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

```
$ kubectl create -f externaldns.yaml
```

## Deploying an Nginx Service

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
    protocol: tcp
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

When using external-dns with ingress objects it will automatically create DNS records based on host names specified in ingress objects that match the domain-filter argument in the external-dns deployment manifest. When those host names are removed or renamed the corresponding DNS records are also altered.

Create the deployment, service and ingress object:

```
$ kubectl create -f nginx.yaml
```

Since your external IP would have already been assigned to the nginx-ingress service, the DNS records pointing to the IP of the nginx-ingress service should be created within a minute. 

## Verifying Azure DNS records

Run the following command to view the A records for your Azure DNS zone:

```
$ az network dns record-set a list -g externaldns -z example.com
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain ('@' indicates the record is for the zone itself).

## Delete Azure Resource Group

Now that we have verified that ExternalDNS will automatically manage Azure DNS records, we can delete the tutorial's
resource group:

```
$ az group delete -n externaldns
```
