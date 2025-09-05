# Managing and Importing Existing DNS Records with ExternalDNS

Sometimes DNS records are created manually (e.g., through Route53, CloudDNS, or AzureDNS), but you still want ExternalDNS to take ownership of them for ongoing management. This tutorial shows how to “import” such records into ExternalDNS by creating the appropriate TXT records.

---

## Prerequisites

- A working Kubernetes cluster
- ExternalDNS installed and configured with your DNS provider
- Manually created DNS records that you want to manage

---

## Example: Importing a Manually Created A Record

Let’s assume you already have the following A record created manually in Route53:

```text
grafana.dev.example.com  → A record → pointing to NLB
```

This entry is referenced in an Istio Gateway resource but was not created by ExternalDNS.

This is how a gateway.yaml file looks like:

```yaml
apiVersion: networking.istio.io/v1
kind: Gateway
metadata:
  name: gateway
  namespace: istio-system
spec:
  selector:
    istio: gateway
  servers:
  - hosts:
    - grafana.dev.example.com
    port:
      name: http
      number: 80
      protocol: HTTP
```

ExternalDNS deployment file example:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: kube-system
spec:
  minReadySeconds: 15
  replicas: 2
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: external-dns
  strategy:
    rollingUpdate:
      maxSurge: 50%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      automountServiceAccountToken: true
      containers:
      - args:
        - --source=service
        - --source=ingress
        - --source=istio-gateway
        - --domain-filter=dev.example.com.
        - --provider=aws
        - --policy=sync
        - --aws-zone-type=private
        - --registry=txt
        - --events
        - --txt-owner-id=dev.example.com
        - --log-level=info
        env:
        - name: AWS_DEFAULT_REGION
          value: us-west-2
        image: registry.k8s.io/external-dns/external-dns:v0.19.0
        imagePullPolicy: IfNotPresent
        name: external-dns
      securityContext:
        fsGroup: 65534
        runAsNonRoot: false
      serviceAccount: external-dns
```

---

## Step 1: Create Corresponding TXT Records

To let ExternalDNS take ownership of the existing A record, you must add TXT records that follow the ExternalDNS format. For example:

```text
aaaa-grafana.dev.example.com  → TXT → "heritage=external-dns,external-dns/owner=dev.example.com,external-dns/resource=gateway/istio/gateway"
cname-grafana.dev.example.com → TXT → "heritage=external-dns,external-dns/owner=dev.example.com,external-dns/resource=gateway/istio/gateway"
```

Note: The easiest way to determine the correct TXT value is to create a dummy record with ExternalDNS. This will generate the required TXT entries, which you can then copy and apply to your manually created records.

These TXT records tell ExternalDNS:

- Which resource owns the record (`external-dns/resource=...`) (in this case, it's istio)
- Which owner identifier is managing it (`external-dns/owner=...`)

---

## Step 2: Verify ExternalDNS Behavior

After creating the TXT records, wait for the next reconciliation loop. You should now see ExternalDNS managing the record without errors.

- With `policy=sync`: if you remove the entry from the Kubernetes resource (e.g., Istio Gateway), ExternalDNS will also remove the corresponding DNS record from your provider.
- With `policy=upsert-only`: ExternalDNS will not delete existing records, even if you remove them from Kubernetes resources.

---

## Notes

- TXT records are required because they serve as ownership markers, preventing conflicts between multiple ExternalDNS controllers.
- This approach is especially useful during migrations, where DNS records pre-exist but you want to avoid downtime or duplication.

---

With this setup, ExternalDNS will manage both newly created and previously existing records in a consistent way.

