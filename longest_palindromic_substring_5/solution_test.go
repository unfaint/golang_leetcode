package main

import (
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"bb", "bb"},
		{"bbcbbb", "bbcbb"},
		{"asdfbbabbasdf", "bbabb"},
	}

	for i, c := range cases {
		res := longestPalindrome(c.s)

		if res != c.expected {
			t.Errorf("failed example #%v: '%v'", i, c.s)
		}
	}
}
