package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/linki/instrumented_http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe("127.0.0.1:9099", nil))
	}()

	config := aws.NewConfig()

	config = config.WithHTTPClient(
		instrumented_http.NewClient(config.HTTPClient, &instrumented_http.Callbacks{
			PathProcessor: instrumented_http.LastPathElementProcessor,
		}),
	)

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatal(err)
	}

	client := route53.New(session)

	for {
		zones, err := client.ListHostedZones(&route53.ListHostedZonesInput{})
		if err != nil {
			log.Fatal(err)
		}

		for _, z := range zones.HostedZones {
			fmt.Println(aws.StringValue(z.Name))
		}

		time.Sleep(10 * time.Second)
	}
}

// expected result:
//
// $ curl -Ss 127.0.0.1:9099/metrics | grep http
//
// http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.5"} 0.463922
// http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.9"} 0.969598
// http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.99"} 0.969598
// http_request_duration_seconds_sum{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200"} 2.363297242
// http_request_duration_seconds_count{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200"} 4
