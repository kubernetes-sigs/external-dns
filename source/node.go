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

package source

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

type nodeSource struct {
	client       kubernetes.Interface
	fqdnTemplate *template.Template
	nodeInformer coreinformers.NodeInformer
}

func NewNodeSource(kubeClient kubernetes.Interface, fqdnTemplate string) (Source, error) {
	var (
		tmpl *template.Template
		err  error
	)

	if fqdnTemplate != "" {
		tmpl, err = template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdnTemplate)
		if err != nil {
			return nil, err
		}
	}

	// Use shared informers to listen for add/update/delete of services/pods/nodes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0)
	nodeInformer := informerFactory.Core().V1().Nodes()

	// Add default resource event handler to properly initialize informer.
	nodeInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("node added")
			},
		},
	)

	// TODO informer is not explicitly stopped since controller is not passing in its channel.
	informerFactory.Start(wait.NeverStop)

	// wait for the local cache to be populated.
	err = wait.Poll(time.Second, 60*time.Second, func() (bool, error) {
		return nodeInformer.Informer().HasSynced() == true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sync cache: %v", err)
	}

	return &nodeSource{
		client:       kubeClient,
		fqdnTemplate: tmpl,
		nodeInformer: nodeInformer,
	}, nil
}

// Endpoints returns endpoint objects for each service that should be processed.
func (ns *nodeSource) Endpoints() ([]*endpoint.Endpoint, error) {
	nodes, err := ns.nodeInformer.Lister().List(labels.Everything())
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	// create endpoints for all nodes
	for _, node := range nodes {
		log.Debugf("creating endpoint for node %s", node.Name)

		// create new endpoint with the information we already have
		ep := &endpoint.Endpoint{
			RecordType: "A", // hardcoded DNS record type
		}

		if ns.fqdnTemplate != nil {
			// Process the whole template string
			var buf bytes.Buffer
			err := ns.fqdnTemplate.Execute(&buf, node)
			if err != nil {
				return nil, fmt.Errorf("failed to apply template on node %s: %v", node.Name, err)
			}

			ep.DNSName = buf.String()
			log.Debugf("applied template for %s, converting to %s", node.Name, ep.DNSName)
		} else {
			ep.DNSName = node.Name
			log.Debugf("not applying template for %s", node.Name)
		}

		addr, err := ns.nodeAddress(node)
		if err != nil {
			log.Error(err)
			continue
		}

		ep.Targets = endpoint.Targets([]string{addr})

		log.Debugf("adding endpoint %s", ep)
		endpoints = append(endpoints, ep)
	}

	return endpoints, nil
}

// nodeAddress returns node's externalIP and if that's not found, node's internalIP
// basically what k8s.io/kubernetes/pkg/util/node.GetPreferredNodeAddress does
func (ns *nodeSource) nodeAddress(node *v1.Node) (string, error) {
	for _, t := range []v1.NodeAddressType{v1.NodeExternalIP, v1.NodeInternalIP} {
		for _, addr := range node.Status.Addresses {
			if addr.Type == t {
				return addr.Address, nil
			}
		}
	}

	return "", fmt.Errorf("could not find node address for %s", node.Name)
}
