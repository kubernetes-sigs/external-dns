## v0.5.9 - 2018-11-22

  - Core: Update delivery.yaml to new format (#782) @linki
  - Core: Adjust gometalinter timeout by setting env var (#778) @njuettner
  - Provider **Google**: Panic assignment to entry in nil map (#776) @njuettner
  - Docs: Fix typos (#769) @mooncak
  - Docs: Remove duplicated words (#768) @mooncak
  - Provider **Alibaba**: Alibaba Cloud Provider Fix Multiple Subdomains Bug (#767) @xianlubird
  - Core: Add Traefik to the supported list of ingress controllers (#764) @coderanger
  - Provider **Dyn**: Fix some typos in returned messages in dyn.go (#760) @AdamDang
  - Docs: Update Azure documentation (#756) @pascalgn
  - Provider **Oracle**: Oracle doc fix (add "key:" to secret) (#750) @CaptTofu
  - Core: Docker MAINTAINER is deprecated - using LABEL instead (#747) @helgi
  - Core: Feature add alias annotation (#742) @vaegt
  - Provider **RFC2136**: Fix rfc2136 - setup fails issue and small docs (#741) @antlad
  - Core: Fix nil map access of endpoint labels (#739) @shashidharatd
  - Provider **PowerDNS**: PowerDNS Add DomainFilter support (#737) @ottoyiu
  - Core: Fix domain-filter matching logic to not match similar domain names (#736) @ottoyiu
  - Core: Matching entire string for wildcard in txt records with prefixes (#727) @etopeter
  - Provider **Designate**: Fix TLS issue with OpenStack auth (#717) @FestivalBobcats
  - Provider **AWS**: Add helper script to update route53 txt owner entries (#697) @efranford
  - Provider **CoreDNS**: Migrate to use etcd client v3 for CoreDNS provider (#686) @shashidharatd
  - Core: Create a non-root user to run the container process (#684) @coderanger
  - Core: Do not replace TXT records with A/CNAME records in planner (#581) @jchv

## v0.5.8 - 2018-10-11

  - New Provider: RFC2136 (#702) @antlad
  - Add Linode to list of supported providers (#730) @cliedeman
  - Correctly populate target health check on existing records (#724) @linki
  - Don't erase Endpoint labels (#713) @sebastien-prudhomme

## v0.5.7 - 2018-09-27

  - Pass all relevant CLI flags to AWS provider (#719) @linki
  - Replace glog with a noop logger (#714) @linki
  - Fix handling of custom TTL values with Google DNS. (#704) @kevinmdavis
  - Continue even if node listing fails (#701) @pascalgn
  - Fix Host field in HTTP request when using pdns provider (#700) @peterbale
  - Allow AWS batching to fully sync on each run (#699) @bartelsielski

## v0.5.6 - 2018-09-07
  
  - Alibaba Cloud (#696) @xianlubird  
  - Add Source implementation for Istio Gateway (#694) @jonasrmichel
  - CRD source based on getting endpoints from CRD (#657) @shashidharatd
  - Add filter by service type feature (#653) @Devatoria
  - Add generic metrics for Source & Registry Errors (#652) @wleese

## v0.5.5 - 2018-08-17

  - Configure req timeout calling k8s APIs (#681) @jvassev
  - Adding assume role to aws_sd provider (#676) @lb-saildrone
  - Dyn: cache records per zone using zone's serial number (#675) @jvassev
  - Linode provider (#674) @cliedeman
  - Cloudflare Link Language Specificity (#673) @christopherhein
  - Retry calls to dyn on ErrRateLimited (#671) @jvassev
  - Add support to configure TTLs on DigitalOcean (#667) @andrewsomething
  - Log level warning option (#664) @george-angel
  - Fix usage of k8s.io/client-go package (#655) @shashidharatd
  - Fix for empty target annotation (#647) @rdrgmnzs
  - Fix log message for #592 when no updates in hosted zones (#634) @audip
  - Add aws-evaluate-target-health flag (#628) @peterbale
  - Exoscale provider (#625) @FaKod @greut
  - Oracle Cloud Infrastructure DNS provider (#626) @prydie
  - Update DO CNAME type API request to prevent error 422 (#624) @nenadilic84
  - Fix typo in cloudflare.md (#623) @derekperkins
  - Infoblox-go-client was only setting timeout for http.Transport.ResponseHeaderTimeout instead of for http.Client (#615) @khrisrichardson
  - Adding a flag to optionally publish hostIP instead of podIP for headless services (#597) @Arttii

## v0.5.4 - 2018-06-28

  - Only store endpoints with their labels in the cache (#612) @njuettner
  - Read hostnames from spec.tls.hosts on Ingress object (#611) @ysoldak
  - Reorder provider/aws suitable-zones tests (#608) @elordahl
  - Adds TLS flags for pdns provider (#607) @jhoch-palantir
  - Update RBAC for external-dns to list nodes (#600) @njuettner
  - Add aws max change count flag (#596) @peterbale
  - AWS provider: Properly check suitable domains (#594) @elordahl
  - Annotation with upper-case hostnames block further updates (#579) @njuettner
  
## v0.5.3 - 2018-06-15

  - Print a message if no hosted zones match (aws provider) (#592) @svend
  - Add support for NodePort services (#559) @grimmy
  - Update azure.md to fix protocol value (#593) @JasonvanBrackel
  - Add cache to limit calls to providers (#589) @jessfraz
  - Add Azure MSI support (#578) @r7vme
  - CoreDNS/SkyDNS provider (#253) @istalker2

## v0.5.2 - 2018-05-31

  - DNSimple: Make DNSimple tolerant of unknown zones (#574) @jbowes
  - Cloudflare: Custom record TTL (#572) @njuettner
  - AWS ServiceDiscovery: Implementation of AWS ServiceDiscovery provider (#483) @vanekjar
  - Update docs to latest changes (#563) @Raffo
  - New source - connector (#552) @shashidharatd
  - Update AWS SDK dependency to v1.13.7 @vanekjar

## v0.5.1 - 2018-05-16

  - Refactor implementation of sync loop to use `time.Ticker` (#553) @r0fls
  - Document how ExternalDNS gets permission to change AWS Route53 entries (#557) @hjacobs
  - Fix CNAME support for the PowerDNS provider (#547) @kciredor
  - Add support for hostname annotation in Ingress resource (#545) @rajatjindal
  - Fix for TTLs being ignored on headless Services (#546) @danbondd
  - Fix failing tests by giving linters more time to do their work (#548) @linki
  - Fix misspelled flag for the OpenStack Designate provider (#542) @zentale
  - Document additional RBAC rules needed to read Pods (#538) @danbondd

## v0.5.0 - 2018-04-23

  - Google: Correctly filter records that don't match all filters (#533) @prydie @linki
  - AWS: add support for AWS Network Load Balancers (#531) @linki
  - Add a flag that allows FQDN template and annotations to combine (#513) @helgi
  - Fix: Use PodIP instead of HostIP for headless Services (#498) @nrobert13
  - Support a comma separated list for the FQDN template (#512) @helgi
  - Google Provider: Add auto-detection of Google Project when running on GCP (#492) @drzero42
  - Add custom TTL support for DNSimple (#477) @jbowes
  - Fix docker build and delete vendor files which were not deleted (#473) @njuettner
  - DigitalOcean: DigitalOcean creates entries with host in them twice (#459) @njuettner
  - Bugfix: Retrive all DNSimple response pages (#468) @jbowes
  - external-dns does now provide support for multiple targets for A records. This is currently only supported by the Google Cloud DNS provider (#418) @dereulenspiegel
  - Graceful handling of misconfigure password for dyn provider (#470) @jvassev
  - Don't log sensitive data on start (#463) @jvassev
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
