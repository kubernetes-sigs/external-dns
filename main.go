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

	"github.com/kubernetes-incubator/external-dns/controller"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns/validation"
)

func main() {
	cfg := externaldns.NewConfig()
	cfg.ParseFlags()
	if err := validation.ValidateConfig(cfg); err != nil {
		log.Errorf("config validation failed: %v", err)
	}

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	stopChan := make(chan struct{}, 1)

	go registerHandlers(cfg.HealthPort)
	go handleSigterm(stopChan)

	controller.Run(stopChan)
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
