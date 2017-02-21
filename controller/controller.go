package controller

import (
	log "github.com/Sirupsen/logrus"
)

// Run runs the main controller loop
func Run(stopChan <-chan struct{}) {
	<-stopChan
	log.Infoln("terminating main controller loop")
}
