package source

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

const (
	// The annotation used for figuring out which controller is responsible
	controllerAnnotationKey = "external-dns.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	hostnameAnnotationKey = "external-dns.kubernetes.io/hostname"
	// The value of the controller annotation so that we feel resposible
	controllerAnnotationValue = "dns-controller"
)

// ServiceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' external entrypoints it will return a corresponding
// Endpoint object.
type ServiceSource struct {
	Client kubernetes.Interface
}

// Endpoints returns endpoint objects for each service that should be processed.
func (sc *ServiceSource) Endpoints() ([]endpoint.Endpoint, error) {
	services, err := sc.Client.Core().Services(v1.NamespaceAll).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []endpoint.Endpoint{}

	for _, svc := range services.Items {
		// Check controller annotation to see if we are responsible.
		controller, exists := svc.Annotations[controllerAnnotationKey]
		if exists && controller != controllerAnnotationValue {
			continue
		}

		// Get the desired hostname of the service from the annotation.
		hostname, exists := svc.Annotations[hostnameAnnotationKey]
		if !exists {
			continue
		}

		// Create an endpoint matching the desired hostname.
		endpoint := endpoint.Endpoint{
			DNSName: hostname,
		}

		// Create a corresponding endpoint for each configured external entrypoint.
		for _, lb := range svc.Status.LoadBalancer.Ingress {
			endpoint.Target = lb.IP
			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints, nil
}
