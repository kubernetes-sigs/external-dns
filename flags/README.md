| Flag | Default | Type | Description |
|:----|:----|:----|:----|
| **Flags related to Kubernetes** | | | |
|--server|string| |The Kubernetes API server to connect to (default: auto-detect)|
|--kubeconfig|string| |Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)|
|--request-timeout|time.Duration|time.Second * 30|Request timeout when calling Kubernetes APIs. 0s means no timeout|
| **Flags related to cloud foundry** ** | | | |
|--cf-api-endpoint|string| |The fully-qualified domain name of the cloud foundry instance you are targeting|
|--cf-username|string| |The username to log into the cloud foundry API|
|--cf-password|string| |The password to log into the cloud foundry API|
| **Flags related to Contour** | | | |
|--contour-load-balancer|string|heptio-contour/contour|The fully-qualified name of the Contour load balancer service. (default: heptio-contour/contour)|
| **Flags related to Gloo** | | | |
|--gloo-namespace|string|gloo-system|Gloo namespace. (default: gloo-system)|
| **Flags related to Skipper RouteGroup** | | | |
|--skipper-routegroup-groupversion|string|zalando.org/v1|The resource version for skipper routegroup|
| **Flags related to processing source** | | | |
|--source|string| |The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, node, fake, connector, gateway-httproute, gateway-tlsroute, gateway-tcproute, gateway-udproute, istio-gateway, istio-virtualservice, cloudfoundry, contour-ingressroute, contour-httpproxy, gloo-proxy, crd, empty, skipper-routegroup, openshift-route, ambassador-host, kong-tcpingress)|
|--openshift-router-name|string| |if source is openshift-route then you can pass the ingress controller name. Based on this name external-dns will select the respective router from the route status and map that routerCanonicalHostname to the route host while creating a CNAME record.|
|--namespace|string| |Limit sources of endpoints to a specific namespace (default: all namespaces)|
|--annotation-filter|string| |Filter sources managed by external-dns via annotation using label selector semantics (default: all sources)|
|--label-filter|string|labels.Everything().String()|Filter sources managed by external-dns via label selector when listing all resources; currently supported by source types CRD, ingress, service and openshift-route|
|--fqdn-template|string| |A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional). Accepts comma separated list for multiple global FQDN.|
|--combine-fqdn-annotation|bool|FALSE|Combine FQDN template and Annotations instead of overwriting|
|--ignore-hostname-annotation|bool|FALSE|Ignore hostname annotation when generating DNS names, valid only when using fqdn-template is set (optional, default: false)|
|--ignore-ingress-tls-spec|bool|FALSE|Ignore tls spec section in ingresses resources, applicable only for ingress sources (optional, default: false)|
|--gateway-namespace|string| |Limit Gateways of Route endpoints to a specific namespace (default: all namespaces)|
|--gateway-label-filter|string| |Filter Gateways of Route endpoints via label selector (default: all gateways)|
|--compatibility|string| |Process annotation semantics from legacy implementations (optional, options: mate, molecule, kops-dns-controller)|
|--ignore-ingress-rules-spec|bool| |Ignore rules spec section in ingresses resources, applicable only for ingress sources (optional, default: false)|
|--publish-internal-services|bool|FALSE|Allow external-dns to publish DNS records for ClusterIP services (optional)|
|--publish-host-ip|bool|FALSE|Allow external-dns to publish host-ip for headless services (optional)|
|--always-publish-not-ready-addresses|bool| |Always publish also not ready addresses for headless services (optional)|
|--connector-source-server|string|localhost:8080|The server to connect for connector source, valid only when using connector source|
|--crd-source-apiversion|string|externaldns.k8s.io/v1alpha1|API version of the CRD for crd source, e.g. `externaldns.k8s.io/v1alpha1`, valid only when using crd source|
|--crd-source-kind|string|DNSEndpoint|Kind of the CRD for the crd source in API group and version specified by crd-source-apiversion|
|--service-type-filter|[]string|[]string{}|The service types to take care about (default: all, expected: ClusterIP, NodePort, LoadBalancer or ExternalName)|
|--managed-record-types|[]string|[]string{endpoint.RecordTypeA endpoint.RecordTypeCNAME}|Comma separated list of record types to manage (default: A, CNAME) (supported records: CNAME, A, NS)|
|--default-targets|[]string|[]string{}|Set globally default IP address that will apply as a target instead of source addresses. Specify multiple times for multiple targets (optional)|
|--target-net-filter|[]string|[]string{}|Limit possible targets by a net filter; Specify multiple times for multiple possible nets (optional)|
|--exclude-target-net|[]string|[]string{}|Exclude target nets (optional)|
| **Flags related to providers** | | | |
|--provider|string| |The DNS provider where the DNS records will be created (required, options: aws, aws-sd, godaddy, google, azure, azure-dns, azure-private-dns, bluecat, cloudflare, rcodezero, digitalocean, dnsimple, akamai, infoblox, dyn, designate, coredns, skydns, ibmcloud, inmemory, ovh, pdns, oci, exoscale, linode, rfc2136, ns1, transip, vinyldns, rdns, scaleway, vultr, ultradns, gandi, safedns, tencentcloud)|
|--domain-filter|[]string|[]string{}|Limit possible target zones by a domain suffix; Specify multiple times for multiple domains (optional)|
|--exclude-domains|[]string|[]string{}|Exclude subdomains (optional)|
|--regex-domain-filter|*regexp.Regexp|regexp.MustCompile("")|Limit possible domains and target zones by a Regex filter; Overrides domain-filter (optional)|
|--regex-domain-exclusion|*regexp.Regexp|regexp.MustCompile("")|Regex filter that excludes domains and target zones matched by regex-domain-filter (optional)|
|--zone-name-filter|[]string| |Filter target zones by zone domain (For now, only AzureDNS provider is using this flag); Specify multiple times for multiple zones (optional)|
|--zone-id-filter|[]string|[]string{}|Filter target zones by hosted zone id; Specify multiple times for multiple zones (optional)|
| **Flags related to Google provider** | | | |
|--google-project|string| |When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.|
|--google-batch-change-size|int|1000|When using the Google provider, set the maximum number of changes that will be applied in each batch.|
|--google-batch-change-interval|time.Duration|time.Second|When using the Google provider, set the interval between batch changes.|
|--google-zone-visibility|string| |When using the Google provider, filter for zones with this visibility (optional, options: public, private)|
| **Flags related to Alibaba provider** | | | |
|--alibaba-cloud-config-file|string|/etc/kubernetes/alibaba-cloud.json|When using the Alibaba Cloud provider, specify the Alibaba Cloud configuration file (required when --provider=alibabacloud|
|--alibaba-cloud-zone-type|string| |When using the Alibaba Cloud provider, filter for zones of this type (optional, options: public, private)|
| **Flags related to AWS provider** | | | |
|--aws-zone-type|string| |When using the AWS provider, filter for zones of this type (optional, options: public, private)|
|--aws-zone-tags|[]string|[]string{}|When using the AWS provider, filter for zones with these tags|
|--aws-assume-role|string| |When using the AWS provider, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns` (optional)|
|--aws-assume-role-external-id|string| |When using the AWS provider and assuming a role then specify this external ID` (optional)|
|--aws-batch-change-size|int|1000|When using the AWS provider, set the maximum number of changes that will be applied in each batch.|
|--aws-batch-change-interval|time.Duration|time.Second|When using the AWS provider, set the interval between batch changes.|
|--aws-evaluate-target-health|bool|TRUE|When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)|
|--aws-api-retries|int|3|When using the AWS provider, set the maximum number of retries for API calls before giving up.|
|--aws-prefer-cname|bool|FALSE|When using the AWS provider, prefer using CNAME instead of ALIAS (default: disabled)|
|--aws-zones-cache-duration|time.Duration|0 * time.Second|When using the AWS provider, set the zones list cache TTL (0s to disable|
|--aws-sd-service-cleanup|bool|FALSE|When using the AWS CloudMap provider, delete empty Services without endpoints (default: disabled)|
| **Flags related to Azure provider** | | | |
|--azure-config-file|string|/etc/kubernetes/azure.json|When using the Azure provider, specify the Azure configuration file (required when --provider=azure|
|--azure-resource-group|string| |When using the Azure provider, override the Azure resource group to use (required when --provider=azure-private-dns)|
|--azure-subscription-id|string| |When using the Azure provider, specify the Azure configuration file (required when --provider=azure-private-dns)|
|--azure-user-assigned-identity-client-id|string| |When using the Azure provider, override the client id of user assigned identity in config file (optional)|
| **Flags related to Tencent provider** | | | |
|--tencent-cloud-config-file|string|/etc/kubernetes/tencent-cloud.json|When using the Tencent Cloud provider, specify the Tencent Cloud configuration file (required when --provider=tencentcloud|
|--tencent-cloud-zone-type|string| |When using the Tencent Cloud provider, filter for zones with visibility (optional, options: public, private)|
| **Flags related to BlueCat provider** | | | |
|--bluecat-dns-configuration|string| |When using the Bluecat provider, specify the Bluecat DNS configuration string (optional when --provider=bluecat)|
|--bluecat-config-file|string|/etc/kubernetes/bluecat.json|When using the Bluecat provider, specify the Bluecat configuration file (optional when --provider=bluecat)|
|--bluecat-dns-view|string| |When using the Bluecat provider, specify the Bluecat DNS view string (optional when --provider=bluecat)|
|--bluecat-gateway-host|string| |When using the Bluecat provider, specify the Bluecat Gateway Host (optional when --provider=bluecat)|
|--bluecat-root-zone|string| |When using the Bluecat provider, specify the Bluecat root zone (optional when --provider=bluecat)|
|--bluecat-skip-tls-verify|bool| |When using the Bluecat provider, specify to skip TLS verification (optional when --provider=bluecat) (default: false)|
|--bluecat-dns-server-name|string| |When using the Bluecat provider, specify the Bluecat DNS Server to initiate deploys against. This is only used if --bluecat-dns-deploy-type is not 'no-deploy' (optional when --provider=bluecat)|
|--bluecat-dns-deploy-type|string|no-deploy|When using the Bluecat provider, specify the type of DNS deployment to initiate after records are updated. Valid options are 'full-deploy' and 'no-deploy'. Deploy will only execute if --bluecat-dns-server-name is set (optional when --provider=bluecat)|
| **Flags related to CloudFlare provider** | | | |
|--cloudflare-proxied|bool|FALSE|When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)|
|--cloudflare-zones-per-page|int|50|When using the Cloudflare provider, specify how many zones per page listed, max. possible 50 (default: 50)|
| **Flags related to CoreDNS provider** | | | |
|--coredns-prefix|string|/skydns/|When using the CoreDNS provider, specify the prefix name|
| **Flags related to Akamai provider** | | | |
|--akamai-serviceconsumerdomain|string| |When using the Akamai provider, specify the base URL (required when --provider=akamai and edgerc-path not specified)|
|--akamai-client-token|string| |When using the Akamai provider, specify the client token (required when --provider=akamai and edgerc-path not specified)|
|--akamai-client-secret|string| |When using the Akamai provider, specify the client secret (required when --provider=akamai and edgerc-path not specified)|
|--akamai-access-token|string| |When using the Akamai provider, specify the access token (required when --provider=akamai and edgerc-path not specified)|
|--akamai-edgerc-path|string| |When using the Akamai provider, specify the .edgerc file path. Path must be reachable form invocation environment. (required when --provider=akamai and *-token, secret serviceconsumerdomain not specified)|
|--akamai-edgerc-section|string| |When using the Akamai provider, specify the .edgerc file path (Optional when edgerc-path is specified)|
| **Flags related to Infoblox provider** | | | |
|--infoblox-grid-host|string| |When using the Infoblox provider, specify the Grid Manager host (required when --provider=infoblox)|
|--infoblox-wapi-port|int|443|When using the Infoblox provider, specify the WAPI port (default: 443)|
|--infoblox-wapi-username|string|admin|When using the Infoblox provider, specify the WAPI username (default: admin)|
|--infoblox-wapi-password|string| |When using the Infoblox provider, specify the WAPI password (required when --provider=infoblox)|
|--infoblox-wapi-version|string|2.3.1|When using the Infoblox provider, specify the WAPI version (default: 2.3.1)|
|--infoblox-ssl-verify|bool|TRUE|When using the Infoblox provider, specify whether to verify the SSL certificate (default: true, disable with --no-infoblox-ssl-verify)|
|--infoblox-view|string| |DNS view (default: \"\")|
|--infoblox-max-results|int|0|Add _max_results as query parameter to the URL on all API requests. The default is 0 which means _max_results is not set and the default of the server is used.|
|--infoblox-fqdn-regex|string| |Apply this regular expression as a filter for obtaining zone_auth objects. This is disabled by default.|
|--infoblox-create-ptr|bool|FALSE|When using the Infoblox provider, create a ptr entry in addition to an entry|
|--infoblox-cache-duration|int|0|When using the Infoblox provider, set the record TTL (0s to disable|
| **Flags related to Dyn provider** | | | |
|--dyn-customer-name|string| |When using the Dyn provider, specify the Customer Name|
|--dyn-username|string| |When using the Dyn provider, specify the Username|
|--dyn-password|string| |When using the Dyn provider, specify the password|
|--dyn-min-ttl|int| |Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.|
| **Flags related to OCI provider** | | | |
|--oci-config-file|string|/etc/kubernetes/oci.yaml|When using the OCI provider, specify the OCI configuration file (required when --provider=oci|
| **Flags related to Rcodezero provider** | | | |
|--rcodezero-txt-encrypt|bool|FALSE|When using the Rcodezero provider with txt registry option, set if TXT rrs are encrypted (default: false)|
| **Flags related to Inmemory provider** | | | |
|--inmemory-zone|[]string|[]string{}|Provide a list of pre-configured zones for the inmemory provider; Specify multiple times for multiple zones (optional)|
| **Flags related to OVH provider** | | | |
|--ovh-endpoint|string|ovh-eu|When using the OVH provider, specify the endpoint (default: ovh-eu)|
|--ovh-api-rate-limit|int|20|When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)|
| **Flags related to PowerDNS provider** | | | |
|--pdns-server|string|http://localhost:8081|When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)|
|--pdns-api-key|string `secure:"yes"`| |When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)|
|--pdns-tls-enabled|bool|FALSE|When using the PowerDNS/PDNS provider, specify whether to use TLS (default: false, requires --tls-ca, optionally specify --tls-client-cert and --tls-client-cert-key)|
| **Flags related to NS1 provider** | | | |
|--NS1-endpoint|string| |When using the NS1 provider, specify the URL of the API endpoint to target (default: https://api.nsone.net/v1/)|
|--NS1-ignoressl|bool|FALSE|When using the NS1 provider, specify whether to verify the SSL certificate (default: false)|
|--NS1-min-ttl|int| |Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.|
| **Flags related to DigitalOcean provider** | | | |
|--digitalocean-api-page-size|int|50|Configure the page size used when querying the DigitalOcean API.|
| **Flags related to IBM Cloud provider** | | | |
|--ibmcloud-config-file|string|/etc/kubernetes/ibmcloud.json|When using the IBM Cloud provider, specify the IBM Cloud configuration file (required when --provider=ibmcloud|
|--ibmcloud-proxied|bool|FALSE|When using the IBM provider, specify if the proxy mode must be enabled (default: disabled)|
| **Flags related to GoDaddy provider** | | | |
|--godaddy-api-key|string `secure:"yes"`| |When using the GoDaddy provider, specify the API Key (required when --provider=godaddy)|
|--godaddy-api-secret|string `secure:"yes"`| |When using the GoDaddy provider, specify the API secret (required when --provider=godaddy)|
|--godaddy-api-ttl|int64|600|TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is not provided.|
|--godaddy-api-ote|bool|FALSE|When using the GoDaddy provider, use OTE api (optional, default: false, when --provider=godaddy)|
| **Flags related to TLS communication** | | | |
|--tls-ca|string| |When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)|
|--tls-client-cert|string| |When using TLS communication, the path to the certificate to present as a client (not required for TLS)|
|--tls-client-cert-key|string| |When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)|
| **Flags related to Exoscale provider** | | | |
|--exoscale-endpoint|string|https://api.exoscale.ch/dns|Provide the endpoint for the Exoscale provider|
|--exoscale-apikey|string `secure:"yes"`| |Provide your API Key for the Exoscale provider|
|--exoscale-apisecret|string `secure:"yes"`| |Provide your API Secret for the Exoscale provider|
| **Flags related to RFC2136 provider** | | | |
|--RFC2136-host|string| |When using the RFC2136 provider, specify the host of the DNS server|
|--RFC2136-port|int|0|When using the RFC2136 provider, specify the port of the DNS server|
|--RFC2136-zone|string| |When using the RFC2136 provider, specify the zone entry of the DNS server to use|
|--RFC2136-insecure|bool|FALSE|When using the RFC2136 provider, specify whether to attach TSIG or not (default: false, requires --rfc2136-tsig-keyname and rfc2136-tsig-secret)|
|--RFC2136-tsig-keyname|string| |When using the RFC2136 provider, specify the TSIG key to attached to DNS messages (required when --rfc2136-insecure=false)|
|--RFC2136-tsig-secret|string `secure:"yes"`| |When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)|
|--RFC2136-tsig-secret-alg|string| |When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)|
|--RFC2136-tsig-axfr|bool|TRUE|When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)|
|--RFC2136-min-ttl|time.Duration|0|When using the RFC2136 provider, specify minimal TTL (in duration format) for records. This value will be used if the provided TTL for a service/ingress is lower than this|
|--RFC2136-gss-tsig|bool|FALSE|When using the RFC2136 provider, specify whether to use secure updates with GSS-TSIG using Kerberos (default: false, requires --rfc2136-kerberos-realm, --rfc2136-kerberos-username, and rfc2136-kerberos-password)|
|--RFC2136-kerberos-username|string| |When using the RFC2136 provider with GSS-TSIG, specify the username of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)|
|--RFC2136-kerberos-password|string| |When using the RFC2136 provider with GSS-TSIG, specify the password of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)|
|--RFC2136-kerberos-realm|string `secure:"yes"`| |When using the RFC2136 provider with GSS-TSIG, specify the realm of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)|
|--RFC2136-batch-change-size|int|50|When using the RFC2136 provider, set the maximum number of changes that will be applied in each batch.|
| **Flags related to TransIP provider** | | | |
|--transip-account|string| |When using the TransIP provider, specify the account name (required when --provider=transip)|
|--transip-keyfile|string| |When using the TransIP provider, specify the path to the private key file (required when --provider=transip)|
| **Flags related to Pihole provider** | | | |
|--pihole-server|string| |When using the Pihole provider, the base URL of the Pihole web server (required when --provider=pihole)|
|--pihole-password|string `secure:"yes"`| |When using the Pihole provider, the password to the server if it is protected|
|--pihole-tls-skip-verify|bool|FALSE|When using the Pihole provider, disable verification of any TLS certificates|
| **Flags related to the Plural provider** | | | |
|--plural-cluster|string| |When using the plural provider, specify the cluster name you're running with|
|--plural-provider|string| |When using the plural provider, specify the provider name you're running with|
| **Flags related to policies** | | | |
|--policy|string|sync|Modify how DNS records are synchronized between sources and providers (default: sync, options: sync, upsert-only, create-only)|
| **Flags related to the registry** | | | |
|--registry|string|txt|The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, aws-sd)|
|--txt-owner-id|string|default|When using the TXT registry, a name that identifies this instance of ExternalDNS (default: default)|
|--txt-prefix|string| |When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional). Could contain record type template like '%{record_type}-prefix-'. Mutual exclusive with txt-suffix!|
|--txt-suffix|string| |When using the TXT registry, a custom string that's suffixed to the host portion of each ownership DNS record (optional). Could contain record type template like '-%{record_type}-suffix'. Mutual exclusive with txt-prefix!|
|--txt-wildcard-replacement|string| |When using the TXT registry, a custom string that's used instead of an asterisk for TXT records corresponding to wildcard DNS records (optional)|
| **Flags related to the main control loop** | | | |
|--txt-cache-interval|time.Duration|0|The interval between cache synchronizations in duration format (default: disabled)|
|--interval|time.Duration|time.Minute|The interval between two consecutive synchronizations in duration format (default: 1m)|
|--min-event-sync-interval|time.Duration|5 * time.Second|The minimum interval between two consecutive synchronizations triggered from kubernetes events in duration format (default: 5s)|
|--once|bool|FALSE|When enabled, exits the synchronization loop after the first iteration (default: disabled)|
|--dry-run|bool|FALSE|When enabled, prints DNS record changes rather than actually performing them (default: disabled)|
|--events|bool|FALSE|When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)|
| **Miscellaneous flags** | | | |
|--log-format|string|text|The format in which log messages are printed (default: text, options: text, json)|
|--metrics-address|string|:7979|Specify where to serve the metrics and health check endpoint (default: :7979)|
|--log-level|string|logrus.InfoLevel.String()|Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal|
