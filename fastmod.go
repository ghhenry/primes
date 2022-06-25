package primes

import "math/big"

// wordPower is the smallest x such that 2 squared x times does not fit into an uint.
const wordPower = int(5 + (^uint(0) >> 32 & 1))

// Fastmod calculates |a| % m.
// This is faster than the big.Int modulo function.
func Fastmod(a *big.Int, m uint32) uint32 {
	var b uint64 = 2
	mb := uint64(m)
	for i := 0; i < wordPower; i++ {
		b = b * b % mb
	}
	w := a.Bits()
	var r uint64 = 0
	for i := len(w) - 1; i >= 0; i-- {
		// r*b can be at most 2**64 - 2**34 + 4, so the sum cannot overflow
		r = (r*b + uint64(w[i])%mb) % mb
	}
	return uint32(r)
}
