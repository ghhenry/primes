package primes

import (
	"math/big"
	"testing"
)

func pow10(exp int) *big.Int {
	x := big.NewInt(10)
	y := big.NewInt(1)
	for exp > 0 {
		for exp&1 == 0 {
			exp /= 2
			x.Mul(x, x)
		}
		exp--
		y.Mul(x, y)
	}
	return y
}

func TestFastmod(t *testing.T) {
	type args struct {
		a *big.Int
		m uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "100 % 43",
			args: args{big.NewInt(100), 43},
		},
		{
			name: "10**40 % 43",
			args: args{pow10(40), 43},
		},
		{
			name: "10**40 % max",
			args: args{pow10(40), 0xffffffff},
		},
		{
			name: "10**80 % max",
			args: args{pow10(80), 0xffffffff},
		},
		{
			name: "benchmark",
			args: args{a, mi},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := uint32(new(big.Int).Mod(tt.args.a, big.NewInt(int64(tt.args.m))).Int64())
			if got := Fastmod(tt.args.a, tt.args.m); got != want {
				t.Errorf("Fastmod() = %v, want %v", got, want)
			}
		})
	}
}

var a = pow10(1000)
var mi uint32 = 128347619
var mb = big.NewInt(int64(mi))

func BenchmarkBigMod(b *testing.B) {
	var z = new(big.Int)
	for n := 0; n < b.N; n++ {
		z.Mod(a, mb)
	}
}

func BenchmarkFastMod(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fastmod(a, mi)
	}
}
