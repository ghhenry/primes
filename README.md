# package primes

This repository contains a generator for small primes.
The primes must fit into an uint32.

## Usage

    package primes // import "github.com/ghhenry/primes"

    func Iterate(min, max uint32, consumer func(p uint32) bool)

The consumer is called with all primes in the range [min, max] (borders inclusive).
It may stop the iteration early by returning true.
