module sigs.k8s.io/external-dns

go 1.16

require (
	cloud.google.com/go v0.97.0
	git.blindage.org/21h/hcloud-dns v0.0.0-20200807003420-f768ffe03f8d
	github.com/Azure/azure-sdk-for-go v46.4.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.21
	github.com/Azure/go-autorest/autorest/adal v0.9.16
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/akamai/AkamaiOPEN-edgegrid-golang v1.1.1
	github.com/StackExchange/dnscontrol v0.2.8
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38 // indirect
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/kingpin v2.2.5+incompatible
	github.com/alecthomas/repr v0.0.0-20200325044227-4184120f674c // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.357
	github.com/aws/aws-sdk-go v1.40.53
	github.com/bodgit/tsig v0.0.2
	github.com/cloudflare/cloudflare-go v0.13.2
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20190201205600-f136f9222381
	github.com/datawire/ambassador v1.6.0
	github.com/denverdino/aliyungo v0.0.0-20190125010748-a747050bb1ba
	github.com/digitalocean/godo v1.69.1
	github.com/dnsimple/dnsimple-go v0.60.0
	github.com/exoscale/egoscale v0.73.2
	github.com/ffledgling/pdns-go v0.0.0-20180219074714-524e7daccd99
	github.com/go-gandi/go-gandi v0.0.0-20200921091836-0d8a64b9cc09
	github.com/go-logr/logr v1.1.0 // indirect
	github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	github.com/google/go-cmp v0.5.6
	github.com/gophercloud/gophercloud v0.21.0
	github.com/hooklift/gowsdl v0.5.0
	github.com/infobloxopen/infoblox-go-client v1.1.1
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/linki/instrumented_http v0.3.0
	github.com/linode/linodego v0.32.2
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/maxatome/go-testdeep v1.10.1
	github.com/miekg/dns v1.1.36-0.20210109083720-731b191cabd1
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/nesv/go-dynect v0.6.0
	github.com/nic-at/rc0go v1.1.1
	github.com/onsi/gomega v1.14.0 // indirect
	github.com/openshift/api v0.0.0-20200605231317-fb2a6ca106ae
	github.com/openshift/client-go v0.0.0-20200608144219-584632b8fc73
	github.com/oracle/oci-go-sdk v21.4.0+incompatible
	github.com/ovh/go-ovh v0.0.0-20181109152953-ba5adb4cf014
	github.com/pkg/errors v0.9.1
	github.com/projectcontour/contour v1.18.1
	github.com/prometheus/client_golang v1.11.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.7.0.20210127161313-bd30bebeac4f
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/gunit v1.3.4 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/terra-farm/udnssdk v1.3.5 // indirect
	github.com/transip/gotransip/v6 v6.6.2
	github.com/ultradns/ultradns-sdk-go v0.0.0-20200616202852-e62052662f60
	github.com/vinyldns/go-vinyldns v0.0.0-20200211145900-fe8a3d82e556
	github.com/vultr/govultr/v2 v2.9.0
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/api v0.58.0
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190322154155-0dafb5275fd1
	gopkg.in/yaml.v2 v2.4.0
	istio.io/api v0.0.0-20210128181506-0c4b8e54850f
	istio.io/client-go v0.0.0-20210128182905-ee2edd059e02
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/klog/v2 v2.20.0 // indirect
	k8s.io/utils v0.0.0-20210820185131-d34e5cb4466e // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace k8s.io/klog/v2 => github.com/Raffo/knolog v0.0.0-20211016155154-e4d5e0cc970a
