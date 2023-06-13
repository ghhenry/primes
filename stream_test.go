package primes

import (
	"context"
	"testing"
)

func expectNumber(t *testing.T, pc <-chan uint32, n uint32) {
	p, ok := <-pc
	if !ok {
		t.Error("unexpected channel close, want", n)
	}
	if p != n {
		t.Error("got", p, "but want", n)
	}
}

func expectClose(t *testing.T, pc <-chan uint32) {
	if p, ok := <-pc; ok {
		t.Error("got", p, "but expected channel close")
	}
}

func TestCompleteRun(t *testing.T) {
	pc := Stream(context.Background(), 2, 10)
	expectNumber(t, pc, 2)
	expectNumber(t, pc, 3)
	expectNumber(t, pc, 5)
	expectNumber(t, pc, 7)
	expectClose(t, pc)
}

func TestEarlyStop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	pc := Stream(ctx, 2, 10)
	expectNumber(t, pc, 2)
	expectNumber(t, pc, 3)
	cancel()
	expectClose(t, pc)
}
