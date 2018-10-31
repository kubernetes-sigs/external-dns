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

package srv

import (
	"fmt"
	"strings"
)

const (
	// DummyTarget is used as an SRV target as records are passed around as
	// composite DNS names, where the target is encoded in the DNS name
	DummyTarget = "dummy"
)

// Name holds SRV name specific fields.
type Name struct {
	// Service name e.g. _http.
	Service string
	// Protocol name e.g. _tcp.
	Proto string
	// DNS Domain name.
	Name string
}

// ParseName accepts a name in the for _service._proto.domain and
// unmarshalls into a Name.
func ParseName(name string) Name {
	parts := strings.Split(name, ".")
	return Name{
		Service: parts[0],
		Proto:   parts[1],
		Name:    strings.Join(parts[2:], "."),
	}
}

// Format returns the stringified version of a Name.
func (s Name) Format() string {
	return fmt.Sprintf("%s.%s.%s", s.Service, s.Proto, s.Name)
}

// Target holds SRV target specific fields.
type Target struct {
	// Priority is the priority with which a record will be considered by a client.
	Priority int
	// Weight is the weight of requests as issued by a client.
	Weight int
	// Port is the port number the service can be dialed.
	Port int
	// Target is the DNS record that this target refers to.
	Target string
}

// ParseTarget takes an endpoint target string and unmarshals into an SRV record.
func ParseTarget(target string) Target {
	s := Target{}
	fmt.Sscanf(target, "%d %d %d %s", &s.Priority, &s.Weight, &s.Port, &s.Target)
	return s
}

// Format takes an SRV record and formats an endpoint target string.
func (s Target) Format() string {
	return fmt.Sprintf("%d %d %d %s", s.Priority, s.Weight, s.Port, s.Target)
}

// Record allows SRV record names and targets to be manipulated in
// supported providers.
type Record struct {
	Name   Name
	Target Target
}

// ParseRecord decodes a DNS name into an SRV record for providers.
func ParseRecord(name, target string) Record {
	return Record{
		Name:   ParseName(name),
		Target: ParseTarget(target),
	}
}

// ParseAnnotation creates an SRV record based on an SRV annotation.
func ParseAnnotation(annotation string) Record {
	var name string
	var priority, weight, target int

	fmt.Sscanf(annotation, "%s %d %d %d", &name, &priority, &weight, &target)

	return Record{
		Name: ParseName(name),
		Target: Target{
			Priority: priority,
			Weight:   weight,
			Port:     target,
		},
	}
}
