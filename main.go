/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kubernetes-incubator/external-dns/controller"
	"github.com/kubernetes-incubator/external-dns/dnsprovider"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns/validation"
	"github.com/kubernetes-incubator/external-dns/source"
)

var (
	version = "unknown"
)

func main() {
	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	if cfg.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	if err := validation.ValidateConfig(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if cfg.DryRun {
		log.Info("running in dry-run mode. No changes to DNS records will be made.")
	}
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	stopChan := make(chan struct{}, 1)

	go registerHandlers(cfg.HealthPort)
	go handleSigterm(stopChan)

	client, err := newClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	source.Register("service", source.NewServiceSource(client, cfg.Namespace))
	source.Register("ingress", source.NewIngressSource(client, cfg.Namespace))

	sources := source.NewMultiSource(source.LookupMultiple(cfg.Sources...)...)

	googleProvider, err := dnsprovider.NewGoogleProvider(cfg.GoogleProject, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}

	awsProvider, err := dnsprovider.NewAWSProvider(cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}

	dnsprovider.Register("google", googleProvider)
	dnsprovider.Register("aws", awsProvider)

	ctrl := controller.Controller{
		Zone:        cfg.Zone,
		Source:      sources,
		DNSProvider: dnsprovider.Lookup(cfg.DNSProvider),
	}

	if cfg.Once {
		err := ctrl.RunOnce()
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	ctrl.Run(stopChan)
	for {
		log.Infoln("pod waiting to be deleted")
		time.Sleep(time.Second * 30)
	}
}

func registerHandlers(port string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handleSigterm(stopChan chan struct{}) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Infoln("received SIGTERM. Terminating...")
	close(stopChan)
}

func newClient(cfg *externaldns.Config) (*kubernetes.Clientset, error) {
	if !cfg.InCluster && cfg.KubeConfig == "" {
		cfg.KubeConfig = clientcmd.RecommendedHomeFile
	}

	config, err := clientcmd.BuildConfigFromFlags("", cfg.KubeConfig)
	if err != nil {
		return nil, err
	}

	log.Infof("targeting cluster at %s", config.Host)

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
