package primes

// Stream creates a goroutine that sends the primes in the interval [min,max] over the channel it returns.
// The stream can be stopped early by closing the done channel.
func Stream(min, max uint32, done <-chan struct{}) <-chan uint32 {
	out := make(chan uint32)
	go func() {
		Iterate(min, max, func(p uint32) bool {
			select {
			// done should have priority
			case <-done:
				return true
			default:
			}
			select {
			case <-done:
				return true
			case out <- p:
				return false
			}
		})
		close(out)
	}()
	return out
}
