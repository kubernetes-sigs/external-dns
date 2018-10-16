package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/linki/instrumented_http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe("127.0.0.1:9099", nil))
	}()

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatal(err)
	}

	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return instrumented_http.NewTransport(rt, &instrumented_http.Callbacks{
			PathProcessor: instrumented_http.LastPathElementProcessor,
		})
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	for {
		pods, err := client.CoreV1().Pods(metav1.NamespaceDefault).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range pods.Items {
			fmt.Println(p.Name)
		}

		time.Sleep(10 * time.Second)
	}
}

// expected result:
//
// $ curl -Ss 127.0.0.1:9099/metrics | grep http
//
// http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.5"} 0.83626
// http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.9"} 0.736648
// http_request_duration_seconds{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200",quantile="0.99"} 0.736648
// http_request_duration_seconds_sum{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200"} 0.820274243
// http_request_duration_seconds_count{handler="instrumented_http",host="my-cluster.example.org",method="GET",path="pods",query="",scheme="https",status="200"} 2
