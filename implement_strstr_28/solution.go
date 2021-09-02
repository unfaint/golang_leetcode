package main

import "fmt"

func strStr(haystack string, needle string) int {
	switch {
	case len(needle) == 0:
		return 0
	case len(needle) > len(haystack):
		return -1
	case haystack == needle:
		return 0
	}

	hbytes := []byte(haystack)
	nbytes := []byte(needle)

	for i, j := 0, 0; i < len(hbytes); i++ {
		if hbytes[i] != nbytes[j] {
			i -= j
			j = 0
		} else {
			if j == len(nbytes)-1 {
				return i - j
			}
			j++
		}
	}
	return -1
}

func main() {
	fmt.Println("------------")
	fmt.Println(strStr("hello", "ll"))
}
