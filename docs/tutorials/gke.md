# Setting up ExternalDNS on Google Kubernetes Engine

This tutorial describes how to setup ExternalDNS for usage within a [GKE](https://cloud.google.com/kubernetes-engine) ([Google Kuberentes Engine](https://cloud.google.com/kubernetes-engine)) cluster. Make sure to use **>=0.11.0** version of ExternalDNS for this tutorial

## Single project test scenario using access scopes

*If you prefer to try-out ExternalDNS in one of the existing environments you can skip this step*

The following instructions use [access scopes](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) to provide ExternalDNS with the permissions it needs to manage DNS records within a single [project](https://cloud.google.com/docs/overview#projects), the organizing entity to allocate resources.

Note that since these permissions are associated with the instance, all pods in the cluster will also have these permissions. As such, this approach is not suitable for anything but testing environments.

This solution will only work when both CloudDNS and GKE are provisioned in the same project.  If the CloudDNS zone is in a different project, this solution will not work.

### Configure Project Environment

Setup your environment to work with Google Cloud Platform. Fill in your variables as needed, e.g. target project.

```bash
# set variables to the appropriate desired values
PROJECT_ID="my-external-dns-test"
REGION="europe-west1"
ZONE="europe-west1-d"
ClOUD_BILLING_ACCOUNT="<my-cloud-billing-account>"
# set default settings for project
gcloud config set project $PROJECT_ID
gcloud config set compute/region $REGION
gcloud config set compute/zone $ZONE
# enable billing and APIs if not done already
gcloud beta billing projects link $PROJECT_ID \
  --billing-account $BILLING_ACCOUNT
gcloud services enable "dns.googleapis.com"
gcloud services enable "container.googleapis.com"
```

### Create GKE Cluster

```bash
gcloud container clusters create $GKE_CLUSTER_NAME \
  --num-nodes 1 \
  --scopes "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
```

**WARNING**: Note that this cluster will use the default [compute engine GSA](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) that contians the overly permissive project editor (`roles/editor`) role. So essentially, anything on the cluster could potentially grant escalated privileges.  Also, as mentioned earlier, the access scope `ndev.clouddns.readwrite` will allow anything running on the cluster to have read/write permissions on all Cloud DNS zones within the same project.

### Cloud DNS Zone

Create a DNS zone which will contain the managed DNS records. If using your own domain that was registered with a third-party domain registrar, you should point your domain's name servers to the values under the `nameServers` key. Please consult your registrar's documentation on how to do that.  This tutorial will use example domain of  `example.com`.

```bash
gcloud dns managed-zones create "example-com" --dns-name "example.com." \
  --description "Automatically managed zone by kubernetes.io/external-dns"
```

Make a note of the nameservers that were assigned to your new zone.

```bash
gcloud dns record-sets list \
    --zone "example-com" --name "example.com." --type NS
```

Outputs:

```
NAME          TYPE  TTL    DATA
example.com.  NS    21600  ns-cloud-e1.googledomains.com.,ns-cloud-e2.googledomains.com.,ns-cloud-e3.googledomains.com.,ns-cloud-e4.googledomains.com.
```

In this case it's `ns-cloud-{e1-e4}.googledomains.com.` but your's could slightly differ, e.g. `{a1-a4}`, `{b1-b4}` etc.

## Cross project access scenario using Google Service Account

More often, following best practices in regards to security and operations, Cloud DNS zones will be managed in a separate project from the Kubernetes cluster.  This section shows how setup ExternalDNS to access Cloud DNS from a different project. These steps will also work for single project scenarios as well.

ExternalDNS will need permissions to make changes to the Cloud DNS zone. There are three ways to configure the access needed:

* [Worker Node Service Account](#worker-node-service-account)
* [Static Credentials](#static-credentials)
* [Work Load Identity](#work-load-identity)

### Setup Cloud DNS and GKE

Below are examples on how you can configure Cloud DNS and GKE in separate projects, and then use one of the three methods to grant access to ExternalDNS.  Replace the environment variables to values that make sense in your environment.

#### Configure Projects

For this process, create projects with the appropriate APIs enabled.

```bash
# set variables to appropriate desired values
GKE_PROJECT_ID="my-workload-project"
DNS_PROJECT_ID="my-cloud-dns-project"
ClOUD_BILLING_ACCOUNT="<my-cloud-billing-account>"
# enable billing and APIs for DNS project if not done already
gcloud config set project $DNS_PROJECT_ID
gcloud beta billing projects link $CLOUD_DNS_PROJECT \
  --billing-account $ClOUD_BILLING_ACCOUNT
gcloud services enable "dns.googleapis.com"
# enable billing and APIs for GKE project if not done already
gcloud config set project $GKE_PROJECT_ID
gcloud beta billing projects link $CLOUD_DNS_PROJECT \
  --billing-account $ClOUD_BILLING_ACCOUNT
gcloud services enable "container.googleapis.com"
```

#### Provisioning Cloud DNS

Create a Cloud DNS zone in the designated DNS project.  

```bash
gcloud dns managed-zones create "example-com" --project $DNS_PROJECT_ID \
  --description "example.com" --dns-name="example.com." --visibility=public
```

If using your own domain that was registered with a third-party domain registrar, you should point your domain's name servers to the values under the `nameServers` key.  Please consult your registrar's documentation on how to do that. The example domain of `example.com` will be used for this tutorial.

#### Provisioning a GKE cluster for cross project access

Create a GSA (Google Service Account) and grant it the [minimal set of privileges required](https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#use_least_privilege_sa) for GKE nodes:

```bash
GKE_CLUSTER_NAME="my-external-dns-cluster"
GKE_REGION="us-central1"
GKE_SA_NAME="worker-nodes-sa"
GKE_SA_EMAIL="$GKE_SA_NAME@${GKE_PROJECT_ID}.iam.gserviceaccount.com"

ROLES=(
  roles/logging.logWriter
  roles/monitoring.metricWriter
  roles/monitoring.viewer
  roles/stackdriver.resourceMetadata.writer
)

gcloud iam service-accounts create $GKE_SA_NAME \
  --display-name $GKE_SA_NAME --project $GKE_PROJECT_ID

# assign google service account to roles in GKE project
for ROLE in ${ROLES[*]}; do
  gcloud projects add-iam-policy-binding $GKE_PROJECT_ID \
    --member "serviceAccount:$GKE_SA_EMAIL" \
    --role $ROLE
done
```

Create a cluster using this service account and enable [workload identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity):

```bash
gcloud container clusters create $GKE_CLUSTER_NAME \
  --project $GKE_PROJECT_ID --region $GKE_REGION --num-nodes 1 \
  --service-account "$GKE_SA_EMAIL" \
  --workload-pool "$GKE_PROJECT_ID.svc.id.goog"
```

### Worker Node Service Account method

In this method, the GSA (Google Service Account) that is associated with GKE worker nodes will be configured to have access to Cloud DNS.  

**WARNING**: This will grant access to modify the Cloud DNS zone records for all containers running on cluster, not just ExternalDNS, so use this option with caution.  This is not recommended for production environments.

```bash
GKE_SA_EMAIL="$GKE_SA_NAME@${GKE_PROJECT_ID}.iam.gserviceaccount.com"

# assign google service account to dns.admin role in the cloud dns project
gcloud projects add-iam-policy-binding $DNS_PROJECT_ID \
  --member serviceAccount:$GKE_SA_EMAIL \
  --role roles/dns.admin
```

After this, follow the steps in [Deploy ExternalDNS](#deploy-externaldns).  Make sure to set the `--google-project` flag to match the Cloud DNS project name.

### Static Credentials

In this scenario, a new GSA (Google Service Account) is created that has access to the CloudDNS zone.  The credentials for this GSA are saved and installed as a Kubernetes secret that will be used by ExternalDNS.  

This allows only containers that have access to the secret, such as ExternalDNS to update records on the Cloud DNS Zone.

#### Create GSA for use with static credentials

```bash
DNS_SA_NAME="external-dns-sa"
DNS_SA_EMAIL="$DNS_SA_NAME@${GKE_PROJECT_ID}.iam.gserviceaccount.com"

# create GSA used to access the Cloud DNS zone
gcloud iam service-accounts create $DNS_SA_NAME --display-name $DNS_SA_NAME

# assign google service account to dns.admin role in cloud-dns project
gcloud projects add-iam-policy-binding $DNS_PROJECT_ID \
  --member serviceAccount:$DNS_SA_EMAIL --role "roles/dns.admin"
```

#### Create Kubernetes secret using static credentials

Generate static credentials from the ExternalDNS GSA.

```bash
# download static credentials
gcloud iam service-accounts keys create /local/path/to/credentials.json \
  --iam-account $DNS_SA_EMAIL
```

Create a Kubernetes secret with the credentials in the same namespace of ExternalDNS.

```bash
kubectl create secret generic "external-dns" --namespace ${EXTERNALDNS_NS:-"default"} \
  --from-file /local/path/to/credentials.json
```

After this, follow the steps in [Deploy ExternalDNS](#deploy-externaldns).  Make sure to set the `--google-project` flag to match Cloud DNS project name. Make sure to uncomment out the section that mounts the secret to the ExternalDNS pods.
### Workload Identity

[Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) allows workloads in your GKE cluster to impersonate GSA (Google Service Accounts) using KSA (Kubernetes Service Accounts) configured during deployemnt.  These are the steps to use this feature with ExternalDNS.

#### Create GSA for use with Workload Identity

```bash
DNS_SA_NAME="external-dns-sa"
DNS_SA_EMAIL="$DNS_SA_NAME@${GKE_PROJECT_ID}.iam.gserviceaccount.com"

gcloud iam service-accounts create $DNS_SA_NAME --display-name $DNS_SA_NAME
gcloud projects add-iam-policy-binding $DNS_PROJECT_ID \
   --member serviceAccount:$DNS_SA_EMAIL --role "roles/dns.admin"
```

#### Link KSA to GSA

Add an IAM policy binding bewtween the workload identity GSA and ExternalDNS GSA.  This will link the ExternalDNS KSA to ExternalDNS GSA.

```bash
gcloud iam service-accounts add-iam-policy-binding $DNS_SA_EMAIL \
  --role "roles/iam.workloadIdentityUser" \
  --member "serviceAccount:$GKE_PROJECT_ID.svc.id.goog[${EXTERNALDNS_NS:-"default"}/external-dns]"
```

#### Deploy External DNS

Deploy ExternalDNS with the following steps below, documented under [Deploy ExternalDNS](#deploy-externaldns).  Set the `--google-project` flag to the Cloud DNS project name.

#### Link KSA to GSA in Kubernetes

Add the proper workload identity annotation to the ExternalDNS KSA.

```bash
kubectl annotate serviceaccount "external-dns" \
  --namespace ${EXTERNALDNS_NS:-"default"} \
  "iam.gke.io/gcp-service-account=$DNS_SA_EMAIL"
```

#### Update ExternalDNS pods

Update the Pod spec to schedule the workloads on nodes that use Workload Identity and to use the annotated Kubernetes service account.

```bash
kubectl patch deployment "external-dns" \
  --namespace ${EXTERNALDNS_NS:-"default"} \
  --patch \
 '{"spec": {"template": {"spec": {"nodeSelector": {"iam.gke.io/gke-metadata-server-enabled": "true"}}}}}'
```

After all of these steps you may see several messages with `googleapi: Error 403: Forbidden, forbidden`.  After several minutes when the token is refreshed, these error messages will go away, and you should see info messages, such as: `All records are already up to date`.

## Deploy ExternalDNS

Then apply the following manifests file to deploy ExternalDNS.

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  labels:
    app.kubernetes.io/name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
  labels:
    app.kubernetes.io/name: external-dns
rules:
  - apiGroups: [""]
    resources: ["services","endpoints","pods","nodes"]
    verbs: ["get","watch","list"]
  - apiGroups: ["extensions","networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get","watch","list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns-viewer
  labels:
    app.kubernetes.io/name: external-dns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
  - kind: ServiceAccount
    name: external-dns
    namespace: default # change if namespace is not 'default'
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  labels:
    app.kubernetes.io/name: external-dns  
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app.kubernetes.io/name: external-dns
  template:
    metadata:
      labels:
        app.kubernetes.io/name: external-dns
    spec:
      serviceAccountName: external-dns
      containers:
        - name: external-dns
          image: k8s.gcr.io/external-dns/external-dns:v0.11.0
          args:
            - --source=service
            - --source=ingress
            - --domain-filter=example.com # will make ExternalDNS see only the hosted zones matching provided domain, omit to process all available hosted zones
            - --provider=google
            - --log-format=json # google cloud logs parses severity of the "text" log format incorrectly
    #        - --google-project=my-cloud-dns-project # Use this to specify a project different from the one external-dns is running inside
            - --google-zone-visibility=public # Use this to filter to only zones with this visibility. Set to either 'public' or 'private'. Omitting will match public and private zones
            - --policy=upsert-only # would prevent ExternalDNS from deleting any records, omit to enable full synchronization
            - --registry=txt
            - --txt-owner-id=my-identifier
      #     # uncomment below if static credentials are used  
      #     env:
      #       - name: GOOGLE_APPLICATION_CREDENTIALS
      #         value: /etc/secrets/service-account/credentials.json
      #     volumeMounts:
      #       - name: google-service-account
      #         mountPath: /etc/secrets/service-account/
      # volumes:
      #   - name: google-service-account
      #     secret:
      #       secretName: external-dns
```

Create the deployment for ExternalDNS:

```bash
kubectl create --namespace "default" --filename externaldns.yaml
```

## Verify ExternalDNS works

The following will deploy a small nginx server that will be used to demonstrate that ExternalDNS is working.

### Verify using an external load balancer

Create the following sample application to test that ExternalDNS works.  This example will provision a L4 load balancer.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    # change nginx.example.com to match an appropriate value
    external-dns.alpha.kubernetes.io/hostname: nginx.example.com
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: nginx
---
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
```

Create the deployment and service objects:

```bash
kubectl create --namespace "default" --filename nginx.yaml
```

After roughly two minutes check that a corresponding DNS record for your service was created.

```bash
gcloud dns record-sets list --zone "example-com" --name "nginx.example.com."
```

Example output:

```
NAME                TYPE  TTL  DATA
nginx.example.com.  A     300  104.155.60.49
nginx.example.com.  TXT   300  "heritage=external-dns,external-dns/owner=my-identifier"
```

Note created `TXT` record alongside `A` record. `TXT` record signifies that the corresponding `A` record is managed by ExternalDNS. This makes ExternalDNS safe for running in environments where there are other records managed via other means.

Let's check that we can resolve this DNS name. We'll ask the nameservers assigned to your zone first.

```bash
dig +short @ns-cloud-e1.googledomains.com. nginx.example.com.
104.155.60.49
```

Given you hooked up your DNS zone with its parent zone you can use `curl` to access your site.

```bash
curl nginx.example.com
```

### Verify using an ingress

Let's check that Ingress works as well. Create the following Ingress.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx
spec:
  rules:
    - host: server.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: nginx-svc
                port:
                  number: 80
```

Create the ingress objects with:

```bash
kubectl create --namespace "default" --filename ingress.yaml
```

Note that this will ingress object will use the default ingress controller that comes with GKE to create a L7 load balancer in addition to the L4 load balancer previously with the service object.  To use only the L7 load balancer, update the service manafest to change the Service type to `NodePort` and remove the ExternalDNS annotation.

After roughly two minutes check that a corresponding DNS record for your Ingress was created.

```bash
gcloud dns record-sets list \
    --zone "example-com" \
    --name "server.example.com." \
```
Output:

```
NAME                 TYPE  TTL  DATA
server.example.com.  A     300  130.211.46.224
server.example.com.  TXT   300  "heritage=external-dns,external-dns/owner=my-identifier"
```

Let's check that we can resolve this DNS name as well.

```bash
dig +short @ns-cloud-e1.googledomains.com. server.example.com.
130.211.46.224
```

Try with `curl` as well.

```bash
curl server.example.com
```

### Clean up

Make sure to delete all Service and Ingress objects before terminating the cluster so all load balancers get cleaned up correctly.

```bash
kubectl delete service nginx
kubectl delete ingress nginx
```

Give ExternalDNS some time to clean up the DNS records for you. Then delete the managed zone and cluster.

```bash
gcloud dns managed-zones delete "example-com"
gcloud container clusters delete "external-dns"
```
