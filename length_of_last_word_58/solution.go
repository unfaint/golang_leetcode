package lengthoflastword58

func lengthOfLastWord(s string) int {
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res++
		} else {
			if res > 0 {
				break
			}
		}
	}
	return res
}
