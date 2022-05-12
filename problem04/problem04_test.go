package problem04_test

import (
	"halan-assignment/problem04"
	"testing"
)

func TestComposition(t *testing.T) {
	tt := []struct {
		input  int
		target int
	}{
		{6, 49},
		{7, 64},
		{1, 4},
		{0, 1},
		{-1, 0},
		{-2, 1},
	}

	inc := func(x int) int {
		return x + 1
	}

	square := func(x int) int {
		return x * x
	}

	for _, tc := range tt {
		fn := problem04.Compose(square, inc)
		output := fn(tc.input)
		if output != tc.target {
			t.Errorf("Expected (%v) but got (%v). input: (%v)", tc.target, output, tc.input)
		}
	}
}
