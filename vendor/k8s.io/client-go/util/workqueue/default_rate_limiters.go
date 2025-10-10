/*
Copyright 2016 The Kubernetes Authors.

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

package workqueue

import (
	"math"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Deprecated: RateLimiter is deprecated, use TypedRateLimiter instead.
type RateLimiter TypedRateLimiter[any]

type TypedRateLimiter[T comparable] interface {
	// When gets an item and gets to decide how long that item should wait
<<<<<<< HEAD
	When(item interface{}) time.Duration
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	When(item interface{}) time.Duration
=======
	When(item T) time.Duration
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for failing
	// or for success, we'll stop tracking it
	Forget(item T)
	// NumRequeues returns back how many failures the item has had
	NumRequeues(item T) int
}

// DefaultControllerRateLimiter is a no-arg constructor for a default rate limiter for a workqueue.  It has
// both overall and per-item rate limiting.  The overall is a token bucket and the per-item is exponential
//
// Deprecated: Use DefaultTypedControllerRateLimiter instead.
func DefaultControllerRateLimiter() RateLimiter {
	return DefaultTypedControllerRateLimiter[any]()
}

// DefaultTypedControllerRateLimiter is a no-arg constructor for a default rate limiter for a workqueue.  It has
// both overall and per-item rate limiting.  The overall is a token bucket and the per-item is exponential
func DefaultTypedControllerRateLimiter[T comparable]() TypedRateLimiter[T] {
	return NewTypedMaxOfRateLimiter(
		NewTypedItemExponentialFailureRateLimiter[T](5*time.Millisecond, 1000*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&TypedBucketRateLimiter[T]{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

// Deprecated: BucketRateLimiter is deprecated, use TypedBucketRateLimiter instead.
type BucketRateLimiter = TypedBucketRateLimiter[any]

// TypedBucketRateLimiter adapts a standard bucket to the workqueue ratelimiter API
type TypedBucketRateLimiter[T comparable] struct {
	*rate.Limiter
}

var _ RateLimiter = &BucketRateLimiter{}

func (r *TypedBucketRateLimiter[T]) When(item T) time.Duration {
	return r.Limiter.Reserve().Delay()
}

func (r *TypedBucketRateLimiter[T]) NumRequeues(item T) int {
	return 0
}

func (r *TypedBucketRateLimiter[T]) Forget(item T) {
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// ItemBucketRateLimiter implements a workqueue ratelimiter API using standard rate.Limiter.
// Each key is using a separate limiter.
type ItemBucketRateLimiter struct {
	r     rate.Limit
	burst int

	limitersLock sync.Mutex
	limiters     map[interface{}]*rate.Limiter
}

var _ RateLimiter = &ItemBucketRateLimiter{}

// NewItemBucketRateLimiter creates new ItemBucketRateLimiter instance.
func NewItemBucketRateLimiter(r rate.Limit, burst int) *ItemBucketRateLimiter {
	return &ItemBucketRateLimiter{
		r:        r,
		burst:    burst,
		limiters: make(map[interface{}]*rate.Limiter),
	}
}

// When returns a time.Duration which we need to wait before item is processed.
func (r *ItemBucketRateLimiter) When(item interface{}) time.Duration {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	limiter, ok := r.limiters[item]
	if !ok {
		limiter = rate.NewLimiter(r.r, r.burst)
		r.limiters[item] = limiter
	}

	return limiter.Reserve().Delay()
}

// NumRequeues returns always 0 (doesn't apply to ItemBucketRateLimiter).
func (r *ItemBucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

// Forget removes item from the internal state.
func (r *ItemBucketRateLimiter) Forget(item interface{}) {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	delete(r.limiters, item)
}

>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// ItemBucketRateLimiter implements a workqueue ratelimiter API using standard rate.Limiter.
// Each key is using a separate limiter.
type ItemBucketRateLimiter struct {
	r     rate.Limit
	burst int

	limitersLock sync.Mutex
	limiters     map[interface{}]*rate.Limiter
}

var _ RateLimiter = &ItemBucketRateLimiter{}

// NewItemBucketRateLimiter creates new ItemBucketRateLimiter instance.
func NewItemBucketRateLimiter(r rate.Limit, burst int) *ItemBucketRateLimiter {
	return &ItemBucketRateLimiter{
		r:        r,
		burst:    burst,
		limiters: make(map[interface{}]*rate.Limiter),
	}
}

// When returns a time.Duration which we need to wait before item is processed.
func (r *ItemBucketRateLimiter) When(item interface{}) time.Duration {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	limiter, ok := r.limiters[item]
	if !ok {
		limiter = rate.NewLimiter(r.r, r.burst)
		r.limiters[item] = limiter
	}

	return limiter.Reserve().Delay()
}

// NumRequeues returns always 0 (doesn't apply to ItemBucketRateLimiter).
func (r *ItemBucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

// Forget removes item from the internal state.
func (r *ItemBucketRateLimiter) Forget(item interface{}) {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	delete(r.limiters, item)
}

=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// ItemBucketRateLimiter implements a workqueue ratelimiter API using standard rate.Limiter.
// Each key is using a separate limiter.
type ItemBucketRateLimiter struct {
	r     rate.Limit
	burst int

	limitersLock sync.Mutex
	limiters     map[interface{}]*rate.Limiter
}

var _ RateLimiter = &ItemBucketRateLimiter{}

// NewItemBucketRateLimiter creates new ItemBucketRateLimiter instance.
func NewItemBucketRateLimiter(r rate.Limit, burst int) *ItemBucketRateLimiter {
	return &ItemBucketRateLimiter{
		r:        r,
		burst:    burst,
		limiters: make(map[interface{}]*rate.Limiter),
	}
}

// When returns a time.Duration which we need to wait before item is processed.
func (r *ItemBucketRateLimiter) When(item interface{}) time.Duration {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	limiter, ok := r.limiters[item]
	if !ok {
		limiter = rate.NewLimiter(r.r, r.burst)
		r.limiters[item] = limiter
	}

	return limiter.Reserve().Delay()
}

// NumRequeues returns always 0 (doesn't apply to ItemBucketRateLimiter).
func (r *ItemBucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

// Forget removes item from the internal state.
func (r *ItemBucketRateLimiter) Forget(item interface{}) {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	delete(r.limiters, item)
}

>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// ItemBucketRateLimiter implements a workqueue ratelimiter API using standard rate.Limiter.
// Each key is using a separate limiter.
type ItemBucketRateLimiter struct {
	r     rate.Limit
	burst int

	limitersLock sync.Mutex
	limiters     map[interface{}]*rate.Limiter
}

var _ RateLimiter = &ItemBucketRateLimiter{}

// NewItemBucketRateLimiter creates new ItemBucketRateLimiter instance.
func NewItemBucketRateLimiter(r rate.Limit, burst int) *ItemBucketRateLimiter {
	return &ItemBucketRateLimiter{
		r:        r,
		burst:    burst,
		limiters: make(map[interface{}]*rate.Limiter),
	}
}

// When returns a time.Duration which we need to wait before item is processed.
func (r *ItemBucketRateLimiter) When(item interface{}) time.Duration {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	limiter, ok := r.limiters[item]
	if !ok {
		limiter = rate.NewLimiter(r.r, r.burst)
		r.limiters[item] = limiter
	}

	return limiter.Reserve().Delay()
}

// NumRequeues returns always 0 (doesn't apply to ItemBucketRateLimiter).
func (r *ItemBucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

// Forget removes item from the internal state.
func (r *ItemBucketRateLimiter) Forget(item interface{}) {
	r.limitersLock.Lock()
	defer r.limitersLock.Unlock()

	delete(r.limiters, item)
}

=======
>>>>>>> 6b7ce455e (update vendored files)
// ItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
// dealing with max failures and expiration are up to the caller
type ItemExponentialFailureRateLimiter struct {
	failuresLock sync.Mutex
	failures     map[interface{}]int

	baseDelay time.Duration
	maxDelay  time.Duration
}

var _ RateLimiter = &ItemExponentialFailureRateLimiter{}

func NewItemExponentialFailureRateLimiter(baseDelay time.Duration, maxDelay time.Duration) RateLimiter {
	return &ItemExponentialFailureRateLimiter{
		failures:  map[interface{}]int{},
		baseDelay: baseDelay,
		maxDelay:  maxDelay,
	}
}

func DefaultItemBasedRateLimiter() RateLimiter {
	return NewItemExponentialFailureRateLimiter(time.Millisecond, 1000*time.Second)
}

func (r *ItemExponentialFailureRateLimiter) When(item interface{}) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	exp := r.failures[item]
	r.failures[item] = r.failures[item] + 1

	// The backoff is capped such that 'calculated' value never overflows.
	backoff := float64(r.baseDelay.Nanoseconds()) * math.Pow(2, float64(exp))
	if backoff > math.MaxInt64 {
		return r.maxDelay
	}

	calculated := time.Duration(backoff)
	if calculated > r.maxDelay {
		return r.maxDelay
	}

	return calculated
}

func (r *ItemExponentialFailureRateLimiter) NumRequeues(item interface{}) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *ItemExponentialFailureRateLimiter) Forget(item interface{}) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// ItemFastSlowRateLimiter does a quick retry for a certain number of attempts, then a slow retry after that
type ItemFastSlowRateLimiter struct {
	failuresLock sync.Mutex
	failures     map[interface{}]int

	maxFastAttempts int
	fastDelay       time.Duration
	slowDelay       time.Duration
}

var _ RateLimiter = &ItemFastSlowRateLimiter{}

func NewItemFastSlowRateLimiter(fastDelay, slowDelay time.Duration, maxFastAttempts int) RateLimiter {
	return &ItemFastSlowRateLimiter{
		failures:        map[interface{}]int{},
		fastDelay:       fastDelay,
		slowDelay:       slowDelay,
		maxFastAttempts: maxFastAttempts,
	}
}

func (r *ItemFastSlowRateLimiter) When(item interface{}) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	r.failures[item] = r.failures[item] + 1

	if r.failures[item] <= r.maxFastAttempts {
		return r.fastDelay
	}

	return r.slowDelay
}

func (r *ItemFastSlowRateLimiter) NumRequeues(item interface{}) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *ItemFastSlowRateLimiter) Forget(item interface{}) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// MaxOfRateLimiter calls every RateLimiter and returns the worst case response
// When used with a token bucket limiter, the burst could be apparently exceeded in cases where particular items
// were separately delayed a longer time.
type MaxOfRateLimiter struct {
	limiters []RateLimiter
}

func (r *MaxOfRateLimiter) When(item interface{}) time.Duration {
	ret := time.Duration(0)
	for _, limiter := range r.limiters {
		curr := limiter.When(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

func NewMaxOfRateLimiter(limiters ...RateLimiter) RateLimiter {
	return &MaxOfRateLimiter{limiters: limiters}
}

func (r *MaxOfRateLimiter) NumRequeues(item interface{}) int {
	ret := 0
	for _, limiter := range r.limiters {
		curr := limiter.NumRequeues(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

func (r *MaxOfRateLimiter) Forget(item interface{}) {
	for _, limiter := range r.limiters {
		limiter.Forget(item)
	}
}

// WithMaxWaitRateLimiter have maxDelay which avoids waiting too long
type WithMaxWaitRateLimiter struct {
	limiter  RateLimiter
	maxDelay time.Duration
}

func NewWithMaxWaitRateLimiter(limiter RateLimiter, maxDelay time.Duration) RateLimiter {
	return &WithMaxWaitRateLimiter{limiter: limiter, maxDelay: maxDelay}
}

func (w WithMaxWaitRateLimiter) When(item interface{}) time.Duration {
	delay := w.limiter.When(item)
	if delay > w.maxDelay {
		return w.maxDelay
	}

	return delay
}

func (w WithMaxWaitRateLimiter) Forget(item interface{}) {
	w.limiter.Forget(item)
}

func (w WithMaxWaitRateLimiter) NumRequeues(item interface{}) int {
	return w.limiter.NumRequeues(item)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Forget indicates that an item is finished being retried.  Doesn't matter whether its for perm failing
||||||| parent of 4d7e5ad26 (update vendored files)
	// Forget indicates that an item is finished being retried.  Doesn't matter whether its for perm failing
=======
	// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for failing
>>>>>>> 4d7e5ad26 (update vendored files)
	// or for success, we'll stop tracking it
	Forget(item interface{})
	// NumRequeues returns back how many failures the item has had
	NumRequeues(item interface{}) int
}

// DefaultControllerRateLimiter is a no-arg constructor for a default rate limiter for a workqueue.  It has
// both overall and per-item rate limiting.  The overall is a token bucket and the per-item is exponential
func DefaultControllerRateLimiter() RateLimiter {
	return NewMaxOfRateLimiter(
		NewItemExponentialFailureRateLimiter(5*time.Millisecond, 1000*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

// BucketRateLimiter adapts a standard bucket to the workqueue ratelimiter API
type BucketRateLimiter struct {
	*rate.Limiter
}

var _ RateLimiter = &BucketRateLimiter{}

func (r *BucketRateLimiter) When(item interface{}) time.Duration {
	return r.Limiter.Reserve().Delay()
}

func (r *BucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

func (r *BucketRateLimiter) Forget(item interface{}) {
}

// ItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
// ItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
=======
// Deprecated: ItemExponentialFailureRateLimiter is deprecated, use TypedItemExponentialFailureRateLimiter instead.
type ItemExponentialFailureRateLimiter = TypedItemExponentialFailureRateLimiter[any]

// TypedItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
// dealing with max failures and expiration are up to the caller
type TypedItemExponentialFailureRateLimiter[T comparable] struct {
	failuresLock sync.Mutex
	failures     map[T]int

	baseDelay time.Duration
	maxDelay  time.Duration
}

var _ RateLimiter = &ItemExponentialFailureRateLimiter{}

// Deprecated: NewItemExponentialFailureRateLimiter is deprecated, use NewTypedItemExponentialFailureRateLimiter instead.
func NewItemExponentialFailureRateLimiter(baseDelay time.Duration, maxDelay time.Duration) RateLimiter {
	return NewTypedItemExponentialFailureRateLimiter[any](baseDelay, maxDelay)
}

func NewTypedItemExponentialFailureRateLimiter[T comparable](baseDelay time.Duration, maxDelay time.Duration) TypedRateLimiter[T] {
	return &TypedItemExponentialFailureRateLimiter[T]{
		failures:  map[T]int{},
		baseDelay: baseDelay,
		maxDelay:  maxDelay,
	}
}

// Deprecated: DefaultItemBasedRateLimiter is deprecated, use DefaultTypedItemBasedRateLimiter instead.
func DefaultItemBasedRateLimiter() RateLimiter {
	return DefaultTypedItemBasedRateLimiter[any]()
}

func DefaultTypedItemBasedRateLimiter[T comparable]() TypedRateLimiter[T] {
	return NewTypedItemExponentialFailureRateLimiter[T](time.Millisecond, 1000*time.Second)
}

func (r *TypedItemExponentialFailureRateLimiter[T]) When(item T) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	exp := r.failures[item]
	r.failures[item] = r.failures[item] + 1

	// The backoff is capped such that 'calculated' value never overflows.
	backoff := float64(r.baseDelay.Nanoseconds()) * math.Pow(2, float64(exp))
	if backoff > math.MaxInt64 {
		return r.maxDelay
	}

	calculated := time.Duration(backoff)
	if calculated > r.maxDelay {
		return r.maxDelay
	}

	return calculated
}

func (r *TypedItemExponentialFailureRateLimiter[T]) NumRequeues(item T) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *TypedItemExponentialFailureRateLimiter[T]) Forget(item T) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// ItemFastSlowRateLimiter does a quick retry for a certain number of attempts, then a slow retry after that
// Deprecated: Use TypedItemFastSlowRateLimiter instead.
type ItemFastSlowRateLimiter = TypedItemFastSlowRateLimiter[any]

// TypedItemFastSlowRateLimiter does a quick retry for a certain number of attempts, then a slow retry after that
type TypedItemFastSlowRateLimiter[T comparable] struct {
	failuresLock sync.Mutex
	failures     map[T]int

	maxFastAttempts int
	fastDelay       time.Duration
	slowDelay       time.Duration
}

var _ RateLimiter = &ItemFastSlowRateLimiter{}

// Deprecated: NewItemFastSlowRateLimiter is deprecated, use NewTypedItemFastSlowRateLimiter instead.
func NewItemFastSlowRateLimiter(fastDelay, slowDelay time.Duration, maxFastAttempts int) RateLimiter {
	return NewTypedItemFastSlowRateLimiter[any](fastDelay, slowDelay, maxFastAttempts)
}

func NewTypedItemFastSlowRateLimiter[T comparable](fastDelay, slowDelay time.Duration, maxFastAttempts int) TypedRateLimiter[T] {
	return &TypedItemFastSlowRateLimiter[T]{
		failures:        map[T]int{},
		fastDelay:       fastDelay,
		slowDelay:       slowDelay,
		maxFastAttempts: maxFastAttempts,
	}
}

func (r *TypedItemFastSlowRateLimiter[T]) When(item T) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	r.failures[item] = r.failures[item] + 1

	if r.failures[item] <= r.maxFastAttempts {
		return r.fastDelay
	}

	return r.slowDelay
}

func (r *TypedItemFastSlowRateLimiter[T]) NumRequeues(item T) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *TypedItemFastSlowRateLimiter[T]) Forget(item T) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// MaxOfRateLimiter calls every RateLimiter and returns the worst case response
// When used with a token bucket limiter, the burst could be apparently exceeded in cases where particular items
// were separately delayed a longer time.
//
// Deprecated: Use TypedMaxOfRateLimiter instead.
type MaxOfRateLimiter = TypedMaxOfRateLimiter[any]

// TypedMaxOfRateLimiter calls every RateLimiter and returns the worst case response
// When used with a token bucket limiter, the burst could be apparently exceeded in cases where particular items
// were separately delayed a longer time.
type TypedMaxOfRateLimiter[T comparable] struct {
	limiters []TypedRateLimiter[T]
}

func (r *TypedMaxOfRateLimiter[T]) When(item T) time.Duration {
	ret := time.Duration(0)
	for _, limiter := range r.limiters {
		curr := limiter.When(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

// Deprecated: NewMaxOfRateLimiter is deprecated, use NewTypedMaxOfRateLimiter instead.
func NewMaxOfRateLimiter(limiters ...TypedRateLimiter[any]) RateLimiter {
	return NewTypedMaxOfRateLimiter(limiters...)
}

func NewTypedMaxOfRateLimiter[T comparable](limiters ...TypedRateLimiter[T]) TypedRateLimiter[T] {
	return &TypedMaxOfRateLimiter[T]{limiters: limiters}
}

func (r *TypedMaxOfRateLimiter[T]) NumRequeues(item T) int {
	ret := 0
	for _, limiter := range r.limiters {
		curr := limiter.NumRequeues(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

func (r *TypedMaxOfRateLimiter[T]) Forget(item T) {
	for _, limiter := range r.limiters {
		limiter.Forget(item)
	}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// WithMaxWaitRateLimiter have maxDelay which avoids waiting too long
type WithMaxWaitRateLimiter struct {
	limiter  RateLimiter
	maxDelay time.Duration
}

func NewWithMaxWaitRateLimiter(limiter RateLimiter, maxDelay time.Duration) RateLimiter {
	return &WithMaxWaitRateLimiter{limiter: limiter, maxDelay: maxDelay}
}

func (w WithMaxWaitRateLimiter) When(item interface{}) time.Duration {
	delay := w.limiter.When(item)
	if delay > w.maxDelay {
		return w.maxDelay
	}

	return delay
}

func (w WithMaxWaitRateLimiter) Forget(item interface{}) {
	w.limiter.Forget(item)
}

func (w WithMaxWaitRateLimiter) NumRequeues(item interface{}) int {
	return w.limiter.NumRequeues(item)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Forget indicates that an item is finished being retried.  Doesn't matter whether its for perm failing
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// Forget indicates that an item is finished being retried.  Doesn't matter whether its for perm failing
=======
	// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for failing
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// or for success, we'll stop tracking it
	Forget(item interface{})
	// NumRequeues returns back how many failures the item has had
	NumRequeues(item interface{}) int
}

// DefaultControllerRateLimiter is a no-arg constructor for a default rate limiter for a workqueue.  It has
// both overall and per-item rate limiting.  The overall is a token bucket and the per-item is exponential
func DefaultControllerRateLimiter() RateLimiter {
	return NewMaxOfRateLimiter(
		NewItemExponentialFailureRateLimiter(5*time.Millisecond, 1000*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

// BucketRateLimiter adapts a standard bucket to the workqueue ratelimiter API
type BucketRateLimiter struct {
	*rate.Limiter
}

var _ RateLimiter = &BucketRateLimiter{}

func (r *BucketRateLimiter) When(item interface{}) time.Duration {
	return r.Limiter.Reserve().Delay()
}

func (r *BucketRateLimiter) NumRequeues(item interface{}) int {
	return 0
}

func (r *BucketRateLimiter) Forget(item interface{}) {
}

// ItemExponentialFailureRateLimiter does a simple baseDelay*2^<num-failures> limit
// dealing with max failures and expiration are up to the caller
type ItemExponentialFailureRateLimiter struct {
	failuresLock sync.Mutex
	failures     map[interface{}]int

	baseDelay time.Duration
	maxDelay  time.Duration
}

var _ RateLimiter = &ItemExponentialFailureRateLimiter{}

func NewItemExponentialFailureRateLimiter(baseDelay time.Duration, maxDelay time.Duration) RateLimiter {
	return &ItemExponentialFailureRateLimiter{
		failures:  map[interface{}]int{},
		baseDelay: baseDelay,
		maxDelay:  maxDelay,
	}
}

func DefaultItemBasedRateLimiter() RateLimiter {
	return NewItemExponentialFailureRateLimiter(time.Millisecond, 1000*time.Second)
}

func (r *ItemExponentialFailureRateLimiter) When(item interface{}) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	exp := r.failures[item]
	r.failures[item] = r.failures[item] + 1

	// The backoff is capped such that 'calculated' value never overflows.
	backoff := float64(r.baseDelay.Nanoseconds()) * math.Pow(2, float64(exp))
	if backoff > math.MaxInt64 {
		return r.maxDelay
	}

	calculated := time.Duration(backoff)
	if calculated > r.maxDelay {
		return r.maxDelay
	}

	return calculated
}

func (r *ItemExponentialFailureRateLimiter) NumRequeues(item interface{}) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *ItemExponentialFailureRateLimiter) Forget(item interface{}) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// ItemFastSlowRateLimiter does a quick retry for a certain number of attempts, then a slow retry after that
type ItemFastSlowRateLimiter struct {
	failuresLock sync.Mutex
	failures     map[interface{}]int

	maxFastAttempts int
	fastDelay       time.Duration
	slowDelay       time.Duration
}

var _ RateLimiter = &ItemFastSlowRateLimiter{}

func NewItemFastSlowRateLimiter(fastDelay, slowDelay time.Duration, maxFastAttempts int) RateLimiter {
	return &ItemFastSlowRateLimiter{
		failures:        map[interface{}]int{},
		fastDelay:       fastDelay,
		slowDelay:       slowDelay,
		maxFastAttempts: maxFastAttempts,
	}
}

func (r *ItemFastSlowRateLimiter) When(item interface{}) time.Duration {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	r.failures[item] = r.failures[item] + 1

	if r.failures[item] <= r.maxFastAttempts {
		return r.fastDelay
	}

	return r.slowDelay
}

func (r *ItemFastSlowRateLimiter) NumRequeues(item interface{}) int {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	return r.failures[item]
}

func (r *ItemFastSlowRateLimiter) Forget(item interface{}) {
	r.failuresLock.Lock()
	defer r.failuresLock.Unlock()

	delete(r.failures, item)
}

// MaxOfRateLimiter calls every RateLimiter and returns the worst case response
// When used with a token bucket limiter, the burst could be apparently exceeded in cases where particular items
// were separately delayed a longer time.
type MaxOfRateLimiter struct {
	limiters []RateLimiter
}

func (r *MaxOfRateLimiter) When(item interface{}) time.Duration {
	ret := time.Duration(0)
	for _, limiter := range r.limiters {
		curr := limiter.When(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

func NewMaxOfRateLimiter(limiters ...RateLimiter) RateLimiter {
	return &MaxOfRateLimiter{limiters: limiters}
}

func (r *MaxOfRateLimiter) NumRequeues(item interface{}) int {
	ret := 0
	for _, limiter := range r.limiters {
		curr := limiter.NumRequeues(item)
		if curr > ret {
			ret = curr
		}
	}

	return ret
}

func (r *MaxOfRateLimiter) Forget(item interface{}) {
	for _, limiter := range r.limiters {
		limiter.Forget(item)
	}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// WithMaxWaitRateLimiter have maxDelay which avoids waiting too long
// Deprecated: Use TypedWithMaxWaitRateLimiter instead.
type WithMaxWaitRateLimiter = TypedWithMaxWaitRateLimiter[any]

// TypedWithMaxWaitRateLimiter have maxDelay which avoids waiting too long
type TypedWithMaxWaitRateLimiter[T comparable] struct {
	limiter  TypedRateLimiter[T]
	maxDelay time.Duration
}

// Deprecated: NewWithMaxWaitRateLimiter is deprecated, use NewTypedWithMaxWaitRateLimiter instead.
func NewWithMaxWaitRateLimiter(limiter RateLimiter, maxDelay time.Duration) RateLimiter {
	return NewTypedWithMaxWaitRateLimiter[any](limiter, maxDelay)
}

func NewTypedWithMaxWaitRateLimiter[T comparable](limiter TypedRateLimiter[T], maxDelay time.Duration) TypedRateLimiter[T] {
	return &TypedWithMaxWaitRateLimiter[T]{limiter: limiter, maxDelay: maxDelay}
}

func (w TypedWithMaxWaitRateLimiter[T]) When(item T) time.Duration {
	delay := w.limiter.When(item)
	if delay > w.maxDelay {
		return w.maxDelay
	}

	return delay
}

func (w TypedWithMaxWaitRateLimiter[T]) Forget(item T) {
	w.limiter.Forget(item)
}

func (w TypedWithMaxWaitRateLimiter[T]) NumRequeues(item T) int {
	return w.limiter.NumRequeues(item)
}
