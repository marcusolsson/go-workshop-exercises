package slices

import "testing"

func TestFindMaximum(t *testing.T) {
	if got, want := maximum([]int{6, 2, 56, 1, 64, 32, 7}), 64; got != want {
		t.Fatalf("got = %d; want = %d", got, want)
	}
}
