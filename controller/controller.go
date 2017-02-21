package controller

import (
	log "github.com/Sirupsen/logrus"
)

// Controller main controlling object
type Controller struct {
}

// Run runs the main controller loop
func Run(stopChan <-chan struct{}) {
	<-stopChan
	log.Infoln("terminating main controller loop")
}
