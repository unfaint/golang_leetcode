package maximumsubarray53

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > curSum {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}

	return maxSum
}
