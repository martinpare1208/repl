package main

import (
	"testing"
)

func TestInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " ",
			expected: []string{},
		},
		{
			input: " input  ",
			expected: []string{"input"},
		},
		{
			input: " input1 input2 ",
			expected: []string{"input1","input2"},
		},
		{
			input: "henlO",
			expected: []string{"henlo"},
		},

	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: %v vs %v", actual, c.expected)
			continue
		}
		for i, _ := range actual {
			letter := actual[i]
			expectedLetter := c.expected[i]
			if letter != expectedLetter {
				t.Errorf("letters don't match %v != %v", actual, c.expected )
			}
		}
	}
}