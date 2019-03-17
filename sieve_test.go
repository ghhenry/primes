package primes

import (
	"testing"

	"github.com/ghhenry/intsqrt"
)

// slow but safe primality test
func isPrime(n uint32) bool {
	if n&1 == 0 {
		return n == 2
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

func TestInitMoving(t *testing.T) {
	var tests = []struct {
		start uint32
	}{
		{0},
		{1},
		{500},
		{1 << 16},
		{1 << 20},
		{1 << 24},
		{1 << 28},
		{1<<31 - 128},
		{0x7fffff06},
	}
	for _, test := range tests {
		s := make([]uint64, 2)
		initMoving(test.start, s)
		for i := uint32(0); i < 128; i++ {
			n := 2*(i+test.start) + 1
			if isPrime(n) {
				if s[i/64]&(1<<uint(i%64)) != 0 {
					t.Errorf("%v is prime but indicated as composite", n)
				}
			} else {
				if s[i/64]&(1<<uint(i%64)) == 0 {
					t.Errorf("%v is composite but indicated as prime", n)
				}
			}
		}
	}
}

func TestIterateSmall(t *testing.T) {
	var tests = []struct {
		min, max, count uint32
	}{
		{0, 1000, 168},
		{1001, 2000, 135},
		{2001, 5000, 366},
		{10000, 1 << 16, 5313},
		{20000, 1 << 17, 9989},
		{1<<28 - 500, 1<<28 + 500, 53},
		{1<<32 - 500, 1<<32 - 1, 18},
		{65000, 70000, 442},
		{65000, 200000, 11491},
		{65000, 1000000, 72005},
		//{0, 10000000, 664579},
	}
	for _, test := range tests {
		count := uint32(0)
		Iterate(test.min, test.max, func(p uint32) bool {
			if !isPrime(p) {
				t.Errorf("%v is composite but indicated as prime", p)
			}
			count++
			return false
		})
		if count != test.count {
			t.Errorf("found %v primes, but there should be %v", count, test.count)
		}
	}
}

func BenchmarkSieve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var count int
		Iterate(0, 10000000, func(p uint32) bool {
			count++
			return false
		})
		if count != 664579 {
			b.Errorf("wrong count %v, want 664579", count)
		}
	}
}
