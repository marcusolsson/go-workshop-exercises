package ids

import "testing"

func TestGenerateIDs(t *testing.T) {
	for _, tt := range []struct {
		In  string
		Out string
	}{
		{In: "A New Hope", Out: "a-new-hope"},
		{In: "Empire Strikes Back", Out: "empire-strikes-back"},
		{In: "Return of the Jedi", Out: "return-of-the-jedi"},
	} {
		if got := GenerateID(tt.In); got != tt.Out {
			t.Errorf("got = %s; want = %s", got, tt.Out)
		}
	}
}
