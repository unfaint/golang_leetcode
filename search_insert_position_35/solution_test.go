package searchinsertposition35

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		nums             []int
		needle, expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
		{[]int{}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, 7},
		{[]int{1, 2, 3, 4, 6, 7}, 5, 4},
	}

	for _, c := range cases {
		res := searchInsert(c.nums, c.needle)
		if res != c.expected {
			t.Errorf("%v (%v) failed - expected %v, got %v", c.nums, c.needle, c.expected, res)
		}
	}
}
