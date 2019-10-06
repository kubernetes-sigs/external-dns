module github.com/kubernetes-incubator/external-dns

go 1.13

require (
	cloud.google.com/go v0.37.4
	github.com/Azure/azure-sdk-for-go v10.0.4-beta+incompatible
	github.com/Azure/go-autorest v10.9.0+incompatible
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38 // indirect
	github.com/alecthomas/colour v0.0.0-20160524082231-60882d9e2721 // indirect
	github.com/alecthomas/kingpin v2.2.5+incompatible
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20180828111155-cad214d7d71f
	github.com/aws/aws-sdk-go v1.23.18
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.10.1
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20190201205600-f136f9222381
	github.com/coreos/bbolt v1.3.2 // indirect
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/coreos/go-semver v0.2.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/denverdino/aliyungo v0.0.0-20180815121905-69560d9530f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/digitalocean/godo v1.19.0
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/dnsimple/dnsimple-go v0.14.0
	github.com/exoscale/egoscale v0.18.1
	github.com/ffledgling/pdns-go v0.0.0-20180219074714-524e7daccd99
	github.com/go-resty/resty v1.8.0 // indirect
	github.com/gobs/pretty v0.0.0-20180724170744-09732c25a95b // indirect
	github.com/gophercloud/gophercloud v0.0.0-20190126172459-c818fa66e4c8
	github.com/gorilla/websocket v0.0.0-20170926233335-4201258b820c // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v0.0.0-20190222133341-cfaf5686ec79 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.8.5 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/heptio/contour v0.13.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/infobloxopen/infoblox-go-client v0.0.0-20180606155407-61dc5f9b0a65
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/linki/instrumented_http v0.2.0
	github.com/linode/linodego v0.3.0
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/miekg/dns v1.0.8
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nesv/go-dynect v0.6.0
	github.com/nic-at/rc0go v1.1.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/oracle/oci-go-sdk v1.8.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/sanyu/dynectsoap v0.0.0-20181203081243-b83de5edc4e0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/smartystreets/go-aws-auth v0.0.0-20180515143844-0c1422d1fdb9 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/smartystreets/gunit v1.0.2 // indirect
	github.com/soheilhy/cmux v0.1.3 // indirect
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20170815181823-89b8d40f7ca8 // indirect
	github.com/transip/gotransip v5.8.2+incompatible
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	github.com/vinyldns/go-vinyldns v0.0.0-20190611170422-7119fe55ed92
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/yl2chen/cidranger v0.0.0-20180214081945-928b519e5268 // indirect
	go.etcd.io/bbolt v1.3.2 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	google.golang.org/api v0.3.1
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190322154155-0dafb5275fd1
	gopkg.in/yaml.v2 v2.2.2
	istio.io/api v0.0.0-20190321180614-db16d82d3672
	istio.io/istio v0.0.0-20190322063008-2b1331886076
	k8s.io/api v0.0.0-20190503184017-f1b257a4ce96
	k8s.io/apiextensions-apiserver v0.0.0-20190503184539-c338b28ceaa1 // indirect
	k8s.io/apimachinery v0.0.0-20190223001710-c182ff3b9841
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190401085232-94e1e7b7574c // indirect
)

replace k8s.io/code-generator v0.0.0-20190409092313-b1289fc74931 => k8s.io/code-generator v0.0.0-20181128191024-b1289fc74931

replace github.com/golang/glog => github.com/kubermatic/glog-logrus v0.0.0-20180829085450-3fa5b9870d1d
