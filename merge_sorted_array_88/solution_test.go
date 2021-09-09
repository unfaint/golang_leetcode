package mergesortedarray88

import "testing"

func compareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}
	for i, c := range cases {
		before := make([]int, len(c.nums1))
		copy(before, c.nums1)
		merge(c.nums1, c.m, c.nums2, c.n)
		if !compareSlices(c.nums1, c.expected) {
			t.Errorf("failed %v %v+%v with %v", i, before, c.nums2, c.nums1)
		}
	}
}
