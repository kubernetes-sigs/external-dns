/*
 *
 * Copyright 2020 gRPC authors.
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

// Package networktype declares the network type to be used in the default
// dialer. Attribute of a resolver.Address.
package networktype

import (
	"google.golang.org/grpc/resolver"
)

// keyType is the key to use for storing State in Attributes.
type keyType string

const key = keyType("grpc.internal.transport.networktype")

// Set returns a copy of the provided address with attributes containing networkType.
func Set(address resolver.Address, networkType string) resolver.Address {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	address.Attributes = address.Attributes.WithValues(key, networkType)
||||||| parent of 6b7ce455e (update vendored files)
=======
	address.Attributes = address.Attributes.WithValue(key, networkType)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	address.Attributes = address.Attributes.WithValue(key, networkType)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	address.Attributes = address.Attributes.WithValue(key, networkType)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return address
}

// Get returns the network type in the resolver.Address and true, or "", false
// if not present.
func Get(address resolver.Address) (string, bool) {
	v := address.Attributes.Value(key)
	if v == nil {
		return "", false
	}
	return v.(string), true
}
