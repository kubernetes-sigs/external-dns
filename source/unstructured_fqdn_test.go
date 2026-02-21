/*
Copyright 2026 The Kubernetes Authors.

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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	discoveryfake "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/dynamic"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
)

func TestUnstructuredFqdnTemplatingExamples(t *testing.T) {
	type cfg struct {
		resources          []string
		fqdnTemplate       string
		targetFqdnTemplate string
		hostTargetTemplate string
		labelFilter        string
	}
	for _, tt := range []struct {
		title    string
		cfg      cfg
		objects  []*unstructured.Unstructured
		expected []*endpoint.Endpoint
	}{
		{
			title: "ConfigMap with comma-separated hostnames",
			cfg: cfg{
				resources:          []string{"configmaps.v1"},
				fqdnTemplate:       `{{index .Object.data "hostnames"}}`,
				targetFqdnTemplate: `{{index .Object.data "target"}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "v1",
						"kind":       "ConfigMap",
						"metadata": map[string]any{
							"name":      "multi-dns",
							"namespace": "default",
							"annotations": map[string]any{
								annotations.ControllerKey: annotations.ControllerValue,
							},
						},
						"data": map[string]any{
							"hostnames": "entry1.internal.tld,entry2.example.tld",
							"target":    "10.10.10.10",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("entry1.internal.tld", endpoint.RecordTypeA, "10.10.10.10").
					WithLabel(endpoint.ResourceLabelKey, "configmap/default/multi-dns"),
				endpoint.NewEndpoint("entry2.example.tld", endpoint.RecordTypeA, "10.10.10.10").
					WithLabel(endpoint.ResourceLabelKey, "configmap/default/multi-dns"),
			},
		},
		{
			title: "with IP address",
			cfg: cfg{
				resources:          []string{"virtualmachineinstances.v1.kubevirt.io"},
				fqdnTemplate:       `{{.Name}}.{{index .Status.interfaces 0 "name"}}.vmi.com`,
				targetFqdnTemplate: `{{index .Status.interfaces 0 "ipAddress"}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "default",
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{
									"ipAddress": "10.244.1.50",
									"name":      "main",
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vm.main.vmi.com", endpoint.RecordTypeA, "10.244.1.50").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/my-vm"),
			},
		},
		{
			title: "Crossplane RDSInstance with endpoint",
			cfg: cfg{
				resources:          []string{"rdsinstances.v1alpha1.rds.aws.crossplane.io"},
				fqdnTemplate:       "{{.Name}}.db.example.com",
				targetFqdnTemplate: "{{.Status.atProvider.endpoint.address}}",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "rds.aws.crossplane.io/v1alpha1",
						"kind":       "RDSInstance",
						"metadata": map[string]any{
							"name":      "prod-postgres",
							"namespace": "default",
						},
						"status": map[string]any{
							"atProvider": map[string]any{
								"endpoint": map[string]any{
									"address": "prod-postgres.abc123.us-east-1.rds",
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("prod-postgres.db.example.com", endpoint.RecordTypeCNAME, "prod-postgres.abc123.us-east-1.rds.").
					WithLabel(endpoint.ResourceLabelKey, "rdsinstance/default/prod-postgres"),
			},
		},
		{
			title: "multiple VirtualMachineInstances",
			cfg: cfg{
				resources:          []string{"virtualmachineinstances.v1.kubevirt.io"},
				fqdnTemplate:       "{{.Name}}.vmi.example.com",
				targetFqdnTemplate: `{{index .Status.interfaces 0 "ipAddress"}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "vm-one",
							"namespace": "default",
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{"ipAddress": "10.244.1.10"},
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "vm-two",
							"namespace": "default",
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{"ipAddress": "10.244.1.20"},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("vm-one.vmi.example.com", endpoint.RecordTypeA, "10.244.1.10").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/vm-one"),
				endpoint.NewEndpoint("vm-two.vmi.example.com", endpoint.RecordTypeA, "10.244.1.20").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/vm-two"),
			},
		},
		{
			title: "multiple hosts from template",
			cfg: cfg{
				resources:          []string{"proxyservices.v1beta1.proxyconfigs.acme.corp"},
				fqdnTemplate:       "{{.Name}}.mesh.com,{{.Name}}.internal.com",
				targetFqdnTemplate: "{{index .Spec.hosts 0}}",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "proxyconfigs.acme.corp/v1beta1",
						"kind":       "ProxyService",
						"metadata": map[string]any{
							"name":      "reviews",
							"namespace": "default",
						},
						"spec": map[string]any{
							"hosts": []any{
								"promo.svc.local",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("reviews.internal.com", endpoint.RecordTypeCNAME, "promo.svc.local").
					WithLabel(endpoint.ResourceLabelKey, "proxyservice/default/reviews"),
				endpoint.NewEndpoint("reviews.mesh.com", endpoint.RecordTypeCNAME, "promo.svc.local").
					WithLabel(endpoint.ResourceLabelKey, "proxyservice/default/reviews"),
			},
		},
		{
			title: "with labels",
			cfg: cfg{
				resources:          []string{"applications.v1alpha1.argoproj.io"},
				fqdnTemplate:       `{{index .Labels "app.kubernetes.io/instance"}}.apps.com`,
				targetFqdnTemplate: "{{.Status.loadBalancer}}",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "argoproj.io/v1alpha1",
						"kind":       "Application",
						"metadata": map[string]any{
							"name":      "guestbook",
							"namespace": "default",
							"labels": map[string]any{
								"app.kubernetes.io/instance": "guestbook-prod",
							},
						},
						"status": map[string]any{
							"loadBalancer": "lb.example.com",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("guestbook-prod.apps.com", endpoint.RecordTypeCNAME, "lb.example.com").
					WithLabel(endpoint.ResourceLabelKey, "application/default/guestbook"),
			},
		},
		{
			title: "with ttl annotation set",
			cfg: cfg{
				resources:          []string{"applications.v1alpha1.argoproj.io"},
				fqdnTemplate:       `{{index .Labels "app.kubernetes.io/instance"}}.apps.com`,
				targetFqdnTemplate: "{{.Status.loadBalancer}}",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "argoproj.io/v1alpha1",
						"kind":       "Application",
						"metadata": map[string]any{
							"name":      "guestbook",
							"namespace": "ns",
							"labels": map[string]any{
								"app.kubernetes.io/instance": "guestbook-prod",
							},
							"annotations": map[string]any{
								annotations.TtlKey: "300",
							},
						},
						"status": map[string]any{
							"loadBalancer": "lb.example.com",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("guestbook-prod.apps.com", endpoint.RecordTypeCNAME, 300, "lb.example.com").
					WithLabel(endpoint.ResourceLabelKey, "application/ns/guestbook"),
			},
		},
		{
			title: "two different resource types - VirtualMachineInstance and RDSInstance",
			cfg: cfg{
				resources: []string{
					"virtualmachineinstances.v1.kubevirt.io",
					"rdsinstances.v1alpha1.rds.aws.crossplane.io",
				},
				fqdnTemplate: "{{.Name}}.{{.Namespace}}.com",
				targetFqdnTemplate: `
{{if .Status.interfaces}}{{index .Status.interfaces 0 "ipAddress"}}{{else}}{{.Status.atProvider.endpoint.address}}{{end}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "vms",
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{
									"ipAddress": "10.244.1.100",
								},
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "rds.aws.crossplane.io/v1alpha1",
						"kind":       "RDSInstance",
						"metadata": map[string]any{
							"name":      "my-db",
							"namespace": "databases",
						},
						"status": map[string]any{
							"atProvider": map[string]any{
								"endpoint": map[string]any{
									"address": "my-db.abc123.us-west-2.rds",
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-db.databases.com", endpoint.RecordTypeCNAME, "my-db.abc123.us-west-2.rds").
					WithLabel(endpoint.ResourceLabelKey, "rdsinstance/databases/my-db"),
				endpoint.NewEndpoint("my-vm.vms.com", endpoint.RecordTypeA, "10.244.1.100").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/vms/my-vm"),
			},
		},
		{
			title: "two different resource types with same template",
			cfg: cfg{
				resources: []string{
					"virtualmachineinstances.v1.kubevirt.io",
					"targetgroupbindings.v1beta1.elbv2.k8s.aws",
				},
				fqdnTemplate:       "{{.Name}}.{{.Kind}}.example.com",
				targetFqdnTemplate: "{{.Status.target}}",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "web-server",
							"namespace": "default",
						},
						"status": map[string]any{
							"target": "192.168.1.10",
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "elbv2.k8s.aws/v1beta1",
						"kind":       "TargetGroupBinding",
						"metadata": map[string]any{
							"name":      "api-tgb",
							"namespace": "default",
						},
						"status": map[string]any{
							"target": "lb.api.example.com",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api-tgb.TargetGroupBinding.example.com", endpoint.RecordTypeCNAME, "lb.api.example.com").
					WithLabel(endpoint.ResourceLabelKey, "targetgroupbinding/default/api-tgb"),
				endpoint.NewEndpoint("web-server.VirtualMachineInstance.example.com", endpoint.RecordTypeA, "192.168.1.10").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/web-server"),
			},
		},
		{
			title: "combined fqdn annotation and template",
			cfg: cfg{
				resources:          []string{"virtualmachineinstances.v1.kubevirt.io"},
				fqdnTemplate:       "{{.Name}}.template.example.com",
				targetFqdnTemplate: `{{index .Status.interfaces 0 "ipAddress"}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "default",
							"annotations": map[string]any{
								annotations.HostnameKey: "my-vm.annotation.example.com",
								annotations.TargetKey:   "192.168.1.100",
							},
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{
									"ipAddress": "10.244.1.50",
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vm.annotation.example.com", endpoint.RecordTypeA, "192.168.1.100").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/my-vm"),
				endpoint.NewEndpoint("my-vm.template.example.com", endpoint.RecordTypeA, "10.244.1.50").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/my-vm"),
			},
		},
		{
			title: "three different resource types",
			cfg: cfg{
				resources: []string{
					"virtualmachineinstances.v1.kubevirt.io",
					"targetgroupbinding.v1beta1.elbv2.k8s.aws",
					"apisixroute.v2.apisix.apache.org",
				},
				fqdnTemplate: `
{{if eq .Kind "VirtualMachineInstance"}}{{.Name}}.vm.com{{end}},
{{if eq .Kind "TargetGroupBinding"}}{{.Name}}.tgb.com{{end}},
{{if eq .Kind "ApisixRoute"}}{{.Name}}.route.com{{end}}`,
				targetFqdnTemplate: `
{{if eq .Kind "VirtualMachineInstance"}}{{index .Status.interfaces 0 "ipAddress"}}{{end}},
{{if eq .Kind "TargetGroupBinding"}}{{.Status.loadBalancerHostname}}{{end}},
{{if eq .Kind "ApisixRoute"}}{{.Status.apisix.gateway}}{{end}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "vms",
						},
						"status": map[string]any{
							"interfaces": []any{
								map[string]any{
									"ipAddress": "10.0.0.1",
									"name":      "default",
								},
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "elbv2.k8s.aws/v1beta1",
						"kind":       "TargetGroupBinding",
						"metadata": map[string]any{
							"name":      "my-tgb",
							"namespace": "apps",
						},
						"status": map[string]any{
							"loadBalancerHostname": "my-alb.us-east-1.elb.amazonaws.com",
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "apisix.apache.org/v2",
						"kind":       "ApisixRoute",
						"metadata": map[string]any{
							"name":      "httpbin",
							"namespace": "ingress-apisix",
						},
						"spec": map[string]any{
							"ingressClassName": "apisix",
							"http": []any{
								map[string]any{
									"name": "httpbin",
									"match": map[string]any{
										"paths": []any{"/ip"},
									},
									"backends": []any{
										map[string]any{
											"serviceName": "httpbin",
											"servicePort": int64(80),
										},
									},
								},
							},
						},
						"status": map[string]any{
							"apisix": map[string]any{
								"gateway": "apisix-gateway.ingress-apisix.svc.cluster.local",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("httpbin.route.com", endpoint.RecordTypeCNAME, "apisix-gateway.ingress-apisix.svc.cluster.local").
					WithLabel(endpoint.ResourceLabelKey, "apisixroute/ingress-apisix/httpbin"),
				endpoint.NewEndpoint("my-tgb.tgb.com", endpoint.RecordTypeCNAME, "my-alb.us-east-1.elb.amazonaws.com").
					WithLabel(endpoint.ResourceLabelKey, "targetgroupbinding/apps/my-tgb"),
				endpoint.NewEndpoint("my-vm.vm.com", endpoint.RecordTypeA, "10.0.0.1").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/vms/my-vm"),
			},
		},
		{
			title: "ACK S3 Bucket with FieldExport to ConfigMap",
			cfg: cfg{
				resources: []string{
					"buckets.v1alpha1.s3.services.k8s.aws",
					"fieldexports.v1alpha1.services.k8s.aws",
					"configmap.v1"},
				fqdnTemplate: `{{if eq .Kind "ConfigMap"}}{{.Name}}.s3.example.com{{end}}`,
				targetFqdnTemplate: `
{{if eq .Kind "ConfigMap"}}{{$url := index .Object.data "default.export-user-data-bucket"}}{{trimSuffix (trimPrefix $url "https://") "/"}}{{end}}`,
				labelFilter: "app.kubernetes.io/name=example-app",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "s3.services.k8s.aws/v1alpha1",
						"kind":       "Bucket",
						"metadata": map[string]any{
							"name":      "application-user-data",
							"namespace": "default",
							"labels": map[string]any{
								"app.kubernetes.io/name":    "example-app",
								"app.kubernetes.io/part-of": "exported-config",
							},
							"annotations": map[string]any{
								annotations.ControllerKey: annotations.ControllerValue,
							},
						},
						"spec": map[string]any{
							"name": "doc-example-bucket",
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "services.k8s.aws/v1alpha1",
						"kind":       "FieldExport",
						"metadata": map[string]any{
							"name":      "export-user-data-bucket",
							"namespace": "default",
							"labels": map[string]any{
								"app.kubernetes.io/name":    "example-app",
								"app.kubernetes.io/part-of": "exported-config",
							},
							"annotations": map[string]any{
								annotations.ControllerKey: annotations.ControllerValue,
							},
						},
						"spec": map[string]any{
							"to": map[string]any{
								"name":      "application-user-data-cm",
								"namespace": "default",
								"kind":      "configmap",
							},
							"from": map[string]any{
								"path": ".status.location",
								"resource": map[string]any{
									"group":     "s3.services.k8s.aws",
									"kind":      "Bucket",
									"name":      "application-user-data",
									"namespace": "default",
								},
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "v1",
						"kind":       "ConfigMap",
						"metadata": map[string]any{
							"name":      "application-user-data-cm",
							"namespace": "default",
							"labels": map[string]any{
								"app.kubernetes.io/name":    "example-app",
								"app.kubernetes.io/part-of": "exported-config",
							},
							"annotations": map[string]any{
								annotations.ControllerKey: annotations.ControllerValue,
							},
						},
						"data": map[string]any{
							"default.export-user-data-bucket": "https://doc-example-bucket.s3.amazonaws.com/",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("application-user-data-cm.s3.example.com", endpoint.RecordTypeCNAME, "doc-example-bucket.s3.amazonaws.com").
					WithLabel(endpoint.ResourceLabelKey, "configmap/default/application-user-data-cm"),
			},
		},
		{
			title: "EndpointSlice for headless service with per-pod DNS",
			cfg: cfg{
				resources:          []string{"endpointslices.v1.discovery.k8s.io"},
				hostTargetTemplate: `{{if and (eq .Kind "EndpointSlice") (hasKey .Labels "service.kubernetes.io/headless")}}{{range $ep := .Object.endpoints}}{{if $ep.conditions.ready}}{{range $ep.addresses}}{{$ep.targetRef.name}}.pod.com:{{.}},{{end}}{{end}}{{end}}{{end}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "discovery.k8s.io/v1",
						"kind":       "EndpointSlice",
						"metadata": map[string]any{
							"name":      "test-headless-abc12",
							"namespace": "default",
							"labels": map[string]any{
								"endpointslice.kubernetes.io/managed-by": "endpointslice-controller.k8s.io",
								"kubernetes.io/service-name":             "test-headless",
								"service.kubernetes.io/headless":         "",
							},
						},
						"addressType": "IPv4",
						"endpoints": []any{
							map[string]any{
								"addresses": []any{"10.244.1.2", "2001:db8::1"},
								"conditions": map[string]any{
									"ready": true,
								},
								"nodeName": "worker1",
								"targetRef": map[string]any{
									"kind":      "Pod",
									"name":      "app-abc12",
									"namespace": "default",
								},
							},
							map[string]any{
								"addresses": []any{"10.244.2.3", "10.244.2.4"},
								"conditions": map[string]any{
									"ready": true,
								},
								"nodeName": "worker2",
								"targetRef": map[string]any{
									"kind":      "Pod",
									"name":      "app-def34",
									"namespace": "default",
								},
							},
						},
						"ports": []any{
							map[string]any{
								"name":     "http",
								"port":     int64(80),
								"protocol": "TCP",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("app-abc12.pod.com", endpoint.RecordTypeA, "10.244.1.2").
					WithLabel(endpoint.ResourceLabelKey, "endpointslice/default/test-headless-abc12"),
				endpoint.NewEndpoint("app-abc12.pod.com", endpoint.RecordTypeAAAA, "2001:db8::1").
					WithLabel(endpoint.ResourceLabelKey, "endpointslice/default/test-headless-abc12"),
				endpoint.NewEndpoint("app-def34.pod.com", endpoint.RecordTypeA, "10.244.2.3", "10.244.2.4").
					WithLabel(endpoint.ResourceLabelKey, "endpointslice/default/test-headless-abc12"),
			},
		},
		{
			title: "hostTargetTemplate returns no values when condition not met",
			cfg: cfg{
				resources:          []string{"endpointslices.v1.discovery.k8s.io"},
				hostTargetTemplate: `{{if and (eq .Kind "EndpointSlice") (hasKey .Labels "service.kubernetes.io/headless")}}{{range $ep := .Object.endpoints}}{{if $ep.conditions.ready}}{{range $ep.addresses}}{{$ep.targetRef.name}}.pod.com:{{.}},{{end}}{{end}}{{end}}{{end}}`,
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "discovery.k8s.io/v1",
						"kind":       "EndpointSlice",
						"metadata": map[string]any{
							"name":      "regular-service-abc12",
							"namespace": "default",
							"labels": map[string]any{
								"endpointslice.kubernetes.io/managed-by": "endpointslice-controller.k8s.io",
								"kubernetes.io/service-name":             "regular-service",
								// Note: missing service.kubernetes.io/headless label
							},
						},
						"addressType": "IPv4",
						"endpoints": []any{
							map[string]any{
								"addresses": []any{"10.244.1.2"},
								"conditions": map[string]any{
									"ready": true,
								},
								"targetRef": map[string]any{
									"kind":      "Pod",
									"name":      "app-abc12",
									"namespace": "default",
								},
							},
						},
					},
				},
			},
			expected: nil,
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient, dynamicClient := setupUnstructuredTestClients(t, tt.cfg.resources, tt.objects)

			var selector labels.Selector
			if tt.cfg.labelFilter != "" {
				var err error
				selector, err = labels.Parse(tt.cfg.labelFilter)
				require.NoError(t, err)
			} else {
				selector = labels.Everything()
			}

			src, err := NewUnstructuredFQDNSource(
				t.Context(),
				dynamicClient,
				kubeClient,
				"",
				"",
				selector,
				tt.cfg.resources,
				tt.cfg.fqdnTemplate,
				tt.cfg.targetFqdnTemplate,
				tt.cfg.hostTargetTemplate,
				true,
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestUnstructuredWrapper_Templating(t *testing.T) {
	tests := []struct {
		name    string
		tmpl    string
		obj     *unstructured.Unstructured
		want    []string
		wantErr bool
	}{
		{
			name: "typed-style Name and Namespace access",
			tmpl: "{{.Name}}.{{.Namespace}}.example.com",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "kubevirt.io/v1",
					"kind":       "VirtualMachineInstance",
					"metadata": map[string]any{
						"name":      "my-vm",
						"namespace": "default",
					},
				},
			},
			want: []string{"my-vm.default.example.com"},
		},
		{
			name: "raw Metadata map access",
			tmpl: "{{.Metadata.name}}.{{.Metadata.namespace}}.example.com",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "kubevirt.io/v1",
					"kind":       "VirtualMachineInstance",
					"metadata": map[string]any{
						"name":      "my-vm",
						"namespace": "default",
					},
				},
			},
			want: []string{"my-vm.default.example.com"},
		},
		{
			name: "nested Status field access",
			tmpl: "{{.Status.atProvider.endpoint.address}}",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "rds.aws.crossplane.io/v1alpha1",
					"kind":       "RDSInstance",
					"metadata": map[string]any{
						"name":      "prod-db",
						"namespace": "default",
					},
					"status": map[string]any{
						"atProvider": map[string]any{
							"endpoint": map[string]any{
								"address": "prod-db.abc123.rds.amazonaws.com",
							},
						},
					},
				},
			},
			want: []string{"prod-db.abc123.rds.amazonaws.com"},
		},
		{
			name: "array index access via Status",
			tmpl: `{{index .Status.interfaces 0 "ipAddress"}}`,
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "kubevirt.io/v1",
					"kind":       "VirtualMachineInstance",
					"metadata": map[string]any{
						"name":      "vm-1",
						"namespace": "default",
					},
					"status": map[string]any{
						"interfaces": []any{
							map[string]any{
								"ipAddress": "10.244.1.50",
								"name":      "default",
							},
						},
					},
				},
			},
			want: []string{"10.244.1.50"},
		},
		{
			name: "typed-style Labels access",
			tmpl: `{{index .Labels "app.kubernetes.io/instance"}}.example.com`,
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "argoproj.io/v1alpha1",
					"kind":       "Application",
					"metadata": map[string]any{
						"name":      "guestbook",
						"namespace": "argocd",
						"labels": map[string]any{
							"app.kubernetes.io/instance": "guestbook-prod",
						},
					},
				},
			},
			want: []string{"guestbook-prod.example.com"},
		},
		{
			name: "Kind and APIVersion access",
			tmpl: "{{.Kind}}.{{.APIVersion}}.example.com",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "kubevirt.io/v1",
					"kind":       "VirtualMachineInstance",
					"metadata": map[string]any{
						"name":      "test",
						"namespace": "default",
					},
				},
			},
			want: []string{"VirtualMachineInstance.kubevirt.io/v1.example.com"},
		},
		{
			name: "Spec hosts array",
			tmpl: `{{index .Spec.hosts 0}}`,
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "networking.istio.io/v1beta1",
					"kind":       "VirtualService",
					"metadata": map[string]any{
						"name":      "reviews",
						"namespace": "bookinfo",
					},
					"spec": map[string]any{
						"hosts": []any{
							"reviews.bookinfo.svc.cluster.local",
						},
					},
				},
			},
			want: []string{"reviews.bookinfo.svc.cluster.local"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := fqdn.ParseTemplate(tt.tmpl)
			require.NoError(t, err)

			wrapped := newUnstructuredWrapper(tt.obj)
			got, err := fqdn.ExecTemplate(tmpl, wrapped)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// buildAPIResourceLists build API resource list from test resources
func buildAPIResourceLists(t *testing.T, resources []string, objects []*unstructured.Unstructured) []*metav1.APIResourceList {
	t.Helper()

	// Build a map of GVR to Kind from objects
	gvrToKind := make(map[string]string)
	for _, obj := range objects {
		apiVersion := obj.GetAPIVersion()
		kind := obj.GetKind()
		gvrToKind[apiVersion] = kind
	}

	result := make([]*metav1.APIResourceList, 0, len(resources))
	for _, res := range resources {
		if strings.Count(res, ".") == 1 {
			res += "."
		}
		gvr, _ := schema.ParseResourceArg(res)
		require.NotNil(t, gvr, "invalid resource identifier: %s", res)

		// Find matching kind from objects
		kind := gvrToKind[gvr.GroupVersion().String()]
		if kind == "" {
			// Fallback to first object's kind
			kind = objects[0].GetKind()
		}

		result = append(result, &metav1.APIResourceList{
			GroupVersion: gvr.GroupVersion().String(),
			APIResources: []metav1.APIResource{{
				Name:       gvr.Resource,
				Namespaced: true,
				Kind:       kind,
			}},
		})
	}
	return result
}

// setupUnstructuredTestClients creates fake kube and dynamic clients with the given resources and objects.
func setupUnstructuredTestClients(t *testing.T, resources []string, objects []*unstructured.Unstructured) (
	kubernetes.Interface, dynamic.Interface,
) {
	t.Helper()

	kubeClient := fake.NewClientset()
	fakeDiscovery := kubeClient.Discovery().(*discoveryfake.FakeDiscovery)
	fakeDiscovery.Resources = buildAPIResourceLists(t, resources, objects)

	// Build GVR to ListKind map and apiVersion to GVR map
	gvrToListKind := make(map[schema.GroupVersionResource]string)
	apiVersionToGVR := make(map[string]schema.GroupVersionResource)
	for _, res := range resources {
		if strings.Count(res, ".") == 1 {
			res += "."
		}
		gvr, _ := schema.ParseResourceArg(res)

		require.NotNil(t, gvr, "invalid resource identifier: %s", res)
		apiVersionToGVR[gvr.GroupVersion().String()] = *gvr
	}

	// Determine list kinds from objects
	for _, obj := range objects {
		apiVersion := obj.GetAPIVersion()
		if gvr, ok := apiVersionToGVR[apiVersion]; ok {
			gvrToListKind[gvr] = obj.GetKind() + "List"
		}
	}

	scheme := runtime.NewScheme()
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, gvrToListKind)

	// Create each object using its matching GVR
	for _, obj := range objects {
		apiVersion := obj.GetAPIVersion()
		gvr, ok := apiVersionToGVR[apiVersion]
		require.True(t, ok, "no resource found for apiVersion %s", apiVersion)
		_, err := dynamicClient.Resource(gvr).Namespace(obj.GetNamespace()).Create(
			t.Context(), obj, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	return kubeClient, dynamicClient
}
