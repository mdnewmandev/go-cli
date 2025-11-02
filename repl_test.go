package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:		"  hello world  ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		"    Hello BIG FAT STUPID WOrld",
			expected:	[]string{"hello", "big", "fat", "stupid", "world"},
		},
		{
			input:		"Charmander Bulbasaur PIKACHU", 
			expected:	[]string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] == %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}