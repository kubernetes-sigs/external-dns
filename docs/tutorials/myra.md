# Myra ExternalDNS Webhook

This guide provides quick instructions for setting up and testing the [Myra ExternalDNS Webhook](https://github.com/Myra-Security-GmbH/external-dns-myrasec-webhook) in a Kubernetes environment.

## Prerequisites

- Kubernetes cluster (v1.19+)
- `kubectl` configured to access your cluster
- Docker for building the container image
- MyraSec API credentials (API key and secret)
- Domain registered with MyraSec

## Quick Installation

### 1. Build and Push the Docker Image

```bash
# From the project root
docker build -t myra-webhook:latest .

# Tag the image for your container registry
docker tag myra-webhook:latest YOUR_REGISTRY/myra-webhook:latest

# Push to your container registry
docker push YOUR_REGISTRY/myra-webhook:latest
```

> **Important**: The image must be pushed to a container registry accessible by your Kubernetes cluster. Update the image reference in the deployment YAML file to match your registry path.

### 2. Configure API Credentials

Create a secret with your MyraSec API credentials:

```bash
kubectl create secret generic myra-webhook-secrets \
  --from-literal=myrasec-api-key=YOUR_API_KEY \
  --from-literal=myrasec-api-secret=YOUR_API_SECRET \
  --from-literal=domain-filter=YOUR_DOMAIN.com
```

Alternatively, apply the provided secret template after editing:

```bash
# Edit the secret file first
vi deploy/myra-webhook-secrets.yaml

# Then apply
kubectl apply -f deploy/myra-webhook-secrets.yaml
```

### 3. Deploy the Webhook and ExternalDNS

```bash
# Apply the combined deployment
kubectl apply -f deploy/combined-deployment.yaml
```

This deploys:

- ConfigMap with webhook configuration
- ServiceAccount, ClusterRole, and ClusterRoleBinding for RBAC
- Deployment with two containers:
  - myra-webhook: The webhook provider implementation
  - external-dns: The ExternalDNS controller using the webhook provider

### 4. Verify Deployment

```bash
# Check if pods are running
kubectl get pods -l app=myra-externaldns

# Check logs for the webhook container
kubectl logs -l app=myra-externaldns -c myra-webhook

# Check logs for the external-dns container
kubectl logs -l app=myra-externaldns -c external-dns
```

## Manual Testing with NGINX Demo

### 1. Deploy the NGINX Demo Application

```bash
# Edit the domain in the nginx-demo.yaml file to match your domain
vi deploy/nginx-demo.yaml

# Most important part is to set the correct domain in the external-dns.alpha.kubernetes.io/hostname annotation
# Example:
# annotations:
#   external-dns.alpha.kubernetes.io/enabled: "true"
#   external-dns.alpha.kubernetes.io/hostname: "nginx-demo.dummydomainforkubes.de"
#   external-dns.alpha.kubernetes.io/target: "9.2.3.4"

# Apply the demo resources
kubectl apply -f deploy/nginx-demo.yaml
```

This creates:

- NGINX Deployment
- Service for the deployment
- Ingress resource with ExternalDNS annotations

### 2. Verify DNS Record Creation

After deploying the demo application, ExternalDNS should automatically create DNS records in MyraSec:

```bash
# Check external-dns logs to see record creation
kubectl logs -l app=myra-externaldns -c external-dns | grep "nginx-demo"

# Verify the webhook logs
kubectl logs -l app=myra-externaldns -c myra-webhook | grep "Created DNS record"
```

You can also verify through the MyraSec dashboard that the records were created.

### 3. Testing Record Deletion

To test record deletion:

```bash
# Delete the nginx-demo resources or remove annotation from ingress
kubectl delete -f deploy/nginx-demo.yaml

# Delete the ingress resource or remove annotation from ingress
# If resource is still active, external dns might still see the record and manage it
kubectl delete ingress nginx-demo -n default

# Check external-dns logs to see record deletion
kubectl logs -l app=myra-externaldns -c external-dns | grep "nginx-demo" | grep "delete"

# Verify the webhook logs
kubectl logs -l app=myra-externaldns -c myra-webhook | grep "Deleted DNS record"
```

## Configuration Options

The webhook can be configured through the ConfigMap:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `dry-run` | Run in dry-run mode without making actual changes | `"false"` |
| `environment` | Environment name (affects private IP handling) | `"prod"` |
| `log-level` | Logging level (debug, info, warn, error) | `"debug"` |
| `ttl` | Default TTL for DNS records | `"300"` |
| `webhook-listen-address` | Address and port for the webhook server | `":8080"` |

## Troubleshooting

### Common Issues

1. **Webhook not receiving requests**
   - Ensure the `webhook-provider-url` in the external-dns args is correct
   - Check network connectivity between containers

2. **DNS records not being created**
   - Verify MyraSec API credentials are correct
   - Check if the domain filter is properly configured
   - Look for error messages in the webhook and external-dns logs

3. **Permissions issues**
   - Ensure the ServiceAccount has the correct RBAC permissions

### Getting Help

For more detailed logs:

```bash
# Set log level to debug in the ConfigMap
kubectl edit configmap myra-externaldns-config
# Change log-level to "debug"

# Restart the pods
kubectl rollout restart deployment myra-externaldns
```

## Environment Configuration

The webhook supports different environment configurations through the `environment` setting in the ConfigMap:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: myra-externaldns-config
data:
  environment: "prod"  # Can be "prod", "staging", "dev", etc.
```

The environment setting affects how the webhook handles certain operations:

| Environment | Behavior |
|-------------|----------|
| `prod`, `production`, `staging` | Strict mode: Skips private IP records, enforces stricter validation |
| `dev`, `development`, `test`, etc. | Development mode: Allows private IP records, more permissive validation |

To modify the environment:

```bash
# Edit the ConfigMap directly
kubectl edit configmap myra-externaldns-config

# Or apply an updated YAML file
kubectl apply -f updated-config.yaml
```

## Advanced Configuration

For production deployments, consider:

1. Using a proper image registry instead of `latest` tag
2. Setting resource limits appropriate for your environment
3. Configuring horizontal pod autoscaling
4. Using Helm for deployment management
