package problem06_test

import (
	"halan-assignment/problem06"
	"testing"
)

func TestTranspose(t *testing.T) {
	tt := []struct {
		input  [][]int
		target [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}}},
		{[][]int{}, [][]int{}},
		{[][]int{{1}, {2}}, [][]int{{1, 2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
	}

	for _, tc := range tt {
		output := problem06.Transpose(tc.input)
		if OK := int2dSlicesEqual(output, tc.target); !OK {
			t.Errorf("Expected %v but got %v on input %v", tc.target, output, tc.input)
		}
	}
}

func int2dSlicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		for j, v2 := range v {
			if v2 != b[i][j] {
				return false
			}
		}
	}
	return true
}
