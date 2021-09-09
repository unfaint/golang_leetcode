package mergesortedarray88

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, len(nums1))
	copy(tmp, nums1)

	for i, im, in := 0, 0, 0; i < len(tmp); i++ {
		if (im < m) && (in < n) {
			if tmp[im] <= nums2[in] {
				nums1[i] = tmp[im]
				im++
			} else {
				nums1[i] = nums2[in]
				in++
			}
		} else {
			if im == m {
				nums1[i] = nums2[in]
				in++
			} else {
				nums1[i] = tmp[im]
				im++
			}
		}
	}
}
