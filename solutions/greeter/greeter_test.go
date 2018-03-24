package greeter

import "testing"

var tests = []struct {
	name     string
	greeting string
}{
	{name: "Jane Doe", greeting: "Hello, Jane Doe!"},
	{name: "Patrick Swayze", greeting: "Hello, Patrick Swayze!"},
}

func TestGreeter(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := greeter{
				name: tt.name,
			}
			if got, want := g.greet(), tt.greeting; got != want {
				t.Fatalf("got = %s; want = %s", got, want)
			}
		})
	}
}
