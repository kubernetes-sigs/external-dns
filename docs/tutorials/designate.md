# Setting up ExternalDNS for Services on OpenStack Designate

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using OpenStack Designate DNS.

## Authenticating with OpenStack

We are going to use OpenStack CLI - `openstack` utility, which is an umbrella application for most of OpenStack clients including `designate`.

All OpenStack CLIs require authentication parameters to be provided. These parameters include:
* URL of the OpenStack identity service (`keystone`) which is responsible for user authentication and also served as a registry for other
  OpenStack services. Designate endpoints must be registered in `keystone` in order to ExternalDNS and OpenStack CLI be able to find them.
* OpenStack region name
* User login name.
* User project (tenant) name.
* User domain (only when using keystone API v3)

Although these parameters can be passed explicitly through the CLI flags, traditionally it is done by sourcing `openrc` file (`source ~/openrc`) that is a
shell snippet that sets environment variables that all OpenStack CLI understand by convention.

Recent versions of OpenStack Dashboard have a nice UI to download `openrc` file for both v2 and v3 auth protocols. Both protocols can be used with ExternalDNS.
v3 is generally preferred over v2, but might not be available in some OpenStack installations.

## Installing OpenStack Designate

Please refer to the Designate deployment [tutorial](https://docs.openstack.org/project-install-guide/dns/ocata/install.html) for instructions on how
to install and test Designate with BIND backend. You will be required to have admin rights in existing OpenStack installation to do this. One convenient
way to get yourself an OpenStack installation to play with is to use [DevStack](https://docs.openstack.org/devstack/latest/).

## Creating DNS zones

All domain names that ExternalDNS is going to create must belong to one of the DNS zones created in advance. Here is an example of how to create the `external-dns-test.my-org.com` DNS zone:
```console
$ openstack zone create --email dnsmaster@external-dns-test.my-org.com external-dns-test.my-org.com.
```

It is important to manually create all the zones that are going to be used for kubernetes entities (ExternalDNS sources) before starting ExternalDNS.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then apply one of the following manifests file to deploy ExternalDNS. You can check if your cluster has RBAC by `kubectl api-versions | grep rbac.authorization.k8s.io`.

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
        image: registry.opensource.zalan.do/teapot/external-dns
        args:
        - --source=service 
        - --source=ingress 
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=designate
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --registry=txt
        - --txt-owner-id=my-hostedzone-identifier
        env: # values from openrc file
        - name: OS_AUTH_URL
          value: http://controller/identity/v3
        - name: OS_REGION_NAME
          value: RegionOne
        - name: OS_USERNAME
          value: admin
        - name: OS_PASSWORD
          value: p@ssw0rd
        - name: OS_PROJECT_NAME
          value: demo
        - name: OS_USER_DOMAIN_NAME
          value: Default
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
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list","watch"]
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
        image: registry.opensource.zalan.do/teapot/external-dns
        args:
        - --source=service 
        - --source=ingress 
        - --domain-filter=external-dns-test.my-org.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
        - --provider=designate
        - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
        - --registry=txt
        - --txt-owner-id=my-hostedzone-identifier
        env: # values from openrc file
        - name: OS_AUTH_URL
          value: http://controller/identity/v3
        - name: OS_REGION_NAME
          value: RegionOne
        - name: OS_USERNAME
          value: admin
        - name: OS_PASSWORD
          value: p@ssw0rd
        - name: OS_PROJECT_NAME
          value: demo
        - name: OS_USER_DOMAIN_NAME
          value: Default
```

## Verify ExternalDNS works (Service example)

Create the following sample application to test that ExternalDNS works.

> For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

> If you want to give multiple names to service, you can set it to external-dns.alpha.kubernetes.io/hostname with a comma separator.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: nginx.external-dns-test.my-org.com
spec:
  type: LoadBalancer
  ports:
  - port: 80
    name: http
    targetPort: 80
  selector:
    app: nginx

---

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
          name: http
```

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and notify Designate, which in turn synchronize DNS records with underlying DNS server backend.

## Verifying DNS records

To verify that DNS record was indeed created, you can use the following command:

```console
$ openstack recordset list external-dns-test.my-org.com.
```

There should be a record for my-app.example.com having `ACTIVE` status. And of course, the ultimate method to verify is to issue a DNS query:

```console
$ dig nginx.external-dns-test.my-org.com @8.8.8.8
```

## Clean up

Make sure to delete all Service objects before terminating the cluster so all load balancers get cleaned up correctly.

```console
$ kubectl delete service nginx
```
Give ExternalDNS some time to clean up the DNS records for you. Then delete the DNS zone if you created one for the testing purpose.

```console
$ openstack zone delete external-dns-test.my-org.com. 
```
