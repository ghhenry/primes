package primes

import "context"

// Stream creates a goroutine that sends the primes in the interval [min,max] over the channel it returns.
// The stream can be stopped early by cancelling the context.
func Stream(ctx context.Context, min, max uint32) <-chan uint32 {
	out := make(chan uint32)
	go func() {
		Iterate(min, max, func(p uint32) bool {
			select {
			// done should have priority
			case <-ctx.Done():
				return true
			default:
			}
			select {
			case <-ctx.Done():
				return true
			case out <- p:
				return false
			}
		})
		close(out)
	}()
	return out
}
