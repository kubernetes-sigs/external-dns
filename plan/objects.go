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

package plan

import "sigs.k8s.io/external-dns/pkg/events"

// PlannedObject is a Kubernetes object that contributed desired endpoints to a plan.
type PlannedObject struct {
	Ref *events.ObjectReference
	// Endpoints left after the domain and record-type filters; 0 if none.
	Endpoints int
}

// PlannedObjects returns every object behind a desired endpoint, changed or not,
// once each in first-seen order. Call it on the result of Calculate.
func (p *Plan) PlannedObjects() []PlannedObject {
	byKey := map[string]*PlannedObject{}
	var keys []string
	for _, ep := range p.Desired {
		for _, ref := range ep.RefObjects() {
			if ref == nil {
				continue
			}
			if _, ok := byKey[ref.Key()]; !ok {
				byKey[ref.Key()] = &PlannedObject{Ref: ref}
				keys = append(keys, ref.Key())
			}
		}
	}

	for _, ep := range p.Planned {
		for _, ref := range ep.RefObjects() {
			if ref == nil {
				continue
			}
			if obj, ok := byKey[ref.Key()]; ok {
				obj.Endpoints++
			}
		}
	}

	objects := make([]PlannedObject, 0, len(keys))
	for _, key := range keys {
		objects = append(objects, *byKey[key])
	}
	return objects
}
