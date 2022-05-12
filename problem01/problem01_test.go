package problem01_test

import (
	"halan-assignment/problem01"
	"testing"
)

func TestIsPalindrom(t *testing.T) {
	// Test table
	tt := []struct {
		str   string
		state bool
	}{
		{"anna", true},
		{"apple", false},
		{"a", true},
		{"ab", false},
		{"", true},
		{"lala", false},
	}

	for _, tc := range tt {
		output := problem01.IsPalindrom(tc.str)
		if output != tc.state {
			t.Errorf("Expected %v but instead got %v on string: (\"%v\")", tc.state, output, tc.str)
		}
	}
}
