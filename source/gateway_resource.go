/*
Copyright 2021 The Kubernetes Authors.

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
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	gwinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
	srctypes "sigs.k8s.io/external-dns/source/types"
)

func NewGatewaySource(ctx context.Context, clients ClientGenerator, config *Config) (Source, error) {
	return newGatewaySource(ctx, clients, config, newGatewayInformerFactory)
}

// +externaldns:source:name=gateway
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API Gateway resources annotated with external-dns.kubernetes.io/hostname
// +externaldns:source:resources=Gateway.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type gatewaySource struct {
	gwName        string
	gwNamespace   string
	gwLabels      labels.Selector
	gwAnnotations labels.Selector
	gwInformer    informers_v1.GatewayInformer

	templateEngine           template.Engine
	ignoreHostnameAnnotation bool
}

func newGatewaySource(
	ctx context.Context,
	clients ClientGenerator,
	config *Config,
	newInformerFactory func(gateway.Interface, string, labels.Selector) gwinformers.SharedInformerFactory,
) (Source, error) {
	gwLabels := config.LabelFilter
	if gwLabels == nil {
		gwLabels = labels.Everything()
	}
	gwAnnotations := config.AnnotationFilter
	if gwAnnotations == nil {
		gwAnnotations = labels.Everything()
	}

	client, err := clients.GatewayClient()
	if err != nil {
		return nil, err
	}

	gwInformerFactory := newInformerFactory(client, config.Namespace(), gwLabels)
	gwInformer := gwInformerFactory.Gateway().V1().Gateways()
	informers.MustSetTransform(gwInformer.Informer(), informers.TransformerWithOptions[*v1.Gateway](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
		informers.TransformRemoveStatusConditions(),
	))

	gwInformerFactory.Start(ctx.Done())
	if err := informers.WaitForCacheSync(ctx, gwInformerFactory); err != nil {
		return nil, err
	}

	return &gatewaySource{
		gwName:                   config.GatewayName,
		gwNamespace:              config.Namespace(),
		gwLabels:                 gwLabels,
		gwAnnotations:            gwAnnotations,
		gwInformer:               gwInformer,
		templateEngine:           config.TemplateEngine,
		ignoreHostnameAnnotation: config.IgnoreHostnameAnnotation,
	}, nil
}

func (src *gatewaySource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handlers for Gateway")
	informers.MustAddEventHandler(src.gwInformer.Informer(), eventHandlerFunc(handler))
}

func (src *gatewaySource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	gateways, err := src.gwInformer.Lister().Gateways(src.gwNamespace).List(src.gwLabels)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	for _, gw := range gateways {
		if src.gwName != "" && src.gwName != gw.Name {
			continue
		}

		meta := &gw.ObjectMeta
		annots := meta.Annotations

		if !src.gwAnnotations.Matches(labels.Set(annots)) {
			continue
		}

		if annotations.IsControllerMismatch(meta, gatewayKind) {
			continue
		}

		hostnames, err := src.hostnames(gw)
		if err != nil {
			return nil, err
		}
		if len(hostnames) == 0 {
			log.Debugf("No endpoints could be generated from Gateway %s/%s", gw.Namespace, gw.Name)
			continue
		}

		targets := src.targets(gw)
		if len(targets) == 0 {
			log.Debugf("No targets found for Gateway %s/%s", gw.Namespace, gw.Name)
			continue
		}

		resource := fmt.Sprintf("gateway/%s/%s", gw.Namespace, gw.Name)
		providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(annots)
		ttl := annotations.TTLFromAnnotations(annots, resource)

		var gwEndpoints []*endpoint.Endpoint
		for _, host := range hostnames {
			gwEndpoints = append(gwEndpoints, endpoint.EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
		if endpoint.HasNoEmptyEndpoints(gwEndpoints, srctypes.Gateway, gw) {
			continue
		}
		log.Debugf("Endpoints generated from Gateway %s/%s: %v", gw.Namespace, gw.Name, gwEndpoints)
		endpoint.AttachRefObject(gwEndpoints, events.NewObjectReference(gw, srctypes.Gateway))
		endpoints = append(endpoints, gwEndpoints...)
	}
	return endpoint.MergeEndpoints(endpoints), nil
}

func (src *gatewaySource) hostnames(gw *v1.Gateway) ([]string, error) {
	var hostnames []string
	if !src.ignoreHostnameAnnotation {
		hostnames = append(hostnames, annotations.HostnamesFromAnnotations(gw.Annotations)...)
	}
	return appendFQDNTemplate(src.templateEngine, hostnames, gw)
}

func (src *gatewaySource) targets(gw *v1.Gateway) endpoint.Targets {
	override := annotations.TargetsFromTargetAnnotation(gw.Annotations)
	if len(override) > 0 {
		return override
	}
	var ips, hostnames endpoint.Targets
	for _, addr := range gw.Status.Addresses {
		if addr.Value == "" {
			continue
		}
		if addr.Type == nil || *addr.Type == v1.IPAddressType {
			ips = append(ips, addr.Value)
		} else if *addr.Type == v1.HostnameAddressType {
			hostnames = append(hostnames, addr.Value)
		}
	}
	if len(ips) > 0 {
		return ips
	}
	return hostnames
}
