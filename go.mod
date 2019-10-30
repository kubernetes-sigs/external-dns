module github.com/kubernetes-sigs/external-dns

go 1.13

require (
	cloud.google.com/go v0.37.4
	github.com/Azure/azure-sdk-for-go v10.0.4-beta+incompatible
	github.com/Azure/go-autorest v10.9.0+incompatible
	github.com/alecthomas/kingpin v2.2.5+incompatible
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20180828111155-cad214d7d71f
	github.com/aws/aws-sdk-go v1.23.18
	github.com/cloudflare/cloudflare-go v0.10.1
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20190201205600-f136f9222381
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/denverdino/aliyungo v0.0.0-20180815121905-69560d9530f5
	github.com/digitalocean/godo v1.19.0
	github.com/dnsimple/dnsimple-go v0.14.0
	github.com/exoscale/egoscale v0.18.1
	github.com/ffledgling/pdns-go v0.0.0-20180219074714-524e7daccd99
	github.com/gophercloud/gophercloud v0.0.0-20190126172459-c818fa66e4c8
	github.com/heptio/contour v0.13.0
	github.com/infobloxopen/infoblox-go-client v0.0.0-20180606155407-61dc5f9b0a65
	github.com/kubernetes-incubator/external-dns v0.5.17
	github.com/linki/instrumented_http v0.2.0
	github.com/linode/linodego v0.3.0
	github.com/miekg/dns v1.0.8
	github.com/nesv/go-dynect v0.6.0
	github.com/nic-at/rc0go v1.1.0
	github.com/oracle/oci-go-sdk v1.8.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/sanyu/dynectsoap v0.0.0-20181203081243-b83de5edc4e0
	github.com/sirupsen/logrus v1.4.1
	github.com/stretchr/testify v1.4.0
	github.com/transip/gotransip v5.8.2+incompatible
	github.com/vinyldns/go-vinyldns v0.0.0-20190611170422-7119fe55ed92
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	google.golang.org/api v0.3.1
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190322154155-0dafb5275fd1
	gopkg.in/yaml.v2 v2.2.2
	istio.io/api v0.0.0-20190321180614-db16d82d3672
	istio.io/istio v0.0.0-20190322063008-2b1331886076
	k8s.io/api v0.0.0-20190503184017-f1b257a4ce96
	k8s.io/apimachinery v0.0.0-20190223001710-c182ff3b9841
	k8s.io/client-go v8.0.0+incompatible
)

replace k8s.io/code-generator v0.0.0-20190409092313-b1289fc74931 => k8s.io/code-generator v0.0.0-20181128191024-b1289fc74931

replace github.com/golang/glog => github.com/kubermatic/glog-logrus v0.0.0-20180829085450-3fa5b9870d1d
