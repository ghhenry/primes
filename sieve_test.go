package primes

import (
	"testing"

	"github.com/ghhenry/intsqrt"
)

// slow but safe primality test
func isPrime(n uint32) bool {
	if n&1 == 0 {
		return false
	}
	lim := intsqrt.Sqrt32(n)
	var i uint32
	for i = 3; i <= lim; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestInitBase(t *testing.T) {
	initBase()
	for n := uint32(3); n < 1<<16; n += 2 {
		bit := n / 2
		if isPrime(n) {
			if baseSieve[bit/64]&(1<<uint(bit%64)) != 0 {
				t.Errorf("%v is prime but indicated as composite", n)
			}
		} else {
			if baseSieve[bit/64]&(1<<uint(bit%64)) == 0 {
				t.Errorf("%v is composite but indicated as prime", n)
			}
		}
	}
}
