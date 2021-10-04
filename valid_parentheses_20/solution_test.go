package validparentheses20

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for i, c := range cases {
		res := isValid(c.s)
		if res != c.expected {
			t.Errorf("failed case %v (%v) with %v", i, c.s, res)
		}
	}
}
