package primes

import (
	"sync"

	"github.com/ghhenry/intsqrt"
)

// baseSieve has one bit for each odd number below 2^16.
// A set bit indicates the number is not prime.
// From n to bit: n/2
// From bit to n: 2*bit+1
var baseSieve [512]uint64
var initMutex sync.Mutex

func initBase() {
	initMutex.Lock()
	defer initMutex.Unlock()
	if baseSieve[0]&1 != 0 {
		// already initialized
		return
	}
	for bit := 1; bit < 1<<7; bit++ {
		if baseSieve[bit/64]&(1<<uint(bit%64)) != 0 {
			continue
		}
		n := 2*bit + 1
		for i := n * n / 2; i < 512*64; i += n {
			baseSieve[i/64] |= 1 << uint(i%64)
		}
	}
	baseSieve[0] |= 1
}

// Iterate calls consumer with all primes between min and max (inclusive).
func Iterate(min, max uint32, consumer func(p uint32) bool) {
	initBase()
	// 2 is a special case
	if min <= 2 && 2 <= max && consumer(2) {
		return
	}
	// primes from base sieve
	if min < 1<<16 {
		var bit uint32
		bound := max / 2
		if bound >= 1<<15 {
			bound = 1<<15 - 1
		}
		for bit = min / 2; bit <= bound; bit++ {
			if baseSieve[bit/64]&(1<<uint(bit%64)) == 0 && consumer(2*bit+1) {
				return
			}
		}
		min = 1 << 16
	}
	// primes after the base sieve
}

// initMoving initializes a moving sieve
func initMoving(start uint32, sieve []uint64) {
	end := uint32(len(sieve) * 64)
	pmax := intsqrt.Sqrt32(2*(start+end) - 1)
	Iterate(3, pmax, func(p uint32) bool {
		s := p * p / 2
		if s >= start {
			s -= start
		} else {
			s = p - 1 - (start-s-1)%p
		}
		for i := s; i < end; i += p {
			sieve[i/64] |= 1 << uint(i%64)
		}
		return false
	})
}
