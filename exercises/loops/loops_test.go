package loops

import "testing"

var tests = []struct {
	begin, end int
	expected   int
}{
	{begin: 1, end: 10, expected: -1},
	{begin: 10, end: 50, expected: 35},
	{begin: 123, end: 234, expected: 140},
	{begin: 1900, end: 2000, expected: 1925},
}

func TestFullName(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got, want := compute(tt.begin, tt.end), tt.expected; got != want {
				t.Fatalf("got = %d; want = %d", got, want)
			}
		})
	}
}
