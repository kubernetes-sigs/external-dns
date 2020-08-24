# Frequently asked questions

### How is ExternalDNS useful to me?

You've probably created many deployments. Typically, you expose your deployment to the Internet by creating a Service with `type=LoadBalancer`. Depending on your environment, this usually assigns a random publicly available endpoint to your service that you can access from anywhere in the world. On Google Container Engine, this is a public IP address:

```console
$ kubectl get svc
NAME      CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
nginx     10.3.249.226   35.187.104.85   80:32281/TCP   1m
```

But dealing with IPs for service discovery isn't nice, so you register this IP with your DNS provider under a better name—most likely, one that corresponds to your service name. If the IP changes, you update the DNS record accordingly.

Those times are over! ExternalDNS takes care of that last step for you by keeping your DNS records synchronized with your external entry points.

ExternalDNS' usefulness also becomes clear when you use Ingresses to allow external traffic into your cluster. Via Ingress, you can tell Kubernetes to route traffic to different services based on certain HTTP request attributes, e.g. the Host header:

```console
$ kubectl get ing
NAME         HOSTS                                      ADDRESS         PORTS     AGE
entrypoint   frontend.example.org,backend.example.org   35.186.250.78   80        1m
```

But there's nothing that actually makes clients resolve those hostnames to the Ingress' IP address. Again, you normally have to register each entry with your DNS provider. Only if you're lucky can you use a wildcard, like in the example above.

ExternalDNS can solve this for you as well.

### Which DNS providers are supported?

Currently, the following providers are supported:

- Google Cloud DNS
- AWS Route 53
- AzureDNS
- CloudFlare
- DigitalOcean
- DNSimple
- Infoblox
- Dyn
- OpenStack Designate
- PowerDNS
- CoreDNS
- Exoscale
- Oracle Cloud Infrastructure DNS
- Linode DNS
- RFC2136
- TransIP

As stated in the README, we are currently looking for stable maintainers for those providers, to ensure that bugfixes and new features will be available for all of those.

### Which Kubernetes objects are supported?

Services exposed via `type=LoadBalancer`, `type=ExternalName` and for the hostnames defined in Ingress objects as well as headless hostPort services. An initial effort to support type `NodePort` was started as of May 2018 and it is in progress at the time of writing.

### What are all of the CLI options that can be used to configure external-dns?
```console
$ external-dns --help
usage: external-dns --source=source --provider=provider [<flags>]

ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.

Note that all flags may be replaced with env vars - `--flag` -> `EXTERNAL_DNS_FLAG=1` or `--flag value` -> `EXTERNAL_DNS_FLAG=value`

Flags:
  --help                        Show context-sensitive help (also try --help-long and --help-man).
  --version                     Show application version.
  --server=""                   The Kubernetes API server to connect to (default: auto-detect)
  --kubeconfig=""               Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)
  --request-timeout=30s         Request timeout when calling Kubernetes APIs. 0s means no timeout
  --cf-api-endpoint=""          The fully-qualified domain name of the cloud foundry instance you are targeting
  --cf-username=""              The username to log into the cloud foundry API
  --cf-password=""              The password to log into the cloud foundry API
  --contour-load-balancer="heptio-contour/contour"
                                The fully-qualified name of the Contour load balancer service. (default: heptio-contour/contour)
  --skipper-routegroup-groupversion="zalando.org/v1"
                                The resource version for skipper routegroup
  --source=source ...           The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, node, fake, connector, istio-gateway,
                                istio-virtualservice, cloudfoundry, contour-ingressroute, crd, empty, skipper-routegroup,openshift-route)
  --namespace=""                Limit sources of endpoints to a specific namespace (default: all namespaces)
  --annotation-filter=""        Filter sources managed by external-dns via annotation using label selector semantics (default: all sources)
  --fqdn-template=""            A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source
                                (optional). Accepts comma separated list for multiple global FQDN.
  --combine-fqdn-annotation     Combine FQDN template and Annotations instead of overwriting
  --ignore-hostname-annotation  Ignore hostname annotation when generating DNS names, valid only when using fqdn-template is set (optional, default: false)
  --compatibility=              Process annotation semantics from legacy implementations (optional, options: mate, molecule)
  --publish-internal-services   Allow external-dns to publish DNS records for ClusterIP services (optional)
  --publish-host-ip             Allow external-dns to publish host-ip for headless services (optional)
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
  --provider=provider           The DNS provider where the DNS records will be created (required, options: aws, aws-sd, google, azure, azure-dns, azure-private-dns, cloudflare, rcodezero, digitalocean,
                                hetzner, dnsimple, akamai, infoblox, dyn, designate, coredns, skydns, inmemory, ovh, pdns, oci, exoscale, linode, rfc2136, ns1, transip, vinyldns, rdns, scaleway, vultr,
                                ultradns)
  --domain-filter= ...          Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)
  --exclude-domains= ...        Exclude subdomains (optional)
  --zone-id-filter= ...         Filter target zones by hosted zone id; specify multiple times for multiple zones (optional)
  --google-project=""           When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.
  --google-batch-change-size=1000
                                When using the Google provider, set the maximum number of changes that will be applied in each batch.
  --google-batch-change-interval=1s
                                When using the Google provider, set the interval between batch changes.
  --alibaba-cloud-config-file="/etc/kubernetes/alibaba-cloud.json"
                                When using the Alibaba Cloud provider, specify the Alibaba Cloud configuration file (required when --provider=alibabacloud
  --alibaba-cloud-zone-type=    When using the Alibaba Cloud provider, filter for zones of this type (optional, options: public, private)
  --aws-zone-type=              When using the AWS provider, filter for zones of this type (optional, options: public, private)
  --aws-zone-tags= ...          When using the AWS provider, filter for zones with these tags
  --aws-assume-role=""          When using the AWS provider, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns`
                                (optional)
  --aws-batch-change-size=1000  When using the AWS provider, set the maximum number of changes that will be applied in each batch.
  --aws-batch-change-interval=1s
                                When using the AWS provider, set the interval between batch changes.
  --aws-evaluate-target-health  When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)
  --aws-api-retries=3           When using the AWS provider, set the maximum number of retries for API calls before giving up.
  --aws-prefer-cname            When using the AWS provider, prefer using CNAME instead of ALIAS (default: disabled)
  --azure-config-file="/etc/kubernetes/azure.json"
                                When using the Azure provider, specify the Azure configuration file (required when --provider=azure
  --azure-resource-group=""     When using the Azure provider, override the Azure resource group to use (required when --provider=azure-private-dns)
  --azure-subscription-id=""    When using the Azure provider, specify the Azure configuration file (required when --provider=azure-private-dns)
  --azure-user-assigned-identity-client-id=""
                                When using the Azure provider, override the client id of user assigned identity in config file (optional)
  --cloudflare-proxied          When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)
  --cloudflare-zones-per-page=50
                                When using the Cloudflare provider, specify how many zones per page listed, max. possible 50 (default: 50)
  --coredns-prefix="/skydns/"   When using the CoreDNS provider, specify the prefix name
  --akamai-serviceconsumerdomain=""
                                When using the Akamai provider, specify the base URL (required when --provider=akamai)
  --akamai-client-token=""      When using the Akamai provider, specify the client token (required when --provider=akamai)
  --akamai-client-secret=""     When using the Akamai provider, specify the client secret (required when --provider=akamai)
  --akamai-access-token=""      When using the Akamai provider, specify the access token (required when --provider=akamai)
  --infoblox-grid-host=""       When using the Infoblox provider, specify the Grid Manager host (required when --provider=infoblox)
  --infoblox-wapi-port=443      When using the Infoblox provider, specify the WAPI port (default: 443)
  --infoblox-wapi-username="admin"
                                When using the Infoblox provider, specify the WAPI username (default: admin)
  --infoblox-wapi-password=""   When using the Infoblox provider, specify the WAPI password (required when --provider=infoblox)
  --infoblox-wapi-version="2.3.1"
                                When using the Infoblox provider, specify the WAPI version (default: 2.3.1)
  --infoblox-ssl-verify         When using the Infoblox provider, specify whether to verify the SSL certificate (default: true, disable with --no-infoblox-ssl-verify)
  --infoblox-view=""            DNS view (default: "")
  --infoblox-max-results=0      Add _max_results as query parameter to the URL on all API requests. The default is 0 which means _max_results is not set and the default of the server is used.
  --dyn-customer-name=""        When using the Dyn provider, specify the Customer Name
  --dyn-username=""             When using the Dyn provider, specify the Username
  --dyn-password=""             When using the Dyn provider, specify the pasword
  --dyn-min-ttl=DYN-MIN-TTL     Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.
  --oci-config-file="/etc/kubernetes/oci.yaml"
                                When using the OCI provider, specify the OCI configuration file (required when --provider=oci
  --rcodezero-txt-encrypt       When using the Rcodezero provider with txt registry option, set if TXT rrs are encrypted (default: false)
  --inmemory-zone= ...          Provide a list of pre-configured zones for the inmemory provider; specify multiple times for multiple zones (optional)
  --ovh-endpoint="ovh-eu"       When using the OVH provider, specify the endpoint (default: ovh-eu)
  --ovh-api-rate-limit=20       When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)
  --pdns-server="http://localhost:8081"
                                When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)
  --pdns-api-key=""             When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)
  --pdns-tls-enabled            When using the PowerDNS/PDNS provider, specify whether to use TLS (default: false, requires --tls-ca, optionally specify --tls-client-cert and --tls-client-cert-key)
  --ns1-endpoint=""             When using the NS1 provider, specify the URL of the API endpoint to target (default: https://api.nsone.net/v1/)
  --ns1-ignoressl               When using the NS1 provider, specify whether to verify the SSL certificate (default: false)
  --digitalocean-api-page-size=50
                                Configure the page size used when querying the DigitalOcean API.
  --tls-ca=""                   When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)
  --tls-client-cert=""          When using TLS communication, the path to the certificate to present as a client (not required for TLS)
  --tls-client-cert-key=""      When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)
  --exoscale-endpoint="https://api.exoscale.ch/dns"
                                Provide the endpoint for the Exoscale provider
  --exoscale-apikey=""          Provide your API Key for the Exoscale provider
  --exoscale-apisecret=""       Provide your API Secret for the Exoscale provider
  --rfc2136-host=""             When using the RFC2136 provider, specify the host of the DNS server
  --rfc2136-port=0              When using the RFC2136 provider, specify the port of the DNS server
  --rfc2136-zone=""             When using the RFC2136 provider, specify the zone entry of the DNS server to use
  --rfc2136-insecure            When using the RFC2136 provider, specify whether to attach TSIG or not (default: false, requires --rfc2136-tsig-keyname and rfc2136-tsig-secret)
  --rfc2136-tsig-keyname=""     When using the RFC2136 provider, specify the TSIG key to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-secret=""      When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-secret-alg=""  When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-tsig-axfr           When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)
  --rfc2136-min-ttl=0s          When using the RFC2136 provider, specify minimal TTL (in duration format) for records. This value will be used if the provided TTL for a service/ingress is lower than this
  --transip-account=""          When using the TransIP provider, specify the account name (required when --provider=transip)
  --transip-keyfile=""          When using the TransIP provider, specify the path to the private key file (required when --provider=transip)
  --policy=sync                 Modify how DNS records are synchronized between sources and providers (default: sync, options: sync, upsert-only, create-only)
  --registry=txt                The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, aws-sd)
  --txt-owner-id="default"      When using the TXT registry, a name that identifies this instance of ExternalDNS (default: default)
  --txt-prefix=""               When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional). Mutual exclusive with txt-suffix!
  --txt-suffix=""               When using the TXT registry, a custom string that's suffixed to the host portion of each ownership DNS record (optional). Mutual exclusive with txt-prefix!
  --txt-cache-interval=0s       The interval between cache synchronizations in duration format (default: disabled)
  --interval=1m0s               The interval between two consecutive synchronizations in duration format (default: 1m)
  --once                        When enabled, exits the synchronization loop after the first iteration (default: disabled)
  --dry-run                     When enabled, prints DNS record changes rather than actually performing them (default: disabled)
  --events                      When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)
  --log-format=text             The format in which log messages are printed (default: text, options: text, json)
  --metrics-address=":7979"     Specify where to serve the metrics and health check endpoint (default: :7979)
  --log-level=info              Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal
```

### How do I specify a DNS name for my Kubernetes objects?

There are three sources of information for ExternalDNS to decide on DNS name. ExternalDNS will pick one in order as listed below:

1. For ingress objects ExternalDNS will create a DNS record based on the host specified for the ingress object. For services ExternalDNS will look for the annotation `external-dns.alpha.kubernetes.io/hostname` on the service and use the corresponding value.

2. If compatibility mode is enabled (e.g. `--compatibility={mate,molecule}` flag), External DNS will parse annotations used by Zalando/Mate, wearemolecule/route53-kubernetes. Compatibility mode with Kops DNS Controller is planned to be added in the future.

3. If `--fqdn-template` flag is specified, e.g. `--fqdn-template={{.Name}}.my-org.com`, ExternalDNS will use service/ingress specifications for the provided template to generate DNS name.

### Can I specify multiple global FQDN templates?

Yes, you can. Pass in a comma separated list to `--fqdn-template`. Beaware this will double (triple, etc) the amount of DNS entries based on how many services, ingresses and so on you have and will get you faster towards the API request limit of your DNS provider.

### Which Service and Ingress controllers are supported?

Regarding Services, we'll support the OSI Layer 4 load balancers that Kubernetes creates on AWS and Google Container Engine, and possibly other clusters running on Google Compute Engine.

Regarding Ingress, we'll support:
* Google's Ingress Controller on GKE that integrates with their Layer 7 load balancers (GLBC)
* nginx-ingress-controller v0.9.x with a fronting Service
* Zalando's [AWS Ingress controller](https://github.com/zalando-incubator/kube-ingress-aws-controller), based on AWS ALBs and [Skipper](https://github.com/zalando/skipper)
* [Traefik](https://github.com/containous/traefik) 1.7 and above, when [`kubernetes.ingressEndpoint`](https://docs.traefik.io/v1.7/configuration/backends/kubernetes/#ingressendpoint) is configured (`kubernetes.ingressEndpoint.useDefaultPublishedService` in the [Helm chart](https://github.com/helm/charts/tree/HEAD/stable/traefik#configuration))

### Are other Ingress Controllers supported?

For Ingress objects, ExternalDNS will attempt to discover the target hostname of the relevant Ingress Controller automatically. If you are using an Ingress Controller that is not listed above you may have issues with ExternalDNS not discovering Endpoints and consequently not creating any DNS records. As a workaround, it is possible to force create an Endpoint by manually specifying a target host/IP for the records to be created by setting the annotation `external-dns.alpha.kubernetes.io/target` in the Ingress object.

Another reason you may want to override the ingress hostname or IP address is if you have an external mechanism for handling failover across ingress endpoints. Possible scenarios for this would include using [keepalived-vip](https://github.com/kubernetes/contrib/tree/HEAD/keepalived-vip) to manage failover faster than DNS TTLs might expire.

Note that if you set the target to a hostname, then a CNAME record will be created. In this case, the hostname specified in the Ingress object's annotation must already exist. (i.e. you have a Service resource for your Ingress Controller with the `external-dns.alpha.kubernetes.io/hostname` annotation set to the same value.)

### What about other projects similar to ExternalDNS?

ExternalDNS is a joint effort to unify different projects accomplishing the same goals, namely:

* Kops' [DNS Controller](https://github.com/kubernetes/kops/tree/HEAD/dns-controller)
* Zalando's [Mate](https://github.com/linki/mate)
* Molecule Software's [route53-kubernetes](https://github.com/wearemolecule/route53-kubernetes)

We strive to make the migration from these implementations a smooth experience. This means that, for some time, we'll support their annotation semantics in ExternalDNS and allow both implementations to run side-by-side. This enables you to migrate incrementally and slowly phase out the other implementation.

### How does it work with other implementations and legacy records?

ExternalDNS will allow you to opt into any Services and Ingresses that you want it to consider, by an annotation. This way, it can co-exist with other implementations running in the same cluster if they also support this pattern. However, we'll most likely declare ExternalDNS to be the default implementation. This means that ExternalDNS will consider Services and Ingresses that don't specifically declare which controller they want to be processed by; this is similar to the `ingress.class` annotation on GKE.

### I'm afraid you will mess up my DNS records!

ExternalDNS since v0.3 implements the concept of owning DNS records. This means that ExternalDNS will keep track of which records it has control over, and will never modify any records over which it doesn't have control. This is a fundamental requirement to operate ExternalDNS safely when there might be other actors creating DNS records in the same target space.

For now ExternalDNS uses TXT records to label owned records, and there might be other alternatives coming in the future releases.

### Does anyone use ExternalDNS in production?

Yes, multiple companies are using ExternalDNS in production. Zalando, as an example, has been using it in production since its v0.3 release, mostly using the AWS provider.

### How can we start using ExternalDNS?

Check out the following descriptive tutorials on how to run ExternalDNS in [GKE](tutorials/gke.md) and [AWS](tutorials/aws.md) or any other supported provider.

### Why is ExternalDNS only adding a single IP address in Route 53 on AWS when using the `nginx-ingress-controller`? How do I get it to use the FQDN of the ELB assigned to my `nginx-ingress-controller` Service instead?

By default the `nginx-ingress-controller` assigns a single IP address to an Ingress resource when it's created. ExternalDNS uses what's assigned to the Ingress resource, so it too will use this single IP address when adding the record in Route 53.

In most AWS deployments, you'll instead want the Route 53 entry to be the FQDN of the ELB that is assigned to the `nginx-ingress-controller` Service. To accomplish this, when you create the `nginx-ingress-controller` Deployment, you need to provide the `--publish-service` option to the `/nginx-ingress-controller` executable under `args`. Once this is deployed new Ingress resources will get the ELB's FQDN and ExternalDNS will use the same when creating records in Route 53.

According to the `nginx-ingress-controller` [docs](https://kubernetes.github.io/ingress-nginx/) the value you need to provide `--publish-service` is:

> Service fronting the ingress controllers. Takes the form namespace/name. The controller will set the endpoint records on the ingress objects to reflect those on the service.

For example if your `nginx-ingress-controller` Service's name is `nginx-ingress-controller-svc` and it's in the `default` namespace the start of your resource YAML might look like the following. Note the second to last line.

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-ingress-controller
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx-ingress
  template:
    metadata:
      labels:
        app: nginx-ingress
    spec:
      hostNetwork: false
      containers:
        - name: nginx-ingress-controller
          image: "gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.11"
          imagePullPolicy: "IfNotPresent"
          args:
            - /nginx-ingress-controller
            - --default-backend-service={your-backend-service}
            - --publish-service=default/nginx-ingress-controller-svc
            - --configmap={your-configmap}
```

### I have a Service/Ingress but it's ignored by ExternalDNS. Why?

ExternalDNS can be configured to only use Services or Ingresses as source. In case Services or Ingresses seem to be ignored in your setup, consider checking how the flag `--source` was configured when deployed. For reference, see the issue https://github.com/kubernetes-sigs/external-dns/issues/267.

### I'm using an ELB with TXT registry but the CNAME record clashes with the TXT record. How to avoid this?

CNAMEs cannot co-exist with other records, therefore you can use the `--txt-prefix` flag which makes sure to create a TXT record with a name following the pattern `prefix.<CNAME record>`. For reference, see the issue https://github.com/kubernetes-sigs/external-dns/issues/262.

### Can I force ExternalDNS to create CNAME records for ELB/ALB?

The default logic is: when a target looks like an ELB/ALB, ExternalDNS will create ALIAS records for it.
Under certain circumstances you want to force ExternalDNS to create CNAME records instead. If you want to do that, start ExternalDNS with the `--aws-prefer-cname` flag.

Why should I want to force ExternalDNS to create CNAME records for ELB/ALB? Some motivations of users were:

> "Our hosted zones records are synchronized with our enterprise DNS. The record type ALIAS is an AWS proprietary record type and AWS allows you to set a DNS record directly on AWS resources. Since this is not a DNS RfC standard and therefore can not be transferred and created in our enterprise DNS. So we need to force CNAME creation instead."

or

> "In case of ALIAS if we do nslookup with domain name, it will return only IPs of ELB. So it is always difficult for us to locate ELB in AWS console to which domain is pointing. If we configure it with CNAME it will return exact ELB CNAME, which is more helpful.!"

### Which permissions do I need when running ExternalDNS on a GCE or GKE node.

You need to add either https://www.googleapis.com/auth/ndev.clouddns.readwrite or https://www.googleapis.com/auth/cloud-platform on your instance group's scope.

### What metrics can I get from ExternalDNS and what do they mean?

ExternalDNS exposes 2 types of metrics: Sources and Registry errors.

`Source`s are mostly Kubernetes API objects. Examples of `source` errors may be connection errors to the Kubernetes API server itself or missing RBAC permissions. It can also stem from incompatible configuration in the objects itself like invalid characters, processing a broken fqdnTemplate, etc.

`Registry` errors are mostly Provider errors, unless there's some coding flaw in the registry package. Provider errors often arise due to accessing their APIs due to network or missing cloud-provider permissions when reading records. When applying a changeset, errors will arise if the changeset applied is incompatible with the current state.

In case of an increased error count, you could correlate them with the `http_request_duration_seconds{handler="instrumented_http"}` metric which should show increased numbers for status codes 4xx (permissions, configuration, invalid changeset) or 5xx (apiserver down).

You can use the host label in the metric to figure out if the request was against the Kubernetes API server (Source errors) or the DNS provider API (Registry/Provider errors).

### How can I run ExternalDNS under a specific GCP Service Account, e.g. to access DNS records in other projects?

Have a look at https://github.com/linki/mate/blob/v0.6.2/examples/google/README.md#permissions

### How do I configure multiple Sources via environment variables? (also applies to domain filters)

Separate the individual values via a line break. The equivalent of `--source=service --source=ingress` would be `service\ningress`. However, it can be tricky do define that depending on your environment. The following examples work (zsh):

Via docker:

```console
$ docker run \
  -e EXTERNAL_DNS_SOURCE=$'service\ningress' \
  -e EXTERNAL_DNS_PROVIDER=google \
  -e EXTERNAL_DNS_DOMAIN_FILTER=$'foo.com\nbar.com' \
  registry.opensource.zalan.do/teapot/external-dns:latest
time="2017-08-08T14:10:26Z" level=info msg="config: &{APIServerURL: KubeConfig: Sources:[service ingress] Namespace: ...
```

Locally:

```console
$ export EXTERNAL_DNS_SOURCE=$'service\ningress'
$ external-dns --provider=google
INFO[0000] config: &{APIServerURL: KubeConfig: Sources:[service ingress] Namespace: ...
```

```
$ EXTERNAL_DNS_SOURCE=$'service\ningress' external-dns --provider=google
INFO[0000] config: &{APIServerURL: KubeConfig: Sources:[service ingress] Namespace: ...
```

In a Kubernetes manifest:

```yaml
spec:
  containers:
  - name: external-dns
    args:
    - --provider=google
    env:
    - name: EXTERNAL_DNS_SOURCE
      value: "service\ningress"
```

Or preferably:

```yaml
spec:
  containers:
  - name: external-dns
    args:
    - --provider=google
    env:
    - name: EXTERNAL_DNS_SOURCE
      value: |-
        service
        ingress
```


### Running an internal and external dns service

Sometimes you need to run an internal and an external dns service.
The internal one should provision hostnames used on the internal network (perhaps inside a VPC), and the external
one to expose DNS to the internet.

To do this with ExternalDNS you can use the `--annotation-filter` to specifically tie an instance of ExternalDNS to
an instance of a ingress controller. Let's assume you have two ingress controllers `nginx-internal` and `nginx-external`
then you can start two ExternalDNS providers one with `--annotation-filter=kubernetes.io/ingress.class=nginx-internal`
and one with `--annotation-filter=kubernetes.io/ingress.class=nginx-external`.

### Can external-dns manage(add/remove) records in a hosted zone which is setup in different AWS account?

Yes, give it the correct cross-account/assume-role permissions and use the `--aws-assume-role` flag https://github.com/kubernetes-sigs/external-dns/pull/524#issue-181256561

### How do I provide multiple values to the annotation `external-dns.alpha.kubernetes.io/hostname`?

Separate them by `,`.


### Are there official Docker images provided?

When we tag a new release, we push a Docker image on Zalando's public Docker registry with the following name: 

```
registry.opensource.zalan.do/teapot/external-dns
```

As tags, you can use your version of choice or use `latest` that always resolves to the latest tag.

If you wish to build your own image, you can use the provided [Dockerfile](../Dockerfile) as a starting point.

We are currently working with the Kubernetes community to provide official images for the project similarly to what is done with the other official Kubernetes projects, but we don't have an ETA on when those images will be available.

### Why am I seeing time out errors even though I have connectivity to my cluster?

If you're seeing an error such as this:
```
FATA[0060] failed to sync cache: timed out waiting for the condition
```

You may not have the correct permissions required to query all the necessary resources in your kubernetes cluster. Specifically, you may be running in a `namespace` that you don't have these permissions in. By default, commands are run against the `default` namespace. Try changing this to your particular namespace to see if that fixes the issue.
