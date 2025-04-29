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
			input:    "hello    world",
			expected: []string{"hello", "world"},
		}, {
			input:    "dora la    exploradora",
			expected: []string{"dora", "la", "exploradora"},
		}, {
			input:    "    bad to the       bone",
			expected: []string{"bad", "to", "the", "bone"},
		},
	}
	for _, v := range cases {
		actual := cleanInput(v.input)
		if len(actual) != len(v.expected) {
			t.Errorf("Not matching length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := v.expected[i]
			if word != expectedWord {
				t.Errorf("Not matching words")
			}
		}
	}
}

/*func TestCommandExplore(t *testing.T) {
	cases := []struct {
		input    string
		expected error
	}{
		{
			input:    "https://pokeapi.co/api/v2/location-area/mt-coronet-3f",
			expected: nil,
		},
	}
}*/
