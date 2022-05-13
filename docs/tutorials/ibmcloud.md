# Setting up ExternalDNS for Services on IBMCloud

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using IBMCloud DNS.

This tutorial uses [IBMCloud CLI](https://cloud.ibm.com/docs/cli?topic=cli-getting-started) for all
IBM Cloud commands and assumes that the Kubernetes cluster was created via IBM Cloud Kubernetes Service and `kubectl` commands
are being run on an orchestration node.

## Creating a IBMCloud DNS zone
The IBMCloud provider for ExternalDNS will find suitable zones for domains it manages; it will
not automatically create zones.
For public zone, This tutorial assume that the [IBMCloud Internet Services](https://cloud.ibm.com/catalog/services/internet-services) was provisioned and the [cis cli plugin](https://cloud.ibm.com/docs/cis?topic=cis-cli-plugin-cis-cli) was installed with IBMCloud CLI
For private zone, This tutorial assume that the [IBMCloud DNS Services](https://cloud.ibm.com/catalog/services/dns-services) was provisioned and the [dns cli plugin](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-cli-plugin-dns-services-cli-commands) was installed with IBMCloud CLI

### Public Zone
For this tutorial, we create public zone named `example.com` on IBMCloud Internet Services instance `external-dns-public`
```
$ ibmcloud cis domain-add example.com -i external-dns-public
```
Follow [step](https://cloud.ibm.com/docs/cis?topic=cis-getting-started#configure-your-name-servers-with-the-registrar-or-existing-dns-provider) to active your zone

### Private Zone
For this tutorial, we create private zone named `example.com` on IBMCloud DNS Services instance `external-dns-private`
```
$ ibmcloud dns zone-create example.com -i external-dns-private
```

## Creating configuration file

The preferred way to inject the configuration file is by using a Kubernetes secret. The secret should contain an object named azure.json with content similar to this:

```
{
  "apiKey": "1234567890abcdefghijklmnopqrstuvwxyz",
  "instanceCrn": "crn:v1:bluemix:public:internet-svcs:global:a/bcf1865e99742d38d2d5fc3fb80a5496:b950da8a-5be6-4691-810e-36388c77b0a3::"
}
```

You can create or find the `apiKey` in your ibmcloud IAM --> [API Keys page](https://cloud.ibm.com/iam/apikeys)

You can find the `instanceCrn` in your service instance details

Now you can create a file named 'ibmcloud.json' with values gathered above and with the structure of the example above. Use this file to create a Kubernetes secret:
```
$ kubectl create secret generic ibmcloud-config-file --from-file=/local/path/to/ibmcloud.json
```
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
        image: k8s.gcr.io/external-dns/external-dns:v0.12.0
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ibmcloud
        - --ibmcloud-proxied # (optional) enable the proxy feature of IBMCloud
        volumeMounts:
        - name: ibmcloud-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: ibmcloud-config-file
        secret:
          secretName: ibmcloud-config-file
          items:
          - key: externaldns-config.json
            path: ibmcloud.json
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
        image: k8s.gcr.io/external-dns/external-dns:v0.12.0
        args:
        - --source=service # ingress is also possible
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=ibmcloud
        - --ibmcloud-proxied # (optional) enable the proxy feature of IBMCloud public zone
        volumeMounts:
        - name: ibmcloud-config-file
          mountPath: /etc/kubernetes
          readOnly: true
      volumes:
      - name: ibmcloud-config-file
        secret:
          secretName: ibmcloud-config-file
          items:
          - key: externaldns-config.json
            path: ibmcloud.json
```

## Deploying an Nginx Service

Create a service file called `nginx.yaml` with the following contents:

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
    external-dns.alpha.kubernetes.io/hostname: www.example.com
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

Note the annotation on the service; use the hostname as the IBMCloud DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

By setting the TTL annotation on the service, you have to pass a valid TTL, which must be 120 or above.
This annotation is optional, if you won't set it, it will be 1 (automatic) which is 300.

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the IBMCloud DNS records.

## Verifying IBMCloud DNS records
Run the following command to view the A records:

### Public Zone
```
# Get the domain ID with below command on IBMCloud Internet Services instance `external-dns-public`
$ ibmcloud cis domains -i external-dns-public
# Get the records with domain ID
$ ibmcloud cis dns-records DOMAIN_ID  -i external-dns-public
```

### Private Zone
```
# Get the domain ID with below command on IBMCloud DNS Services instance `external-dns-private`
$ ibmcloud dns zones -i external-dns-private
# Get the records with domain ID
$ ibmcloud dns resource-records ZONE_ID  -i external-dns-public
```
This should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage IBMCloud DNS records, we can delete the tutorial's example:

```
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```

## Setting proxied records on public zone

Using the `external-dns.alpha.kubernetes.io/ibmcloud-proxied: "true"` annotation on your ingress or service, you can specify if the proxy feature of IBMCloud public DNS should be enabled for that record. This setting will override the global `--ibmcloud-proxied` setting.

## Active priviate zone with VPC allocated

By default, IBMCloud DNS Services don't active your private zone with new zone added, with externale DNS, you can use `external-dns.alpha.kubernetes.io/ibmcloud-vpc: "crn:v1:bluemix:public:is:us-south:a/bcf1865e99742d38d2d5fc3fb80a5496::vpc:r006-74353823-a60d-42e4-97c5-5e2551278435"` annotation on your ingress or service, it will active your private zone with in specific VPC for that record created in. this setting won't work if the private zone was active already.

Note: the annotaion value is the VPC CRN, every IBM Cloud service have a valid CRN.