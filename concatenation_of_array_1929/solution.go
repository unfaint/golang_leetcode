package concatenationofarray1929

func getConcatenation(nums []int) []int {
	res := make([]int, 2*len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = nums[i%len(nums)]
	}
	return res
}
