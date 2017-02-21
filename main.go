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
)

func main() {
	cfg := newConfig()
	cfg.parseFlags()
	if err := cfg.validate(); err != nil {
		log.Errorf("config validation failed: %v", err)
	}

	if cfg.logFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if cfg.debug {
		log.SetLevel(log.DebugLevel)
	}

	stopChan := make(chan struct{}, 1)

	go registerHandlers(cfg.healthPort)
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
