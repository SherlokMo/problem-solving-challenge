package problem03_test

import (
	"halan-assignment/problem03"
	"testing"
)

func TestRunLengthEncode(t *testing.T) {
	tt := []struct {
		input  string
		target string
	}{
		{"aaaaaaaaaabbbaxxxxyyyzyx", "a10b3a1x4y3z1y1x1"},
		{"a", "a1"},
		{"aabbab", "a2b2a1b1"},
		{"", ""},
	}

	for _, tc := range tt {
		output := problem03.RunLengthEncode(tc.input)
		if output != tc.target {
			t.Errorf("Expected \"%v\" but got \"%v\". Input: %v", tc.target, output, tc.input)
		}
	}
}

func TestRunLengthDecode(t *testing.T) {
	tt := []struct {
		input  string
		target string
	}{
		{"a10b3a1x4y3z1y1x1", "aaaaaaaaaabbbaxxxxyyyzyx"},
		{"a1", "a"},
		{"a2b2a1b1", "aabbab"},
		{"", ""},
		{"a1b1c2d3", "abccddd"},
	}

	for _, tc := range tt {
		output := problem03.RunLengthDecode(tc.input)
		if output != tc.target {
			t.Errorf("Expected \"%v\" but got \"%v\". Input: %v", tc.target, output, tc.input)
		}
	}

}
