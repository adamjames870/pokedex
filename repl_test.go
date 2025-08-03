package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "adam",
			expected: []string{"adam"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "3 WORDS   or MoRe  than 2 ",
			expected: []string{"3", "words", "or", "more", "than", "2"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %v words, got %v", len(c.expected), len(actual))
			t.Fail()
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word [%v] expected %v, got %v", i, expectedWord, word)
				t.Fail()
			}
		}
	}

}
