CLI Usage
=========

```
usage: external-dns --source=source --provider=provider [<flags>]

ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.

Note that all flags may be replaced with env vars - `--flag` -> `EXTERNAL_DNS_FLAG=1` or `--flag value` -> `EXTERNAL_DNS_FLAG=value`

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --version                      Show application version.
  --server=""                    The Kubernetes API server to connect to (default: auto-detect)
  --kubeconfig=""                Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)
  --request-timeout=30s          Request timeout when calling Kubernetes APIs. 0s means no timeout
  --cf-api-endpoint=""           The fully-qualified domain name of the cloud foundry instance you are targeting
  --cf-username=""               The username to log into the cloud foundry API
  --cf-password=""               The password to log into the cloud foundry API
  --contour-load-balancer="heptio-contour/contour"
                                 The fully-qualified name of the Contour load balancer service. (default: heptio-contour/contour)
  --gloo-namespace="gloo-system"
                                 Gloo namespace. (default: gloo-system)
  --skipper-routegroup-groupversion="zalando.org/v1"
                                 The resource version for skipper routegroup
  --source=source ...            The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, node, fake, connector, istio-gateway, istio-virtualservice, cloudfoundry, contour-ingressroute,
                                 contour-httpproxy, gloo-proxy, crd, empty, skipper-routegroup, openshift-route, ambassador-host, kong-tcpingress)
  --openshift-router-name=OPENSHIFT-ROUTER-NAME
                                 if source is openshift-route then you can pass the ingress controller name. Based on this name external-dns will select the respective router from the route status and map that routerCanonicalHostname to the route host while creating a
                                 CNAME record.
  --namespace=""                 Limit sources of endpoints to a specific namespace (default: all namespaces)
  --annotation-filter=""         Filter sources managed by external-dns via annotation using label selector semantics (default: all sources)
  --label-filter=""              Filter sources managed by external-dns via label selector when listing all resources; currently supported by source types CRD, ingress, service and openshift-route
  --fqdn-template=""             A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional). Accepts comma separated list for multiple global FQDN.
  --combine-fqdn-annotation      Combine FQDN template and Annotations instead of overwriting
  --ignore-hostname-annotation   Ignore hostname annotation when generating DNS names, valid only when using fqdn-template is set (optional, default: false)
  --ignore-ingress-tls-spec      Ignore tls spec section in ingresses resources, applicable only for ingress sources (optional, default: false)
  --compatibility=               Process annotation semantics from legacy implementations (optional, options: mate, molecule, kops-dns-controller)
  --ignore-ingress-rules-spec    Ignore rules spec section in ingresses resources, applicable only for ingress sources (optional, default: false)
  --publish-internal-services    Allow external-dns to publish DNS records for ClusterIP services (optional)
  --publish-host-ip              Allow external-dns to publish host-ip for headless services (optional)
  --always-publish-not-ready-addresses
                                 Always publish also not ready addresses for headless services (optional)
  --connector-source-server="localhost:8080"
                                 The server to connect for connector source, valid only when using connector source
  --crd-source-apiversion="externaldns.k8s.io/v1alpha1"
                                 API version of the CRD for crd source, e.g. `externaldns.k8s.io/v1alpha1`, valid only when using crd source
  --crd-source-kind="DNSEndpoint"
                                 Kind of the CRD for the crd source in API group and version specified by crd-source-apiversion
  --service-type-filter=SERVICE-TYPE-FILTER ...
                                 The service types to take care about (default: all, expected: ClusterIP, NodePort, LoadBalancer or ExternalName)
  --managed-record-types=A... ...
                                 Comma separated list of record types to manage (default: A, CNAME) (supported records: CNAME, A, NS
  --default-targets=DEFAULT-TARGETS ...
                                 Set globally default IP address that will apply as a target instead of source addresses. Specify multiple times for multiple targets (optional)
  --provider=provider            The DNS provider where the DNS records will be created (required, options: aws, aws-sd, godaddy, google, azure, azure-dns, azure-private-dns, bluecat, cloudflare, rcodezero, digitalocean, dnsimple, akamai, infoblox, dyn, designate,
                                 coredns, skydns, inmemory, ovh, pdns, oci, exoscale, linode, rfc2136, ns1, transip, vinyldns, rdns, scaleway, vultr, ultradns, gandi, safedns)
  --domain-filter= ...           Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)
  --exclude-domains= ...         Exclude subdomains (optional)
  --regex-domain-filter=         Limit possible domains and target zones by a Regex filter; Overrides domain-filter (optional)
  --regex-domain-exclusion=      Regex filter that excludes domains and target zones matched by regex-domain-filter (optional)
  --zone-name-filter= ...        Filter target zones by zone domain (For now, only AzureDNS provider is using this flag); specify multiple times for multiple zones (optional)
  --zone-id-filter= ...          Filter target zones by hosted zone id; specify multiple times for multiple zones (optional)
  --google-project=""            When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.
  --google-batch-change-size=1000
                                 When using the Google provider, set the maximum number of changes that will be applied in each batch.
  --google-batch-change-interval=1s
                                 When using the Google provider, set the interval between batch changes.
  --google-zone-visibility=      When using the Google provider, filter for zones with this visibility (optional, options: public, private)
  --alibaba-cloud-config-file="/etc/kubernetes/alibaba-cloud.json"
                                 When using the Alibaba Cloud provider, specify the Alibaba Cloud configuration file (required when --provider=alibabacloud
  --alibaba-cloud-zone-type=     When using the Alibaba Cloud provider, filter for zones of this type (optional, options: public, private)
  --aws-zone-type=               When using the AWS provider, filter for zones of this type (optional, options: public, private)
  --aws-zone-tags= ...           When using the AWS provider, filter for zones with these tags
  --aws-assume-role=""           When using the AWS provider, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns` (optional)
  --aws-batch-change-size=1000   When using the AWS provider, set the maximum number of changes that will be applied in each batch.
  --aws-batch-change-interval=1s
                                 When using the AWS provider, set the interval between batch changes.
  --aws-evaluate-target-health   When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)
  --aws-api-retries=3            When using the AWS provider, set the maximum number of retries for API calls before giving up.
  --aws-prefer-cname             When using the AWS provider, prefer using CNAME instead of ALIAS (default: disabled)
  --aws-zones-cache-duration=0s  When using the AWS provider, set the zones list cache TTL (0s to disable).
  --aws-sd-service-cleanup       When using the AWS CloudMap provider, delete empty Services without endpoints (default: disabled)
  --azure-config-file="/etc/kubernetes/azure.json"
                                 When using the Azure provider, specify the Azure configuration file (required when --provider=azure
  --azure-resource-group=""      When using the Azure provider, override the Azure resource group to use (required when --provider=azure-private-dns)
  --azure-subscription-id=""     When using the Azure provider, specify the Azure configuration file (required when --provider=azure-private-dns)
  --azure-user-assigned-identity-client-id=""
                                 When using the Azure provider, override the client id of user assigned identity in config file (optional)
  --bluecat-dns-configuration=""
                                 When using the Bluecat provider, specify the Bluecat DNS configuration string (optional when --provider=bluecat)
  --bluecat-config-file="/etc/kubernetes/bluecat.json"
                                 When using the Bluecat provider, specify the Bluecat configuration file (optional when --provider=bluecat)
  --bluecat-dns-view=""          When using the Bluecat provider, specify the Bluecat DNS view string (optional when --provider=bluecat)
  --bluecat-gateway-host=""      When using the Bluecat provider, specify the Bluecat Gateway Host (optional when --provider=bluecat)
  --bluecat-root-zone=""         When using the Bluecat provider, specify the Bluecat root zone (optional when --provider=bluecat)
  --bluecat-skip-tls-verify      When using the Bluecat provider, specify to skip TLS verification (optional when --provider=bluecat) (default: false)
  --bluecat-dns-server-name=""   When using the Bluecat provider, specify the Bluecat DNS Server to initiate deploys against. This is only used if --bluecat-dns-deploy-type is not 'no-deploy' (optional when --provider=bluecat)
  --bluecat-dns-deploy-type="no-deploy"
                                 When using the Bluecat provider, specify the type of DNS deployment to initiate after records are updated. Valid options are 'full-deploy' and 'no-deploy'. Deploy will only execute if --bluecat-dns-server-name is set (optional when
                                 --provider=bluecat)
  --cloudflare-proxied           When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)
  --cloudflare-zones-per-page=50
                                 When using the Cloudflare provider, specify how many zones per page listed, max. possible 50 (default: 50)
  --coredns-prefix="/skydns/"    When using the CoreDNS provider, specify the prefix name
  --akamai-serviceconsumerdomain=""
                                 When using the Akamai provider, specify the base URL (required when --provider=akamai and edgerc-path not specified)
  --akamai-client-token=""       When using the Akamai provider, specify the client token (required when --provider=akamai and edgerc-path not specified)
  --akamai-client-secret=""      When using the Akamai provider, specify the client secret (required when --provider=akamai and edgerc-path not specified)
  --akamai-access-token=""       When using the Akamai provider, specify the access token (required when --provider=akamai and edgerc-path not specified)
  --akamai-edgerc-path=""        When using the Akamai provider, specify the .edgerc file path. Path must be reachable form invocation environment. (required when --provider=akamai and *-token, secret serviceconsumerdomain not specified)
  --akamai-edgerc-section=""     When using the Akamai provider, specify the .edgerc file path (Optional when edgerc-path is specified)
  --infoblox-grid-host=""        When using the Infoblox provider, specify the Grid Manager host (required when --provider=infoblox)
  --infoblox-wapi-port=443       When using the Infoblox provider, specify the WAPI port (default: 443)
  --infoblox-wapi-username="admin"
                                 When using the Infoblox provider, specify the WAPI username (default: admin)
  --infoblox-wapi-password=""    When using the Infoblox provider, specify the WAPI password (required when --provider=infoblox)
  --infoblox-wapi-version="2.3.1"
                                 When using the Infoblox provider, specify the WAPI version (default: 2.3.1)
  --infoblox-ssl-verify          When using the Infoblox provider, specify whether to verify the SSL certificate (default: true, disable with --no-infoblox-ssl-verify)
  --infoblox-view=""             DNS view (default: "")
  --infoblox-max-results=0       Add _max_results as query parameter to the URL on all API requests. The default is 0 which means _max_results is not set and the default of the server is used.
  --infoblox-fqdn-regex=""       Apply this regular expression as a filter for obtaining zone_auth objects. This is disabled by default.
  --infoblox-create-ptr          When using the Infoblox provider, create a ptr entry in addition to an entry
  --infoblox-cache-duration=0    When using the Infoblox provider, set the record TTL (0s to disable).
  --dyn-customer-name=""         When using the Dyn provider, specify the Customer Name
  --dyn-username=""              When using the Dyn provider, specify the Username
  --dyn-password=""              When using the Dyn provider, specify the password
  --dyn-min-ttl=DYN-MIN-TTL      Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.
  --oci-config-file="/etc/kubernetes/oci.yaml"
                                 When using the OCI provider, specify the OCI configuration file (required when --provider=oci
  --rcodezero-txt-encrypt        When using the Rcodezero provider with txt registry option, set if TXT rrs are encrypted (default: false)
  --inmemory-zone= ...           Provide a list of pre-configured zones for the inmemory provider; specify multiple times for multiple zones (optional)
  --ovh-endpoint="ovh-eu"        When using the OVH provider, specify the endpoint (default: ovh-eu)
  --ovh-api-rate-limit=20        When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)
  --pdns-server="http://localhost:8081"
                                 When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)
  --pdns-api-key=""              When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)
  --pdns-tls-enabled             When using the PowerDNS/PDNS provider, specify whether to use TLS (default: false, requires --tls-ca, optionally specify --tls-client-cert and --tls-client-cert-key)
  --ns1-endpoint=""              When using the NS1 provider, specify the URL of the API endpoint to target (default: https://api.nsone.net/v1/)
  --ns1-ignoressl                When using the NS1 provider, specify whether to verify the SSL certificate (default: false)
  --ns1-min-ttl=NS1-MIN-TTL      Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.
  --digitalocean-api-page-size=50
                                 Configure the page size used when querying the DigitalOcean API.
  --godaddy-api-key=""           When using the GoDaddy provider, specify the API Key (required when --provider=godaddy)
  --godaddy-api-secret=""        When using the GoDaddy provider, specify the API secret (required when --provider=godaddy)
  --godaddy-api-ttl=GODADDY-API-TTL
                                 TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is not provided.
  --godaddy-api-ote              When using the GoDaddy provider, use OTE api (optional, default: false, when --provider=godaddy)
  --tls-ca=""                    When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)
  --tls-client-cert=""           When using TLS communication, the path to the certificate to present as a client (not required for TLS)
  --tls-client-cert-key=""       When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)
  --exoscale-endpoint="https://api.exoscale.ch/dns"
                                 Provide the endpoint for the Exoscale provider
  --exoscale-apikey=""           Provide your API Key for the Exoscale provider
  --exoscale-apisecret=""        Provide your API Secret for the Exoscale provider
  --rfc2136-host=""              When using the RFC2136 provider, specify the host of the DNS server
  --rfc2136-port=0               When using the RFC2136 provider, specify the port of the DNS server
  --rfc2136-zone=""              When using the RFC2136 provider, specify the zone entry of the DNS server to use
  --rfc2136-insecure             When using the RFC2136 provider, specify whether to attach TSIG or not (default: false, requires --rfc2136-tsig-keyname and rfc2136-tsig-secret)
  --rfc2136-tsig-keyname=""      When using the RFC2136 provider, specify the TSIG key to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-secret=""       When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-secret-alg=""   When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-axfr            When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-min-ttl=0s           When using the RFC2136 provider, specify minimal TTL (in duration format) for records. This value will be used if the provided TTL for a service/ingress is lower than this
  --rfc2136-gss-tsig             When using the RFC2136 provider, specify whether to use secure updates with GSS-TSIG using Kerberos (default: false, requires --rfc2136-kerberos-realm, --rfc2136-kerberos-username, and rfc2136-kerberos-password)
  --rfc2136-kerberos-username=""
                                 When using the RFC2136 provider with GSS-TSIG, specify the username of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)
  --rfc2136-kerberos-password=""
                                 When using the RFC2136 provider with GSS-TSIG, specify the password of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)
  --rfc2136-kerberos-realm=""    When using the RFC2136 provider with GSS-TSIG, specify the realm of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)
  --rfc2136-batch-change-size=50
                                 When using the RFC2136 provider, set the maximum number of changes that will be applied in each batch.
  --transip-account=""           When using the TransIP provider, specify the account name (required when --provider=transip)
  --transip-keyfile=""           When using the TransIP provider, specify the path to the private key file (required when --provider=transip)
  --policy=sync                  Modify how DNS records are synchronized between sources and providers (default: sync, options: sync, upsert-only, create-only)
  --registry=txt                 The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, aws-sd)
  --txt-owner-id="default"       When using the TXT registry, a name that identifies this instance of ExternalDNS (default: default)
  --txt-prefix=""                When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional). Could contain record type template like '%{record_type}-prefix-'. Mutual exclusive with txt-suffix!
  --txt-suffix=""                When using the TXT registry, a custom string that's suffixed to the host portion of each ownership DNS record (optional). Could contain record type template like '-%{record_type}-suffix'. Mutual exclusive with txt-prefix!
  --txt-wildcard-replacement=""  When using the TXT registry, a custom string that's used instead of an asterisk for TXT records corresponding to wildcard DNS records (optional)
  --txt-cache-interval=0s        The interval between cache synchronizations in duration format (default: disabled)
  --interval=1m0s                The interval between two consecutive synchronizations in duration format (default: 1m)
  --min-event-sync-interval=5s   The minimum interval between two consecutive synchronizations triggered from kubernetes events in duration format (default: 5s)
  --once                         When enabled, exits the synchronization loop after the first iteration (default: disabled)
  --dry-run                      When enabled, prints DNS record changes rather than actually performing them (default: disabled)
  --events                       When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)
  --log-format=text              The format in which log messages are printed (default: text, options: text, json)
  --metrics-address=":7979"      Specify where to serve the metrics and health check endpoint (default: :7979)
  --log-level=info               Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal
  ```