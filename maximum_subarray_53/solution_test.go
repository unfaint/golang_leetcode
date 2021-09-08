package maximumsubarray53

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{1, 2}, 3},
	}

	for i, c := range cases {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		res := maxSubArray(nums)
		if res != c.expected {
			t.Errorf("case #%v %v (%v) failed with result %v", i, nums, c.expected, res)
		} else {
			t.Logf("case #%v %v (%v) ok", i, nums, c.expected)
		}
	}
}
