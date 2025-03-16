/*
Copyright 2025 The Kubernetes Authors.

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

package istio

import (
	"reflect"

	log "github.com/sirupsen/logrus"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
)

const (
	ApiGatewayV1Alpha3        = "gateway.networking.istio.io/v1alpha3"
	ApiVirtualServiceV1Alpha3 = "virtualservice.networking.istio.io/v1alpha3"
)

func GatewayAddFuncEventHander(obj interface{}) {
	if log.IsLevelEnabled(log.DebugLevel) {
		service, ok := obj.(*networkingv1alpha3.Gateway)
		if !ok {
			log.Errorf("event handler not added. want '%s' got '%s'", ApiGatewayV1Alpha3, reflect.TypeOf(obj).String())
			return
		} else {
			log.Debugf("event handler added for '%s' in 'namespace:%s' with 'name:%s'", ApiGatewayV1Alpha3, service.Namespace, service.Name)
		}
	}
}

func VirtualServiceAddedEvent(obj interface{}) {
	if log.IsLevelEnabled(log.DebugLevel) {
		service, ok := obj.(*networkingv1alpha3.VirtualService)
		if !ok {
			log.Errorf("event handler not added. want '%s' got '%s'", ApiVirtualServiceV1Alpha3, reflect.TypeOf(obj).String())
			return
		} else {
			log.Debugf("event handler added for '%s' in 'namespace:%s' with 'name:%s'", ApiVirtualServiceV1Alpha3, service.Namespace, service.Name)
		}
	}
}
