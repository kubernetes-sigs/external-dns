  - Google: Improve logging to help trace misconfigurations (#388) @stealthybox
  - AWS: In addition to the one best public hosted zone, records will be added to all matching private hosted zones (#356) @coreypobrien
  - Every record managed by External DNS is now mapped to a kubernetes resource (service/ingress) @ideahitme
    - New field is stored in TXT DNS record which reflects which kubernetes resource has acquired the DNS name
    - Target of DNS record is changed only if corresponding kubernetes resource target changes
    - If kubernetes resource is deleted, then another resource may acquire DNS name
    - "Flapping" target issue is resolved by providing a consistent and defined mechanism for choosing a target
  - New `--zone-id-filter` parameter allows filtering by zone id (#422) @vboginskey
  - TTL annotation check for azure records (#436) @stromming
  - Switch from glide to dep (#435) @bkochendorfer

## v0.4.8 - 2017-11-22

  - Allow filtering by source annotation via `--annotation-filter` (#354) @khrisrichardson
  - Add support for Headless hostPort services (#324) @Arttii
  - AWS: Added change batch limiting to a maximum of 4000 Route53 updates in one API call.  Changes exceeding the limit will be dropped but all related changes by hostname are preserved within the limit. (#368) @bitvector2
  - Google: Support configuring TTL by annotation: `external-dns.alpha.kubernetes.io/ttl`. (#389) @stealthybox
  - Infoblox: add option `--no-infoblox-ssl-verify` (#378) @khrisrichardson
  - Inmemory: add support to specify zones for inmemory provider via command line (#366) @ffledgling

## v0.4.7 - 2017-10-18

  - CloudFlare: Disable proxy mode for TXT and others (#361) @dunglas

## v0.4.6 - 2017-10-12

  - [AWS Route53 provider] Support customization of DNS record TTL through the use of annotation `external-dns.alpha.kubernetes.io/ttl` on services or ingresses (#320) @kevinjqiu
  - Added support for [DNSimple](https://dnsimple.com/) as DNS provider (#224) @jose5918
  - Added support for [Infoblox](https://www.infoblox.com/products/dns/) as DNS provider (#349) @khrisrichardson

## v0.4.5 - 2017-09-24

  - Add `--log-level` flag to control log verbosity and remove `--debug` flag in favour of `--log-level=debug` (#339) @ultimateboy
  - AWS: Allow filtering for private and public zones via `--aws-zone-type` flag (#329) @linki
  - CloudFlare: Add `--cloudflare-proxied` flag to toggle CloudFlare proxy feature (#340) @dunglas
  - Kops Compatibility: Isolate ALIAS type in AWS provider (#248) @sethpollack

## v0.4.4 - 2017-08-17

  - ExternalDNS now services of type `ClusterIP` with the use of the `--publish-internal-services`.  Enabling this will now create the apprioriate A records for the given service's internal ip.  @jrnt30
  - Fix to have external target annotations on ingress resources replace existing endpoints instead of appending to them (#318)

## v0.4.3 - 2017-08-10

  - Support new `external-dns.alpha.kubernetes.io/target` annotation for Ingress (#312)
  - Fix for wildcard domains in Route53 (#302)

## v0.4.2 - 2017-08-03

  - Fix to support multiple hostnames for Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes) compatibility (#301)

## v0.4.1 - 2017-07-28

  - Fix incorrect order of constructor parameters (#298)

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
