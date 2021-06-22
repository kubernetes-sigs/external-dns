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

package grpc

import (
<<<<<<< HEAD
	"context"
	"strings"
	"sync"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/pretty"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// resolverStateUpdater wraps the single method used by ccResolverWrapper to
// report a state update from the actual resolver implementation.
type resolverStateUpdater interface {
	updateResolverState(s resolver.State, err error) error
}

// ccResolverWrapper is a wrapper on top of cc for resolvers.
// It implements resolver.ClientConn interface.
type ccResolverWrapper struct {
	// The following fields are initialized when the wrapper is created and are
	// read-only afterwards, and therefore can be accessed without a mutex.
	cc                  resolverStateUpdater
	channelzID          *channelz.Identifier
	ignoreServiceConfig bool
	opts                ccResolverWrapperOpts
	serializer          *grpcsync.CallbackSerializer // To serialize all incoming calls.
	serializerCancel    context.CancelFunc           // To close the serializer, accessed only from close().

	// All incoming (resolver --> gRPC) calls are guaranteed to execute in a
	// mutually exclusive manner as they are scheduled on the serializer.
	// Fields accessed *only* in these serializer callbacks, can therefore be
	// accessed without a mutex.
	curState resolver.State

	// mu guards access to the below fields.
	mu       sync.Mutex
	closed   bool
	resolver resolver.Resolver // Accessed only from outgoing calls.
}

// ccResolverWrapperOpts wraps the arguments to be passed when creating a new
// ccResolverWrapper.
type ccResolverWrapperOpts struct {
	target     resolver.Target       // User specified dial target to resolve.
	builder    resolver.Builder      // Resolver builder to use.
	bOpts      resolver.BuildOptions // Resolver build options to use.
	channelzID *channelz.Identifier  // Channelz identifier for the channel.
}

// newCCResolverWrapper uses the resolver.Builder to build a Resolver and
// returns a ccResolverWrapper object which wraps the newly built resolver.
func newCCResolverWrapper(cc resolverStateUpdater, opts ccResolverWrapperOpts) (*ccResolverWrapper, error) {
	ctx, cancel := context.WithCancel(context.Background())
	ccr := &ccResolverWrapper{
		cc:                  cc,
		channelzID:          opts.channelzID,
		ignoreServiceConfig: opts.bOpts.DisableServiceConfig,
		opts:                opts,
		serializer:          grpcsync.NewCallbackSerializer(ctx),
		serializerCancel:    cancel,
	}

	// Cannot hold the lock at build time because the resolver can send an
	// update or error inline and these incoming calls grab the lock to schedule
	// a callback in the serializer.
	r, err := opts.builder.Build(opts.target, ccr, opts.bOpts)
	if err != nil {
		cancel()
		return nil, err
	}

	// Any error reported by the resolver at build time that leads to a
	// re-resolution request from the balancer is dropped by grpc until we
	// return from this function. So, we don't have to handle pending resolveNow
	// requests here.
	ccr.mu.Lock()
	ccr.resolver = r
	ccr.mu.Unlock()

	return ccr, nil
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	ccr.mu.Lock()
	defer ccr.mu.Unlock()

	// ccr.resolver field is set only after the call to Build() returns. But in
	// the process of building, the resolver may send an error update which when
	// propagated to the balancer may result in a re-resolution request.
	if ccr.closed || ccr.resolver == nil {
		return
	}
	ccr.resolver.ResolveNow(o)
}

func (ccr *ccResolverWrapper) close() {
	ccr.mu.Lock()
	if ccr.closed {
		ccr.mu.Unlock()
		return
	}

	channelz.Info(logger, ccr.channelzID, "Closing the name resolver")

	// Close the serializer to ensure that no more calls from the resolver are
	// handled, before actually closing the resolver.
	ccr.serializerCancel()
	ccr.closed = true
	r := ccr.resolver
	ccr.mu.Unlock()

	// Give enqueued callbacks a chance to finish.
	<-ccr.serializer.Done

	// Spawn a goroutine to close the resolver (since it may block trying to
	// cleanup all allocated resources) and return early.
	go r.Close()
}

// serializerScheduleLocked is a convenience method to schedule a function to be
// run on the serializer while holding ccr.mu.
func (ccr *ccResolverWrapper) serializerScheduleLocked(f func(context.Context)) {
	ccr.mu.Lock()
	ccr.serializer.Schedule(f)
	ccr.mu.Unlock()
}

// UpdateState is called by resolver implementations to report new state to gRPC
// which includes addresses and service config.
func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
	errCh := make(chan error, 1)
	ok := ccr.serializer.Schedule(func(context.Context) {
		ccr.addChannelzTraceEvent(s)
		ccr.curState = s
		if err := ccr.cc.updateResolverState(ccr.curState, nil); err == balancer.ErrBadResolverState {
			errCh <- balancer.ErrBadResolverState
			return
		}
		errCh <- nil
	})
	if !ok {
		// The only time when Schedule() fail to add the callback to the
		// serializer is when the serializer is closed, and this happens only
		// when the resolver wrapper is closed.
		return nil
	}
	return <-errCh
}

// ReportError is called by resolver implementations to report errors
// encountered during name resolution to gRPC.
func (ccr *ccResolverWrapper) ReportError(err error) {
	ccr.serializerScheduleLocked(func(_ context.Context) {
		channelz.Warningf(logger, ccr.channelzID, "ccResolverWrapper: reporting error to cc: %v", err)
		ccr.cc.updateResolverState(resolver.State{}, err)
	})
}

// NewAddress is called by the resolver implementation to send addresses to
// gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	ccr.serializerScheduleLocked(func(_ context.Context) {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
		ccr.curState.Addresses = addrs
		ccr.cc.updateResolverState(ccr.curState, nil)
	})
}

// NewServiceConfig is called by the resolver implementation to send service
// configs to gRPC.
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
	ccr.serializerScheduleLocked(func(_ context.Context) {
		channelz.Infof(logger, ccr.channelzID, "ccResolverWrapper: got new service config: %s", sc)
		if ccr.ignoreServiceConfig {
			channelz.Info(logger, ccr.channelzID, "Service config lookups disabled; ignoring config")
			return
		}
		scpr := parseServiceConfig(sc)
		if scpr.Err != nil {
			channelz.Warningf(logger, ccr.channelzID, "ccResolverWrapper: error parsing service config: %v", scpr.Err)
			return
		}
		ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
		ccr.curState.ServiceConfig = scpr
		ccr.cc.updateResolverState(ccr.curState, nil)
	})
}

// ParseServiceConfig is called by resolver implementations to parse a JSON
// representation of the service config.
func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	return parseServiceConfig(scJSON)
}

// addChannelzTraceEvent adds a channelz trace event containing the new
// state received from resolver implementations.
func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	var updates []string
	var oldSC, newSC *ServiceConfig
	var oldOK, newOK bool
	if ccr.curState.ServiceConfig != nil {
		oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
	}
	if s.ServiceConfig != nil {
		newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
	}
	if oldOK != newOK || (oldOK && newOK && oldSC.rawJSONString != newSC.rawJSONString) {
		updates = append(updates, "service config updated")
	}
	if len(ccr.curState.Addresses) > 0 && len(s.Addresses) == 0 {
		updates = append(updates, "resolver returned an empty address list")
	} else if len(ccr.curState.Addresses) == 0 && len(s.Addresses) > 0 {
		updates = append(updates, "resolver returned new addresses")
	}
<<<<<<< HEAD
<<<<<<< HEAD
	channelz.AddTraceEvent(logger, ccr.cc.channelzID, 0, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
		Severity: channelz.CtInfo,
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"time"
||||||| parent of 5ce8c7613 (update vendored files)
	"time"
=======
>>>>>>> 5ce8c7613 (update vendored files)

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// ccResolverWrapper is a wrapper on top of cc for resolvers.
// It implements resolver.ClientConn interface.
type ccResolverWrapper struct {
	cc         *ClientConn
	resolverMu sync.Mutex
	resolver   resolver.Resolver
	done       *grpcsync.Event
	curState   resolver.State

	incomingMu sync.Mutex // Synchronizes all the incoming calls.
}

// newCCResolverWrapper uses the resolver.Builder to build a Resolver and
// returns a ccResolverWrapper object which wraps the newly built resolver.
func newCCResolverWrapper(cc *ClientConn, rb resolver.Builder) (*ccResolverWrapper, error) {
	ccr := &ccResolverWrapper{
		cc:   cc,
		done: grpcsync.NewEvent(),
	}

	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
		credsClone = creds.Clone()
	}
	rbo := resolver.BuildOptions{
		DisableServiceConfig: cc.dopts.disableServiceConfig,
		DialCreds:            credsClone,
		CredsBundle:          cc.dopts.copts.CredsBundle,
		Dialer:               cc.dopts.copts.Dialer,
	}

	var err error
	// We need to hold the lock here while we assign to the ccr.resolver field
	// to guard against a data race caused by the following code path,
	// rb.Build-->ccr.ReportError-->ccr.poll-->ccr.resolveNow, would end up
	// accessing ccr.resolver which is being assigned here.
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
		return nil, err
	}
	return ccr, nil
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	ccr.resolverMu.Lock()
	if !ccr.done.HasFired() {
		ccr.resolver.ResolveNow(o)
	}
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) close() {
	ccr.resolverMu.Lock()
	ccr.resolver.Close()
	ccr.done.Fire()
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return nil
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending update to cc: %v", s)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(s)
	}
	ccr.curState = s
	if err := ccr.cc.updateResolverState(ccr.curState, nil); err == balancer.ErrBadResolverState {
		return balancer.ErrBadResolverState
	}
	return nil
}

func (ccr *ccResolverWrapper) ReportError(err error) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: reporting error to cc: %v", err)
	ccr.cc.updateResolverState(resolver.State{}, err)
}

// NewAddress is called by the resolver implementation to send addresses to gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending new addresses to cc: %v", addrs)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
	}
	ccr.curState.Addresses = addrs
	ccr.cc.updateResolverState(ccr.curState, nil)
}

// NewServiceConfig is called by the resolver implementation to send service
// configs to gRPC.
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: got new service config: %v", sc)
	if ccr.cc.dopts.disableServiceConfig {
		channelz.Info(logger, ccr.cc.channelzID, "Service config lookups disabled; ignoring config")
		return
	}
	scpr := parseServiceConfig(sc)
	if scpr.Err != nil {
		channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: error parsing service config: %v", scpr.Err)
		return
	}
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
	}
	ccr.curState.ServiceConfig = scpr
	ccr.cc.updateResolverState(ccr.curState, nil)
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	return parseServiceConfig(scJSON)
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	var updates []string
	var oldSC, newSC *ServiceConfig
	var oldOK, newOK bool
	if ccr.curState.ServiceConfig != nil {
		oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
	}
	if s.ServiceConfig != nil {
		newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
	}
	if oldOK != newOK || (oldOK && newOK && oldSC.rawJSONString != newSC.rawJSONString) {
		updates = append(updates, "service config updated")
	}
	if len(ccr.curState.Addresses) > 0 && len(s.Addresses) == 0 {
		updates = append(updates, "resolver returned an empty address list")
	} else if len(ccr.curState.Addresses) == 0 && len(s.Addresses) > 0 {
		updates = append(updates, "resolver returned new addresses")
	}
	channelz.AddTraceEvent(logger, ccr.cc.channelzID, 0, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
<<<<<<< HEAD
		Severity: channelz.CtINFO,
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		Severity: channelz.CtINFO,
=======
		Severity: channelz.CtInfo,
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"time"
||||||| parent of 6b7ce455e (update vendored files)
	"time"
=======
>>>>>>> 6b7ce455e (update vendored files)

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// ccResolverWrapper is a wrapper on top of cc for resolvers.
// It implements resolver.ClientConn interface.
type ccResolverWrapper struct {
	cc         *ClientConn
	resolverMu sync.Mutex
	resolver   resolver.Resolver
	done       *grpcsync.Event
	curState   resolver.State

	incomingMu sync.Mutex // Synchronizes all the incoming calls.
}

// newCCResolverWrapper uses the resolver.Builder to build a Resolver and
// returns a ccResolverWrapper object which wraps the newly built resolver.
func newCCResolverWrapper(cc *ClientConn, rb resolver.Builder) (*ccResolverWrapper, error) {
	ccr := &ccResolverWrapper{
		cc:   cc,
		done: grpcsync.NewEvent(),
	}

	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
		credsClone = creds.Clone()
	}
	rbo := resolver.BuildOptions{
		DisableServiceConfig: cc.dopts.disableServiceConfig,
		DialCreds:            credsClone,
		CredsBundle:          cc.dopts.copts.CredsBundle,
		Dialer:               cc.dopts.copts.Dialer,
	}

	var err error
	// We need to hold the lock here while we assign to the ccr.resolver field
	// to guard against a data race caused by the following code path,
	// rb.Build-->ccr.ReportError-->ccr.poll-->ccr.resolveNow, would end up
	// accessing ccr.resolver which is being assigned here.
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
		return nil, err
	}
	return ccr, nil
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	ccr.resolverMu.Lock()
	if !ccr.done.HasFired() {
		ccr.resolver.ResolveNow(o)
	}
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) close() {
	ccr.resolverMu.Lock()
	ccr.resolver.Close()
	ccr.done.Fire()
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return nil
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending update to cc: %v", s)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(s)
	}
	ccr.curState = s
	if err := ccr.cc.updateResolverState(ccr.curState, nil); err == balancer.ErrBadResolverState {
		return balancer.ErrBadResolverState
	}
	return nil
}

func (ccr *ccResolverWrapper) ReportError(err error) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: reporting error to cc: %v", err)
	ccr.cc.updateResolverState(resolver.State{}, err)
}

// NewAddress is called by the resolver implementation to send addresses to gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending new addresses to cc: %v", addrs)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
	}
	ccr.curState.Addresses = addrs
	ccr.cc.updateResolverState(ccr.curState, nil)
}

// NewServiceConfig is called by the resolver implementation to send service
// configs to gRPC.
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: got new service config: %v", sc)
	if ccr.cc.dopts.disableServiceConfig {
		channelz.Info(logger, ccr.cc.channelzID, "Service config lookups disabled; ignoring config")
		return
	}
	scpr := parseServiceConfig(sc)
	if scpr.Err != nil {
		channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: error parsing service config: %v", scpr.Err)
		return
	}
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
	}
	ccr.curState.ServiceConfig = scpr
	ccr.cc.updateResolverState(ccr.curState, nil)
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	return parseServiceConfig(scJSON)
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	var updates []string
	var oldSC, newSC *ServiceConfig
	var oldOK, newOK bool
	if ccr.curState.ServiceConfig != nil {
		oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
	}
	if s.ServiceConfig != nil {
		newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
	}
	if oldOK != newOK || (oldOK && newOK && oldSC.rawJSONString != newSC.rawJSONString) {
		updates = append(updates, "service config updated")
	}
	if len(ccr.curState.Addresses) > 0 && len(s.Addresses) == 0 {
		updates = append(updates, "resolver returned an empty address list")
	} else if len(ccr.curState.Addresses) == 0 && len(s.Addresses) > 0 {
		updates = append(updates, "resolver returned new addresses")
	}
	channelz.AddTraceEvent(logger, ccr.cc.channelzID, 0, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
<<<<<<< HEAD
		Severity: channelz.CtINFO,
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		Severity: channelz.CtINFO,
=======
		Severity: channelz.CtInfo,
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"time"
||||||| parent of 4d7e5ad26 (update vendored files)
	"time"
=======
>>>>>>> 4d7e5ad26 (update vendored files)

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// ccResolverWrapper is a wrapper on top of cc for resolvers.
// It implements resolver.ClientConn interface.
type ccResolverWrapper struct {
	cc         *ClientConn
	resolverMu sync.Mutex
	resolver   resolver.Resolver
	done       *grpcsync.Event
	curState   resolver.State

	incomingMu sync.Mutex // Synchronizes all the incoming calls.
}

// newCCResolverWrapper uses the resolver.Builder to build a Resolver and
// returns a ccResolverWrapper object which wraps the newly built resolver.
func newCCResolverWrapper(cc *ClientConn, rb resolver.Builder) (*ccResolverWrapper, error) {
	ccr := &ccResolverWrapper{
		cc:   cc,
		done: grpcsync.NewEvent(),
	}

	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
		credsClone = creds.Clone()
	}
	rbo := resolver.BuildOptions{
		DisableServiceConfig: cc.dopts.disableServiceConfig,
		DialCreds:            credsClone,
		CredsBundle:          cc.dopts.copts.CredsBundle,
		Dialer:               cc.dopts.copts.Dialer,
	}

	var err error
	// We need to hold the lock here while we assign to the ccr.resolver field
	// to guard against a data race caused by the following code path,
	// rb.Build-->ccr.ReportError-->ccr.poll-->ccr.resolveNow, would end up
	// accessing ccr.resolver which is being assigned here.
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
		return nil, err
	}
	return ccr, nil
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	ccr.resolverMu.Lock()
	if !ccr.done.HasFired() {
		ccr.resolver.ResolveNow(o)
	}
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) close() {
	ccr.resolverMu.Lock()
	ccr.resolver.Close()
	ccr.done.Fire()
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return nil
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending update to cc: %v", s)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(s)
	}
	ccr.curState = s
	if err := ccr.cc.updateResolverState(ccr.curState, nil); err == balancer.ErrBadResolverState {
		return balancer.ErrBadResolverState
	}
	return nil
}

func (ccr *ccResolverWrapper) ReportError(err error) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: reporting error to cc: %v", err)
	ccr.cc.updateResolverState(resolver.State{}, err)
}

// NewAddress is called by the resolver implementation to send addresses to gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: sending new addresses to cc: %v", addrs)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
	}
	ccr.curState.Addresses = addrs
	ccr.cc.updateResolverState(ccr.curState, nil)
}

// NewServiceConfig is called by the resolver implementation to send service
// configs to gRPC.
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
	ccr.incomingMu.Lock()
	defer ccr.incomingMu.Unlock()
	if ccr.done.HasFired() {
		return
	}
	channelz.Infof(logger, ccr.cc.channelzID, "ccResolverWrapper: got new service config: %v", sc)
	if ccr.cc.dopts.disableServiceConfig {
		channelz.Info(logger, ccr.cc.channelzID, "Service config lookups disabled; ignoring config")
		return
	}
	scpr := parseServiceConfig(sc)
	if scpr.Err != nil {
		channelz.Warningf(logger, ccr.cc.channelzID, "ccResolverWrapper: error parsing service config: %v", scpr.Err)
		return
	}
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
	}
	ccr.curState.ServiceConfig = scpr
	ccr.cc.updateResolverState(ccr.curState, nil)
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	return parseServiceConfig(scJSON)
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	var updates []string
	var oldSC, newSC *ServiceConfig
	var oldOK, newOK bool
	if ccr.curState.ServiceConfig != nil {
		oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
	}
	if s.ServiceConfig != nil {
		newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
	}
	if oldOK != newOK || (oldOK && newOK && oldSC.rawJSONString != newSC.rawJSONString) {
		updates = append(updates, "service config updated")
	}
	if len(ccr.curState.Addresses) > 0 && len(s.Addresses) == 0 {
		updates = append(updates, "resolver returned an empty address list")
	} else if len(ccr.curState.Addresses) == 0 && len(s.Addresses) > 0 {
		updates = append(updates, "resolver returned new addresses")
	}
	channelz.AddTraceEvent(logger, ccr.cc.channelzID, 0, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
<<<<<<< HEAD
		Severity: channelz.CtINFO,
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		Severity: channelz.CtINFO,
=======
		Severity: channelz.CtInfo,
>>>>>>> 4d7e5ad26 (update vendored files)
	})
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	channelz.AddTraceEvent(logger, ccr.cc.channelzID, 0, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
		Severity: channelz.CtInfo,
	})
=======
	channelz.Infof(logger, ccr.cc.channelzID, "Resolver state updated: %s (%v)", pretty.ToJSON(s), strings.Join(updates, "; "))
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
	channelz.Infof(logger, ccr.cc.channelzID, "Resolver state updated: %s (%v)", pretty.ToJSON(s), strings.Join(updates, "; "))
=======
	channelz.Infof(logger, ccr.channelzID, "Resolver state updated: %s (%v)", pretty.ToJSON(s), strings.Join(updates, "; "))
>>>>>>> 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"fmt"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

// ccResolverWrapper is a wrapper on top of cc for resolvers.
// It implements resolver.ClientConnection interface.
type ccResolverWrapper struct {
	cc         *ClientConn
	resolverMu sync.Mutex
	resolver   resolver.Resolver
	done       *grpcsync.Event
	curState   resolver.State

	pollingMu sync.Mutex
	polling   chan struct{}
}

// split2 returns the values from strings.SplitN(s, sep, 2).
// If sep is not found, it returns ("", "", false) instead.
func split2(s, sep string) (string, string, bool) {
	spl := strings.SplitN(s, sep, 2)
	if len(spl) < 2 {
		return "", "", false
	}
	return spl[0], spl[1], true
}

// parseTarget splits target into a struct containing scheme, authority and
// endpoint.
//
// If target is not a valid scheme://authority/endpoint, it returns {Endpoint:
// target}.
func parseTarget(target string) (ret resolver.Target) {
	var ok bool
	ret.Scheme, ret.Endpoint, ok = split2(target, "://")
	if !ok {
		return resolver.Target{Endpoint: target}
	}
	ret.Authority, ret.Endpoint, ok = split2(ret.Endpoint, "/")
	if !ok {
		return resolver.Target{Endpoint: target}
	}
	return ret
}

// newCCResolverWrapper uses the resolver.Builder stored in the ClientConn to
// build a Resolver and returns a ccResolverWrapper object which wraps the
// newly built resolver.
func newCCResolverWrapper(cc *ClientConn) (*ccResolverWrapper, error) {
	rb := cc.dopts.resolverBuilder
	if rb == nil {
		return nil, fmt.Errorf("could not get resolver for scheme: %q", cc.parsedTarget.Scheme)
	}

	ccr := &ccResolverWrapper{
		cc:   cc,
		done: grpcsync.NewEvent(),
	}

	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
		credsClone = creds.Clone()
	}
	rbo := resolver.BuildOptions{
		DisableServiceConfig: cc.dopts.disableServiceConfig,
		DialCreds:            credsClone,
		CredsBundle:          cc.dopts.copts.CredsBundle,
		Dialer:               cc.dopts.copts.Dialer,
	}

	var err error
	// We need to hold the lock here while we assign to the ccr.resolver field
	// to guard against a data race caused by the following code path,
	// rb.Build-->ccr.ReportError-->ccr.poll-->ccr.resolveNow, would end up
	// accessing ccr.resolver which is being assigned here.
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
		return nil, err
	}
	return ccr, nil
}

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	ccr.resolverMu.Lock()
	if !ccr.done.HasFired() {
		ccr.resolver.ResolveNow(o)
	}
	ccr.resolverMu.Unlock()
}

func (ccr *ccResolverWrapper) close() {
	ccr.resolverMu.Lock()
	ccr.resolver.Close()
	ccr.done.Fire()
	ccr.resolverMu.Unlock()
}

// poll begins or ends asynchronous polling of the resolver based on whether
// err is ErrBadResolverState.
func (ccr *ccResolverWrapper) poll(err error) {
	ccr.pollingMu.Lock()
	defer ccr.pollingMu.Unlock()
	if err != balancer.ErrBadResolverState {
		// stop polling
		if ccr.polling != nil {
			close(ccr.polling)
			ccr.polling = nil
		}
		return
	}
	if ccr.polling != nil {
		// already polling
		return
	}
	p := make(chan struct{})
	ccr.polling = p
	go func() {
		for i := 0; ; i++ {
			ccr.resolveNow(resolver.ResolveNowOptions{})
			t := time.NewTimer(ccr.cc.dopts.resolveNowBackoff(i))
			select {
			case <-p:
				t.Stop()
				return
			case <-ccr.done.Done():
				// Resolver has been closed.
				t.Stop()
				return
			case <-t.C:
				select {
				case <-p:
					return
				default:
				}
				// Timer expired; re-resolve.
			}
		}
	}()
}

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) {
	if ccr.done.HasFired() {
		return
	}
	grpclog.Infof("ccResolverWrapper: sending update to cc: %v", s)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(s)
	}
	ccr.curState = s
	ccr.poll(ccr.cc.updateResolverState(ccr.curState, nil))
}

func (ccr *ccResolverWrapper) ReportError(err error) {
	if ccr.done.HasFired() {
		return
	}
	grpclog.Warningf("ccResolverWrapper: reporting error to cc: %v", err)
	if channelz.IsOn() {
		channelz.AddTraceEvent(ccr.cc.channelzID, &channelz.TraceEventDesc{
			Desc:     fmt.Sprintf("Resolver reported error: %v", err),
			Severity: channelz.CtWarning,
		})
	}
	ccr.poll(ccr.cc.updateResolverState(resolver.State{}, err))
}

// NewAddress is called by the resolver implementation to send addresses to gRPC.
func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	if ccr.done.HasFired() {
		return
	}
	grpclog.Infof("ccResolverWrapper: sending new addresses to cc: %v", addrs)
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: addrs, ServiceConfig: ccr.curState.ServiceConfig})
	}
	ccr.curState.Addresses = addrs
	ccr.poll(ccr.cc.updateResolverState(ccr.curState, nil))
}

// NewServiceConfig is called by the resolver implementation to send service
// configs to gRPC.
func (ccr *ccResolverWrapper) NewServiceConfig(sc string) {
	if ccr.done.HasFired() {
		return
	}
	grpclog.Infof("ccResolverWrapper: got new service config: %v", sc)
	if ccr.cc.dopts.disableServiceConfig {
		grpclog.Infof("Service config lookups disabled; ignoring config")
		return
	}
	scpr := parseServiceConfig(sc)
	if scpr.Err != nil {
		grpclog.Warningf("ccResolverWrapper: error parsing service config: %v", scpr.Err)
		if channelz.IsOn() {
			channelz.AddTraceEvent(ccr.cc.channelzID, &channelz.TraceEventDesc{
				Desc:     fmt.Sprintf("Error parsing service config: %v", scpr.Err),
				Severity: channelz.CtWarning,
			})
		}
		ccr.poll(balancer.ErrBadResolverState)
		return
	}
	if channelz.IsOn() {
		ccr.addChannelzTraceEvent(resolver.State{Addresses: ccr.curState.Addresses, ServiceConfig: scpr})
	}
	ccr.curState.ServiceConfig = scpr
	ccr.poll(ccr.cc.updateResolverState(ccr.curState, nil))
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	return parseServiceConfig(scJSON)
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	var updates []string
	var oldSC, newSC *ServiceConfig
	var oldOK, newOK bool
	if ccr.curState.ServiceConfig != nil {
		oldSC, oldOK = ccr.curState.ServiceConfig.Config.(*ServiceConfig)
	}
	if s.ServiceConfig != nil {
		newSC, newOK = s.ServiceConfig.Config.(*ServiceConfig)
	}
	if oldOK != newOK || (oldOK && newOK && oldSC.rawJSONString != newSC.rawJSONString) {
		updates = append(updates, "service config updated")
	}
	if len(ccr.curState.Addresses) > 0 && len(s.Addresses) == 0 {
		updates = append(updates, "resolver returned an empty address list")
	} else if len(ccr.curState.Addresses) == 0 && len(s.Addresses) > 0 {
		updates = append(updates, "resolver returned new addresses")
	}
	channelz.AddTraceEvent(ccr.cc.channelzID, &channelz.TraceEventDesc{
		Desc:     fmt.Sprintf("Resolver state updated: %+v (%v)", s, strings.Join(updates, "; ")),
		Severity: channelz.CtINFO,
	})
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
