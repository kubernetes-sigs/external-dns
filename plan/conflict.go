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

package plan

import (
	"fmt"
	"sort"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

// ConflictResolver is used to make a decision in case of two or more different kubernetes resources
// are trying to acquire same DNS name
type ConflictResolver interface {
	Resolve(currents []*endpoint.Endpoint, candidates []*endpoint.Endpoint) (changes Changes, err error)
}

// PerResource allows only one resource to own a given dns name
type PerResource struct{}

// ResolveUpdate is invoked when dns name is already owned by "current" endpoint
// ResolveUpdate uses "current" record as base and updates it accordingly with new version of same resource
// if it doesn't exist then pick min

type wrapEndpoint struct {
	endpoint  *endpoint.Endpoint
	isCurrent bool
	targetIdx int
	order     int
}

func (w wrapEndpoint) target() string {
	return w.endpoint.Targets[w.targetIdx]
}

func (w wrapEndpoint) setIdentifier() string {
	return w.endpoint.SetIdentifier
}

func (w wrapEndpoint) recordTTL() endpoint.TTL {
	return w.endpoint.RecordTTL
}

func (w wrapEndpoint) recordType() string {
	return w.endpoint.RecordType
}

func (w wrapEndpoint) dnsName() string {
	return w.endpoint.DNSName
}

type wrapEndpoints []wrapEndpoint

func (ws wrapEndpoints) Len() int {
	return len(ws)
}

// this sorts the same target in the input order
func (ws wrapEndpoints) Less(i, j int) bool {
	wi := ws[i]
	wj := ws[j]
	// we don't care about dnsName, recordType, setIdentifier
	// they must be the same
	val := strings.Compare(wi.dnsName(), wj.dnsName())
	if val != 0 {
		return val < 0
	}
	val = strings.Compare(wi.recordType(), wj.recordType())
	if val != 0 {
		return val < 0
	}
	val = strings.Compare(wi.setIdentifier(), wj.setIdentifier())
	if val != 0 {
		return val < 0
	}
	val = strings.Compare(wi.target(), wj.target())
	if val != 0 {
		return val < 0
	}
	// this sorts the same targets in the input order
	// to have a defined marge behavior
	return wi.order < wj.order
}

func (ws wrapEndpoints) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func toWrapEndpoints(ep *endpoint.Endpoint, order int, isCurrent bool) wrapEndpoints {
	result := wrapEndpoints{}
	workingEp := ep.DeepCopy()
	// there is a situation where we have more than 65536 endpoints
	// but i think that will break the dns provider
	order = order << 16
	for idx, _ := range ep.Targets {
		result = append(result, wrapEndpoint{
			endpoint:  workingEp,
			isCurrent: isCurrent,
			targetIdx: idx,
			order:     order + idx,
		})
	}
	return result
}

// func equalWrapNameType(a, b *wrapEndpoint) bool {
// 	if a.dnsName() != b.dnsName() {
// 		return false
// 	}
// 	if a.recordType() != b.recordType() {
// 		return false
// 	}
// 	return true
// }

func equalWrapEndpoint(a, b *wrapEndpoint) bool {
	if a.dnsName() != b.dnsName() {
		return false
	}
	if a.recordType() != b.recordType() {
		return false
	}
	if a.setIdentifier() != b.setIdentifier() {
		return false
	}
	if a.recordTTL() != b.recordTTL() {
		return false
	}
	if a.target() != b.target() {
		return false
	}
	return true
}

func uniqueWrapEndpoints(ws wrapEndpoints) wrapEndpoints {
	// sort by dnsName, recordType, target and order to have a defined merge behavior
	sort.Sort(ws)
	result := make(wrapEndpoints, 0, len(ws))
	for i := 0; i < len(ws); i++ {
		if i == 0 {
			result = append(result, ws[i])
			continue
		}
		if equalWrapEndpoint(&ws[i], &ws[i-1]) {
			// mergeSetIdentifier(ws[i-1].endpoint, ws[i].endpoint)
			ws[i-1].endpoint.Labels = srcWinsMergeLabels(ws[i-1].endpoint.Labels, ws[i].endpoint.Labels)
			mergeProviderSpecific(ws[i-1].endpoint, ws[i].endpoint)
			continue
		}
		result = append(result, ws[i])
	}
	return result
}

func mergeCurrentAndCandidates(ws wrapEndpoints) wrapEndpoints {
	sort.Sort(ws)
	currents := findCurrents(ws)
	candidates := uniqueWrapEndpoints(findCandidates(ws))
	return append(currents, candidates...)
}

func hasCurrentAndCandidate(ws wrapEndpoints) bool {
	current := false
	candidate := false
	for _, wep := range ws {
		if wep.isCurrent {
			current = true
		} else {
			candidate = true
		}
	}
	return current && candidate
}

func labels2ProviderSpecific(lbs endpoint.Labels) endpoint.ProviderSpecific {
	ret := make(endpoint.ProviderSpecific, 0, len(lbs))
	for k, v := range lbs {
		ret = append(ret, endpoint.ProviderSpecificProperty{
			Name:  k,
			Value: v,
		})
	}
	return ret
}

type myProviderSpecific endpoint.ProviderSpecific

func (ps myProviderSpecific) Len() int {
	return len(ps)
}

func (ps myProviderSpecific) Less(i, j int) bool {
	x := strings.Compare(ps[i].Name, ps[j].Name)
	if x < 0 {
		return true
	} else if x == 0 {
		if strings.Compare(ps[i].Value, ps[j].Value) < 0 {
			return true
		}
	}
	return false
}

func (ps myProviderSpecific) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func equalProviderSpecific(ps1, ps2 endpoint.ProviderSpecific) bool {
	if len(ps1) != len(ps2) {
		return false
	}
	sort.Sort(myProviderSpecific(ps1))
	sort.Sort(myProviderSpecific(ps2))
	for i := 0; i < len(ps1); i++ {
		if ps1[i].Name != ps2[i].Name || ps1[i].Value != ps2[i].Value {
			return false
		}
	}
	return true
}

func equalLabels(l1, l2 endpoint.Labels) bool {
	if len(l1) != len(l2) {
		return false
	}
	return equalProviderSpecific(labels2ProviderSpecific(l1), labels2ProviderSpecific(l2))
}

func onlyCurrent(ws wrapEndpoints) bool {
	ret := true
	for _, wep := range ws {
		ret = ret && wep.isCurrent
	}
	return ret
}

func findCandidates(ws wrapEndpoints) wrapEndpoints {
	ret := wrapEndpoints{}
	for _, wep := range ws {
		if !wep.isCurrent {
			ret = append(ret, wep)
		}
	}
	return ret
}

// func findCandidate(ws wrapEndpoints) *wrapEndpoint {
// 	return &findCandidates(ws)[0]
// }

func findCurrents(ws wrapEndpoints) wrapEndpoints {
	ret := wrapEndpoints{}
	for _, wep := range ws {
		if wep.isCurrent {
			ret = append(ret, wep)
		}
	}
	return ret
}

func findCurrentEndpoints(ws wrapEndpoints) []*endpoint.Endpoint {
	sort.Sort(ws)
	weps := map[string]wrapEndpoint{}
	for _, wep := range findCurrents(ws) {
		weps[fmt.Sprintf("%p", wep.endpoint)] = wep
	}
	orderWeps := make(wrapEndpoints, 0, len(weps))
	for _, wep := range weps {
		orderWeps = append(orderWeps, wep)
	}
	sort.Slice(orderWeps, func(i, j int) bool {
		return orderWeps[i].order < orderWeps[j].order
	})
	ret := make([]*endpoint.Endpoint, 0, len(weps))
	for _, wep := range orderWeps {
		ret = append(ret, wep.endpoint)
	}
	return ret
}

// func findCurrent(ws wrapEndpoints) *wrapEndpoint {
// 	return &findCurrents(ws)[0]
// }

func equalEndpointKey(a, b *endpoint.Endpoint) bool {
	return a.DNSName == b.DNSName && a.RecordType == b.RecordType && a.SetIdentifier == b.SetIdentifier
}

func (wes wrapEndpoints) append(ep *endpoint.Endpoint, isCurrent bool) (wrapEndpoints, error) {
	if len(wes) != 0 {
		prev := wes[len(wes)-1]
		// ensure that the group is homogeneous
		if !equalEndpointKey(prev.endpoint, ep) {
			return wes, fmt.Errorf("cannot append endpoint %v to group %v", ep, wes)
		}
	}
	order := len(wes)
	return append(wes, toWrapEndpoints(ep, order, isCurrent)...), nil
}

// func mergeSetIdentifier(dest, src *endpoint.Endpoint) {
// 	if src.SetIdentifier != "" {
// 		dest.SetIdentifier = src.SetIdentifier
// 	}
// }

// func destWinsMergeLabels(dest, src *endpoint.Endpoint) {
// 	if src.Labels == nil {
// 		return
// 	}
// 	if dest.Labels == nil {
// 		dest.Labels = make(endpoint.Labels)
// 	}
// 	for k, v := range src.Labels {
// 		_, found := dest.Labels[k]
// 		if found {
// 			continue
// 		}
// 		dest.Labels[k] = v
// 	}
// }

func srcWinsMergeLabels(dst, src map[string]string) map[string]string {
	if src == nil || len(src) == 0 {
		return dst
	}
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func mergeProviderSpecific(dest, src *endpoint.Endpoint) {
	destMap := make(map[string]string)
	for _, p := range dest.ProviderSpecific {
		destMap[p.Name] = p.Value
	}
	for _, p := range src.ProviderSpecific {
		_, found := destMap[p.Name]
		if found {
			continue
		}
		dest.ProviderSpecific = append(dest.ProviderSpecific, p)
	}
}

func equalEndpointWithoutLabels(e1, e2 *endpoint.Endpoint) bool {
	if e1.DNSName != e2.DNSName {
		return false
	}
	if !e1.Targets.Same(e2.Targets) {
		return false
	}
	if e1.RecordType != e2.RecordType {
		return false
	}
	if e1.SetIdentifier != e2.SetIdentifier {
		return false
	}
	if e1.RecordTTL != e2.RecordTTL {
		return false
	}
	if !equalProviderSpecific(e1.ProviderSpecific, e2.ProviderSpecific) {
		return false
	}
	return true
}

func equalEndpoint(e1, e2 *endpoint.Endpoint) bool {
	if !equalEndpointWithoutLabels(e1, e2) {
		return false
	}
	if !equalLabels(e1.Labels, e2.Labels) {
		return false
	}
	return true
}

func mergeEndpointTargets(ep *endpoint.Endpoint) error {
	if len(ep.Targets) == 0 {
		return fmt.Errorf("empty targets are not allowed")
	}
	switch ep.RecordType {
	case endpoint.RecordTypePTR:
		for i := 1; i < len(ep.Targets); i++ {
			if ep.Targets[i] != ep.Targets[0] {
				return fmt.Errorf("inconsistent targets for PTR record")
			}
		}
		ep.Targets = []string{ep.Targets[0]}
	case endpoint.RecordTypeCNAME:
		for i := 1; i < len(ep.Targets); i++ {
			if ep.Targets[i] != ep.Targets[0] {
				return fmt.Errorf("inconsistent targets for CNAME record")
			}
		}
		ep.Targets = []string{ep.Targets[0]}

	default:
		// we don't merge targets for other record types
	}
	return nil
}

func toEndpoint(ws wrapEndpoints) (*endpoint.Endpoint, error) {
	if len(ws) == 0 {
		return nil, fmt.Errorf("empty wrapEndpoints")
	}
	// sort by target
	sort.Sort(ws)

	ret := ws[0].endpoint.DeepCopy()
	ret.Targets = []string{ws[0].target()}
	for i := 1; i < len(ws); i++ {
		wep := ws[i]
		if ret.DNSName != wep.dnsName() {
			return nil, fmt.Errorf("inconsistent DNSName")
		}
		if ret.RecordType != wep.recordType() {
			return nil, fmt.Errorf("inconsistent RecordType")
		}
		if ret.SetIdentifier != wep.setIdentifier() {
			return nil, fmt.Errorf("inconsistent SetIdentifier")
		}
		if ret.RecordTTL != wep.recordTTL() {
			return nil, fmt.Errorf("inconsistent RecordTTL")
		}
		// check if target is duplicated
		if ws[i-1].target() != wep.target() {
			ret.Targets = append(ret.Targets, wep.target())
		}
		// src wins merge
		ret.Labels = srcWinsMergeLabels(ret.Labels, wep.endpoint.Labels)
		mergeProviderSpecific(ret, wep.endpoint)
	}
	err := mergeEndpointTargets(ret)

	return ret, err
}

func (s PerResource) Resolve(currents []*endpoint.Endpoint, candidates []*endpoint.Endpoint) (changes Changes, err error) {
	wes := make(wrapEndpoints, 0, len(currents)+len(candidates))
	for _, ep := range currents {
		wes, err = wes.append(ep, true)
		if err != nil {
			return
		}
	}
	for _, ep := range candidates {
		wes, err = wes.append(ep, false)
		if err != nil {
			return
		}
	}

	mcac := mergeCurrentAndCandidates(wes)
	var ep *endpoint.Endpoint
	if onlyCurrent(mcac) {
		// ep, err = toEndpoint(mcac)
		// if err != nil {
		// 	return
		// }
		changes.Delete = append(changes.Delete, findCurrentEndpoints(mcac)...)
		return
	}
	if hasCurrentAndCandidate(mcac) {
		currents := findCurrents(mcac)
		sort.Sort(currents) // sort by target
		var currentMergedLabels map[string]string
		for _, current := range currents {
			currentMergedLabels = srcWinsMergeLabels(currentMergedLabels, current.endpoint.Labels)
		}
		candidates := findCandidates(mcac)
		sort.Sort(candidates) // sort by target

		var candidateEP *endpoint.Endpoint
		candidateEP, err = toEndpoint(candidates)
		if err != nil {
			return
		}
		candidateEP.Labels = srcWinsMergeLabels(currentMergedLabels, candidateEP.Labels)
		oldEPs := findCurrentEndpoints(mcac)
		foundCandidate := false
		for _, oldEp := range oldEPs {
			if equalEndpoint(oldEp, candidateEP) {
				foundCandidate = true
				continue
			}
			changes.UpdateOld = append(changes.UpdateOld, oldEp)
		}
		if !foundCandidate {
			changes.UpdateNew = append(changes.UpdateNew, candidateEP)
		}
		return
	}
	ep, err = toEndpoint(mcac)
	if err != nil {
		return
	}
	changes.Create = append(changes.Create, ep)
	return // s.ResolveCreate(candidates)
}

// less returns true if endpoint x is less than y
// func (s PerResource) less(x, y *endpoint.Endpoint) bool {
// 	return x.Targets.IsLess(y.Targets)
// }

// TODO: with cross-resource/cross-cluster setup alternative variations of ConflictResolver can be used
