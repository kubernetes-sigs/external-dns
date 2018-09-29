package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/linki/instrumented_http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/dns/v1"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe("127.0.0.1:9099", nil))
	}()

	defaultClient, err := google.DefaultClient(context.TODO(), dns.NdevClouddnsReadwriteScope)
	if err != nil {
		log.Fatal(err)
	}

	instrumentedClient := instrumented_http.NewClient(defaultClient, &instrumented_http.Callbacks{
		PathProcessor: instrumented_http.LastPathElementProcessor,
	})

	dnsClient, err := dns.New(instrumentedClient)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zones, err := dnsClient.ManagedZones.List("<replace-with-your-google-project>").Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, z := range zones.ManagedZones {
			fmt.Println(z.Name)
		}

		time.Sleep(10 * time.Second)
	}
}

// expected result:
//
// $ curl -Ss 127.0.0.1:9099/metrics | grep http
//
// http_request_duration_seconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.5"} 0.642468
// http_request_duration_seconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.9"} 0.660945
// http_request_duration_seconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.99"} 0.660945
// http_request_duration_seconds_sum{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200"} 1.303413521
// http_request_duration_seconds_count{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200"} 2
