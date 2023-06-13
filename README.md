# package primes

This repository contains a generator for small primes.
The primes must fit into an uint32.

## Usage

### Synchronous

```go
package primes // import "github.com/ghhenry/primes"

func Iterate(min, max uint32, consumer func(p uint32) bool)
```

The consumer is called with all primes in the range `[min, max]` (borders inclusive).
The consumer may stop the iteration early by returning `true`.

### Via channel

```go
package primes // import "github.com/ghhenry/primes"
import "context"

func Stream(ctx context.Context, min, max uint32) <-chan uint32
```
Stream creates a goroutine that sends the primes in the interval `[min,max]` over the channel it returns.
The stream can be stopped early by cancelling the context `ctx`.
