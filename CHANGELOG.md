## v0.4.0 - 2017-07-21

  - ExternalDNS now supports three more DNS providers:
    * [AzureDNS](https://azure.microsoft.com/en-us/services/dns) @peterhuene
    * [CloudFlare](https://www.cloudflare.com/de/dns) @njuettner
    * [DigitalOcean](https://www.digitalocean.com/products/networking) @njuettner
  - Fixed a bug that prevented ExternalDNS to be run on Tectonic clusters @sstarcher
  - ExternalDNS is now a full replace for Molecule Software's `route53-kubernetes` @iterion
  - The `external-dns.alpha.kubernetes.io/hostname` annotation accepts now a comma separated list of hostnames and a trailing period is not required anymore. @totallyunknown
  - The flag `--domain-filter` can be repeated multiple times like `--domain-filter=example.com --domain-filter=company.org.`. @totallyunknown
  - A trailing period is not required anymore for `--domain-filter` when AWS (or any other) provider is used. @totallyunknown
  - We added a FakeSource that generates random endpoints and allows to run ExternalDNS without a Kubernetes cluster (e.g. for testing providers) @ismith
  - All HTTP requests to external APIs (e.g. DNS providers) generate client side metrics. @linki
  - The `--zone` parameter was removed in favor of a provider independent `--domain-filter` flag. @linki
  - All flags can now also be set via environment variables. @linki

## v0.3.0 - 2017-05-08

Features:

  - Changed the flags to the v0.3 semantics, the following has changed:
    1. The TXT registry is used by default and has an owner ID of `default`
    2. `--dry-run` is disabled by default
    3. The `--compatibility` flag was added and takes a string instead of a boolean
    4. The `--in-cluster` flag has been dropped for auto-detection
    5. The `--zone` specifier has been replaced by a `--domain-filter` that filters domains by suffix
  - Improved logging output
  - Generate DNS Name from template for services/ingress if annotation is missing but `--fqdn-template` is specified
  - Route 53, Google CloudDNS: Support creation of records in multiple hosted zones.
  - Route 53: Support creation of ALIAS records when endpoint target is a ELB/ALB.
  - Ownership via TXT records
    1. Create TXT records to mark the records managed by External DNS
    2. Supported for AWS Route53 and Google CloudDNS
    3. Configurable TXT record DNS name format
  - Add support for altering the DNS record modification behavior via policies.

## v0.2.0 - 2017-04-07

Features:

  - Support creation of CNAME records when endpoint target is a hostname.
  - Allow omitting the trailing dot in Service annotations.
  - Expose basic Go metrics via Prometheus.

Documentation:

  - Add documentation on how to setup ExternalDNS for Services on AWS.

## v0.1.1 - 2017-04-03

Bug fixes:

  - AWS Route 53: Do not submit request when there are no changes.

## v0.1.0 - 2017-03-30 (KubeCon)

Features:

  - Manage DNS records for Services with `Type=LoadBalancer` on Google CloudDNS.
