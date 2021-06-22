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

package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
)

// Returns string version of ResourceName.
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (rn ResourceName) String() string {
	return string(rn)
}

// Cpu returns the Cpu limit if specified.
func (rl *ResourceList) Cpu() *resource.Quantity {
	return rl.Name(ResourceCPU, resource.DecimalSI)
}

// Memory returns the Memory limit if specified.
func (rl *ResourceList) Memory() *resource.Quantity {
	return rl.Name(ResourceMemory, resource.BinarySI)
}

// Storage returns the Storage limit if specified.
func (rl *ResourceList) Storage() *resource.Quantity {
	return rl.Name(ResourceStorage, resource.BinarySI)
}

// Pods returns the list of pods
func (rl *ResourceList) Pods() *resource.Quantity {
	return rl.Name(ResourcePods, resource.DecimalSI)
}

// StorageEphemeral returns the list of ephemeral storage volumes, if any
func (rl *ResourceList) StorageEphemeral() *resource.Quantity {
	return rl.Name(ResourceEphemeralStorage, resource.BinarySI)
}

// Name returns the resource with name if specified, otherwise it returns a nil quantity with default format.
func (rl *ResourceList) Name(name ResourceName, defaultFormat resource.Format) *resource.Quantity {
	if val, ok := (*rl)[name]; ok {
		return &val
	}
	return &resource.Quantity{Format: defaultFormat}
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
func (self ResourceName) String() string {
	return string(self)
||||||| parent of 5ce8c7613 (update vendored files)
func (self ResourceName) String() string {
	return string(self)
=======
func (rn ResourceName) String() string {
	return string(rn)
>>>>>>> 5ce8c7613 (update vendored files)
}

// Cpu returns the Cpu limit if specified.
func (rl *ResourceList) Cpu() *resource.Quantity {
	return rl.Name(ResourceCPU, resource.DecimalSI)
}

// Memory returns the Memory limit if specified.
func (rl *ResourceList) Memory() *resource.Quantity {
	return rl.Name(ResourceMemory, resource.BinarySI)
}

// Storage returns the Storage limit if specified.
func (rl *ResourceList) Storage() *resource.Quantity {
	return rl.Name(ResourceStorage, resource.BinarySI)
}

// Pods returns the list of pods
func (rl *ResourceList) Pods() *resource.Quantity {
	return rl.Name(ResourcePods, resource.DecimalSI)
}

// StorageEphemeral returns the list of ephemeral storage volumes, if any
func (rl *ResourceList) StorageEphemeral() *resource.Quantity {
	return rl.Name(ResourceEphemeralStorage, resource.BinarySI)
}

// Name returns the resource with name if specified, otherwise it returns a nil quantity with default format.
func (rl *ResourceList) Name(name ResourceName, defaultFormat resource.Format) *resource.Quantity {
	if val, ok := (*rl)[name]; ok {
		return &val
	}
<<<<<<< HEAD
	return &resource.Quantity{}
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return &resource.Quantity{}
=======
	return &resource.Quantity{Format: defaultFormat}
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
func (self ResourceName) String() string {
	return string(self)
}

// Returns the CPU limit if specified.
func (self *ResourceList) Cpu() *resource.Quantity {
	if val, ok := (*self)[ResourceCPU]; ok {
		return &val
	}
	return &resource.Quantity{Format: resource.DecimalSI}
}

// Returns the Memory limit if specified.
func (self *ResourceList) Memory() *resource.Quantity {
	if val, ok := (*self)[ResourceMemory]; ok {
		return &val
	}
	return &resource.Quantity{Format: resource.BinarySI}
}

// Returns the Storage limit if specified.
func (self *ResourceList) Storage() *resource.Quantity {
	if val, ok := (*self)[ResourceStorage]; ok {
		return &val
	}
	return &resource.Quantity{Format: resource.BinarySI}
}

func (self *ResourceList) Pods() *resource.Quantity {
	if val, ok := (*self)[ResourcePods]; ok {
		return &val
	}
	return &resource.Quantity{}
}

func (self *ResourceList) StorageEphemeral() *resource.Quantity {
	if val, ok := (*self)[ResourceEphemeralStorage]; ok {
		return &val
	}
	return &resource.Quantity{}
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
