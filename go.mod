module sigs.k8s.io/external-dns

go 1.13

require (
	cloud.google.com/go v0.50.0
	github.com/Azure/azure-sdk-for-go v36.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.4
	github.com/Azure/go-autorest/autorest/adal v0.6.0
	github.com/Azure/go-autorest/autorest/azure/auth v0.0.0-00010101000000-000000000000
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/akamai/AkamaiOPEN-edgegrid-golang v0.9.5
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38 // indirect
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/kingpin v2.2.5+incompatible
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20180828111155-cad214d7d71f
	github.com/aws/aws-sdk-go v1.27.4
	github.com/cloudflare/cloudflare-go v0.10.1
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20190201205600-f136f9222381
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/denverdino/aliyungo v0.0.0-20180815121905-69560d9530f5
	github.com/digitalocean/godo v1.19.0
	github.com/dnsimple/dnsimple-go v0.14.0
	github.com/exoscale/egoscale v0.18.1
	github.com/ffledgling/pdns-go v0.0.0-20180219074714-524e7daccd99
	github.com/go-resty/resty v1.8.0 // indirect
	github.com/gobs/pretty v0.0.0-20180724170744-09732c25a95b // indirect
	github.com/gophercloud/gophercloud v0.1.0
	github.com/infobloxopen/infoblox-go-client v0.0.0-20180606155407-61dc5f9b0a65
	github.com/linki/instrumented_http v0.2.0
	github.com/linode/linodego v0.3.0
	github.com/miekg/dns v1.1.4
	github.com/nesv/go-dynect v0.6.0
	github.com/nic-at/rc0go v1.1.0
	github.com/oracle/oci-go-sdk v1.8.0
	github.com/pkg/errors v0.8.1
	github.com/projectcontour/contour v1.0.0
	github.com/prometheus/client_golang v1.1.0
	github.com/sanyu/dynectsoap v0.0.0-20181203081243-b83de5edc4e0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/go-aws-auth v0.0.0-20180515143844-0c1422d1fdb9 // indirect
	github.com/smartystreets/gunit v1.1.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/transip/gotransip v5.8.2+incompatible
	github.com/vinyldns/go-vinyldns v0.0.0-20190611170422-7119fe55ed92
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.15.0
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190322154155-0dafb5275fd1
	gopkg.in/yaml.v2 v2.2.8
	istio.io/api v0.0.0-20200227124350-5ce1c9cc553f
	istio.io/istio v0.0.0-20200227051249-587f118ac5da
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	k8s.io/kube-proxy v0.16.7 // indirect
	k8s.io/kubernetes v1.16.7
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.1+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.1
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.6.0
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.3.0
	github.com/golang/glog => github.com/kubermatic/glog-logrus v0.0.0-20180829085450-3fa5b9870d1d
	github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.14
	k8s.io/klog => github.com/mikkeloscar/knolog v0.0.0-20190326191552-80742771eb6b
	k8s.io/api => k8s.io/api v0.16.7
	k8s.io/api-machinery => k8s.io/api-machinery v0.16.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.7
	k8s.io/apiserver => k8s.io/apiserver v0.16.7
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.7
	k8s.io/client-go => k8s.io/client-go v0.16.7
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.7
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.7
	k8s.io/code-generator => k8s.io/code-generator v0.16.7
	k8s.io/component-base => k8s.io/component-base v0.16.7
	k8s.io/cri-api => k8s.io/cri-api v0.16.7
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.7
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.7
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.7
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.7
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.7
	k8s.io/kubectl => k8s.io/kubectl v0.16.7
	k8s.io/kubelet => k8s.io/kubelet v0.16.7
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.7
	k8s.io/metrics => k8s.io/metrics v0.16.7
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.7
)
