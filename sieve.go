package primes

// baseSieve has one bit for each odd number below 2^16.
// A set bit indicates the number is not prime.
// From n to bit: n/2
// From bit to n: 2*bit+1
var baseSieve [512]uint64

func initBase() {
	for bit := 1; bit < 1<<7; bit++ {
		if baseSieve[bit/64]&(1<<uint(bit%64)) != 0 {
			continue
		}
		n := 2*bit + 1
		for i := n * n / 2; i < 512*64; i += n {
			baseSieve[i/64] |= 1 << uint(i%64)
		}
	}
}
