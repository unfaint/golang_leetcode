package lengthoflastword58

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
	}

	for i, c := range cases {
		res := lengthOfLastWord(c.input)
		if res != c.output {
			t.Errorf("failed #%v \"%v (%v)\" with result %v", i, c.input, c.output, res)
		}
	}
}
