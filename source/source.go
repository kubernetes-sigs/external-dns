/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package source

import (
	"net"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

const (
	// The annotation used for figuring out which controller is responsible
	controllerAnnotationKey = "external-dns.alpha.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	hostnameAnnotationKey = "external-dns.alpha.kubernetes.io/hostname"
	// The annotation used for defining the desired ingress target
	targetAnnotationKey = "external-dns.alpha.kubernetes.io/target"
	// The annotation used for defining the desired weight scope for the record
	weightScopeAnnotationKey = "external-dns.alpha.kubernetes.io/weight-scope"
	// The annotation used for defining the desired weight for the record
	awsRoute53WeightAnnotationKey = "external-dns.alpha.kubernetes.io/aws-route53-weight"
	// The value of the suffix for the weighted policy id
	awsRoute53SetIdentifierAnnotationKey = "external-dns.alpha.kubernetes.io/aws-route53-set-identifier"
	// The value of the controller annotation so that we feel responsible
	controllerAnnotationValue = "dns-controller"
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints() ([]*endpoint.Endpoint, error)
}

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A for IPs and type CNAME for everything else.
func suitableType(target string) string {
	if net.ParseIP(target) != nil {
		return endpoint.RecordTypeA
	}
	return endpoint.RecordTypeCNAME
}
