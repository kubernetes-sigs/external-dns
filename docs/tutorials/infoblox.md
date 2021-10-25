# Setting up ExternalDNS for Infoblox

This tutorial describes how to setup ExternalDNS for usage with Infoblox.

Make sure to use **>=0.4.6** version of ExternalDNS for this tutorial. The only WAPI version that
has been validated is **v2.3.1**. It is assumed that the API user has rights to create objects of
the following types: `zone_auth`, `record:a`, `record:cname`, `record:txt`.

This tutorial assumes you have substituted the correct values for the following environment variables:

```
export GRID_HOST=127.0.0.1
export WAPI_PORT=443
export WAPI_VERSION=2.3.1
export WAPI_USERNAME=admin
export WAPI_PASSWORD=infoblox
```

## Creating an Infoblox DNS zone

The Infoblox provider for ExternalDNS will find suitable zones for domains it manages; it will
not automatically create zones.

Create an Infoblox DNS zone for "example.com":

```
$ curl -kl \
      -X POST \
      -d fqdn=example.com \
      -u ${WAPI_USERNAME}:${WAPI_PASSWORD} \
         https://${GRID_HOST}:${WAPI_PORT}/wapi/v${WAPI_VERSION}/zone_auth
```

Substitute a domain you own for "example.com" if desired.

## Creating an Infoblox Configuration Secret

For ExternalDNS to access the Infoblox API, create a Kubernetes secret.

To create the secret:

```
$ kubectl create secret generic external-dns \
      --from-literal=EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME=${WAPI_USERNAME} \
      --from-literal=EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD=${WAPI_PASSWORD}
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --domain-filter=example.com       # (optional) limit to only example.com domains.
        - --provider=infoblox
        - --infoblox-grid-host=${GRID_HOST} # (required) IP of the Infoblox Grid host.
        - --infoblox-wapi-port=443          # (optional) Infoblox WAPI port. The default is "443".
        - --infoblox-wapi-version=2.3.1     # (optional) Infoblox WAPI version. The default is "2.3.1"
        - --infoblox-ssl-verify             # (optional) Use --no-infoblox-ssl-verify to skip server certificate verification.
        - --infoblox-create-ptr             # (optional) Use --infoblox-create-ptr to create a ptr entry in addition to an entry.
        env:
        - name: EXTERNAL_DNS_INFOBLOX_HTTP_POOL_CONNECTIONS
          value: "10" # (optional) Infoblox WAPI request connection pool size. The default is "10".
        - name: EXTERNAL_DNS_INFOBLOX_HTTP_REQUEST_TIMEOUT
          value: "60" # (optional) Infoblox WAPI request timeout in seconds. The default is "60".
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service
        - --domain-filter=example.com       # (optional) limit to only example.com domains.
        - --provider=infoblox
        - --infoblox-grid-host=${GRID_HOST} # (required) IP of the Infoblox Grid host.
        - --infoblox-wapi-port=443          # (optional) Infoblox WAPI port. The default is "443".
        - --infoblox-wapi-version=2.3.1     # (optional) Infoblox WAPI version. The default is "2.3.1"
        - --infoblox-ssl-verify             # (optional) Use --no-infoblox-ssl-verify to skip server certificate verification.
        - --infoblox-create-ptr             # (optional) Use --infoblox-create-ptr to create a ptr entry in addition to an entry.
        env:
        - name: EXTERNAL_DNS_INFOBLOX_HTTP_POOL_CONNECTIONS
          value: "10" # (optional) Infoblox WAPI request connection pool size. The default is "10".
        - name: EXTERNAL_DNS_INFOBLOX_HTTP_REQUEST_TIMEOUT
          value: "60" # (optional) Infoblox WAPI request timeout in seconds. The default is "60".
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD
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
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the Infoblox DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.com').

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

It takes a little while for the Infoblox cloud provider to create an external IP for the service.  Check the status by running
`kubectl get services nginx`.  If the `EXTERNAL-IP` field shows an address, the service is ready to be accessed externally.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the Infoblox DNS records.

## Verifying Infoblox DNS records

Run the following command to view the A records for your Infoblox DNS zone:

```
$ curl -kl \
      -X GET \
      -u ${WAPI_USERNAME}:${WAPI_PASSWORD} \
         https://${GRID_HOST}:${WAPI_PORT}/wapi/v${WAPI_VERSION}/record:a?zone=example.com
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain ('@' indicates the record is for the zone itself).

## Clean up

Now that we have verified that ExternalDNS will automatically manage Infoblox DNS records, we can delete the tutorial's
DNS zone:

```
$ curl -kl \
      -X DELETE \
      -u ${WAPI_USERNAME}:${WAPI_PASSWORD} \
         https://${GRID_HOST}:${WAPI_PORT}/wapi/v${WAPI_VERSION}/zone_auth?fqdn=example.com
```

## Ability to filter results from the zone auth API using a regular expression

There is also the ability to filter results from the Infoblox zone_auth service based upon a regular expression.  See the [Infoblox API document](https://www.infoblox.com/wp-content/uploads/infoblox-deployment-infoblox-rest-api.pdf) for examples.  To use this feature for the zone_auth service, set the parameter infoblox-fqdn-regex for external-dns to a regular expression that makes sense for you.  For instance, to only return hosted zones that start with staging in the test.com domain (like staging.beta.test.com, or staging.test.com), use the following command line option when starting external-dns

```
--infoblox-fqdn-regex=^staging.*test.com$
```

## Infoblox PTR record support

There is an option to enable PTR records support for infoblox provider. PTR records allow to do reverse dns search. To enable PTR records support, add following into arguments for external-dns:  
`--infoblox-create-ptr` to allow management of PTR records.  
You can also add a filter for reverse dns zone to limit PTR records to specific zones only:  
`--domain-filter=10.196.0.0/16` change this to the reverse zone(s) as defined in your infoblox.  
Now external-dns will manage PTR records for you.
