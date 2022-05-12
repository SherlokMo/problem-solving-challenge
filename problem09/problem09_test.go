package problem09_test

import (
	"halan-assignment/problem09"
	"testing"
)

func TestSumTree(t *testing.T) {

	tt := []struct {
		input  *problem09.Node
		target int
	}{
		{
			problem09.NewNode(1,
				&[]problem09.Node{
					*problem09.NewNode(2, &[]problem09.Node{}),
					*problem09.NewNode(3, &[]problem09.Node{}),
				}),
			6,
		},
		{
			problem09.NewNode(1, &[]problem09.Node{}),
			1,
		},
		{
			problem09.NewNode(1,
				&[]problem09.Node{
					*problem09.NewNode(2, &[]problem09.Node{}),
				}),
			3,
		},
	}

	for index, tc := range tt {
		output := problem09.SumTree(tc.input, 0)
		if output != tc.target {
			t.Errorf("Expected %v but got %v on test case %v", tc.target, output, index)
		}
	}
}
