# Setting up ExternalDNS for Infoblox

This tutorial describes how to set ExternalDNS up for usage with Infoblox.

Make sure to use **>=0.10.2** version of ExternalDNS for this tutorial. The only WAPI version that
has been validated is **v2.3.1**. It is assumed that the API user has rights to create objects of
the following types: `zone_auth`, `record:a`, `record:cname`, `record:txt`.

This tutorial assumes you have substituted the correct values for the following environment variables:

```
export GRID_HOST=172.18.20.1
export WAPI_PORT=443
export WAPI_VERSION=2.3.1
export WAPI_USERNAME=admin
export WAPI_PASSWORD=infoblox

# these values are for specifying a separate endpoint for read-only requests
export GRID_HOST_RO=172.18.20.2
export WAPI_USERNAME_RO=admin-ro
export WAPI_PASSWORD=infoblox123
```

**Warning**: Grid's Cloud platform member is not supported yet as a target host to be used by ExternalDNS plugin.

You may use Grid Master Candidate node for 'read-only' requests, to load off the Grid Master node.
The parameters which may be defined additionally for this case are:

* infoblox-grid-host-ro
* infoblox-wapi-port-ro
* infoblox-wapi-username-ro
* infoblox-wapi-password-ro
* infoblox-wapi-version-ro
* infoblox-ssl-verify-ro

The parameters may be the same as for 'write' requests. In this case, all the requests go to the same node,
the same as in the case when no 'read-only' parameters are specified.

### More complicated cases when specifying authentication parameters

If you specify either Grid Host or port number for read-only requests which are different from those for 'write' requests,
then Infoblox provider plugin considers a separate set of authentication parameters. Port number is never copied from
the main set of parameters. If 'infoblox-wapi-port-ro' is omited, the default value is used: 443.
The same is about 'infoblox-ssl-verify-ro' and 'infoblox-wapi-version-ro': the default values will be used if omitted,
not the values from 'infoblox-ssl-verify' and 'infoblox-wapi-version' respectively.

You may completely omit them. In this case, they will be assumed the same as for 'write' requests,
with the same parameter's values. But if you specify at least one of the parameters,
then you have to specify a full set of parameters for read-only operations.

There is an ability to use either of password-based or client's certificate-based authentication, but not both
at the same time for the given endpoint ('read' endpoint and 'write' endpoint). But you may use, for example,
password-based authentication for 'write' endpoint and client's certificate-based authentication for 'read' endpoint.
And vice versa.

If you use password-based authentication for a given endpoint, you must specify both user's name AND password.

If you use client's certificate-based authentication, you must specify paths both to certificate's private part file and
to certificate's public part.

By default, 'infoblox-ssl-verify' and 'infoblox-ssl-verify-ro' are set to 'true' which means that
server's certificate must be valid and must be signed by a certificate authority which is
familiar to and trusted by the client (Infoblox provider for External DNS).

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
        image: registry.k8s.io/external-dns/external-dns:v0.13.1
        args:
        - --source=service
        - --domain-filter=example.com             # (optional) limit to only example.com domains.
        - --provider=infoblox
        - --infoblox-grid-host=${GRID_HOST}       # (required) domain name or IP of the Infoblox Grid host.
        - --infoblox-wapi-port=443                # (optional) Infoblox WAPI port. The default is "443".
        - --infoblox-wapi-version=2.3.1           # (optional) Infoblox WAPI version. The default is "2.3.1"
        - --infoblox-ssl-verify                   # (optional) Use --no-infoblox-ssl-verify to skip server certificate verification.
        - --infoblox-create-ptr                   # (optional) Use --infoblox-create-ptr to create a ptr entry in addition to an entry.
        - --infoblox-grid-host-ro=${GRID_HOST_RO} # (required for a separate read-only endpoint) domain name or IP of the Infoblox Grid host.
        - --infoblox-wapi-port-ro=443             # (optional, for a separate read-only endpoint) Infoblox WAPI port. The default is "443".
        - --infoblox-ssl-verify-ro                # (optional, for a separate read-only endpoint) Use --no-infoblox-ssl-verify-ro to skip server certificate verification.
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
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME_RO
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME_RO
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD_RO
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD_RO
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
        image: registry.k8s.io/external-dns/external-dns:v0.13.1
        args:
          - --source=service
          - --domain-filter=example.com             # (optional) limit to only example.com domains.
          - --provider=infoblox
          - --infoblox-grid-host=${GRID_HOST}       # (required) domain name or IP of the Infoblox Grid host.
          - --infoblox-wapi-port=443                # (optional) Infoblox WAPI port. The default is "443".
          - --infoblox-wapi-version=2.3.1           # (optional) Infoblox WAPI version. The default is "2.3.1"
          - --infoblox-ssl-verify                   # (optional) Use --no-infoblox-ssl-verify to skip server certificate verification.
          - --infoblox-create-ptr                   # (optional) Use --infoblox-create-ptr to create a ptr entry in addition to an entry.
          - --infoblox-grid-host-ro=${GRID_HOST_RO} # (required for a separate read-only endpoint) domain name or IP of the Infoblox Grid host.
          - --infoblox-wapi-port-ro=443             # (optional, for a separate read-only endpoint) Infoblox WAPI port. The default is "443".
          - --infoblox-ssl-verify-ro                # (optional, for a separate read-only endpoint) Use --no-infoblox-ssl-verify-ro to skip server certificate verification.
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
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME_RO
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_USERNAME_RO
        - name: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD_RO
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_INFOBLOX_WAPI_PASSWORD_RO
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

There is also the ability to filter results from the Infoblox zone_auth service
based upon a regular expression.
See the [Infoblox API document](https://www.infoblox.com/wp-content/uploads/infoblox-deployment-infoblox-rest-api.pdf)
for examples. To use this feature for the zone_auth service,
set the parameter 'infoblox-fqdn-regex' for external-dns to a regular expression that
makes sense for you.  For instance, to only return hosted zones that start with
'staging' in the 'test.com' domain (like staging.beta.test.com, or staging.test.com),
use the following command line option when starting external-dns:

```
--infoblox-fqdn-regex='^staging.*test.com$'
```

The regular expression must take into account all the zones (including reverse-mapping)
which are to be populated by the plugin. Any zone which does not match the regular
expression will be ignored, no matter if it is a forward- or reverse-mapping zone.
This note is specifically for the case when 'infoblox-create-ptr' is in use. 

To select one forward-mapping zone and all of reverse-mapping zones,
you may use the following CLI option:
```
--infoblox-fqdn-regex='(^example\.org$)|(^(\d{1,3}\.){3}\d{1,3}(\/\d{1,3})?$)'
```

The option 'infoblox-fqdn-regex' is similar to 'domain-filter' but works on NIOS side,
only zones which match the regular expression are returned by NIOS server.
You may use both 'infoblox-fqdn-regex' and 'domain-filter' when it makes sense for you;
but remember that if a zone matches only the regular expression from 'infoblox-fqdn-regex'
option but not any of values from 'domain-filter' options or vise versa,
then the zone will not be impacted or taken into account by Infoblox provider for ExternalDNS.

## Infoblox PTR record support

There is an option to enable PTR records support for infoblox provider.
PTR records allow to do reverse dns search. To enable PTR records support,
add following into arguments for external-dns:  
`--infoblox-create-ptr` to allow management of PTR records.  
You can also add a filter for reverse DNS zone to limit PTR records
to specific zones only:  
```--domain-filter=10.0.0.0/8```
Now external-dns will manage PTR records for you.

**Important note 1**: you may expect that the reverse-mapping zone must be in the
form of a subdomain of the 'in-addr.arpa' domain, but in case of ExternalDNS's Infoblox provider and
'domain-filter' option you must specify the subnet, with its network mask,
in CIDR-form as in the example above.

**Important note 2**: if Infoblox provider was in use without 'infoblox-create-ptr' option but
you change this eventually, new PTR-records will be created only for newly created A-records.
Already existing A-records will be left without PTR-records unless endpoints which
were the sources for the A-records will be changed in the Kubernetes cluster. 

## TXT-records

If '--registry txt' option is used to configure ExternalDNS plugin, then TXT-records
will be created along with corresponding A-records to indicate that this A-record is
managed by ExternalDNS plugin.

Even more, starting from the version released on 12th of April, 2022,
an additional TXT-record will be created, prefixed by record-type ('a-' for A-records).
This is for transitional period of time. For details, see [registry.md](../registry.md) document.
