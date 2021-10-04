package validparentheses20

func isValid(s string) bool {
	opening := map[rune]bool{
		rune('('): true,
		rune('{'): true,
		rune('['): true,
	}
	closing := map[rune]rune{
		rune(')'): rune('('),
		rune('}'): rune('{'),
		rune(']'): rune('['),
	}
	var lifo []rune
	for _, r := range s {
		_, ok := opening[r]
		if ok {
			lifo = append(lifo, r)
		} else {
			if len(lifo) == 0 {
				return false
			}
			op, ok := closing[r]
			if ok {
				if op != lifo[len(lifo)-1] {
					return false
				} else {
					lifo = lifo[:len(lifo)-1]
				}
			}
		}
	}
	return true
}
