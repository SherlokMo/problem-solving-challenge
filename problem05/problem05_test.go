package problem05_test

import (
	"halan-assignment/problem05"
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tt := []struct {
		input  []string
		target []string
	}{
		{[]string{"apples", "oranges", "flowers", "apples"}, []string{"oranges", "flowers"}},
		{[]string{"apples", "apples"}, []string{}},
		{[]string{"A", "A", "B", "C"}, []string{"B", "C"}},
		{[]string{""}, []string{""}},
		{[]string{}, []string{}},
	}

	for _, tc := range tt {
		output := problem05.Unique(tc.input)
		if equal := reflect.DeepEqual(output, tc.target); !equal {
			t.Errorf("Expected %v but got %v at input: %v", tc.target, output, tc.input)
		}
	}
}
