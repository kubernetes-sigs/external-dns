package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
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
		PathProcessor: func(path string) string {
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
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
// http_request_duration_microseconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.5"} 642468
// http_request_duration_microseconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.9"} 660945
// http_request_duration_microseconds{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200",quantile="0.99"} 660945
// http_request_duration_microseconds_sum{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200"} 1.303413e+06
// http_request_duration_microseconds_count{handler="instrumented_http",host="www.googleapis.com",method="GET",path="managedZones",query="",scheme="https",status="200"} 2
