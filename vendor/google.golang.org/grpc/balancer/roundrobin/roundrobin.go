/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package roundrobin defines a roundrobin balancer. Roundrobin balancer is
// installed as one of the default balancers in gRPC, users don't need to
// explicitly install this balancer.
package roundrobin

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
=======
	"math/rand"
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	"sync/atomic"
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	"math/rand"
	"sync/atomic"
=======
	"fmt"
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/endpointsharding"
	"google.golang.org/grpc/balancer/pickfirst"
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

// Name is the name of round_robin balancer.
const Name = "round_robin"

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
var logger = grpclog.Component("roundrobin")

// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	logger.Infof("roundrobinPicker: newPicker called with info: %v", info)
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
var logger = grpclog.Component("roundrobin")

>>>>>>> 5ce8c7613 (update vendored files)
// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	logger.Infof("roundrobinPicker: newPicker called with info: %v", info)
	if len(info.ReadySCs) == 0 {
<<<<<<< HEAD
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
=======
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
var logger = grpclog.Component("roundrobin")

>>>>>>> 6b7ce455e (update vendored files)
// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	logger.Infof("roundrobinPicker: Build called with info: %v", info)
	if len(info.ReadySCs) == 0 {
<<<<<<< HEAD
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
=======
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
>>>>>>> 6b7ce455e (update vendored files)
	}
	scs := make([]balancer.SubConn, 0, len(info.ReadySCs))
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
var logger = grpclog.Component("roundrobin")

>>>>>>> 4d7e5ad26 (update vendored files)
// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	logger.Infof("roundrobinPicker: Build called with info: %v", info)
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
<<<<<<< HEAD
	var scs []balancer.SubConn
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	var scs []balancer.SubConn
=======
	scs := make([]balancer.SubConn, 0, len(info.ReadySCs))
>>>>>>> 4d7e5ad26 (update vendored files)
	for sc := range info.ReadySCs {
		scs = append(scs, sc)
	}
	return &rrPicker{
		subConns: scs,
		// Start at a random index, as the same RR balancer rebuilds a new
		// picker when SubConn states change, and we don't want to apply excess
		// load to the first server in the list.
		next: uint32(grpcrand.Intn(len(scs))),
	}
}

type rrPicker struct {
	// subConns is the snapshot of the roundrobin balancer when this picker was
	// created. The slice is immutable. Each Get() will do a round robin
	// selection from it and return the selected SubConn.
	subConns []balancer.SubConn
	next     uint32
}

func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	subConnsLen := uint32(len(p.subConns))
	nextIndex := atomic.AddUint32(&p.next, 1)

	sc := p.subConns[nextIndex%subConnsLen]
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"sync"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"sync"
=======
	"sync/atomic"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/grpcrand"
)

// Name is the name of round_robin balancer.
const Name = "round_robin"

var logger = grpclog.Component("roundrobin")

func init() {
	balancer.Register(builder{})
}

type builder struct{}

func (bb builder) Name() string {
	return Name
}

func (bb builder) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
	childBuilder := balancer.Get(pickfirst.Name).Build
	bal := &rrBalancer{
		cc:       cc,
		Balancer: endpointsharding.NewBalancer(cc, opts, childBuilder, endpointsharding.Options{}),
	}
	bal.logger = internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf("[%p] ", bal))
	bal.logger.Infof("Created")
	return bal
}

type rrBalancer struct {
	balancer.Balancer
	cc     balancer.ClientConn
	logger *internalgrpclog.PrefixLogger
}

<<<<<<< HEAD
func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
<<<<<<< HEAD
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
=======
	subConnsLen := uint32(len(p.subConns))
	nextIndex := atomic.AddUint32(&p.next, 1)

	sc := p.subConns[nextIndex%subConnsLen]
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return balancer.PickResult{SubConn: sc}, nil
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	subConnsLen := uint32(len(p.subConns))
	nextIndex := atomic.AddUint32(&p.next, 1)

	sc := p.subConns[nextIndex%subConnsLen]
	return balancer.PickResult{SubConn: sc}, nil
=======
func (b *rrBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	return b.Balancer.UpdateClientConnState(balancer.ClientConnState{
		// Enable the health listener in pickfirst children for client side health
		// checks and outlier detection, if configured.
		ResolverState: pickfirst.EnableHealthListener(ccs.ResolverState),
	})
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
}
