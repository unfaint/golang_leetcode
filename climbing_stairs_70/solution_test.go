package climbingstairs70

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for i, c := range cases {
		res := climbStairs(c.n)
		if res != c.expected {
			t.Errorf("failed #%v: %v (%v) with %v", i, c.expected, c.n, res)
		}
	}
}
