package function

import "testing"

func TestFullName(t *testing.T) {
	if got, want := fullName("Jane", "Doe"), "Jane Doe"; got != want {
		t.Fatalf("got = %s; want = %s", got, want)
	}
}
