package concatenationofarray1929

import "testing"

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSolution(t *testing.T) {
	cases := []struct {
		array    []int
		expected []int
	}{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}
	for i, c := range cases {
		res := getConcatenation(c.array)
		if !compareSlices(res, c.expected) {
			t.Errorf("failed case #%v %v with result %v", i, c.array, res)
		}
	}
}
