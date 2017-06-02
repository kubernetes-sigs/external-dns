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

package controller

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

var _ Controller = &LoopedController{}

// LoopedController invokes another controller's Run() method in an endless loop.
type LoopedController struct {
	ctrl   Controller
	config Config
}

// NewLoopedController returns a new endlessly looping Controller that wraps the given controller.
func NewLoopedController(ctrl Controller, config Config) Controller {
	return &LoopedController{ctrl: ctrl, config: config}
}

// Run runs the wrapped controller's Run in a loop with a delay until stopChan receives a value.
func (c *LoopedController) Run() error {
	for {
		if err := c.ctrl.Run(); err != nil {
			log.Error(err)
		}

		select {
		case <-time.After(c.config.Interval):
		case <-c.config.StopChan:
			log.Info("Terminating LoopedController")
			return nil
		}
	}
}
