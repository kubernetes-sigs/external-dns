/*
Copyright 2014 The Kubernetes Authors.

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

package cache

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/clock"
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"k8s.io/apimachinery/pkg/util/clock"
||||||| parent of 4d7e5ad26 (update vendored files)
	"k8s.io/apimachinery/pkg/util/clock"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	"k8s.io/apimachinery/pkg/util/sets"
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"k8s.io/utils/clock"
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"k8s.io/apimachinery/pkg/util/clock"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"k8s.io/apimachinery/pkg/util/clock"
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"k8s.io/apimachinery/pkg/util/sets"
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"k8s.io/utils/clock"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)

type fakeThreadSafeMap struct {
	ThreadSafeStore
	deletedKeys chan<- string
}

func (c *fakeThreadSafeMap) Delete(key string) {
	if c.deletedKeys != nil {
		c.ThreadSafeStore.Delete(key)
		c.deletedKeys <- key
	}
}

// FakeExpirationPolicy keeps the list for keys which never expires.
type FakeExpirationPolicy struct {
	NeverExpire     sets.String
	RetrieveKeyFunc KeyFunc
}

// IsExpired used to check if object is expired.
func (p *FakeExpirationPolicy) IsExpired(obj *TimestampedEntry) bool {
	key, _ := p.RetrieveKeyFunc(obj)
	return !p.NeverExpire.Has(key)
}

// NewFakeExpirationStore creates a new instance for the ExpirationCache.
func NewFakeExpirationStore(keyFunc KeyFunc, deletedKeys chan<- string, expirationPolicy ExpirationPolicy, cacheClock clock.Clock) Store {
	cacheStorage := NewThreadSafeStore(Indexers{}, Indices{})
	return &ExpirationCache{
		cacheStorage:     &fakeThreadSafeMap{cacheStorage, deletedKeys},
		keyFunc:          keyFunc,
		clock:            cacheClock,
		expirationPolicy: expirationPolicy,
	}
}
