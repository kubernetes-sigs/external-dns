# Setting up ExternalDNS for EfficientIP

This tutorial describes how to setup ExternalDNS for usage with EfficientIP.

Make sure to use **>=0.7.7** version of ExternalDNS for this tutorial. The SOLIDserver version that
has been validated is **7.3.x** and **8.0.x**. It is assumed that the API user has rights to create and list objects of
the following types: `Zones`, `Records`.

This tutorial assumes you have substituted the correct values for the following environment variables:

```
export EFFICIENTIP_HOST=127.0.0.1
export EFFICIENTIP_PORT=443
export EFFICIENTIP_USERNAME=username
export EFFICIENTIP_PASSWORD=password
```

## Creating an EfficientIP DNS zone

The EfficientIP provider for ExternalDNS will find suitable zones for domains it manages; it will
not automatically create zones.

Create an EfficientIP DNS zone for "example.org":

```
$ curl -kl \
      -X POST \
      -d server_name=server.name \
      -d zone_name=example.org \
      -d zone_type=master \
      -u ${EFFICIENTIP_USERNAME}:${EFFICIENTIP_PASSWORD} \
         https://${EFFICIENTIP_HOST}:${EFFICIENTIP_PORT}/api/v2.0/dns/zone/add
```

Substitute "server.name" by your own server name in the SOLIDserver and use the domain of your choice "example.org".

## Creating an EfficientIP Configuration Secret

For ExternalDNS to access the EfficientIP API, create a Kubernetes secret.

To create the secret:

```
$ kubectl create secret generic external-dns \
      --from-literal=EXTERNAL_DNS_EFFICIENTIP_USERNAME=${EFFICIENTIP_USERNAME} \
      --from-literal=EXTERNAL_DNS_EFFICIENTIP_PASSWORD=${EFFICIENTIP_PASSWORD}
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
        image: k8s.gcr.io/external-dns/external-dns:v0.7.7
        args:
        - --source=service
        - --domain-filter=example.org            # (optional) limit to only example.org domains.
        - --provider=efficientip
        - --efficientip-host=${EFFICIENTIP_HOST} # (required) IP of the EfficientIP SOLIDserver host.
        - --efficientip-port=443                 # (optional) EfficientIP SOLIDserver port. The default is "443".
        - --efficientip-ssl-verify               # (optional) Use --no-efficientip-ssl-verify to skip server certificate verification.
        env:
        - name: EXTERNAL_DNS_EFFICIENTIP_USERNAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_EFFICIENTIP_USERNAME
        - name: EXTERNAL_DNS_EFFICIENTIP_PASSWORD
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_EFFICIENTIP_PASSWORD
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
        - --domain-filter=example.org            # (optional) limit to only example.org domains.
        - --provider=efficientip
        - --efficientip-host=${EFFICIENTIP_HOST} # (required) IP of the EfficientIP SOLIDserver host.
        - --efficientip-port=443                 # (optional) EfficientIP SOLIDserver port. The default is "443".
        - --efficientip-ssl-verify               # (optional) Use --no-efficientip-ssl-verify to skip server certificate verification.
        env:
        - name: EXTERNAL_DNS_EFFICIENTIP_USERNAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_EFFICIENTIP_USERNAME
        - name: EXTERNAL_DNS_EFFICIENTIP_PASSWORD
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_EFFICIENTIP_PASSWORD
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
    external-dns.alpha.kubernetes.io/hostname: example.org
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Note the annotation on the service; use the same hostname as the EfficientIP DNS zone created above. The annotation may also be a subdomain
of the DNS zone (e.g. 'www.example.org').

ExternalDNS uses this annotation to determine what services should be registered with DNS.  Removing the annotation
will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```
$ kubectl create -f nginx.yaml
```

It takes a little while for the EfficientIP cloud provider to create an external IP for the service.  Check the status by running
`kubectl get services nginx`.  If the `EXTERNAL-IP` field shows an address, the service is ready to be accessed externally.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize
the EfficientIP DNS records.

## Verifying EfficientIP DNS records

Run the following command to view the A records for your EfficientIP DNS zone:

```
$ curl -kl \
      -X GET \
      -u ${EFFICIENTIP_USERNAME}:${EFFICIENTIP_PASSWORD} \
         https://${EFFICIENTIP_HOST}:${EFFICIENTIP_PORT}/api/v2.0/dns/rr/list?where=zone_name%3D%27example.org%27+and+rr_type+%3D+%27A%27
```

Substitute the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain

## Clean up

Now that we have verified that ExternalDNS will automatically manage EfficientIP DNS records, we can delete the tutorial's
DNS zone and examples:
```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
$ curl -kl \
      -X DELETE \
      -u ${EFFICIENTIP_USERNAME}:${EFFICIENTIP_PASSWORD} \
         https://${EFFICIENTIP_HOST}:${EFFICIENTIP_PORT}/api/v2.0/dns/zone/delete?server_name=server.name\&zone_name=example.org
```

Do not forget to replace "server.name" and "example.org" by your correct informations

## Ability to filter results from the zone API using a regular expression

There is also the ability to filter results from the EfficientIP zone service based upon a regular expression. To use this feature for the zone service, set the parameter regex-domain-filter for external-dns to a regular expression that makes sense for you.  For instance, to only return hosted zones that start with staging in the test.com domain (like staging.beta.test.com, or staging.test.com), use the following command line option when starting external-dns

```
--regex-domain-filter=^staging.*test.com$
```
