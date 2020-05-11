module sigs.k8s.io/external-dns

go 1.14

require (
	cloud.google.com/go v0.44.3
	github.com/Azure/azure-sdk-for-go v36.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal v0.6.0
	github.com/Azure/go-autorest/autorest/azure/auth v0.0.0-00010101000000-000000000000
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/akamai/AkamaiOPEN-edgegrid-golang v0.9.11
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38 // indirect
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/kingpin v2.2.5+incompatible
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/ultradns/ultradns-sdk-go v1.3.7
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20180828111155-cad214d7d71f
	github.com/aws/aws-sdk-go v1.27.4
	github.com/cloudflare/cloudflare-go v0.10.1
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20190201205600-f136f9222381
	github.com/denverdino/aliyungo v0.0.0-20180815121905-69560d9530f5
	github.com/digitalocean/godo v1.34.0
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/dnsimple/dnsimple-go v0.14.0
	github.com/exoscale/egoscale v0.18.1
	github.com/ffledgling/pdns-go v0.0.0-20180219074714-524e7daccd99
	github.com/go-resty/resty v1.8.0 // indirect
	github.com/gobs/pretty v0.0.0-20180724170744-09732c25a95b // indirect
	github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	github.com/gophercloud/gophercloud v0.1.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/heptio/contour v0.15.0
	github.com/infobloxopen/infoblox-go-client v0.0.0-20180606155407-61dc5f9b0a65
	github.com/linki/instrumented_http v0.2.0
	github.com/linode/linodego v0.3.0
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/miekg/dns v1.1.25
	github.com/nesv/go-dynect v0.6.0
	github.com/nic-at/rc0go v1.1.0
	github.com/openshift/api v0.0.0-20190322043348-8741ff068a47
	github.com/openshift/client-go v3.9.0+incompatible
	github.com/oracle/oci-go-sdk v1.8.0
	github.com/ovh/go-ovh v0.0.0-20181109152953-ba5adb4cf014
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.0.0
	github.com/sanyu/dynectsoap v0.0.0-20181203081243-b83de5edc4e0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/go-aws-auth v0.0.0-20180515143844-0c1422d1fdb9 // indirect
	github.com/smartystreets/gunit v1.1.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/transip/gotransip v5.8.2+incompatible
	github.com/vinyldns/go-vinyldns v0.0.0-20190611170422-7119fe55ed92
	github.com/vultr/govultr v0.3.2
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200401174654-e694b7bb0875
	golang.org/x/net v0.0.0-20190923162816-aa69164e4478
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.9.0
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190322154155-0dafb5275fd1
	gopkg.in/yaml.v2 v2.2.5
	istio.io/api v0.0.0-20190820204432-483f2547d882
	istio.io/istio v0.0.0-20190322063008-2b1331886076
	k8s.io/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/klog v0.3.1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.1+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.1
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.6.0
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.3.0
	github.com/golang/glog => github.com/kubermatic/glog-logrus v0.0.0-20180829085450-3fa5b9870d1d
	istio.io/api => istio.io/api v0.0.0-20190820204432-483f2547d882
	istio.io/istio => istio.io/istio v0.0.0-20190911205955-c2bd59595ce6
	k8s.io/api => k8s.io/api v0.0.0-20190817221950-ebce17126a01
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190919022157-e8460a76b3ad
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817221809-bf4de9df677c
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190817222206-ee6c071a42cf
	k8s.io/klog => github.com/mikkeloscar/knolog v0.0.0-20190326191552-80742771eb6b
)
