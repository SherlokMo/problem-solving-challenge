package problem08_test

import (
	"halan-assignment/problem08"
	"testing"
)

func TestIndexOfFirstDuplicate(t *testing.T) {
	tt := []struct {
		input  []int
		target int
	}{
		{[]int{5, 17, 3, 17, 4, -1}, 3},
		{[]int{5, 5, 2, 6}, 1},
		{[]int{-1, 2, 3, 4, 3}, 4},
		{[]int{}, -1},
		{[]int{1}, -1},
	}

	for _, tc := range tt {
		output := problem08.IndexOfFirstDuplicate(tc.input)
		if output != tc.target {
			t.Errorf("Expected (%v) but got (%v). input: %v", tc.target, output, tc.input)
		}
	}
}
