/*
Copyright 2025 The Kubernetes Authors.

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

package types

type Type = string

const (
	Node                Type = "node"
	Service             Type = "service"
	Ingress             Type = "ingress"
	Pod                 Type = "pod"
	GatewayHttpRoute    Type = "gateway-httproute"
	GatewayGrpcRoute    Type = "gateway-grpcroute"
	GatewayTlsRoute     Type = "gateway-tlsroute"
	GatewayTcpRoute     Type = "gateway-tcproute"
	GatewayUdpRoute     Type = "gateway-udproute"
	IstioGateway        Type = "istio-gateway"
	IstioVirtualService Type = "istio-virtualservice"
	Cloudfoundry        Type = "cloudfoundry"
	AmbassadorHost      Type = "ambassador-host"
	ContourHTTPProxy    Type = "contour-httpproxy"
	GlooProxy           Type = "gloo-proxy"
	TraefikProxy        Type = "traefik-proxy"
	OpenShiftRoute      Type = "openshift-route"
	Fake                Type = "fake"
	Connector           Type = "connector"
	CRD                 Type = "crd"
	SkipperRouteGroup   Type = "skipper-routegroup"
	KongTCPIngress      Type = "kong-tcpingress"
	F5VirtualServer     Type = "f5-virtualserver"
	F5TransportServer   Type = "f5-transportserver"
)
