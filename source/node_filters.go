/*
Copyright 2026 The Kubernetes Authors.

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
	"strconv"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/client-go/tools/cache"
)

// Taints autoscalers add before they drain and delete a Node.
const (
	toBeDeletedByClusterAutoscalerTaint = "ToBeDeletedByClusterAutoscaler"
	karpenterDisruptedTaint             = "karpenter.sh/disrupted"
)

// unschedulableReason returns why --exclude-unschedulable skips the Node, or "" if it doesn't.
func unschedulableReason(node *v1.Node) string {
	if node.Spec.Unschedulable {
		return "unschedulable"
	}
	if node.DeletionTimestamp != nil {
		return "being deleted"
	}
	for _, taint := range node.Spec.Taints {
		if taint.Key == toBeDeletedByClusterAutoscalerTaint || taint.Key == karpenterDisruptedTaint {
			return "tainted " + taint.Key
		}
	}
	return ""
}

// isExcludedFromLoadBalancers reports whether the Node has the exclude-from-external-load-balancers
// label with a value that does not parse as false, like the cloud-provider service controller.
func isExcludedFromLoadBalancers(node *v1.Node) bool {
	value, ok := node.Labels[v1.LabelNodeExcludeBalancers]
	if !ok {
		return false
	}
	exclude, err := strconv.ParseBool(value)
	return err != nil || exclude
}

// nodePortEventHandler calls handler on Node add and delete, and on updates that change NodePort targets.
func nodePortEventHandler(handler func()) cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(_ any) { handler() },
		UpdateFunc: func(oldObj, newObj any) {
			if nodePortTargetsChanged(oldObj, newObj) {
				handler()
			}
		},
		DeleteFunc: func(_ any) { handler() },
	}
}

// nodePortTargetsChanged reports whether an update changed the Node's addresses, labels, taints, unschedulable flag or deletion state.
func nodePortTargetsChanged(oldObj, newObj any) bool {
	oldNode, okOld := oldObj.(*v1.Node)
	newNode, okNew := newObj.(*v1.Node)
	if !okOld || !okNew {
		return true
	}
	return !equality.Semantic.DeepEqual(oldNode.Status.Addresses, newNode.Status.Addresses) ||
		!equality.Semantic.DeepEqual(oldNode.Labels, newNode.Labels) ||
		!equality.Semantic.DeepEqual(oldNode.Spec.Taints, newNode.Spec.Taints) ||
		oldNode.Spec.Unschedulable != newNode.Spec.Unschedulable ||
		(oldNode.DeletionTimestamp == nil) != (newNode.DeletionTimestamp == nil)
}
