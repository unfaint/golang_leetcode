package climbingstairs70

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	x, y := 2, 3
	for i := 4; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
