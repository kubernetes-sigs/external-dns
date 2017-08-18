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
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"strconv"
	log "github.com/Sirupsen/logrus"
)

const (
	// The annotation used for figuring out which controller is responsible
	controllerAnnotationKey = "external-dns.alpha.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	hostnameAnnotationKey = "external-dns.alpha.kubernetes.io/hostname"
	// The annotation used for defining the desired ingress target
	targetAnnotationKey = "external-dns.alpha.kubernetes.io/target"
	// The annotation used for defining the desired DNS record TTL
	ttlAnnotationKey = "external-dns.alpha.kubernetes.io/ttl"
	// The value of the controller annotation so that we feel responsible
	controllerAnnotationValue = "dns-controller"
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints() ([]*endpoint.Endpoint, error)
}

func getTTLFromAnnotations(annotations map[string]string) endpoint.TTL {
	ttlAnnotation, exists := annotations[ttlAnnotationKey]
	if !exists {
		return endpoint.TTL(0)
	}
	ttlValue, err := strconv.ParseInt(ttlAnnotation, 10, 64)
	if err != nil {
		log.Warnf("%v is not a valid TTL value", ttlAnnotation)
		return endpoint.TTL(0)
	}
	if ttlValue < 0 {
		log.Warnf("TTL must be a non-negative integer", ttlAnnotation)
		return endpoint.TTL(0)
	}
	return endpoint.TTL(ttlValue)
}
