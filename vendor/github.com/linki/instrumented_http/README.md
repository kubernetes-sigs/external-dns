# instrumented_http

A Go `http.RoundTripper` that exports request statistics via [Prometheus](https://prometheus.io/).

# Example

Transparently inject `instrumented_http` into any `http.Client` or `http.RoundTripper` and get metrics about all requests made.

```console
$ curl -Ss 127.0.0.1:9099/metrics | grep http
http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.5"} 0.83626
http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.9"} 0.736648
http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.99"} 0.736648
http_request_duration_seconds_sum{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200"} 0.820274243
http_request_duration_seconds_count{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200"} 2
```

# Usage

Browse the [examples](examples) directory to see how `instrumented_http` works with:
* [http.DefaultClient](https://golang.org/pkg/net/http/#Client): [examples/default-client](examples/default-client)
* a custom [http.Transport](https://golang.org/pkg/net/http/#Transport): [examples/custom-transport](examples/custom-transport)
* the [Google CloudDNS](https://godoc.org/google.golang.org/api/dns/v1) client: [examples/googledns](examples/googledns)
* the [AWS Route53](https://godoc.org/github.com/aws/aws-sdk-go/service/route53) client: [examples/route53](examples/route53)
* the [Kubernetes](https://godoc.org/k8s.io/client-go) client: [examples/kubernetes](examples/kubernetes)
* [Resty](https://github.com/go-resty/resty): [examples/resty](examples/resty)
* [Sling](https://github.com/dghubble/sling): [examples/sling](examples/sling)
* [Gentleman](https://github.com/h2non/gentleman): [examples/gentleman](examples/gentleman)
