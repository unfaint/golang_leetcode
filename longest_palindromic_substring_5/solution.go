package main

import (
	"fmt"
)

func makeMap(s string) map[byte]map[int]bool {
	b := []byte(s)
	res := make(map[byte]map[int]bool)

	for i := range b {
		_, ok := res[b[i]]
		if !ok {
			res[b[i]] = make(map[int]bool)
		}
		res[b[i]][i] = true
	}

	return res
}

func longestPalindrome(s string) string {
	m := makeMap(s)
	b := []byte(s)
	var longest string
	n := 0

	for i := 0; i < len(b); i++ {
		for k := range m[b[i]] {
			switch {
			case k-i == 1:
				if k-i > n {
					longest = string(b[i : k+1])
					n = k - i
				}
			case k > i:
				success := true
				middle := i + (k-i)/2
				for ii, kk := i+1, k-1; ii <= middle; ii, kk = ii+1, kk-1 {
					if b[ii] != b[kk] {
						success = false
					}
				}
				if success {
					if k-i > n {
						longest = string(b[i : k+1])
						n = k - i
					}
				}
			}
		}
	}
	if n == 0 {
		longest = string(b[0])
	}

	return longest
}

func main() {
	s := "bbcbbb"
	res := longestPalindrome(s)
	fmt.Printf("%v %v \n", s, res)
}
