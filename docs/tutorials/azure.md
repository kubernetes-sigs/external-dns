# Setting up ExternalDNS for Services on Azure

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster on Azure.

Make sure to use **>=0.4.0** version of ExternalDNS for this tutorial.

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

When your Kubernetes cluster is created by Azure Container Services, a file named `/etc/kubernetes/azure.json` is created to store
the Azure credentials for API access.  Kubernetes uses this file for the Azure cloud provider.

For ExternalDNS to access the Azure API, it also needs access to this file.  However, we will be deploying ExternalDNS inside of
the Kubernetes cluster so we will need to use a Kubernetes secret.

The Azure DNS provider expects, by default, that the configuration file is at `/etc/kubernetes/azure.json`.  This can be overridden with
the `--azure-config-file` option when starting ExternalDNS.

To create the secret:

```
$ kubectl create secret generic azure-config-file --from-file=/etc/kubernetes/azure.json
```

## Deploy ExternalDNS

Create a deployment file called `externaldns.yaml` with the following contents:

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
        image: registry.opensource.zalan.do/teapot/external-dns:v0.4.0
        args:
        - --source=service
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
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the Azure DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

It takes a little while for the Azure cloud provider to create an external IP for the service.  Check the status by running
`kubectl get services nginx`.  If the `EXTERNAL-IP` field shows an address, the service is ready to be accessed externally.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the Azure DNS records.

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
