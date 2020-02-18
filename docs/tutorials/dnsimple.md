# Setting up ExternalDNS for Services on DNSimple


This tutorial describes how to setup ExternalDNS for usage with DNSimple.

Make sure to use **>=0.4.6** version of ExternalDNS for this tutorial.

## Created a DNSimple API Access Token

A DNSimple API access token can be acquired by following the [provided documentation from DNSimple](https://support.dnsimple.com/articles/api-access-token/)

The environment variable `DNSIMPLE_OAUTH` must be set to the API token generated for to run ExternalDNS with DNSimple.

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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone you create in DNSimple.
        - --provider=dnsimple
        - --registry=txt
        env:
        - name: DNSIMPLE_OAUTH
          value: "YOUR_DNSIMPLE_API_KEY"
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
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]
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
        image: registry.opensource.zalan.do/teapot/external-dns:latest
        args:
        - --source=service
        - --domain-filter=example.com # (optional) limit to only example.com domains; change to match the zone you create in DNSimple.
        - --provider=dnsimple
        - --registry=txt
        env:
        - name: DNSIMPLE_OAUTH
          value: "YOUR_DNSIMPLE_API_KEY"
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
    external-dns.alpha.kubernetes.io/hostname: validate-external-dns.example.com
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the DNSimple DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```sh
$ kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service. Check the status by running
`kubectl get services nginx`.  If the `EXTERNAL-IP` field shows an address, the service is ready to be accessed externally.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the DNSimple DNS records.

## Verifying DNSimple DNS records

### Getting your DNSimple Account ID

If you do not know your DNSimple account ID it can be aquired using the [whoami](https://developer.dnsimple.com/v2/identity/#whoami) endpoint from the DNSimple Identity API

```sh
curl -H "Authorization: Bearer $DNSIMPLE_ACCOUNT_TOKEN" \
    -H 'Accept: application/json' \
    https://api.dnsimple.com/v2/whoami
{
  "data": {
    "user": null,
    "account": {
      "id": 1,
      "email": "example-account@example.com",
      "plan_identifier": "dnsimple-professional",
      "created_at": "2015-09-18T23:04:37Z",
      "updated_at": "2016-06-09T20:03:39Z"
    }
  }
}
```

### Looking at the DNSimple Dashboard

You can view your DNSimple Record Editor at https://dnsimple.com/a/YOUR_ACCOUNT_ID/domains/example.com/records. Ensure you substitute the value `YOUR_ACCOUNT_ID` with the ID of your DNSimple account and `example.com` with the correct domain that you used during validation.

### Using the DNSimple Zone Records API

This approach allows for you to use the DNSimple [List records for a zone](https://developer.dnsimple.com/v2/zones/records/#listZoneRecords) endpoint to verify the creation of the A and TXT record. Ensure you substitute the value `YOUR_ACCOUNT_ID` with the ID of your DNSimple account and `example.com` with the correct domain that you used during validation.

```sh
curl -H "Authorization: Bearer $DNSIMPLE_ACCOUNT_TOKEN" \
    -H 'Accept: application/json' \
    'https://api.dnsimple.com/v2/YOUR_ACCOUNT_ID/zones/example.com/records&name=validate-external-dns'
```

## Clean up

Now that we have verified that ExternalDNS will automatically manage DNSimple DNS records, we can delete the tutorial's example:

```sh
$ kubectl delete -f nginx.yaml
$ kubectl delete -f externaldns.yaml
```

### Deleting Created Records

The created records can be deleted using the record IDs from the verification step and the [Delete a zone record](https://developer.dnsimple.com/v2/zones/records/#deleteZoneRecord) endpoint.
