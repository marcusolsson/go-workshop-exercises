package division

import "testing"

func TestDivisionByOneShouldReturnDividend(t *testing.T) {
	got, err := Divide(7.0, 1.0)
	if err != nil {
		t.Fatal("found an error")
	}

	want := 7.0

	if got != want {
		t.Fatalf("got = %f; want = %f", got, want)
	}
}

func TestDivisionByZeroShouldReturnError(t *testing.T) {
	_, err := Divide(7.0, 0.0)
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	// go test -v
	t.Logf("received error: %s", err)
}
