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

// Package base defines a balancer base that can be used to build balancers with
// different picking algorithms.
//
// The base balancer creates a new SubConn for each resolved address. The
// provided picker will only be notified about READY SubConns.
//
// This package is the base of round_robin balancer, its purpose is to be used
// to build round_robin like balancers with complex picking algorithms.
// Balancers with more complicated logic should try to implement a balancer
// builder from scratch.
//
// All APIs in this package are experimental.
package base

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

// PickerBuilder creates balancer.Picker.
type PickerBuilder interface {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
// balancer.
type SubConnInfo struct {
	Address resolver.Address // the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:          name,
		pickerBuilder: pb,
		config:        config,
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
||||||| parent of 5ce8c7613 (update vendored files)
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
=======
>>>>>>> 5ce8c7613 (update vendored files)
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
// balancer.
type SubConnInfo struct {
	Address resolver.Address // the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:          name,
		pickerBuilder: pb,
		config:        config,
	}
}
<<<<<<< HEAD

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}
}
||||||| parent of 5ce8c7613 (update vendored files)

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
||||||| parent of 6b7ce455e (update vendored files)
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
=======
>>>>>>> 6b7ce455e (update vendored files)
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
// balancer.
type SubConnInfo struct {
	Address resolver.Address // the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:          name,
		pickerBuilder: pb,
		config:        config,
	}
}
<<<<<<< HEAD

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
||||||| parent of 4d7e5ad26 (update vendored files)
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
// balancer.
type SubConnInfo struct {
	Address resolver.Address // the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:          name,
		pickerBuilder: pb,
		config:        config,
	}
}
<<<<<<< HEAD

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// Build takes a slice of ready SubConns, and returns a picker that will be
	// used by gRPC to pick a SubConn.
	Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker
}

// V2PickerBuilder creates balancer.V2Picker.
type V2PickerBuilder interface {
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) balancer.Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[balancer.SubConn]SubConnInfo
}

// SubConnInfo contains information about a SubConn created by the base
// balancer.
type SubConnInfo struct {
	Address resolver.Address // the address used to create this SubConn
}

// Config contains the config info about the base balancer builder.
type Config struct {
	// HealthCheck indicates whether health checking should be enabled for this specific balancer.
	HealthCheck bool
}

// NewBalancerBuilder returns a base balancer builder configured by the provided config.
func NewBalancerBuilder(name string, pb PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:          name,
		pickerBuilder: pb,
		config:        config,
	}
}
<<<<<<< HEAD

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

// NewBalancerBuilderV2 returns a base balancer builder configured by the provided config.
func NewBalancerBuilderV2(name string, pb V2PickerBuilder, config Config) balancer.Builder {
	return &baseBuilder{
		name:            name,
		v2PickerBuilder: pb,
		config:          config,
	}
}
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
