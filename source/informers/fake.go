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

package informers

import (
	"github.com/stretchr/testify/mock"
	corev1lister "k8s.io/client-go/listers/core/v1"
	discoveryv1lister "k8s.io/client-go/listers/discovery/v1"
	"k8s.io/client-go/tools/cache"
)

type FakeServiceInformer struct {
	mock.Mock
}

func (f *FakeServiceInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeServiceInformer) Lister() corev1lister.ServiceLister {
	return corev1lister.NewServiceLister(f.Informer().GetIndexer())
}

type FakeEndpointSliceInformer struct {
	mock.Mock
}

func (f *FakeEndpointSliceInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeEndpointSliceInformer) Lister() discoveryv1lister.EndpointSliceLister {
	return discoveryv1lister.NewEndpointSliceLister(f.Informer().GetIndexer())
}

type FakeNodeInformer struct {
	mock.Mock
}

func (f *FakeNodeInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeNodeInformer) Lister() corev1lister.NodeLister {
	return corev1lister.NewNodeLister(f.Informer().GetIndexer())
}
