<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
# Go rate limiter [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] ![test][test-img]

This package provides a Golang implementation of the leaky-bucket rate limit algorithm.
This implementation refills the bucket based on the time elapsed between
requests instead of requiring an interval clock to fill the bucket discretely.

Create a rate limiter with a maximum number of operations to perform per second.
Call Take() before each operation. Take will sleep until you can continue.

```go
import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
    rl := ratelimit.New(100) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}
```

[cov-img]: https://codecov.io/gh/uber-go/ratelimit/branch/master/graph/badge.svg?token=zhLeUjjrm2
[cov]: https://codecov.io/gh/uber-go/ratelimit
[doc-img]: https://pkg.go.dev/badge/go.uber.org/ratelimit
[doc]: https://pkg.go.dev/go.uber.org/ratelimit
[test-img]: https://github.com/uber-go/ratelimit/workflows/test/badge.svg
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
# Go rate limiter
||||||| parent of 5ce8c7613 (update vendored files)
# Go rate limiter
=======
# Go rate limiter [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] ![test][test-img]
>>>>>>> 5ce8c7613 (update vendored files)

This package provides a Golang implementation of the leaky-bucket rate limit algorithm.
This implementation refills the bucket based on the time elapsed between
requests instead of requiring an interval clock to fill the bucket discretely.

Create a rate limiter with a maximum number of operations to perform per second.
Call Take() before each operation. Take will sleep until you can continue.

```go
import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
    rl := ratelimit.New(100) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}
```
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======

[cov-img]: https://codecov.io/gh/uber-go/ratelimit/branch/master/graph/badge.svg?token=zhLeUjjrm2
[cov]: https://codecov.io/gh/uber-go/ratelimit
[doc-img]: https://pkg.go.dev/badge/go.uber.org/ratelimit
[doc]: https://pkg.go.dev/go.uber.org/ratelimit
[test-img]: https://github.com/uber-go/ratelimit/workflows/test/badge.svg
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
# Go rate limiter
||||||| parent of 6b7ce455e (update vendored files)
# Go rate limiter
=======
# Go rate limiter [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] ![test][test-img]
>>>>>>> 6b7ce455e (update vendored files)

This package provides a Golang implementation of the leaky-bucket rate limit algorithm.
This implementation refills the bucket based on the time elapsed between
requests instead of requiring an interval clock to fill the bucket discretely.

Create a rate limiter with a maximum number of operations to perform per second.
Call Take() before each operation. Take will sleep until you can continue.

```go
import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
    rl := ratelimit.New(100) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}
```
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======

[cov-img]: https://codecov.io/gh/uber-go/ratelimit/branch/master/graph/badge.svg?token=zhLeUjjrm2
[cov]: https://codecov.io/gh/uber-go/ratelimit
[doc-img]: https://pkg.go.dev/badge/go.uber.org/ratelimit
[doc]: https://pkg.go.dev/go.uber.org/ratelimit
[test-img]: https://github.com/uber-go/ratelimit/workflows/test/badge.svg
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
# Go rate limiter
||||||| parent of 4d7e5ad26 (update vendored files)
# Go rate limiter
=======
# Go rate limiter [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] ![test][test-img]
>>>>>>> 4d7e5ad26 (update vendored files)

This package provides a Golang implementation of the leaky-bucket rate limit algorithm.
This implementation refills the bucket based on the time elapsed between
requests instead of requiring an interval clock to fill the bucket discretely.

Create a rate limiter with a maximum number of operations to perform per second.
Call Take() before each operation. Take will sleep until you can continue.

```go
import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
    rl := ratelimit.New(100) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}
```
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

[cov-img]: https://codecov.io/gh/uber-go/ratelimit/branch/master/graph/badge.svg?token=zhLeUjjrm2
[cov]: https://codecov.io/gh/uber-go/ratelimit
[doc-img]: https://pkg.go.dev/badge/go.uber.org/ratelimit
[doc]: https://pkg.go.dev/go.uber.org/ratelimit
[test-img]: https://github.com/uber-go/ratelimit/workflows/test/badge.svg
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
# Go rate limiter
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
# Go rate limiter
=======
# Go rate limiter [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] ![test][test-img]
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

This package provides a Golang implementation of the leaky-bucket rate limit algorithm.
This implementation refills the bucket based on the time elapsed between
requests instead of requiring an interval clock to fill the bucket discretely.

Create a rate limiter with a maximum number of operations to perform per second.
Call Take() before each operation. Take will sleep until you can continue.

```go
import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
    rl := ratelimit.New(100) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take()
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}
```
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

## FAQ:
- What's the major diff v.s. https://pkg.go.dev/golang.org/x/time/rate? (based on #77)

  This ratelimiter was meant to have a (1) simple API and (2) minimal overhead. For more complex use-cases [x/time/rate] is a great choice.  See [here][redit] for historical context, and [here][bench] for benchmarks (from 2016).

- Why does example_test.go fail when I run it locally on Windows? (based on #80)

  Windows has some known issues with timers precision. See golang/go#44343. We don't expect to work around it.

[cov-img]: https://codecov.io/gh/uber-go/ratelimit/branch/master/graph/badge.svg?token=zhLeUjjrm2
[cov]: https://codecov.io/gh/uber-go/ratelimit
[doc-img]: https://pkg.go.dev/badge/go.uber.org/ratelimit
[doc]: https://pkg.go.dev/go.uber.org/ratelimit
[test-img]: https://github.com/uber-go/ratelimit/workflows/test/badge.svg
[redit]: https://www.reddit.com/r/golang/comments/59k2bi/ubergoratelimit_a_golang_blocking_leakybucket/d99ob9q
[x/time/rate]: https://pkg.go.dev/golang.org/x/time/rate
[bench]: https://gist.github.com/prashantv/26016a7dbc6fc1ec52d8c2b6591f3582
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
