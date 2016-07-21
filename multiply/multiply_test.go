package multiply

import "testing"
// import "fmt"

func TestMultiply(t *testing.T) {
	for _, c := range []struct {
		in1, in2, want string
	}{
		{"10", "10", "100"},
		{"", "", ""},
	} {
		got := multiply(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Multiply(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}
