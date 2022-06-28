package primes

import "testing"

func expectNumber(t *testing.T, pc <-chan uint32, n uint32) {
	if p, ok := <-pc; ok {
		if p != n {
			t.Error("got", p, "but want", n)
		}
		return
	}
	t.Error("unexpected channel close, want", n)
}

func expectClose(t *testing.T, pc <-chan uint32) {
	if p, ok := <-pc; ok {
		t.Error("got", p, "but expected channel close")
	}
}

func TestCompleteRun(t *testing.T) {
	pc := Stream(2, 10, nil)
	expectNumber(t, pc, 2)
	expectNumber(t, pc, 3)
	expectNumber(t, pc, 5)
	expectNumber(t, pc, 7)
	expectClose(t, pc)
}

func TestEarlyStop(t *testing.T) {
	done := make(chan struct{})
	pc := Stream(2, 10, done)
	expectNumber(t, pc, 2)
	expectNumber(t, pc, 3)
	close(done)
	expectClose(t, pc)
}
