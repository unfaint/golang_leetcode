package main

import "testing"

func TestEmptyNeedle(t *testing.T) {
	haystack := "abc"
	needle := ""
	expected := 0

	res := strStr(haystack, needle)

	if res != expected {
		t.Errorf("failed empty needle")
	}
}

func TestExamples(t *testing.T) {
	cases := []struct {
		haystack, needle string
		expected         int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
	}

	for i, c := range cases {
		res := strStr(c.haystack, c.needle)

		if res != c.expected {
			t.Errorf("failed example %q", i+1)
		}
	}
}

func TestExampleMyCases(t *testing.T) {
	cases := []struct {
		haystack, needle string
		expected         int
	}{
		{".ab.abc..", "abc", 4},
		{"ab", "abc", -1},
		{"", "", 0},
		{"mississippi", "issip", 4},
		{"mississippi", "pi", 9},
	}

	for i, c := range cases {
		res := strStr(c.haystack, c.needle)

		if res != c.expected {
			t.Errorf("failed my example %v", i+1)
		}
	}
}
