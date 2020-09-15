package models

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_NewCharacterMapSlice(t *testing.T) {
	testCases := []struct {
		word     string
		expected map[string]int
	}{
		{
			word:     "This",
			expected: map[string]int{"t": 1, "h": 1, "i": 1, "s": 1},
		},
		{
			word:     "is",
			expected: map[string]int{"i": 1, "s": 1},
		},
		{
			word:     "Sparta",
			expected: map[string]int{"s": 1, "p": 1, "a": 2, "t": 1, "r": 1},
		},
		{
			word:     "banana",
			expected: map[string]int{"b": 1, "a": 3, "n": 2},
		},
		{
			word:     "foo",
			expected: map[string]int{"f": 1, "o": 2},
		},
	}

	for _, tc := range testCases {
		cMapSlice := NewCharacterMapSlice(tc.word)
		for _, d := range cMapSlice {
			assert.Equal(t, d.Value, tc.expected[d.Key])
		}
	}
}

func Test_SortByValue(t *testing.T) {
	testCases := []struct {
		word     string
		expected CharacterMapSlice
	}{
		{
			word: "foo",
			expected: CharacterMapSlice{KeyValuePair{Key: "f", Value: 1},
				KeyValuePair{Key: "o", Value: 2}},
		},
		{
			word: "affoooo",
			expected: CharacterMapSlice{KeyValuePair{Key: "a", Value: 1},
				KeyValuePair{Key: "f", Value: 2},
				KeyValuePair{Key: "o", Value: 4}},
		},
		{
			word: "qwweeerrrrtttttyyyyyy",
			expected: CharacterMapSlice{KeyValuePair{Key: "q", Value: 1},
				KeyValuePair{Key: "w", Value: 2},
				KeyValuePair{Key: "e", Value: 3},
				KeyValuePair{Key: "r", Value: 4},
				KeyValuePair{Key: "t", Value: 5},
				KeyValuePair{Key: "y", Value: 6}},
		},
	}

	for _, tc := range testCases {
		cMapSlice := NewCharacterMapSlice(tc.word)
		cMapSlice.SortByValue()
		assert.Equal(t, cMapSlice, tc.expected)
	}
}
