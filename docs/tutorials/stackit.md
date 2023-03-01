# Setting up ExternalDNS for Services on STACKIT

This tutorial describes how to set up ExternalDNS for usage within a Kubernetes cluster using [STACKIT DNS](https://www.stackit.de/en/).

The following steps are required to use STACKIT with ExternalDNS:

1. Create a STACKIT [customer account](https://portal.stackit.cloud/customer-accounts)
2. Create a STACKIT [project](https://portal.stackit.cloud/projects/new)
3. Add your zone to STACKIT DNS
4. Create a service account in the STACKIT project
5. Create a bearer token from the service account
6. Give the service account the permission to manage DNS records
7. Deploy ExternalDNS to use the STACKIT provider
8. Verify the setup by deploying a test service (optional)

## Creating a STACKIT DNS zone

Before records can be added to your domain name automatically, you need to add your domain name to the set of zones 
managed by STACKIT. In order to add the zone, perform the following steps:

- Create the zone in the STACKIT portal in your project
- Or use the STACKIT DNS API to create the zone `curl -X POST -H "Authorziation: Bearer <your token>" 
-H "Content-Type: application/json" -d '{"name": "example", "dnsName": example.com"}' https://api.dns.stackit.cloud/v1/projects/<projectId>/zones`

Note that "SECONDARY" domains cannot be managed by ExternalDNS, because this would not allow modification of records in the zone.

## Deploy an example service

To use ExternalDNS, you need to deploy a service from type LoadBalancer or NodePort.
The following section shows an example service.

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
    external-dns.alpha.kubernetes.io/hostname: example.com # replace example.com with your zone
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Change the file as follows:

- Replace the annotation of the service; use the same hostname as the STACKIT DNS zone created above

Create the deployment and service:

```bash
$ kubectl create -f nginx.yaml
```

Depending on your cloud provider, it might take a while to create an external IP for the service. Once an external IP 
address is assigned to the service, ExternalDNS will notice the new address and synchronize the STACKIT DNS records accordingly.

## Use ExternalDNS

To use ExternalDNS, you need to deploy it to your cluster. The following section shows how to deploy ExternalDNS using 
a deployment without rbac to illustrate what flags need to be configured.

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
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns
        args:
        - --source=service
        - --source="ingress"
        - --policy="sync" # (optional) ExternalDNS will create and delete records in the zone
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone created above.
        - --provider=stackit
        - --stackit-project-id=<projectId>
        - --stackit-client-token=<bearerToken> # specify the token created in the previous step here or use the env var
        - --stackit-base-url=https://api.dns.stackit.cloud # (optional) use this to override the default base url
        env:
        - name: EXTERNAL_DNS_STACKIT_CLIENT_TOKEN # use this to specify the token as an environment variable if not specified as a flag
          value: "YOUR_STACKIT_BEARER_TOKEN"
```
