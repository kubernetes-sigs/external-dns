# About

A source in ExternalDNS defines where DNS records are discovered from within your infrastructure. Each source corresponds to a specific Kubernetes resource or external system that declares DNS names.

ExternalDNS watches the specified sources for hostname information and uses it to create, update, or delete DNS records accordingly. Multiple sources can be configured simultaneously to support diverse environments.

| Source                                  | Resources                                                                     | annotation-filter | label-filter |
| --------------------------------------- | ----------------------------------------------------------------------------- | ----------------- | ------------ |
| ambassador-host                         | Host.getambassador.io                                                         | Yes               | Yes          |
| connector                               |                                                                               |                   |              |
| contour-httpproxy                       | HttpProxy.projectcontour.io                                                   | Yes               |              |
| cloudfoundry                            |                                                                               |                   |              |
| [crd](crd.md)                           | DNSEndpoint.externaldns.k8s.io                                                | Yes               | Yes          |
| [f5-virtualserver](f5-virtualserver.md) | VirtualServer.cis.f5.com                                                      | Yes               |              |
| [gateway-grpcroute](gateway.md)         | GRPCRoute.gateway.networking.k8s.io                                           | Yes               | Yes          |
| [gateway-httproute](gateway.md)         | HTTPRoute.gateway.networking.k8s.io                                           | Yes               | Yes          |
| [gateway-tcproute](gateway.md)          | TCPRoute.gateway.networking.k8s.io                                            | Yes               | Yes          |
| [gateway-tlsroute](gateway.md)          | TLSRoute.gateway.networking.k8s.io                                            | Yes               | Yes          |
| [gateway-udproute](gateway.md)          | UDPRoute.gateway.networking.k8s.io                                            | Yes               | Yes          |
| [gloo-proxy](gloo-proxy.md)             | Proxy.gloo.solo.io                                                            |                   |              |
| [ingress](ingress.md)                   | Ingress.networking.k8s.io                                                     | Yes               | Yes          |
| [istio-gateway](istio.md)               | Gateway.networking.istio.io                                                   | Yes               |              |
| [istio-virtualservice](istio.md)        | VirtualService.networking.istio.io                                            | Yes               |              |
| [kong-tcpingress](kong.md)              | TCPIngress.configuration.konghq.com                                           | Yes               |              |
| [node](nodes.md)                        | Node                                                                          | Yes               | Yes          |
| [openshift-route](openshift.md)         | Route.route.openshift.io                                                      | Yes               | Yes          |
| [pod](pod.md)                           | Pod                                                                           |                   |              |
| [service](service.md)                   | Service                                                                       | Yes               | Yes          |
| skipper-routegroup                      | RouteGroup.zalando.org                                                        | Yes               |              |
| [traefik-proxy](traefik-proxy.md)       | IngressRoute.traefik.io IngressRouteTCP.traefik.io IngressRouteUDP.traefik.io | Yes               |              |
