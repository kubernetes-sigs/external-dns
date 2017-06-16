package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/linki/instrumented_http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe("127.0.0.1:9099", nil))
	}()

	var originalTransport http.RoundTripper = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	instrumentedTransport := instrumented_http.NewTransport(originalTransport,
		&instrumented_http.Callbacks{
			PathProcessor:  instrumented_http.IdentityProcessor,
			QueryProcessor: instrumented_http.IdentityProcessor,
		},
	)

	client := &http.Client{Transport: instrumentedTransport}

	for {
		func() {
			resp, err := client.Get("https://kubernetes.io/docs/search/?q=pods")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			fmt.Printf("%d\n", resp.StatusCode)
		}()

		time.Sleep(10 * time.Second)
	}
}

// expected result:
//
// $ curl -Ss 127.0.0.1:9099/metrics | grep http
//
// http_request_duration_seconds{handler="instrumented_http",host="kubernetes.io",method="GET",path="/docs/search/",query="q=pods",scheme="https",status="200",quantile="0.5"} 0.526252
// http_request_duration_seconds{handler="instrumented_http",host="kubernetes.io",method="GET",path="/docs/search/",query="q=pods",scheme="https",status="200",quantile="0.9"} 0.663617
// http_request_duration_seconds{handler="instrumented_http",host="kubernetes.io",method="GET",path="/docs/search/",query="q=pods",scheme="https",status="200",quantile="0.99"} 0.663617
// http_request_duration_seconds_sum{handler="instrumented_http",host="kubernetes.io",method="GET",path="/docs/search/",query="q=pods",scheme="https",status="200"} 1.189869154
// http_request_duration_seconds_count{handler="instrumented_http",host="kubernetes.io",method="GET",path="/docs/search/",query="q=pods",scheme="https",status="200"} 2
