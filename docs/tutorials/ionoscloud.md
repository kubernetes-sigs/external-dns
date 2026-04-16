# IONOS Cloud

This tutorial describes how to set up ExternalDNS for use within a Kubernetes cluster using IONOS Cloud DNS.
For more details, visit the [IONOS external-dns webhook repository](https://github.com/ionos-cloud/external-dns-ionos-webhook).
You can also find the [external-dns-ionos-webhook container image](https://github.com/ionos-cloud/external-dns-ionos-webhook/pkgs/container/external-dns-ionos-webhook) required for this setup.

## Creating a DNS Zone with IONOS Cloud DNS

If you are new to IONOS Cloud DNS, we recommend you first read the following instructions for creating a DNS zone:

- [Manage DNS Zones in Data Centre Designer](https://docs.ionos.com/cloud/network-services/cloud-dns/dcd-how-tos/manage-dns-zone)
- [Creating a DNS Zone using the IONOS Cloud DNS API](https://docs.ionos.com/cloud/network-services/cloud-dns/api-how-tos/create-dns-zone)

### Steps to Create a DNS Zone

1. Log in to the [IONOS Cloud Data Center Designer](https://dcd.ionos.com/).
2. Navigate to the **Network Services** section and select **Cloud DNS**.
3. Click on **Create Zone** and provide the following details:
   - **Zone Name**: Enter the domain name (e.g., `example.com`).
   - **Description**: It is optional to provide a description of your zone.
4. Save the zone configuration.

For more advanced configurations, such as adding records or managing subdomains, refer to the [IONOS Cloud DNS Documentation](https://docs.ionos.com/cloud/network-services/cloud-dns/).

## Creating an IONOS API Token

To use ExternalDNS with IONOS Cloud DNS, you need an API token with sufficient privileges to manage DNS zones and records. Follow these steps to create an API token:

1. Log in to the [IONOS Cloud Data Center Designer](https://dcd.ionos.com/).
2. Navigate to the **Management** section in the top right corner and select **Token Manager**.
3. Select the Time To Live(TTL) of the token and click on **Create Token**.
4. Copy the generated token and store it securely. You will use this token to authenticate ExternalDNS.

## Deploy ExternalDNS

### Step 1: Create a Kubernetes Secret for the IONOS API Token

Store your IONOS API token securely in a Kubernetes secret:

```bash
kubectl create secret generic ionos-credentials --from-literal=api-key='<IONOS_API_TOKEN>'
```

Replace `<IONOS_API_TOKEN>` with your actual IONOS API token.

### Step 2: Configure ExternalDNS

Create a Helm values file for the ExternalDNS Helm chart that includes the webhook configuration. In this example, the values file is called `external-dns-ionos-values.yaml` .

```yaml
logLevel: debug # ExternalDNS Log level, reduce in production

namespaced: false # if true, ExternalDNS will run in a namespaced scope (Role and Rolebinding will be namespaced too).
triggerLoopOnEvent: true # if true, ExternalDNS will trigger a loop on every event (create/update/delete) on the resources it watches.

logLevel: debug
sources:
  - ingress
  - service
provider:
  name: webhook
  webhook:
    image:
      repository: ghcr.io/ionos-cloud/external-dns-ionos-webhook
      tag: latest
      pullPolicy: IfNotPresent
    env:
      - name: IONOS_API_KEY
        valueFrom:
          secretKeyRef:
            name: ionos-credentials
            key: api-key
      - name: SERVER_PORT
        value: "8888"
      - name: METRICS_PORT
        value: "8080"
      - name: DRY_RUN
        value: "false"
```

### Step 3: Install ExternalDNS Using Helm

Install ExternalDNS with the IONOS webhook provider:

```bash
helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
helm upgrade --install external-dns external-dns/external-dns -f external-dns-ionos-values.yaml
```

## Deploying an Example Application

### Step 1: Create a Deployment

In this step we will create `echoserver` application manifest with the following content:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echoserver
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: echoserver
  template:
    metadata:
      labels:
        app: echoserver
    spec:
      containers:
      - name: echoserver
        image: ealen/echo-server:latest
        ports:
        - containerPort: 80
```

Deployment manifest can be saved in `echoserver-deployment.yaml` file.

Next, we will apply the deployment:

```bash
kubectl apply -f echoserver-deployment.yaml
```

### Step 2: Create a Service

In this step, we will create a `Service` manifest to expose the `echoserver` application within the cluster. The service will also include an annotation for ExternalDNS to create a DNS record for the specified hostname.

Save the following content in a file named `echoserver-service.yaml`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: echoserver
  annotations:
    external-dns.alpha.kubernetes.io/hostname: app.example.com
spec:
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: echoserver
```

 **Note:** Replace `app.example.com` with a subdomain of your DNS zone configured in IONOS Cloud DNS. For example, if your DNS zone is `example.com`, you can use a subdomain like `app.example.com`.

Next, apply the service:

```bash
kubectl apply -f echoserver-service.yaml
```

This service will expose the echoserver application on port 80 and instruct ExternalDNS to create a DNS record for `app.example.com`.

### Step 3: Create an Ingress

In this step, we will create an `Ingress` resource to expose the `echoserver` application externally. The ingress will route HTTP traffic to the `echoserver` service and include a hostname that ExternalDNS will use to create the corresponding DNS record.

Save the following content in a file named `echoserver-ingress.yaml` :

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: echoserver
spec:
  rules:
  - host: app.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: echoserver
            port:
              number: 80
```

 **Note:** Replace `app.example.com` with a subdomain of your DNS zone configured in IONOS Cloud DNS. For example, if your DNS zone is `example.com`, you can use a subdomain like `app.example.com`.

Next, apply the ingress manifest:

```bash
kubectl apply -f echoserver-ingress.yaml
```

This ingress will expose the `echoserver` application at `http://app.example.com` and instruct ExternalDNS to create a DNS record for the specified hostname.

## Accessing the Application

Once the `Ingress` resource has been applied and the DNS records have been created, you can access the application using the hostname specified in the ingress (`app.example.com`).

### Verify Application Access

Use the following `curl` command to verify that the application is accessible:

```bash
curl -I http://app.example.com
```

Replace app.example.com with the subdomain you configured in your DNS zone.

 **Note:** Ensure that your DNS changes have propagated and that the hostname resolves to the correct IP address before running the command.

### Expected result

You should see an HTTP response header indicating that the application is running, such as:

```bash
HTTP/1.1 200 OK
```

> **Troubleshooting:**
>
>If you encounter any issues, verify the following:
>
> - The DNS record for `app.example.com` (replace with your own subdomain configured in IONOS Cloud DNS) has been created in IONOS Cloud DNS.
> - The ingress controller is running and properly configured in your Kubernetes cluster.
> - The `echoserver` application is running and accessible within the cluster.

## Verifying IONOS Cloud DNS Records

Use the IONOS Cloud Console or API to verify that the A and TXT records for your domain have been created. For example, you can use the following API call:

```bash
curl --location --request GET 'https://dns.de-fra.ionos.com/records?filter.name=app' \
--header 'Authorization: Bearer <IONOS_API_TOKEN>'
```

Replace `<IONOS_API_TOKEN>` with your actual API token.

The API response should include the `A` and `TXT` records for the subdomain you configured.

> **Note:** DNS changes may take a few minutes to propagate. If the records are not visible immediately, wait and try again.

## Cleanup

> **Optional:** Perform the cleanup step only if you no longer need the deployed resources.

Once you have verified the setup, you can clean up the resources created during this tutorial:

```bash
kubectl delete -f echoserver-deployment.yaml
kubectl delete -f echoserver-service.yaml
kubectl delete -f echoserver-ingress.yaml
```

## Summary

In this tutorial, you successfully deployed ExternalDNS webhook with IONOS Cloud DNS as the provider.
You created a Kubernetes deployment, service, and ingress, and verified that DNS records were created and the application was accessible.
You also learned how to clean up the resources when they are no longer needed.
