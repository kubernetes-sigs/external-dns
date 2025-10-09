# Split Horizon DNS

Split horizon DNS allows you to serve different DNS responses based on the client's location - internal clients receive private IPs while external clients receive public IPs. External-DNS supports split horizon DNS by running multiple instances with different annotation prefixes.

## Overview

By default, all external-dns instances use the same annotation prefix: `external-dns.alpha.kubernetes.io/`. This means all instances process the same annotations. To enable split horizon DNS, you can configure each instance to use a different annotation prefix via the `--annotation-prefix` flag.

## Use Cases

- **Internal/External separation**: Internal DNS points to private IPs (ClusterIP), external DNS points to public Load Balancer IPs
- **Multiple DNS providers**: Route different services to different DNS providers (e.g., internal to CoreDNS, external to Route53)
- **Geographic split**: Different DNS records for different regions

## Configuration

### Basic Split Horizon Setup

**Internal DNS Instance:**

```bash
external-dns \
  --annotation-prefix=internal.company.io/ \
  --source=service \
  --source=ingress \
  --provider=aws \
  --aws-zone-type=private \
  --domain-filter=internal.company.com \
  --txt-owner-id=internal-dns
```

**External DNS Instance:**

```bash
external-dns \
  --annotation-prefix=external-dns.alpha.kubernetes.io/ \  # default, can be omitted
  --source=service \
  --source=ingress \
  --provider=aws \
  --aws-zone-type=public \
  --domain-filter=company.com \
  --txt-owner-id=external-dns
```

### Service with Both Annotations

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp
  annotations:
    # Internal DNS reads this
    internal.company.io/hostname: myapp.internal.company.com
    internal.company.io/ttl: "300"
    internal.company.io/target: 10.0.1.50  # Private IP

    # External DNS reads this
    external-dns.alpha.kubernetes.io/hostname: myapp.company.com
    external-dns.alpha.kubernetes.io/ttl: "60"
    # No target = uses LoadBalancer IP automatically
spec:
  type: LoadBalancer
  clusterIP: 10.0.1.50
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: myapp
```

**Result:**

- **Internal DNS** (Route53 Private Zone `internal.company.com`): `myapp.internal.company.com → 10.0.1.50`
- **External DNS** (Route53 Public Zone `company.com`): `myapp.company.com → 203.0.113.10` (LoadBalancer IP)

### Helm Chart Configuration

You can use the Helm chart to deploy multiple instances:

**values-internal.yaml:**

```yaml
annotationPrefix: "internal.company.io/"

provider:
  name: aws

aws:
  zoneType: private

domainFilters:
  - internal.company.com

txtOwnerId: internal-dns

sources:
  - service
  - ingress
```

**values-external.yaml:**

```yaml
# annotationPrefix defaults to "external-dns.alpha.kubernetes.io/"
# can be omitted or set explicitly:
# annotationPrefix: "external-dns.alpha.kubernetes.io/"

provider:
  name: aws

aws:
  zoneType: public

domainFilters:
  - company.com

txtOwnerId: external-dns

sources:
  - service
  - ingress
```

**Deploy:**

```bash
# Internal instance
helm install external-dns-internal external-dns/external-dns \
  --namespace external-dns-internal \
  --create-namespace \
  --values values-internal.yaml

# External instance
helm install external-dns-external external-dns/external-dns \
  --namespace external-dns-external \
  --create-namespace \
  --values values-external.yaml
```

## Advanced Examples

### Three-Way Split (Internal / DMZ / External)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  annotations:
    # Internal (private network only)
    internal.company.io/hostname: api.internal.company.com
    internal.company.io/ttl: "300"

    # DMZ (accessible from office network)
    dmz.company.io/hostname: api.dmz.company.com
    dmz.company.io/ttl: "120"

    # External (public internet)
    external-dns.alpha.kubernetes.io/hostname: api.company.com
    external-dns.alpha.kubernetes.io/ttl: "60"
    external-dns.alpha.kubernetes.io/cloudflare-proxied: "true"
spec:
  type: LoadBalancer
  # ...
```

**Deploy three instances:**

```bash
# Internal
--annotation-prefix=internal.company.io/ --provider=aws --aws-zone-type=private

# DMZ
--annotation-prefix=dmz.company.io/ --provider=aws --aws-zone-type=private

# External
--annotation-prefix=external-dns.alpha.kubernetes.io/ --provider=cloudflare
```

### Different Providers Per Instance

```yaml
apiVersion: v1
kind: Service
metadata:
  name: webapp
  annotations:
    # Route53 for AWS internal
    aws.company.io/hostname: webapp.aws.company.com
    aws.company.io/aws-alias: "true"

    # Cloudflare for public
    cf.company.io/hostname: webapp.company.com
    cf.company.io/cloudflare-proxied: "true"
spec:
  type: LoadBalancer
  # ...
```

**Deploy:**

```bash
# AWS instance
--annotation-prefix=aws.company.io/ --provider=aws

# Cloudflare instance
--annotation-prefix=cf.company.io/ --provider=cloudflare
```

## Important Notes

1. **Annotation prefix must end with `/`** - The validation will fail if the prefix doesn't end with a forward slash.

2. **Backward compatibility** - If you don't specify `--annotation-prefix`, the default `external-dns.alpha.kubernetes.io/` is used, maintaining full backward compatibility.

3. **All annotations use the same prefix** - When you set a custom prefix, ALL external-dns annotations (hostname, ttl, target, cloudflare-proxied, etc.) must use that prefix.

4. **TXT ownership records** - Each instance should have a unique `--txt-owner-id` to avoid conflicts in ownership tracking.

5. **Provider-specific annotations** - Provider-specific annotations (like `cloudflare-proxied`, `aws-alias`) also use the custom prefix:

   ```yaml
   custom.io/hostname: example.com
   custom.io/cloudflare-proxied: "true"  # NOT external-dns.alpha.kubernetes.io/cloudflare-proxied
   ```

## Troubleshooting

### Both instances processing the same resources

**Problem:** Both internal and external instances are creating records for the same service.

**Solution:** Make sure you're using different annotation prefixes and that your services have the correct annotations:

```yaml
# ✅ Correct - different prefixes
internal.company.io/hostname: internal.example.com
external-dns.alpha.kubernetes.io/hostname: example.com

# ❌ Wrong - same prefix
external-dns.alpha.kubernetes.io/hostname: internal.example.com
external-dns.alpha.kubernetes.io/hostname: example.com  # Second one overwrites first
```

### Validation error: "annotation-prefix must end with '/'"

**Problem:** The annotation prefix doesn't end with a forward slash.

**Solution:** Always end your custom prefix with `/`:

```bash
# ✅ Correct
--annotation-prefix=custom.io/

# ❌ Wrong
--annotation-prefix=custom.io
```

### Provider-specific annotations not working

**Problem:** Cloudflare/AWS-specific annotations are not being applied.

**Solution:** Provider-specific annotations must use the same prefix as the hostname:

```yaml
# If using custom prefix
custom.io/hostname: example.com
custom.io/cloudflare-proxied: "true"
custom.io/ttl: "60"
```

## See Also

- [Configuration Precedence](configuration-precedence.md) - Understanding how external-dns processes configuration
- [FAQ](../faq.md) - Frequently asked questions
- [AWS Provider](../tutorials/aws.md) - AWS Route53 configuration
- [Cloudflare Provider](../tutorials/cloudflare.md) - Cloudflare configuration
