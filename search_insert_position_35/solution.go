package searchinsertposition35

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	left := len(nums)/2 - 1
	right := left + 1
	if target == nums[left] {
		return left
	}
	if target == nums[right] {
		return right
	}

	if nums[left] > target {
		return searchInsert(nums[:left], target)
	} else {
		if nums[right] > target {
			return right
		} else {
			// go right
			return searchInsert(nums[right:], target) + right
		}
	}
}
