package primes

import "math/big"

// wordPower is the smallest x such that 2 to the power x does not fit into an uint
const wordPower = int(5 + (^uint(0) >> 32 & 1))

func Fastmod(a *big.Int, m uint32) uint32 {
	var b uint64 = 2
	for i := 0; i < wordPower; i++ {
		b = b * b % uint64(m)
	}
	w := a.Bits()
	var r uint64 = 0
	for i := len(w) - 1; i >= 0; i-- {
		// r*b can be at most 2**64 - 2**34 + 4, so the sum cannot overflow
		r = (r*b + uint64(w[i])%uint64(m)) % uint64(m)
	}
	return uint32(r)
}
