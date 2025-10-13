# Anexia Engine

This tutorial describes how to set up ExternalDNS for use within a Kubernetes cluster using Anexia Cloud DNS.
To view the source repository and the container image visit the [GitHub source of anexia/k8s-external-dns-webhook](https://github.com/anexia/k8s-external-dns-webhook).

## Creating a DNS Zone with Anexia Engine

Make sure to familiarize yourself with Anexia CloudDNS and read the following documentation to know how to create a new DNS Zone:

- [Setting up CloudDNS](https://engine.anexia-it.com/docs/en/module/clouddns/getting-started/setting-up-clouddns)

### Adding a DNS Zone

1. Log in to [Anexia Engine](https://engine.anexia-it.com/).
2. Navigate to the **Domains & DNS** section and select **CloudDNS**.
3. Click on **+ Create Zone** and provide the following details:
   - **Zone Name**: Enter the domain name (e.g., `anexia-app.com`).
   - Continue through the wizard and complete the setup
4. Save the zone configuration.

> Advanced configuration can be located at the [Anexia CloudDNS Documentation](https://engine.anexia-it.com/docs/en/module/clouddns/overview)

## Creating an Anexia API Token

To use ExternalDNS with CloudDNS, you need an API token to manage DNS zones and records.
For production usage it is advised to use a **service account** to manage DNS records.
View the official [documentation how to create a service account and retrieve the respective API token](https://engine.anexia-it.com/docs/en/module/engine-core/getting-started/managing-users-and-teams/managing-users).

Once the user is created the retrieval of the API token itself is very similar to retrieving it from your personal user:

1. Log in to [Anexia Engine](https://engine.anexia-it.com/).
2. Navigate to your profile in the top right corner and select **API**.
3. In the API section, tick the box at **Token Authentication** and your token will be generated.
4. Copy the token and save it for later use with ExternalDNS.

## Deploy ExternalDNS

### Step 1: Create a Kubernetes Secret for the Anexia API Token

Create a Kubernetes secret to store your token inside Kubernetes.

```bash
export ANX_TOKEN='<ANEXIA_API_TOKEN>' # store the token in the local environment, for later use
kubectl create secret generic anexia-credentials --from-literal=token="${ANX_TOKEN}"
```

Replace `<ANEXIA_API_TOKEN>` with your actual API token.

### Step 2: Configure ExternalDNS

Create a Helm values file for the ExternalDNS Helm chart that includes the webhook configuration. In this example, the values file is called `external-dns-anexia-values.yaml` .

```yaml
# -- ExternalDNS Log level.
logLevel: debug # reduce in production

# -- if true, _ExternalDNS_ will run in a namespaced scope (Role and Rolebinding will be namespaced too).
namespaced: false

# -- _Kubernetes_ resources to monitor for DNS entries.
sources:
  - ingress
  - service
  - crd

extraArgs:
  ## must override the default value with port 8888 with port 8080 because this is hard-coded in the helm chart
  - --webhook-provider-url=http://localhost:8080
  ## You should filter the domains that can be requested to limit the amount of requests done to the anexia engine.
  ## This will help you avoid running into rate limiting
  - --domain-filter=web-demo.anexia-app.com

provider:
  name: webhook
  webhook:
    image:
      repository: ghcr.io/anexia/k8s-external-dns-webhook
      tag: v0.2.4
    env:
      - name: LOG_LEVEL
        value: debug # reduce in production
      - name: ANEXIA_API_TOKEN
        valueFrom:
          secretKeyRef:
            name: anexia-credentials
            key: token
      - name: SERVER_HOST
        value: "0.0.0.0"
      - name: SERVER_PORT
        value: "8080"
      - name: DRY_RUN
        value: "false"
```

### Step 3: Install ExternalDNS Using Helm

Install ExternalDNS with the Anexia webhook provider:

```bash
helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
helm upgrade -i external-dns-anexia external-dns/external-dns -f external-dns-anexia-values.yaml
```

## Deploying an Example Application

### Step 1: Create a Deployment

Create a sample deployment named `web-demo`.

```bash
kubectl create deployment web-demo --image=gcr.io/google-samples/hello-app:2.0
```

### Step 2: Create a Service

Create a `Service` manifest to expose the `hello-app` application within the cluster.
It's important to include the annotation for ExternalDNS to create DNS records for the specified hostname.

Save the following content in a file named `web-demo-service.yaml`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: web-demo
  annotations:
    external-dns.alpha.kubernetes.io/hostname: web-demo.anexia-app.com
spec:
  ports:
    - port: 80
      targetPort: 8080
  selector:
    app: web-demo
```

 > Replace `web-demo.anexia-app.com` with a subdomain of your DNS zone configured in Anexia CloudDNS.

To create the service apply it with the following command:

```bash
kubectl apply -f web-demo-service.yaml
```

The service is going to expose the application on port 80.

### Step 3: Create an Ingress

Create an `Ingress` resource to expose the `hello-app` application externally.
The ingress will route HTTP traffic to the `hello-app` service and include a hostname that ExternalDNS will use to create the corresponding DNS record.

Save the following content in a file named `web-demo-ingress.yaml` :

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: web-demo
spec:
  rules:
  - host: web-demo.anexia-app.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: web-demo
            port:
              number: 8080
```

Make sure to have an ingress controller running inside your cluster, otherwise it will not work.

> Replace `web-demo.anexia-app.com` with a subdomain of your DNS zone configured in Anexia CloudDNS.
> Make sure to **NOT** create a DNS record yourself at first.
> ExternalDNS will not change records, which were not created by it.

Next, apply the ingress manifest:

```bash
kubectl apply -f web-demo-ingress.yaml
```

This ingress will expose the `hello-app` application at `http://web-demo.anexia-app.com` and instruct ExternalDNS to create a DNS record for the specified hostname.

### Verifying everything worked

Verification if the DNS records were created can be done in multiple ways:
* Looking it up in the Anexia Engine web interface
* Using the Anexia API
* Querying the Anexia DNS server

#### Using the Anexia API

The following `curl` command can be used to check if the records were already created in your DNS zone.
For this, you have to find your `zone ID` first, either through the web interface or API call.
While editing your zone, the `zone ID` will be in the URL of your browser: `https://engine.anexia-it.com/clouddns/zone/view/<ZONE_ID>`

Use the following `curl` command to query the Anexia Engine:

```bash
# get the ZONE_ID via the api
# list all your zones and locate your desired zone -> look up the field 'identifier' in the JSON
curl --header "Authorization: Token ${ANX_TOKEN}" https://engine.anexia-it.com/api/clouddns/v1/zone.json | jq

export ZONE_ID='<ZONE_ID>' # set your zone id first
curl --header "Authorization: Token ${ANX_TOKEN}" https://engine.anexia-it.com/api/clouddns/v1/zone.json/$ZONE_ID/records | jq
```

It should list all records in a JSON format.
Note that there is an `A` record and a `TXT` record created.
The `TXT` record is used internally for ExternalDNS to show that this record is managed by ExternalDNS.

> If you don't have the tool `jq` installed, this will lead to an error.
> You can omit the last part of the command then, but no pretty printing will be used.

#### Using DNS Query

To query DNS records, use the following `dig` command:

```bash
dig @acns01.xaas.systems +answer +multiline web-demo.anexia-app.com any
```

It should list the records in a text format.

### Accessing your app

Using `curl` you can access your application now.

```bash
curl http://web-demo.anexia-app.com
```

You should see something like this as response:

```bash
Hello, world!
Version: 2.0.0
Hostname: web-demo-cb7595bfd-82htv
```

> **Troubleshooting:**
>
>If you encounter any issues, verify the following:
>
> - The DNS record for `web-demo.anexia-app.com` has been created in Anexia CloudDNS.
> - The ingress controller is running and properly configured in your Kubernetes cluster.

## Cleanup

> **Optional:** Only cleanup your changes if they are no longer needed.

```bash
kubectl delete deployment/web-demo
kubectl delete -f web-demo-service.yaml
kubectl delete -f web-demo-ingress.yaml
```
